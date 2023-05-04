// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Espat is ESP-AT based TCP echo server. It uses the espat package directly
// (instead of espat/espnet) because of the insufficient RAM in STM32L476. See
// ../espnet for the example of same TCP server implemented using espat/espnet.
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

	time.Sleep(time.Second) // waiting for the ESP-AT module in case of power up

	dev := espat.NewDevice("esp0", u, u)
	fatalErr(dev.Init(true))
	_, err := dev.Cmd("+CIPMUX=1")
	fatalErr(err)
	_, err = dev.Cmd("+CIPRECVMODE=1")
	fatalErr(err)

	println("waiting for an IP address...")
	for msg := range dev.Async() {
		fatalErr(msg.Err)
		println(msg.Str)
		if msg.Str == "WIFI GOT IP" {
			break
		}
	}

	dev.SetServer(true)
	_, err = dev.Cmd("+CIPSERVER=1,1111")
	fatalErr(err)

	println()
	println("listen on :1111")
	for conn := range dev.Server() {
		go handle(conn)
	}
}

var welcome = []byte("Echo Server\n\n")

func handle(c *espat.Conn) {
	println("connected:", c.ID)
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
	println("closed:  ", c.ID)
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
