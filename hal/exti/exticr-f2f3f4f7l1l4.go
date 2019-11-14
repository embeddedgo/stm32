// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f303 stm32f407 stm32f7x6 stm32l4x6

package exti

import (
	"embedded/mmio"

	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
	"github.com/embeddedgo/stm32/p/syscfg"
)

func exticr(n int) *mmio.U32 {
	return (*mmio.U32)(&syscfg.SYSCFG().EXTICR[n].U32)
}

func exticrEna() {
	internal.AtomicSetBits(&rcc.RCC().APB2ENR.U32, uint32(rcc.SYSCFGEN))
	rcc.RCC().APB2ENR.Load()
}

func exticrDis() {
	internal.AtomicClearBits(&rcc.RCC().APB2ENR.U32, uint32(rcc.SYSCFGEN))
}
