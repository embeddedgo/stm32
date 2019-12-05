// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spi

import (
	"embedded/mmio"
	"math/bits"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/internal/apb"
	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/spi"
)

// Periph represents SPI peripheral.
type Periph struct {
	raw spi.Periph
}

// Bus returns a bus to which p is connected to.
func (p *Periph) Bus() bus.Bus {
	return p.raw.Bus()
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

type Config uint16

const (
	CPHA0  = Config(0)        // Sample on first edge.
	CPHA1  = Config(spi.CPHA) // Sample on second edge.
	CPOL0  = Config(0)        // Clock idle state is 0.
	CPOL1  = Config(spi.CPOL) // Clock idle state is 1.
	Slave  = Config(0)        // Slave mode.
	Master = Config(spi.MSTR) // Master mode.

	BR2   = Config(0)            // Baud rate = PCLK/2
	BR4   = Config(1 << spi.BRn) // Baud rate = PCLK/4.
	BR8   = Config(2 << spi.BRn) // Baud rate = PCLK/8.
	BR16  = Config(3 << spi.BRn) // Baud rate = PCLK/16.
	BR32  = Config(4 << spi.BRn) // Baud rate = PCLK/32.
	BR64  = Config(5 << spi.BRn) // Baud rate = PCLK/64.
	BR128 = Config(6 << spi.BRn) // Baud rate = PCLK/128.
	BR256 = Config(7 << spi.BRn) // Baud rate = PCLK/256.

	MSBF = Config(0)            // Most significant bit first.
	LSBF = Config(spi.LSBFIRST) // Least significant bit first.

	HardSS = Config(0)       // Hardware slave select.
	SoftSS = Config(spi.SSM) // Software slave select (use ISSLow, ISSHigh).

	ISSLow  = Config(0)       // Set NSS internally to low (requires SoftSS).
	ISSHigh = Config(spi.SSI) // Set NSS internally to high (requires SoftSS).

	ThreeWire = Config(0)            // Three-wire mode (SCK, MOSI, MISO).
	TwoWire   = Config(spi.BIDIMODE) // Two-wire mode (SCK, MOSI/MISO).

	Rx = Config(0)          // Two-wire receive-only
	Tx = Config(spi.BIDIOE) // Two-wire transmit-only.
)

// BR calculates the baud rate bits of configuration. BR guarantees that
// the returned bits corresponds to the value closest to but not greater than
// baudrate.
func (p *Periph) BR(baudrate int) Config {
	pclk := p.Bus().Clock()
	div := uint32((pclk + int64(baudrate) - 1) / int64(baudrate))
	switch {
	case div < 2:
		div = 2
	case div > 256:
		div = 256
	}
	br := 31 - bits.LeadingZeros32(uint32(div-1))
	return Config(br << spi.BRn)
}

// Config returns the current p's configuration.
func (p *Periph) Config() Config {
	return Config(p.raw.CR1.LoadBits(cr1Mask))
}

// SetConfig configures p (p must be disabled). If baudrate > 0 it replaces the
// BR bits in conf with the ones calculated from baudrate.
func (p *Periph) SetConfig(conf Config, baudrate int) {
	if baudrate != 0 {
		conf = conf&^BR256 | p.BR(baudrate)
	}
	p.raw.CR1.StoreBits(cr1Mask, spi.CR1(conf))
}

// Baudrate returns the currently configured baudrate.
func (p *Periph) Baudrate() int {
	br := p.raw.CR1.LoadBits(cr1Mask) >> spi.BRn & 7
	return int(p.Bus().Clock() >> (br + 1))
}

// SetBaudrate sets the baudrate (p must be disabled).
func (p *Periph) SetBaudrate(baudrate int) {
	p.raw.CR1.StoreBits(spi.CR1(BR256), spi.CR1(p.BR(baudrate)))
}

// WordSize return currently used word size.
func (p *Periph) WordSize() int {
	return p.wordSize()
}

// SetWordSize sets size of data word. All families support 8 and 16-bit words,
// F0, F3, L4 supports 4 to 16-bit.  Some families require disabled peripheral
// before use this function. SetWordSize is mandatory for F0, F3, L4 because
// the default reset configuration does not work.
func (p *Periph) SetWordSize(size int) {
	p.setWordSize(size)
}

// TwoWireSetRx sets data direction to recevie in two-wire mode.
func (p *Periph) TwoWireSetRx() {
	p.raw.CR1.ClearBits(spi.BIDIOE)
}

// TwoWireSetTx sets data direction to transmit in two-wire mode.
func (p *Periph) TwoWireSetTx() {
	p.raw.CR1.SetBits(spi.BIDIOE)
}

// Event is a bitfield that encodes possible peripheral events.
type Event uint16

const (
	Err        = Event(1)             // Some hardware error occurs.
	RxNotEmpty = Event(spi.RXNE) << 1 // Receive buffer not empty.
	TxEmpty    = Event(spi.TXE) << 1  // Transmit buffer empty.
	Busy       = Event(spi.BSY) << 1  // Periph is busy (not a real event).

	realEventMask = RxNotEmpty | TxEmpty
	errEventMask  = realEventMask | Err
	bsyEventMask  = realEventMask | Busy
)

// Error is a bitfield that encodes possible peripheral errors.
type Error byte

const (
	ErrCRC     = Error(spi.CRCERR >> 3)
	ErrMode    = Error(spi.MODF >> 3)
	ErrOverrun = Error(spi.OVR >> 3)

	errorMask = ErrCRC | ErrMode | ErrOverrun
)

func (e Error) Error() string {
	var (
		s string
		d Error
	)
	switch {
	case e&ErrCRC != 0:
		d = ErrCRC
		s = "SPI CRC error+"
	case e&ErrMode != 0:
		d = ErrMode
		s = "SPI mode fault+"
	case e&ErrOverrun != 0:
		d = ErrOverrun
		s = "SPI overrun+"
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
	sr := p.raw.SR.Load()
	err := Error(sr>>3) & errorMask
	ev := Event(sr<<1) & bsyEventMask
	if err != 0 {
		ev |= Err
	}
	return ev, err
}

// EnableIRQ enables generating of IRQ by events e.
func (p *Periph) EnableIRQ(e Event) {
	if e &= errEventMask; e != 0 {
		p.raw.CR2.SetBits(spi.CR2(e) << spi.ERRIEn)
	}
}

// DisableIRQ disables generating of IRQ by events e.
func (p *Periph) DisableIRQ(e Event) {
	if e &= errEventMask; e != 0 {
		p.raw.CR2.ClearBits(spi.CR2(e) << spi.ERRIEn)
	}
}

// EnableDMA enables generating of DMA requests by events e.
func (p *Periph) EnableDMA(e Event) {
	if e &= realEventMask; e != 0 {
		p.raw.CR2.SetBits(spi.CR2(e) >> 1)
	}
}

// DisableDMA disables generating of DMA requests by events e.
func (p *Periph) DisableDMA(e Event) {
	if e &= realEventMask; e != 0 {
		p.raw.CR2.ClearBits(spi.CR2(e) >> 1)
	}
}

// Enabled reports whether p is enabled.
func (p *Periph) Enabled() bool {
	return p.raw.SPE().Load() != 0
}

// Enable enables p.
func (p *Periph) Enable() {
	p.raw.SPE().Set()
}

// Disable disables p.
func (p *Periph) Disable() {
	p.raw.SPE().Clear()
}

// Duplex describes duplex mode. In full-duplex mode transmission is performed
// using MOSI and MISO lines. In half-duplex mode only MOSI at master side and
// MISO at slave side are used.
type Duplex uint16

const (
	Full    = Duplex(0)                         // Full-duplex mode.
	HalfIn  = Duplex(spi.BIDIMODE)              // Half-duplex input mode.
	HalfOut = Duplex(spi.BIDIMODE | spi.BIDIOE) // Half-duplex output mode.
)

func (p *Periph) Duplex() Duplex {
	return Duplex(p.raw.CR1.Load()) & HalfOut
}

func (p *Periph) SetDuplex(duplex Duplex) {
	cr1 := p.raw.CR1.Load()
	if cr1&spi.CR1(HalfOut) != spi.CR1(duplex) {
		p.raw.CR1.Store(cr1&^spi.CR1(HalfOut) | spi.CR1(duplex))
	}
}

// StoreWord16 stores a 16-bit word to the data register. Use it only when 16-bit
// word or data packing is configured.
func (p *Periph) StoreWord16(v uint16) {
	(*mmio.U16)(unsafe.Pointer(&p.raw.DR)).Store(v)
}

// LoadWord16 loads a 16-bit word from the data register. Use it only when 16-bit
// word or data packing is configured.
func (p *Periph) LoadWord16() uint16 {
	return (*mmio.U16)(unsafe.Pointer(&p.raw.DR)).Load()
}

// StoreByte stores a byte to the data register. Use it only when 8-bit frame is
// configured.
func (p *Periph) StoreByte(v byte) {
	(*mmio.U8)(unsafe.Pointer(&p.raw.DR)).Store(v)
}

// LoadByte loads a byte from the data register. Use it only when 8-bit frame is
// configured.
func (p *Periph) LoadByte() byte {
	return (*mmio.U8)(unsafe.Pointer(&p.raw.DR)).Load()
}
