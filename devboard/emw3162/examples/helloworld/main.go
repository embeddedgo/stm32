// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"

	_ "github.com/embeddedgo/stm32/devboard/emw3162/board/system"
)

func main() {
	for {
		println("Hello World!")
		time.Sleep(time.Second / 4)
	}
}
