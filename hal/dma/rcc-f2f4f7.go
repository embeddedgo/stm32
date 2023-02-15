// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32f215 || stm32f407 || stm32f7x6

package dma

import (
	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
)

func (d *Controller) enableClock(lp bool) {
	n := d.num()
	mask := rcc.AHB1ENR(1) << (rcc.DMA1ENn + n)
	internal.AtomicStoreBits(&rcc.RCC().AHB1ENR, mask, mask)

	n += rcc.DMA1LPENn
	internal.AtomicStoreBits(&rcc.RCC().AHB1LPENR, 1<<n, rcc.AHB1LPENR(internal.BoolUint32(lp)<<n))
	rcc.RCC().AHB1LPENR.Load() // RCC delay (workaround for silicon bugs)
}

func (d *Controller) disableClock() {
	internal.AtomicStoreBits(&rcc.RCC().AHB1ENR, 1<<(rcc.DMA1ENn+d.num()), 0)
}

func (d *Controller) reset() {
	mask := rcc.AHB1RSTR(1) << (rcc.DMA1RSTn + d.num())
	internal.AtomicStoreBits(&rcc.RCC().AHB1RSTR, mask, mask)
	internal.AtomicStoreBits(&rcc.RCC().AHB1RSTR, mask, 0)
}
