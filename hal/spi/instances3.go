// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f215 stm32l4x6

package spi

import (
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

func SPI1() *Periph { return (*Periph)(unsafe.Pointer(mmap.SPI1_BASE)) }
func SPI2() *Periph { return (*Periph)(unsafe.Pointer(mmap.SPI2_BASE)) }
func SPI3() *Periph { return (*Periph)(unsafe.Pointer(mmap.SPI3_BASE)) }
