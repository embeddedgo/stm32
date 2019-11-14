// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows the basic usage of available LED and button.
package main

import (
	"github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board"
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
		board.UserLED.SetOn()
		delay()
		board.UserLED.SetOff()
		delay()
	}
}
