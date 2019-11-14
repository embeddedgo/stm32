// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package exti

const (
	RTCALR  Lines = 1 << 17 // Real Time Clock Alarm event.
	USB     Lines = 1 << 18 // USB wakeup.
	RTCTTS  Lines = 1 << 19 // RTC Tamper and TimeStamp events.
	RTCWKUP Lines = 1 << 20 // RTC Wakeup event.
)
