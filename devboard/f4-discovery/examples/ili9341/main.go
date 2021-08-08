// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"github.com/embeddedgo/display/tft"
	"github.com/embeddedgo/display/tft/ili9341"
	"github.com/embeddedgo/stm32/dci/tftdci"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi2"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/init"
)

func main() {
	pb := gpio.B()
	pb.EnableClock(true)
	sck := pb.Pin(13)
	mosi := pb.Pin(15)

	pd := gpio.D()
	pd.EnableClock(true)
	reset := pd.Pin(8)
	cs := pd.Pin(9)
	dc := pd.Pin(10)

	cfg := &gpio.Config{Mode: gpio.Out, Speed: gpio.Low}
	dc.Clear()
	dc.Setup(cfg)
	reset.Clear() // assert RESET
	reset.Setup(cfg)
	cs.Clear() // assert CS
	cs.Setup(cfg)

	spidrv := spi2.Driver()
	spidrv.UsePinMaster(sck, spi.SCK)
	spidrv.UsePinMaster(mosi, spi.MOSI)
	spidrv.Setup(spi.Master|spi.MSBF|spi.CPOL0|spi.CPHA0|spi.SoftSS|spi.ISSHigh, 21e6)
	spidrv.SetWordSize(8)
	spidrv.Enable()

	disp := tft.NewDisplay(tftdci.NewSPI(spidrv, dc), 240, 320, ili9341.FillRect)

	line := make([]uint16, 240)
	for i := 0; ; i++ {
		reset.Clear()
		time.Sleep(time.Millisecond)
		reset.Set()

		ili9341.Init(disp.DCI(), ili9341.BGR|ili9341.MX, false)

		for k := range line {
			line[k] = 0
		}
		for i := 0; i < 320; i++ {
			disp.DCI().WriteWords(line...)
		}

		time.Sleep(time.Second)

		for i := 0; i < 320; i++ {
			for k := range line {
				x := 64 * (i + k) / (320 + 240)
				x = (x/2)<<11 | x<<5 | (x / 2)
				line[k] = uint16(x)
			}
			disp.DCI().WriteWords(line...)
			time.Sleep(5 * time.Millisecond)
		}

		time.Sleep(time.Second)
	}
}
