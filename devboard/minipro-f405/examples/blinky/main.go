// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows the basic usage of available LEDs and button.
package main

import (
	"time"

	"github.com/embeddedgo/stm32/devboard/minipro-f405/board/buttons"
	"github.com/embeddedgo/stm32/devboard/minipro-f405/board/leds"
)

func delay() {
	if buttons.User.Read() != 0 {
		time.Sleep(time.Second / 7)
	} else {
		time.Sleep(time.Second / 2)
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
