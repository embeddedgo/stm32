// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows the basic usage of available LEDs and button.
package main

import (
	"time"

	"github.com/embeddedgo/stm32/devboard/mcudev-mini-pro/board/buttons"
	"github.com/embeddedgo/stm32/devboard/mcudev-mini-pro/board/leds"
)

func delay() {
	if buttons.User.Read() != 0 {
		time.Sleep(time.Second / 10)
	} else {
		time.Sleep(time.Second / 4)
	}
}

func main() {
	for {
		leds.User.SetOff()
		delay()
		leds.User.SetOn()
		delay()
	}
}
