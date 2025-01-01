// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package nocache

import (
	"embedded/arch/cortexm/mpu/mpu7"
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
	runtime.LockOSThread()
	pl, _ := rtos.SetPrivLevel(0)

	mpu7.SetRegion(base|mpu7.VALID|7, mpu7.ENA|mpu7.SIZE(logSiz)|mpu7.Arwrw|mpu7.TEX1)

	rtos.SetPrivLevel(pl)
	runtime.UnlockOSThread()

	clear((*[1 << logSiz]byte)(unsafe.Pointer(base))[:])
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
