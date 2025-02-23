// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package dmairq

import (
	"sync/atomic"
	"unsafe"

	"github.com/embeddedgo/stm32/hal/dma"
	"github.com/embeddedgo/stm32/hal/irq"
)

var handlers [16]unsafe.Pointer // func()

func setISR(c dma.Channel, isr func()) {
	h := *(*unsafe.Pointer)(unsafe.Pointer(&isr))
	cn := c.Num()
	dn := c.Controller().Num() - 1
	atomic.StorePointer(&handlers[dn*8+cn], h)
}

//go:nosplit
func isr(dn, cn int) {
	if h := atomic.LoadPointer(&handlers[dn*8+cn]); h != nil {
		(*(*func())(unsafe.Pointer(&h)))()
	}
}

func enableIRQs(prio int) {
	irq.DMA1_STR0.Enable(prio, 0)
	irq.DMA1_STR1.Enable(prio, 0)
	irq.DMA1_STR2.Enable(prio, 0)
	irq.DMA1_STR3.Enable(prio, 0)
	irq.DMA1_STR4.Enable(prio, 0)
	irq.DMA1_STR5.Enable(prio, 0)
	irq.DMA1_STR6.Enable(prio, 0)
	irq.DMA1_STR7.Enable(prio, 0)
	irq.DMA2_STR0.Enable(prio, 0)
	irq.DMA2_STR1.Enable(prio, 0)
	irq.DMA2_STR2.Enable(prio, 0)
	irq.DMA2_STR3.Enable(prio, 0)
	irq.DMA2_STR4.Enable(prio, 0)
	irq.DMA2_STR5.Enable(prio, 0)
	irq.DMA2_STR6.Enable(prio, 0)
	irq.DMA2_STR7.Enable(prio, 0)
}

// TODO: Think about avoid 16 separate handlers below. Currently they take ~1280
// bytes of Flash and pollute I-Cache. We can use VTOR and check interrupt
// number or check active and enabled DMA interrupts.

//go:interrupthandler
func _DMA1_STR0_Handler() { isr(0, 0) }

//go:interrupthandler
func _DMA1_STR1_Handler() { isr(0, 1) }

//go:interrupthandler
func _DMA1_STR2_Handler() { isr(0, 2) }

//go:interrupthandler
func _DMA1_STR3_Handler() { isr(0, 3) }

//go:interrupthandler
func _DMA1_STR4_Handler() { isr(0, 4) }

//go:interrupthandler
func _DMA1_STR5_Handler() { isr(0, 5) }

//go:interrupthandler
func _DMA1_STR6_Handler() { isr(0, 6) }

//go:interrupthandler
func _DMA1_STR7_Handler() { isr(0, 7) }

//go:interrupthandler
func _DMA2_STR0_Handler() { isr(1, 0) }

//go:interrupthandler
func _DMA2_STR1_Handler() { isr(1, 1) }

//go:interrupthandler
func _DMA2_STR2_Handler() { isr(1, 2) }

//go:interrupthandler
func _DMA2_STR3_Handler() { isr(1, 3) }

//go:interrupthandler
func _DMA2_STR4_Handler() { isr(1, 4) }

//go:interrupthandler
func _DMA2_STR5_Handler() { isr(1, 5) }

//go:interrupthandler
func _DMA2_STR6_Handler() { isr(1, 6) }

//go:interrupthandler
func _DMA2_STR7_Handler() { isr(1, 7) }

//go:linkname _DMA1_STR0_Handler IRQ11_Handler
//go:linkname _DMA1_STR1_Handler IRQ12_Handler
//go:linkname _DMA1_STR2_Handler IRQ13_Handler
//go:linkname _DMA1_STR3_Handler IRQ14_Handler
//go:linkname _DMA1_STR4_Handler IRQ15_Handler
//go:linkname _DMA1_STR5_Handler IRQ16_Handler
//go:linkname _DMA1_STR6_Handler IRQ17_Handler
//go:linkname _DMA1_STR7_Handler IRQ47_Handler
//go:linkname _DMA2_STR0_Handler IRQ56_Handler
//go:linkname _DMA2_STR1_Handler IRQ57_Handler
//go:linkname _DMA2_STR2_Handler IRQ58_Handler
//go:linkname _DMA2_STR3_Handler IRQ59_Handler
//go:linkname _DMA2_STR4_Handler IRQ60_Handler
//go:linkname _DMA2_STR5_Handler IRQ68_Handler
//go:linkname _DMA2_STR6_Handler IRQ69_Handler
//go:linkname _DMA2_STR7_Handler IRQ70_Handler
