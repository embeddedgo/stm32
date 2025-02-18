// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image/color"
	"time"

	"github.com/embeddedgo/rgbled/ws281x/wsuart"
	_ "github.com/embeddedgo/stm32/devboard/devebox-h743/board/system"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/usart"
	"github.com/embeddedgo/stm32/hal/usart/usart1"
)

func main() {
	// GPIO pins assignment
	pa := gpio.PA()
	pa.EnableClock(true)
	tx := pa.Pin(9) // CN5 D8

	// Configure and enable USART
	u := usart1.Driver()
	u.UsePin(tx, usart.TXD)
	u.Setup(usart.Word8b|usart.InvTx, wsuart.BaudWS2812)
	u.EnableTx()

	colors := []color.RGBA{
		{127, 0, 0, 255},
		{255, 0, 0, 255},
		{0, 127, 0, 255},
		{0, 255, 0, 255},
		{0, 0, 127, 255},
		{0, 0, 255, 255},
		{127, 127, 0, 255},
		{255, 255, 0, 255},
		{0, 127, 127, 255},
		{0, 255, 255, 255},
		{127, 0, 127, 255},
		{255, 0, 255, 255},
		{127, 127, 127, 255},
		{255, 255, 255, 255},
	}
	grb := wsuart.GRB
	strip := wsuart.Make(24)

	for i := 0; ; i++ {
		for k := range strip {
			strip.Clear()
			strip[k] = grb.Pixel(colors[i%len(colors)])
			u.Write(strip.Bytes())
			time.Sleep(time.Second / 2)
		}
	}
}
