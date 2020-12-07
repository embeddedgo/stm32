// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32l4x6

// Clock setup for USB
//
// Goal is to provide 48 MHz for USB FS using PLL48M1CLK. USB can be clocked
// also from second PLL (PLLSAI1) which gives more flexibility for clocks setup
// but additional PLL means more power (see Current consumption).
//
// PLLVCO must satisfy the equation:
//
//  PLLVCO = 48 MHz * Q
//
// where Q can be:
//
//  Q = 2, 4, 6
//
// which means that PLLVCO can be: 96, 192, 288 MHz. We cannot use Q = 8 (PLLVCO
// = 384 MHz) because allowed PLLVCO range is 64 MHz to 344 MHz.
//
//  PLLIN = CLKSRC / M
//
// PLLIN must be 4 MHz to 16 MHz. Allowed M range is 1 to 8.
//
//  PLLVCO = PLLIN * N
//
// Allowed PLLVCO range is 64 MHz to 344 MHz, allowed N range is 8 to 86.
//
// Taking all this into account, N can be:
//
//  24, 16, 12, 8           for PLLVCO =  96 MHz (PLLIN: 4, 6, 8, 12 MHz),
//  48, 32, 24, 16, 12      for PLLVCO = 192 MHz (PLLIN: 4, 6, 8, 12, 16 MHz),
//  72, 48, 36, 32, 24, 18, for PLLVCO = 288 MHz (PLLIN: 4, 6, 8, 9, 12, 16 MHz)
//
// USB friendly values of CLKSRC:
//
//  - MSI: 48 Mhz (must be in LSE PLL mode: 32768 Hz * 1465 = 48005120 Hz).
//  - HSE: 4, 6, 8, 9, 12, 16, 18, 20, 21, 24, 27, 30, 32, 36, 40, 42, 48 MHz.
//
// Current consumption
//
//  MSI:  1   2   4   8   16   24   32   48  MHz
//  --------------------------------------------
//  Typ:  5   7  11  19   62   85  110  155  µA
//  Max:  6   9  15  25   80  110  130  190  µA
//
//  HSI16 (typ/max): 155/190 µA
//
//  HSE:   8   48  MHz
//  ------------------
//  Typ: 450  940  µA
//
//  PLLVCO:  64   96  192  344  MHz
//  --------------------------------
//  Typ:    150  200  300  520  µA
//  Max:    200  260  380  650  µA
//
package system

import (
	"runtime"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/flash"
	"github.com/embeddedgo/stm32/p/pwr"
	"github.com/embeddedgo/stm32/p/rcc"
)

// SetupPLL setups MCU for best performance (prefetch on, I/D cache on, minimum
// allowed Flash latency) using integrated PLL as system clock source.
//
// Clksrc configures clock source for PLL.
//
// Positive clksrc selects HSE as PLL clock source and informs about external
// clock signal frequency in MHz (alowed values: 4 to 48 MHz).
//
// Zero clksrc selects HSI16 as PLL clock surce.
//
// Negative clksrc selects MSI as PLL clock source and setups its frequency to
// (-clksrc) MHz (allowed values -4, -8, -16, -24, -32, -48).
//
// PLL input freq. is equal to clock source divided by M and must be in range
// 4 to 16 MHz.
//
// PLL VCO is equal to (input clock) / M * N and must be in range 64 to 344 MHz.
// Allowed M values: 1 to 8. Allowed N values: 8 to 86.
//
// P is VCO divider for PLLCAI2CLK. Allowed P values: 0 (disabled), 7, 17.
//
// Q is VCO divider for PLL48M1CLK (USB, RNG, SDMMC). Allowed Q values: 0
// (disabled), 2, 4, 6, 8.
//
// R is VCO divider for SYSCLK. Allowed R values: 2, 4, 6, 8.
//
// Voltage scaling Range 1 (high-performance) is configured if SYSCLK > 26 MHz
// or VCO > 128 MHz, otherwise Range 2 (low-power).
func SetupPLL(clksrc, M, N, P, Q, R int) {
	RCC := rcc.RCC()

	RCC.CIER.Store(0) // Disable clock interrupts.

	// Reset RCC clock configuration.
	RCC.CR.SetBits(rcc.MSION)
	for RCC.CR.LoadBits(rcc.MSIRDY) == 0 {
		runtime.Gosched()
	}
	RCC.CR.Store(6<<rcc.MSIRANGEn | rcc.MSIRGSEL | rcc.MSION)
	RCC.CFGR.Store(0) // MSI selected as system clock. APBCLK, AHBCLK = SYSCLK.

	// Calculate system clock.
	if M < 1 || M > 8 {
		panic("bad M")
	}
	if N < 8 || N > 86 {
		panic("bad N")
	}
	if P != 0 && P != 7 && P != 17 {
		panic("bad P")
	}
	if Q&1 != 0 || Q < 0 || Q > 8 {
		panic("bad Q")
	}
	if R&1 != 0 || R < 2 || R > 8 {
		panic("bad R")
	}

	var osc uint

	switch clksrc {
	case -4, -8, -16, -24, -32, -48:
		osc = uint(-clksrc)
	case 0:
		RCC.CR.SetBits(rcc.HSION)
		osc = 16
	default:
		if clksrc < 4 || clksrc > 48 {
			panic("bad clksrc")
		}
		// HSE needs milliseconds to stabilize, so enable it now.
		RCC.CR.SetBits(rcc.HSEON)
		osc = uint(clksrc)
	}

	pllin := osc * 1e6 / uint(M)
	if pllin < 4e6 || pllin > 16e6 {
		panic("bad PLLIN")
	}
	vco := pllin * uint(N)
	if vco < 64e6 || vco > 344e6 {
		panic("bad VCO")
	}
	sysclk := vco / uint(R)

	// Setup PWR and Flash.
	var (
		vos     pwr.CR1
		latency flash.ACR
	)
	if sysclk > 26e6 || vco > 128e6 {
		// Range 1: High-performance.
		vos = 1
		switch {
		case sysclk <= 16e6:
			latency = 0
		case sysclk <= 32e6:
			latency = 1
		case sysclk <= 48e6:
			latency = 2
		case sysclk <= 64e6:
			latency = 3
		default:
			latency = 4
		}
	} else {
		// Range 2: Low-power.
		vos = 2
		switch {
		case sysclk <= 6e6:
			latency = 0
		case sysclk <= 12e6:
			latency = 1
		case sysclk <= 18e6:
			latency = 2
		default:
			latency = 3
		}
	}
	RCC.APB1ENR1.SetBits(rcc.PWREN)
	RCC.APB1ENR1.Load()
	PWR := pwr.PWR()
	PWR.CR1.Store(vos << pwr.VOSn)
	flash.FLASH().ACR.Store(flash.DCEN | flash.ICEN | flash.PRFTEN | latency)

	// Setup clock dividers for AHB, APB1, APB2 bus.
	ahbclk := sysclk
	cfgr := rcc.HPRE_DIV1
	var apb1clk, apb2clk uint
	switch {
	case ahbclk <= 1*maxAPB1Clk:
		cfgr |= rcc.PPRE1_DIV1
		apb1clk = ahbclk / 1
	case ahbclk <= 2*maxAPB1Clk:
		cfgr |= rcc.PPRE1_DIV2
		apb1clk = ahbclk / 2
	case ahbclk <= 4*maxAPB1Clk:
		cfgr |= rcc.PPRE1_DIV4
		apb1clk = ahbclk / 4
	case ahbclk <= 8*maxAPB1Clk:
		cfgr |= rcc.PPRE1_DIV8
		apb1clk = ahbclk / 8
	default:
		cfgr |= rcc.PPRE1_DIV16
		apb1clk = ahbclk / 16
	}
	switch {
	case ahbclk <= 1*maxAPB2Clk:
		cfgr |= rcc.PPRE2_DIV1
		apb2clk = ahbclk / 1
	case ahbclk <= 2*maxAPB2Clk:
		cfgr |= rcc.PPRE2_DIV2
		apb2clk = ahbclk / 2
	case ahbclk <= 4*maxAPB2Clk:
		cfgr |= rcc.PPRE2_DIV4
		apb2clk = ahbclk / 4
	case ahbclk <= 8*maxAPB2Clk:
		cfgr |= rcc.PPRE2_DIV8
		apb2clk = ahbclk / 8
	default:
		cfgr |= rcc.PPRE2_DIV16
		apb2clk = ahbclk / 16
	}
	bus.Core.SetClock(int64(sysclk))
	bus.AHB1.SetClock(int64(ahbclk))
	bus.AHB2.SetClock(int64(ahbclk))
	bus.AHB3.SetClock(int64(ahbclk))
	bus.APB1.SetClock(int64(apb1clk))
	bus.APB2.SetClock(int64(apb2clk))

	// Setup PLL.
	for PWR.SR2.LoadBits(pwr.VOSF) != 0 {
		runtime.Gosched()
	}
	RCC.APB1ENR1.ClearBits(rcc.PWREN)
	var src rcc.PLLCFGR
	if clksrc == 0 {
		src = rcc.PLLSRC_HSI
		for RCC.CR.LoadBits(rcc.HSIRDY) == 0 {
			runtime.Gosched()
		}
	} else if clksrc > 0 {
		src = rcc.PLLSRC_HSE
		for RCC.CR.LoadBits(rcc.HSERDY) == 0 {
			runtime.Gosched()
		}
	} else {
		src = rcc.PLLSRC_MSI
		var msirange rcc.CR
		switch clksrc {
		case -4:
			// Current freq.
		case -48:
			msirange = 11
		default:
			msirange = rcc.CR(-clksrc/8 + 6)
		}
		if msirange != 0 {
			RCC.CR.StoreBits(rcc.MSIRANGE, msirange<<rcc.MSIRANGEn)
			for RCC.CR.LoadBits(rcc.MSIRDY) == 0 {
				runtime.Gosched()
			}
		}
	}
	mnpqr := rcc.PLLCFGR(M-1)<<rcc.PLLMn | rcc.PLLCFGR(N)<<rcc.PLLNn
	if P != 0 {
		mnpqr |= rcc.PLLPEN
		if P == 17 {
			mnpqr |= rcc.PLLP
		}
	}
	if Q != 0 {
		mnpqr |= rcc.PLLQEN | rcc.PLLCFGR(Q/2-1)
	}
	mnpqr |= rcc.PLLREN | rcc.PLLCFGR(R/2-1)
	RCC.PLLCFGR.Store(mnpqr | src)
	RCC.CR.SetBits(rcc.PLLON)
	for RCC.CR.LoadBits(rcc.PLLRDY) == 0 {
		runtime.Gosched()
	}

	// Set system clock source to PLL.
	RCC.CFGR.Store(cfgr | rcc.SW_PLL)
	for RCC.CFGR.LoadBits(rcc.SWS) != rcc.SWS_PLL {
		runtime.Gosched()
	}
	if osc >= 0 {
		RCC.CR.ClearBits(rcc.MSION)
	}
}

func Setup80(P, Q int) {
	SetupPLL(-48, 6, 20, P, Q, 2)
}

var msiRanges = [...]uint16{
	100, 200, 400, 800, 1e3, 2e3, 4e3, 8e3, 16e3, 24e3, 32e3, 48e3,
}

// SetupMSI setups MCU for best performance (prefetch on, I/D cache on, minimum
// allowed Flash latency) using integrated multi-speed oscilator as system
// clock source.
func SetupMSI(msikHz int) {
	RCC := rcc.RCC()

	// Reset RCC clock configuration.
	RCC.CR.SetBits(rcc.MSION)
	for RCC.CR.LoadBits(rcc.MSIRDY) == 0 {
		runtime.Gosched()
	}
	RCC.CR.Store(6<<rcc.MSIRANGEn | rcc.MSIRGSEL | rcc.MSION)
	RCC.CFGR.Store(0) // MSI selected as system clock. APBCLK, AHBCLK = SYSCLK.
	RCC.PLLCFGR.Store(0x1000)
	RCC.CIER.Store(0) // Disable clock interrupts.

	// Calculate system clock.
	var msirange rcc.CR
	for int(msirange) < len(msiRanges) {
		if int(msiRanges[msirange]) == msikHz {
			break
		}
		msirange++
	}
	if int(msirange) == len(msiRanges) {
		panic("bad msikHz")
	}
	sysclk := uint(msikHz) * 1e3

	// Setup PWR and Flash.
	var (
		vos     pwr.CR1
		latency flash.ACR
	)
	if sysclk > 26e6 {
		// Range 1: High-performance.
		vos = 1
		switch {
		case sysclk <= 16e6:
			latency = 0
		case sysclk <= 32e6:
			latency = 1
		case sysclk <= 48e6:
			latency = 2
		case sysclk <= 64e6:
			latency = 3
		default:
			latency = 4
		}
	} else {
		// Range 2: Low-power.
		vos = 2
		switch {
		case sysclk <= 6e6:
			latency = 0
		case sysclk <= 12e6:
			latency = 1
		case sysclk <= 18e6:
			latency = 2
		default:
			latency = 3
		}
	}

	RCC.APB1ENR1.SetBits(rcc.PWREN)
	PWR := pwr.PWR()
	PWR.CR1.Store(vos << pwr.VOSn)
	flash.FLASH().ACR.Store(flash.DCEN | flash.ICEN | flash.PRFTEN | latency)

	// Setup clock dividers for AHB, APB1, APB2 bus.
	ahbclk := sysclk
	cfgr := rcc.HPRE_DIV1
	var apb1clk, apb2clk uint
	switch {
	case ahbclk <= 1*maxAPB1Clk:
		cfgr |= rcc.PPRE1_DIV1
		apb1clk = ahbclk / 1
	case ahbclk <= 2*maxAPB1Clk:
		cfgr |= rcc.PPRE1_DIV2
		apb1clk = ahbclk / 2
	case ahbclk <= 4*maxAPB1Clk:
		cfgr |= rcc.PPRE1_DIV4
		apb1clk = ahbclk / 4
	case ahbclk <= 8*maxAPB1Clk:
		cfgr |= rcc.PPRE1_DIV8
		apb1clk = ahbclk / 8
	default:
		cfgr |= rcc.PPRE1_DIV16
		apb1clk = ahbclk / 16
	}
	switch {
	case ahbclk <= 1*maxAPB2Clk:
		cfgr |= rcc.PPRE2_DIV1
		apb2clk = ahbclk / 1
	case ahbclk <= 2*maxAPB2Clk:
		cfgr |= rcc.PPRE2_DIV2
		apb2clk = ahbclk / 2
	case ahbclk <= 4*maxAPB2Clk:
		cfgr |= rcc.PPRE2_DIV4
		apb2clk = ahbclk / 4
	case ahbclk <= 8*maxAPB2Clk:
		cfgr |= rcc.PPRE2_DIV8
		apb2clk = ahbclk / 8
	default:
		cfgr |= rcc.PPRE2_DIV16
		apb2clk = ahbclk / 16
	}

	bus.Core.SetClock(int64(sysclk))
	bus.AHB1.SetClock(int64(ahbclk))
	bus.AHB2.SetClock(int64(ahbclk))
	bus.AHB3.SetClock(int64(ahbclk))
	bus.APB1.SetClock(int64(apb1clk))
	bus.APB2.SetClock(int64(apb2clk))

	// Setup MSI freq.
	for PWR.SR2.LoadBits(pwr.VOSF) != 0 {
	}
	RCC.APB1ENR1.ClearBits(rcc.PWREN)
	if msirange != 6 {
		RCC.CR.StoreBits(rcc.MSIRANGE, msirange<<rcc.MSIRANGEn)
	}
}
