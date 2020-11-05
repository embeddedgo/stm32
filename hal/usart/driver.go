// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usart

import (
	"embedded/mmio"
	"embedded/rtos"
	"time"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
)

type DriverError byte

const (
	ErrBufOverflow DriverError = iota + 1
	ErrTimeout
)

func (e DriverError) Error() string {
	switch e {
	case ErrBufOverflow:
		return "buffer overflow"
	case ErrTimeout:
		return "timeout"
	default:
		return ""
	}
}

type Driver struct {
	timeoutRx time.Duration
	timeoutTx time.Duration
	p         *Periph
	txDMA     dma.Channel
	rxDMA     dma.Channel
	rxBuf     []byte // Rx ring buffer for RxDMA.
	txDone    rtos.Note
	rxReady   rtos.Note
	rxp       int
}

// NewDriver returns a new driver for p.
func NewDriver(p *Periph, txdma, rxdma dma.Channel) *Driver {
	return &Driver{
		timeoutRx: -1,
		timeoutTx: -1,
		p:         p,
		txDMA:     txdma,
		rxDMA:     rxdma,
	}
}

func (d *Driver) Periph() *Periph {
	return d.p
}

func (d *Driver) TxDMA() dma.Channel {
	return d.txDMA
}

func (d *Driver) RxDMA() dma.Channel {
	return d.rxDMA
}

func setupDMA(ch dma.Channel, mode dma.Mode, paddr uintptr) {
	ch.Setup(mode)
	ch.SetWordSize(1, 1)
	ch.SetAddrP(unsafe.Pointer(paddr))
}

func startDMA(ch dma.Channel, maddr uintptr, mlen int, irq bool) {
	ch.SetAddrM(unsafe.Pointer(maddr))
	ch.SetLen(mlen)
	ch.Clear(dma.EvAll, dma.ErrAll)
	if irq {
		ch.EnableIRQ(dma.Complete, dma.ErrAll&^dma.ErrFIFO)
	}
	mmio.MB() // ensure the data in RAM are ready; BUG: not enough with D-cache
	ch.Enable()
}
