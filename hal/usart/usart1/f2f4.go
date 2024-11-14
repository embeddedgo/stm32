// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32f215 || stm32f407

package usart1

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
	driver = usart.NewDriver(usart.USART1(), d.Channel(2, 4), d.Channel(7, 4))
	irq.USART1.Enable(rtos.IntPrioLow, 0)
	irq.DMA2_STREAM7.Enable(rtos.IntPrioLow, 0)
}

//go:interrupthandler
func _USART1_Handler() { driver.RxISR() }

//go:interrupthandler
func _DMA2_STREAM7_Handler() { driver.TxDMAISR() }

//go:linkname _USART1_Handler IRQ37_Handler
//go:linkname _DMA2_STREAM7_Handler IRQ70_Handler
