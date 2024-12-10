// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32f215 || stm32f407 || stm32f412 || stm32f7x6

package exti

import (
	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/exti"
)

type lines uint32

func riseTrigEnabled() Lines {
	return Lines(exti.EXTI().RTSR.Load())
}

func enableRiseTrig(li Lines) {
	internal.ExclusiveStoreBits(&exti.EXTI().RTSR, uint32(li), uint32(li))
}

func disableRiseTrig(li Lines) {
	internal.ExclusiveStoreBits(&exti.EXTI().RTSR, uint32(li), 0)
}

func fallTrigEnabled() Lines {
	return Lines(exti.EXTI().FTSR.Load())
}

func enableFallTrig(li Lines) {
	internal.ExclusiveStoreBits(&exti.EXTI().FTSR, uint32(li), uint32(li))
}

func disableFallTrig(li Lines) {
	internal.ExclusiveStoreBits(&exti.EXTI().FTSR, uint32(li), 0)
}

func trigger(li Lines) {
	exti.EXTI().SWIER.Store(uint32(li))
}

func irqEnabled() Lines {
	return Lines(exti.EXTI().IMR.Load())
}

func enableIRQ(li Lines) {
	internal.ExclusiveStoreBits(&exti.EXTI().IMR, uint32(li), uint32(li))
}

func disableIRQ(li Lines) {
	internal.ExclusiveStoreBits(&exti.EXTI().IMR, uint32(li), 0)
}

func eventEnabled() Lines {
	return Lines(exti.EXTI().EMR.Load())
}

func enableEvent(li Lines) {
	internal.ExclusiveStoreBits(&exti.EXTI().EMR, uint32(li), uint32(li))
}

func disableEvent(li Lines) {
	internal.ExclusiveStoreBits(&exti.EXTI().EMR, uint32(li), 0)
}

func pending() Lines {
	return Lines(exti.EXTI().PR.Load())
}

func clearPending(li Lines) {
	exti.EXTI().PR.Store(uint32(li))
}
