// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Display draws on the connected display.
package main

import (
	"time"

	"github.com/embeddedgo/display/pix/displays"
	"github.com/embeddedgo/display/pix/examples"

	"github.com/embeddedgo/stm32/dci/tftdci"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi2"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/system"
)

func main() {
	// Assign GPIO pins

	pb := gpio.PB()
	pb.EnableClock(true)
	csn := pb.Pin(12)
	sck := pb.Pin(13)
	miso := pb.Pin(14)
	mosi := pb.Pin(15)

	pd := gpio.PD()
	pd.EnableClock(true)
	reset := pd.Pin(8) // optional
	dc := pd.Pin(10)

	// Configure peripherals

	reset.Set()
	reset.Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.Low})

	spidrv := spi2.Driver()
	spidrv.UsePinMaster(sck, spi.SCK)
	spidrv.UsePinMaster(miso, spi.MISO)
	spidrv.UsePinMaster(mosi, spi.MOSI)
	spidrv.UsePinMaster(csn, spi.NSS) // use hardware slave-select

	// Hardware reset - optional

	reset.Clear()
	time.Sleep(time.Millisecond)
	reset.Set()

	//dp := displays.Adafruit_0i96_128x64_OLED_SSD1306()
	//dp := displays.Adafruit_1i5_128x128_OLED_SSD1351()
	//dp := displays.Adafruit_1i54_240x240_IPS_ST7789()
	dp := displays.Adafruit_2i8_240x320_TFT_ILI9341()
	//dp := displays.ERTFTM_1i54_240x240_IPS_ST7789()
	//dp := displays.MSP4022_4i0_320x480_TFT_ILI9486()
	//dp := displays.Waveshare_1i5_128x128_OLED_SSD1351()

	// Most of the displays accept significant overclocking.
	//dp.MaxReadClk *= 2
	//dp.MaxWriteClk *= 2

	dci := tftdci.NewSPI(spidrv, dc, spi.CPOL0|spi.CPHA0, dp.MaxWriteClk, dp.MaxWriteClk)
	//dci.UseCSN(csn, false) // use software slave-select

	// Run

	disp := dp.New(dci)
	for {
		examples.RotateDisplay(disp)
		//examples.DrawText(disp)
		//examples.GraphicsTest(disp)
	}
}
