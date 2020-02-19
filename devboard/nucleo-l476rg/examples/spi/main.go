// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// SPI loop test: wire PA6 and PA7 together.
package main

import (
	"time"

	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi1"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/init"
)

func main() {
	// Allocate GPIO pins

	pa := gpio.A()
	pa.EnableClock(true)
	sck := pa.Pin(5)
	miso := pa.Pin(6)
	mosi := pa.Pin(7)

	// Configure SPI pins

	spi1.UsePinMaster(spi.SCK, sck)
	spi1.UsePinMaster(spi.MOSI, mosi)
	spi1.UsePinMaster(spi.MISO, miso)

	// Configure and enable SPI

	d := spi1.Driver()
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
