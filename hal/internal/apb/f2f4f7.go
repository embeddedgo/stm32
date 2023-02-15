// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32f215 || stm32f407 || stm32f7x6

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
		panic("bad periph addr")
	}
	enr := (*mmio.R32[uint32])(unsafe.Pointer(rcc.RCC().APB1ENR.Addr() + rn*2))
	lpenr := (*mmio.R32[uint32])(unsafe.Pointer(rcc.RCC().APB1LPENR.Addr() + rn*2))
	mask := uint32(1) << bn
	internal.AtomicStoreBits(enr, mask, mask)
	internal.AtomicStoreBits(lpenr, mask, internal.BoolUint32(lp)<<bn)
	lpenr.Load() // RCC delay (workaround for silicon bugs)
}

func DisableClock(addr unsafe.Pointer) {
	rn, bn := rnbn(addr)
	if rn&1 != 0 || rn > 2 {
		panic("bad periph addr")
	}
	enr := (*mmio.R32[uint32])(unsafe.Pointer(rcc.RCC().APB1ENR.Addr() + rn*2))
	internal.AtomicStoreBits(enr, 1<<bn, 0)
	enr.Load() // RCC delay (workaround for silicon bugs)
}

func Reset(addr unsafe.Pointer) {
	rn, bn := rnbn(addr)
	if rn&1 != 0 || rn > 2 {
		panic("bad periph addr")
	}
	rstr := (*mmio.R32[uint32])(unsafe.Pointer(rcc.RCC().APB1RSTR.Addr() + rn*2))
	mask := uint32(1) << bn
	internal.AtomicStoreBits(rstr, mask, mask)
	internal.AtomicStoreBits(rstr, mask, 0)
	rstr.Load() // RCC delay (workaround for silicon bugs)
}
