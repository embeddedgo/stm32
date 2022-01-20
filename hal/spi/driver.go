// Copyright 2019 The Embedded Go authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spi

import (
	"embedded/rtos"
	"runtime"
	"sync/atomic"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
)

// Driver implements DMA based driver to an SPI peripheral working in master
// mode.
type Driver struct {
	p      *Periph
	rxDMA  dma.Channel
	txDMA  dma.Channel
	done   rtos.Note
	dmacnt int32
	err    uint32
}

// NewDriver provides convenient way to create heap allocated Driver struct.
func NewDriver(p *Periph, rxdma, txdma dma.Channel) *Driver {
	return &Driver{p: p, rxDMA: rxdma, txDMA: txdma}
}

// Periph returns the SPI peripheral used by Driver.
func (d *Driver) Periph() *Periph {
	return d.p
}

// Setup enables clock source, resets peripheral, then calls
// d.Periph().SetConfig(conf, baudrate) to configure it.
func (d *Driver) Setup(conf Config, baudrate int) {
	d.p.EnableClock(true)
	d.p.Reset()
	d.p.SetConfig(conf, baudrate)
}

// SetBaudrate calls d.Periph().SetBaudrate(baudrate).
func (d *Driver) SetBaudrate(baudrate int) {
	d.p.SetBaudrate(baudrate)
}

// SetWordSize calls d.Periph().SetWordSize(size).
func (d *Driver) SetWordSize(size int) {
	d.p.setWordSize(size)
}

// Enable calls d.Periph().Enable().
func (d *Driver) Enable() {
	d.p.Enable()
}

// Disable waits for the peripheral idle state and calls d.Periph().Disable().
func (d *Driver) Disable() {
	for {
		if ev, _ := d.p.Status(); ev == TxEmpty {
			break
		}
		runtime.Gosched()
	}
	d.p.Disable()
}

// TwoWireSetRx can be used in two-wire mode to set data direction to recevie.
func (d *Driver) TwoWireSetRx() {
	d.p.EditConfig(Tx, Rx)
}

// TwoWireSetRx can be used in two-wire mode to set data direction to transimit.
func (d *Driver) TwoWireSetTx() {
	d.p.EditConfig(Rx, Tx)
}

func (d *Driver) RxDMA() dma.Channel {
	return d.rxDMA
}

func (d *Driver) TxDMA() dma.Channel {
	return d.txDMA
}

func (d *Driver) DMAISR(ch dma.Channel) {
	ev, err := ch.Status()
	if ev&dma.Complete != 0 || err&^dma.ErrFIFO != 0 {
		ch.DisableIRQ(dma.EvAll, dma.ErrAll)
		if atomic.AddInt32(&d.dmacnt, -1) == 0 {
			d.done.Wakeup()
		}
	}
}

func (d *Driver) ISR() {
	d.p.DisableIRQ(Err)
	d.done.Wakeup()
}

// WriteReadByte writes and reads byte.
func (d *Driver) WriteReadByte(b byte) byte {
	if d.err != 0 {
		return 0
	}
	d.p.EditConfig(TwoWire|Tx, ThreeWire|TxRx)
	d.p.StoreByte(b)
	// wait in loop, interrupts are too expensive for one byte transfer
	for {
		ev, err := d.p.Status()
		if ev&RxNotEmpty != 0 {
			break
		}
		if err != 0 {
			d.err = uint32(err) << 8
			return 0
		}
	}
	return d.p.LoadByte()
}

// WriteReadWord16 writes and reads 16-bit word.
func (d *Driver) WriteReadWord16(w uint16) uint16 {
	if d.err != 0 {
		return 0
	}
	d.p.EditConfig(TwoWire|Tx, ThreeWire|TxRx)
	d.p.StoreWord16(w)
	// wait in loop, interrupts are too expensive for one word transfer
	for {
		ev, err := d.p.Status()
		if ev&RxNotEmpty != 0 {
			break
		}
		if err != 0 {
			d.err = uint32(err) << 8
			return 0
		}
	}
	return d.p.LoadWord16()
}

func (d *Driver) setupDMA(ch dma.Channel, mode dma.Mode, wordSize uintptr) {
	ch.Setup(mode)
	ch.SetWordSize(wordSize, wordSize)
	ch.SetAddrP(unsafe.Pointer(&d.p.dr))
}

const minIRQLen = 8

func startDMA(ch dma.Channel, addr uintptr, n int) {
	ch.SetAddrM(unsafe.Pointer(addr))
	ch.SetLen(n)
	ch.Clear(dma.EvAll, dma.ErrAll)
	if n >= minIRQLen {
		ch.EnableIRQ(dma.Complete, dma.ErrAll&^dma.ErrFIFO)
	}
	ch.Enable()
}

func waitDMA(d *Driver, n int, rx bool) {
	if n >= minIRQLen {
		d.done.Sleep(-1)
	}
	tx := true
	for {
		if _, err := d.p.Status(); err != 0 {
			d.err |= uint32(err) << 8
			d.txDMA.DisableIRQ(dma.EvAll, dma.ErrAll)
			d.rxDMA.DisableIRQ(dma.EvAll, dma.ErrAll)
			d.txDMA.Disable()
			d.rxDMA.Disable()
			tx, rx = false, false
		}
		if tx {
			ev, err := d.txDMA.Status()
			err &^= dma.ErrFIFO
			if ev&dma.Complete != 0 || err != 0 {
				d.err |= uint32(err)
				tx = false
				d.txDMA.Disable() // required by non-stream DMA (F0,F1,F3,L1,L4)
			}
		}
		if rx {
			ev, err := d.rxDMA.Status()
			err &^= dma.ErrFIFO
			if ev&dma.Complete != 0 || err != 0 {
				d.err |= uint32(err)
				rx = false
				d.rxDMA.Disable() // required by non-stream DMA (F0,F1,F3,L1,L4)
			}
		}
		if !(tx || rx) {
			break
		}
		runtime.Gosched()
	}
}

func (d *Driver) writeReadDMA(out, in uintptr, olen, ilen int, wsize uintptr) (n int) {
	txdmacfg := dma.MTP | dma.FT4
	if olen > 1 {
		txdmacfg |= dma.IncM
	}
	d.setupDMA(d.txDMA, txdmacfg, wsize)
	d.setupDMA(d.rxDMA, dma.PTM|dma.IncM|dma.FT4, wsize)
	p := d.p
	p.EditConfig(TwoWire|Tx, ThreeWire|TxRx)
	p.EnableDMA(RxNotEmpty | TxEmpty)
	p.EnableIRQ(Err)
	for {
		m := ilen - n
		if m == 0 {
			break
		}
		if m > 0xffff {
			m = 0xffff
		}
		if m >= minIRQLen {
			d.dmacnt = 2
			d.done.Clear()
		}
		startDMA(d.rxDMA, in, m)
		startDMA(d.txDMA, out, m)
		if olen > 1 {
			out += uintptr(m)
		}
		in += uintptr(m)
		n += m
		waitDMA(d, m, true)
		if d.err != 0 {
			n -= d.rxDMA.Len()
			break
		}
	}
	p.DisableDMA(RxNotEmpty | TxEmpty)
	p.DisableIRQ(Err)
	return
}

func (d *Driver) writeDMA(out uintptr, n int, wsize uintptr, incm dma.Mode) {
	d.setupDMA(d.txDMA, dma.MTP|incm|dma.FT4, wsize)
	p := d.p
	p.EditConfig(ThreeWire|TxRx, TwoWire|Tx) // avoid ErrOverflow
	p.EnableDMA(TxEmpty)
	p.EnableIRQ(Err)
	for n > 0 {
		m := n
		if m > 0xffff {
			m = 0xffff
		}
		if m >= minIRQLen {
			d.dmacnt = 1
			d.done.Clear()
		}
		startDMA(d.txDMA, out, m)
		n -= m
		if incm != 0 {
			out += uintptr(m)
		}
		waitDMA(d, m, false)
		if d.err != 0 {
			break
		}
	}
	p.DisableDMA(TxEmpty)
	p.DisableIRQ(Err)
	// DMA finished but SPI can still send buffered data. Wait for end.
	for {
		if ev, _ := p.Status(); ev&Busy == 0 {
			break
		}
		runtime.Gosched()
	}
}

// Err returns the saved error and clears it if the clear is true. If an error
// occurs during any command it is saved and the subsequent commands are ignored
// until the error is cleared.
func (d *Driver) Err(clear bool) error {
	e := d.err
	if e == 0 {
		return nil
	}
	if clear {
		d.err = 0
	}
	if err := Error(e >> 8); err != 0 {
		if err&ErrOverrun != 0 && clear {
			d.p.LoadByte()
			d.p.Status()
		}
		return err
	}
	return dma.Error(e)
}

func (d *Driver) writeRead(oaddr, iaddr uintptr, olen, ilen int, wsize uintptr) int {
	if olen > ilen {
		var n int
		if ilen > 0 {
			n = d.writeReadDMA(oaddr, iaddr, ilen, ilen, wsize)
			if d.err != 0 {
				return n
			}
			olen -= ilen
			oaddr += uintptr(ilen)
		}
		d.writeDMA(oaddr, olen, wsize, dma.IncM)
		return n
	}
	if ilen > olen {
		var n int
		ffff := uint16(0xffff)
		if olen > 0 {
			n = d.writeReadDMA(oaddr, iaddr, olen, olen, wsize)
			if d.err != 0 {
				return n
			}
			ilen -= olen
			iaddr += uintptr(olen)
			oaddr += uintptr(olen - 1)
		} else {
			oaddr = uintptr(unsafe.Pointer(&ffff))
		}
		return n + d.writeReadDMA(oaddr, iaddr, 1, ilen, wsize)
	}
	return d.writeReadDMA(oaddr, iaddr, ilen, ilen, wsize)
}

func (d *Driver) WriteStringRead(out string, in []byte) int {
	olen := len(out)
	ilen := len(in)
	if d.err != 0 || olen == 0 && ilen == 0 {
		return 0
	}
	if olen <= 1 && ilen <= 1 {
		// Avoid DMA for one byte transfers.
		b := byte(0xff)
		if olen != 0 {
			b = out[0]
		}
		b = d.WriteReadByte(b)
		if ilen != 0 {
			in[0] = b
			return 1
		}
		return 0
	}
	oaddr := *(*uintptr)(unsafe.Pointer(&out))
	iaddr := *(*uintptr)(unsafe.Pointer(&in))
	return d.writeRead(oaddr, iaddr, olen, ilen, 1)
}

func (d *Driver) WriteRead(out, in []byte) int {
	return d.WriteStringRead(*(*string)(unsafe.Pointer(&out)), in)
}

func (d *Driver) WriteByteN(b byte, n int) {
	if d.err != 0 {
		return
	}
	switch {
	case n > 1:
		d.writeDMA(uintptr(unsafe.Pointer(&b)), n, 1, 0)
	case n == 1:
		// Avoid DMA for one byte transfers.
		d.WriteReadByte(b)
	}
}

func (d *Driver) WriteRead16(out, in []uint16) int {
	olen := len(out)
	ilen := len(in)
	if d.err != 0 || olen == 0 && ilen == 0 {
		return 0
	}
	if olen <= 1 && ilen <= 1 {
		// Avoid DMA for one word transfers.
		w := uint16(0xffff)
		if olen != 0 {
			w = out[0]
		}
		w = d.WriteReadWord16(w)
		if ilen != 0 {
			in[0] = w
			return 1
		}
		return 0
	}
	oaddr := *(*uintptr)(unsafe.Pointer(&out))
	iaddr := *(*uintptr)(unsafe.Pointer(&in))
	return d.writeRead(oaddr, iaddr, olen, ilen, 2)
}

func (d *Driver) WriteWord16N(w uint16, n int) {
	if d.err != 0 {
		return
	}
	switch {
	case n > 1:
		d.writeDMA(uintptr(unsafe.Pointer(&w)), n, 2, 0)
	case n == 1:
		// Avoid DMA for one word transfers.
		d.WriteReadWord16(w)
	}
}
