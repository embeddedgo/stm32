// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to use SPI peripheral to read the on-board LIS3DSH
// accelerometer.
package main

import (
	"embedded/rtos"

	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi1"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/init"
)

func main() {
	// Allocate GPIO pins

	pa := gpio.A()
	pa.EnableClock(true)
	sck := pa.Pin(5)
	miso := pa.Pin(6)
	mosi := pa.Pin(7)

	pe := gpio.E()
	pe.EnableClock(true)
	cs := pe.Pin(3)

	// Configure SPI pins

	spi1.UsePinsMaster(sck, mosi, miso)
	cs.Set() // CS active state is low
	cs.Setup(&gpio.Config{Mode: gpio.Out})

	// Configure and enable SPI

	d := spi1.Driver()
	d.Setup(spi.Master|spi.CPOL1|spi.CPHA1|spi.SoftSS|spi.ISSHigh, 10e6)
	d.SetWordSize(8)
	d.Enable()

	// LIS3DSH

	const (
		READ      = 1 << 7
		OFF_X     = 0x10
		CTRL_REG4 = 0x20
		OUT_X_L   = 0x28
	)

	buf := make([]byte, 7)

	// initialize
	buf[0] = CTRL_REG4
	buf[1] = 0x57 // 50 Hz, block data update, all axis enabled
	cs.Clear()
	d.WriteRead(buf[:2], nil)
	cs.Set()

	// read correction offsets
	buf[0] = READ | OFF_X
	cs.Clear()
	d.WriteRead(buf[:4], nil)
	cs.Set()
	ox := int32(int8(buf[1]))
	oy := int32(int8(buf[2]))
	oz := int32(int8(buf[3]))
	println("correction offsets:", ox, oy, oz)

	// print values in G/1000 units
	for i := 0; ; i++ {
		buf[0] = READ | OUT_X_L
		cs.Clear()
		d.WriteRead(buf, buf)
		cs.Set()
		x := int32(int16(buf[1])|int16(buf[2])<<8) + ox
		y := int32(int16(buf[3])|int16(buf[4])<<8) + oy
		z := int32(int16(buf[5])|int16(buf[6])<<8) + oz
		x = x * 2000 / 32768
		y = y * 2000 / 32768
		z = z * 2000 / 32768
		println(x, y, z)
		rtos.Nanosleep(1e8)
	}
}
