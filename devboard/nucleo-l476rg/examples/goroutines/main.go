// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example runs two goroutines that set on/off the user LED concurently.
package main

import (
	"math/rand"
	"time"

	"github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/leds"
)

func randDelay() {
	time.Sleep(time.Second / 1024 * time.Duration(1024 + rand.Intn(1024)))
}


func ledOn() {
	for {
		randDelay()
		leds.User.SetOn()
	}
}

func main() {
	go ledOn()
	for {
		randDelay()
		leds.User.SetOff()
	}
}
