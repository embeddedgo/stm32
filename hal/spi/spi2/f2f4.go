// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407

package spi2

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/hal/spi"
)

func setupDriver() {
	d := dma.DMA(1)
	d.EnableClock(true)
	driver = spi.NewDriver(spi.SPI2(), d.Channel(3, 0), d.Channel(4, 0))
	irq.SPI2.Enable(rtos.IntPrioLow, 0)
	irq.DMA1_STREAM3.Enable(rtos.IntPrioLow, 0)
	irq.DMA1_STREAM4.Enable(rtos.IntPrioLow, 0)
}

//go:interrupthandler
func _SPI2_Handler() { driver.ISR() }

//go:interrupthandler
func _DMA1_STREAM3_Handler() { driver.DMAISR(driver.RxDMA()) }

//go:interrupthandler
func _DMA1_STREAM4_Handler() { driver.DMAISR(driver.TxDMA()) }

//go:linkname _SPI2_Handler IRQ36_Handler
//go:linkname _DMA1_STREAM3_Handler IRQ14_Handler
//go:linkname _DMA1_STREAM4_Handler IRQ15_Handler
