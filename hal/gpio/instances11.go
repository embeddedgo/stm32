// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32f7x6 || stm32h7x3
// +build stm32f7x6 stm32h7x3

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
func PI() *Port { return (*Port)(unsafe.Pointer(mmap.GPIOA_BASE + I*pstep)) }
func PJ() *Port { return (*Port)(unsafe.Pointer(mmap.GPIOA_BASE + J*pstep)) }
func PK() *Port { return (*Port)(unsafe.Pointer(mmap.GPIOA_BASE + K*pstep)) }

const pnum = 11

// compile-time checks
const (
	c2  = byte((mmap.GPIOC_BASE - mmap.GPIOB_BASE - pstep) * 256)
	c3  = byte((mmap.GPIOD_BASE - mmap.GPIOC_BASE - pstep) * 256)
	c4  = byte((mmap.GPIOE_BASE - mmap.GPIOD_BASE - pstep) * 256)
	c5  = byte((mmap.GPIOF_BASE - mmap.GPIOE_BASE - pstep) * 256)
	c6  = byte((mmap.GPIOG_BASE - mmap.GPIOF_BASE - pstep) * 256)
	c7  = byte((mmap.GPIOH_BASE - mmap.GPIOG_BASE - pstep) * 256)
	c8  = byte((mmap.GPIOI_BASE - mmap.GPIOH_BASE - pstep) * 256)
	c9  = byte((mmap.GPIOJ_BASE - mmap.GPIOI_BASE - pstep) * 256)
	c10 = byte((mmap.GPIOK_BASE - mmap.GPIOJ_BASE - pstep) * 256)
)
