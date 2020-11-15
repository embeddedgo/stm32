// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407 stm32f412

package spi

import "github.com/embeddedgo/stm32/hal/gpio"

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
