// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Upsexti uses EXTI to detect changes in the button state. The LED should light
// up mmediately after pressing the button and go out 5 seconds after button
// release.
//
// This algorithm was originaly designed to switch fast to backup power source
// in the event of a mains failure and switch back when mains restores but with
// some dealy to avoid flapping.
package main

import (
	"time"

	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/rcc"
	"github.com/embeddedgo/stm32/p/tim"

	"github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/leds"
)

func main() {
	inp := gpio.PA().Pin(0) // PA0
	trg := gpio.PA().Pin(1) // PA1
	out := leds.User.Pin()  // PA5

	// Connect output pin (PA5) to TIM8_CH1N
	out.Setup(&gpio.Config{Mode: gpio.Alt, Speed: gpio.Low})
	out.SetAltFunc(gpio.AF3)

	trg.Clear()
	trg.Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.Low})

	// Connect input pin (PA0) to TIM8_ETR
	inp.Setup(&gpio.Config{Mode: gpio.AltIn})
	inp.SetAltFunc(gpio.AF3)

	RRC := rcc.RCC()
	RRC.APB2ENR.SetBits(rcc.TIM8EN)

	pclk := bus.APB2.Clock()
	if RRC.CFGR.LoadBits(rcc.PPRE2) != rcc.PPRE2_DIV1 {
		pclk *= 2
	}

	timFreq := int64(1e4) // Hz
	pre := pclk/timFreq - 1
	if pre > 0xffff {
		panic("timer frequency too low")
	}

	t := tim.TIM8()
	t.PSC.Store(uint32(pre))
	t.ARR.Store(uint32(2*timFreq - 1))

	// Configure CH1 as output in the retriggerable one pulse mode and CH2
	// as input that triggers on falling edge.
	t.CCMR1.Store(1<<tim.OC1M_2n | tim.OC1PE)
	t.CCER.Store(tim.CC1NE)
	t.SMCR.Store(1<<tim.SMS_2n | 7<<tim.TSn)

	t.BDTR.Store(tim.MOE)
	t.CR1.Store(tim.OPM)
	t.EGR.Store(tim.UG) // trigger update event to update registers

	for {
		trg.Clear()
		println("trg=0 sleep5")
		time.Sleep(5 * time.Second)
		println("trg=1 sleep5")
		trg.Set()
		time.Sleep(5 * time.Second)
	}
}
