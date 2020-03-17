// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows the basic usage of available LEDs.
package main

import (
	"time"

	"github.com/embeddedgo/stm32/devboard/az3166/board/buttons"
	"github.com/embeddedgo/stm32/devboard/az3166/board/leds"
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
		leds.Blue.SetOff()
		leds.User.SetOn()
		delay()

		leds.User.SetOff()
		leds.Azure.SetOn()
		delay()

		leds.Azure.SetOff()
		leds.WiFi.SetOn()
		delay()

		leds.WiFi.SetOff()
		leds.Red.SetOn()
		delay()

		leds.Red.SetOff()
		leds.Green.SetOn()
		delay()

		leds.Green.SetOff()
		leds.Blue.SetOn()
		delay()
	}
}
