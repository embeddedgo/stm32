// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32f303 || stm32f412
// +build stm32f303 stm32f412

package gpio

import (
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

func PA() *Port { return (*Port)(unsafe.Pointer(mmap.GPIOA_BASE + A*pstep)) }
func PB() *Port { return (*Port)(unsafe.Pointer(mmap.GPIOA_BASE + B*pstep)) }
func PC() *Port { return (*Port)(unsafe.Pointer(mmap.GPIOA_BASE + C*pstep)) }
func PD() *Port { return (*Port)(unsafe.Pointer(mmap.GPIOA_BASE + D*pstep)) }
func PE() *Port { return (*Port)(unsafe.Pointer(mmap.GPIOA_BASE + E*pstep)) }
func PF() *Port { return (*Port)(unsafe.Pointer(mmap.GPIOA_BASE + F*pstep)) }
func PG() *Port { return (*Port)(unsafe.Pointer(mmap.GPIOA_BASE + G*pstep)) }
func PH() *Port { return (*Port)(unsafe.Pointer(mmap.GPIOA_BASE + H*pstep)) }

const pnum = 8

// compile-time checks
const (
	c2 = byte((mmap.GPIOC_BASE - mmap.GPIOB_BASE - pstep) * 256)
	c3 = byte((mmap.GPIOD_BASE - mmap.GPIOC_BASE - pstep) * 256)
	c4 = byte((mmap.GPIOE_BASE - mmap.GPIOD_BASE - pstep) * 256)
	c5 = byte((mmap.GPIOF_BASE - mmap.GPIOE_BASE - pstep) * 256)
	c6 = byte((mmap.GPIOG_BASE - mmap.GPIOF_BASE - pstep) * 256)
	c7 = byte((mmap.GPIOH_BASE - mmap.GPIOG_BASE - pstep) * 256)
)
