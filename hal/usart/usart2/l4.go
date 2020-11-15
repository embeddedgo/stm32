// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32l4x6

package usart2

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/hal/usart"
)

func setupDriver() {
	d := dma.DMA(1)
	d.EnableClock(true)
	driver = usart.NewDriver(usart.USART2(), d.Channel(6, 2), d.Channel(7, 2))
	irq.USART2.Enable(rtos.IntPrioLow, 0)
	irq.DMA1_CH7.Enable(rtos.IntPrioLow, 0)
}

//go:interrupthandler
func _USART2_Handler() { driver.RxISR() }

//go:interrupthandler
func _DMA1_CH7_Handler() { driver.TxDMAISR() }

//go:linkname _USART2_Handler IRQ38_Handler
//go:linkname _DMA1_CH7_Handler IRQ17_Handler
