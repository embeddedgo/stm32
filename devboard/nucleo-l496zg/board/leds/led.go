// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leds

import (
	"github.com/embeddedgo/stm32/hal/gpio"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l496zg/board/init"
)

// Onboard LEDs
const (
	LD1 LED = 0x27 // PC7
	LD2 LED = 0x17 // PB7
	LD3 LED = 0x1E // PB14

	Green = LD1
	Blue  = LD2
	Red   = LD3

	User = LD3
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
	cfg := &gpio.Config{Mode: gpio.Out, Speed: gpio.Low}

	gpio.B().EnableClock(true)
	LD2.Pin().Setup(cfg)
	LD3.Pin().Setup(cfg)

	gpio.C().EnableClock(true)
	LD1.Pin().Setup(cfg)
}
