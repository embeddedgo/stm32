// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build stm32l4x6

package spi

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CR1    mmio.R32[CR1]
	CR2    mmio.R32[CR2]
	SR     mmio.R32[SR]
	DR     mmio.R32[uint32]
	CRCPR  mmio.R32[uint32]
	RXCRCR mmio.R32[uint32]
	TXCRCR mmio.R32[uint32]
}

func SPI1() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI1_BASE))) }
func SPI2() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI2_BASE))) }
func SPI3() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI3_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	switch p.BaseAddr() {
	default:
		return bus.APB1
	case mmap.SPI1_BASE:
		return bus.APB2
	}
}

type CR1 uint32

func CPHA_(p *Periph) mmio.RM32[CR1]     { return mmio.RM32[CR1]{&p.CR1, CPHA} }
func CPOL_(p *Periph) mmio.RM32[CR1]     { return mmio.RM32[CR1]{&p.CR1, CPOL} }
func MSTR_(p *Periph) mmio.RM32[CR1]     { return mmio.RM32[CR1]{&p.CR1, MSTR} }
func BR_(p *Periph) mmio.RM32[CR1]       { return mmio.RM32[CR1]{&p.CR1, BR} }
func SPE_(p *Periph) mmio.RM32[CR1]      { return mmio.RM32[CR1]{&p.CR1, SPE} }
func LSBFIRST_(p *Periph) mmio.RM32[CR1] { return mmio.RM32[CR1]{&p.CR1, LSBFIRST} }
func SSI_(p *Periph) mmio.RM32[CR1]      { return mmio.RM32[CR1]{&p.CR1, SSI} }
func SSM_(p *Periph) mmio.RM32[CR1]      { return mmio.RM32[CR1]{&p.CR1, SSM} }
func RXONLY_(p *Periph) mmio.RM32[CR1]   { return mmio.RM32[CR1]{&p.CR1, RXONLY} }
func DFF_(p *Periph) mmio.RM32[CR1]      { return mmio.RM32[CR1]{&p.CR1, DFF} }
func CRCNEXT_(p *Periph) mmio.RM32[CR1]  { return mmio.RM32[CR1]{&p.CR1, CRCNEXT} }
func CRCEN_(p *Periph) mmio.RM32[CR1]    { return mmio.RM32[CR1]{&p.CR1, CRCEN} }
func BIDIOE_(p *Periph) mmio.RM32[CR1]   { return mmio.RM32[CR1]{&p.CR1, BIDIOE} }
func BIDIMODE_(p *Periph) mmio.RM32[CR1] { return mmio.RM32[CR1]{&p.CR1, BIDIMODE} }

type CR2 uint32

func RXDMAEN_(p *Periph) mmio.RM32[CR2] { return mmio.RM32[CR2]{&p.CR2, RXDMAEN} }
func TXDMAEN_(p *Periph) mmio.RM32[CR2] { return mmio.RM32[CR2]{&p.CR2, TXDMAEN} }
func SSOE_(p *Periph) mmio.RM32[CR2]    { return mmio.RM32[CR2]{&p.CR2, SSOE} }
func NSSP_(p *Periph) mmio.RM32[CR2]    { return mmio.RM32[CR2]{&p.CR2, NSSP} }
func FRF_(p *Periph) mmio.RM32[CR2]     { return mmio.RM32[CR2]{&p.CR2, FRF} }
func ERRIE_(p *Periph) mmio.RM32[CR2]   { return mmio.RM32[CR2]{&p.CR2, ERRIE} }
func RXNEIE_(p *Periph) mmio.RM32[CR2]  { return mmio.RM32[CR2]{&p.CR2, RXNEIE} }
func TXEIE_(p *Periph) mmio.RM32[CR2]   { return mmio.RM32[CR2]{&p.CR2, TXEIE} }
func DS_(p *Periph) mmio.RM32[CR2]      { return mmio.RM32[CR2]{&p.CR2, DS} }
func FRXTH_(p *Periph) mmio.RM32[CR2]   { return mmio.RM32[CR2]{&p.CR2, FRXTH} }
func LDMA_RX_(p *Periph) mmio.RM32[CR2] { return mmio.RM32[CR2]{&p.CR2, LDMA_RX} }
func LDMA_TX_(p *Periph) mmio.RM32[CR2] { return mmio.RM32[CR2]{&p.CR2, LDMA_TX} }

type SR uint32

func RXNE_(p *Periph) mmio.RM32[SR]   { return mmio.RM32[SR]{&p.SR, RXNE} }
func TXE_(p *Periph) mmio.RM32[SR]    { return mmio.RM32[SR]{&p.SR, TXE} }
func CRCERR_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, CRCERR} }
func MODF_(p *Periph) mmio.RM32[SR]   { return mmio.RM32[SR]{&p.SR, MODF} }
func OVR_(p *Periph) mmio.RM32[SR]    { return mmio.RM32[SR]{&p.SR, OVR} }
func BSY_(p *Periph) mmio.RM32[SR]    { return mmio.RM32[SR]{&p.SR, BSY} }
func TIFRFE_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, TIFRFE} }
func FRLVL_(p *Periph) mmio.RM32[SR]  { return mmio.RM32[SR]{&p.SR, FRLVL} }
func FTLVL_(p *Periph) mmio.RM32[SR]  { return mmio.RM32[SR]{&p.SR, FTLVL} }
