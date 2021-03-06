// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32l4x6

package dma

import (
	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
)

func (d *Controller) enableClock(_ bool) {
	internal.AtomicSetBits(&rcc.RCC().AHB1ENR.U32, 1<<(rcc.DMA1ENn+d.num()))
	rcc.RCC().AHB1ENR.Load() // RCC delay (workaround for silicon bugs)
}

func (d *Controller) disableClock() {
	internal.AtomicClearBits(&rcc.RCC().AHB1ENR.U32, 1<<(rcc.DMA1ENn+d.num()))
}

func (d *Controller) reset() {
	mask := uint32(1) << (rcc.DMA1RSTn + d.num())
	internal.AtomicSetBits(&rcc.RCC().AHB1RSTR.U32, mask)
	internal.AtomicClearBits(&rcc.RCC().AHB1RSTR.U32, mask)
}
