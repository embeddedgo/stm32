// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32f303

package gpio

import (
	"embedded/mmio"

	"github.com/embeddedgo/stm32/p/rcc"
)

const (
	veryLow  = -1 // Not supported.
	low      = -1 // 2 MHz ???
	_        = 0  // 18 MHz ???
	high     = 2  // 36 MHz (CL = 47 pF, VDD = 3.3 V)
	veryHigh = 2  // Not supported.
)

func enreg() *mmio.R32[rcc.AHBENR]   { return &rcc.RCC().AHBENR }
func rstreg() *mmio.R32[rcc.AHBRSTR] { return &rcc.RCC().AHBRSTR }

func lpenaclk(pnum uint) {}
func lpdisclk(pnum uint) {}

const (
	AF0_RTC_REFIN     = AF0
	AF0_MCO           = AF0
	AF0_SWDIO_FTNS    = AF0
	AF0_SWCLK_JTCK    = AF0
	AF0_JTDI          = AF0
	AF0_JTDO_TRACESWO = AF0
	AF0_JRST          = AF0
	AF0_TRACECK       = AF0
	AF0_TRACED0       = AF0
	AF0_TRACED1       = AF0
	AF0_TRACED2       = AF0
	AF0_TRACED3       = AF0

	AF1_TIM2     = AF1
	AF1_TIM15    = AF1
	AF1_TIM16    = AF1
	AF1_TIM17    = AF1
	AF1_EVENTOUT = AF1

	AF2_I2C3  = AF2
	AF2_TIM1  = AF2
	AF2_TIM2  = AF2
	AF2_TIM3  = AF2
	AF2_TIM4  = AF2
	AF2_TIM8  = AF2
	AF2_TIM20 = AF2
	AF2_TIM15 = AF2
	AF2_COMP1 = AF2

	AF3_I2C3  = AF3
	AF3_TIM8  = AF3
	AF3_TIM20 = AF3
	AF3_TIM15 = AF3
	AF3_COMP7 = AF3
	AF3_TSC   = AF3

	AF4_I2C1  = AF4
	AF4_I2C2  = AF4
	AF4_TIM1  = AF4
	AF4_TIM8  = AF4
	AF4_TIM16 = AF4
	AF4_TIM17 = AF4

	AF5_SPI1  = AF5
	AF5_SPI2  = AF5
	AF5_I2S2  = AF5
	AF5_SPI3  = AF5
	AF5_I2S3  = AF5
	AF5_SPI4  = AF5
	AF5_UART4 = AF5
	AF5_UART5 = AF5
	AF5_TIM8  = AF5
	AF5_IR    = AF5

	AF6_SPI2  = AF6
	AF6_I2S2  = AF6
	AF6_SPI3  = AF6
	AF6_I2S3  = AF6
	AF6_TIM1  = AF6
	AF6_TIM8  = AF6
	AF6_TIM20 = AF6
	AF6_IR    = AF6

	AF7_USART1 = AF7
	AF7_USART2 = AF7
	AF7_USART3 = AF7
	AF7_CAN    = AF7
	AF7_COMP3  = AF7
	AF7_COMP5  = AF7
	AF7_COMP6  = AF7

	AF8_I2C3  = AF8
	AF8_COMP1 = AF8
	AF8_COMP2 = AF8
	AF8_COMP3 = AF8
	AF8_COMP4 = AF8
	AF8_COMP5 = AF8
	AF8_COMP6 = AF8

	AF9_CAN   = AF9
	AF9_TIM1  = AF9
	AF9_TIM8  = AF9
	AF9_TIM15 = AF9

	AF10_TIM2  = AF10
	AF10_TIM3  = AF10
	AF10_TIM4  = AF10
	AF10_TIM8  = AF10
	AF10_TIM17 = AF10

	AF11_TIM1 = AF11
	AF11_TIM8 = AF11

	AF12_FSMC = AF12
	AF12_TIM1 = AF12
)
