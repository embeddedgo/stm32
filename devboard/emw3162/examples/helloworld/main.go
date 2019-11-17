// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/embeddedgo/stm32/devboard/emw3162/board"
	"github.com/embeddedgo/x/time"
)

func main() {
	board.Setup(true)
	for {
		println("Hello World!")
		time.Sleep(time.Second / 4)
	}
}
