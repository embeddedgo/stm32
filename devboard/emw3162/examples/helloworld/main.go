// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/embeddedgo/stm32/devboard/emw3162/board"
	"github.com/embeddedgo/time"
)

func main() {
	board.Setup()
	for {
		board.Red.SetOn()
		println("Hello World!")
		board.Red.SetOff()
		time.Sleep(time.Second / 4)
	}
}
