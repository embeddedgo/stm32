// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to handle the push-button with interrupts and with
// the software debouncing done right.
package main

import (
	"embedded/rtos"

	"github.com/embeddedgo/stm32/devboard/nucleo-l496zg/board"
	"github.com/embeddedgo/stm32/hal/exti"
	"github.com/embeddedgo/stm32/hal/irq"
)

func main() {
	board.Setup(true)

	pin := board.UserBtn.Pin() // PC13
	line := exti.Lines(1 << pin.Index())
	line.Connect(pin.Port())
	line.EnableRiseTrig()
	line.EnableFallTrig()

	irq.EXTI15_10.Enable(rtos.IntPrioLow)

	for {
		board.Red.SetOff()
		board.Green.SetOn()

		waitBtn(0)
		waitBtn(1)

		board.Green.SetOff()
		board.Blue.SetOn()

		waitBtn(0)
		waitBtn(1)

		board.Blue.SetOff()
		board.Red.SetOn()

		waitBtn(0)
		waitBtn(1)
	}
}

var note rtos.Note

func waitBtn(state int) {
	line := exti.Lines(1 << board.UserBtn.Pin().Index())
	for {
		note.Clear()
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
func EXTI15_10_Handler() {
	p := exti.Pending() & (exti.L15<<1 - exti.L10)
	p.DisableIRQ()
	p.ClearPending()
	if pin := board.UserBtn.Pin(); p>>pin.Index()&1 != 0 {
		note.Wakeup()
	}
}
