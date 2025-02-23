// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package usart2

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/dma/dmairq"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/hal/usart"
)

func setupDriver() {
	d := dma.DMA(1)
	d.EnableClock(true)
	rxdma := d.AllocChannel()
	rxdma.SetMux(dma.USART2_RX)
	txdma := d.AllocChannel()
	txdma.SetMux(dma.USART2_TX)
	driver = usart.NewDriver(usart.USART2(), rxdma, txdma)
	irq.USART2.Enable(rtos.IntPrioLow, 0)
	dmairq.SetISR(txdma, driver.TxDMAISR)
}

//go:interrupthandler
func _USART2_Handler() { driver.RxISR() }

//go:linkname _USART2_Handler IRQ38_Handler
