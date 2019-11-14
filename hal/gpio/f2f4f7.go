// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407 stm32f7x6

package gpio

import (
	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
)

const (
	veryLow  = -1 // Not supported.
	low      = -1 //   2 MHz (CL = 50 pF, VDD > 2.7 V)
	_        = 0  //  25 MHz (CL = 50 pF, VDD > 2.7 V)
	high     = 1  //  50 MHz (CL = 40 pF, VDD > 2.7 V)
	veryHigh = 2  // 100 MHz (CL = 30 pF, VDD > 2.7 V)
)

func enreg() *rcc.RAHB1ENR   { return &rcc.RCC().AHB1ENR }
func rstreg() *rcc.RAHB1RSTR { return &rcc.RCC().AHB1RSTR }

func lpenaclk(pnum uint) {
	internal.AtomicSetBits(&rcc.RCC().AHB1LPENR.U32, uint32(rcc.GPIOALPEN<<pnum))
}

func lpdisclk(pnum uint) {
	internal.AtomicClearBits(&rcc.RCC().AHB1LPENR.U32, uint32(rcc.GPIOALPEN<<pnum))
}
