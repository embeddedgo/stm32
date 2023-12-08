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
		pwmmax  = 10
		pwmmode = 6 // Mode 1
	)

	var pwmfreq int64 = 1000 // Hz

	// Configure TIM1 to output PWM waveform on CH1N.
	pclk := bus.APB2.Clock()
	if RRC.CFGR.LoadBits(rcc.PPRE2) != rcc.PPRE2_DIV1 {
		pclk *= 2
	}
	t := tim.TIM1()
	t.ARR.Store(pwmmax - 1)
	t.CCMR1.Store(pwmmode<<tim.OC1Mn | tim.OC1PE)
	t.CCER.Store(tim.CC1E | tim.CC1NE)
	t.BDTR.Store(tim.MOE)
	t.EGR.Store(tim.UG)
	t.CR1.Store(tim.ARPE | tim.CEN)

	t.CCR1.Store(pwmmax / 2)

	r := bufio.NewReader(os.Stdin)
	for {
		t.PSC.Store(uint32(pclk/(pwmmax*pwmfreq) - 1))
	again:
		fmt.Printf("Freq: %d Hz. New freq [Hz]: ", pwmfreq)
		s, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			goto again
		}
		pwmfreq, err = strconv.ParseInt(s[:len(s)-1], 10, 64)
		if err != nil {
			fmt.Println("new freq:", err)
			goto again
		}
		if pwmfreq <= 0 {
			fmt.Println("new freq <= 0")
			goto again
		}
		if pwmfreq > pclk/pwmmax {
			fmt.Println("new freq >", pclk/pwmmax)
			goto again
		}
	}
}
