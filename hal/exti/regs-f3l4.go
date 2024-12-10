// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32f303 || stm32l4x6

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

func enableRiseTrig(li Lines) {
	if m := uint32(li); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().RTSR1, m, m)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().RTSR2, m, m)
	}
}

func disableRiseTrig(li Lines) {
	if m := uint32(li); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().RTSR1, m, 0)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().RTSR2, m, 0)
	}
}

func fallTrigEnabled() Lines {
	return Lines(exti.EXTI().FTSR1.Load()) |
		Lines(exti.EXTI().FTSR2.Load())<<32
}

func enableFallTrig(li Lines) {
	if m := uint32(li); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().FTSR1, m, m)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().FTSR2, m, m)
	}
}

func disableFallTrig(li Lines) {
	if m := uint32(li); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().FTSR1, m, 0)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().FTSR2, m, 0)
	}
}

func trigger(li Lines) {
	if m := uint32(li); m != 0 {
		exti.EXTI().SWIER1.Store(m)
	}
	if m := uint32(li >> 32); m != 0 {
		exti.EXTI().SWIER2.Store(m)
	}
}

func irqEnabled() Lines {
	return Lines(exti.EXTI().IMR1.Load()) |
		Lines(exti.EXTI().IMR2.Load())<<32
}

func enableIRQ(li Lines) {
	if m := uint32(li); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().IMR1, m, m)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().IMR2, m, m)
	}
}

func disableIRQ(li Lines) {
	if m := uint32(li); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().IMR1, m, 0)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().IMR2, m, 0)
	}
}

func eventEnabled() Lines {
	return Lines(exti.EXTI().EMR1.Load() | exti.EXTI().EMR2.Load()<<32)
}

func enableEvent(li Lines) {
	if m := uint32(li); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().EMR1, m, m)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().EMR2, m, m)
	}
}

func disableEvent(li Lines) {
	if m := uint32(li); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().EMR1, m, 0)
	}
	if m := uint32(li >> 32); m != 0 {
		internal.ExclusiveStoreBits(&exti.EXTI().EMR2, m, 0)
	}
}

func pending() Lines {
	return Lines(exti.EXTI().PR1.Load() | exti.EXTI().PR2.Load()<<32)
}

func clearPending(li Lines) {
	if m := uint32(li); m != 0 {
		exti.EXTI().PR1.Store(m)
	}
	if m := uint32(li >> 32); m != 0 {
		exti.EXTI().PR2.Store(m)
	}
}
