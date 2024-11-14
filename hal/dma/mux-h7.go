// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package dma

import (
	"github.com/embeddedgo/stm32/p/dmamux/dmamux1"
	"github.com/embeddedgo/stm32/p/mmap"
)

// A Mux represents a configuration of DMAMUX for a DMA channel.
type Mux uint32

const (
	REQ_DISA     Mux = 0
	REQ_GEN0     Mux = 1
	REQ_GEN1     Mux = 2
	REQ_GEN2     Mux = 3
	REQ_GEN3     Mux = 4
	REQ_GEN4     Mux = 5
	REQ_GEN5     Mux = 6
	REQ_GEN6     Mux = 7
	REQ_GEN7     Mux = 8
	ADC1         Mux = 9
	ADC2         Mux = 10
	TIM1_CH1     Mux = 11
	TIM1_CH2     Mux = 12
	TIM1_CH3     Mux = 13
	TIM1_CH4     Mux = 14
	TIM1_UP      Mux = 15
	TIM1_TRIG    Mux = 16
	TIM1_COM     Mux = 17
	TIM2_CH1     Mux = 18
	TIM2_CH2     Mux = 19
	TIM2_CH3     Mux = 20
	TIM2_CH4     Mux = 21
	TIM2_UP      Mux = 22
	TIM3_CH1     Mux = 23
	TIM3_CH2     Mux = 24
	TIM3_CH3     Mux = 25
	TIM3_CH4     Mux = 26
	TIM3_UP      Mux = 27
	TIM3_TRIG    Mux = 28
	TIM4_CH1     Mux = 29
	TIM4_CH2     Mux = 30
	TIM4_CH3     Mux = 31
	TIM4_UP      Mux = 32
	I2C1_RX      Mux = 33
	I2C1_TX      Mux = 34
	I2C2_RX      Mux = 35
	I2C2_TX      Mux = 36
	SPI1_RX      Mux = 37
	SPI1_TX      Mux = 38
	SPI2_RX      Mux = 39
	SPI2_TX      Mux = 40
	USART1_RX    Mux = 41
	USART1_TX    Mux = 42
	USART2_RX    Mux = 43
	USART2_TX    Mux = 44
	USART3_RX    Mux = 45
	USART3_TX    Mux = 46
	TIM8_CH1     Mux = 47
	TIM8_CH2     Mux = 48
	TIM8_CH3     Mux = 49
	TIM8_CH4     Mux = 50
	TIM8_UP      Mux = 51
	TIM8_TRIG    Mux = 52
	TIM8_COM     Mux = 53
	TIM5_CH1     Mux = 55
	TIM5_CH2     Mux = 56
	TIM5_CH3     Mux = 57
	TIM5_CH4     Mux = 58
	TIM5_UP      Mux = 59
	TIM5_TRIG    Mux = 60
	SPI3_RX      Mux = 61
	SPI3_TX      Mux = 62
	UART4_RX     Mux = 63
	UART4_TX     Mux = 64
	UART5_RX     Mux = 65
	UART5_TX     Mux = 66
	DAC_CH1      Mux = 67
	DAC_CH2      Mux = 68
	TIM6_UP      Mux = 69
	TIM7_UP      Mux = 70
	USART6_RX    Mux = 71
	USART6_TX    Mux = 72
	I2C3_RX      Mux = 73
	I2C3_TX      Mux = 74
	DCMI         Mux = 75
	CRYP_IN      Mux = 76
	CRYP_OUT     Mux = 77
	HASH_IN      Mux = 78
	UART7_RX     Mux = 79
	UART7_TX     Mux = 80
	UART8_RX     Mux = 81
	UART8_TX     Mux = 82
	SPI4_RX      Mux = 83
	SPI4_TX      Mux = 84
	SPI5_RX      Mux = 85
	SPI5_TX      Mux = 86
	SAI1A        Mux = 87
	SAI1B        Mux = 88
	SAI2A        Mux = 89
	SAI2B        Mux = 90
	SWPMI_RX     Mux = 91
	SWPMI_TX     Mux = 92
	SPDIFRX_DAT  Mux = 93
	SPDIFRX_CTRL Mux = 94
	HR_REQ1      Mux = 95
	HR_REQ2      Mux = 96
	HR_REQ3      Mux = 97
	HR_REQ4      Mux = 98
	HR_REQ5      Mux = 99
	HR_REQ6      Mux = 100
	DFSDM10      Mux = 101
	DFSDM11      Mux = 102
	DFSDM12      Mux = 103
	DFSDM13      Mux = 104
	TIM15_CH1    Mux = 105
	TIM15_UP     Mux = 106
	TIM15_TRIG   Mux = 107
	TIM15_COM    Mux = 108
	TIM16_CH1    Mux = 109
	TIM16_UP     Mux = 110
	TIM17_CH1    Mux = 111
	TIM17_UP     Mux = 112
	SAI3_A       Mux = 113
	SAI3_B       Mux = 114
	ADC3         Mux = 115

	SOIE         Mux = 1 << 8   //+ Synchronization overrun interrupt enable
	EGE          Mux = 1 << 9   //+ Event generation enable
	SE           Mux = 1 << 16  //+ Synchronization enable
	SPOL         Mux = 3 << 17  //+ Synchronization polarity Defines the edge polarity of the selected synchronization input:
	SPOL_NONE    Mux = 0 << 17  //  No event, i.e. no synchronization nor detection.
	SPOL_RISING  Mux = 1 << 17  //  Rising edge
	SPOL_FALLING Mux = 2 << 17  //  Falling edge
	SPOL_BOTH    Mux = 3 << 17  //  Rising and falling edges
	NBREQ        Mux = 31 << 19 //+ Number of DMA requests minus 1 to forward after a synchronization event, and/or the number of DMA requests before an output event is generated. This field shall only be written when both SE and EGE bits are low.
	SYNC_ID      Mux = 7 << 24  //+ Synchronization input
)

func (c Channel) SetMux(mux Mux) {
	sn := snum(c)
	dn := (c.h&^0x3ff - mmap.DMA1_BASE) / 0x400
	dmamux1.DMAMUX1().CCR[dn*8+sn].Store(dmamux1.CCR(mux))
}

/*

See imxrt.

var chanMask uint32 = 0xffff

func (d *Controller) AllocChannel() Channel {

}

*/