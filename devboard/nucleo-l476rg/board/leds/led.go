// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leds

import (
	"github.com/embeddedgo/stm32/hal/gpio"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/init"
)

// Onboard LED
const (
	LD2 LED = 0x05 // PA5

	User = LD2
)

type LED uint8

func (d LED) prt() int  { return int(d) >> 4 }
func (d LED) pin() uint { return uint(d) & 15 }

func (d LED) SetOn() {
	gpio.P(d.prt()).SetPins(1 << d.pin())
}
func (d LED) SetOff() {
	gpio.P(d.prt()).ClearPins(1 << d.pin())
}
func (d LED) Set(on int) {
	gpio.P(d.prt()).StorePins(1<<d.pin(), gpio.Pins((on&1)<<d.pin()))
}
func (d LED) Get() int {
	return int(gpio.P(d.prt()).LoadOut()>>d.pin()) & 1
}
func (d LED) Pin() gpio.Pin {
	return gpio.P(d.prt()).Pin(int(d.pin()))
}

func init() {
	gpio.A().EnableClock(true)
	LD2.Pin().Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.Low})
}
