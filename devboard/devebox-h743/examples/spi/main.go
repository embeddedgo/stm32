// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/mmio"
	"fmt"
	"time"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spih"
	"github.com/embeddedgo/stm32/hal/system/console/uartcon"
	"github.com/embeddedgo/stm32/hal/usart"
	"github.com/embeddedgo/stm32/hal/usart/usart1"

	"github.com/embeddedgo/stm32/devboard/devebox-h743/board/leds"
	_ "github.com/embeddedgo/stm32/devboard/devebox-h743/board/leds"
	_ "github.com/embeddedgo/stm32/devboard/devebox-h743/board/system"
)

var (
	buf      [8]byte
	baudrate int
)

func main() {
	// GPIO pins assignment
	pa := gpio.PA()
	pa.EnableClock(true)
	conTx := pa.Pin(9)
	conRx := pa.Pin(10)

	pb := gpio.PB()
	pb.EnableClock(true)
	//blk := pb.Pin(0)
	//dc := pb.Pin(1)
	csn := pb.Pin(12)
	sck := pb.Pin(13)
	miso := pb.Pin(14)
	mosi := pb.Pin(15)

	// Serial console
	uartcon.Setup(usart1.Driver(), conRx, conTx, usart.Word8b, 115200, "USART1")
	outCfg := gpio.Config{Mode: gpio.Alt, Speed: gpio.VeryHigh}
	csn.Setup(&outCfg)
	csn.SetAltFunc(gpio.AF5)
	sck.Setup(&outCfg)
	sck.SetAltFunc(gpio.AF5)
	mosi.Setup(&outCfg)
	mosi.SetAltFunc(gpio.AF5)
	miso.Setup(&gpio.Config{Mode: gpio.AltIn})
	miso.SetAltFunc(gpio.AF5)

	d := spih.NewMaster(spih.SPI(2), dma.Channel{}, dma.Channel{})
	baudrate = d.Setup(spih.Word8b, 400e3)
	p := d.Periph()

	txdr := (*mmio.U8)(unsafe.Pointer(&p.TXDR))
	rxdr := (*mmio.U8)(unsafe.Pointer(&p.RXDR))
	for i := 0; ; i++ {
		fmt.Print(i, ": ")
		p.CR2.Store(8)
		d.Enable()
		txdr.Store(byte(i))
		txdr.Store(byte(i - 1))
		txdr.Store(byte(i - 2))
		p.TXDR.Store(uint32(i))
		txdr.Store(byte(i - 3))
		for p.SR.LoadBits(spih.RXP) == 0 {
		}
		buf[0] = rxdr.Load()
		buf[1] = rxdr.Load()
		buf[2] = rxdr.Load()
		for p.SR.LoadBits(spih.RXP) == 0 {
		}
		v := p.RXDR.Load()
		buf[3] = byte(v)
		buf[4] = byte(v >> 8)
		buf[5] = byte(v >> 16)
		buf[6] = byte(v >> 24)
		for p.SR.LoadBits(spih.RXP) == 0 {
		}
		buf[7] = rxdr.Load()
		d.Disable()
		leds.User.Toggle()
		time.Sleep(time.Second / 4)
	}
}
