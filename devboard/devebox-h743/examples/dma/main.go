// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"time"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/dma/dmairq"

	"github.com/embeddedgo/stm32/devboard/devebox-h743/board/leds"
)

var (
	ch  dma.Channel
	tce rtos.Note
)

func copyDMA(dst, src unsafe.Pointer, n int, mode dma.Mode) dma.Error {
	tce.Clear()
	ch.Setup(dma.MTM | dma.IncP | dma.IncM | mode)
	ch.SetWordSize(4, 4)
	ch.SetLen(n)
	ch.SetAddrP(src)
	ch.SetAddrM(dst)
	ch.Clear(dma.EvAll, dma.ErrAll)
	ch.EnableIRQ(dma.Complete, dma.ErrAll)
	ch.Enable()
	tce.Sleep(-1)
	_, err := ch.Status()
	return err
}

func main() {
	d := dma.DMA(2)
	d.EnableClock(true)
	ch = d.AllocChannel()
	dmairq.SetISR(ch, dmaisr)

	n := 4000
	src := dma.MakeSlice[uint32](n, n)
	dst := dma.MakeSlice[uint32](n, n)
	srcAddr := unsafe.Pointer(&src[0])
	dstAddr := unsafe.Pointer(&dst[0])

	for i := range src {
		src[i] = uint32(i)
	}
	rtos.CacheMaint(rtos.DCacheFlush, srcAddr, n*4)
	rtos.CacheMaint(rtos.DCacheFlushInval, dstAddr, n*4)

	delay := time.Second
	if copyDMA(dstAddr, srcAddr, n, dma.FT4|dma.PB4|dma.MB4) != 0 {
		delay /= 4
	} else {
		for i, v := range src {
			if v != dst[i] {
				delay /= 4
				break
			}
		}
	}

	for {
		leds.User.Toggle()
		time.Sleep(delay)
	}

}

//go:nosplit
func dmaisr() {
	ch.DisableIRQ(dma.EvAll, dma.ErrAll)
	ch.Disable()
	tce.Wakeup()
}
