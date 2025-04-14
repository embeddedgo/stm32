// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spih

import (
	"sync"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
)

// A Master is a driver to the SPI peripheral used in master mode.
type Master struct {
	sync.Mutex // helps in case of concurent use of Master (not used internally)

	p          *Periph
	rdma, wdma dma.Channel
	master
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

// Periph returns the underlying SPI peripheral.
func (d *Master) Periph() *Periph {
	return d.p
}

// Disable waits for the last bit of the last transfer to be sent and next it
// disables the SPI peripheral.
func (d *Master) Disable() {
	masterDisable(d)
}

func (d *Master) WaitTxDone() {
	masterWaitTxDone(d)
}

func (d *Master) ClkDiv(baudrate int) Config {
	return masterClkDiv(d.p, baudrate)
}

func (d *Master) Baudrate() int {
	return masterBaudrate(d)
}

// Config returns the current configuration used to communicate.
func (d *Master) Config() Config {
	return masterConfig(d)
}

// SetConfig sets the configuration to be used to communicate..
func (d *Master) SetConfig(cfg Config) {
	masterSetConfig(d, cfg)
}

// Setup resets the underlying SPI peripheral and configures it according to
// the master driver needs.
func (d *Master) Setup(cfg Config, baudrate int) (actualBaud int) {
	p := d.p
	p.EnableClock(true)
	p.Reset()
	if baudrate > 0 {
		cfg = cfg&^ClkDiv | masterClkDiv(p, baudrate)
	}
	masterSetConfig(d, cfg)
	return masterBaudrate(d)
}

func (d *Master) WriteRead(out, in []byte) (n int) {
	n = min(len(out), len(in))
	if n == 0 {
		return
	}
	return masterWriteRead(d, unsafe.Pointer(&out[0]), unsafe.Pointer(&in[0]), n)
}
