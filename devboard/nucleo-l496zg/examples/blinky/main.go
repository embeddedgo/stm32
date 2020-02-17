// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows the basic usage of available LEDs and button.
package main

import (
	"time"

	"github.com/embeddedgo/stm32/devboard/nucleo-l496zg/board/buttons"
	"github.com/embeddedgo/stm32/devboard/nucleo-l496zg/board/leds"
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
		leds.Red.SetOff()
		leds.Green.SetOn()
		delay()

		leds.Green.SetOff()
		leds.Blue.SetOn()
		delay()

		leds.Blue.SetOff()
		leds.Red.SetOn()
		delay()
	}
}
