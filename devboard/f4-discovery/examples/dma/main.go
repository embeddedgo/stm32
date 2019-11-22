// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example forces the garbage collector to work hard and periodically
// prints statistics of the memory allocator.
package main

import (
	"embedded/rtos"
	"unsafe"

	"github.com/embeddedgo/stm32/devboard/f4-discovery/board"
	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/irq"
)

const n = 16 * 1024 / 4

var (
	ch     dma.Channel
	dmaErr dma.Error
	tce    rtos.Note

	src = make([]uint32, n)
	dst = make([]uint32, n)
)

func copyDMA(mode dma.Mode) {
	ch.Setup(dma.MTM | dma.IncP | dma.IncM | mode)
	ch.SetWordSize(unsafe.Sizeof(src[0]), unsafe.Sizeof(dst[0]))
	ch.SetLen(n)
	ch.SetAddrP(unsafe.Pointer(&src[0]))
	ch.SetAddrM(unsafe.Pointer(&dst[0]))
	tce.Clear()
	ch.Enable()
	tce.Sleep(-1)
	if dmaErr != 0 {
		println(dmaErr.Error())
	}
}

func check() {
	for i := range dst {
		if dst[i] != uint32(i) {
			println(" dst != src\n")
			return
		}
		dst[i] = 0
	}
}

var p *int

func main() {
	board.Setup(true)
	
	a := *p

	d := dma.DMA(2)
	d.EnableClock(true)
	ch = d.Channel(0, a)
	ch.EnableIRQ(dma.Complete, dma.ErrAll)
	irq.DMA2_Stream0.Enable(rtos.IntPrioLow)

	println("Initialize src")
	for i := range src {
		src[i] = uint32(i)
	}

	for {
		println("DMA")
		copyDMA(0)
		check()

		println("DMA FT1")
		copyDMA(dma.FT1)
		check()

		println("DMA FT2")
		copyDMA(dma.FT2)
		check()

		println("DMA FT3")
		copyDMA(dma.FT3)
		check()

		println("DMA FT4")
		copyDMA(dma.FT4)
		check()

		println("DMA FT4 PB4 MB4")
		copyDMA(dma.FT4 | dma.PB4 | dma.MB4)
		check()
	}
}

//go:interrupthandler
func DMA2_Stream0_Handler() {
	ev, err := ch.Status()
	ch.Clear(ev, err)
	if ev&dma.Complete != 0 || err != 0 {
		dmaErr = err
		tce.Wakeup()
	}
}
