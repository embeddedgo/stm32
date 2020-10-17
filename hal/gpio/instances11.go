// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f7x6 stm32h7x3

package gpio

import "github.com/embeddedgo/stm32/p/mmap"

func A() *Port { return P(0) }
func B() *Port { return P(1) }
func C() *Port { return P(2) }
func D() *Port { return P(3) }
func E() *Port { return P(4) }
func F() *Port { return P(5) }
func G() *Port { return P(6) }
func H() *Port { return P(7) }
func I() *Port { return P(8) }
func J() *Port { return P(9) }
func K() *Port { return P(10) }

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
