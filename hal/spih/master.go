// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package spih

import (
	"embedded/rtos"
	"sync"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
)

// A Master is a driver to the SPI peripheral used in master mode.
type Master struct {
	sync.Mutex // helps in case of concurent use of Master (not used internally)

	p     *Periph
	slow  bool
	wonly bool
	repw  uint16

	rdma, wdma dma.Channel

	done rtos.Note
}

// NewMaster returns a new master-mode driver for p. If valid DMA channels are
// given, the DMA will be used for bigger data transfers.
func NewMaster(p *Periph, rdma, wdma dma.Channel) *Master {
	d := &Master{p: p, rdma: rdma, wdma: wdma}
	if rdma.IsValid() {
		rdma.SetAddrP(unsafe.Pointer(&p.RXDR))
	}
	if wdma.IsValid() {
		wdma.SetAddrP(unsafe.Pointer(&p.TXDR))
	}
	return d
}
