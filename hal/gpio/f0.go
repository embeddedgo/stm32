// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package gpio

import (
	"stm32/hal/raw/rcc"
)

const (
	veryLow  = -1 // Not supported.
	low      = -1 // 2 MHz (CL = 50 pF, VDD >= 2.4 V)
	_        = 0  // 10 MHz (CL = 50 pF, VDD >= 2.4 V)
	high     = 2  // 30 MHz (CL = 50 pF, VDD >= 2.7 V)
	veryHigh = 2  // Not supported.
)

func enreg() *rcc.RAHBENR   { return &rcc.RCC().AHBENR }
func rstreg() *rcc.RAHBRSTR { return &rcc.RCC().AHBRSTR }

func lpenaclk(pnum uint) {}
func lpdisclk(pnum uint) {}

const (
	AF0_EVENTOUT = AF0
	AF0_TIM15    = AF0
	AF0_SPI1     = AF0
	AF0_MCO      = AF0
	AF0_TIM17    = AF0
	AF0_SWDIO    = AF0
	AF0_SWCLK    = AF0
	AF0_TIM14    = AF0
	AF0_USART1   = AF0
	AF0_IR_OUT   = AF0
	AF0_SPI2     = AF0
	AF0_TIM3     = AF0
	AF0_USART4   = AF0
	AF0_TIM3_ETR = AF0

	AF1_USART1   = AF1
	AF1_USART2   = AF1
	AF1_TIM3     = AF1
	AF1_I2C1     = AF1
	AF1_I2C2     = AF1
	AF1_EVENTOUT = AF1
	AF1_TIM15    = AF1
	AF1_SPI2     = AF1
	AF1_USART3   = AF1

	AF2_TIM16  = AF2
	AF2_TIM17  = AF2
	AF2_TIM1   = AF2
	AF2_USART6 = AF2
	AF2_USART5 = AF2

	AF3_EVENTOUT = AF3
	AF3_I2C1     = AF3
	AF3_TIM15    = AF3

	AF4_USART4 = AF4
	AF4_TIM14  = AF4
	AF4_USART3 = AF4
	AF4_I2C1L  = AF4
	AF4_USART5 = AF4

	AF5_TIM15  = AF5
	AF5_USART6 = AF5
	AF5_TIM16  = AF5
	AF5_TIM17  = AF5
	AF5_MCO    = AF5
	AF5_SPI2   = AF5
	AF5_I2C2   = AF5

	AF6_EVENTOUT = AF6
)
