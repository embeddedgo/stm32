// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spih

import (
	"embedded/rtos"
	"runtime"
	"sync"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
)

// A Master is a driver to the SPI peripheral used in master mode.
type Master struct {
	sync.Mutex // helps in case of concurent use of Master (not used internally)

	p    *Periph
	slow bool

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

// Periph returns the underlying SPI peripheral.
func (d *Master) Periph() *Periph {
	return d.p
}

func (d *Master) Enable() {
	p := d.p
	cr1 := p.CR1.Load()
	cr1 |= SPE | MASRX
	p.CR1.Store(SPE)
	cr1 |= CSTART
	p.CR1.Store(cr1)
}

// Disable waits for the last bit of the last transfer to be sent and next it
// disables the SPI peripheral.
func (d *Master) Disable() {
	d.WaitTxDone()
	d.p.CR1.ClearBits(SPE)
}

func (d *Master) WaitTxDone() {
	p, slow := d.p, d.slow
	for p.SR.LoadBits(TXC) == 0 {
		if slow {
			runtime.Gosched()
		}
	}
}

type Config uint64

const (
	FrameFormat = Config(sp) // Frame format:
	MSPI        = Config(0)  // - Motorola SPI
	SyncSerial  = Config(ti) // - Texas Instruments Synchronous Serial

	CPHA0 = Config(0)    // Sample on leading edge.
	CPHA1 = Config(cpha) // Sample on trailing edge.
	CPOL0 = Config(0)    // Clock idle state is 0.
	CPOL1 = Config(cpol) // Clock idle state is 1.

	WordLen = Config(wlen) // Data word length (no support for SPIv3 17-24 bit):
	Word4b  = Config(w4)   // - 4 bit (SPIv2+)
	Word5b  = Config(w5)   // - 5 bit (SPIv2+)
	Word6b  = Config(w6)   // - 6 bit (SPIv2+)
	Word7b  = Config(w7)   // - 7 bit (SPIv2+)
	Word8b  = Config(w8)   // - 8 bit
	Word9b  = Config(w9)   // - 9 bit (SPIv2+)
	Word10b = Config(w10)  // - 10 bit (SPIv2+)
	Word11b = Config(w11)  // - 11 bit (SPIv2+)
	Word12b = Config(w12)  // - 12 bit (SPIv2+)
	Word13b = Config(w13)  // - 13 bit (SPIv2+)
	Word14b = Config(w14)  // - 14 bit (SPIv2+)
	Word15b = Config(w15)  // - 15 bit (SPIv2+)
	Word16b = Config(w16)  // - 16 bit
	Word17b = Config(w17)  // - 17 bit (SPIv3)
	Word18b = Config(w18)  // - 18 bit (SPIv3)
	Word19b = Config(w19)  // - 19 bit (SPIv3)
	Word20b = Config(w20)  // - 20 bit (SPIv3)
	Word21b = Config(w21)  // - 21 bit (SPIv3)
	Word22b = Config(w22)  // - 22 bit (SPIv3)
	Word23b = Config(w23)  // - 23 bit (SPIv3)
	Word24b = Config(w24)  // - 24 bit (SPIv3)
	Word25b = Config(w25)  // - 25 bit (SPIv3)
	Word26b = Config(w26)  // - 26 bit (SPIv3)
	Word27b = Config(w27)  // - 27 bit (SPIv3)
	Word28b = Config(w28)  // - 28 bit (SPIv3)
	Word29b = Config(w29)  // - 29 bit (SPIv3)
	Word30b = Config(w30)  // - 30 bit (SPIv3)
	Word31b = Config(w31)  // - 31 bit (SPIv3)
	Word32b = Config(w32)  // - 32 bit (SPIv3)

	ClkDiv    = Config(clkDiv)      // Clock divider (see also SetBaudrate):
	ClkDiv2   = Config(0 * clkDiv4) // - baud rate = ker_ck/2
	ClkDiv4   = Config(1 * clkDiv4) // - baud rate = ker_ck/4
	ClkDiv8   = Config(2 * clkDiv4) // - baud rate = ker_ck/8
	ClkDiv16  = Config(3 * clkDiv4) // - baud rate = ker_ck/16
	ClkDiv32  = Config(4 * clkDiv4) // - baud rate = ker_ck/32
	ClkDiv64  = Config(5 * clkDiv4) // - baud rate = ker_ck/64
	ClkDiv128 = Config(6 * clkDiv4) // - baud rate = ker_ck/128
	ClkDiv256 = Config(7 * clkDiv4) // - baud rate = ker_ck/256

	MSBF = Config(0)        // Most significant bit first.
	LSBF = Config(lsbfirst) // Least significant bit first.

	Duplex     = Config(dir) // Tranfer mode:
	FullDuplex = Config(0)   // - full-duplex
	HalfDuplex = Config(hd)  // - half-duplex (MOSI is used for Tx and Rx)
)

func (d *Master) ClkDiv(baudrate int) Config {
	return configClkDiv(kernHz(d.p), baudrate)
}

func (d *Master) Baudrate() int {
	return baudrate(d.p)
}

// Config returns the current configuration of the SPI peripheral.
func (d *Master) Config() Config {
	return config(d.p)
}

// SetConfig configures p.
func (d *Master) SetConfig(cfg Config) {
	d.WaitTxDone()
	p := d.p
	cr1 := p.CR1.Load()
	if cr1&SPE != 0 {
		p.CR1.Store(cr1 &^ SPE)
	}
	setConfig(p, cfg)
	if cr1&SPE != 0 {
		p.CR1.Store(cr1)
	}

}

// Setup resets the underlying SPI peripheral and configures it according to
// the master driver needs. Next it calls d.SetConfig(cfg) and
// d.SetBaudrate(baudrate) if baudrate>0.
func (d *Master) Setup(cfg Config, baudrate int) (actualBaud int) {
	p := d.p
	p.EnableClock(true)
	p.Reset()
	kernHz := kernHz(p)
	if baudrate > 0 {
		cfg = cfg&^ClkDiv | configClkDiv(kernHz, baudrate)
	}
	d.SetConfig(cfg)
	return int(kernHz >> (cfg&ClkDiv/ClkDiv4 + 1))
}
