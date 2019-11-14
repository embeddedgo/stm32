// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32l4x6

package gpio

import (
	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
)

const (
	veryLow  = -1 // Not supported.
	low      = -1 // v MHz (CL = 50 pF, VDD > 2.7 V)
	_        = 0  // x MHz (CL = 50 pF, VDD > 2.7 V)
	high     = 1  // y MHz (CL = 40 pF, VDD > 2.7 V)
	veryHigh = 2  // z MHz (CL = 30 pF, VDD > 2.7 V)
)

func enreg() *rcc.RAHB2ENR   { return &rcc.RCC().AHB2ENR }
func rstreg() *rcc.RAHB2RSTR { return &rcc.RCC().AHB2RSTR }

func lpenaclk(pnum uint) {
	internal.AtomicSetBits(&rcc.RCC().AHB2SMENR.U32, uint32(rcc.GPIOASMEN<<pnum))
}

func lpdisclk(pnum uint) {
	internal.AtomicClearBits(&rcc.RCC().AHB2SMENR.U32, uint32(rcc.GPIOASMEN<<pnum))
}

const (
	AF0_MCO = AF0

	AF1_TIM2   = AF1
	AF1_TIM1   = AF1
	AF1_IR_OUT = AF1
	AF1_LPTIM1 = AF1
	AF1_TIM8   = AF1

	AF2_TIM5 = AF2
	AF2_TIM2 = AF2
	AF2_TIM3 = AF2
	AF2_TIM1 = AF2
	AF2_TIM4 = AF2

	AF3_TIM8 = AF3
	AF3_TIM1 = AF3

	AF4_I2C3 = AF4
	AF4_I2C1 = AF4
	AF4_I2C2 = AF4

	AF5_SPI1 = AF5
	AF5_SPI2 = AF5

	AF6_SPI3  = AF6
	AF6_DFSDM = AF6

	AF7_USART2 = AF7
	AF7_USART3 = AF7
	AF7_USART1 = AF7

	AF8_UART4   = AF8
	AF8_UART5   = AF8
	AF8_LPUART1 = AF8

	AF9_CAN1 = AF9
	AF9_TSC  = AF9

	AF10_QUADSPI = AF10
	AF10_OTG_FS  = AF10

	AF11_LCD = AF11

	AF12_TIM1   = AF12
	AF12_COMP1  = AF12
	AF12_TIM8   = AF12
	AF12_FMC    = AF12
	AF12_SDMMC1 = AF12
	AF12_SWPMI1 = AF12

	AF13_SAI1 = AF13
	AF13_SAI2 = AF13
	AF13_TIM8 = AF13

	AF14_TIM2   = AF14
	AF14_TIM15  = AF14
	AF14_LPTIM2 = AF14
	AF14_TIM16  = AF14
	AF14_TIM17  = AF14
	AF14_TIM8   = AF14
)
