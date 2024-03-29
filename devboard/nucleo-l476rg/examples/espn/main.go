// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Espn is ESP-AT based TCP echo server. See also ../espat for slightly less
// memory consuming version of this program that uses the espat package
// directly. Because of the insufficient RAM in STM32L476 you cannot use
// espat/espnet on this MCU.
package main

import (
	"io"
	"time"

	"github.com/embeddedgo/espat"
	"github.com/embeddedgo/espat/espn"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/system"
	"github.com/embeddedgo/stm32/hal/system/timer/rtcst"
	"github.com/embeddedgo/stm32/hal/usart"
	"github.com/embeddedgo/stm32/hal/usart/usart1"
)

func logErr(err error) bool {
	if err != nil {
		println("error:", err.Error())
		return true
	}
	return false
}

func fatalErr(err error) {
	for err != nil {
		println("error:", err.Error())
		time.Sleep(time.Second)
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
	fatalErr(espn.SetPasvRecv(dev, true))

	println("Waiting for an IP address...")
	for msg := range dev.Async() {
		fatalErr(msg.Err)
		println(msg.Str)
		if msg.Str == "WIFI GOT IP" {
			break
		}
	}

	ls, err := espn.ListenDev(dev, "tcp", ":1111")
	fatalErr(err)

	println("Listen on:", ls.Addr().String())
	for {
		c, err := ls.Accept()
		fatalErr(err)
		go handle(c)
	}
}

func handle(c *espn.Conn) {
	var buf [64]byte
	println("Connected:", c.RemoteAddr().String())
	_, err := io.WriteString(c, "Echo Server\n\n")
	if logErr(err) {
		return
	}
	for {
		n, err := c.Read(buf[:])
		if err == io.EOF {
			break
		}
		if logErr(err) {
			return
		}
		_, err = c.Write(buf[:n])
		if logErr(err) {
			return
		}
	}
	c.Close()
	println("Closed:   ", c.RemoteAddr().String())
}
