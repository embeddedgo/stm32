// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3
// +build stm32h7x3

package system

import (
	"runtime"

	//"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/flash"
	"github.com/embeddedgo/stm32/p/pwr"
	"github.com/embeddedgo/stm32/p/rcc"
	"github.com/embeddedgo/stm32/p/syscfg"
)

// SetupPLL setups MCU for the best performance (prefetch on, I/D cache on,
// minimum allowed Flash latency) using integrated PLL as system clock source.
//
// Clksrc configures clock source for PLL1.
//
// Positive clksrc selects HSE as PLL1 clock source and informs about external
// clock signal frequency in MHz (alowed values: 4 to 48 MHz).
//
// Zero clksrc selects CSI as PLL clock surce (frequency about 4 MHz).
//
// Negative clksrc selects HSI as PLL clock source and setups its frequency to
// (-clksrc) MHz (allowed values -8, -16, -32, -64).
//
// The PLL input clock is ref_ck = |clksrc| / M and must be 1..16 MHz. The VCO
// freq is Fvco = ref_ck * N and must be 150..420 MHz for ref_ck <= 2 MHz or
// 192..960 MHz for ref_ck >= 2 MHz.
func SetupPLL(clksrc, M, N, P, Q, R int) {
	RCC := rcc.RCC()

	RCC.CIER.Store(0) // disable clock interrupts

	// Reset RCC clock configuration.
	RCC.CR.SetBits(rcc.HSION)
	for RCC.CR.LoadBits(rcc.HSIRDY) == 0 {
		runtime.Gosched() // Wait for HSI...
	}
	RCC.CFGR.Store(0) // select HSI as system clock
	for RCC.CFGR.LoadBits(rcc.SWS) != rcc.SWS_HSI {
		runtime.Gosched() // wait for system clock setup...
	}
	RCC.CR.Store(rcc.HSION) // disables other clocks and all PLLs

	// Calculate and configure some clock dividers.

	if M < 1 || M > 63 {
		panic("bad M")
	}
	if N < 4 || N > 512 {
		panic("bad N")
	}
	if P&1 != 0 && P != 1 || P < 1 || P > 128 {
		panic("bad P")
	}
	pllcfgr := rcc.DIVP1EN
	switch {
	case Q == 0: // disabled
		Q = 128
	case Q >= 1 && Q <= 128:
		pllcfgr |= rcc.DIVQ1EN
	default:
		panic("bad Q")
	}
	switch {
	case R == 0: // disabled
		R = 128
	case R >= 1 && R <= 128:
		pllcfgr |= rcc.DIVR1EN
	default:
		panic("bad R")
	}
	divs := (R-1)<<rcc.DIVR1n | (Q-1)<<rcc.DIVQ1n | (P-1)<<rcc.DIVP1n | (N-1)<<rcc.DIVN1n
	RCC.PLL1DIVR.Store(rcc.PLL1DIVR(divs))

	pllckselr := rcc.PLLCKSELR(M) << rcc.DIVM1n
	var osc int
	switch clksrc {
	case -HSIClk / 8, -HSIClk / 4, -HSIClk / 2, -HSIClk:
		pllckselr |= rcc.PLLSRC_HSI
		osc = -clksrc
		cr := rcc.HSION
		switch clksrc {
		case HSIClk / 8:
			cr |= 3 << rcc.HSIDIV
		case HSIClk / 4:
			cr |= 2 << rcc.HSIDIV
		case HSIClk / 2:
			cr |= 1 << rcc.HSIDIV
		}
		RCC.CR.Store(cr)
	case 0:
		pllckselr |= rcc.PLLSRC_CSI
		RCC.CR.SetBits(rcc.CSION)
		osc = CSIClk
	default:
		if clksrc < 4 || clksrc > 48 {
			panic("bad clksrc")
		}
		pllckselr |= rcc.PLLSRC_HSE
		osc = clksrc
		RCC.CR.SetBits(rcc.HSEON)
	}
	RCC.PLLCKSELR.Store(pllckselr)

	pllin := osc * 1e6 / M
	switch {
	case pllin < 1e6:
		panic("PLLIN too low")
	case pllin <= 2e6:
		pllcfgr |= rcc.PLL1VCOSEL
	case pllin <= 4e6:
		pllcfgr |= 1 << rcc.PLL1RGEn
	case pllin <= 8e6:
		pllcfgr |= 2 << rcc.PLL1RGEn
	case pllin <= 16e6:
		pllcfgr |= 3 << rcc.PLL1RGEn
	default:
		panic("PLLIN too high")
	}
	RCC.PLLCFGR.Store(pllcfgr)

	sysclk := uint(pllin) * uint(N) / uint(P)

	// Calculate required Flash latency and connfigure PWR.

	var (
		vos       pwr.D3CR
		maxAHBClk uint
		maxAPBClk uint
	)
	switch {
	case sysclk <= maxSysClk3:
		vos = 3 // 0.95-1.05V
		maxAHBClk = maxAHBClk3
		maxAPBClk = maxAPBClk3
	case sysclk <= maxSysClk2:
		vos = 2 // 1.05-1.15V
		maxAHBClk = maxAHBClk2
		maxAPBClk = maxAPBClk2
	default:
		vos = 1 // 1.15-1.26V
		maxAHBClk = maxAHBClk1
		maxAPBClk = maxAPBClk1
	}
	PWR := pwr.PWR()
	PWR.D3CR.Store(vos << pwr.VOSn)
	for PWR.D3CR.LoadBits(pwr.VOSRDY) == 0 {
		runtime.Gosched() // wait for voltage regulator...
	}
	if sysclk > maxSysClk1 {
		maxAPBClk = maxAPBClk0
		RCC.APB4ENR.SetBits(rcc.SYSCFGEN)
		syscfg.SYSCFG().PWRCR.SetBits(syscfg.ODEN) // 1.26-1.40V (Vcore boost)
		RCC.APB4ENR.ClearBits(rcc.SYSCFGEN)
		for PWR.D3CR.LoadBits(pwr.VOSRDY) == 0 {
			runtime.Gosched()
		}
	}

	var (
		hclk, pclk uint
		d1cfgr     rcc.D1CFGR
		d2cfgr     rcc.D2CFGR
		d3cfgr     rcc.D3CFGR
	)
	switch {
	case sysclk <= 1*maxAHBClk:
		hclk = sysclk / 1
		d1cfgr = 0 << rcc.HPREn
	case sysclk <= 2*maxAHBClk:
		hclk = sysclk / 2
		d1cfgr = 8 << rcc.HPREn
	case sysclk <= 4*maxAHBClk:
		hclk = sysclk / 4
		d1cfgr = 9 << rcc.HPREn
	default:
		hclk = sysclk / 8
		d1cfgr = 10 << rcc.HPREn
	}
	switch {
	case hclk <= 1*maxAPBClk:
		pclk = hclk / 1
		d1cfgr |= 0 << rcc.D1PPREn
		d2cfgr = 0<<rcc.D2PPRE2n | 0<<rcc.D2PPRE1n
		d3cfgr = 0 << rcc.D3PPREn
	case hclk <= 2*maxAPBClk:
		pclk = hclk / 2
		d1cfgr |= 4 << rcc.D1PPREn
		d2cfgr = 4<<rcc.D2PPRE2n | 4<<rcc.D2PPRE1n
		d3cfgr = 4 << rcc.D3PPREn
	case hclk <= 4*maxAPBClk:
		pclk = hclk / 4
		d1cfgr |= 5 << rcc.D1PPREn
		d2cfgr = 5<<rcc.D2PPRE2n | 5<<rcc.D2PPRE1n
		d3cfgr = 5 << rcc.D3PPREn
	case hclk <= 8*maxAPBClk:
		pclk = hclk / 8
		d1cfgr |= 6 << rcc.D1PPREn
		d2cfgr = 6<<rcc.D2PPRE2n | 6<<rcc.D2PPRE1n
		d3cfgr = 6 << rcc.D3PPREn
	default:
		pclk = hclk / 16
		d1cfgr |= 7 << rcc.D1PPREn
		d2cfgr = 7<<rcc.D2PPRE2n | 7<<rcc.D2PPRE1n
		d3cfgr = 7 << rcc.D3PPREn
	}

	bus.Core.SetClock(int64(sysclk))
	for b := bus.AHB1; b <= bus.AHBLast; b++ {
		b.SetClock(int64(hclk))
	}
	for b := bus.APB1; b <= bus.APBLast; b++ {
		b.SetClock(int64(pclk))
	}

	var acr flash.ACR
	switch vos {
	case 1:
		switch {
		case hclk <= 70e6:
			acr = 0<<flash.WRHIGHFREQn | 0<<flash.LATENCYn
		case hclk <= 140e6:
			acr = 1<<flash.WRHIGHFREQn | 1<<flash.LATENCYn
		case hclk <= 185e6:
			acr = 1<<flash.WRHIGHFREQn | 2<<flash.LATENCYn
		case hclk <= 210e6:
			acr = 2<<flash.WRHIGHFREQn | 2<<flash.LATENCYn
		case hclk <= 225e6:
			acr = 2<<flash.WRHIGHFREQn | 3<<flash.LATENCYn
		default:
			acr = 2<<flash.WRHIGHFREQn | 4<<flash.LATENCYn
		}
	case 2:
		switch {
		case hclk <= 55e6:
			acr = 0<<flash.WRHIGHFREQn | 0<<flash.LATENCYn
		case hclk <= 110e6:
			acr = 1<<flash.WRHIGHFREQn | 1<<flash.LATENCYn
		case hclk <= 165e6:
			acr = 1<<flash.WRHIGHFREQn | 2<<flash.LATENCYn
		case hclk <= 225e6:
			acr = 2<<flash.WRHIGHFREQn | 3<<flash.LATENCYn
		default:
			acr = 2<<flash.WRHIGHFREQn | 4<<flash.LATENCYn
		}
	default: // 3
		switch {
		case hclk <= 45e6:
			acr = 0<<flash.WRHIGHFREQn | 0<<flash.LATENCYn
		case hclk <= 90e6:
			acr = 1<<flash.WRHIGHFREQn | 1<<flash.LATENCYn
		case hclk <= 135e6:
			acr = 1<<flash.WRHIGHFREQn | 2<<flash.LATENCYn
		case hclk <= 180e6:
			acr = 2<<flash.WRHIGHFREQn | 3<<flash.LATENCYn
		default:
			acr = 2<<flash.WRHIGHFREQn | 4<<flash.LATENCYn
		}
	}

	// Enable PLL1 and make it a system clock source.

	RCC.CR.SetBits(rcc.PLL1ON)
	FLASH := flash.FLASH1() // PLL need some time to stabilize so
	FLASH.ACR.Store(acr)    // reconfigure Flash in the the meantime.
	for FLASH.ACR.LoadBits(flash.WRHIGHFREQ|flash.LATENCY) != acr {
		runtime.Gosched()
	}
	for RCC.CR.LoadBits(rcc.PLL1RDY) == 0 {
		runtime.Gosched()
	}

	RCC.D1CFGR.Store(d1cfgr)
	RCC.D2CFGR.Store(d2cfgr)
	RCC.D3CFGR.Store(d3cfgr)
	RCC.CFGR.StoreBits(rcc.SW, rcc.SW_PLL1)
	for RCC.CFGR.LoadBits(rcc.SWS) != rcc.SWS_PLL1 {
		runtime.Gosched()
	}
	if osc >= 0 {
		RCC.CR.ClearBits(rcc.HSION)
	}
}
