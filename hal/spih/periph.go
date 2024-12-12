// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package spih

import (
	"embedded/mmio"
	"unsafe"
)

type Periph struct {
	CR1     mmio.R32[CR1]
	CR2     mmio.R32[CR2]
	CFG1    mmio.R32[CFG1]
	CFG2    mmio.R32[CFG2]
	IER     mmio.R32[IER]
	SR      mmio.R32[SR]
	IFCR    mmio.R32[IFCR]
	_       uint32
	TXDR    mmio.R32[uint32]
	_       [3]uint32
	RXDR    mmio.R32[uint32]
	_       [3]uint32
	CRCPOLY mmio.R32[uint32]
	TXCRC   mmio.R32[uint32]
	RXCRC   mmio.R32[uint32]
	UDRDR   mmio.R32[uint32]
	CGFR    mmio.R32[CGFR]
}

// SPI returns the SPIn peripheral.
func SPI(n int) *Periph {
	var up uintptr
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

/*
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
*/

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

type Config uint64

// CFG1
const (
	mbrn = MBRn
)

// CFG2
const (
	cpha     = uint64(CPHAn) << 32
	cpol     = uint64(CPOLn) << 32
	lsbfirst = uint64(LSBFRST) << 32
	ssm      = uint64(SSM) << 32
	cmn      = COMMn + 32
)

const (
	CPHA0 = Config(0)    // Sample on leading edge.
	CPHA1 = Config(cpha) // Sample on trailing edge.
	CPOL0 = Config(0)    // Clock idle state is 0.
	CPOL1 = Config(cpol) // Clock idle state is 1.

	BR2   = Config(0)         // Baud rate = PCLK/2
	BR4   = Config(1 << mbrn) // Baud rate = PCLK/4.
	BR8   = Config(2 << mbrn) // Baud rate = PCLK/8.
	BR16  = Config(3 << mbrn) // Baud rate = PCLK/16.
	BR32  = Config(4 << mbrn) // Baud rate = PCLK/32.
	BR64  = Config(5 << mbrn) // Baud rate = PCLK/64.
	BR128 = Config(6 << mbrn) // Baud rate = PCLK/128.
	BR256 = Config(7 << mbrn) // Baud rate = PCLK/256.

	MSBF = Config(0)        // Most significant bit first.
	LSBF = Config(lsbfirst) // Least significant bit first.

	HardSS = Config(0)   // Hardware slave select.
	SoftSS = Config(ssm) // Software slave select (use ISSLow, ISSHigh)

	FullDuplex = Config(0)        // Full-duplex mode: SCK, MOSI, MISO
	TxOnly     = Config(1 << cmn) // Tx-only: SCK, MOSI(master), MISO(slave)
	RxOnly     = Config(2 << cmn) // Rx-only: SCK, MISO(master), MOSI(slave)
	HalfDuplex = Config(3 << cmn) // Half-duplex: SCK, MOSI(master), MISO(slave)
)

// Config returns the current configuration of the SPI peripheral.
func (p *Periph) Config() Config {
	cfg1 := p.cfg1.Load()
	cfg2 := p.cfg2.Load()
	return Config(cfg1) | Config(cfg2)<<32
}

// SetConfig configures p. If baudrate > 0 it replaces the BR bits in conf with
// the ones calculated from baudrate.
func (p *Periph) SetConfig(conf Config, baudrate int) {
	if baudrate > 0 {
		//conf = conf&^BR256 | p.BR(baudrate)
	}
	p.cfg1.Store(CFG1(conf))
	p.cfg2.Store(CFG2(conf >> 32))
}
