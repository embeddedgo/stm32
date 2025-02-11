// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usart

import (
	"runtime"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/hal/mem/nocache"
)

// EnableRx enables the UART receive part. It allocates an internal ring buffer
// of bufLen size. In most cases bufLen=64 is good choise (minimum is 2). Use
// bufLen=0 to reenable the UART with the previous buffer size.
//
// EnableRx setups Rx DMA channel in circular mode and enables it to continuous
// reception of data. Driver assumes that it has exclusive access to the Rx part
// of the underlying USART peripheral and the Rx DMA channel between EnableRx
// and DisableRx.
//
// The Rx algorithm and the DMA configuration is optimized for speed. The DMA
// copies received data to the internal ring buffer. If there is no room for new
// data the DMA overwrites the oldest ones. The detection of buffer overflow is
// weak (some cases may not be detected).
//
// You can not rely on the error reporting by read methods at all. The detected
// UART errors are asynchronous with the received data and can be used only as
// an indicator of poor connection quality. The ErrBufOverflow error indicates
// that the rxbuf is too small or the data processing is too slow. You cannot
// determine which data has been affected by an error (use other techniques to
// ensure data integrity).
//
// As the DMA is usded to receive data it does not make much sense to enable
// hardware RTS signaling unless the DMA is very busy.
func (d *Driver) EnableRx(bufLen int) {
	if len(d.rxBuf) != 0 {
		panic("enabled before")
	}
	if cap(d.rxBuf) >= bufLen {
		if bufLen == 0 {
			bufLen = cap(d.rxBuf)
		} else if bufLen < 2 {
			bufLen = 2
		}
		d.rxBuf = d.rxBuf[:bufLen]
	} else {
		if bufLen < 2 {
			bufLen = 2
		}
		d.rxBuf = nocache.MakeSlice[byte](1, bufLen, bufLen)
	}
	p := d.p
	ch := d.rxDMA
	internal.ExclusiveStoreBits(&p.cr1, ue|re, ue|re)
	internal.ExclusiveStoreBits(&p.cr3, dmar, dmar)
	setupDMA(ch, dma.PTM|dma.IncM|dma.Circ|dma.TrBuf, rdr(p).Addr())
	startDMA(ch, uintptr(unsafe.Pointer(&d.rxBuf[0])), len(d.rxBuf), false)
}

// DisableRx disables USART receiver and frees memory allocated for the internal
// ring buffer.
func (d *Driver) DisableRx() {
	ch := d.rxDMA
	ch.Disable()
	ch.DisableIRQ(dma.EvAll, dma.ErrAll)
	internal.ExclusiveStoreBits(&d.p.cr3, dmar, 0)
	internal.ExclusiveStoreBits(&d.p.cr1, re, 0)
	d.rxBuf = d.rxBuf[:0]
	d.rxp = 0
	// wait for DMA to really stop
	for ch.Enabled() {
		runtime.Gosched()
	}
	return
}

// DiscardRx discards all rceived data.
func (d *Driver) DiscardRx() {
	rxBuf := d.rxBuf
	d.DisableRx()
	d.rxBuf = rxBuf
	d.Periph().Status()
	d.Periph().Load()
	d.EnableRx(-len(rxBuf))
}

// dmaPos returns the current Rx DMA position, tc extends the positon
// information, e.g. pos=1 and tc=true is after pos=9 and tc=false.
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

// Read implements io.Reader interface.
func (d *Driver) Read(buf []byte) (n int, err error) {
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
		// wait for the DMA to read any data to the d.rxBuf
		for {
			dmap, tc, e = dmaPos(d)
			if e != 0 {
				return 0, e
			}
			if dmap != d.rxp {
				goto again
			}
		}
	default: // DMA position is in between d.rxp and the end of d.rxBuf
		n = copy(buf, d.rxBuf[d.rxp:dmap])
	}
	d.rxp += n
	if d.rxp >= len(d.rxBuf) {
		d.rxp -= len(d.rxBuf)
		d.rxDMA.Clear(dma.Complete, 0)
	}
	return
}

func (d *Driver) ReadByte() (b byte, err error) {
	var buf [1]byte
	if _, err = d.Read(buf[:]); err == nil {
		b = buf[0]
	}
	return
}

// RxISR is an interrupt handler for the UART interrupt. The UART interrupt is
// only need for reception. No other interrupt is need for Rx-only use case.
func (d *Driver) RxISR() {
	d.disableRxIRQ()
	d.rxReady.Wakeup()
}
