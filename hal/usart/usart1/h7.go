// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package usart1

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
	rxdma := d.Channel(0, 0)
	rxdma.SetMux(dma.USART1_RX)
	txdma := d.Channel(1, 0)
	txdma.SetMux(dma.USART1_TX)
	driver = usart.NewDriver(usart.USART1(), rxdma, txdma)
	irq.USART1.Enable(rtos.IntPrioLow, 0)
	irq.DMA1_STR1.Enable(rtos.IntPrioLow, 0)
}

//go:interrupthandler
func _USART1_Handler() { driver.RxISR() }

//go:interrupthandler
func _DMA1_STR1_Handler() { driver.TxDMAISR() }

//go:linkname _USART1_Handler IRQ37_Handler
//go:linkname _DMA1_STR1_Handler IRQ12_Handler
