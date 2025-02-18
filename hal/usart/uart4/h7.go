// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package uart4

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
	rxdma := d.Channel(2, 0)
	rxdma.SetMux(dma.UART4_RX)
	txdma := d.Channel(3, 0)
	txdma.SetMux(dma.UART4_TX)
	driver = usart.NewDriver(usart.UART4(), rxdma, txdma)
	irq.UART4.Enable(rtos.IntPrioLow, 0)
	irq.DMA1_STR3.Enable(rtos.IntPrioLow, 0)
}

//go:interrupthandler
func _UART4_Handler() { driver.RxISR() }

//go:interrupthandler
func _DMA1_STR3_Handler() { driver.TxDMAISR() }

//go:linkname _UART4_Handler IRQ52_Handler
//go:linkname _DMA1_STR3_Handler IRQ14_Handler
