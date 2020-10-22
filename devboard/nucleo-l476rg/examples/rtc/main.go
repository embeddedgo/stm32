package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/leds"
)

func main() {
	for {
		leds.User.SetOn()
		print(fmt.Sprintln(time.Now()))
		time.Sleep(time.Second)
		leds.User.SetOff()
		print(fmt.Sprintln(time.Now()))
		time.Sleep(time.Second)

	}
}

/*
func main() {
	rtcst.Setup(rtcst.LSE, 1, 32768)
	leds.User.SetOn()
	RTC := rtc.RTC()
	for {
		dr := RTC.DR.Load()
		tr := RTC.TR.Load()
		ns := rtcst.Nanotime()
		println(fmt.Sprintf(
			"%06x %d %06x %d",
			dr&^rtc.WDU, rtcst.Dayofmonth(uint(ns/(24*60*60*1e9))), tr, ns,
		))
		time.Sleep(time.Second)
	}
}
*/
