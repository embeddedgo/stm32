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
	"github.com/embeddedgo/stm32/hal/system"
	"github.com/embeddedgo/stm32/hal/system/timer/systick"

	"github.com/embeddedgo/display/eve/examples/gameduino/nightstrike"
)

var eveInt rtos.Note

func main() {
	system.Setup80(0, 0)
	systick.Setup(2e6)

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

	sd := spi2.Driver()
	sd.UsePinMaster(csn, spi.NSS)
	sd.UsePinMaster(sck, spi.SCK)
	sd.UsePinMaster(mosi, spi.MOSI)
	sd.UsePinMaster(miso, spi.MISO)

	lcd := eve.NewDriver(evedci.NewSPI(sd, pdn), 256)
	lcd.SetNote(&eveInt)

	////

	// Dithering causes distortion of vertical gradients on my FT811CB-HY50HD
	// (horizontal darker lines appear). The effect is most visible if both
	// Dither and Spread are set at the same time (default after reset).
	// Reducing PCLK to 15 MHz eliminates this problem but slows down the
	// display to 30 FPS.
	//dispCfg := &eve.Default800x480
	//dispCfg.ClkMHz = 15
	dispCfg := &eve.Default480x272
	if err := lcd.Init(dispCfg, nil); err != nil {
		println(err.Error())
		for {
		}
	}
	lcd.SetBacklight(96)

	nightstrike.Run(lcd)
}

//go:interrupthandler
func EXTI1_Handler() {
	exti.L1.ClearPending()
	eveInt.Wakeup()
}
