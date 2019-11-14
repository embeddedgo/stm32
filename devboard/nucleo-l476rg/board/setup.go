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
	system.Setup80(0, 0)
	systick.Setup(2e6)

	if !all {
		return
	}

	gpio.A().EnableClock(true)
	LD2.Pin().Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.Low})

	gpio.C().EnableClock(true)
	B1.Pin().Setup(&gpio.Config{Mode: gpio.In})
}
