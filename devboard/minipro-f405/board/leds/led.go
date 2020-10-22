// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leds

import (
	"github.com/embeddedgo/stm32/hal/gpio"

	_ "github.com/embeddedgo/stm32/devboard/minipro-f405/board/init"
)

// Onboard LEDs
const (
	D1 LED = 0x08 // PA8

	User = D1
)

type LED uint8

func (d LED) prt() int  { return int(d) >> 4 }
func (d LED) pin() uint { return uint(d) & 15 }

func (d LED) SetOn() {
	gpio.P(d.prt()).ClearPins(1 << d.pin())
}
func (d LED) SetOff() {
	gpio.P(d.prt()).SetPins(1 << d.pin())
}
func (d LED) Set(on int) {
	gpio.P(d.prt()).StorePins(1<<d.pin(), gpio.Pins((on&1^1)<<d.pin()))
}
func (d LED) Get() int {
	return int(gpio.P(d.prt()).LoadOut()>>d.pin())&1 ^ 1
}
func (d LED) Pin() gpio.Pin {
	return gpio.P(d.prt()).Pin(int(d.pin()))
}

func init() {
	gpio.A().EnableClock(true)
	cfg := &gpio.Config{Mode: gpio.Out, Driver: gpio.OpenDrain, Speed: gpio.Low}
	D1.Pin().Setup(cfg)
}
