// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407

package spi3

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/hal/spi"
)

const altFunc = gpio.AF6_SPI3

func setupDriver() {
	d := dma.DMA(1)
	d.EnableClock(true)
	driver = spi.NewDriver(spi.SPI3(), d.Channel(0, 0), d.Channel(7, 0))
	irq.SPI3.Enable(rtos.IntPrioLow, 0)
	irq.DMA1_Stream0.Enable(rtos.IntPrioLow, 0)
	irq.DMA1_Stream7.Enable(rtos.IntPrioLow, 0)
}

//go:interrupthandler
func _SPI3_Handler() { driver.ISR() }

//go:interrupthandler
func _DMA1_Stream0_Handler() { driver.DMAISR(driver.RxDMA()) }

//go:interrupthandler
func _DMA1_Stream7_Handler() { driver.DMAISR(driver.TxDMA()) }

//go:linkname _SPI3_Handler IRQ51_Handler
//go:linkname _DMA1_Stream0_Handler IRQ11_Handler
//go:linkname _DMA1_Stream7_Handler IRQ47_Handler
