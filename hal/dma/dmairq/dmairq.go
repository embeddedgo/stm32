// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dmairq

import (
	"embedded/rtos"

	"github.com/embeddedgo/stm32/hal/dma"
)

func SetISR(c dma.Channel, isr func()) { setISR(c, isr) }
func SetPrio(prio int)                 { enableIRQs(prio) }

func init() { enableIRQs(rtos.IntPrioLow) }
