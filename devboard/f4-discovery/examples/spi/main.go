// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// SPI loop test: wire PB14 and PB15 together.
package main

import (
	"time"

	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi2"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/init"
)

func main() {
	// Allocate GPIO pins

	pb := gpio.B()
	pb.EnableClock(true)
	sck := pb.Pin(13)
	miso := pb.Pin(14)
	mosi := pb.Pin(15)

	// Configure SPI pins

	sd := spi2.Driver()
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
