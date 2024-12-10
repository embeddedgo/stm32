// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package gpio

import (
	"embedded/mmio"

	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
)

const (
	veryLow  = -1 // Not supported.
	low      = -1 //  12 MHz (CL = 50 pF, VDD > 2.7 V)
	_        = 0  //  60 MHz (CL = 50 pF, VDD > 2.7 V)
	high     = 1  // 110 MHz (CL = 30 pF, VDD > 2.7 V)
	veryHigh = 2  // 133 MHz (CL = 30 pF, VDD > 2.7 V)
)

func enreg() *mmio.R32[rcc.AHB4ENR]   { return &rcc.RCC().AHB4ENR }
func rstreg() *mmio.R32[rcc.AHB4RSTR] { return &rcc.RCC().AHB4RSTR }

func lpenaclk(pnum uint) {
	internal.ExclusiveStoreBits(&rcc.RCC().AHB4LPENR, rcc.GPIOALPEN<<pnum, rcc.GPIOALPEN<<pnum)
}

func lpdisclk(pnum uint) {
	internal.ExclusiveStoreBits(&rcc.RCC().AHB4LPENR, rcc.GPIOALPEN<<pnum, 0)
}
