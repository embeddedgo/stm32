// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"github.com/embeddedgo/stm32/hal/system"
	"github.com/embeddedgo/stm32/hal/system/timer/systick"
)

func init() {
	// 480 MHz CPU clock not supported, requires disabling Vcore overdrive
	// before any WFE instruction

	// PLL1 configuration (TODO: PLL2 and PLL3)
	const (
		osc = 25 // MHz
		m   = 5  // ref_ck = osc / 5 = 5 MHz
		n   = 80 // Fvco = ref_ck * n = 400 MHz
		p   = 1  // p_ck = Fvco / p = 400 MHz (CPU)
		q   = 2  // q_ck = Fvco / q = 200 MHz (SPI(I2S)1,2,3, SAI)
		r   = 2  // r_ck = Fvco / r = 200 MHz (trace port)
	)
	system.SetupPLL(osc, m, n, p, q, r)
	systick.Setup(2e6)
	//rtcst.Setup(rtcst.LSE, 1, 32768)
}
