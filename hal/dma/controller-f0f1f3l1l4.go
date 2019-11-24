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

func (d *Controller) channel(n, _ int) Channel {
	n-- // channels are numbered from 1
	if uint(n) > 7 {
		panic("dma: bad stream")
	}
	return Channel{uintptr(unsafe.Pointer(&d.p.C[n]))}
}

func (c Channel) periph() *dma.Periph {
	return (*dma.Periph)(unsafe.Pointer(c.h &^ 0x3ff))
}

func (c Channel) channel() *dma.RC {
	return (*dma.RC)(unsafe.Pointer(c.h))
}

func (c Channel) num() uintptr {
	off := c.h & 0x3ff
	step := unsafe.Sizeof(dma.RC{})
	return (off - 8) / step
}

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
	return byte(isr >> (c.num() * 4) & 0xf)
}

func (c Channel) clear(flags byte) {
	mask := dma.IFCR(flags&0xf) << (c.num() * 4)
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

func (c Channel) request() Request {
	n := c.num() * 4
	return Request(c.periph().CSELR.LoadBits(0xf << n))
}

func (c Channel) setRequest(req Request) {
	n := c.num() * 4
	internal.AtomicStoreBits(&c.periph().CSELR.U32, 0xf<<n, uint32(req)<<n)
}
