// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Blinky2 shows the basic usage of the available LED and button. It differs
// from blinky in that it uses time.Timer instead of more obvious time.Sleep.

package main

import (
	"time"

	"github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/buttons"
	"github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/leds"
)

func delay() time.Duration {
	if buttons.User.Read() != 0 {
		return time.Second / 8
	}
	return time.Second / 2
}

func main() {
	tim := time.NewTimer(delay())
	for range tim.C {
		leds.User.Toggle()
		tim.Reset(delay())
	}
}
