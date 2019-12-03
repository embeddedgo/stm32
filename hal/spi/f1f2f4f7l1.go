// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407

package spi

import (
	"github.com/embeddedgo/stm32/p/spi"
)

const cr1Mask = ^spi.CR1(spi.DFF | spi.SPE | spi.BIDIMODE | spi.BIDIOE)

func (p *Periph) setWordSize(size int) {
	if size == 16 {
		p.raw.DFF().Set()
	} else {
		p.raw.DFF().Clear()
	}
}

func (p *Periph) wordSize() int {
	if p.raw.DFF().Load() != 0 {
		return 16
	}
	return 8
}
