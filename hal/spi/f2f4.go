// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407 stm32f412

package spi

import (
	"unsafe"

	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

func busForAddr(p *Periph) bus.Bus {
	switch uintptr(unsafe.Pointer(p)) {
	default:
		return bus.APB2
	case mmap.SPI2_BASE, mmap.SPI3_BASE:
		return bus.APB1
	}
}

func altFunc(p *Periph, pin gpio.Pin) gpio.AltFunc {
	switch p {
	case SPI3():
		if pin.Port() == gpio.D() {
			return gpio.AF5
		}
		return gpio.AF6
	default: // SPI1, SPI2, SPI4, SPI5, SPI6
		return gpio.AF5
	}
}
