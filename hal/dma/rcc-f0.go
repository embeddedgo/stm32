// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package dma

import (
	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
)

func (d *Controller) enableClock(_ bool) {
	internal.AtomicSetBits(&rcc.RCC().AHBENR.U32, uint32(1)<<rcc.DMAENn)
	rcc.RCC().AHBENR.Load() // RCC delay (workaround for silicon bugs)
}

func (d *Controller) disableClock() {
	internal.AtomicClearBits(&rcc.RCC().AHBENR.U32, uint32(1)<<rcc.DMAENn)
}

func (d *Controller) reset() {}
