// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows the basic usage of available LEDs and button.
package main

import (
	"github.com/embeddedgo/stm32/devboard/f4-discovery/board"
	"github.com/embeddedgo/x/time"
)

func delay() {
	if board.UserBtn.Read() != 0 {
		time.Sleep(time.Second / 8)
	} else {
		time.Sleep(time.Second / 2)
	}
}

func main() {
	board.Setup(true)
	for {
		board.Green.SetOff()
		board.Orange.SetOn()
		delay()

		board.Orange.SetOff()
		board.Red.SetOn()
		delay()

		board.Red.SetOff()
		board.Blue.SetOn()
		delay()

		board.Blue.SetOff()
		board.Green.SetOn()
		delay()
	}
}
