// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"github.com/embeddedgo/stm32/devboard/devebox-h743/board/leds"
)

//go:noinline
func delay(sec float64) {
	time.Sleep(time.Duration(sec * float64(time.Second)))
}

func main() {
	for {
		leds.User.SetOff()
		delay(.9)
		leds.User.SetOn()
		delay(.1)
	}
}
