// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32l4x6

package spi

import "github.com/embeddedgo/stm32/p/spi"

func (p *Periph) setWordSize(size int) {
	ds := spi.CR2((size - 1) & 0xf << spi.DSn)
	if size <= 8 {
		ds |= spi.FRXTH
	}
	p.raw.CR2.StoreBits(spi.FRXTH|spi.DS, ds)
}

func (p *Periph) wordSize() int {
	return int(p.raw.CR2.LoadBits(spi.DS))>>spi.DSn&0xf + 1
}
