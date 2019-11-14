// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package systick allows to setup ARMv7-M SysTick timer as ticking system timer.
package systick

import (
	"embedded/arch/cortexm/systim"
	"embedded/rtos"
	"runtime"

	"github.com/embeddedgo/stm32/p/bus"
)

// Setup setups the SysTick as system timer. This is a ticking timer. Use tickless
// timer (eg. RTC based) if available. SysTick runs the thread scheduler every
// periodns nanoseconds.
func Setup(periodns int64) {
	runtime.LockOSThread()
	pl, _ := rtos.SetPrivLevel(0)
	systim.Setup(periodns, bus.AHB1.Clock()/8, true)
	rtos.SetPrivLevel(pl)
	runtime.UnlockOSThread()
	rtos.SetSystemTimer(systim.Nanotime, nil)
}
