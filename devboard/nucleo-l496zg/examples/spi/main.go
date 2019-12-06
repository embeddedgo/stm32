// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/hal/spi"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l496zg/board/init"
)

var sd *spi.Driver

func main() {
	pa := gpio.A()
	pa.EnableClock(true)

	spiport, sck, miso, mosi := pa, gpio.Pin5, gpio.Pin6, gpio.Pin7

	spiport.Setup(sck|mosi, &gpio.Config{Mode: gpio.Alt, Speed: gpio.High})
	spiport.Setup(miso, &gpio.Config{Mode: gpio.AltIn})
	spiport.SetAltFunc(sck|miso|mosi, gpio.AF5_SPI1)
	d := dma.DMA(1)
	d.EnableClock(true)
	txdc, rxdc := d.Channel(3, 0), d.Channel(2, 0)
	txdc.SetRequest(dma.DMA1_SPI1)
	rxdc.SetRequest(dma.DMA1_SPI1)
	sd = spi.NewDriver(spi.SPI1(), txdc, rxdc)
	irq.SPI1.Enable(rtos.IntPrioLow)
	irq.DMA1_CH2.Enable(rtos.IntPrioLow)
	irq.DMA1_CH3.Enable(rtos.IntPrioLow)

	p := sd.Periph()
	p.EnableClock(true)
	p.Disable()
	p.SetConfig(spi.Master|spi.SoftSS|spi.ISSHigh, 1e6)
	p.SetWordSize(8)
	p.Enable()

	var buf [16]byte
	for {
		n := sd.WriteStringRead("Hello world!", buf[:12])
		println(n, string(buf[:n]))
		rtos.Nanosleep(1e9)
	}
}

//go:interrupthandler
func SPI1_Handler() { sd.ISR() }

//go:interrupthandler
func DMA1_CH2_Handler() { sd.DMAISR(sd.RxDMA()) }

//go:interrupthandler
func DMA1_CH3_Handler() { sd.DMAISR(sd.TxDMA()) }
