// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407

package usart6

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/hal/usart"
)

func setupDriver() {
	d := dma.DMA(2)
	d.EnableClock(true)
	driver = usart.NewDriver(usart.USART6(), d.Channel(1, 5), d.Channel(7, 5))
	irq.USART6.Enable(rtos.IntPrioLow, 0)
	irq.DMA2_STREAM7.Enable(rtos.IntPrioLow, 0)
}

//go:interrupthandler
func _USART6_Handler() { driver.RxISR() }

//go:interrupthandler
func _DMA2_STREAM7_Handler() { driver.TxDMAISR() }

//go:linkname _USART6_Handler IRQ71_Handler
//go:linkname _DMA2_STREAM7_Handler IRQ70_Handler
