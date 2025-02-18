// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32l4x6

package uart4

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
	driver = usart.NewDriver(usart.UART4(), d.Channel(5, 2), d.Channel(3, 2))
	irq.UART4.Enable(rtos.IntPrioLow, 0)
	irq.DMA2_CH3.Enable(rtos.IntPrioLow, 0)
}

//go:interrupthandler
func _UART4_Handler() { driver.RxISR() }

//go:interrupthandler
func _DMA2_CH3_Handler() { driver.TxDMAISR() }

//go:linkname _UART4_Handler IRQ52_Handler
//go:linkname _DMA2_CH3_Handler IRQ58_Handler
