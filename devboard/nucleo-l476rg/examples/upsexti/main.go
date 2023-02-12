// Copyright 2019 Michal Derkacz. All rights reserved.
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
	"embedded/rtos"
	"sync/atomic"
	"time"

	"github.com/embeddedgo/stm32/hal/exti"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/irq"

	"github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/buttons"
	"github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/leds"
)

// Input pin states.
const (
	failure = 0 // low state on the input pin
	ok      = 1 // high state on the input pin
)

// System state
const (
	mains  = 0 // user LED is off
	backup = 1 // user LED is on
)

type input struct {
	line     exti.Lines
	pin      gpio.Pin
	detected int32
	note     rtos.Note
}

func (inp *input) detectFailure() {
	inp.line.DisableRiseTrig()
	inp.line.EnableFallTrig()
	inp.line.EnableIRQ()
	if inp.pin.Load() == failure {
		inp.line.Trig() // ensure detection of a possibly missed change
	}
}

func (inp *input) detectRecovery() {
	inp.line.DisableFallTrig()
	inp.line.EnableRiseTrig()
	inp.line.EnableIRQ()
	if inp.pin.Load() == ok {
		inp.line.Trig() // ensure detection of a possibly missed change
	}
}

func (inp *input) ISR() {
	trigFail := exti.FallTrigEnabled()&inp.line != 0
	detected := int32(ok)
	if trigFail {
		detected = failure
	}
	atomic.StoreInt32(&inp.detected, detected) // must be before setting LED
	if trigFail {
		leds.User.Set(backup) // respond fast
	}
	inp.note.Wakeup()
}

var inp input

func main() {
	// Setup the input struct and EXTI.
	inp.pin = buttons.User.Pin() // PC13
	inp.line = exti.Lines(1 << inp.pin.Num())
	inp.line.Connect(inp.pin.Port())
	irq.EXTI15_10.Enable(rtos.IntPrioMid, 0)

	// Start with the mains state.
	leds.User.Set(mains)
	inp.detectFailure()
	timeout := time.Duration(-1)

	for {
		// Wait for the input change.
		if !inp.note.Sleep(timeout) {
			// Note.Sleep can return false only in case of positive timeout so
			// the mains returned and there was no failure for timeout period.
			leds.User.Set(mains)

			// Check for the possible race with ISR.
			if atomic.LoadInt32(&inp.detected) == failure {
				leds.User.Set(backup) // fix it
			}

			// Wait for the mains failure.
			inp.note.Sleep(-1)
		}
		inp.note.Clear()

		// Reduce the interrupt rate to less than 1 Hz.
		time.Sleep(time.Second)

		if inp.detected == failure {
			// The system state switched to backup by ISR.
			inp.detectRecovery()
			timeout = -1
		} else {
			// Switch the system state to mains after some delay.
			inp.detectFailure()
			timeout = 4 * time.Second
		}
	}
}

//go:interrupthandler
func EXTI15_10_Handler() {
	p := exti.Pending() & (exti.L15<<1 - exti.L10)
	p.DisableIRQ()
	p.ClearPending()

	// This interrupt may be shared by multiple exti lines. Call individual ISRs
	// below. Note that their corresponding EXTI lines pending bits were
	// disabled and cleared (see above).
	if p&inp.line != 0 {
		inp.ISR()
	}
}
