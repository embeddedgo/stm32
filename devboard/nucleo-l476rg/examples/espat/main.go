// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Espat is an ESP-AT based TCP echo server. It uses the espat package directly
// (instead of espat/espnet). See ../espn for the example of the same TCP
// server implemented using simpler interface of the espat/espn package.
// Because of the insufficient RAM in STM32L476 you cannot use espat/espnet on
// this MCU.
package main

import (
	"time"

	"github.com/embeddedgo/espat"

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
	_, err := dev.Cmd("+CIPMUX=1")
	fatalErr(err)
	_, err = dev.Cmd("+CIPRECVMODE=1")
	fatalErr(err)

	println("Waiting for an IP address...")
	for msg := range dev.Async() {
		fatalErr(msg.Err)
		println(msg.Str)
		if msg.Str == "WIFI GOT IP" {
			break
		}
	}

	const port = "1111"

	dev.SetServer(true)
	_, err = dev.Cmd("+CIPSERVER=1," + port)
	fatalErr(err)

	println("Listen on :" + port)
	for conn := range dev.Server() {
		go handle(conn)
	}
}

var welcome = []byte("Echo Server\n\n")

func handle(c *espat.Conn) {
	println("Connected:", c.ID)
	if logErr(send(c, welcome)) {
		return
	}
	var buf [64]byte
	for {
		if _, ok := <-c.Ch; !ok {
			break // connection closed by remote part
		}
		n, err := c.Dev.CmdInt("+CIPRECVDATA=", buf[:], c.ID, len(buf))
		if logErr(err) {
			return
		}
		if logErr(send(c, buf[:n])) {
			return
		}
	}
	println("Closed:   ", c.ID)
}

func send(c *espat.Conn, p []byte) error {
	c.Dev.Lock()
	defer c.Dev.Unlock()
	if _, err := c.Dev.UnsafeCmd("+CIPSEND=", c.ID, len(p)); err != nil {
		return err
	}
	if _, err := c.Dev.UnsafeWrite(p); err != nil {
		return err
	}
	_, err := c.Dev.UnsafeCmd("")
	return err
}
