// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package apb

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
)

func EnableClock(addr unsafe.Pointer, lp bool) {
	rn, bn := rnbn(addr)
	if rn > 2 {
		panic("bad periph addr")
	}
	RCC := rcc.RCC()
	enr := (*mmio.R32[uint32])(unsafe.Pointer(RCC.APB1LENR.Addr() + rn*4))
	lpenr := (*mmio.R32[uint32])(unsafe.Pointer(RCC.APB1LLPENR.Addr() + rn*4))
	mask := uint32(1) << bn
	internal.AtomicStoreBits(enr, mask, mask)
	internal.AtomicStoreBits(lpenr, mask, internal.BoolUint32(lp)<<bn)
	lpenr.Load() // RCC delay (workaround for silicon bugs)
}

func DisableClock(addr unsafe.Pointer) {
	rn, bn := rnbn(addr)
	if rn > 2 {
		panic("bad periph addr")
	}
	enr := (*mmio.R32[uint32])(unsafe.Pointer(rcc.RCC().APB1LENR.Addr() + rn*4))
	internal.AtomicStoreBits(enr, 1<<bn, 0)
	enr.Load() // RCC delay (workaround for silicon bugs)
}

func Reset(addr unsafe.Pointer) {
	rn, bn := rnbn(addr)
	if rn > 2 {
		panic("bad periph addr")
	}
	rstr := (*mmio.R32[uint32])(unsafe.Pointer(rcc.RCC().APB1LRSTR.Addr() + rn*4))
	mask := uint32(1) << bn
	internal.AtomicStoreBits(rstr, mask, mask)
	internal.AtomicStoreBits(rstr, mask, 0)
	rstr.Load() // RCC delay (workaround for silicon bugs)
}
