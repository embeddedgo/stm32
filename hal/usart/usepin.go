// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usart

import "github.com/embeddedgo/stm32/hal/gpio"

type Signal uint8

const (
	RTSn Signal = iota
	TXD
	CTSn
	RXD
)

// UsePin is a helper function that can be used to configure GPIO pins as
// required by USART peripheral. Only some pins can be used (see datasheet).
func (d *Driver) UsePin(pin gpio.Pin, sig Signal) {
	var cfg gpio.Config
	if sig <= TXD {
		cfg.Mode = gpio.Alt
		cfg.Speed = gpio.VeryHigh
	} else {
		cfg.Mode = gpio.AltIn
	}
	pin.Setup(&cfg)
	pin.SetAltFunc(altFunc(d.p))
}
