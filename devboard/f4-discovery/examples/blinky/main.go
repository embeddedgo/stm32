// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows the basic usage of available LEDs and button.
package main

import (
	"time"

	"github.com/embeddedgo/stm32/devboard/f4-discovery/board/buttons"
	"github.com/embeddedgo/stm32/devboard/f4-discovery/board/leds"
)

func delay() {
	if buttons.User.Read() != 0 {
		time.Sleep(time.Second / 8)
	} else {
		time.Sleep(time.Second / 2)
	}
}

func main() {
	for {
		leds.Green.SetOff()
		leds.Orange.SetOn()
		delay()

		leds.Orange.SetOff()
		leds.Red.SetOn()
		delay()

		leds.Red.SetOff()
		leds.Blue.SetOn()
		delay()

		leds.Blue.SetOff()
		leds.Green.SetOn()
		delay()
	}
}
