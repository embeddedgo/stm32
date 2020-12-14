// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package dma

import (
	"unsafe"

	"stm32/hal/raw/mmap"
)

func controller(n int) *Controller {
	if uint(n) != 1 {
		panic("bad DMA number")
	}
	return (*Controller)(unsafe.Pointer(mmap.DMA1_BASE)
}
