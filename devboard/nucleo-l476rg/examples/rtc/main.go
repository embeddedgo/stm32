// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"github.com/embeddedgo/stm32/hal/system/timer/rtcst"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/system"
)

func main() {
	bi := 0 // index of the first RTC backup register used to save t0
	t0, _ := rtcst.LoadTime(bi)
	u0 := time.Unix(0, 0)
	if t0.Equal(u0) {
		// We have lost the time settings. Let's ask someone for the current
		// time and adjust the internal "wall clock".
		t0 = time.Set(askForTime())
		rtcst.StoreTime(t0, bi)
	} else {
		// The RTC has survived the system reset. Let's adjust the wall clock
		// based on the saved t0.
		time.Set(u0, t0)
	}

	for {
		println(time.Now().String())
		time.Sleep(time.Second)
	}
}

func askForTime() (old, new time.Time) {
	return time.Now(), time.Date(2020, 2, 28, 23, 59, 0, 0, time.UTC)
}
