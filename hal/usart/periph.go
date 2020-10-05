// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usart

import (
	"unsafe"

	"github.com/embeddedgo/stm32/hal/internal/apb"
	"github.com/embeddedgo/stm32/p/bus"
)

// Periph represents USART peripheral.
type Periph struct {
	registers
}

// Bus returns a bus to which p is connected to.
func (p *Periph) Bus() bus.Bus {
	return busForAddr(p)
}

// EnableClock enables clock for p.
// lp determines whether the clock remains on in low power (sleep) mode.
func (p *Periph) EnableClock(lp bool) {
	apb.EnableClock(unsafe.Pointer(p), lp)
}

// DisableClock disables clock for p.
func (p *Periph) DisableClock() {
	apb.DisableClock(unsafe.Pointer(p))
}

// Reset resets p.
func (p *Periph) Reset() {
	apb.Reset(unsafe.Pointer(p))
}

// Event is bitmask that describes events in USART peripheral.
type Event uint16

const (
	Idle       = Event(1 << 4) // IDLE line detected.
	RxNotEmpty = Event(1 << 5) // Read data register not empty.
	TxDone     = Event(1 << 6) // Transmission complete.
	TxEmpty    = Event(1 << 7) // Transmit data register empty.
	LINBreak   = Event(1 << 8) // LIN break detection flag.
	CTS        = Event(1 << 9) // Change on CTS status line

	EvAll = Idle | RxNotEmpty | TxDone | TxEmpty | LINBreak | CTS
)

// Error is bitmask that describes errors that can be detected by USART hardware
// when receiving data.
type Error uint8

const (
	ErrParity  = Error(1 << 0) // Parity error.
	ErrFraming = Error(1 << 1) // Framing error.
	ErrNoise   = Error(1 << 2) // Noise error flag.
	ErrOverrun = Error(1 << 3) // Overrun error.

	ErrAll = ErrParity | ErrFraming | ErrNoise | ErrOverrun
)

func (e Error) Error() string {
	var (
		s string
		d Error
	)
	switch {
	case e&ErrOverrun != 0:
		d = ErrOverrun
		s = "USART overrun+"
	case e&ErrNoise != 0:
		d = ErrNoise
		s = "USART noise+"
	case e&ErrFraming != 0:
		d = ErrFraming
		s = "USART framing+"
	case e&ErrParity != 0:
		d = ErrParity
		s = "USART parity+"
	default:
		return ""
	}
	if e&^d == 0 {
		s = s[:len(s)-1]
	}
	return s
}

// Status return current status of p.
func (p *Periph) Status() (Event, Error) {
	sr := p.sr.Load()
	return Event(sr) & EvAll, Error(sr) & ErrAll
}

// Clear clears events ev and errors err. For MCUs that have no USART_ICR
// register (older than L4, F7 series) only RxNotEmpty, TxDone, LINBreak and CTS
// events can be cleared this way. Other events can be cleared only by specific
// sequence of reading status register and read or write data register.
func (p *Periph) Clear(ev Event, err Error) {
	clear(p, ev, err)
}

/*
// EnableIRQ enables generating of IRQ by events e.
func (p *Periph) EnableIRQ(e Event) {
	if cr1e := e & (Idle | RxNotEmpty | TxDone | TxEmpty); cr1e != 0 {
		p.raw.CR1.SetBits(usart.CR1(cr1e) << 4)
	}
	if e&LINBreak != 0 {
		p.raw.CR2.SetBits(lbdie)
	}
	if e&CTS != 0 {
		p.raw.CTSIE().Set()
	}
}

// DisableIRQ disables generating of IRQ by events e.
func (p *Periph) DisableIRQ(e Event) {
	if cr1e := e & (Idle | RxNotEmpty | TxDone | TxEmpty); cr1e != 0 {
		p.raw.CR1.ClearBits(usart.CR1(cr1e) << 4)
	}
	if e&LINBreak != 0 {
		p.raw.CR2.ClearBits(lbdie)
	}
	if e&CTS != 0 {
		p.raw.CTSIE().Clear()
	}
}

// EnableErrorIRQ enables generating of IRQ by ErrNoise, ErrOverrun, ErrFraming
// errors when DMA is used to handle incoming data.
func (p *Periph) EnableErrorIRQ() {
	p.raw.EIE().Set()
}

func (p *Periph) DisableErrorIRQ() {
	p.raw.EIE().Clear()
}

// SetBaudRate sets baud rate [sym/s]. APB1 and APB2 clock in stm32/hal/system
// package must be set properly before use this function.
func (p *Periph) SetBaudRate(baudrate int) {
	br := uint(baudrate)
	pclk := p.Bus().Clock()
	usartdiv := (pclk + br/2) / br
	if uint(br) > pclk/16 {
		// Oversampling = 8
		p.raw.OVER8().Set() // BUG: Not supported by F1xx.
		usartdiv = usartdiv&^7<<1 | usartdiv&7
	} else {
		// Oversampling = 16
		p.raw.OVER8().Clear()
	}
	p.raw.BRR.Store(usart.BRR(usartdiv))
}

// Enable enables p.
func (p *Periph) Enable() {
	p.raw.UE().Set()
}

// Disable disables p at end of the current byte transfer.
func (p *Periph) Disable() {
	p.raw.UE().Clear()
}

type Conf1 uint32

const (
	RxEna   Conf1 = 1 << 2  // Receiver enabled.
	TxEna   Conf1 = 1 << 3  // Transmiter enabled.
	ParEven Conf1 = 2 << 9  // Parity control enabled: even.
	ParOdd  Conf1 = 3 << 9  // Parity control enabled: odd.
	Word9b  Conf1 = 1 << 12 // 9 bit data word instead of 8 bit.
	over8         = 1 << 15
	Word7b  Conf1 = 1 << 28 // 7 bit data word instead of 8 bit.
)

func (p *Periph) Conf1() Conf1 {
	return Conf1(p.raw.CR1.Load() &^ over8)
}

func (p *Periph) SetConf1(c Conf1) {
	cr1 := p.raw.CR1.Load()&over8 | usart.CR1(c)&^over8
	p.raw.CR1.Store(cr1)
}

type Conf2 uint32

const (
	Stop0b5 Conf2 = 1 << 12 // Use 0.5 stop bits insted of 1.
	Stop2b  Conf2 = 2 << 12 // Use 2 stop bits instead of 1.
	Stop1b5 Conf2 = 3 << 12 // Use 1.5 stop bits instead of 1.
	Swap    Conf2 = 1 << 15 // Swap Tx/Rx pins.
	RxInv   Conf2 = 1 << 16 // Invert Rx signal.
	TxInv   Conf2 = 1 << 17 // Invert Tx signal.
	DataInv Conf2 = 1 << 18 // Ivert data bits for Tx and Rx.
)

func (p *Periph) Conf2() Conf2 {
	return Conf2(p.raw.CR2.Load())
}

func (p *Periph) SetConf2(c Conf2) {
	p.raw.CR2.Store(usart.CR2(c))
}

type Conf3 uint32

const (
	// HalfDuplex enables half-duplx operation.
	HalfDuplex Conf3 = 1 << 3

	// OneBit sets sampling method to one bit and disables Noise error
	// detection.
	OneBit Conf3 = 1 << 11
)

func (p *Periph) Conf3() Conf3 {
	return Conf3(p.raw.CR3.Load())
}

func (p *Periph) SetConf3(c Conf3) {
	p.raw.CR3.Store(usart.CR3(c))
}

func (p *Periph) Store(d int) {
	p.store(d)
}

func (p *Periph) Load() int {
	return p.load()
}

*/
