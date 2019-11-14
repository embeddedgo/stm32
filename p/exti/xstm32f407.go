// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32f407

package exti

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	IMR   RIMR
	EMR   REMR
	RTSR  RRTSR
	FTSR  RFTSR
	SWIER RSWIER
	PR    RPR
}

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func EXTI() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.EXTI_BASE))) }

type IMR uint32

type RIMR struct{ mmio.U32 }

func (r *RIMR) Bits(mask IMR) IMR     { return IMR(r.U32.Bits(uint32(mask))) }
func (r *RIMR) StoreBits(mask, b IMR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIMR) SetBits(mask IMR)      { r.U32.SetBits(uint32(mask)) }
func (r *RIMR) ClearBits(mask IMR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIMR) Load() IMR             { return IMR(r.U32.Load()) }
func (r *RIMR) Store(b IMR)           { r.U32.Store(uint32(b)) }

type RMIMR struct{ mmio.UM32 }

func (rm RMIMR) Load() IMR   { return IMR(rm.UM32.Load()) }
func (rm RMIMR) Store(b IMR) { rm.UM32.Store(uint32(b)) }

type EMR uint32

type REMR struct{ mmio.U32 }

func (r *REMR) Bits(mask EMR) EMR     { return EMR(r.U32.Bits(uint32(mask))) }
func (r *REMR) StoreBits(mask, b EMR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REMR) SetBits(mask EMR)      { r.U32.SetBits(uint32(mask)) }
func (r *REMR) ClearBits(mask EMR)    { r.U32.ClearBits(uint32(mask)) }
func (r *REMR) Load() EMR             { return EMR(r.U32.Load()) }
func (r *REMR) Store(b EMR)           { r.U32.Store(uint32(b)) }

type RMEMR struct{ mmio.UM32 }

func (rm RMEMR) Load() EMR   { return EMR(rm.UM32.Load()) }
func (rm RMEMR) Store(b EMR) { rm.UM32.Store(uint32(b)) }

type RTSR uint32

type RRTSR struct{ mmio.U32 }

func (r *RRTSR) Bits(mask RTSR) RTSR    { return RTSR(r.U32.Bits(uint32(mask))) }
func (r *RRTSR) StoreBits(mask, b RTSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRTSR) SetBits(mask RTSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRTSR) ClearBits(mask RTSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRTSR) Load() RTSR             { return RTSR(r.U32.Load()) }
func (r *RRTSR) Store(b RTSR)           { r.U32.Store(uint32(b)) }

type RMRTSR struct{ mmio.UM32 }

func (rm RMRTSR) Load() RTSR   { return RTSR(rm.UM32.Load()) }
func (rm RMRTSR) Store(b RTSR) { rm.UM32.Store(uint32(b)) }

type FTSR uint32

type RFTSR struct{ mmio.U32 }

func (r *RFTSR) Bits(mask FTSR) FTSR    { return FTSR(r.U32.Bits(uint32(mask))) }
func (r *RFTSR) StoreBits(mask, b FTSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFTSR) SetBits(mask FTSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFTSR) ClearBits(mask FTSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFTSR) Load() FTSR             { return FTSR(r.U32.Load()) }
func (r *RFTSR) Store(b FTSR)           { r.U32.Store(uint32(b)) }

type RMFTSR struct{ mmio.UM32 }

func (rm RMFTSR) Load() FTSR   { return FTSR(rm.UM32.Load()) }
func (rm RMFTSR) Store(b FTSR) { rm.UM32.Store(uint32(b)) }

type SWIER uint32

type RSWIER struct{ mmio.U32 }

func (r *RSWIER) Bits(mask SWIER) SWIER   { return SWIER(r.U32.Bits(uint32(mask))) }
func (r *RSWIER) StoreBits(mask, b SWIER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSWIER) SetBits(mask SWIER)      { r.U32.SetBits(uint32(mask)) }
func (r *RSWIER) ClearBits(mask SWIER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSWIER) Load() SWIER             { return SWIER(r.U32.Load()) }
func (r *RSWIER) Store(b SWIER)           { r.U32.Store(uint32(b)) }

type RMSWIER struct{ mmio.UM32 }

func (rm RMSWIER) Load() SWIER   { return SWIER(rm.UM32.Load()) }
func (rm RMSWIER) Store(b SWIER) { rm.UM32.Store(uint32(b)) }

type PR uint32

type RPR struct{ mmio.U32 }

func (r *RPR) Bits(mask PR) PR      { return PR(r.U32.Bits(uint32(mask))) }
func (r *RPR) StoreBits(mask, b PR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPR) SetBits(mask PR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPR) ClearBits(mask PR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPR) Load() PR             { return PR(r.U32.Load()) }
func (r *RPR) Store(b PR)           { r.U32.Store(uint32(b)) }

type RMPR struct{ mmio.UM32 }

func (rm RMPR) Load() PR   { return PR(rm.UM32.Load()) }
func (rm RMPR) Store(b PR) { rm.UM32.Store(uint32(b)) }
