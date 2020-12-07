// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407

package spi

import "github.com/embeddedgo/stm32/p/spi"

func (p *Periph) setWordSize(size int) {
	if size == 16 {
		p.raw.CR1.SetBits(spi.DFF)
	} else {
		p.raw.CR1.ClearBits(spi.DFF)
	}
}

func (p *Periph) wordSize() int {
	if p.raw.CR1.LoadBits(spi.DFF) != 0 {
		return 16
	}
	return 8
}
