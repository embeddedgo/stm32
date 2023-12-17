// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Espnet is an ESP-AT based TCP echo server. See also ../espat for the less
// memory consuming  ../espat program that uses the espat package directly and
// has much lower memory requirements. See also the same example written for
// Nucleo-L476RG and Teensy 4 development boards.
package main

import (
	"io"
	"net"
	"strings"
	"time"

	"github.com/embeddedgo/espat"
	"github.com/embeddedgo/espat/espnet"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/usart"
	"github.com/embeddedgo/stm32/hal/usart/usart2"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/system"
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
	// GPIO pins assignment
	pa := gpio.PA()
	pa.EnableClock(true)
	tx := pa.Pin(2)
	rx := pa.Pin(3)

	// Configure and enable USART
	u := usart2.Driver()
	u.UsePin(tx, usart.TXD)
	u.UsePin(rx, usart.RXD)
	u.Setup(usart.Word8b, 115200)
	u.EnableTx()
	u.EnableRx(256)

	print("Initializing ESP-AT module... ")
	dev := espat.NewDevice("esp0", u, u)
	fatalErr(dev.Init(true))
	fatalErr(espnet.SetPasvRecv(dev, true))
	println("OK")

	println("Waiting for an IP address...")
	for msg := range dev.Async() {
		fatalErr(msg.Err)
		println(msg.Str)
		if msg.Str == "WIFI GOT IP" {
			break
		}
	}
	txt, err := dev.CmdStr("+CIPSTA?")
	fatalErr(err)
	println(strings.ReplaceAll(txt, "+CIPSTA:", ""))

	ls, err := espnet.ListenDev(dev, "tcp", ":1111")
	fatalErr(err)

	println("Listen on:", ls.Addr().String())
	for {
		c, err := ls.Accept()
		fatalErr(err)
		go handle(c)
	}
}

func handle(c net.Conn) {
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
