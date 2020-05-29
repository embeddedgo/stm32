// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example tests different ways of coping memory. It also shows how to use
// DMA for memory to memory transfers.
package main

import (
	"embedded/rtos"
	"time"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/irq"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/init"
)

var (
	src, dst [5000]uint32
	ch       dma.Channel
	tce      rtos.Note
)

func copyDMA(n int) {
	tce.Clear()
	ch.Setup(dma.MTM | dma.IncP | dma.IncM)
	ch.SetWordSize(unsafe.Sizeof(src[0]), unsafe.Sizeof(dst[0]))
	ch.SetLen(n)
	ch.SetAddrP(unsafe.Pointer(&src[0]))
	ch.SetAddrM(unsafe.Pointer(&dst[0]))
	ch.Clear(dma.EvAll, dma.ErrAll)
	ch.EnableIRQ(dma.Complete, dma.ErrAll)
	ch.Enable()
	tce.Sleep(-1)
	if _, err := ch.Status(); err != 0 {
		println(err.Error())
	}
}

func printSpeed(t time.Duration, n int, check bool) {
	t1 := rtos.Nanotime()
	t2 := rtos.Nanotime()
	dt := int64(t1-t) - int64(t2-t1)
	if check {
		for i := 0; i < n; i++ {
			if dst[i] != uint32(i) {
				println(" dst != src\n")
				return
			}
			dst[i] = 0
		}
	}
	println("", (int64(uintptr(n)*unsafe.Sizeof(src[0]))*1e6+dt/2)/dt, "kB/s")
}

func main() {
	d := dma.DMA(1)
	d.EnableClock(true)
	ch = d.Channel(1, 0)
	irq.DMA1_CH1.Enable(rtos.IntPrioLow, -1)

	for {
		for n := 1000; n <= len(src); n += 1000 {
			println("\nTransfer length:", n, "words\n")

			print("Initialize src                      ")
			t := rtos.Nanotime()
			for i := 0; i < n; i++ {
				src[i] = uint32(i)
			}
			printSpeed(t, n, false)

			print("for i:=0; i<n; i++ { dst[i]=src[i] }")
			t = rtos.Nanotime()
			for i := 0; i < n; i++ {
				dst[i] = src[i]
			}
			printSpeed(t, n, true)

			print("copy(dst[:n], src[:n])              ")
			t = rtos.Nanotime()
			copy(dst[:n], src[:n])
			printSpeed(t, n, true)

			print("DMA                                 ")
			t = rtos.Nanotime()
			copyDMA(n)
			printSpeed(t, n, true)

			time.Sleep(2 * time.Second)
		}
		time.Sleep(2 * time.Second)
		println("\n---")
	}
}

//go:interrupthandler
func DMA1_CH1_Handler() {
	ch.DisableIRQ(dma.EvAll, dma.ErrAll)
	ch.Disable()
	tce.Wakeup()
}
