// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package usart

import (
	"runtime"
	"sync/atomic"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
)

// EnableRx enables the UART receiver. If rxbuf is not nil the Driver uses the
// provided slice to buffer received data. Othewrise it allocates a small buffer
// itself. At least 2-byte buffer is required. EnableRx setups Tx DMA channel
// in circular mode and enables it to continuous reception of data. Driver
// assumes that it has exclusive access to the underlying USART peripheral and
// Rx DMA channel between EnableRx and DisableRx.
func (d *Driver) EnableRx(rxbuf []byte) {
	if d.rxBuf != nil {
		panic("enabled before")
	}
	if rxbuf == nil {
		rxbuf = make([]byte, 64)
	} else if len(rxbuf) < 2 {
		panic("rxbuf too short")
	}
	d.rxBuf = rxbuf
	ch := d.rxDMA
	d.p.cr1.SetBits(re)
	d.p.cr3.SetBits(dmar)
	setupDMA(ch, dma.PTM|dma.IncM|dma.Circ, rdr(d.p).Addr())
	startDMA(ch, uintptr(unsafe.Pointer(&d.rxBuf[0])), len(d.rxBuf), dma.Complete)
}

// DisableRx disables USART receiver and resets the state of the internal ring
// buffer.
func (d *Driver) DisableRx() (rxbuf []byte) {
	ch := d.rxDMA
	ch.Disable()
	ch.DisableIRQ(dma.EvAll, dma.ErrAll)
	d.p.cr1.ClearBits(re)
	d.p.cr3.ClearBits(dmar)
	d.rxH = 0
	d.rxL = 0
	// wait for DMA to really stop
	for ch.Enabled() {
		runtime.Gosched()
	}
	d.dmaH = 0
	rxbuf = d.rxBuf
	d.rxBuf = nil
	return
}

func (d *Driver) dmaHL() (h, l uint32) {
	ch := d.rxDMA
	h = atomic.LoadUint32(&d.dmaH)
	for {
		cl := ch.Len()
		nh := atomic.LoadUint32(&d.dmaH)
		if h == nh {
			return h, uint32(len(d.rxBuf) - cl)
		}
		h = nh
	}
}

func (d *Driver) rxHLadd(n int) {
	d.rxL += uint32(n)
	if d.rxL >= uint32(len(d.rxBuf)) {
		d.rxL -= uint32(len(d.rxBuf))
		d.rxH++
	}
}

func (d *Driver) disableRxIRQ() {
	d.p.DisableIRQ(RxNotEmpty)
	d.p.Clear(RxNotEmpty, 0)
	d.p.DisableErrIRQ()
}

func (d *Driver) Read(buf []byte) (int, error) {
start:
	dmaH, dmaL := d.dmaHL()
	switch dmaH - d.rxH {
	case 0:
		if dmaL == d.rxL {
			// No data in ring buffer. Wait for RxNotEmpty or error (DMA IRQs
			// are useless because they can only signal half or full transfer.
			d.rxReady.Clear()
			d.p.EnableIRQ(RxNotEmpty)
			d.p.EnableErrIRQ()
			dmaH, dmaL = d.dmaHL()
			if dmaL != d.rxL || dmaH != d.rxH {
				d.disableRxIRQ()
				goto start
			}
			if !d.rxReady.Sleep(d.timeoutRx) {
				return 0, ErrTimeout
			}
			if _, e := d.p.Status(); e != 0 {
				// clear errors
				d.p.Load()      // for older MCUs (complete read SR then DR seq)
				d.p.Clear(0, e) // for newer MCUs.
				return 0, e
			}
			if _, e := d.rxDMA.Status(); e != 0 {
				return 0, e
			}
			goto start
		}
		if dmaL == 0 {
			// Belated RxDMAISR: dmaHL read NDTR just after it was reloaded and
			// before TC interrupt was taken.
			dmaL = uint32(len(d.rxBuf))
		}
		n := copy(buf, d.rxBuf[d.rxL:dmaL])
		d.rxHLadd(n)
		return n, nil
	case 1:
		if dmaL > d.rxL {
			break
		}
		n := copy(buf, d.rxBuf[d.rxL:])
		if n < len(buf) {
			n += copy(buf[n:], d.rxBuf[:dmaL])
		}
		dmaH, dmaL = d.dmaHL()
		if dmaH-d.rxH != 1 || dmaL > d.rxL {
			// There is no certainty that we managed to copy before overwriting.
			break
		}
		d.rxHLadd(n)
		return n, nil
	}
	d.rxH = dmaH
	d.rxL = dmaL
	return 0, ErrBufOverflow
}

func (d *Driver) RxDMAISR() {
	ch := d.rxDMA
	ev, err := ch.Status()
	if err != 0 {
		ch.DisableIRQ(dma.EvAll, dma.ErrAll)
		d.rxReady.Wakeup()
		return
	}
	if ev&dma.Complete != 0 {
		ch.Clear(dma.Complete, 0)
		atomic.StoreUint32(&d.dmaH, d.dmaH+1)
	}
}

func (d *Driver) ISR() {
	d.disableRxIRQ()
	d.rxReady.Wakeup()
}
