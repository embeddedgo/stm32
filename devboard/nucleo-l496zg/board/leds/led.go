// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leds

import (
	"github.com/embeddedgo/stm32/hal/gpio"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l496zg/board/system"
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
	cfg := &gpio.Config{Mode: gpio.Out, Speed: gpio.Low}

	gpio.PB().EnableClock(true)
	LD2.Pin().Setup(cfg)
	LD3.Pin().Setup(cfg)

	gpio.PC().EnableClock(true)
	LD1.Pin().Setup(cfg)
}
