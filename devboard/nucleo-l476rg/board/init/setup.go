// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package init

import (
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/system"
	"github.com/embeddedgo/stm32/hal/system/console/uartcon"
	"github.com/embeddedgo/stm32/hal/system/timer/rtcst"
	"github.com/embeddedgo/stm32/hal/usart"
	"github.com/embeddedgo/stm32/hal/usart/usart2"
)

func init() {
	system.Setup80(0, 0)
	rtcst.Setup(rtcst.LSE, 1, 32768)
	uart := usart2.Driver()
	rx := gpio.A().Pin(3)
	tx := gpio.A().Pin(2)
	// Use "light" version of console because of small RAM.
	uartcon.SetupLight(uart, rx, tx, usart.Word8b, 115200, "USART2")
}
