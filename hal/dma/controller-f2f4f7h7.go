// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32f215 || stm32f407 || stm32f7x6 || stm32h7x3

package dma

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/internal"
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

//go:nosplit
func (d *Controller) channel(sn, cn int) Channel {
	if uint(sn) > 7 {
		panic("bad stream")
	}
	if uint(cn) > 7 {
		panic("bad channel")
	}
	return Channel{uintptr(unsafe.Pointer(&d.s[sn])) | uintptr(cn)}
}

//go:nosplit
func cctrl(c Channel) *Controller {
	return (*Controller)(unsafe.Pointer(c.h &^ 0x3ff))
}

//go:nosplit
func st(c Channel) *stream {
	return (*stream)(unsafe.Pointer(c.h &^ 7))
}

//go:nosplit
func snum(c Channel) uintptr {
	off := c.h & 0x3ff
	step := unsafe.Sizeof(stream{})
	return (off - 0x10) / step
}

//go:nosplit
func csnum(c Channel) uintptr { return snum(c) }

//go:nosplit
func cnum(c Channel) int {
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

//go:nosplit
func (c Channel) status() uint8 {
	n := snum(c)
	isr := &cctrl(c).isr[n/4]
	n = n % 4 * 6
	if n > 6 {
		n += 4
	}
	return uint8(isr.Load() >> n & flmask)
}

//go:nosplit
func (c Channel) clear(flags byte) {
	n := snum(c)
	ifcr := &cctrl(c).ifcr[n/4]
	n = n % 4 * 6
	if n > 6 {
		n += 4
	}
	ifcr.Store(uint32(flags&flmask) << n)
}

//go:nosplit
func (c Channel) enable() {
	st(c).cr.SetBits(en)
}

//go:nosplit
func (c Channel) disable() {
	st(c).cr.ClearBits(en)
}

func (c Channel) enabled() bool {
	return st(c).cr.LoadBits(en) != 0
}

//go:nosplit
func (c Channel) irqEnabled() byte {
	s := st(c)
	ev := byte(s.cr.Load() & irqs << 1) //| byte(s.fcr.Load()>>7&1)
	return ev
}

//go:nosplit
func (c Channel) enableIRQ(flags byte) {
	s := st(c)
	s.cr.SetBits(uint32(flags) >> 1 & irqs)
	//s.fcr.SetBits(uint32(flags) & 1 << 7) do not use
}

//go:nosplit
func (c Channel) disableIRQ(flags byte) {
	s := st(c)
	s.cr.ClearBits(uint32(flags) >> 1 & irqs)
	//s.fcr.ClearBits(uint32(flags) & 1 << 7)
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

	trbuf = 1 << 20 // UART/USART/LPUART transfers. H7 only

	pb4  = 1 << 21
	pb8  = 2 << 21
	pb16 = 3 << 21

	mb4  = 1 << 23
	mb8  = 2 << 23
	mb16 = 3 << 23
)

//go:nosplit
func (c Channel) setup(m Mode) {
	mask := uint32(pfc | dir | circ | incP | incM | pburst | mburst | chsel)
	if internal.H7 {
		mask |= trbuf
	}
	cr := uint32(cnum(c))<<25 | uint32(m)
	s := st(c)
	s.cr.StoreBits(mask, cr)
	s.fcr.StoreBits(fth|dmdis, uint32(m))
}

const (
	prioM = 1
	prioH = 2
	prioV = 3
)

//go:nosplit
func (c Channel) setPrio(prio Prio) {
	st(c).cr.StoreBits(pl, uint32(prio)<<16)
}

//go:nosplit
func (c Channel) prio() Prio {
	return Prio(st(c).cr.LoadBits(pl) >> 16)
}

//go:nosplit
func (c Channel) wordSize() (p, m uintptr) {
	cr := uintptr(st(c).cr.Load())
	p = 1 << (cr >> 11 & 3)
	m = 1 << (cr >> 13 & 3)
	return
}

//go:nosplit
func (c Channel) setWordSize(p, m uintptr) {
	cr := p&6<<10 | m&6<<12
	st(c).cr.StoreBits(0x7800, uint32(cr))
}

//go:nosplit
func (c Channel) len() int {
	return int(st(c).ndtr.Load() & 0xFFFF)
}

//go:nosplit
func (c Channel) setLen(n int) {
	st(c).ndtr.Store(uint32(n) & 0xFFFF)
}

//go:nosplit
func (c Channel) setAddrP(a unsafe.Pointer) {
	st(c).par.Store(uint32(uintptr(a)))
}

//go:nosplit
func (c Channel) setAddrM(a unsafe.Pointer) {
	st(c).m0ar.Store(uint32(uintptr(a)))
}
