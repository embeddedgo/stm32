// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"time"

	"github.com/embeddedgo/stm32/hal/exti"
	"github.com/embeddedgo/stm32/hal/irq"

	"github.com/embeddedgo/stm32/devboard/f4-discovery/board/buttons"
)

type LatM struct {
	t0   time.Duration
	note rtos.Note
}

var latm LatM

func main() {
	pin := buttons.User.Pin() // PA0
	line := exti.Lines(1 << pin.Num())
	line.Connect(pin.Port())
	line.EnableRiseTrig()
	irq.EXTI0.Enable(rtos.IntPrioLow, 0)

	go waitBtn(&latm)
	//rtos.NewRawTask(func() { waitBtn(&latm) })

	for {
		time.Sleep(time.Hour)
	}
}

//go:nowritebarrierrec
func waitBtn(lm *LatM) {
	line := exti.Lines(1 << buttons.User.Pin().Num())
	for {
		lm.note.Clear()
		line.ClearPending()
		line.EnableIRQ()
		lm.note.Sleep(-1)
		t1 := rtos.Nanotime()
		println(t1 - lm.t0)
		time.Sleep(time.Second / 2)
	}
}

//go:interrupthandler
func EXTI0_Handler() {
	line := exti.Lines(1 << buttons.User.Pin().Num())
	line.DisableIRQ()
	line.ClearPending()
	latm.t0 = rtos.Nanotime()
	latm.note.Wakeup()
}
