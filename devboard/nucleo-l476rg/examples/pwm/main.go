// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// PWM uses timer to implement PWM blinking of the onboard LED.
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
	out := leds.User.Pin() // PA5

	// Connect output pin (PA5) to TIM8_CH1N
	out.Setup(&gpio.Config{Mode: gpio.Alt, Speed: gpio.Low})
	out.SetAltFunc(gpio.AF3)

	RRC := rcc.RCC()

	// Enable TIM8.
	RRC.APB2ENR.SetBits(rcc.TIM8EN)

	const (
		pwmmax  = 1e4
		pwmfreq = 200 // Hz
		pwmmode = 6   // Mode 1
	)

	// Configure TIM8 to output PWM waveform on CH1N.
	pclk := bus.APB2.Clock()
	if RRC.CFGR.LoadBits(rcc.PPRE2) != rcc.PPRE2_DIV1 {
		pclk *= 2
	}
	t := tim.TIM8()
	t.PSC.Store(uint32(pclk/(pwmmax*pwmfreq) - 1))
	t.ARR.Store(pwmmax - 1)
	t.CCMR1.Store(pwmmode<<tim.OC1Mn | tim.OC1PE)
	t.CCER.Store(tim.CC1NE)
	t.BDTR.Store(tim.MOE)
	t.EGR.Store(tim.UG)
	t.CR1.Store(tim.ARPE | tim.CEN)

	// Slowly change the duty cycle.
	var (
		x    uint32
		back bool
	)
	for {
		t.CCR1.Store(x)
		if x <= 1 {
			x = 1
			back = false
		}
		if x >= pwmmax {
			x = pwmmax
			back = true

		}
		if back {
			x /= 2
		} else {
			x *= 2
		}
		time.Sleep(50 * time.Millisecond)
	}
}
