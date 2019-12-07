// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407 stm32f7x6

package dma

import (
	"unsafe"

	"github.com/embeddedgo/stm32/p/dma"
)

func (d *Controller) channel(sn, cn int) Channel {
	if uint(sn) > 7 {
		panic("dma: bad stream")
	}
	if uint(cn) > 7 {
		panic("dma: bad channel")
	}
	return Channel{uintptr(unsafe.Pointer(&d.p.S[sn])) | uintptr(cn)}
}

func (c Channel) periph() *dma.Periph {
	return (*dma.Periph)(unsafe.Pointer(c.h &^ 0x3ff))
}

func (c Channel) stream() *dma.RS {
	return (*dma.RS)(unsafe.Pointer(c.h &^ 7))
}

func (c Channel) snum() uintptr {
	off := c.h & 0x3ff
	step := unsafe.Sizeof(dma.RS{})
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
	isr := &c.periph().ISR[n/4]
	n = n % 4 * 6
	if n > 6 {
		n += 4
	}
	return uint8(isr.Load() >> n & flmask)
}

func (c Channel) clear(flags byte) {
	n := c.snum()
	ifcr := &c.periph().IFCR[n/4]
	n = n % 4 * 6
	if n > 6 {
		n += 4
	}
	ifcr.Store(dma.IFCR(flags&flmask) << n)
}

func (c Channel) enable() {
	c.stream().CR.SetBits(dma.EN)
}

func (c Channel) disable() {
	c.stream().CR.ClearBits(dma.EN)
}

func (c Channel) enabled() bool {
	return c.stream().CR.LoadBits(dma.EN) != 0
}

func (c Channel) irqEnabled() byte {
	s := c.stream()
	ev := byte(s.CR.Load()&0x1e<<1) | byte(s.FCR.Load()>>7&1)
	return ev
}

func (c Channel) enableIRQ(flags byte) {
	s := c.stream()
	s.CR.SetBits(dma.CR(flags) >> 1 & 0x1e)
	//s.FCR.SetBits(dma.FCR(flags) & 1 << 7) do not use
}

func (c Channel) disableIRQ(flags byte) {
	s := c.stream()
	s.CR.ClearBits(dma.CR(flags) >> 1 & 0x1e)
	s.FCR.ClearBits(dma.FCR(flags) & 1 << 7)
}

const (
	ft1 = 4 << 0
	ft2 = 5 << 0
	ft3 = 6 << 0
	ft4 = 7 << 0

	pfc = 1 << 5

	mtp = 1 << 6
	mtm = 2 << 6

	circ = 1 << 8
	incP = 1 << 9
	incM = 1 << 10

	pb4  = 1 << 21
	pb8  = 2 << 21
	pb16 = 3 << 21
	mb4  = 1 << 23
	mb8  = 2 << 23
	mb16 = 3 << 23

	chSel = 7 << 25 // only to see that others don't interfere with CHSEL.
)

func (c Channel) setup(m Mode) {
	cr := dma.CR(c.cnum())<<dma.CHSELn | dma.CR(m)
	mask := dma.PFCTRL | dma.DIR | dma.CIRC | dma.PINC | dma.MINC |
		dma.PBURST | dma.MBURST | dma.CHSEL
	s := c.stream()
	s.CR.StoreBits(mask, cr)
	s.FCR.StoreBits(dma.DMDIS|dma.FTH, dma.FCR(m))
}

const (
	prioM = 1
	prioH = 2
	prioV = 3
)

func (c Channel) setPrio(prio Prio) {
	c.stream().CR.StoreBits(dma.PL, dma.CR(prio)<<dma.PLn)
}

func (c Channel) prio() Prio {
	return Prio(c.stream().CR.LoadBits(dma.PL) >> dma.PLn)
}

func (c Channel) wordSize() (p, m uintptr) {
	cr := uintptr(c.stream().CR.Load())
	p = 1 << (cr >> 11 & 3)
	m = 1 << (cr >> 13 & 3)
	return
}

func (c Channel) setWordSize(p, m uintptr) {
	cr := p&6<<10 | m&6<<12
	c.stream().CR.StoreBits(0x7800, dma.CR(cr))
}

func (c Channel) len() int {
	return int(c.stream().NDTR.Load() & 0xFFFF)
}

func (c Channel) setLen(n int) {
	c.stream().NDTR.Store(dma.NDTR(n) & 0xFFFF)
}

func (c Channel) setAddrP(a unsafe.Pointer) {
	c.stream().PAR.Store(dma.PAR(uintptr(a)))
}

func (c Channel) setAddrM(a unsafe.Pointer) {
	c.stream().M0AR.Store(dma.M0AR(uintptr(a)))
}
