// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"github.com/embeddedgo/stm32/hal/system"
	"github.com/embeddedgo/stm32/hal/system/timer/systick"
)

func init() {
	system.Setup168(8)
	systick.Setup(2e6)
}
