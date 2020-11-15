// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407 stm32f412

package usart

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/gpio"
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

const (
	// cr1
	re    = 1 << 2
	te    = 1 << 3
	ps    = 1 << 9
	pce   = 1 << 10
	m0    = 1 << 12
	ue    = 1 << 13
	over8 = 1 << 15

	// cr2
	lbdie   = 1 << 6
	stop0b5 = 1 << 12
	stop2b  = 2 << 12
	stop1b5 = 3 << 12

	// cr3
	eie    = 1 << 0
	iren   = 1 << 1
	irlp   = 1 << 2
	hdsel  = 1 << 3
	nack   = 1 << 4
	scen   = 1 << 5
	dmar   = 1 << 6
	dmat   = 1 << 7
	rtse   = 1 << 8
	ctse   = 1 << 9
	ctsie  = 1 << 10
	onebit = 1 << 11

	// unsupported
	m1       = 0
	swap     = 0
	rxinv    = 0
	txinv    = 0
	datainv  = 0
	msbfirst = 0
	abren    = 0
)

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

func tdr(p *Periph) *mmio.U32 {
	return &p.dr
}

func rdr(p *Periph) *mmio.U32 {
	return &p.dr
}

func altFunc(p *Periph, pin gpio.Pin) gpio.AltFunc {
	switch p {
	case USART1(), USART2(), USART3():
		return gpio.AF7
	default: // UART4, UART5, USART6, UART7, UART8
		return gpio.AF8
	}
}
