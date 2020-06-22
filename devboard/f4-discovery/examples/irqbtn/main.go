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

	"github.com/embeddedgo/stm32/devboard/f4-discovery/board/buttons"
	"github.com/embeddedgo/stm32/devboard/f4-discovery/board/leds"
)

func main() {
	pin := buttons.User.Pin() // PA0
	line := exti.Lines(1 << pin.Num())
	line.Connect(pin.Port())
	line.EnableRiseTrig()
	line.EnableFallTrig()

	irq.EXTI0.Enable(rtos.IntPrioLow, 0)

	for {
		leds.Green.SetOff()
		leds.Orange.SetOn()

		waitBtn(0)
		waitBtn(1)

		leds.Orange.SetOff()
		leds.Red.SetOn()

		waitBtn(0)
		waitBtn(1)

		leds.Red.SetOff()
		leds.Blue.SetOn()

		waitBtn(0)
		waitBtn(1)

		leds.Blue.SetOff()
		leds.Green.SetOn()

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
func EXTI0_Handler() {
	line := exti.Lines(1 << buttons.User.Pin().Num())
	line.DisableIRQ()
	line.ClearPending()
	note.Wakeup()
}
