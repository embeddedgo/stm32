// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407 stm32f7x6

package exti

import (
	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/exti"
)

type lines uint32

func riseTrigEnabled() Lines {
	return Lines(exti.EXTI().RTSR.Load())
}

func (li Lines) enableRiseTrig() {
	internal.AtomicSetBits(&exti.EXTI().RTSR.U32, uint32(li))
}

func (li Lines) disableRiseTrig() {
	internal.AtomicClearBits(&exti.EXTI().RTSR.U32, uint32(li))
}

func fallTrigEnabled() Lines {
	return Lines(exti.EXTI().FTSR.Load())
}

func (li Lines) enableFallTrig() {
	internal.AtomicSetBits(&exti.EXTI().FTSR.U32, uint32(li))
}

func (li Lines) disableFallTrig() {
	internal.AtomicClearBits(&exti.EXTI().FTSR.U32, uint32(li))
}

func (li Lines) trigger() {
	exti.EXTI().SWIER.Store(exti.SWIER(li))
}

func irqEnabled() Lines {
	return Lines(exti.EXTI().IMR.Load())
}

func (li Lines) enableIRQ() {
	internal.AtomicSetBits(&exti.EXTI().IMR.U32, uint32(li))
}

func (li Lines) disableIRQ() {
	internal.AtomicClearBits(&exti.EXTI().IMR.U32, uint32(li))
}

func eventEnabled() Lines {
	return Lines(exti.EXTI().EMR.Load())
}

func (li Lines) enableEvent() {
	internal.AtomicSetBits(&exti.EXTI().EMR.U32, uint32(li))
}

func (li Lines) disableEvent() {
	internal.AtomicClearBits(&exti.EXTI().EMR.U32, uint32(li))
}

func pending() Lines {
	return Lines(exti.EXTI().PR.Load())
}

func (li Lines) clearPending() {
	exti.EXTI().PR.Store(exti.PR(li))
}
