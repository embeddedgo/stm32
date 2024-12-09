// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package nocache

import (
	"embedded/mmio"
	"embedded/rtos"
	"runtime"
	"sync/atomic"
	"unsafe"
)

// AHBSRAM(1+2) in STM32H743
const (
	base   uintptr = 0x3000_0000
	logSiz         = 18 // log2(256 * 1024)
	end            = base + 1<<logSiz
)

var free = base

func init() {
	// TODO: use mpu package (see internal/cpu/cortexm/mpu)
	const (
		ENA   = 1 << 0
		B     = 1 << 16
		C     = 1 << 17
		S     = 1 << 18
		TEX1  = 1 << 19
		Arwrw = 3 << 24
	)
	const VALID = 1 << 4
	rbar := (*mmio.R32[uint32])(unsafe.Pointer(uintptr(0xE000ED90 + 3*4)))
	rasr := (*mmio.R32[uint32])(unsafe.Pointer(uintptr(0xE000ED90 + 4*4)))

	runtime.LockOSThread()
	pl, _ := rtos.SetPrivLevel(0)

	rbar.Store(uint32(base) | VALID | 7)               // tasker uses regions 0-3
	rasr.Store(ENA | (logSiz-1)<<1 | Arwrw | S | TEX1) // normal noncacheable
	//rasr.Store(ENA | (logSiz-1)<<1 | Arwrw | B) // normal noncacheable

	rtos.SetPrivLevel(pl)
	runtime.UnlockOSThread()

	runtime_memclrNoHeapPointers(unsafe.Pointer(base), 1<<logSiz)
}

func alloc(align, size uintptr) unsafe.Pointer {
	align--
	for {
		oldFree := atomic.LoadUintptr(&free)
		ptr := (oldFree + align) &^ align
		if ptr+size > end {
			panic("out of noncacheable RAM")
		}
		if atomic.CompareAndSwapUintptr(&free, oldFree, ptr+size) {
			return unsafe.Pointer(ptr)
		}
	}
}

//go:linkname runtime_memclrNoHeapPointers runtime.memclrNoHeapPointers
func runtime_memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)
