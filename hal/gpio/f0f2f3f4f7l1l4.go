// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407 stm32f7x6 stm32l4x6

package gpio

import (
	"embedded/mmio"

	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
)

type registers struct {
	moder   mmio.U32
	otyper  mmio.U32
	ospeedr mmio.U32
	pupdr   mmio.U32
	idr     mmio.U16
	_       uint16
	odr     mmio.U16
	_       uint16
	bsrr    mmio.U32
	lckr    mmio.U32
	afr     [2]mmio.U32
}

const (
	// Mode
	out   = 1
	alt   = 2
	altIn = 2
	ana   = 3

	// Pull
	pullUp   = 1
	pullDown = 2

	// Driver
	openDrain = 1
)

func enableClock(p *Port, lp bool) {
	pnum := portnum(p)
	internal.AtomicSetBits(&enreg().U32, uint32(rcc.GPIOAEN<<pnum))
	if lp {
		lpenaclk(pnum)
	} else {
		lpdisclk(pnum)
	}
	enreg().Load() // RCC delay (workaround for silicon bugs).
}

func disableClock(p *Port) {
	internal.AtomicClearBits(&enreg().U32, uint32(rcc.GPIOAEN<<portnum(p)))
}

func reset(p *Port) {
	pnum := portnum(p)
	internal.AtomicSetBits(&rstreg().U32, uint32(rcc.GPIOARST<<pnum))
	internal.AtomicClearBits(&rstreg().U32, uint32(rcc.GPIOARST<<pnum))
}

func setup(p *Port, n int, cfg *Config) {
	pos := uint(n * 2)
	p.otyper.StoreBit(n, int(cfg.Driver))
	p.ospeedr.StoreBits(3<<pos, uint32(int(cfg.Speed)-veryLow)<<pos)
	p.pupdr.StoreBits(3<<pos, uint32(cfg.Pull)<<pos)
	p.moder.StoreBits(3<<pos, uint32(cfg.Mode)<<pos)
}

type AltFunc byte

const (
	AF0 AltFunc = iota
	AF1
	AF2
	AF3
	AF4
	AF5
	AF6
	AF7
	AF8
	AF9
	AF10
	AF11
	AF12
	AF13
	AF14
	AF15

	AF0_SYSTEM    = AF0  // System function.
	AF15_EVENTOUT = AF15 // Pulse generated by SEV instruction.
)

// SetPinAltFunc sets alternate function for n-th pin.
func (p *Port) SetPinAltFunc(n int, af AltFunc) {
	m := n >> 3 & 1
	o := uint(n) & 7 * 4
	p.afr[m].StoreBits(0xf<<o, uint32(af)<<o)
}

// SetAltFunc sets alternate function for pins.
func (p *Port) SetAltFunc(pins Pins, af AltFunc) {
	for n := 0; n < 16; n++ {
		if pins&(1<<uint(n)) != 0 {
			p.SetPinAltFunc(n, af)
		}
	}
}

// SetAltFunc sets alternate function for pin.
func (p Pin) SetAltFunc(af AltFunc) {
	p.Port().SetPinAltFunc(p.Index(), af)
}
