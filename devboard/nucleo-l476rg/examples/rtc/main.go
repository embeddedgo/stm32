package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/leds"
	"github.com/embeddedgo/stm32/hal/system/timer/rtcst"
)

func main() {
	if rtcst.HasReset() {
		time.Set(time.Now(), time.Date(2020, 11, 2, 0, 45, 53, 0, time.UTC))
	} else {
		time.Set(rtcst.SavedTime())
	}
	for {
		println(fmt.Sprint(time.Now()))
		leds.User.Set(leds.User.Get() + 1)
		time.Sleep(time.Second)

	}
}
