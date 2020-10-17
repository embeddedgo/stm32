// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32h7x3

package gpio

import (
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

func enreg() *rcc.RAHB4ENR   { return &rcc.RCC().AHB4ENR }
func rstreg() *rcc.RAHB4RSTR { return &rcc.RCC().AHB4RSTR }

func lpenaclk(pnum uint) {
	internal.AtomicSetBits(&rcc.RCC().AHB4LPENR.U32, uint32(rcc.GPIOALPEN<<pnum))
}

func lpdisclk(pnum uint) {
	internal.AtomicClearBits(&rcc.RCC().AHB4LPENR.U32, uint32(rcc.GPIOALPEN<<pnum))
}
