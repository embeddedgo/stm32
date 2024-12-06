// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usart

import (
	"embedded/rtos"
	"runtime"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
)

// EnableTx enables Tx part of the USART peripheral and setups Tx DMA channel.
// Driver assumes that it has exclusive access to the underlying USART
// peripheral and Tx DMA channel between EnableTx and DisableTx.
func (d *Driver) EnableTx() {
	d.p.cr1.SetBits(te)
	d.p.cr3.SetBits(dmat)
	setupDMA(d.txDMA, dma.MTP|dma.IncM|dma.FT4|dma.TrBuf, tdr(d.p).Addr())
}

// DisableTx disables Tx part of the USART peripheral.
func (d *Driver) DisableTx() {
	for {
		if ev, _ := d.p.Status(); ev&TxDone != 0 {
			break
		}
		runtime.Gosched()
	}
	d.p.cr1.ClearBits(te)
}

func waitTxNotFull(p *Periph) {
	for {
		ev, _ := p.Status()
		if ev&TxNotFull != 0 {
			break
		}
	}
}

// WriteByte implements io.ByteWriter interface.
func (d *Driver) WriteByte(b byte) error {
	p := d.p
	waitTxNotFull(p)
	p.Store(uint32(b))
	return nil
}

func write(p *Periph, s []byte) {
	for _, b := range s {
		waitTxNotFull(p)
		p.Store(uint32(b))
	}
}

// Write implements io.Writer interface.
func (d *Driver) Write(s []byte) (n int, err error) {
	if len(s) == 0 {
		return
	}
	if rtos.HandlerMode() {
		sysWrite(d, s)
		return len(s), nil
	}
	if len(s) >= 2*dma.MemAlign {
		p := unsafe.Pointer(&s[0])
		ds, de := dma.AlignOffsets(p, uintptr(len(s)))
		dmaStart := int(ds)
		dmaEnd := int(de)
		dmaPtr := unsafe.Add(p, ds)
		dmaN := dmaEnd - dmaStart
		if dmaStart != 0 {
			n = dmaStart
			write(d.p, s[:dmaStart])
		}
		s = s[dmaEnd:]
		ch := d.txDMA
		if dma.CacheMaint {
			rtos.CacheMaint(rtos.DCacheFlush, dmaPtr, dmaN)
		}
		up := uintptr(dmaPtr)
		for dmaN != 0 {
			m := dmaN
			if m > 0xffff {
				m = 0xffff
			}
			dmaN -= m
			d.txDone.Clear()
			clear(d.p, TxDone, 0) // Clear TC.
			startDMA(ch, up, m, true)
			up += uintptr(m)
			n += m
			done := d.txDone.Sleep(d.timeoutTx)
			ch.Disable() // to be compatible with STM32F1.
			if !done {
				ch.DisableIRQ(dma.EvAll, dma.ErrAll)
				return n - ch.Len(), ErrTimeout
			}
			// there is no USART errors for Tx, only DMA errors matter
			_, e := ch.Status()
			if e &^= dma.ErrFIFO; e != 0 {
				return n - ch.Len(), e
			}
		}
	}
	if len(s) != 0 {
		write(d.p, s)
		n += len(s)
	}
	return
}

// WriteString implements io.StringWriter interface.
func (d *Driver) WriteString(s string) (int, error) {
	return d.Write(unsafe.Slice(unsafe.StringData(s), len(s)))
}

// TxDMAISR is an interrupt handler for Tx DMA channel IRQ. The DMA interrupt is
// only needed to receive data. No other interrupt is need for Tx-only use case.
func (d *Driver) TxDMAISR() {
	ch := d.txDMA
	ev, err := ch.Status()
	if err&^dma.ErrFIFO != 0 || ev&dma.Complete != 0 {
		ch.DisableIRQ(dma.EvAll, dma.ErrAll)
		d.txDone.Wakeup()
	}
}

// sysWrite is called in handler mode. It is used by print and println mainly
// to print a stack trace before system halt so it disables DMA to interrupt any
// possible active transfer.
func sysWrite(d *Driver, s []byte) {
	d.txDMA.DisableIRQ(dma.EvAll, dma.ErrAll)
	d.txDMA.Disable()
	p := d.p
	for _, b := range s {
		if b == '\n' {
			// convert "\n" to "\r\n"
			waitTxNotFull(d.p)
			p.Store('\r')
		}
		waitTxNotFull(d.p)
		p.Store(uint32(b))
	}
}
