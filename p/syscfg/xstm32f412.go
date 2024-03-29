// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build stm32f412

package syscfg

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	MEMRM      mmio.R32[MEMRM]
	PMC        mmio.R32[PMC]
	EXTICR     [4]mmio.R32[uint32]
	_          [2]uint32
	CMPCR      mmio.R32[CMPCR]
	_          [2]uint32
	I2C_BUFOUT mmio.R32[I2C_BUFOUT]
}

func SYSCFG() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SYSCFG_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	return bus.APB2
}

type MEMRM uint32

func MEM_MODE_(p *Periph) mmio.RM32[MEMRM] { return mmio.RM32[MEMRM]{&p.MEMRM, MEM_MODE} }

type PMC uint32

func ADC1DC2_(p *Periph) mmio.RM32[PMC] { return mmio.RM32[PMC]{&p.PMC, ADC1DC2} }

type CMPCR uint32

func CMP_PD_(p *Periph) mmio.RM32[CMPCR] { return mmio.RM32[CMPCR]{&p.CMPCR, CMP_PD} }
func READY_(p *Periph) mmio.RM32[CMPCR]  { return mmio.RM32[CMPCR]{&p.CMPCR, READY} }

type I2C_BUFOUT uint32

func I2C4SCL_(p *Periph) mmio.RM32[I2C_BUFOUT] { return mmio.RM32[I2C_BUFOUT]{&p.I2C_BUFOUT, I2C4SCL} }
func I2C4SDA_(p *Periph) mmio.RM32[I2C_BUFOUT] { return mmio.RM32[I2C_BUFOUT]{&p.I2C_BUFOUT, I2C4SDA} }
