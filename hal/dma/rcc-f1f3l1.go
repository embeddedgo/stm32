// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package dma

import (
	"stm32/hal/raw/rcc"

	"github.com/embeddedgo/stm32/hal/internal"
)

func (p *DMA) enableClock(_ bool) {
	mask := uint32(1) << (rcc.DMA1ENn + d.num())
	internal.AtomicStoreBits(&rcc.RCC().AHBENR, mask, mask)
	rcc.RCC().AHBENR.Load() // RCC delay (workaround for silicon bugs)
}

func (p *DMA) disableClock() {
	internal.AtomicClearBits(&rcc.RCC().AHBENR, 1<<(rcc.DMA1ENn+d.num()))
}

func (p *DMA) reset() {}
