// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !stm32h7x3

package systick

import (
	"embedded/arch/cortexm/systim"

	"github.com/embeddedgo/stm32/p/bus"
)

func systimSetup(periodns int64) {
	systim.Setup(periodns, bus.AHB1.Clock()/8, true)
}
