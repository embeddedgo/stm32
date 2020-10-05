// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f407

package usart

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type registers struct {
	sr   mmio.U32
	dr   mmio.U32
	brr  mmio.U32
	cr1  mmio.U32
	cr2  mmio.U32
	cr3  mmio.U32
	gtpr mmio.U32
}

func busForAddr(p *Periph) bus.Bus {
	switch uintptr(unsafe.Pointer(p)) {
	default:
		return bus.APB2
	case mmap.USART2_BASE, mmap.USART3_BASE:
		return bus.APB1
	}
}

func clear(p *Periph, ev Event, _ Error) {
	if ev != 0 {
		p.sr.Store(^uint32(ev))
	}
}
