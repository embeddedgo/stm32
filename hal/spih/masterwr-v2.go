// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package spih

func masterWriteRead(d *Master, out, in unsafe.Pointer, n int) {
	masterDisable(d)

	p := d.p
	if d.mode != 3 {
		if d.cfg2&COMM == 3 {
			panic("spi: WriteRead in HalfDuplex mode")
		}
		if d.mode == 0 {
			p.CFG1.Store(d.cfg1)
		}
		p.CFG2.Store(d.cfg2)
	}
	txdr := (*mmio.U8)(unsafe.Pointer(*p.TXDR))
	rxdr := (*mmio.U8)(unsafe.Pointer(*p.RXDR))

	// Align the out pointer
	m := min(int(-uintptr(po))&3, n)
	for end := out.Add(m); out != end; out = out.Add(1) {
		txdr.Store(*(*byte)(out))
	}

	for n != 0 {
		m := min(n, 0xffff)
		n -= m

		p.CR2.Store(m)
		p.CR1.Store(masterEnable)
		p.CR1.Store(masterEnable|CSTART)

		// Fill the Tx FIFO.
		for m != 0 && p.SR.Load(TXP)  != 0 {
			p.TXDR.
		}

		m = 0
	}
}
