// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package exti provides an interface to the External Interrupt/event
// controller.
//
// All functions in this package except Connect support concurrent access to
// different EXTI lines.
package exti

import "github.com/embeddedgo/stm32/hal/gpio"

// Lines is bitmask that represents EXTI input lines.
type Lines lines

const (
	L0 Lines = 1 << iota
	L1
	L2
	L3
	L4
	L5
	L6
	L7
	L8
	L9
	L10
	L11
	L12
	L13
	L14
	L15

	PVD Lines = 1 << 16 // Programmable Voltage Detector output.
)

// LineIndex returns bitmask for id EXTI line.
func LineIndex(id int) Lines {
	return Lines(1) << uint(id)
}

// Connect connects lines to corresponding pins of GPIO port. After reset lines
// 0..15 are connected to GPIO port A.
//
// Connect enables AFIO/SYSCFG clock before configuration and disables it before
// return. It can not be called concurently with any other function (including
// itself) that enables/disables AFIO/SYSCFG.
func (li Lines) Connect(port *gpio.Port) {
	if li >= L15<<1 {
		panic("exti: can not connect lines to GPIO port")
	}
	exticrEna()
	p := uint32(port.Num())
	var n int
	for li != 0 {
		if li&0xf != 0 {
			r := exticr(n)
			if li&1 != 0 {
				r.StoreBits(0x000f, p)
			}
			if li&2 != 0 {
				r.StoreBits(0x00f0, p<<4)
			}
			if li&4 != 0 {
				r.StoreBits(0x0f00, p<<8)
			}
			if li&8 != 0 {
				r.StoreBits(0xf000, p<<12)
			}
		}
		li = li >> 4
		n++
	}
	exticrDis()
}

// RiseTrigEnabled returns lines that have rising edge detection enabled.
func RiseTrigEnabled() Lines {
	return riseTrigEnabled()
}

// EnableRiseTrig enables rising edge detection for lines.
func (li Lines) EnableRiseTrig() {
	enableRiseTrig(li)
}

// DisableRiseTrig disables rising edge detection for lines.
func (li Lines) DisableRiseTrig() {
	disableRiseTrig(li)
}

// FallTrigEnabled returns lines that have falling edge detection enabled.
func FallTrigEnabled() Lines {
	return fallTrigEnabled()
}

// EnableFallTrig enables falling edge detection for lines.
func (li Lines) EnableFallTrig() {
	enableFallTrig(li)
}

// DisableFallTrig disables falling edge detection for lines.
func (li Lines) DisableFallTrig() {
	disableFallTrig(li)
}

// Trig allows to trigger an interrupt/event request by software. Interrupt
// pending flag on the line is set only when interrupt generation is enabled
// for this line.
func (li Lines) Trigger() {
	trigger(li)
}

// IRQEnabled returns lines that have IRQ generation enabled.
func IRQEnabled() Lines {
	return irqEnabled()
}

// EnableInt enables IRQ generation by lines.
func (li Lines) EnableIRQ() {
	enableIRQ(li)
}

// DisableInt disable IRQ generation by lines.
func (li Lines) DisableIRQ() {
	disableIRQ(li)
}

// EventEnabled returns lines that have event generation enabled.
func EventEnabled() Lines {
	return eventEnabled()
}

// EnableEvent enables event generation by lines.
func (li Lines) EnableEvent() {
	enableEvent(li)
}

// DisableEvent disable event generation by lines.
func (li Lines) DisableEvent() {
	disableEvent(li)
}

// Pending returns lines that have pending interrupt flag set.
func Pending() Lines {
	return pending()
}

// ClearPending clears pending interrupt flag for lines.
func (li Lines) ClearPending() {
	clearPending(li)
}
