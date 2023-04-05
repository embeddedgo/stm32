// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example tests the USART1 interface. The onboard USB to UART
// converter is connected to USART2 and used as system console
// (os.Stdin, os.Stdout, os.Stderr). To see something you need to
// connect an external USB to UART converter to D2 and D8 pins.
package main

import (
	"github.com/embeddedgo/espat"
	"github.com/embeddedgo/espat/espnet"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/system"
	"github.com/embeddedgo/stm32/hal/system/timer/rtcst"
	"github.com/embeddedgo/stm32/hal/usart"
	"github.com/embeddedgo/stm32/hal/usart/usart1"
)

func fatalErr(err error) {
	for err != nil {
		println("error:", err.String())
	}
}

func main() {
	system.Setup80(0, 0)
	rtcst.Setup(rtcst.LSE, 1, 32768)

	// GPIO pins assignment
	pa := gpio.PA()
	pa.EnableClock(true)
	tx := pa.Pin(9)  // CN5 D8
	rx := pa.Pin(10) // CN9 D2

	// Configure and enable USART
	u := usart1.Driver()
	u.UsePin(tx, usart.TXD)
	u.UsePin(rx, usart.RXD)
	u.Setup(usart.Word8b, 115200)
	u.EnableTx()
	u.EnableRx(256)

	dev := espat.NewDevice("esp0", u, u)
	fatalErr(dev.Init(true))
	fatalErr(espnet.SetMultiConn(dev, true))
	fatalErr(espnet.SetPasvRecv(dev, true))

}
