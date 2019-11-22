// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f407

package dma

import (
	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
)

func (d *Controller) enableClock(lp bool) {
	n := d.num()
	internal.AtomicSetBits(&rcc.RCC().AHB1ENR.U32, 1<<(uint(rcc.DMA1ENn)+n))

	n += uint(rcc.DMA1LPENn)
	internal.AtomicStoreBits(&rcc.RCC().AHB1LPENR.U32, 1<<n, internal.BoolUint32(lp)<<n)
	rcc.RCC().AHB1LPENR.Load() // RCC delay (workaround for silicon bugs)
}

func (d *Controller) disableClock() {
	mask := uint32(1) << (uint(rcc.DMA1ENn) + d.num())
	internal.AtomicClearBits(&rcc.RCC().AHB1ENR.U32, mask)
}

func (d *Controller) reset() {
	mask := uint32(1) << (uint(rcc.DMA1ENn) + d.num())
	internal.AtomicSetBits(&rcc.RCC().AHB1RSTR.U32, mask)
	internal.AtomicClearBits(&rcc.RCC().AHB1RSTR.U32, mask)
}
