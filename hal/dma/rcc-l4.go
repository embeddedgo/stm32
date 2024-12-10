// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32l4x6

package dma

import (
	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
)

func (d *Controller) enableClock(_ bool) {
	mask := rcc.AHB1ENR(1) << (rcc.DMA1ENn + d.num())
	internal.ExclusiveStoreBits(&rcc.RCC().AHB1ENR, mask, mask)
	rcc.RCC().AHB1ENR.Load() // RCC delay (workaround for silicon bugs)
}

func (d *Controller) disableClock() {
	internal.ExclusiveStoreBits(&rcc.RCC().AHB1ENR, 1<<(rcc.DMA1ENn+d.num()), 0)
}

func (d *Controller) reset() {
	mask := rcc.AHB1RSTR(1) << (rcc.DMA1RSTn + d.num())
	internal.ExclusiveStoreBits(&rcc.RCC().AHB1RSTR, mask, mask)
	internal.ExclusiveStoreBits(&rcc.RCC().AHB1RSTR, mask, 0)
}
