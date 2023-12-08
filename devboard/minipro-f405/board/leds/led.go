// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leds

import (
	"github.com/embeddedgo/stm32/hal/gpio"

	_ "github.com/embeddedgo/stm32/devboard/minipro-f405/board/system"
)

// Onboard LEDs
const (
	D1 LED = 0x08 // PA8

	User = D1
)

type LED uint8

func prt(d LED) int  { return int(d) >> 4 }
func pin(d LED) uint { return uint(d) & 15 }

func (d LED) SetOn() {
	gpio.P(prt(d)).ClearPins(1 << pin(d))
}
func (d LED) SetOff() {
	gpio.P(prt(d)).SetPins(1 << pin(d))
}
func (d LED) Set(on int) {
	gpio.P(prt(d)).StorePins(1<<pin(d), gpio.Pins((on&1^1)<<pin(d)))
}
func (d LED) Get() int {
	return int(gpio.P(prt(d)).LoadOut()>>pin(d))&1 ^ 1
}
func (d LED) Toggle() {
	d.Set(d.Get() ^ 1)
}
func (d LED) Pin() gpio.Pin {
	return gpio.P(prt(d)).Pin(int(pin(d)))
}

func init() {
	gpio.PA().EnableClock(true)
	cfg := &gpio.Config{Mode: gpio.Out, Driver: gpio.OpenDrain, Speed: gpio.Low}
	D1.SetOff()
	D1.Pin().Setup(cfg)
}
