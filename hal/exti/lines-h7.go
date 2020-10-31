// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32h7x3

package exti

const (
	RTCALR  Lines = 1 << 17 // Real Time Clock Alarm event.
	RTCTTS  Lines = 1 << 18 // RTC Tamper and TimeStamp events.
	RTCWKUP Lines = 1 << 19 // RTC Wakeup event.
	COMP1   Lines = 1 << 20 // Comparator 1 wakeup event.
	COMP2   Lines = 1 << 21 // Comparator 2 wakeup event.
	I2C1    Lines = 1 << 22 // I2C1 wakeup event.
	I2C2    Lines = 1 << 23 // I2C2 wakeup event.
	I2C3    Lines = 1 << 24 // I2C3 wakeup event.
	I2C4    Lines = 1 << 25 // I2C3 wakeup event.
	USART1  Lines = 1 << 26 // USART1 wakeup event.
	USART2  Lines = 1 << 27 // USART2 wakeup event.
	USART3  Lines = 1 << 28 // USART3 wakeup event.
	USART6  Lines = 1 << 29 // USART6 wakeup event.
	USART4  Lines = 1 << 30 // USART4 wakeup event.
	USART5  Lines = 1 << 31 // USART5 wakeup event.

	USART7     Lines = 1 << 32 // USART7 wakeup event.
	USART8     Lines = 1 << 33 // USART8 wakeup event.
	LPUART1RX  Lines = 1 << 34 // LPUART1 Rx wakeup event.
	LPUART1TX  Lines = 1 << 35 // LPUART1 Tx wakeup event.
	SPI1       Lines = 1 << 36 // SPI1 wakeup event.
	SPI2       Lines = 1 << 37 // SPI2 wakeup event.
	SPI3       Lines = 1 << 38 // SPI3 wakeup event.
	SPI4       Lines = 1 << 39 // SPI4 wakeup event.
	SPI5       Lines = 1 << 40 // SPI5 wakeup event.
	SPI6       Lines = 1 << 41 // SPI6 wakeup event.
	MDIO       Lines = 1 << 42 // MDIO wakeup event.
	USB1       Lines = 1 << 43 // USB1 wakeup event.
	USB2       Lines = 1 << 44 // USB2 wakeup event.
	LPTIM1     Lines = 1 << 47 // LPTIM1 wakeup event.
	LPTIM2WKUP Lines = 1 << 48 // LPTIM2 wakeup event.
	LPTIM2OUT  Lines = 1 << 49 // LPTIM2 output event.
	LPTIM3WKUP Lines = 1 << 50 // LPTIM3 wakeup event.
	LPTIM3OUT  Lines = 1 << 51 // LPTIM3 output event.
	LPTIM4     Lines = 1 << 52 // LPTIM4 wakeup event.
	LPTIM5     Lines = 1 << 53 // LPTIM5 wakeup event.
	SWPMI      Lines = 1 << 54 // SWPMI wakeup event.
	WKUP1      Lines = 1 << 55 // WKUP1 wakeup event.
	WKUP2      Lines = 1 << 56 // WKUP2 wakeup event.
	WKUP3      Lines = 1 << 57 // WKUP3 wakeup event.
	WKUP4      Lines = 1 << 58 // WKUP4 wakeup event.
	WKUP5      Lines = 1 << 59 // WKUP5 wakeup event.
	WKUP6      Lines = 1 << 60 // WKUP6 wakeup event.
	RCC        Lines = 1 << 61 // RCC interrupt.
	I2C4EV     Lines = 1 << 62 // I2C4 event interrupt.
	I2C4ERR    Lines = 1 << 63 // I2C4 error interrupt.
	// BUG: lines >63 not supported
)
