// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// SPI loop test: wire PA6 and PA7 together.
package main

import (
	"time"

	"github.com/embeddedgo/stm32/hal/exti"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi2"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/init"
)

func main() {
	// Allocate GPIO pins

	pb := gpio.B()
	pb.EnableClock(true)
	irqn := pb.Pin(1)
	pdn := pb.Pin(11)
	csn := pb.Pin(12)
	sck := pb.Pin(13)
	miso := pb.Pin(14)
	mosi := pb.Pin(15)

	// EVE control lines

	pdn.Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.High})
	irqn.Setup(&gpio.Config{Mode: gpio.In})
	irqline := exti.Lines(irqn.Mask())
	irqline.Connect(irqn.Port())
	irqline.EnableFallTrig()
	irqline.EnableIRQ()
	//rtos.IRQ(irq.EXTI1).Enable()

	// EVE SPI

	spi2.UsePinMaster(spi.NSS, csn)
	spi2.UsePinMaster(spi.SCK, sck)
	spi2.UsePinMaster(spi.MOSI, mosi)
	spi2.UsePinMaster(spi.MISO, miso)

	// Configure and enable SPI

	d := spi2.Driver()
	d.Setup(spi.Master|spi.HardSS|spi.SSOut, 11e6)
	d.SetWordSize(8)
	
	// Init

	pdn.Clear()
	time.Sleep(20 * time.Millisecond)
	pdn.Set()
	time.Sleep(20 * time.Millisecond) // Wait 20 ms for internal oscilator and PLL.

	hostCmd(cmdClkExt, 0) // Select external 12 MHz oscilator as clock source.
	hostCmd(cmdActive, 0)

	// Read both possible REG_ID locations for max. 300 ms, then check CHIPID.
	for i := 0; i < 30; i++ {
		if readByte(ramreg) == 0x7C {
			break
		}
		if readByte(ramreg2) == 0x7C {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if err := d.Err(true); err != nil {
		println("error:", err.Error())
		for {
		}
	}
	chipid := readUint32(chipidAddr)
	println("chip id:", chipid)
}

const chipidAddr = 0x0C0000

// Host commands for initialisation.
const (
	cmdActive = 0
	cmdClkExt = 0x44
)

// REG_ID adresses
const (
	ramreg  = 0x102400
	ramreg2 = 0x302000
)

func hostCmd(cmd, param byte) {
	d := spi2.Driver()
	d.Enable()
	d.WriteRead([]byte{byte(cmd), param, 0}, nil)
	d.Disable()
}

func readByte(addr int) byte {
	buf := [3]byte{byte(addr >> 16), byte(addr >> 8), byte(addr)}
	d := spi2.Driver()
	d.Enable()
	d.WriteRead(buf[:3], nil)
	d.WriteRead(nil, buf[:2])
	d.Disable()
	return buf[1]
}

func readUint32(addr int) uint32 {
	buf := [5]byte{byte(addr >> 16), byte(addr >> 8), byte(addr)}
	d := spi2.Driver()
	d.Enable()
	d.WriteRead(buf[:3], nil)
	d.WriteRead(nil, buf[:5])
	d.Disable()
	return uint32(buf[1]) | uint32(buf[2])<<8 | uint32(buf[3])<<16 |
		uint32(buf[4])<<24
}
