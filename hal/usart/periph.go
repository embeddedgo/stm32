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

// Event is a bitmask that describes events in USART peripheral.
type Event uint16

const (
	Idle       = Event(1 << 4) // IDLE line detected
	RxNotEmpty = Event(1 << 5) // read data register not empty
	TxDone     = Event(1 << 6) // transmission complete
	TxEmpty    = Event(1 << 7) // transmit data register empty
	LINBreak   = Event(1 << 8) // LIN break detection flag
	CTST       = Event(1 << 9) // CTS toggle

	EvAll = Idle | RxNotEmpty | TxDone | TxEmpty | LINBreak | CTST
)

// Error is a bitmask that describes errors that can be detected by USART
// hardware when receiving data.
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

// Status return the events and errors which have occurred since the last Clear.
func (p *Periph) Status() (Event, Error) {
	sr := p.sr.Load()
	return Event(sr) & EvAll, Error(sr) & ErrAll
}

// Clear clears specified events and errors. For MCUs that have no USART_ICR
// register (older than L4, F7 series) only RxNotEmpty, TxDone, LINBreak and CTS
// events can be cleared this way. Other events can be cleared only by specific
// sequence of reading status register and read or write data register.
func (p *Periph) Clear(ev Event, err Error) {
	clear(p, ev, err)
}

// EnableIRQ enables generating interrupts by specified events.
func (p *Periph) EnableIRQ(e Event) {
	if cr1e := e & (Idle | RxNotEmpty | TxDone | TxEmpty); cr1e != 0 {
		p.cr1.SetBits(uint32(cr1e))
	}
	if e&LINBreak != 0 {
		p.cr2.SetBits(lbdie)
	}
	if e&CTST != 0 {
		p.cr3.SetBits(ctsie)
	}
}

// DisableIRQ disables generating interrupts by specified events.
func (p *Periph) DisableIRQ(e Event) {
	if cr1e := e & (Idle | RxNotEmpty | TxDone | TxEmpty); cr1e != 0 {
		p.cr1.ClearBits(uint32(cr1e))
	}
	if e&LINBreak != 0 {
		p.cr2.ClearBits(lbdie)
	}
	if e&CTST != 0 {
		p.cr3.ClearBits(ctsie)
	}
}

// EnableErrIRQ enables generating interrupts by ErrNoise, ErrOverrun,
// ErrFraming errors when DMA is used to handle incoming data.
func (p *Periph) EnableErrIRQ() {
	p.cr3.SetBits(eie)
}

// DisableErrIRQ disables generating of IRQ by ErrNoise, ErrOverrun, ErrFraming
// errors when DMA is used to handle incoming data.
func (p *Periph) DisableErrIRQ() {
	p.cr3.ClearBits(eie)
}

// Baudrate returns the current UART speed [sym/s].
func (p *Periph) Baudrate() int {
	usartdiv := uint(p.brr.Load())
	if p.cr1.LoadBits(over8) != 0 {
		usartdiv = usartdiv&^7>>1 | usartdiv&7
	}
	pclk := uint(p.Bus().Clock())
	return int((pclk + usartdiv/2) / usartdiv)
}

// SetBaudrate sets the UART speed [sym/s].
func (p *Periph) SetBaudrate(baud int) {
	br := uint(baud)
	pclk := uint(p.Bus().Clock())
	if br > pclk/8 {
		panic("bad baudrate")
	}
	usartdiv := (pclk + br/2) / br
	if br > pclk/16 {
		//oversampling = 8
		p.cr1.SetBits(over8) // not supported by STM32F1.
		usartdiv = usartdiv&^7<<1 | usartdiv&7
	} else {
		// oversampling = 16
		p.cr1.ClearBits(over8)
	}
	p.brr.Store(uint32(usartdiv))
}

type Conf1 uint32

const (
	UE  Conf1 = ue  // UART enable
	RE  Conf1 = re  // receiver enable
	TE  Conf1 = te  // transmiter enable
	PCE Conf1 = pce // parity control enable (default even)
	OPS Conf1 = ps  // odd parity select
	W9b Conf1 = m0  // 9 bit data word instead of 8 bit
	W7b Conf1 = m1  // 7 bit data word instead of 8 bit
)

func (p *Periph) Conf1() Conf1 {
	return Conf1(p.cr1.Load() &^ (ue | over8))
}

func (p *Periph) SetConf1(cfg Conf1) {
	p.cr1.Store(p.cr1.Load()&(ue|over8) | uint32(cfg)&^(ue|over8))
}

type Conf2 uint32

const (
	S05      Conf2 = stop0b5  // use 0.5 stop bits insted of 1
	S2       Conf2 = stop2b   // use 2 stop bits instead of 1
	S15      Conf2 = stop1b5  // use 1.5 stop bits instead of 1
	Swap     Conf2 = swap     // swap Tx/Rx pins
	RxInv    Conf2 = rxinv    // invert Rx signal
	TxInv    Conf2 = txinv    // invert Tx signal
	DataInv  Conf2 = datainv  // invert data bits for Tx and Rx
	MSBFirst Conf2 = msbfirst // most significant bit first
	ABREn    Conf2 = abren    // automatic baud rate detection
)

func (p *Periph) Conf2() Conf2 {
	return Conf2(p.cr2.Load())
}

func (p *Periph) SetConf2(cfg Conf2) {
	p.cr2.Store(uint32(cfg))
}

type Conf3 uint32

const (
	IREn   Conf3 = iren   // IrDA mode
	IRLP   Conf3 = irlp   // IrDA low-power
	HDSEL  Conf3 = hdsel  // half-duplex operation
	NACK   Conf3 = nack   // Smart Card auto-retry mode in case of NACK
	SCEN   Conf3 = scen   // Smart Card mode
	DMAR   Conf3 = dmar   // RxNotEmpty event generates DMA request
	DMAT   Conf3 = dmat   // TxEmpty event generates DMA request
	RTSE   Conf3 = rtse   // hardware RTS signal
	CTSE   Conf3 = ctse   // hardware CTS signal
	OneBit Conf3 = onebit // one-bit sampling (noise detection disabled)
)

func (p *Periph) Conf3() Conf3 {
	return Conf3(p.cr3.Load() &^ (eie | ctsie))
}

func (p *Periph) SetConf3(cfg Conf3) {
	p.cr3.Store(p.cr3.Load()&(eie|ctsie) | uint32(cfg)&^(eie|ctsie))
}

// Store stores a word in Tx data register.
func (p *Periph) Store(word uint32) {
	tdr(p).Store(word)
}

// Load loads a word from Rx data register.
func (p *Periph) Load() uint32 {
	return rdr(p).Load()
}

func (p *Periph) EnableTx() {
	p.cr1.SetBits(te)
}

func (p *Periph) DisableTx() {
	p.cr1.ClearBits(te)
}

func (p *Periph) EnableRx() {
	p.cr1.SetBits(re)
}

func (p *Periph) DisableRx() {
	p.cr1.ClearBits(re)
}
