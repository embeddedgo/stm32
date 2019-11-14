// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package gpio

import (
	"github.com/embeddedgo/stm32/hal/internal"
	"github.com/embeddedgo/stm32/p/rcc"
)

const (
	veryLow  = -2 // 400 kHz (CL = 50 pF, VDD > 2.7 V)
	low      = -1 //   2 MHz (CL = 50 pF, VDD > 2.7 V)
	_        = 0  //  10 MHz (CL = 50 pF, VDD > 2.7 V)
	high     = 1  //  50 MHz (CL = 50 pF, VDD > 2.7 V)
	veryHigh = 1  // Not supported.
)

func enreg() *rcc.RAHBENR   { return &rcc.RCC().AHBENR }
func rstreg() *rcc.RAHBRSTR { return &rcc.RCC().AHBRSTR }

func lpenaclk(pnum uint) {
	internal.AtomicSetBits(&rcc.RCC().AHBLPENR.U32, uint32(rcc.GPIOALPEN<<pnum))
}

func lpdisclk(pnum uint) {
	internal.AtomicClearBits(&rcc.RCC().AHBLPENR.U32, uint32(rcc.GPIOALPEN<<pnum))
}

const (
	AF1_TIM2 = AF1

	AF2_TIM3 = AF2
	AF2_TIM4 = AF2
	AF2_TIM5 = AF2

	AF3_TIM9  = AF3
	AF3_TIM10 = AF3
	AF3_TIM11 = AF3

	AF4_I2C1 = AF4
	AF4_I2C2 = AF4

	AF5_SPI1 = AF5
	AF5_SPI2 = AF5

	AF6_SPI3 = AF6

	AF7_USART1 = AF7
	AF7_USART2 = AF7
	AF7_USART3 = AF7

	AF8_UART4 = AF8
	AF8_UART5 = AF8

	AF10_USB = AF10

	AF11_LCD = AF11

	AF12_FSMC = AF12

	AF14_RI = AF14
)
