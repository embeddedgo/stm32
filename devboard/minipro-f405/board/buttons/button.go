// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buttons

import (
	"github.com/embeddedgo/stm32/hal/gpio"

	_ "github.com/embeddedgo/stm32/devboard/minipro-f405/board/system"
)

// Onboard buttons
const (
	K1 Button = 0x00 // PA0

	User = K1
)

type Button uint8

func prt(b Button) int  { return int(b) >> 4 }
func pin(b Button) uint { return uint(b) & 15 }

func (b Button) Read() int {
	return int(gpio.P(prt(b)).Load()>>pin(b))&1 ^ 1
}
func (b Button) Pin() gpio.Pin {
	return gpio.P(prt(b)).Pin(int(pin(b)))
}

func init() {
	gpio.PA().EnableClock(true)
	K1.Pin().Setup(&gpio.Config{Mode: gpio.In})
}
