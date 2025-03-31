// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package spih

import (
	"embedded/mmio"
	"math/bits"
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

// CFG1 (FIFO threshold 4 bytes)
const (
	wlen = uint64(DSIZE | FTHVL)
	w4   = 3 | 3<<FTHVLn
	w5   = 4 | 3<<FTHVLn
	w6   = 5 | 3<<FTHVLn
	w7   = 6 | 3<<FTHVLn
	w8   = 7 | 3<<FTHVLn
	w9   = 8 | 1<<FTHVLn
	w10  = 9 | 1<<FTHVLn
	w11  = 10 | 1<<FTHVLn
	w12  = 11 | 1<<FTHVLn
	w13  = 12 | 1<<FTHVLn
	w14  = 13 | 1<<FTHVLn
	w15  = 14 | 1<<FTHVLn
	w16  = 15 | 1<<FTHVLn
	w17  = 16 | 0<<FTHVLn
	w18  = 17 | 0<<FTHVLn
	w19  = 18 | 0<<FTHVLn
	w20  = 19 | 0<<FTHVLn
	w21  = 20 | 0<<FTHVLn
	w22  = 21 | 0<<FTHVLn
	w23  = 22 | 0<<FTHVLn
	w24  = 23 | 0<<FTHVLn
	w25  = 24 | 0<<FTHVLn
	w26  = 25 | 0<<FTHVLn
	w27  = 26 | 0<<FTHVLn
	w28  = 27 | 0<<FTHVLn
	w29  = 28 | 0<<FTHVLn
	w30  = 29 | 0<<FTHVLn
	w31  = 30 | 0<<FTHVLn
	w32  = 31 | 0<<FTHVLn

	clkDiv  = uint64(MBR)
	clkDiv4 = 1 << MBRn
)

// CFG2
const (
	sp = uint64(SP) << 32
	ti = uint64(1) << (SPn + 32)

	cpha = uint64(CPHA) << 32
	cpol = uint64(CPOL) << 32

	lsbfirst = uint64(LSBFRST) << 32

	dir = uint64(COMM) << 32
	tx  = uint64(1) << (COMMn + 3)
	rx  = uint64(2) << (COMMn + 3)
	hd  = uint64(3) << (COMMn + 3)
)

func kernHz(p *Periph) uint {
	return 100e6 // TODO: determine at runtime
}

const cfgMask = Config(wlen | clkDiv | sp | cpha | cpol | lsbfirst | dir)

func config(p *Periph) Config {
	cfg1 := p.CFG1.Load()
	cfg2 := p.CFG2.Load()
	return (Config(cfg1) | Config(cfg2<<32)) & cfgMask
}

func setConfig(p *Periph, cfg Config) {
	const defaults = MASTER | // master mode
		AFCNTR | // keep pins under controll when disabled
		SSOE // enable hardware Slave Select (CS) output
	cfg &= cfgMask
	p.CFG1.Store(CFG1(cfg))
	p.CFG2.Store(CFG2(cfg>>32) | defaults)
}

func configClkDiv(kernHz uint, baudrate int) Config {
	if baudrate <= 0 {
		panic("spi: baudrate<=0")
	}
	div := (kernHz + uint(baudrate) - 1) / uint(baudrate)
	log2div := uint(bits.Len(div) - 1)
	if div&^(1<<log2div) != 0 {
		log2div++
	}
	if log2div != 0 {
		log2div--
	}
	return Config(log2div << MBRn)
}

func baudrate(p *Periph) int {
	log2div := uint(p.CFG1.LoadBits(MBR)>>MBRn) + 1
	return int(kernHz(p) >> log2div)
}
