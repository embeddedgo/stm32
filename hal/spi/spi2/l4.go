// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32l4x6

package spi2

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/hal/spi"
)

const altFunc = gpio.AF5_SPI2

func setupDriver() {
	d := dma.DMA(1)
	d.EnableClock(true)
	driver = spi.NewDriver(spi.SPI2(), d.Channel(5, 1), d.Channel(4, 1))
	irq.SPI2.Enable(rtos.IntPrioLow)
	irq.DMA1_CH4.Enable(rtos.IntPrioLow)
	irq.DMA1_CH5.Enable(rtos.IntPrioLow)
}

//go:interrupthandler
func _SPI2_Handler() { driver.ISR() }

//go:interrupthandler
func _DMA1_CH4_Handler() { driver.DMAISR(driver.RxDMA()) }

//go:interrupthandler
func _DMA1_CH5_Handler() { driver.DMAISR(driver.TxDMA()) }

//go:linkname _SPI2_Handler IRQ36_Handler
//go:linkname _DMA1_CH4_Handler IRQ14_Handler
//go:linkname _DMA1_CH5_Handler IRQ15_Handler
