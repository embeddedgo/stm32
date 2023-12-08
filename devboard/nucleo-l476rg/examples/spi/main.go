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

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/system"
)

func main() {
	// Allocate GPIO pins

	pa := gpio.PA()
	pa.EnableClock(true)
	sck := pa.Pin(5)
	miso := pa.Pin(6)
	mosi := pa.Pin(7)

	// Configure SPI pins

	sd := spi1.Driver()
	sd.UsePinMaster(sck, spi.SCK)
	sd.UsePinMaster(mosi, spi.MOSI)
	sd.UsePinMaster(miso, spi.MISO)

	// Configure and enable SPI

	sd.Setup(spi.Master|spi.SoftSS|spi.ISSHigh, 1e6)
	sd.SetWordSize(8)
	sd.Enable()

	var buf [40]byte
	for i := 0; ; i++ {
		sd.WriteStringRead("Hello world!", buf[:])
		println(string(buf[:]))
		time.Sleep(time.Second)
	}
}
