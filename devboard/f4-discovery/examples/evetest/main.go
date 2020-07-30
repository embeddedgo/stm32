// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"

	"github.com/embeddedgo/display/eve"
	"github.com/embeddedgo/stm32/dci/evedci"
	"github.com/embeddedgo/stm32/hal/exti"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi2"

	"github.com/embeddedgo/display/eve/examples/evetest"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/init"
)

var eveInt rtos.Note

func main() {
	// Allocate GPIO pins

	pb := gpio.B()
	pb.EnableClock(true)
	intn := pb.Pin(1)
	pdn := pb.Pin(11)
	csn := pb.Pin(12)
	sck := pb.Pin(13)
	miso := pb.Pin(14)
	mosi := pb.Pin(15)

	// EVE driver

	evedci.UseIntPin(intn)
	irq.EXTI1.Enable(rtos.IntPrioLowest, 0)

	spi2.UsePinMaster(csn, spi.NSS)
	spi2.UsePinMaster(sck, spi.SCK)
	spi2.UsePinMaster(mosi, spi.MOSI)
	spi2.UsePinMaster(miso, spi.MISO)

	lcd := eve.NewDriver(evedci.NewSPI(spi2.Driver(), pdn), 256)
	lcd.SetNote(&eveInt)

	////

	// Dithering causes distortion of vertical gradients on my FT811CB-HY50HD
	// (horizontal darker lines appear). The effect is most visible if both
	// Dither and Spread are set at the same time (default after reset).
	// Reducing PCLK to 15 MHz eliminates this problem but slows down the
	// display to 30 FPS.
	dispCfg := &eve.Default800x480
	//dispCfg.ClkMHz = 15
	//dispCfg := &eve.Default480x272
	checkErr(lcd.Init(dispCfg, &eve.Config{Dither: 1}))

	lcd.SetBacklight(96)

	evetest.Run(lcd)
}

func checkErr(err error) {
	if err == nil {
		return
	}
	println(err.Error())
	for {
	}
}

//go:interrupthandler
func EXTI1_Handler() {
	exti.L1.ClearPending()
	eveInt.Wakeup()
}
