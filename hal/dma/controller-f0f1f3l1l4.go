// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32l4x6

package dma

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/internal"
)

type channel struct {
	cr   mmio.U32
	ndtr mmio.U32
	par  mmio.U32
	mar  mmio.U32
	_    uint32
}

type registers struct {
	isr   mmio.U32
	ifcr  mmio.U32
	c     [7]channel
	_     [5]uint32
	cselr mmio.U32
}

// cr
const (
	en   = 1 << 0
	irqs = 7 << 1
	mtp  = 1 << 4 // dir
	circ = 1 << 5
	incP = 1 << 6 // pinc
	incM = 1 << 7 // minc
	pl   = 3 << 12
	mtm  = 1 << 14 // mem2mem
)

// unsupported
const (
	ft1  = 0
	ft2  = 0
	ft3  = 0
	ft4  = 0
	pb4  = 0
	pb8  = 0
	pb16 = 0
	mb4  = 0
	mb8  = 0
	mb16 = 0
	pfc  = 0
)

func (d *Controller) channel(cn, rn int) Channel {
	cn-- // channels are numbered from 1
	if uint(cn) > 6 {
		panic("dma: bad stream")
	}
	if uint(rn) > 15 {
		panic("dma: bad request")
	}
	paddr := uintptr(unsafe.Pointer(&d.c[cn])) &^ 0x3ff
	return Channel{paddr | uintptr(rn)<<4 | 8 | uintptr(cn)}
}

func (c Channel) periph() *Controller {
	return (*Controller)(unsafe.Pointer(c.h &^ 0xff))
}

func (c Channel) ch() *channel {
	addr := c.h&^0xf7 + c.h&7*unsafe.Sizeof(channel{})
	return (*channel)(unsafe.Pointer(addr))
}

func (c Channel) cnum() uintptr { return c.h & 7 }
func (c Channel) rnum() uint32  { return uint32(c.h >> 4 & 15) }

const (
	//gif = 1<<0  do not use global interrupt flag bacause of silicon bugs

	trce = 1 << 1
	htce = 1 << 2

	trerr = 1 << 3
	fferr = 0
	dmerr = 0
)

func (c Channel) status() byte {
	isr := c.periph().isr.Load()
	return byte(isr >> (c.cnum() * 4) & 0xf)
}

func (c Channel) clear(flags byte) {
	mask := uint32(flags&0xf) << (c.cnum() * 4)
	c.periph().ifcr.Store(mask)
}

func (c Channel) enable() {
	c.ch().cr.SetBits(en)
}

func (c Channel) disable() {
	c.ch().cr.ClearBits(en)
}

func (c Channel) enabled() bool {
	return c.ch().cr.LoadBits(en) != 0
}

func (c Channel) irqEnabled() byte {
	return byte(c.ch().cr.LoadBits(irqs))
}

func (c Channel) enableIRQ(flags byte) {
	c.ch().cr.SetBits(uint32(flags) & irqs)
}

func (c Channel) disableIRQ(flags byte) {
	c.ch().cr.ClearBits(uint32(flags) & irqs)
}

func (c Channel) setup(m Mode) {
	const mask = mtp | mtm | circ | incP | incM | pl
	c.ch().cr.StoreBits(mask, uint32(m))
	n := 4 * c.cnum()
	internal.AtomicStoreBits(&c.periph().cselr, 0xf<<n, c.rnum()<<n)
}

const (
	prioM = 1
	prioH = 2
	prioV = 3
)

func (c Channel) setPrio(prio Prio) {
	c.ch().cr.StoreBits(pl, uint32(prio)<<12)
}

func (c Channel) prio() Prio {
	return Prio(c.ch().cr.LoadBits(pl) >> 12)
}

func (c Channel) wordSize() (p, m uintptr) {
	ccr := uintptr(c.ch().cr.Load())
	p = 1 << (ccr >> 8 & 3)
	m = 1 << (ccr >> 10 & 3)
	return
}

func (c Channel) setWordSize(p, m uintptr) {
	pm := p&6<<7 | m&6<<9
	c.ch().cr.StoreBits(0xf00, uint32(pm))
}

func (c Channel) len() int {
	return int(c.ch().ndtr.Load() & 0xFFFF)
}

func (c Channel) setLen(n int) {
	c.ch().ndtr.Store(uint32(n) & 0xFFFF)
}

func (c Channel) setAddrP(a unsafe.Pointer) {
	c.ch().par.Store(uint32(uintptr(a)))
}

func (c Channel) setAddrM(a unsafe.Pointer) {
	c.ch().mar.Store(uint32(uintptr(a)))
}
