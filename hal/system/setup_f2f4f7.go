// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32f407 stm32f412

// Goal is to provide 48 MHz for USB FS (SDIO, RNG). So PLLVCO must satisfy the
// equation:
//
//  PLLVCO = 48 MHz * Q
//
// where Q = 2..15, which means that PLLVCO can be: 96, 144, ... , 720 MHz.
//
// But allowed PLLVCO is between 100 and 432 MHz so useful Q values are:
//
//  Q = 3..9
//
// which means PLLVCO can be: 144, 192, 240, 288, 336, 384, 432 MHz.
//
// PLL multipler N range is 50..432. There is recommendation to use 2 MHz input
// clock to PLL to limit its jitter. Taking this into account PLLVCO can be:
//
//  PLLVCO = N * 2 MHz
//
// PLLVCO should be between 100 and 432 MHz so useful N values are:
//
//  N = 50..216
//
// There is much smaller choice of N values that satisfy USB FS requirements:
//
//  N = 72, 96, 120, 144, 168, 192, 216.
//
// System Clock is derived from PLLVCO as follows:
//
//  SYSCLK = PLLVCO / P
//
// where P = 2, 4, 6, 8.
//
// If 2 MHz PLL input clock is used, System Clock can be calculated as follows:
//
//  SYSCLK = 2 MHz * N / P
//
package system

import (
	"runtime"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/flash"
	"github.com/embeddedgo/stm32/p/rcc"
)

// SetupPLL setups MCU for the best performance (prefetch on, I/D cache on,
// minimum allowed Flash latency) using integrated PLL as system clock source.
//
// osc is freqency of external resonator in MHz. Allowed values are multiples
// of 2, from 4 to 26. Use 0 to select internal HSI oscilator as  system clock
// source.
//
// N is PLL multipler. Allowed values are from 50 to 216 but if USB FS will be
// used, N can be only:
//
//  N(USB) = 72, 96, 120, 144, 168, 192, 216
//
// P is system clock divider. Allowed values: 2, 4, 6, 8.
//
// Both N and P determine the system clock frequency according to the
// formula:
//
//  SysClk = 2e6 * N / P [Hz]
//
func SetupPLL(osc, N, P int) {
	RCC := rcc.RCC()

	RCC.CIR.Store(0) // Disable clock interrupts.

	// Reset RCC clock configuration.
	RCC.HSION().Set()
	for RCC.HSIRDY().Load() == 0 {
		runtime.Gosched() // Wait for HSI...
	}
	RCC.CFGR.Store(0)
	for RCC.SWS().Load() != rcc.SWS_HSI {
		runtime.Gosched() // Wait for system clock setup...
	}
	RCC.CR.Store(rcc.HSION)

	// Calculate system clock.
	if osc != 0 && (osc < 4 || osc > 26 || osc&1 != 0) {
		panic("bad HSE osc freq")
	}
	if N < 72 || N > 216 {
		panic("bad N")
	}
	if P&1 != 0 || P < 2 || P > 8 {
		panic("bad P")
	}
	// HSE needs milliseconds to stabilize, so enable it now.
	if osc != 0 {
		RCC.HSEON().Set()
	}

	sysclk := 2e6 * uint(N) / uint(P)

	// Setup linear voltage regulator scaling. Default scale 1 works well up
	// to 168 MHz. Scale 2 can be enabled in case of sysclk <= 144 MHz but
	// the numerical values of constants are not compatible between different
	// STM32F4 series.
	//RCC.PWREN().Set()
	//pwr.PWR.VOS().StoreBits(pwr.VOS, vos)
	//for rcc.PWR.VOSRDY().Load() == 0 {
	//	runtime.Gosched() // wait for voltage regulator...
	//}
	//RCC.PWREN().Clear()

	// Setup Flash.
	FLASH := flash.FLASH()
	latency := flash.ACR((sysclk-1)/30e6) << flash.LATENCYn // wait states
	const dcen = 1 << 10                                    // F7 has no DCEN.
	const icen = 1 << 9                                     // ICEN in F4, ARTEN in F7.
	FLASH.ACR.Store(dcen | icen | flash.PRFTEN | latency)

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
	for b := bus.AHB1; b <= bus.AHBLast; b++ {
		b.SetClock(int64(ahbclk))
	}
	bus.APB1.SetClock(int64(apb1clk))
	bus.APB2.SetClock(int64(apb2clk))

	// Setup PLL.
	var src rcc.PLLCFGR
	mnpq := rcc.PLLCFGR(N)<<rcc.PLLNn | // PLL multiler.
		rcc.PLLCFGR(P/2-1)<<rcc.PLLPn | // SysClk divider.
		rcc.PLLCFGR(2*N+47)/48<<rcc.PLLQn // 48MHz divider for USB, SDIO, RNG.
	if osc != 0 {
		src = rcc.PLLSRC_HSE
		mnpq |= rcc.PLLCFGR(osc/2) << rcc.PLLMn
		for RCC.HSERDY().Load() == 0 {
			runtime.Gosched()
		}
	} else {
		src = rcc.PLLSRC_HSI
		mnpq |= HSIClk / 2 << rcc.PLLMn
	}
	RCC.PLLCFGR.Store(mnpq | src)
	RCC.PLLON().Set()
	for RCC.PLLRDY().Load() == 0 {
			runtime.Gosched()
	}

	// Change system clock source to PLL.
	RCC.CFGR.Store(cfgr | rcc.SW_PLL)
	for RCC.SWS().Load() != rcc.SWS_PLL {
			runtime.Gosched()
	}
	if osc != 0 {
		RCC.HSION().Clear()
	}
}

// Setup96 wraps SetupPLL to setup MCU to work with 96 MHz clock. See SetupPLL
// for more information.
func Setup96(osc int) {
	SetupPLL(osc, 96, 2)
}

// Setup120 wraps SetupPLL to setup MCU to work with 120 MHz clock. See SetupPLL
// for more information.
func Setup120(osc int) {
	SetupPLL(osc, 120, 2)
}

// Setup168 wraps SetupPLL to setup MCU to work with 168 MHz clock. See SetupPLL
// for more information.
func Setup168(osc int) {
	SetupPLL(osc, 168, 2)
}

// Setup192 wraps SetupPLL to setup MCU to work with 192 MHz clock. See SetupPLL
// for more information.
func Setup192(osc int) {
	SetupPLL(osc, 192, 2)
}
