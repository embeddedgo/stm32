// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32l4x6

package usart

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type registers struct {
	cr1  mmio.U32
	cr2  mmio.U32
	cr3  mmio.U32
	brr  mmio.U32
	gtpr mmio.U32
	rtor mmio.U32
	rqr  mmio.U32
	sr   mmio.U32
	icr  mmio.U32
	rdr  mmio.U32
	tdr  mmio.U32
}

func busForAddr(p *Periph) bus.Bus {
	switch uintptr(unsafe.Pointer(p)) {
	default:
		return bus.APB1
	case mmap.USART1_BASE:
		return bus.APB2
	}
}

const rxfrq = 1 << 3 // receive data flush request

func clear(p *Periph, ev Event, err Error) {
	p.icr.Store(uint32(ev) | uint32(err))
	if ev&RxNotEmpty != 0 {
		p.rqr.Store(rxfrq)
	}
}
