// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to handle the push-button with interrupts and with
// the software debouncing done right.
package main

import (
	"embedded/rtos"
	"time"

	"github.com/embeddedgo/stm32/hal/exti"
	"github.com/embeddedgo/stm32/hal/irq"

	"github.com/embeddedgo/stm32/devboard/nucleo-l496zg/board/buttons"
	"github.com/embeddedgo/stm32/devboard/nucleo-l496zg/board/leds"
)

func main() {
	pin := buttons.User.Pin() // PC13
	line := exti.Lines(1 << pin.Num())
	line.Connect(pin.Port())
	line.EnableRiseTrig()
	line.EnableFallTrig()

	irq.EXTI15_10.Enable(rtos.IntPrioLow, 0)

	for {
		leds.Red.SetOff()
		leds.Green.SetOn()

		waitBtn(0)
		waitBtn(1)

		leds.Green.SetOff()
		leds.Blue.SetOn()

		waitBtn(0)
		waitBtn(1)

		leds.Blue.SetOff()
		leds.Red.SetOn()

		waitBtn(0)
		waitBtn(1)
	}
}

var note rtos.Note

func waitBtn(state int) {
	line := exti.Lines(1 << buttons.User.Pin().Num())
	for {
		note.Clear()
		line.EnableIRQ()
		wait := time.Duration(-1)
		if buttons.User.Read() == state {
			wait = 50 * time.Millisecond // we want 50 ms of stable state
		}
		if !note.Sleep(wait) {
			line.DisableIRQ()
			return
		}
	}
}

//go:interrupthandler
func EXTI15_10_Handler() {
	p := exti.Pending() & (exti.L15<<1 - exti.L10)
	p.DisableIRQ()
	p.ClearPending()
	if pin := buttons.User.Pin(); p>>pin.Num()&1 != 0 {
		note.Wakeup()
	}
}
