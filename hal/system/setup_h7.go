// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32h7x3

package system

import (
	"runtime"

	//"github.com/embeddedgo/stm32/p/bus"
	//"github.com/embeddedgo/stm32/p/flash"
	"github.com/embeddedgo/stm32/p/rcc"
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
func SetupPLL(clksrc, M, N, P int) {
	RCC := rcc.RCC()

	RCC.CIER.Store(0) // disable clock interrupts

	// Reset RCC clock configuration.
	RCC.HSION().Set()
	for RCC.HSIRDY().Load() == 0 {
		runtime.Gosched() // Wait for HSI...
	}
	RCC.CFGR.Store(0) // select HSI as system clock
	for RCC.SWS().Load() != rcc.SWS_HSI {
		runtime.Gosched() // wait for system clock setup...
	}
	RCC.CR.Store(rcc.HSION)

	// Calculate system clock.
	if M < 1 || M > 63 {
		panic("bad M")
	}
	if N < 4 || N > 512 {
		panic("bad N")
	}
	if P&1 != 0 && P != 1 || P < 1 || P > 128 {
		panic("bad P")
	}

	var (
		osc       uint
		pllckselr rcc.PLLCKSELR
	)

	switch clksrc {
	case -8, -16, -32, -64:
		pllckselr = rcc.PLLSRC_HSI
		osc = uint(-clksrc)
	case 0:
		pllckselr = rcc.PLLSRC_CSI
		RCC.CSION().Set()
		osc = 4
	default:
		if clksrc < 4 || clksrc > 48 {
			panic("bad clksrc")
		}
		pllckselr = rcc.PLLSRC_HSE
		// HSE needs milliseconds to stabilize, so enable it now.
		RCC.HSEON().Set()
		osc = uint(clksrc)
	}

	pllin := osc * 1e6 / uint(M)
	switch {
	case pllin >= 1 && pllin <= 2:

	case pllin > 2 && pllin <= 16:

	default:
		panic("bad PLLIN")
	}

	RCC.PLLCKSELR.Store(pllckselr)
}
