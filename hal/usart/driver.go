// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usart

import (
	"embedded/rtos"
	"time"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/internal"
)

type DriverError uint8

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
	rxDMA     dma.Channel
	txDMA     dma.Channel
	rxBuf     []byte // Rx ring buffer for RxDMA.
	rxReady   rtos.Note
	txDone    rtos.Note
	rxp       int
}

// NewDriver returns a new driver for p.
func NewDriver(p *Periph, rxdma, txdma dma.Channel) *Driver {
	return &Driver{
		timeoutRx: -1,
		timeoutTx: -1,
		p:         p,
		rxDMA:     rxdma,
		txDMA:     txdma,
	}
}

func (d *Driver) Periph() *Periph {
	return d.p
}

func (d *Driver) RxDMA() dma.Channel {
	return d.rxDMA
}

func (d *Driver) TxDMA() dma.Channel {
	return d.txDMA
}

// SetReadTimeout sets the read timeout used by Read* functions.
func (d *Driver) SetReadTimeout(timeout time.Duration) {
	d.timeoutRx = timeout
}

// SetWriteTimeout sets the write timeout used by Write* functions.
func (d *Driver) SetWriteTimeout(timeout time.Duration) {
	d.timeoutTx = timeout
}

type Config uint32

const (
	Word7b  = Config(W7b) // L4 or newer
	Word8b  = Config(0)
	Word9b  = Config(W9b)
	ParEven = Config(PCE)
	ParOdd  = Config(PCE | OPS)

	Stop0b5  = Config(S05) << 1
	Stop1b5  = Config(S15) << 1
	Stop2b   = Config(S2) << 1
	SwapRxTx = Config(Swap) << 1     // L4 or newer
	InvRx    = Config(RxInv) << 1    // L4 or newer
	InvTx    = Config(TxInv) << 1    // L4 or newer
	InvData  = Config(DataInv) << 1  // L4 or newer
	InvBits  = Config(MSBFirst) << 1 // L4 or newer
	AutoBR   = Config(ABREn) << 1    // L4 or newer

	IR         = Config(IREn) >> 1
	LowPowerIR = Config(IREn|IRLP) >> 1
	HalfDuplex = Config(HDSEL) >> 1
	SCNACK     = Config(NACK) >> 1
	SmartCard  = Config(SCEN) >> 1
	RTS        = Config(RTSE) >> 1
	CTS        = Config(CTSE) >> 1
)

// Enable enables UART.
func (d *Driver) Enable() {
	d.p.cr1.SetBits(ue)
}

// Disable disables UART.
func (d *Driver) Disable() {
	d.p.cr1.ClearBits(ue)
}

// SetConfig configures UART.
func (d *Driver) SetConfig(conf Config) {
	c := conf & (Word7b | Word9b | ParOdd)
	if internal.H7 {
		c |= fifoen
	}
	d.p.SetConf1(Conf1(c))
	c = conf & (Stop0b5 | Stop1b5 | Stop2b | SwapRxTx | InvRx | InvTx | InvData | InvBits | AutoBR)
	d.p.SetConf2(Conf2(c >> 1))
	c = conf & (LowPowerIR | HalfDuplex | SCNACK | SmartCard | RTS | CTS)
	d.p.SetConf3(Conf3(c << 1))
}

func (d *Driver) SetBaudrate(baudrate int) {
	d.p.SetBaudrate(baudrate)
}

// Setup enables clock source, resets UART and configures it. You still need to
// enable Tx and/or Rx before use it.
func (d *Driver) Setup(conf Config, baudrate int) {
	d.p.EnableClock(true)
	d.p.Reset()
	d.p.SetBaudrate(baudrate)
	d.SetConfig(conf)
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
	ch.Enable()
}
