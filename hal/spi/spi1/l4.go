// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32l4x6

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
	d := dma.DMA(1)
	d.EnableClock(true)
	driver = spi.NewDriver(spi.SPI1(), d.Channel(3, 1), d.Channel(2, 1))
	irq.SPI1.Enable(rtos.IntPrioLow)
	irq.DMA1_CH2.Enable(rtos.IntPrioLow)
	irq.DMA1_CH3.Enable(rtos.IntPrioLow)
}

//go:interrupthandler
func _SPI1_Handler() { driver.ISR() }

//go:interrupthandler
func _DMA1_CH2_Handler() { driver.DMAISR(driver.RxDMA()) }

//go:interrupthandler
func _DMA1_CH3_Handler() { driver.DMAISR(driver.TxDMA()) }

//go:linkname _SPI1_Handler IRQ35_Handler
//go:linkname _DMA1_CH2_Handler IRQ12_Handler
//go:linkname _DMA1_CH3_Handler IRQ13_Handler
