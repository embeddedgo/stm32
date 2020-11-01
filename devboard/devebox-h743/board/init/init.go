// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package init

import (
	"github.com/embeddedgo/stm32/hal/system"
	"github.com/embeddedgo/stm32/hal/system/timer/systick"
)

func init() {
	// 480 MHz not supported, requires disabling Vcore overdrive before any WFE instruction
	system.SetupPLL(25, 5, 80, 1, 2, 2) // 400 MHz
	systick.Setup(2e6)
	//rtcst.Setup(rtcst.LSE, 1, 32768)
}
