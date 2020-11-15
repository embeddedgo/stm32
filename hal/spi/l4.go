// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32l4x6

package spi

import "github.com/embeddedgo/stm32/hal/gpio"

func altFunc(p *Periph, pin gpio.Pin) gpio.AltFunc {
	switch p {
	case SPI1():
		return gpio.AF5
	case SPI2():
		if pin == gpio.A().Pin(9) || pin == gpio.C().Pin(1) {
			return gpio.AF3
		}
		return gpio.AF5
	default: // SPI3
		return gpio.AF6
	}
}
