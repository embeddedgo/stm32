// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package dma

import (
	"stm32/hal/raw/rcc"
)

func (p *DMA) enableClock(_ bool) {
	internal.AtomicSetBits(&rcc.RCC().AHBENR.U32, 1<<(rcc.DMA1ENn+d.num()))
	rcc.RCC().AHBENR.Load() // RCC delay (workaround for silicon bugs)
}

func (p *DMA) disableClock() {
	internal.AtomicClearBits(&rcc.RCC().AHBENR.U32, 1<<(rcc.DMA1ENn+d.num()))
}

func (p *DMA) reset() {}
