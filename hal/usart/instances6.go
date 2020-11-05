// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//+build stm32l4x6

package usart

import (
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

func USART1() *Periph  { return (*Periph)(unsafe.Pointer(mmap.USART1_BASE)) }
func USART2() *Periph  { return (*Periph)(unsafe.Pointer(mmap.USART2_BASE)) }
func USART3() *Periph  { return (*Periph)(unsafe.Pointer(mmap.USART3_BASE)) }
func UART4() *Periph   { return (*Periph)(unsafe.Pointer(mmap.UART4_BASE)) }
func UART5() *Periph   { return (*Periph)(unsafe.Pointer(mmap.UART5_BASE)) }
func LPUART1() *Periph { return (*Periph)(unsafe.Pointer(mmap.LPUART1_BASE)) }
