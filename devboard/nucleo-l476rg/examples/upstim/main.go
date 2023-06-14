// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Upstime uses a timer to detect changes in the button state. The LED should
// light up mmediately after pressing the button and go out 5 seconds after
// button release.
//
// As the onboard button isn't connected to any timer input and two timer
// inputs are required to be connected to the button you have to make two
// external connections so the pins PC13 (button), PA0 (TIM8_ETR) and PB7
// (TIM8_BKIN) are connected togather.
//
// This algorithm was originaly designed to switch immediately to backup power
// source in the event of a mains failure and switch back when mains restores
// but with some dealy to avoid flapping.
package main

import (
	"time"

	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/rcc"
	"github.com/embeddedgo/stm32/p/tim"

	"github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/buttons"
)

func main() {
	pa := gpio.PA()
	pa.EnableClock(true)
	inpm := pa.Pin(0) // PA0
	out := pa.Pin(5)  // PA5 (onboard LED)

	pb := gpio.PB()
	pb.EnableClock(true)
	inpb := pb.Pin(7) // PB7

	// Connect input pins (PA0, PB7) to TIM8_ETR and TIM8_BKIN.
	altIn := &gpio.Config{Mode: gpio.AltIn}
	inpm.SetAltFunc(gpio.AF3) // TIM8_ETR
	inpm.Setup(altIn)
	inpb.SetAltFunc(gpio.AF3) // TIM8_BKIN
	inpb.Setup(altIn)

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

	// Configure timer speed and pulse length.
	t := tim.TIM8()
	t.PSC.Store(uint32(pre))
	t.ARR.Store(uint32(5*timFreq - 1)) // 5 second restore delay

	// Enable the one pulse mode.
	t.CR1.Store(tim.OPM)

	// OC1N (LED) high when MOE=0.
	t.CR2.Store(tim.OIS1N)

	// External trigger configuration
	smcr := tim.SMS_2    // TRGI resets and starts the counter
	smcr |= 7 << tim.TSn // TRGI = ETRF
	t.SMCR.Store(smcr)

	// Output compare mode configuration
	t.CCMR1.Store(1<<tim.OC1Mn | tim.OC1M_2) // retrigerrable one pulse mode 2

	// Configure timer outputs.
	ccer := tim.CC1E  // enabled CC1 (unused in this example)
	ccer |= tim.CC1NE // enable CC1N output (connected to the onboard LED)
	ccer |= tim.CC1NP // CC1NE = ^OC1REF (set LED on in backup mode)
	t.CCER.Store(ccer)

	// Enable BKIN
	t.OR2.Store(tim.BKINE)

	// Configure output break function and enable outputs.
	bdtr := tim.MOE  // enable outputs
	bdtr |= tim.AOE  // automaticaly recover at the update event
	bdtr |= tim.BKE  // enable break function used to switch to backup source
	bdtr |= tim.OSSI // use idle levels set in CR2 register
	t.BDTR.Store(bdtr)

	// Connect LED (PA5) to TIM8_CH1N. To avoid any transients do this only
	// after the timer is set up and ready.
	out.SetAltFunc(gpio.AF3)
	out.Setup(&gpio.Config{Mode: gpio.Alt, Speed: gpio.Low})

	for {
		print("BTN=", buttons.User.Read(), " MOE=", t.BDTR.LoadBits(tim.MOE)>>tim.MOEn, "\n")
		time.Sleep(time.Second / 2)
	}
}
