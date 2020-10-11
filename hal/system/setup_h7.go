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

}
