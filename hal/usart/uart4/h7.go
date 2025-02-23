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
	rxdma := d.AllocChannel()
	rxdma.SetMux(dma.UART4_RX)
	txdma := d.AllocChannel()
	txdma.SetMux(dma.UART4_TX)
	driver = usart.NewDriver(usart.UART4(), rxdma, txdma)
	irq.UART4.Enable(rtos.IntPrioLow, 0)
	dmairq.SetISR(txdma, driver.TxDMAISR)
}

//go:interrupthandler
func _UART4_Handler() { driver.RxISR() }

//go:linkname _UART4_Handler IRQ52_Handler
