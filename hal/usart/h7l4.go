// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3 || stm32l4x6

package usart

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type registers struct {
	cr1   mmio.R32[uint32]
	cr2   mmio.R32[uint32]
	cr3   mmio.R32[uint32]
	brr   mmio.R32[uint32]
	gtpr  mmio.R32[uint32]
	rtor  mmio.R32[uint32]
	rqr   mmio.R32[uint32]
	sr    mmio.R32[uint32]
	icr   mmio.R32[uint32]
	rdr   mmio.R32[uint32]
	tdr   mmio.R32[uint32]
	presc mmio.R32[uint32] // H7 only
}

const (
	// cr1
	ue     = 1 << 0
	re     = 1 << 2
	te     = 1 << 3
	ps     = 1 << 9
	pce    = 1 << 10
	m0     = 1 << 12
	over8  = 1 << 15
	m1     = 1 << 28
	fifoen = 1 << 29 // H7 only

	// cr2
	lbdie    = 1 << 6
	stop0b5  = 1 << 12
	stop2b   = 2 << 12
	stop1b5  = 3 << 12
	swap     = 1 << 15
	rxinv    = 1 << 16
	txinv    = 1 << 17
	datainv  = 1 << 18
	msbfirst = 1 << 19
	abren    = 1 << 20

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

	// rqr
	rxfrq = 1 << 3
)

func busForAddr(p *Periph) bus.Bus {
	switch uintptr(unsafe.Pointer(p)) {
	default:
		return bus.APB1
	case mmap.USART1_BASE:
		return bus.APB2
	}
}

func clear(p *Periph, ev Event, err Error) {
	p.icr.Store(uint32(ev) | uint32(err))
	if ev&RxNotEmpty != 0 {
		p.rqr.Store(rxfrq) // flush Rx data
	}
}

func tdr(p *Periph) *mmio.R32[uint32] {
	return &p.tdr
}

func rdr(p *Periph) *mmio.R32[uint32] {
	return &p.rdr
}

func altFunc(p *Periph, pin gpio.Pin) gpio.AltFunc {
	switch p {
	case USART1(), USART3():
		return gpio.AF7
	case USART2():
		if pin == gpio.PA().Pin(15) {
			return gpio.AF3
		}
		return gpio.AF7
	case UART4():
		if pin.Port().Num() == 0 && (pin.Num() == 12 || pin.Num() == 14) {
			return gpio.AF6 // H7 specific
		}
		return gpio.AF8
	case USART6():
		return gpio.AF7
	default: // UART4, UART5, LPUART1
		return gpio.AF8
	}
}
