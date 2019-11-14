// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package board

import "github.com/embeddedgo/stm32/hal/gpio"

// LEDs
const (
	D1 LED = 0x10 // PB0
	D2 LED = 0x11 // PB1

	Green = D1
	Red   = D2

	UserLED = D1
)

type LED uint8

func (d LED) prt() int  { return int(d) >> 4 }
func (d LED) pin() uint { return uint(d) & 15 }

func (d LED) SetOn() {
	if d&1 == 0 {
		gpio.P(d.prt()).SetPins(1 << d.pin())
	} else {
		gpio.P(d.prt()).ClearPins(1 << d.pin())
	}
}
func (d LED) SetOff() {
	if d&1 == 0 {
		gpio.P(d.prt()).ClearPins(1 << d.pin())
	} else {
		gpio.P(d.prt()).SetPins(1 << d.pin())
	}
}
func (d LED) Set(on int) {
	gpio.P(d.prt()).StorePins(1<<d.pin(), gpio.Pins((on&1^int(d&1))<<d.pin()))
}
func (d LED) Get() int {
	return int(gpio.P(d.prt()).LoadOut()>>d.pin())&1 ^ int(d&1)
}
func (d LED) Pin() gpio.Pin {
	return gpio.P(d.prt()).Pin(int(d.pin()))
}
