// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package init

import (
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/system"
	"github.com/embeddedgo/stm32/hal/system/console/uartcon"
	"github.com/embeddedgo/stm32/hal/system/timer/rtcst"
	"github.com/embeddedgo/stm32/hal/usart"
	"github.com/embeddedgo/stm32/hal/usart/usart1"
)

func init() {
	system.Setup168(8)
	rtcst.Setup(rtcst.LSE, 1, 32768)
	uart := usart1.Driver()
	rx := gpio.A().Pin(9)
	tx := gpio.A().Pin(10)
	uartcon.Setup(uart, rx, tx, usart.Word8b, 115200, "USART1")
}
