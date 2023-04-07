// Copyright 2023 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"net"
	"time"

	"github.com/embeddedgo/espat"
	"github.com/embeddedgo/espat/espnet"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/usart"
	"github.com/embeddedgo/stm32/hal/usart/usart2"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/init"
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

	print("\n* F4 ready *\n\n")

	dev := espat.NewDevice("esp0", u, u)
	dev.Init(false)
	fatalErr(espnet.SetPasvRecv(dev, true))

	/*
		for msg := range dev.Async() {
			println(msg)
			if msg == "WIFI GOT IP" {
				break
			}
		}
	*/

	ls, err := espnet.ListenDev(dev, "tcp", 1111)
	fatalErr(err)

	println("listen on:", ls.Addr().String())
	for {
		c, err := ls.Accept()
		fatalErr(err)
		go handle(c)
	}
}

func handle(c net.Conn) {
	println("connected:", c.RemoteAddr().String())
	_, err := io.WriteString(c, "Echo Server\n\n")
	if logErr(err) {
		return
	}
	var buf [64]byte
	for {
		n, err := c.Read(buf[:])
		if err == io.EOF {
			return
		}
		if logErr(err) {
			return
		}
		_, err = c.Write(buf[:n])
		if logErr(err) {
			return
		}
	}
}
