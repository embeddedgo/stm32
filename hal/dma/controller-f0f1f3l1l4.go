// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32l4x6

package dma

import (
	"unsafe"

	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/dma"
)

func (d *Controller) channel(cn, rn int) Channel {
	cn-- // channels are numbered from 1
	if uint(cn) > 6 {
		panic("dma: bad stream")
	}
	if uint(rn) > 15 {
		panic("dma: bad request")
	}
	paddr := uintptr(unsafe.Pointer(&d.p.C[cn])) &^ 0x3ff
	return Channel{paddr | uintptr(rn)<<4 | 8 | uintptr(cn)}
}

func (c Channel) periph() *dma.Periph {
	return (*dma.Periph)(unsafe.Pointer(c.h &^ 0xff))
}

func (c Channel) channel() *dma.RC {
	addr := c.h&^0xf7 + c.h&7*unsafe.Sizeof(dma.RC{})
	return (*dma.RC)(unsafe.Pointer(addr))
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
	isr := c.periph().ISR.Load()
	return byte(isr >> (c.cnum() * 4) & 0xf)
}

func (c Channel) clear(flags byte) {
	mask := dma.IFCR(flags&0xf) << (c.cnum() * 4)
	c.periph().IFCR.Store(mask)
}

func (c Channel) enable() {
	c.channel().CR.SetBits(dma.EN)
}

func (c Channel) disable() {
	c.channel().CR.ClearBits(dma.EN)
}

func (c Channel) enabled() bool {
	return c.channel().CR.LoadBits(dma.EN) != 0
}

func (c Channel) irqEnabled() byte {
	return byte(c.channel().CR.Load() & 0xe)
}

func (c Channel) enableIRQ(flags byte) {
	c.channel().CR.SetBits(dma.CR(flags) & 0xe)
}

func (c Channel) disableIRQ(flags byte) {
	c.channel().CR.ClearBits(dma.CR(flags) & 0xe)
}

const (
	mtp = 1 << dma.DIRn
	mtm = 1 << dma.MEM2MEMn

	circ = 1 << dma.CIRCn
	incP = 1 << dma.PINCn
	incM = 1 << dma.MINCn

	ft1 = 0
	ft2 = 0
	ft3 = 0
	ft4 = 0

	pb4  = 0
	pb8  = 0
	pb16 = 0
	mb4  = 0
	mb8  = 0
	mb16 = 0

	pfc = 0
)

func (c Channel) setup(m Mode) {
	mask := dma.DIR | dma.MEM2MEM | dma.CIRC | dma.PINC | dma.MINC | dma.PL
	c.channel().CR.StoreBits(mask, dma.CR(m))
	n := 4 * c.cnum()
	internal.AtomicStoreBits(&c.periph().CSELR.U32, 0xf<<n, c.rnum()<<n)
}

const (
	prioM = 1
	prioH = 2
	prioV = 3
)

func (c Channel) setPrio(prio Prio) {
	c.channel().CR.StoreBits(dma.PL, dma.CR(prio)<<dma.PLn)
}

func (c Channel) prio() Prio {
	return Prio(c.channel().CR.LoadBits(dma.PL) >> dma.PLn)
}

func (c Channel) wordSize() (p, m uintptr) {
	ccr := uintptr(c.channel().CR.Load())
	p = 1 << (ccr >> 8 & 3)
	m = 1 << (ccr >> 10 & 3)
	return
}

func (c Channel) setWordSize(p, m uintptr) {
	pm := dma.CR(p&6<<7 | m&6<<9)
	c.channel().CR.StoreBits(0xf00, dma.CR(pm))
}

func (c Channel) len() int {
	return int(c.channel().NDTR.Load() & 0xFFFF)
}

func (c Channel) setLen(n int) {
	c.channel().NDTR.Store(dma.NDTR(n) & 0xFFFF)
}

func (c Channel) setAddrP(a unsafe.Pointer) {
	c.channel().PAR.Store(dma.PAR(uintptr(a)))
}

func (c Channel) setAddrM(a unsafe.Pointer) {
	c.channel().MAR.Store(dma.MAR(uintptr(a)))
}
