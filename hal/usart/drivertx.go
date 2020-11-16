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
	setupDMA(d.txDMA, dma.MTP|dma.IncM|dma.FT4, tdr(d.p).Addr())
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

// WriteString works like Write but accepts string instead of byte slice.
func (d *Driver) WriteString(s string) (int, error) {
	if rtos.HandlerMode() {
		return sysWrite(d, s)
	}
	ch := d.txDMA
	ptr := *(*uintptr)(unsafe.Pointer(&s))
	var n int
	for {
		m := len(s) - n
		if m == 0 {
			break
		}
		if m > 0xffff {
			m = 0xffff
		}
		d.txDone.Clear()
		clear(d.p, TxDone, 0) // Clear TC.
		startDMA(ch, ptr+uintptr(n), m, true)
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
	return n, nil
}

// Write sends bytes from p to the remote party. It returns the number of bytes
// sent and error if detected. It does not provide any guarantee that the bytes
// sent were received by the remote party.
func (d *Driver) Write(p []byte) (int, error) {
	return d.WriteString(*(*string)(unsafe.Pointer(&p)))
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
func sysWrite(d *Driver, s string) (int, error) {
	d.txDMA.DisableIRQ(dma.EvAll, dma.ErrAll)
	d.txDMA.Disable()
	for i := 0; i < len(s); i++ {
		c := int(s[i])
		if c == '\n' {
			// convert "\n" to "\r\n"
			waitTxEmpty(d.p)
			d.p.Store('\r')
		}
		waitTxEmpty(d.p)
		d.p.Store(c)
	}
	return len(s), nil
}

func waitTxEmpty(p *Periph) {
	for {
		ev, _ := p.Status()
		if ev&TxEmpty != 0 {
			break
		}
	}
}
