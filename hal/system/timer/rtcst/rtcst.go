// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build noos

// Package rtcst allows to setup the STM32 RTC type 2 peripheral as tickless
// system timer. Although the RTC type 2 is originally designed to provide a
// wall time only (in contrast to type 1 and type 3  binary mode) this package
// uses it as a source of monotonic time.

package rtcst

import (
	"embedded/rtos"
	"runtime"
	"time"
	_ "unsafe"

	"github.com/embeddedgo/stm32/hal/exti"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/p/pwr"
	"github.com/embeddedgo/stm32/p/rcc"
	"github.com/embeddedgo/stm32/p/rtc"
)

// RTC clock source
const (
	LSE int8 = iota
	LSI
	HSE
)

// Setup setups the RTC peripheral as a tickless system timer.
//
// divA*divS should be equal to the selected clock source frequency. Maximum
// allowed divA is 128, maximum alowed divS is 32768. Greater divA promotes
// lower RTC power usage but slows down the response to some configuration
// register changes. Greater divS increases subsecond accuracy.
//
// The RTC accuracy is equal to 1/divS second and in most cases you should
// probably favor better RTC accuracy and responsiveness than its power
// consumption. The RTC accuracy translates directly to the accuracy of the
// time.Now function and is equal to 30.5 us for maximum allowed divS.
func Setup(clkSrc int8, divA, divS int) {
	if uint(divA) > 128 {
		panic("bad divA")
	}
	if uint(divS) > 32768 {
		panic("bad divS")
	}
	const mask = rcc.LSEON | rcc.RTCSEL | rcc.RTCEN
	cfg := rcc.LSEON | rcc.RTCEN
	switch clkSrc {
	case LSE:
		cfg |= rcc.RTCSEL_LSE
	case LSI:
		cfg |= rcc.RTCSEL_LSI
	case HSE:
		cfg |= rcc.RTCSEL_HSE
	default:
		panic("bad clksrc")
	}
	PWR := pwr.PWR()
	RCC := rcc.RCC()
	RTC := rtc.RTC()

	RCC.PWREN().Set()
	RCC.PWREN().Load()
	PWR.DBP().Set() // enable write access to the backup domain.
	RCC.PWREN().Clear()

	if RCC.BDCR.LoadBits(mask) != cfg {
		// RTC not initialized
		RCC.BDCR.StoreBits(mask, cfg)

		RTC.WPR.Store(0xCA)
		RTC.WPR.Store(0x53)
		for RCC.LSERDY().Load() == 0 {
			runtime.Gosched()
		}
		RTC.ISR.Store(rtc.INIT)
		for RTC.INITF().Load() == 0 {
			runtime.Gosched()
		}
		prer := (divA-1)<<rtc.PREDIV_An | (divS-1)<<rtc.PREDIV_Sn
		RTC.PRER.Store(rtc.PRER(prer))

		// start from March 1 to simplify monoday function
		RTC.DR.Store(1<<rtc.WDUn | 0x0301)
	}

	// disable shadow registers because of the known problems (SMP, errata)
	RTC.CR.Store(rtc.BYPSHAD)
	RTC.ISR.Store(0)
	exti.RTCALR.EnableRiseTrig()
	exti.RTCALR.ClearPending()
	exti.RTCALR.EnableIRQ()
	irq.RTC_ALARM.Enable(rtos.IntPrioSysTimer, 0)

	rtos.SetSystemTimer(nanotime, setAlarm)
}

// Store stores a binary representation of t (without location) in a three
// consecutive RTC backup registers starting from BKPR[n].
func Store(t time.Time, n int) {
	sec := t.Unix()
	nsec := t.Nanosecond()
	bkpr := &rtc.RTC().BKPR
	bkpr[n].Store(rtc.BKPR(nsec))
	bkpr[n+1].Store(rtc.BKPR(sec))
	bkpr[n+2].Store(rtc.BKPR(sec >> 32))
}

// Load returns the local time corresponding to the binary representation of
// time saved previously by Store in the three consecutive RTC backup registers
// starting from BKPR[n].
func Load(n int) time.Time {
	bkpr := &rtc.RTC().BKPR
	nsec := int64(bkpr[n].Load())
	sec := int64(bkpr[n+1].Load()) | int64(bkpr[n+2].Load())<<32
	return time.Unix(sec, nsec)
}

func nanotime() int64 {
	RTC := rtc.RTC()
	tr := uint(RTC.TR.Load())
	dr := uint(RTC.DR.Load())
	ssr := uint(RTC.SSR.Load())
	for {
		tr1 := uint(RTC.TR.Load())
		if tr1 == tr {
			break
		}
		tr = tr1
		dr = uint(RTC.DR.Load())
		ssr = uint(RTC.SSR.Load())
	}
	sspre := uint(RTC.PREDIV_S().Load())
	y := dr>>rtc.YTn&15*10 + dr>>rtc.YUn&15
	m := dr>>rtc.MTn&1*10 + dr>>rtc.MUn&15
	d := dr>>rtc.DTn&3*10 + dr>>rtc.DUn&15
	h := tr>>rtc.HTn&3*10 + tr>>rtc.HUn&15
	mn := tr>>rtc.MNTn&7*10 + tr>>rtc.MNUn&15
	s := int64(tr>>rtc.STn&7*10 + tr>>rtc.SUn&15)
	ns := (1e9 * int64(sspre-ssr&0xFFFF)) / int64(sspre+1)
	return (int64((monoday(y, m, d)*24+h)*60+mn)*60+s)*1e9 + ns
}

func setAlarm(nanosec int64) {
	RTC := rtc.RTC()
	RTC.CR.Store(rtc.BYPSHAD)
	if nanosec < 0 {
		return
	}
	sspre := int(RTC.PREDIV_S().Load())
	sec := nanosec / 1e9
	ns := int(nanosec - sec*1e9)
	ss := sspre - int((int64(sspre+1)*int64(ns)+(1e9-1))/1e9)
	if ss < 0 {
		ss = sspre
		sec++
	}
	monoday := uint(sec / (24 * 60 * 60))
	s := uint(sec - int64(monoday)*(24*60*60))
	m := s / 60
	s -= m * 60
	h := m / 60
	m -= h * 60
	d := dayofmonth(monoday)

	dt := d / 10
	du := d - dt*10
	ht := h / 10
	hu := h - ht*10
	mt := m / 10
	mu := m - mt*10
	st := s / 10
	su := s - st*10
	alarm := dt<<rtc.ADTn | du<<rtc.ADUn |
		ht<<rtc.AHTn | hu<<rtc.AHUn |
		mt<<rtc.AMNTn | mu<<rtc.AMNUn |
		st<<rtc.ASTn | su<<rtc.ASUn

	RTC.ISR.Store(0) // clear ALRAF
	for RTC.ALRAWF().Load() == 0 {
	}
	RTC.ALRMAR.Store(rtc.ALRMR(alarm))
	RTC.ALRMASSR.Store(rtc.AMASKSS | rtc.ALRMSSR(ss))
	RTC.CR.Store(rtc.ALRAIE | rtc.ALRAE | rtc.BYPSHAD)
	if nanosec <= nanotime() && RTC.ALRAF().Load() == 0 {
		schedule() // avoid missed alarm
	}
}

func schedule()

//go:interrupthandler
func _RTC_ALARM_Handler() {
	exti.RTCALR.ClearPending()
	schedule()
}

//go:linkname _RTC_ALARM_Handler IRQ41_Handler
