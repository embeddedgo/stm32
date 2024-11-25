// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32l4x6

package dma

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/internal"
)

type channel struct {
	cr   mmio.R32[uint32]
	ndtr mmio.R32[uint32]
	par  mmio.R32[uint32]
	mar  mmio.R32[uint32]
	_    uint32
}

type registers struct {
	isr   mmio.R32[uint32]
	ifcr  mmio.R32[uint32]
	c     [7]channel
	_     [5]uint32
	cselr mmio.R32[uint32]
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
	ft1   = 0
	ft2   = 0
	ft3   = 0
	ft4   = 0
	trbuf = 0
	pb4   = 0
	pb8   = 0
	pb16  = 0
	mb4   = 0
	mb8   = 0
	mb16  = 0
	pfc   = 0
)

func (d *Controller) channel(cn, rn int) Channel {
	cn-- // channels are numbered from 1
	if uint(cn) > 6 {
		panic("bad stream")
	}
	if uint(rn) > 15 {
		panic("bad request")
	}
	paddr := uintptr(unsafe.Pointer(&d.c[cn])) &^ 0x3ff
	return Channel{paddr | uintptr(rn)<<4 | 8 | uintptr(cn)}
}

func regs(c Channel) *registers {
	return (*registers)(unsafe.Pointer(c.h &^ 0xff))
}

func ch(c Channel) *channel {
	addr := c.h&^0xf7 + c.h&7*unsafe.Sizeof(channel{})
	return (*channel)(unsafe.Pointer(addr))
}

func cnum(c Channel) uintptr { return c.h & 7 }
func rnum(c Channel) uint32  { return uint32(c.h >> 4 & 15) }

const (
	//gif = 1<<0  do not use global interrupt flag bacause of silicon bugs

	trce = 1 << 1
	htce = 1 << 2

	trerr = 1 << 3
	fferr = 0
	dmerr = 0
)

func (c Channel) status() byte {
	isr := regs(c).isr.Load()
	return byte(isr >> (cnum(c) * 4) & 0xf)
}

func (c Channel) clear(flags byte) {
	mask := uint32(flags&0xf) << (cnum(c) * 4)
	regs(c).ifcr.Store(mask)
}

func (c Channel) enable() {
	ch(c).cr.SetBits(en)
}

func (c Channel) disable() {
	ch(c).cr.ClearBits(en)
}

func (c Channel) enabled() bool {
	return ch(c).cr.LoadBits(en) != 0
}

func (c Channel) irqEnabled() byte {
	return byte(ch(c).cr.LoadBits(irqs))
}

func (c Channel) enableIRQ(flags byte) {
	ch(c).cr.SetBits(uint32(flags) & irqs)
}

func (c Channel) disableIRQ(flags byte) {
	ch(c).cr.ClearBits(uint32(flags) & irqs)
}

func (c Channel) setup(m Mode) {
	const mask = mtp | mtm | circ | incP | incM | pl
	ch(c).cr.StoreBits(mask, uint32(m))
	n := 4 * cnum(c)
	internal.AtomicStoreBits(&regs(c).cselr, 0xf<<n, rnum(c)<<n)
}

const (
	prioM = 1
	prioH = 2
	prioV = 3
)

func (c Channel) setPrio(prio Prio) {
	ch(c).cr.StoreBits(pl, uint32(prio)<<12)
}

func (c Channel) prio() Prio {
	return Prio(ch(c).cr.LoadBits(pl) >> 12)
}

func (c Channel) wordSize() (p, m uintptr) {
	ccr := uintptr(ch(c).cr.Load())
	p = 1 << (ccr >> 8 & 3)
	m = 1 << (ccr >> 10 & 3)
	return
}

func (c Channel) setWordSize(p, m uintptr) {
	pm := p&6<<7 | m&6<<9
	ch(c).cr.StoreBits(0xf00, uint32(pm))
}

func (c Channel) len() int {
	return int(ch(c).ndtr.Load() & 0xFFFF)
}

func (c Channel) setLen(n int) {
	ch(c).ndtr.Store(uint32(n) & 0xFFFF)
}

func (c Channel) setAddrP(a unsafe.Pointer) {
	ch(c).par.Store(uint32(uintptr(a)))
}

func (c Channel) setAddrM(a unsafe.Pointer) {
	ch(c).mar.Store(uint32(uintptr(a)))
}
