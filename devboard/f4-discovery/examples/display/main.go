// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"github.com/embeddedgo/display/pix/examples"

	"github.com/embeddedgo/stm32/dci/tftdci"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi2"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/init"
)

// The SPI2 speeds below are max. for this MCU but can be out of display spec.
// If your display works unstable reduce to 10.5e6 and 5.25e6 or even more.
const (
	writeClk = 21e6
	readClk  = 21e6
)

func main() {
	// Assign GPIO pins

	pb := gpio.B()
	pb.EnableClock(true)
	csn := pb.Pin(12)
	sck := pb.Pin(13)
	miso := pb.Pin(14)
	mosi := pb.Pin(15)

	pd := gpio.D()
	pd.EnableClock(true)
	reset := pd.Pin(8)
	dc := pd.Pin(10)

	// Configure peripherals

	reset.Set()
	reset.Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.Low})

	spidrv := spi2.Driver()
	spidrv.UsePinMaster(sck, spi.SCK)
	spidrv.UsePinMaster(miso, spi.MISO)
	spidrv.UsePinMaster(mosi, spi.MOSI)
	spidrv.UsePinMaster(csn, spi.NSS) // use hardware slave-select
	dci := tftdci.NewSPI(spidrv, dc, spi.CPOL0|spi.CPHA0, writeClk, readClk)
	//dci.UseCSN(csn, false) // use software slave-select

	// Reset

	reset.Clear()
	time.Sleep(time.Millisecond)
	reset.Set()

	// Run

	//disp := examples.Adafruit_1i5_128x128_OLED_SSD1351(dci)
	//disp := examples.Adafruit_1i54_240x240_IPS_ST7789(dci)
	//disp := examples.Adafruit_2i8_240x320_TFT_ILI9341(dci)
	//disp := examples.ERTFTM_1i54_240x240_IPS_ST7789(dci)
	//disp := examples.MSP4022_4i0_320x480_TFT_ILI9486(dci)
	disp := examples.Waveshare_1i5_128x128_OLED_SSD1351(dci)

	for {
		//examples.Colors(disp)
		//examples.RotateDisplay(disp)
		//examples.DrawText(disp)
		examples.GraphicsTest(disp)
	}
}
