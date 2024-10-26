// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build stm32g471xx

package dmamux

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CCR   [16]mmio.R32[CCR]
	_     [16]uint32
	CSR   mmio.R32[CSR]
	CFR   mmio.R32[CSR]
	_     [30]uint32
	RGCR  [8]mmio.R32[RGCR]
	_     [8]uint32
	RGSR  mmio.R32[RGSR]
	RGCFR mmio.R32[RGSR]
}

func DMAMUX() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.DMAMUX_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	return bus.AHB1
}

type CCR uint32

func DMAREQ_ID_(p *Periph, i int) mmio.RM32[CCR] { return mmio.RM32[CCR]{&p.CCR[i], DMAREQ_ID} }
func SOIE_(p *Periph, i int) mmio.RM32[CCR]      { return mmio.RM32[CCR]{&p.CCR[i], SOIE} }
func EGE_(p *Periph, i int) mmio.RM32[CCR]       { return mmio.RM32[CCR]{&p.CCR[i], EGE} }
func SE_(p *Periph, i int) mmio.RM32[CCR]        { return mmio.RM32[CCR]{&p.CCR[i], SE} }
func SPOL_(p *Periph, i int) mmio.RM32[CCR]      { return mmio.RM32[CCR]{&p.CCR[i], SPOL} }
func NBREQ_(p *Periph, i int) mmio.RM32[CCR]     { return mmio.RM32[CCR]{&p.CCR[i], NBREQ} }
func SYNC_ID_(p *Periph, i int) mmio.RM32[CCR]   { return mmio.RM32[CCR]{&p.CCR[i], SYNC_ID} }

type CSR uint32

func SOF_(p *Periph) mmio.RM32[CSR] { return mmio.RM32[CSR]{&p.CSR, SOF} }

type RGCR uint32

func SIG_ID_(p *Periph, i int) mmio.RM32[RGCR] { return mmio.RM32[RGCR]{&p.RGCR[i], SIG_ID} }
func OIE_(p *Periph, i int) mmio.RM32[RGCR]    { return mmio.RM32[RGCR]{&p.RGCR[i], OIE} }
func GE_(p *Periph, i int) mmio.RM32[RGCR]     { return mmio.RM32[RGCR]{&p.RGCR[i], GE} }
func GPOL_(p *Periph, i int) mmio.RM32[RGCR]   { return mmio.RM32[RGCR]{&p.RGCR[i], GPOL} }
func GNBREQ_(p *Periph, i int) mmio.RM32[RGCR] { return mmio.RM32[RGCR]{&p.RGCR[i], GNBREQ} }

type RGSR uint32

func OF_(p *Periph) mmio.RM32[RGSR] { return mmio.RM32[RGSR]{&p.RGSR, OF} }
