// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package gpio

import (
	"mmio"

	"stm32/hal/raw/rcc"

	"github.com/embeddedgo/stm32/hal/internal"
)

type registers struct {
	cr   [2]mmio.U32
	idr  mmio.U16
	_    uint16
	odr  mmio.U16
	_    uint16
	bsrr mmio.U32
	brr  mmio.U32
	lckr mmio.U32
}

const (
	// Mode
	out   = 1
	alt   = 8 + 1
	altIn = 0
	ana   = 2

	// Pull
	pullUp   = 8 + 4 + 1
	pullDown = 8 + 4 + 0

	// Driver
	openDrain = 4

	// Speed
	veryLow  = 1
	low      = 1
	high     = 2
	veryHigh = 2
)

func enableClock(p *Port, _ bool) {
	internal.AtomicSetBits(&rcc.RCC().APB2ENR.U32, uint32(rcc.IOPAEN<<portnum(p)))
	rcc.RCC().APB2ENR.Load() // RCC delay (workaround for silicon bugs).
}

func disableClock(p *Port) {
	internal.AtomicClearBits(&rcc.RCC().APB2ENR.U32, uint32(rcc.IOPAEN<<portnum(p)))
}

func reset(p *Port) {
	pnum := portnum(p)
	internal.AtomicSetBits(&rcc.RCC().APB2RSTR.U32, uint32(rcc.IOPARST<<pnum))
	internal.AtomicClearBits(&rcc.RCC().APB2RSTR.U32, uint32(rcc.IOPARST<<pnum))
}

func setup(p *Port, n int, cfg *Config) {
	cr := &p.cr[n>>3]
	pos := uint(n & 7 * 4)
	sel := uint32(0xf) << pos
	switch {
	case cfg.Mode == 0: // In, AltIn.
		cm := uint32(cfg.Pull)&(8+4) ^ 4
		cr.StoreBits(sel, cm<<pos)
		p.StorePins(Pin0<<uint(n), Pins(cfg.Pull)<<uint(n))
	case cfg.Mode&1 != 0: // Out, Alt.
		cm := uint32(cfg.Mode) & 8
		cm |= uint32(cfg.Driver)
		cm |= uint32(cfg.Speed) + 1
		cr.StoreBits(sel, cm<<pos)
	default: // Ana.
		cr.ClearBits(sel)
	}
}
