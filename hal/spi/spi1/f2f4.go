// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407

package spi1

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/hal/spi"
)

const altFunc = gpio.AF5_SPI1

func setupDriver() {
	d := dma.DMA(2)
	d.EnableClock(true)
	driver = spi.NewDriver(spi.SPI1(), d.Channel(3, 3), d.Channel(0, 3))
	irq.SPI1.Enable(rtos.IntPrioLow, -1)
	irq.DMA2_Stream0.Enable(rtos.IntPrioLow, -1)
	irq.DMA2_Stream3.Enable(rtos.IntPrioLow, -1)
}

//go:interrupthandler
func _SPI1_Handler() { driver.ISR() }

//go:interrupthandler
func _DMA2_Stream0_Handler() { driver.DMAISR(driver.RxDMA()) }

//go:interrupthandler
func _DMA2_Stream3_Handler() { driver.DMAISR(driver.TxDMA()) }

//go:linkname _SPI1_Handler IRQ35_Handler
//go:linkname _DMA2_Stream0_Handler IRQ56_Handler
//go:linkname _DMA2_Stream3_Handler IRQ59_Handler
