// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"github.com/embeddedgo/stm32/hal/system"
	"github.com/embeddedgo/stm32/hal/system/timer/rtcst"
)

func init() {
	system.Setup80(0, 0)
	rtcst.Setup(rtcst.LSE, 1, 32768)
}
