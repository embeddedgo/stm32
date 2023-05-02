// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows the basic usage of the available LED and button.
package main

import (
	"embedded/rtos"
	"time"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/system"
	"github.com/embeddedgo/stm32/hal/system/timer/rtcst"
	"github.com/embeddedgo/stm32/p/rcc"
	"github.com/embeddedgo/stm32/p/syscfg"
)

func main() {
	system.Setup80(0, 0)
	rtcst.Setup(rtcst.LSE, 1, 32768)

	pa := gpio.PA()
	pa.EnableClock(true)
	led := pa.Pin(5)

	led.Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.Low})

	for i := 0; i < 10; i++ {
		led.Set()
		time.Sleep(time.Second / 8)
		led.Clear()
		time.Sleep(time.Second / 8)
	}
	RCC := rcc.RCC()
	RCC.APB2ENR.SetBits(rcc.SYSCFGEN)
	syscfg.SYSCFG().MEMRMP.StoreBits(syscfg.MEM_MODE, 1)
	RCC.APB2ENR.ClearBits(rcc.SYSCFGEN)
	rtos.Reset(rtos.UnsafeReboot, unsafe.Pointer(uintptr(0x1FFF0000)))
}
