// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// SPI loop test: wire PA6 and PA7 together.
package main

import (
	"time"

	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi2"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/init"
)

func main() {
	// Allocate GPIO pins

	pb := gpio.B()
	pb.EnableClock(true)
	nss := pb.Pin(12)
	sck := pb.Pin(13)
	miso := pb.Pin(14)
	mosi := pb.Pin(15)

	// Configure SPI pins

	spi2.UsePinMaster(spi.NSS, nss)
	spi2.UsePinMaster(spi.SCK, sck)
	spi2.UsePinMaster(spi.MOSI, mosi)
	spi2.UsePinMaster(spi.MISO, miso)

	// Configure and enable SPI

	d := spi2.Driver()
	d.Setup(spi.Master|spi.SoftSS|spi.ISSHigh, 1e6)
	d.SetWordSize(8)
	d.Enable()

	var buf [40]byte
	for i := 0; ; i++ {
		d.WriteStringRead("Hello world!", buf[:])
		println(string(buf[:]))
		time.Sleep(time.Second)
	}
}
