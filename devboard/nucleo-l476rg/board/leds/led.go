// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leds

import (
	"github.com/embeddedgo/stm32/hal/gpio"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/system"
)

// Onboard LED
const (
	LD2 LED = 0x05 // PA5

	User = LD2
)

type LED uint8

func prt(d LED) int  { return int(d) >> 4 }
func pin(d LED) uint { return uint(d) & 15 }

func (d LED) SetOn() {
	gpio.P(prt(d)).SetPins(1 << pin(d))
}
func (d LED) SetOff() {
	gpio.P(prt(d)).ClearPins(1 << pin(d))
}
func (d LED) Set(on int) {
	gpio.P(prt(d)).StorePins(1<<pin(d), gpio.Pins((on&1)<<pin(d)))
}
func (d LED) Get() int {
	return int(gpio.P(prt(d)).LoadOut()>>pin(d)) & 1
}
func (d LED) Toggle() {
	port := gpio.P(prt(d))
	mask := gpio.Pins(1 << pin(d))
	val := port.LoadOut() ^ mask
	port.StorePins(mask, val)
}
func (d LED) Pin() gpio.Pin {
	return gpio.P(prt(d)).Pin(int(pin(d)))
}

func init() {
	gpio.PA().EnableClock(true)
	LD2.Pin().Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.Low})
}
