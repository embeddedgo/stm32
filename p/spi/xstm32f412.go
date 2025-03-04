// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build stm32f412

package spi

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CR1     mmio.R32[CR1]
	CR2     mmio.R32[CR2]
	SR      mmio.R32[SR]
	DR      mmio.R32[uint32]
	CRCPR   mmio.R32[uint32]
	RXCRCR  mmio.R32[uint32]
	TXCRCR  mmio.R32[uint32]
	I2SCFGR mmio.R32[I2SCFGR]
	I2SPR   mmio.R32[I2SPR]
}

func I2S2ext() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.I2S2ext_BASE))) }
func I2S3ext() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.I2S3ext_BASE))) }
func SPI1() *Periph    { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI1_BASE))) }
func SPI2() *Periph    { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI2_BASE))) }
func SPI3() *Periph    { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI3_BASE))) }
func SPI4() *Periph    { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI4_BASE))) }
func SPI5() *Periph    { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI5_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	switch p.BaseAddr() {
	default:
		return bus.APB2
	case mmap.SPI2_BASE, mmap.SPI3_BASE:
		return bus.APB1
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
func FRF_(p *Periph) mmio.RM32[CR2]     { return mmio.RM32[CR2]{&p.CR2, FRF} }
func ERRIE_(p *Periph) mmio.RM32[CR2]   { return mmio.RM32[CR2]{&p.CR2, ERRIE} }
func RXNEIE_(p *Periph) mmio.RM32[CR2]  { return mmio.RM32[CR2]{&p.CR2, RXNEIE} }
func TXEIE_(p *Periph) mmio.RM32[CR2]   { return mmio.RM32[CR2]{&p.CR2, TXEIE} }

type SR uint32

func RXNE_(p *Periph) mmio.RM32[SR]   { return mmio.RM32[SR]{&p.SR, RXNE} }
func TXE_(p *Periph) mmio.RM32[SR]    { return mmio.RM32[SR]{&p.SR, TXE} }
func CHSIDE_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, CHSIDE} }
func UDR_(p *Periph) mmio.RM32[SR]    { return mmio.RM32[SR]{&p.SR, UDR} }
func CRCERR_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, CRCERR} }
func MODF_(p *Periph) mmio.RM32[SR]   { return mmio.RM32[SR]{&p.SR, MODF} }
func OVR_(p *Periph) mmio.RM32[SR]    { return mmio.RM32[SR]{&p.SR, OVR} }
func BSY_(p *Periph) mmio.RM32[SR]    { return mmio.RM32[SR]{&p.SR, BSY} }
func TIFRFE_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, TIFRFE} }

type I2SCFGR uint32

func CHLEN_(p *Periph) mmio.RM32[I2SCFGR]   { return mmio.RM32[I2SCFGR]{&p.I2SCFGR, CHLEN} }
func DATLEN_(p *Periph) mmio.RM32[I2SCFGR]  { return mmio.RM32[I2SCFGR]{&p.I2SCFGR, DATLEN} }
func CKPOL_(p *Periph) mmio.RM32[I2SCFGR]   { return mmio.RM32[I2SCFGR]{&p.I2SCFGR, CKPOL} }
func I2SSTD_(p *Periph) mmio.RM32[I2SCFGR]  { return mmio.RM32[I2SCFGR]{&p.I2SCFGR, I2SSTD} }
func PCMSYNC_(p *Periph) mmio.RM32[I2SCFGR] { return mmio.RM32[I2SCFGR]{&p.I2SCFGR, PCMSYNC} }
func I2SCFG_(p *Periph) mmio.RM32[I2SCFGR]  { return mmio.RM32[I2SCFGR]{&p.I2SCFGR, I2SCFG} }
func I2SE_(p *Periph) mmio.RM32[I2SCFGR]    { return mmio.RM32[I2SCFGR]{&p.I2SCFGR, I2SE} }
func I2SMOD_(p *Periph) mmio.RM32[I2SCFGR]  { return mmio.RM32[I2SCFGR]{&p.I2SCFGR, I2SMOD} }

type I2SPR uint32

func I2SDIV_(p *Periph) mmio.RM32[I2SPR] { return mmio.RM32[I2SPR]{&p.I2SPR, I2SDIV} }
func ODD_(p *Periph) mmio.RM32[I2SPR]    { return mmio.RM32[I2SPR]{&p.I2SPR, ODD} }
func MCKOE_(p *Periph) mmio.RM32[I2SPR]  { return mmio.RM32[I2SPR]{&p.I2SPR, MCKOE} }
