// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407 stm32f7x6

package apb

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
)

func EnableClock(addr unsafe.Pointer, lp bool) {
	rn, bn := rnbn(addr)
	if rn&1 != 0 || rn > 2 {
		panic("apb: bad periph addr")
	}
	enr := (*mmio.U32)(unsafe.Pointer(rcc.RCC().APB1ENR.U32.Addr() + rn*2))
	lpenr := (*mmio.U32)(unsafe.Pointer(rcc.RCC().APB1LPENR.U32.Addr() + rn*2))
	internal.AtomicSetBits(enr, 1<<bn)
	internal.AtomicSetBits(lpenr, internal.BoolUint32(lp)<<bn)
	lpenr.Load() // RCC delay (workaround for silicon bugs)
}

func DisableClock(addr unsafe.Pointer) {
	rn, bn := rnbn(addr)
	if rn&1 != 0 || rn > 2 {
		panic("apb: bad periph addr")
	}
	enr := (*mmio.U32)(unsafe.Pointer(rcc.RCC().APB1ENR.U32.Addr() + rn*2))
	internal.AtomicClearBits(enr, 1<<bn)
	enr.Load() // RCC delay (workaround for silicon bugs)
}

func Reset(addr unsafe.Pointer) {
	rn, bn := rnbn(addr)
	if rn&1 != 0 || rn > 2 {
		panic("apb: bad periph addr")
	}
	rstr := (*mmio.U32)(unsafe.Pointer(rcc.RCC().APB1RSTR.U32.Addr() + rn*2))
	internal.AtomicSetBits(rstr, 1<<bn)
	internal.AtomicClearBits(rstr, 1<<bn)
	rstr.Load() // RCC delay (workaround for silicon bugs)
}
