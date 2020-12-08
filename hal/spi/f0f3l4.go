// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32l4x6

package spi

func (p *Periph) setWordSize(size int) {
	dsfifo := uint32(size-1) & 15 << dsn
	if size <= 8 {
		dsfifo |= frxth
	}
	p.cr2.StoreBits(frxth|ds, dsfifo)
}

func (p *Periph) wordSize() int {
	return int(p.cr2.LoadBits(ds))>>dsn&15 + 1
}
