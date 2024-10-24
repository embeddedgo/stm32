// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32f215 || stm32f303 || stm32f407 || stm32f7x6 || stm32h7x3 || stm32l4x6

package dma

import (
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

func (d *Controller) num() uint {
	return uint((uintptr(unsafe.Pointer(d)) - mmap.DMA1_BASE) / 0x400)
}

func controller(n int) *Controller {
	n--
	if uint(n) > 1 {
		panic("bad DMA number")
	}
	return (*Controller)(unsafe.Pointer(mmap.DMA1_BASE + uintptr(n)*0x400))
}
