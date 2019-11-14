// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package board

import (
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/system"
	"github.com/embeddedgo/stm32/hal/system/timer/systick"
)

func Setup(all bool) {
	system.Setup120(26)
	systick.Setup(2e6)

	if !all {
		return
	}

	gpio.B().EnableClock(true)
	cfg := &gpio.Config{Mode: gpio.Out, Speed: gpio.Low}
	D1.Pin().Setup(cfg)
	cfg.Driver = gpio.OpenDrain
	D2.Pin().Setup(cfg)
}
