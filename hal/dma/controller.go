// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dma

import "unsafe"

type Controller struct {
	registers
}

// DMA returns n-th DMA controller. The first controller number is 1.
//go:nosplit
func DMA(n int) *Controller {
	return controller(n)
}

func (d *Controller) Num() int {
	return int(dnum(d) + 1)
}

// EnableClock enables clock to the DMA controller. The lp parameter determines
// whether the clock remains on in the low power (sleep) mode.
func (d *Controller) EnableClock(lp bool) {
	d.enableClock(lp)
}

// DisableClock disables clock to the DMA controller.
func (d *Controller) DisableClock() {
	d.disableClock()
}

// Reset resets the DMA controller.
//go:nosplit
func (d *Controller) Reset() {
	d.reset()
}

// Channel returns the value that represents the sn stream (F1/Lx DMA channel)
// together with the connected to it cn request channel (ignored for F1/Lx).
// Channels with the same sn points to the same DMA stream so they can't be
// used at the same time.
//go:nosplit
func (d *Controller) Channel(sn, rn int) Channel {
	return d.channel(sn, rn)
}

type Channel struct {
	h uintptr
}

//go:nosplit
func (c Channel) Num() int {
	return int(csnum(c))
}

type Event uint8

const (
	Complete     Event = trce // Transfer Complete Event.
	HalfComplete Event = htce // Half Transfer Complete Event.

	EvAll = Complete | HalfComplete
)

type Error uint8

const (
	ErrTransfer   Error = trerr // Transfer Error.
	ErrDirectMode Error = dmerr // Direct Mode Error.
	ErrFIFO       Error = fferr // FIFO Error.

	ErrAll = ErrTransfer | ErrDirectMode | ErrFIFO
)

func (err Error) Error() string {
	var (
		s string
		d Error
	)
	switch {
	case err&ErrTransfer != 0:
		d = ErrTransfer
		s = "DMA transfer+"
	case err&ErrDirectMode != 0:
		d = ErrDirectMode
		s = "DMA direct mode+"
	case err&ErrFIFO != 0:
		d = ErrFIFO
		s = "DMA FIFO+"
	default:
		return ""
	}
	if err == d {
		s = s[:len(s)-1]
	}
	return s
}

// Status returns current event and error flags.
//go:nosplit
func (c Channel) Status() (Event, Error) {
	flags := c.status()
	return Event(flags) & EvAll, Error(flags) & ErrAll
}

// Clear clears specified flags.
//go:nosplit
func (c Channel) Clear(ev Event, err Error) {
	c.clear(byte(ev) | byte(err))
}

// Enable enables the channel c. All events and errors should be cleared
// before call this method.
//go:nosplit
func (c Channel) Enable() {
	c.enable()
}

// Disable disables channel.
//go:nosplit
func (c Channel) Disable() {
	c.disable()
}

// Returns true if channel is enabled.
//go:nosplit
func (c Channel) Enabled() bool {
	return c.enabled()
}

// IRQEnabled returns events that are enabled to generate interrupt requests.
//go:nosplit
func (c Channel) IRQEnabled() (Event, Error) {
	flags := c.irqEnabled()
	return Event(flags) & EvAll, Error(flags) & ErrAll
}

// EnableIRQ enables generation of IRQs by ev, err. Documentation does not
// mention it, but IRQ can be not generated if an event was asserted before
// enable IRQ for it. So always enable IRQs before channel. Typically, the
// correct sequence is as follows:
//
//	c.Clear(EvAll, ErrAll)
//	c.EnableIRQ(ev, err)
//	c.Enable()
//go:nosplit
func (c Channel) EnableIRQ(ev Event, err Error) {
	c.enableIRQ(byte(ev) | byte(err))
}

// DisableIRQ disables IRQs generation by ev, err.
//go:nosplit
func (c Channel) DisableIRQ(ev Event, err Error) {
	c.disableIRQ(byte(ev) | byte(err))
}

type Mode uint32

const (
	PTM Mode = 0   // Read from peripheral, write to memory.
	MTP Mode = mtp // Read from memory, write to peripheral.
	MTM Mode = mtm // Read from memory (AddrP), write to memory.

	Circ Mode = circ // Enable circular mode.
	IncP Mode = incP // Peripheral increment mode.
	IncM Mode = incM // Memory increment mode.
	PFC  Mode = pfc  // Peripheral flow controller.

	FT1 Mode = ft1 // FIFO mode, threshold 1/4.
	FT2 Mode = ft2 // FIFO mode, threshold 2/4.
	FT3 Mode = ft3 // FIFO mode, threshold 3/4.
	FT4 Mode = ft4 // FIFO mode, threshold 4/4.

	TrBuf Mode = trbuf // Bufferable transfers. H7 only, for UART/USART/LPUART.

	PB4  Mode = pb4  // Peripheral burst transfer, 4 beats.
	PB8  Mode = pb8  // Peripheral burst transfer, 8 beats.
	PB16 Mode = pb16 // Peripheral burst transfer, 16 beats.
	MB4  Mode = mb4  // Memory burst transfer, 4 beats.
	MB8  Mode = mb4  // Memory burst transfer, 4 beats.
	MB16 Mode = mb4  // Memory burst transfer, 4 beats.
)

// Setup configures the DMA channel channel.
//go:nosplit
func (c Channel) Setup(m Mode) {
	c.setup(m)
}

// Prio describes DMA stream priority level.
type Prio uint8

const (
	Low      Prio = 0     // Low priority.
	Medium   Prio = prioM // Medium priority.
	High     Prio = prioH // High priority.
	VeryHigh Prio = prioV // Very high priority.
)

//go:nosplit
func (c Channel) SetPrio(prio Prio) {
	c.setPrio(prio)
}

//go:nosplit
func (c Channel) Prio() Prio {
	return c.prio()
}

// WordSize returns the current word size (in bytes) for peripheral and memory
// side of transfer.
//go:nosplit
func (c Channel) WordSize() (p, m uintptr) {
	return c.wordSize()
}

// SetWordSize sets the word size (in bytes) for peripheral and memory side of
// transfer.
//go:nosplit
func (c Channel) SetWordSize(p, m uintptr) {
	c.setWordSize(p, m)
}

// Len returns current number of words to transfer.
//go:nosplit
func (c Channel) Len() int {
	return c.len()
}

// SetLen sets number of words to transfer (n <= 65535).
//go:nosplit
func (c Channel) SetLen(n int) {
	c.setLen(n)
}

// SetAddrP sets peripheral address (or memory source address in case of MTM).
//go:nosplit
func (c Channel) SetAddrP(a unsafe.Pointer) {
	c.setAddrP(a)
}

// SetAddrM sets memory address.
//go:nosplit
func (c Channel) SetAddrM(a unsafe.Pointer) {
	c.setAddrM(a)
}

// Controller returns the DMA controller that this channel belongs to.
//go:nosplit
func (c Channel) Controller() *Controller {
	return cctrl(c)
}
