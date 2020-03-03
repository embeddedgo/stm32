// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package evedci

import (
	"github.com/embeddedgo/stm32/hal/exti"
	"github.com/embeddedgo/stm32/hal/gpio"
)

// UseIntPin setups GPIO pin and EXTI peripheral to detect EVE interrupts.
func UseIntPin(pin gpio.Pin) {
	pin.Setup(&gpio.Config{Mode: gpio.In})
	line := exti.Lines(pin.Mask())
	line.Connect(pin.Port())
	line.EnableFallTrig()
	line.EnableIRQ()
}
