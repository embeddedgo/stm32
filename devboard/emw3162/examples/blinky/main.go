// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows the basic usage of available LEDs.
package main

import (
	"github.com/embeddedgo/stm32/devboard/emw3162/board"
	"github.com/embeddedgo/x/time"
)

func main() {
	board.Setup(true)
	for {
		board.Green.SetOff()
		board.Red.SetOn()
		time.Sleep(time.Second / 4)

		board.Red.SetOff()
		board.Green.SetOn()
		time.Sleep(time.Second / 4)
	}
}
