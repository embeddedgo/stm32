// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spi

import "github.com/embeddedgo/stm32/hal/gpio"

type Signal uint8

const (
	SCK Signal = iota
	MOSI
	MISO
	NSS
)

// UsePinMaster is a helper function that can be used to configure GPIO pins as
// required by SPI master device. Only certain pins can be used (see datasheet).
func (d *Driver) UsePinMaster(pin gpio.Pin, sig Signal) {
	var cfg gpio.Config
	if sig == MISO {
		cfg.Mode = gpio.AltIn
	} else {
		if sig == NSS {
			cfg.Pull = gpio.PullUp
		}
		cfg.Mode = gpio.Alt
		cfg.Speed = gpio.VeryHigh
	}
	pin.Setup(&cfg)
	pin.SetAltFunc(altFunc(d.p, pin))
}

// UsePinSlave is a helper function that can be used to configure GPIO pins as
// required by SPI slave device. Only certain pins can be used (see datasheet).
func (d *Driver) UsePinSlave(pin gpio.Pin, sig Signal) {
	var cfg gpio.Config
	if sig == MISO {
		cfg.Mode = gpio.Alt
		cfg.Speed = gpio.VeryHigh
	} else {
		cfg.Mode = gpio.AltIn
	}
	pin.Setup(&cfg)
	pin.SetAltFunc(altFunc(d.p, pin))
}
