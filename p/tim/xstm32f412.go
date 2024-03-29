// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build stm32f412

package tim

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	CR1   mmio.R32[CR1]
	CR2   mmio.R32[CR2]
	SMCR  mmio.R32[SMCR]
	DIER  mmio.R32[DIER]
	SR    mmio.R32[SR]
	EGR   mmio.R32[EGR]
	CCMR1 mmio.R32[CCMR1]
	CCMR2 mmio.R32[CCMR2]
	CCER  mmio.R32[CCER]
	CNT   mmio.R32[uint32]
	PSC   mmio.R32[uint32]
	ARR   mmio.R32[uint32]
	RCR   mmio.R32[uint32]
	CCR1  mmio.R32[uint32]
	CCR2  mmio.R32[uint32]
	CCR3  mmio.R32[uint32]
	CCR4  mmio.R32[uint32]
	BDTR  mmio.R32[BDTR]
	DCR   mmio.R32[DCR]
	DMAR  mmio.R32[uint32]
}

func TIM1() *Periph  { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM1_BASE))) }
func TIM2() *Periph  { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM2_BASE))) }
func TIM3() *Periph  { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM3_BASE))) }
func TIM4() *Periph  { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM4_BASE))) }
func TIM5() *Periph  { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM5_BASE))) }
func TIM6() *Periph  { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM6_BASE))) }
func TIM7() *Periph  { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM7_BASE))) }
func TIM8() *Periph  { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM8_BASE))) }
func TIM9() *Periph  { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM9_BASE))) }
func TIM10() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM10_BASE))) }
func TIM11() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM11_BASE))) }
func TIM12() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM12_BASE))) }
func TIM13() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM13_BASE))) }
func TIM14() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.TIM14_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	switch p.BaseAddr() {
	default:
		return bus.APB1
	case mmap.TIM1_BASE, mmap.TIM8_BASE, mmap.TIM9_BASE, mmap.TIM10_BASE, mmap.TIM11_BASE:
		return bus.APB2
	}
}

type CR1 uint32

func CEN_(p *Periph) mmio.RM32[CR1]  { return mmio.RM32[CR1]{&p.CR1, CEN} }
func UDIS_(p *Periph) mmio.RM32[CR1] { return mmio.RM32[CR1]{&p.CR1, UDIS} }
func URS_(p *Periph) mmio.RM32[CR1]  { return mmio.RM32[CR1]{&p.CR1, URS} }
func OPM_(p *Periph) mmio.RM32[CR1]  { return mmio.RM32[CR1]{&p.CR1, OPM} }
func DIR_(p *Periph) mmio.RM32[CR1]  { return mmio.RM32[CR1]{&p.CR1, DIR} }
func CMS_(p *Periph) mmio.RM32[CR1]  { return mmio.RM32[CR1]{&p.CR1, CMS} }
func ARPE_(p *Periph) mmio.RM32[CR1] { return mmio.RM32[CR1]{&p.CR1, ARPE} }
func CKD_(p *Periph) mmio.RM32[CR1]  { return mmio.RM32[CR1]{&p.CR1, CKD} }

type CR2 uint32

func CCPC_(p *Periph) mmio.RM32[CR2]  { return mmio.RM32[CR2]{&p.CR2, CCPC} }
func CCUS_(p *Periph) mmio.RM32[CR2]  { return mmio.RM32[CR2]{&p.CR2, CCUS} }
func CCDS_(p *Periph) mmio.RM32[CR2]  { return mmio.RM32[CR2]{&p.CR2, CCDS} }
func MMS_(p *Periph) mmio.RM32[CR2]   { return mmio.RM32[CR2]{&p.CR2, MMS} }
func TI1S_(p *Periph) mmio.RM32[CR2]  { return mmio.RM32[CR2]{&p.CR2, TI1S} }
func OIS1_(p *Periph) mmio.RM32[CR2]  { return mmio.RM32[CR2]{&p.CR2, OIS1} }
func OIS1N_(p *Periph) mmio.RM32[CR2] { return mmio.RM32[CR2]{&p.CR2, OIS1N} }
func OIS2_(p *Periph) mmio.RM32[CR2]  { return mmio.RM32[CR2]{&p.CR2, OIS2} }
func OIS2N_(p *Periph) mmio.RM32[CR2] { return mmio.RM32[CR2]{&p.CR2, OIS2N} }
func OIS3_(p *Periph) mmio.RM32[CR2]  { return mmio.RM32[CR2]{&p.CR2, OIS3} }
func OIS3N_(p *Periph) mmio.RM32[CR2] { return mmio.RM32[CR2]{&p.CR2, OIS3N} }
func OIS4_(p *Periph) mmio.RM32[CR2]  { return mmio.RM32[CR2]{&p.CR2, OIS4} }

type SMCR uint32

func SMS_(p *Periph) mmio.RM32[SMCR]  { return mmio.RM32[SMCR]{&p.SMCR, SMS} }
func TS_(p *Periph) mmio.RM32[SMCR]   { return mmio.RM32[SMCR]{&p.SMCR, TS} }
func MSM_(p *Periph) mmio.RM32[SMCR]  { return mmio.RM32[SMCR]{&p.SMCR, MSM} }
func ETF_(p *Periph) mmio.RM32[SMCR]  { return mmio.RM32[SMCR]{&p.SMCR, ETF} }
func ETPS_(p *Periph) mmio.RM32[SMCR] { return mmio.RM32[SMCR]{&p.SMCR, ETPS} }
func ECE_(p *Periph) mmio.RM32[SMCR]  { return mmio.RM32[SMCR]{&p.SMCR, ECE} }
func ETP_(p *Periph) mmio.RM32[SMCR]  { return mmio.RM32[SMCR]{&p.SMCR, ETP} }

type DIER uint32

func UIE_(p *Periph) mmio.RM32[DIER]   { return mmio.RM32[DIER]{&p.DIER, UIE} }
func CC1IE_(p *Periph) mmio.RM32[DIER] { return mmio.RM32[DIER]{&p.DIER, CC1IE} }
func CC2IE_(p *Periph) mmio.RM32[DIER] { return mmio.RM32[DIER]{&p.DIER, CC2IE} }
func CC3IE_(p *Periph) mmio.RM32[DIER] { return mmio.RM32[DIER]{&p.DIER, CC3IE} }
func CC4IE_(p *Periph) mmio.RM32[DIER] { return mmio.RM32[DIER]{&p.DIER, CC4IE} }
func COMIE_(p *Periph) mmio.RM32[DIER] { return mmio.RM32[DIER]{&p.DIER, COMIE} }
func TIE_(p *Periph) mmio.RM32[DIER]   { return mmio.RM32[DIER]{&p.DIER, TIE} }
func BIE_(p *Periph) mmio.RM32[DIER]   { return mmio.RM32[DIER]{&p.DIER, BIE} }
func UDE_(p *Periph) mmio.RM32[DIER]   { return mmio.RM32[DIER]{&p.DIER, UDE} }
func CC1DE_(p *Periph) mmio.RM32[DIER] { return mmio.RM32[DIER]{&p.DIER, CC1DE} }
func CC2DE_(p *Periph) mmio.RM32[DIER] { return mmio.RM32[DIER]{&p.DIER, CC2DE} }
func CC3DE_(p *Periph) mmio.RM32[DIER] { return mmio.RM32[DIER]{&p.DIER, CC3DE} }
func CC4DE_(p *Periph) mmio.RM32[DIER] { return mmio.RM32[DIER]{&p.DIER, CC4DE} }
func COMDE_(p *Periph) mmio.RM32[DIER] { return mmio.RM32[DIER]{&p.DIER, COMDE} }
func TDE_(p *Periph) mmio.RM32[DIER]   { return mmio.RM32[DIER]{&p.DIER, TDE} }

type SR uint32

func UIF_(p *Periph) mmio.RM32[SR]   { return mmio.RM32[SR]{&p.SR, UIF} }
func CC1IF_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, CC1IF} }
func CC2IF_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, CC2IF} }
func CC3IF_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, CC3IF} }
func CC4IF_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, CC4IF} }
func COMIF_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, COMIF} }
func TIF_(p *Periph) mmio.RM32[SR]   { return mmio.RM32[SR]{&p.SR, TIF} }
func BIF_(p *Periph) mmio.RM32[SR]   { return mmio.RM32[SR]{&p.SR, BIF} }
func CC1OF_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, CC1OF} }
func CC2OF_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, CC2OF} }
func CC3OF_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, CC3OF} }
func CC4OF_(p *Periph) mmio.RM32[SR] { return mmio.RM32[SR]{&p.SR, CC4OF} }

type EGR uint32

func UG_(p *Periph) mmio.RM32[EGR]   { return mmio.RM32[EGR]{&p.EGR, UG} }
func CC1G_(p *Periph) mmio.RM32[EGR] { return mmio.RM32[EGR]{&p.EGR, CC1G} }
func CC2G_(p *Periph) mmio.RM32[EGR] { return mmio.RM32[EGR]{&p.EGR, CC2G} }
func CC3G_(p *Periph) mmio.RM32[EGR] { return mmio.RM32[EGR]{&p.EGR, CC3G} }
func CC4G_(p *Periph) mmio.RM32[EGR] { return mmio.RM32[EGR]{&p.EGR, CC4G} }
func COMG_(p *Periph) mmio.RM32[EGR] { return mmio.RM32[EGR]{&p.EGR, COMG} }
func TG_(p *Periph) mmio.RM32[EGR]   { return mmio.RM32[EGR]{&p.EGR, TG} }
func BG_(p *Periph) mmio.RM32[EGR]   { return mmio.RM32[EGR]{&p.EGR, BG} }

type CCMR1 uint32

func CC1S_(p *Periph) mmio.RM32[CCMR1]   { return mmio.RM32[CCMR1]{&p.CCMR1, CC1S} }
func OC1FE_(p *Periph) mmio.RM32[CCMR1]  { return mmio.RM32[CCMR1]{&p.CCMR1, OC1FE} }
func OC1PE_(p *Periph) mmio.RM32[CCMR1]  { return mmio.RM32[CCMR1]{&p.CCMR1, OC1PE} }
func OC1M_(p *Periph) mmio.RM32[CCMR1]   { return mmio.RM32[CCMR1]{&p.CCMR1, OC1M} }
func OC1CE_(p *Periph) mmio.RM32[CCMR1]  { return mmio.RM32[CCMR1]{&p.CCMR1, OC1CE} }
func CC2S_(p *Periph) mmio.RM32[CCMR1]   { return mmio.RM32[CCMR1]{&p.CCMR1, CC2S} }
func OC2FE_(p *Periph) mmio.RM32[CCMR1]  { return mmio.RM32[CCMR1]{&p.CCMR1, OC2FE} }
func OC2PE_(p *Periph) mmio.RM32[CCMR1]  { return mmio.RM32[CCMR1]{&p.CCMR1, OC2PE} }
func OC2M_(p *Periph) mmio.RM32[CCMR1]   { return mmio.RM32[CCMR1]{&p.CCMR1, OC2M} }
func OC2CE_(p *Periph) mmio.RM32[CCMR1]  { return mmio.RM32[CCMR1]{&p.CCMR1, OC2CE} }
func ICPCS_(p *Periph) mmio.RM32[CCMR1]  { return mmio.RM32[CCMR1]{&p.CCMR1, ICPCS} }
func IC1F_(p *Periph) mmio.RM32[CCMR1]   { return mmio.RM32[CCMR1]{&p.CCMR1, IC1F} }
func IC2PCS_(p *Periph) mmio.RM32[CCMR1] { return mmio.RM32[CCMR1]{&p.CCMR1, IC2PCS} }
func IC2F_(p *Periph) mmio.RM32[CCMR1]   { return mmio.RM32[CCMR1]{&p.CCMR1, IC2F} }

type CCMR2 uint32

func CC3S_(p *Periph) mmio.RM32[CCMR2]   { return mmio.RM32[CCMR2]{&p.CCMR2, CC3S} }
func OC3FE_(p *Periph) mmio.RM32[CCMR2]  { return mmio.RM32[CCMR2]{&p.CCMR2, OC3FE} }
func OC3PE_(p *Periph) mmio.RM32[CCMR2]  { return mmio.RM32[CCMR2]{&p.CCMR2, OC3PE} }
func OC3M_(p *Periph) mmio.RM32[CCMR2]   { return mmio.RM32[CCMR2]{&p.CCMR2, OC3M} }
func OC3CE_(p *Periph) mmio.RM32[CCMR2]  { return mmio.RM32[CCMR2]{&p.CCMR2, OC3CE} }
func CC4S_(p *Periph) mmio.RM32[CCMR2]   { return mmio.RM32[CCMR2]{&p.CCMR2, CC4S} }
func OC4FE_(p *Periph) mmio.RM32[CCMR2]  { return mmio.RM32[CCMR2]{&p.CCMR2, OC4FE} }
func OC4PE_(p *Periph) mmio.RM32[CCMR2]  { return mmio.RM32[CCMR2]{&p.CCMR2, OC4PE} }
func OC4M_(p *Periph) mmio.RM32[CCMR2]   { return mmio.RM32[CCMR2]{&p.CCMR2, OC4M} }
func OC4CE_(p *Periph) mmio.RM32[CCMR2]  { return mmio.RM32[CCMR2]{&p.CCMR2, OC4CE} }
func IC3PSC_(p *Periph) mmio.RM32[CCMR2] { return mmio.RM32[CCMR2]{&p.CCMR2, IC3PSC} }
func IC3F_(p *Periph) mmio.RM32[CCMR2]   { return mmio.RM32[CCMR2]{&p.CCMR2, IC3F} }
func IC4PSC_(p *Periph) mmio.RM32[CCMR2] { return mmio.RM32[CCMR2]{&p.CCMR2, IC4PSC} }
func IC4F_(p *Periph) mmio.RM32[CCMR2]   { return mmio.RM32[CCMR2]{&p.CCMR2, IC4F} }

type CCER uint32

func CC1E_(p *Periph) mmio.RM32[CCER]  { return mmio.RM32[CCER]{&p.CCER, CC1E} }
func CC1P_(p *Periph) mmio.RM32[CCER]  { return mmio.RM32[CCER]{&p.CCER, CC1P} }
func CC1NE_(p *Periph) mmio.RM32[CCER] { return mmio.RM32[CCER]{&p.CCER, CC1NE} }
func CC1NP_(p *Periph) mmio.RM32[CCER] { return mmio.RM32[CCER]{&p.CCER, CC1NP} }
func CC2E_(p *Periph) mmio.RM32[CCER]  { return mmio.RM32[CCER]{&p.CCER, CC2E} }
func CC2P_(p *Periph) mmio.RM32[CCER]  { return mmio.RM32[CCER]{&p.CCER, CC2P} }
func CC2NE_(p *Periph) mmio.RM32[CCER] { return mmio.RM32[CCER]{&p.CCER, CC2NE} }
func CC2NP_(p *Periph) mmio.RM32[CCER] { return mmio.RM32[CCER]{&p.CCER, CC2NP} }
func CC3E_(p *Periph) mmio.RM32[CCER]  { return mmio.RM32[CCER]{&p.CCER, CC3E} }
func CC3P_(p *Periph) mmio.RM32[CCER]  { return mmio.RM32[CCER]{&p.CCER, CC3P} }
func CC3NE_(p *Periph) mmio.RM32[CCER] { return mmio.RM32[CCER]{&p.CCER, CC3NE} }
func CC3NP_(p *Periph) mmio.RM32[CCER] { return mmio.RM32[CCER]{&p.CCER, CC3NP} }
func CC4E_(p *Periph) mmio.RM32[CCER]  { return mmio.RM32[CCER]{&p.CCER, CC4E} }
func CC4P_(p *Periph) mmio.RM32[CCER]  { return mmio.RM32[CCER]{&p.CCER, CC4P} }

type BDTR uint32

func DTG_(p *Periph) mmio.RM32[BDTR]  { return mmio.RM32[BDTR]{&p.BDTR, DTG} }
func LOCK_(p *Periph) mmio.RM32[BDTR] { return mmio.RM32[BDTR]{&p.BDTR, LOCK} }
func OSSI_(p *Periph) mmio.RM32[BDTR] { return mmio.RM32[BDTR]{&p.BDTR, OSSI} }
func OSSR_(p *Periph) mmio.RM32[BDTR] { return mmio.RM32[BDTR]{&p.BDTR, OSSR} }
func BKE_(p *Periph) mmio.RM32[BDTR]  { return mmio.RM32[BDTR]{&p.BDTR, BKE} }
func BKP_(p *Periph) mmio.RM32[BDTR]  { return mmio.RM32[BDTR]{&p.BDTR, BKP} }
func AOE_(p *Periph) mmio.RM32[BDTR]  { return mmio.RM32[BDTR]{&p.BDTR, AOE} }
func MOE_(p *Periph) mmio.RM32[BDTR]  { return mmio.RM32[BDTR]{&p.BDTR, MOE} }

type DCR uint32

func DBA_(p *Periph) mmio.RM32[DCR] { return mmio.RM32[DCR]{&p.DCR, DBA} }
func DBL_(p *Periph) mmio.RM32[DCR] { return mmio.RM32[DCR]{&p.DCR, DBL} }
