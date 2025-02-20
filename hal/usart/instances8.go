// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32f407

package usart

import (
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

func USART1() *Periph { return (*Periph)(unsafe.Pointer(mmap.USART1_BASE)) }
func USART2() *Periph { return (*Periph)(unsafe.Pointer(mmap.USART2_BASE)) }
func USART3() *Periph { return (*Periph)(unsafe.Pointer(mmap.USART3_BASE)) }
func UART4() *Periph  { return (*Periph)(unsafe.Pointer(mmap.UART4_BASE)) }
func UART5() *Periph  { return (*Periph)(unsafe.Pointer(mmap.UART5_BASE)) }
func USART6() *Periph { return (*Periph)(unsafe.Pointer(mmap.USART6_BASE)) }
func UART7() *Periph  { return (*Periph)(unsafe.Pointer(mmap.UART7_BASE)) }
func UART8() *Periph  { return (*Periph)(unsafe.Pointer(mmap.UART8_BASE)) }
