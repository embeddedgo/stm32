// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Generator implements a square wave generator using an STM32 timer.
// There are two outupts. PA8 (TIM1_CH1) and PA7 (TIM1_CH1N, negation of
// TIM1_CH1).
//
// This aplication can be controled using the USB serial console. It uses
// termfs.LightFS because of the  little of RAM in the STM32L476 has. The
// simple console provided doesn't support neither CR/LF conversion nor remote
// echo. Use options of your terminal emulator to provide local echo and
// convert CR and LF chars properly. The working example for Linux is:
//
//	picocom -b 115200 --echo --imap lfcrlf --omap crlf /dev/ttyACM0
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/rcc"
	"github.com/embeddedgo/stm32/p/tim"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/system"
)

func main() {
	pa := gpio.PA()
	pa.EnableClock(true)

	// Configure PA7 and PA8 as TIM1 outputs.
	const (
		neg = gpio.Pin7 // TIM1_CH1N
		pos = gpio.Pin8 // TIM1_CH1
	)
	pa.Setup(neg|pos, &gpio.Config{Mode: gpio.Alt, Speed: gpio.High})
	pa.SetAltFunc(neg|pos, gpio.AF1)

	RRC := rcc.RCC()
	RRC.APB2ENR.SetBits(rcc.TIM1EN)

	const (
		maxDuty = 10
		pwmMode = 6 // Mode 1
	)

	// Configure TIM1 to output PWM waveform on CH1N.
	pclk := bus.APB2.Clock()
	if RRC.CFGR.LoadBits(rcc.PPRE2) != rcc.PPRE2_DIV1 {
		pclk *= 2
	}
	t := tim.TIM1()
	t.ARR.Store(maxDuty - 1)
	t.CCMR1.Store(pwmMode<<tim.OC1Mn | tim.OC1PE)
	t.CCER.Store(tim.CC1E | tim.CC1NE)
	t.BDTR.Store(tim.MOE)
	t.EGR.Store(tim.UG)
	t.CR1.Store(tim.ARPE | tim.CEN)

	var (
		freq int64 = 1000 // Hz
		duty int64 = maxDuty / 2
	)
	r := bufio.NewReader(os.Stdin)
	for {
		t.CCR1.Store(uint32(duty))
		t.PSC.Store(uint32(pclk/(maxDuty*freq) - 1))
	again:
		fmt.Printf(
			"Freq: %d Hz, duty: %d/%d.\nNew freq and duty: ",
			freq, duty, maxDuty,
		)
		s, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			goto again
		}
		fields := strings.Fields(s)
		if len(fields) != 2 {
			fmt.Println("FREQ_HZ DUTY expected")
			goto again
		}
		freq, err = strconv.ParseInt(fields[0], 10, 64)
		if err != nil {
			fmt.Println("new freq:", err)
			goto again
		}
		if freq <= 0 {
			fmt.Println("new freq <= 0")
			goto again
		}
		if freq > pclk/maxDuty {
			fmt.Println("new freq >", pclk/maxDuty)
			goto again
		}
		duty, err = strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			fmt.Println("new duty:", err)
			goto again
		}
		if uint64(duty) > 10 {
			fmt.Println("new duty must be 0..10")
			goto again
		}
	}
}
