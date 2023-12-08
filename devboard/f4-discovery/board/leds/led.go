// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leds

import (
	"github.com/embeddedgo/stm32/hal/gpio"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/system"
)

// Onboard LEDs
const (
	LD3 LED = 0x3D // PD13
	LD4 LED = 0x3C // PD12
	LD5 LED = 0x3E // PD14
	LD6 LED = 0x3F // PD15

	Orange = LD3
	Green  = LD4
	Red    = LD5
	Blue   = LD6

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
	d.Set(d.Get() ^ 1)
}
func (d LED) Pin() gpio.Pin {
	return gpio.P(prt(d)).Pin(int(pin(d)))
}

func init() {
	gpio.PD().EnableClock(true)
	cfg := &gpio.Config{Mode: gpio.Out, Speed: gpio.Low}
	LD3.Pin().Setup(cfg)
	LD4.Pin().Setup(cfg)
	LD5.Pin().Setup(cfg)
	LD6.Pin().Setup(cfg)
}
