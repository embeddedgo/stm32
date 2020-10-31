// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32f407 stm32f412 stm32f7x6

package exti

const (
	RTCALR  Lines = 1 << 17 // Real Time Clock Alarm event.
	OTGFS   Lines = 1 << 18 // USB OTG FS Wakeup event.
	Ether   Lines = 1 << 19 // Ethernet Wakeup event.
	OTGHS   Lines = 1 << 20 // USB OTG HS Wakeup event.
	RTCTTS  Lines = 1 << 21 // RTC Tamper and TimeStamp events.
	RTCWKUP Lines = 1 << 22 // RTC Wakeup event.
	LPTIM1  Lines = 1 << 23 // LPTIM1 asynchronous event.
)
