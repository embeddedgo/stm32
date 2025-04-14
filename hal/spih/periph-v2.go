// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package spih

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/internal/apb"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CR1     mmio.R32[CR1]
	CR2     mmio.U32
	CFG1    mmio.R32[CFG1]
	CFG2    mmio.R32[CFG2]
	IER     mmio.R32[SR]
	SR      mmio.R32[SR]
	IFCR    mmio.R32[SR]
	_       uint32
	TXDR    mmio.U32
	_       [3]uint32
	RXDR    mmio.U32
	_       [3]uint32
	CRCPOLY mmio.U32
	TXCRC   mmio.U32
	RXCRC   mmio.U32
	UDRDR   mmio.U32
	CGFR    mmio.R32[CGFR]
}

// SPI returns the SPIn peripheral.
func SPI(n int) *Periph {
	var base uintptr
	switch n {
	case 1:
		base = mmap.SPI1_BASE
	case 2:
		base = mmap.SPI2_BASE
	case 3:
		base = mmap.SPI3_BASE
	case 4:
		base = mmap.SPI4_BASE
	case 5:
		base = mmap.SPI5_BASE
	case 6:
		base = mmap.SPI6_BASE
	default:
		panic("wrong SPI number")
	}
	return (*Periph)(unsafe.Pointer(base))
}

func num(p *Periph) int {
	switch uintptr(unsafe.Pointer(p)) {
	case mmap.SPI1_BASE:
		return 0
	case mmap.SPI2_BASE:
		return 1
	case mmap.SPI3_BASE:
		return 2
	case mmap.SPI4_BASE:
		return 3
	case mmap.SPI5_BASE:
		return 4
	case mmap.SPI6_BASE:
		return 5
	}
	return -1
}

// EnableClock enables clock for p.
// lp determines whether the clock remains on in low power (sleep) mode.
func (p *Periph) EnableClock(lp bool) {
	apb.EnableClock(unsafe.Pointer(p), lp)
}

// DisableClock disables clock for p.
func (p *Periph) DisableClock() {
	apb.DisableClock(unsafe.Pointer(p))
}

// Reset resets p.
func (p *Periph) Reset() {
	apb.Reset(unsafe.Pointer(p))
}
