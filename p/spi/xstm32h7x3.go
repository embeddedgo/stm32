// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32h7x3

package spi

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CR1     RCR1
	CR2     RCR2
	CFG1    RCFG1
	CFG2    RCFG2
	IER     RIER
	SR      RSR
	IFCR    RIFCR
	_       uint32
	TXDR    RTXDR
	_       [3]uint32
	RXDR    RRXDR
	_       [3]uint32
	CRCPOLY RCRCPOLY
	TXCRC   RTXCRC
	RXCRC   RRXCRC
	UDRDR   RUDRDR
	CGFR    RCGFR
}

func SPI1() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI1_BASE))) }
func SPI2() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI2_BASE))) }
func SPI3() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI3_BASE))) }
func SPI4() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI4_BASE))) }
func SPI5() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI5_BASE))) }
func SPI6() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI6_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	switch p.BaseAddr() {
	default:
		return bus.APB2
	case mmap.SPI2_BASE, mmap.SPI3_BASE:
		return bus.APB1
	case mmap.SPI6_BASE:
		return bus.APB4
	}
}

type CR1 uint32

type RCR1 struct{ mmio.U32 }

func (r *RCR1) LoadBits(mask CR1) CR1 { return CR1(r.U32.LoadBits(uint32(mask))) }
func (r *RCR1) StoreBits(mask, b CR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR1) SetBits(mask CR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR1) ClearBits(mask CR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR1) Load() CR1             { return CR1(r.U32.Load()) }
func (r *RCR1) Store(b CR1)           { r.U32.Store(uint32(b)) }

type RMCR1 struct{ mmio.UM32 }

func (rm RMCR1) Load() CR1   { return CR1(rm.UM32.Load()) }
func (rm RMCR1) Store(b CR1) { rm.UM32.Store(uint32(b)) }

func SPE_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SPE)}}
}

func MASRX_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(MASRX)}}
}

func CSTART_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CSTART)}}
}

func CSUSP_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CSUSP)}}
}

func HDDIR_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(HDDIR)}}
}

func SSI_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SSI)}}
}

func CRC33_17_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CRC33_17)}}
}

func RCRCI_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(RCRCI)}}
}

func TCRCI_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(TCRCI)}}
}

func IOLOCK_(p *Periph) RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(IOLOCK)}}
}

type CR2 uint32

type RCR2 struct{ mmio.U32 }

func (r *RCR2) LoadBits(mask CR2) CR2 { return CR2(r.U32.LoadBits(uint32(mask))) }
func (r *RCR2) StoreBits(mask, b CR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR2) SetBits(mask CR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR2) ClearBits(mask CR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR2) Load() CR2             { return CR2(r.U32.Load()) }
func (r *RCR2) Store(b CR2)           { r.U32.Store(uint32(b)) }

type RMCR2 struct{ mmio.UM32 }

func (rm RMCR2) Load() CR2   { return CR2(rm.UM32.Load()) }
func (rm RMCR2) Store(b CR2) { rm.UM32.Store(uint32(b)) }

func TSIZE_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(TSIZE)}}
}

func TSER_(p *Periph) RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(TSER)}}
}

type CFG1 uint32

type RCFG1 struct{ mmio.U32 }

func (r *RCFG1) LoadBits(mask CFG1) CFG1 { return CFG1(r.U32.LoadBits(uint32(mask))) }
func (r *RCFG1) StoreBits(mask, b CFG1)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFG1) SetBits(mask CFG1)       { r.U32.SetBits(uint32(mask)) }
func (r *RCFG1) ClearBits(mask CFG1)     { r.U32.ClearBits(uint32(mask)) }
func (r *RCFG1) Load() CFG1              { return CFG1(r.U32.Load()) }
func (r *RCFG1) Store(b CFG1)            { r.U32.Store(uint32(b)) }

type RMCFG1 struct{ mmio.UM32 }

func (rm RMCFG1) Load() CFG1   { return CFG1(rm.UM32.Load()) }
func (rm RMCFG1) Store(b CFG1) { rm.UM32.Store(uint32(b)) }

func DSIZE_(p *Periph) RMCFG1 {
	return RMCFG1{mmio.UM32{&p.CFG1.U32, uint32(DSIZE)}}
}

func FTHVL_(p *Periph) RMCFG1 {
	return RMCFG1{mmio.UM32{&p.CFG1.U32, uint32(FTHVL)}}
}

func UDRCFG_(p *Periph) RMCFG1 {
	return RMCFG1{mmio.UM32{&p.CFG1.U32, uint32(UDRCFG)}}
}

func UDRDET_(p *Periph) RMCFG1 {
	return RMCFG1{mmio.UM32{&p.CFG1.U32, uint32(UDRDET)}}
}

func RXDMAEN_(p *Periph) RMCFG1 {
	return RMCFG1{mmio.UM32{&p.CFG1.U32, uint32(RXDMAEN)}}
}

func TXDMAEN_(p *Periph) RMCFG1 {
	return RMCFG1{mmio.UM32{&p.CFG1.U32, uint32(TXDMAEN)}}
}

func CRCSIZE_(p *Periph) RMCFG1 {
	return RMCFG1{mmio.UM32{&p.CFG1.U32, uint32(CRCSIZE)}}
}

func CRCEN_(p *Periph) RMCFG1 {
	return RMCFG1{mmio.UM32{&p.CFG1.U32, uint32(CRCEN)}}
}

func MBR_(p *Periph) RMCFG1 {
	return RMCFG1{mmio.UM32{&p.CFG1.U32, uint32(MBR)}}
}

type CFG2 uint32

type RCFG2 struct{ mmio.U32 }

func (r *RCFG2) LoadBits(mask CFG2) CFG2 { return CFG2(r.U32.LoadBits(uint32(mask))) }
func (r *RCFG2) StoreBits(mask, b CFG2)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFG2) SetBits(mask CFG2)       { r.U32.SetBits(uint32(mask)) }
func (r *RCFG2) ClearBits(mask CFG2)     { r.U32.ClearBits(uint32(mask)) }
func (r *RCFG2) Load() CFG2              { return CFG2(r.U32.Load()) }
func (r *RCFG2) Store(b CFG2)            { r.U32.Store(uint32(b)) }

type RMCFG2 struct{ mmio.UM32 }

func (rm RMCFG2) Load() CFG2   { return CFG2(rm.UM32.Load()) }
func (rm RMCFG2) Store(b CFG2) { rm.UM32.Store(uint32(b)) }

func MSSI_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(MSSI)}}
}

func MIDI_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(MIDI)}}
}

func IOSWP_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(IOSWP)}}
}

func COMM_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(COMM)}}
}

func SP_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(SP)}}
}

func MASTER_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(MASTER)}}
}

func LSBFRST_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(LSBFRST)}}
}

func CPHA_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(CPHA)}}
}

func CPOL_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(CPOL)}}
}

func SSM_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(SSM)}}
}

func SSIOP_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(SSIOP)}}
}

func SSOE_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(SSOE)}}
}

func SSOM_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(SSOM)}}
}

func AFCNTR_(p *Periph) RMCFG2 {
	return RMCFG2{mmio.UM32{&p.CFG2.U32, uint32(AFCNTR)}}
}

type IER uint32

type RIER struct{ mmio.U32 }

func (r *RIER) LoadBits(mask IER) IER { return IER(r.U32.LoadBits(uint32(mask))) }
func (r *RIER) StoreBits(mask, b IER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIER) SetBits(mask IER)      { r.U32.SetBits(uint32(mask)) }
func (r *RIER) ClearBits(mask IER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIER) Load() IER             { return IER(r.U32.Load()) }
func (r *RIER) Store(b IER)           { r.U32.Store(uint32(b)) }

type RMIER struct{ mmio.UM32 }

func (rm RMIER) Load() IER   { return IER(rm.UM32.Load()) }
func (rm RMIER) Store(b IER) { rm.UM32.Store(uint32(b)) }

func RXPIE_(p *Periph) RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(RXPIE)}}
}

func TXPIE_(p *Periph) RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TXPIE)}}
}

func DPXPIE_(p *Periph) RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(DPXPIE)}}
}

func EOTIE_(p *Periph) RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(EOTIE)}}
}

func TXTFIE_(p *Periph) RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TXTFIE)}}
}

func UDRIE_(p *Periph) RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(UDRIE)}}
}

func OVRIE_(p *Periph) RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(OVRIE)}}
}

func CRCEIE_(p *Periph) RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(CRCEIE)}}
}

func TIFREIE_(p *Periph) RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TIFREIE)}}
}

func MODFIE_(p *Periph) RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(MODFIE)}}
}

func TSERFIE_(p *Periph) RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TSERFIE)}}
}

type SR uint32

type RSR struct{ mmio.U32 }

func (r *RSR) LoadBits(mask SR) SR  { return SR(r.U32.LoadBits(uint32(mask))) }
func (r *RSR) StoreBits(mask, b SR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR) SetBits(mask SR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR) ClearBits(mask SR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR) Load() SR             { return SR(r.U32.Load()) }
func (r *RSR) Store(b SR)           { r.U32.Store(uint32(b)) }

type RMSR struct{ mmio.UM32 }

func (rm RMSR) Load() SR   { return SR(rm.UM32.Load()) }
func (rm RMSR) Store(b SR) { rm.UM32.Store(uint32(b)) }

func RXP_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(RXP)}}
}

func TXP_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TXP)}}
}

func DXP_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(DXP)}}
}

func EOT_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(EOT)}}
}

func TXTF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TXTF)}}
}

func UDR_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(UDR)}}
}

func OVR_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(OVR)}}
}

func CRCE_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CRCE)}}
}

func TIFRE_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TIFRE)}}
}

func MODF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(MODF)}}
}

func TSERF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TSERF)}}
}

func SUSP_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(SUSP)}}
}

func TXC_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TXC)}}
}

func RXPLVL_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(RXPLVL)}}
}

func RXWNE_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(RXWNE)}}
}

func CTSIZE_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CTSIZE)}}
}

type IFCR uint32

type RIFCR struct{ mmio.U32 }

func (r *RIFCR) LoadBits(mask IFCR) IFCR { return IFCR(r.U32.LoadBits(uint32(mask))) }
func (r *RIFCR) StoreBits(mask, b IFCR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIFCR) SetBits(mask IFCR)       { r.U32.SetBits(uint32(mask)) }
func (r *RIFCR) ClearBits(mask IFCR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RIFCR) Load() IFCR              { return IFCR(r.U32.Load()) }
func (r *RIFCR) Store(b IFCR)            { r.U32.Store(uint32(b)) }

type RMIFCR struct{ mmio.UM32 }

func (rm RMIFCR) Load() IFCR   { return IFCR(rm.UM32.Load()) }
func (rm RMIFCR) Store(b IFCR) { rm.UM32.Store(uint32(b)) }

func EOTC_(p *Periph) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(EOTC)}}
}

func TXTFC_(p *Periph) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(TXTFC)}}
}

func UDRC_(p *Periph) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(UDRC)}}
}

func OVRC_(p *Periph) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(OVRC)}}
}

func CRCEC_(p *Periph) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(CRCEC)}}
}

func TIFREC_(p *Periph) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(TIFREC)}}
}

func MODFC_(p *Periph) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(MODFC)}}
}

func TSERFC_(p *Periph) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(TSERFC)}}
}

func SUSPC_(p *Periph) RMIFCR {
	return RMIFCR{mmio.UM32{&p.IFCR.U32, uint32(SUSPC)}}
}

type TXDR uint32

type RTXDR struct{ mmio.U32 }

func (r *RTXDR) LoadBits(mask TXDR) TXDR { return TXDR(r.U32.LoadBits(uint32(mask))) }
func (r *RTXDR) StoreBits(mask, b TXDR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTXDR) SetBits(mask TXDR)       { r.U32.SetBits(uint32(mask)) }
func (r *RTXDR) ClearBits(mask TXDR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RTXDR) Load() TXDR              { return TXDR(r.U32.Load()) }
func (r *RTXDR) Store(b TXDR)            { r.U32.Store(uint32(b)) }

type RMTXDR struct{ mmio.UM32 }

func (rm RMTXDR) Load() TXDR   { return TXDR(rm.UM32.Load()) }
func (rm RMTXDR) Store(b TXDR) { rm.UM32.Store(uint32(b)) }

func TXDR_(p *Periph) RMTXDR {
	return RMTXDR{mmio.UM32{&p.TXDR.U32, uint32(TXDR)}}
}

type RXDR uint32

type RRXDR struct{ mmio.U32 }

func (r *RRXDR) LoadBits(mask RXDR) RXDR { return RXDR(r.U32.LoadBits(uint32(mask))) }
func (r *RRXDR) StoreBits(mask, b RXDR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRXDR) SetBits(mask RXDR)       { r.U32.SetBits(uint32(mask)) }
func (r *RRXDR) ClearBits(mask RXDR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RRXDR) Load() RXDR              { return RXDR(r.U32.Load()) }
func (r *RRXDR) Store(b RXDR)            { r.U32.Store(uint32(b)) }

type RMRXDR struct{ mmio.UM32 }

func (rm RMRXDR) Load() RXDR   { return RXDR(rm.UM32.Load()) }
func (rm RMRXDR) Store(b RXDR) { rm.UM32.Store(uint32(b)) }

func RXDR_(p *Periph) RMRXDR {
	return RMRXDR{mmio.UM32{&p.RXDR.U32, uint32(RXDR)}}
}

type CRCPOLY uint32

type RCRCPOLY struct{ mmio.U32 }

func (r *RCRCPOLY) LoadBits(mask CRCPOLY) CRCPOLY { return CRCPOLY(r.U32.LoadBits(uint32(mask))) }
func (r *RCRCPOLY) StoreBits(mask, b CRCPOLY)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCRCPOLY) SetBits(mask CRCPOLY)          { r.U32.SetBits(uint32(mask)) }
func (r *RCRCPOLY) ClearBits(mask CRCPOLY)        { r.U32.ClearBits(uint32(mask)) }
func (r *RCRCPOLY) Load() CRCPOLY                 { return CRCPOLY(r.U32.Load()) }
func (r *RCRCPOLY) Store(b CRCPOLY)               { r.U32.Store(uint32(b)) }

type RMCRCPOLY struct{ mmio.UM32 }

func (rm RMCRCPOLY) Load() CRCPOLY   { return CRCPOLY(rm.UM32.Load()) }
func (rm RMCRCPOLY) Store(b CRCPOLY) { rm.UM32.Store(uint32(b)) }

func CRCPOLY_(p *Periph) RMCRCPOLY {
	return RMCRCPOLY{mmio.UM32{&p.CRCPOLY.U32, uint32(CRCPOLY)}}
}

type TXCRC uint32

type RTXCRC struct{ mmio.U32 }

func (r *RTXCRC) LoadBits(mask TXCRC) TXCRC { return TXCRC(r.U32.LoadBits(uint32(mask))) }
func (r *RTXCRC) StoreBits(mask, b TXCRC)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTXCRC) SetBits(mask TXCRC)        { r.U32.SetBits(uint32(mask)) }
func (r *RTXCRC) ClearBits(mask TXCRC)      { r.U32.ClearBits(uint32(mask)) }
func (r *RTXCRC) Load() TXCRC               { return TXCRC(r.U32.Load()) }
func (r *RTXCRC) Store(b TXCRC)             { r.U32.Store(uint32(b)) }

type RMTXCRC struct{ mmio.UM32 }

func (rm RMTXCRC) Load() TXCRC   { return TXCRC(rm.UM32.Load()) }
func (rm RMTXCRC) Store(b TXCRC) { rm.UM32.Store(uint32(b)) }

func TXCRC_(p *Periph) RMTXCRC {
	return RMTXCRC{mmio.UM32{&p.TXCRC.U32, uint32(TXCRC)}}
}

type RXCRC uint32

type RRXCRC struct{ mmio.U32 }

func (r *RRXCRC) LoadBits(mask RXCRC) RXCRC { return RXCRC(r.U32.LoadBits(uint32(mask))) }
func (r *RRXCRC) StoreBits(mask, b RXCRC)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRXCRC) SetBits(mask RXCRC)        { r.U32.SetBits(uint32(mask)) }
func (r *RRXCRC) ClearBits(mask RXCRC)      { r.U32.ClearBits(uint32(mask)) }
func (r *RRXCRC) Load() RXCRC               { return RXCRC(r.U32.Load()) }
func (r *RRXCRC) Store(b RXCRC)             { r.U32.Store(uint32(b)) }

type RMRXCRC struct{ mmio.UM32 }

func (rm RMRXCRC) Load() RXCRC   { return RXCRC(rm.UM32.Load()) }
func (rm RMRXCRC) Store(b RXCRC) { rm.UM32.Store(uint32(b)) }

func RXCRC_(p *Periph) RMRXCRC {
	return RMRXCRC{mmio.UM32{&p.RXCRC.U32, uint32(RXCRC)}}
}

type UDRDR uint32

type RUDRDR struct{ mmio.U32 }

func (r *RUDRDR) LoadBits(mask UDRDR) UDRDR { return UDRDR(r.U32.LoadBits(uint32(mask))) }
func (r *RUDRDR) StoreBits(mask, b UDRDR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RUDRDR) SetBits(mask UDRDR)        { r.U32.SetBits(uint32(mask)) }
func (r *RUDRDR) ClearBits(mask UDRDR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RUDRDR) Load() UDRDR               { return UDRDR(r.U32.Load()) }
func (r *RUDRDR) Store(b UDRDR)             { r.U32.Store(uint32(b)) }

type RMUDRDR struct{ mmio.UM32 }

func (rm RMUDRDR) Load() UDRDR   { return UDRDR(rm.UM32.Load()) }
func (rm RMUDRDR) Store(b UDRDR) { rm.UM32.Store(uint32(b)) }

func UDRDR_(p *Periph) RMUDRDR {
	return RMUDRDR{mmio.UM32{&p.UDRDR.U32, uint32(UDRDR)}}
}

type CGFR uint32

type RCGFR struct{ mmio.U32 }

func (r *RCGFR) LoadBits(mask CGFR) CGFR { return CGFR(r.U32.LoadBits(uint32(mask))) }
func (r *RCGFR) StoreBits(mask, b CGFR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCGFR) SetBits(mask CGFR)       { r.U32.SetBits(uint32(mask)) }
func (r *RCGFR) ClearBits(mask CGFR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RCGFR) Load() CGFR              { return CGFR(r.U32.Load()) }
func (r *RCGFR) Store(b CGFR)            { r.U32.Store(uint32(b)) }

type RMCGFR struct{ mmio.UM32 }

func (rm RMCGFR) Load() CGFR   { return CGFR(rm.UM32.Load()) }
func (rm RMCGFR) Store(b CGFR) { rm.UM32.Store(uint32(b)) }

func I2SMOD_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(I2SMOD)}}
}

func I2SCFG_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(I2SCFG)}}
}

func I2SSTD_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(I2SSTD)}}
}

func PCMSYNC_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(PCMSYNC)}}
}

func DATLEN_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(DATLEN)}}
}

func CHLEN_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(CHLEN)}}
}

func CKPOL_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(CKPOL)}}
}

func FIXCH_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(FIXCH)}}
}

func WSINV_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(WSINV)}}
}

func DATFMT_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(DATFMT)}}
}

func I2SDIV_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(I2SDIV)}}
}

func ODD_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(ODD)}}
}

func MCKOE_(p *Periph) RMCGFR {
	return RMCGFR{mmio.UM32{&p.CGFR.U32, uint32(MCKOE)}}
}
