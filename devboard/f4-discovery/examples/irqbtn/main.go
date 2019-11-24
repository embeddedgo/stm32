// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to handle the push-button with interrupts and with
// the software debouncing done right.
package main

import (
	"embedded/mmio"
	"embedded/rtos"

	"github.com/embeddedgo/stm32/devboard/f4-discovery/board"
	"github.com/embeddedgo/stm32/hal/exti"
	"github.com/embeddedgo/stm32/hal/irq"
)

func main() {
	board.Setup(true)

	pin := board.UserBtn.Pin() // PA0
	line := exti.Lines(1 << pin.Index())
	line.Connect(pin.Port())
	line.EnableRiseTrig()
	line.EnableFallTrig()

	irq.EXTI0.Enable(rtos.IntPrioLow)

	for {
		board.Green.SetOff()
		board.Orange.SetOn()

		waitBtn(0)
		waitBtn(1)

		board.Orange.SetOff()
		board.Red.SetOn()

		waitBtn(0)
		waitBtn(1)

		board.Red.SetOff()
		board.Blue.SetOn()

		waitBtn(0)
		waitBtn(1)

		board.Blue.SetOff()
		board.Green.SetOn()

		waitBtn(0)
		waitBtn(1)
	}
}

var note rtos.Note

func waitBtn(state int) {
	line := exti.Lines(1 << board.UserBtn.Pin().Index())
	for {
		note.Clear()
		mmio.MB() // ensure note.Clear() happens before IRQ in multicore system
		line.EnableIRQ()
		wait := int64(-1)
		if board.UserBtn.Read() == state {
			wait = 50e6 // we want 50 ms of stable state
		}
		if !note.Sleep(wait) {
			line.DisableIRQ()
			return
		}
	}
}

//go:interrupthandler
func EXTI0_Handler() {
	line := exti.Lines(1 << board.UserBtn.Pin().Index())
	line.DisableIRQ()
	line.ClearPending()
	note.Wakeup()
}
