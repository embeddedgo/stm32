// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usart

import (
	"runtime"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
)

// EnableRx enables the UART receiver. If rxbuf is not nil the Driver uses the
// provided slice to buffer received data. Othewrise it allocates a small buffer
// itself. At least 2-byte buffer is required.
//
// EnableRx setups Rx DMA channel in circular mode and enables it to continuous
// reception of data. Driver assumes that it has exclusive access to the
// underlying USART peripheral and Rx DMA channel between EnableRx and
// DisableRx.
//
// The Rx algorithm and the DMA configuration is optimized for speed. The DMA is
// allowed to copy any received data to the internal ring buffer. If there is
// no room for new data the DMA overwrites the oldest. The detection of buffer
// overflow is weak (some cases may not be detected).
//
// You can not rely on the error reporting at all. The detected UART errors are
// asynchronous with the received data and can be used only as an indicator of
// poor connection quality. The ErrBufOverflow indicates that the rxbuf is too
// small or the data processing is too slow. You cannot determine which data has
// been affected (use other techniques to ensure data integrity).
//
// As the DMA reads any received data it does not make much sense to enable
// hardware RTS signaling unless the DMA is very busy.
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
	startDMA(ch, uintptr(unsafe.Pointer(&d.rxBuf[0])), len(d.rxBuf), false)
}

// DisableRx disables USART receiver and resets the state of the internal ring
// buffer.
func (d *Driver) DisableRx() (rxbuf []byte) {
	ch := d.rxDMA
	ch.Disable()
	ch.DisableIRQ(dma.EvAll, dma.ErrAll)
	d.p.cr1.ClearBits(re)
	d.p.cr3.ClearBits(dmar)
	rxbuf = d.rxBuf
	d.rxBuf = nil
	d.rxp = 0
	// wait for DMA to really stop
	for ch.Enabled() {
		runtime.Gosched()
	}
	return
}

func dmaPos(d *Driver) (pos int, tc bool, e dma.Error) {
	ev, err := d.rxDMA.Status()
	if err != 0 {
		goto dmaErr
	}
	pos = d.rxDMA.Len()
	if tc = ev&dma.Complete != 0; tc {
		goto end
	}
	ev, err = d.rxDMA.Status()
	if err != 0 {
		goto dmaErr
	}
	if tc = ev&dma.Complete != 0; tc {
		pos = d.rxDMA.Len()
	}
end:
	return len(d.rxBuf) - pos, tc, 0
dmaErr:
	d.rxDMA.Clear(0, err)
	return -1, false, err
}

func (d *Driver) disableRxIRQ() {
	d.p.DisableIRQ(RxNotEmpty)
	d.p.Clear(RxNotEmpty, 0)
	d.p.DisableErrIRQ()
}

func (d *Driver) Read(buf []byte) (n int, err error) {
start:
	dmap, tc, e := dmaPos(d)
	if e != 0 {
		return 0, e
	}
again:
	// check UART Rx errors
	if _, e := d.p.Status(); e != 0 {
		// clear errors
		d.p.Load()      // for older MCUs (complete read SR then DR seq)
		d.p.Clear(0, e) // for newer MCUs.
		return 0, e
	}
	switch {
	case tc: // DMA reached the end of d.rxBuf and started from the begginnning
		// weak overflow detection
		if dmap > d.rxp {
			d.rxp = dmap
			err = ErrBufOverflow
		}
		n = copy(buf, d.rxBuf[d.rxp:])
		if n < len(buf) {
			n += copy(buf[n:], d.rxBuf[:dmap])
		}
		d.rxDMA.Clear(dma.Complete, 0)
	case dmap == d.rxp: // no new data in the ring buffer
		// wait for RxNotEmpty event or error
		d.rxReady.Clear()
		d.p.EnableIRQ(RxNotEmpty)
		d.p.EnableErrIRQ()
		dmap, tc, e = dmaPos(d) // check one more time before sleep
		if e != 0 {
			d.disableRxIRQ()
			return 0, e
		}
		if tc || dmap != d.rxp {
			d.disableRxIRQ()
			goto again
		}
		if !d.rxReady.Sleep(d.timeoutRx) {
			return 0, ErrTimeout
		}
		goto start
	default: // DMA position is in between d.rxp and the end of d.rxBuf
		n = copy(buf, d.rxBuf[d.rxp:dmap])
	}
	d.rxp += n
	if d.rxp >= len(d.rxBuf) {
		d.rxp -= len(d.rxBuf)
	}
	return
}

// RxISR is an interrupt handler for the UART interrupt. The UART interrupt is
// only need for reception. No other interrupt is need for Rx-only use case.
func (d *Driver) RxISR() {
	d.disableRxIRQ()
	d.rxReady.Wakeup()
}
