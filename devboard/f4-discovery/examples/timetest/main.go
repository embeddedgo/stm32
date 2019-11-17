// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example tesets setting the time.
package main

import (
	"github.com/embeddedgo/stm32/devboard/f4-discovery/board"
	"github.com/embeddedgo/x/time"
	"github.com/embeddedgo/x/time/tz"
)

func main() {
	board.Setup(true)
	for i := 0; i < 3; i++ {
		println(time.Now().String())
		time.Sleep(time.Second)
	}
	time.Local = &tz.EuropeWarsaw
	time.Set(time.Date(2019, 11, 17, 11, 44, 35, 9991, time.Local))
	for {
		println(time.Now().String())
		time.Sleep(time.Second)
	}
}
