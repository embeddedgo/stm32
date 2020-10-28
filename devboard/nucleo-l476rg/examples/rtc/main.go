package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/leds"
)

func main() {
	for {
		leds.User.SetOn()
		println(fmt.Sprint(time.Now()))
		time.Sleep(time.Second)
		leds.User.SetOff()
		println(fmt.Sprint(time.Now()))
		time.Sleep(time.Second)

	}
}
