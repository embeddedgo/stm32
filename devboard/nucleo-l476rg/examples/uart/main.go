// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/hal/usart"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/init"
)

var tts *usart.Driver

func main() {
	pa := gpio.A()
	pa.EnableClock(true)
	tx := pa.Pin(2)
	rx := pa.Pin(3)

	usart.UsePin(tx, gpio.AF7_USART2, usart.TXD)
	usart.UsePin(rx, gpio.AF7_USART2, usart.RXD)

	d := dma.DMA(1)
	d.EnableClock(true) // DMA clock must remain enabled in sleep mode.
	txdma, rxdma := d.Channel(7, 2), d.Channel(6, 2)
	tts = usart.NewDriver(usart.USART2(), txdma, rxdma)
	tts.Setup(usart.Word8b, 115200)
	tts.EnableRx(nil)
	tts.EnableTx()
	irq.USART2.Enable(rtos.IntPrioLow, 0)
	irq.DMA1_CH7.Enable(rtos.IntPrioLow, 0)

	tts.WriteString("\r\nEnter text:\r\n")
	buf := make([]byte, 80)
	for {
		n, err := tts.Read(buf)
		if err != nil {
			println(err.Error())
		}
		tts.WriteString("\r\nrecv: ")
		tts.Write(buf[:n])
		tts.WriteString("\r\n# ")
	}
}

//go:interrupthandler
func USART2_Handler() {
	tts.RxISR()
}

//go:interrupthandler
func DMA1_CH7_Handler() {
	tts.TxDMAISR()
}
