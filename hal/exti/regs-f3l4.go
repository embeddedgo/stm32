// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f303 stm32l4x6

package exti

import (
	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/exti"
)

type lines uint64

func riseTrigEnabled() Lines {
	return Lines(exti.EXTI().RTSR1.Load()) |
		Lines(exti.EXTI().RTSR2.Load())<<32
}

func (li Lines) enableRiseTrig() {
	if m := uint32(li); m != 0 {
		internal.AtomicSetBits(&exti.EXTI().RTSR1.U32, m)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.AtomicSetBits(&exti.EXTI().RTSR2.U32, m)
	}
}

func (li Lines) disableRiseTrig() {
	if m := uint32(li); m != 0 {
		internal.AtomicClearBits(&exti.EXTI().RTSR1.U32, m)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.AtomicClearBits(&exti.EXTI().RTSR2.U32, m)
	}
}

func fallTrigEnabled() Lines {
	return Lines(exti.EXTI().FTSR1.Load()) |
		Lines(exti.EXTI().FTSR2.Load())<<32
}

func (li Lines) enableFallTrig() {
	if m := uint32(li); m != 0 {
		internal.AtomicSetBits(&exti.EXTI().FTSR1.U32, m)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.AtomicSetBits(&exti.EXTI().FTSR2.U32, m)
	}
}

func (li Lines) disableFallTrig() {
	if m := uint32(li); m != 0 {
		internal.AtomicClearBits(&exti.EXTI().FTSR1.U32, m)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.AtomicClearBits(&exti.EXTI().FTSR2.U32, m)
	}
}

func (li Lines) trigger() {
	if m := exti.SWIER1(li); m != 0 {
		exti.EXTI().SWIER1.Store(m)
	}
	if m := exti.SWIER2(li >> 32); m != 0 {
		exti.EXTI().SWIER2.Store(m)
	}
}

func irqEnabled() Lines {
	return Lines(exti.EXTI().IMR1.Load()) |
		Lines(exti.EXTI().IMR2.Load())<<32
}

func (li Lines) enableIRQ() {
	if m := uint32(li); m != 0 {
		internal.AtomicSetBits(&exti.EXTI().IMR1.U32, m)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.AtomicSetBits(&exti.EXTI().IMR2.U32, m)
	}
}

func (li Lines) disableIRQ() {
	if m := uint32(li); m != 0 {
		internal.AtomicClearBits(&exti.EXTI().IMR1.U32, m)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.AtomicClearBits(&exti.EXTI().IMR2.U32, m)
	}
}

func eventEnabled() Lines {
	return Lines(exti.EXTI().EMR1.Load()) |
		Lines(exti.EXTI().EMR2.Load())<<32
}

func (li Lines) enableEvent() {
	if m := uint32(li); m != 0 {
		internal.AtomicSetBits(&exti.EXTI().EMR1.U32, m)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.AtomicSetBits(&exti.EXTI().EMR2.U32, m)
	}
}

func (li Lines) disableEvent() {
	if m := uint32(li); m != 0 {
		internal.AtomicClearBits(&exti.EXTI().EMR1.U32, m)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.AtomicClearBits(&exti.EXTI().EMR2.U32, m)
	}
}

func pending() Lines {
	return Lines(exti.EXTI().PR1.Load()) |
		Lines(exti.EXTI().PR2.Load())<<32
}

func (li Lines) clearPending() {
	if m := exti.PR1(li); m != 0 {
		exti.EXTI().PR1.Store(m)
	}
	if m := exti.PR2(li >> 32); m != 0 {
		exti.EXTI().PR2.Store(m)
	}
}
