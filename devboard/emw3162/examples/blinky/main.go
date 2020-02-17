// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows the basic usage of available LEDs.
package main

import (
	"time"

	"github.com/embeddedgo/stm32/devboard/emw3162/board/leds"
)

func main() {
	for {
		leds.Green.SetOff()
		leds.Red.SetOn()
		time.Sleep(time.Second / 4)

		leds.Red.SetOff()
		leds.Green.SetOn()
		time.Sleep(time.Second / 4)
	}
}
