// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407 stm32f7x6

package dma

import (
	"embedded/mmio"
	"unsafe"
)

type stream struct {
	cr   mmio.U32
	ndtr mmio.U32
	par  mmio.U32
	m0ar mmio.U32
	m1ar mmio.U32
	fcr  mmio.U32
}

type registers struct {
	isr  [2]mmio.U32
	ifcr [2]mmio.U32
	s    [8]stream
}

// cr
const (
	en     = 1 << 0
	irqs   = 15 << 1
	dir    = 3 << 6
	pl     = 3 << 16
	pburst = 3 << 21
	mburst = 3 << 23
	chsel  = 7 << 25
)

// fcr
const (
	fth   = 3 << 0
	dmdis = 1 << 2
)

func (d *Controller) channel(sn, cn int) Channel {
	if uint(sn) > 7 {
		panic("dma: bad stream")
	}
	if uint(cn) > 7 {
		panic("dma: bad channel")
	}
	return Channel{uintptr(unsafe.Pointer(&d.s[sn])) | uintptr(cn)}
}

func (c Channel) periph() *Controller {
	return (*Controller)(unsafe.Pointer(c.h &^ 0x3ff))
}

func (c Channel) stream() *stream {
	return (*stream)(unsafe.Pointer(c.h &^ 7))
}

func (c Channel) snum() uintptr {
	off := c.h & 0x3ff
	step := unsafe.Sizeof(stream{})
	return (off - 0x10) / step
}

func (c Channel) cnum() int {
	return int(c.h & 7)
}

const (
	fferr = 1 << 0
	dmerr = 1 << 2
	trerr = 1 << 3

	htce = 1 << 4
	trce = 1 << 5

	flmask = fferr | dmerr | trerr | htce | trce
)

func (c Channel) status() uint8 {
	n := c.snum()
	isr := &c.periph().isr[n/4]
	n = n % 4 * 6
	if n > 6 {
		n += 4
	}
	return uint8(isr.Load() >> n & flmask)
}

func (c Channel) clear(flags byte) {
	n := c.snum()
	ifcr := &c.periph().ifcr[n/4]
	n = n % 4 * 6
	if n > 6 {
		n += 4
	}
	ifcr.Store(uint32(flags&flmask) << n)
}

func (c Channel) enable() {
	c.stream().cr.SetBits(en)
}

func (c Channel) disable() {
	c.stream().cr.ClearBits(en)
}

func (c Channel) enabled() bool {
	return c.stream().cr.LoadBits(en) != 0
}

func (c Channel) irqEnabled() byte {
	s := c.stream()
	ev := byte(s.cr.Load()&irqs<<1) | byte(s.fcr.Load()>>7&1)
	return ev
}

func (c Channel) enableIRQ(flags byte) {
	s := c.stream()
	s.cr.SetBits(uint32(flags) >> 1 & irqs)
	//s.fcr.SetBits(uint32(flags) & 1 << 7) do not use
}

func (c Channel) disableIRQ(flags byte) {
	s := c.stream()
	s.cr.ClearBits(uint32(flags) >> 1 & irqs)
	s.fcr.ClearBits(uint32(flags) & 1 << 7)
}

const (
	ft1 = 0<<0 | dmdis
	ft2 = 1<<0 | dmdis
	ft3 = 2<<0 | dmdis
	ft4 = 3<<0 | dmdis

	pfc = 1 << 5 // pfctrl

	mtp = 1 << 6
	mtm = 2 << 6

	circ = 1 << 8
	incP = 1 << 9  // pinc
	incM = 1 << 10 // minc

	pb4  = 1 << 21
	pb8  = 2 << 21
	pb16 = 3 << 21

	mb4  = 1 << 23
	mb8  = 2 << 23
	mb16 = 3 << 23
)

func (c Channel) setup(m Mode) {
	const mask = pfc | dir | circ | incP | incM | pburst | mburst | chsel
	cr := uint32(c.cnum())<<25 | uint32(m)
	s := c.stream()
	s.cr.StoreBits(mask, cr)
	s.fcr.StoreBits(fth|dmdis, uint32(m))
}

const (
	prioM = 1
	prioH = 2
	prioV = 3
)

func (c Channel) setPrio(prio Prio) {
	c.stream().cr.StoreBits(pl, uint32(prio)<<16)
}

func (c Channel) prio() Prio {
	return Prio(c.stream().cr.LoadBits(pl) >> 16)
}

func (c Channel) wordSize() (p, m uintptr) {
	cr := uintptr(c.stream().cr.Load())
	p = 1 << (cr >> 11 & 3)
	m = 1 << (cr >> 13 & 3)
	return
}

func (c Channel) setWordSize(p, m uintptr) {
	cr := p&6<<10 | m&6<<12
	c.stream().cr.StoreBits(0x7800, uint32(cr))
}

func (c Channel) len() int {
	return int(c.stream().ndtr.Load() & 0xFFFF)
}

func (c Channel) setLen(n int) {
	c.stream().ndtr.Store(uint32(n) & 0xFFFF)
}

func (c Channel) setAddrP(a unsafe.Pointer) {
	c.stream().par.Store(uint32(uintptr(a)))
}

func (c Channel) setAddrM(a unsafe.Pointer) {
	c.stream().m0ar.Store(uint32(uintptr(a)))
}
