// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Guncon reads the state of PS1 Namco GunCon (G-Con).
package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi2"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/system"
)

func main() {
	// Assign GPIO pins

	pb := gpio.PB()
	pb.EnableClock(true)
	csn := pb.Pin(12)
	sck := pb.Pin(13)
	miso := pb.Pin(14) // need 1 kΩ pull-up resistor (the STM32 built in pull-up is too weak)
	mosi := pb.Pin(15)

	// Configure peripherals

	spid := spi2.Driver()
	spid.UsePinMaster(sck, spi.SCK)
	spid.UsePinMaster(miso, spi.MISO)
	spid.UsePinMaster(mosi, spi.MOSI)
	spid.UsePinMaster(csn, spi.NSS) // use hardware slave-select
	spid.Setup(spi.Master|spi.SSOut|spi.CPOL1|spi.CPHA1|spi.LSBF, 250e3)
	spid.SetWordSize(8)
	spid.Enable()

	time.Sleep(time.Second)

	cmd := []byte{0x01, 0x42, 0x00}
	r := make([]byte, 9)
	for {
		println("read")
		spid.Enable()
		spid.WriteRead(cmd, r)
		spid.Disable()
		print(fmt.Sprintf(
			"%02x %02x %02x  %02x %02x  %02x %02x  %02x %02x\r\n",
			r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8],
		))
		time.Sleep(2 * time.Second)
	}
}
