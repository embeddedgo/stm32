// Copyright 2021 The Embedded Go authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tftdci

import (
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
)

// SPI is an implementation of the display/tft.DCI that uses an SPI peripheral
// to communicate with the display in what is known as 4-line serial mode.
type SPI struct {
	spi     *spi.Driver
	dc      gpio.Pin
	csn     gpio.Pin
	wmode   spi.Config
	rmode   spi.Config
	started bool
	reconf  bool
}

var gpioOut = &gpio.Config{Mode: gpio.Out, Speed: gpio.High}

// NewSPI returns new SPI based implementation of tftdrv.DCI. It properly
// configures the provided SPI driver and DC pin to communicate with a display
// controller. Select the SPI mode (CPOL,CPHA), write and read clock speed
// according to the display controller specification. Note that the maximum
// speed may be limited by the concrete instance of STM32 SPI peripheral, the
// bus topology and the specific display design.
func NewSPI(drv *spi.Driver, dc gpio.Pin, mode spi.Config, wclkHz, rclkHz int) *SPI {
	dci := new(SPI)
	dci.spi = drv
	dci.dc = dc
	mode |= spi.Master | spi.SSOut
	dci.wmode = mode | drv.Periph().BR(wclkHz)
	dci.rmode = mode | drv.Periph().BR(rclkHz)
	dc.Clear()
	dc.Setup(gpioOut)
	drv.Setup(dci.wmode, 0)
	return dci
}

// UseCSN setups the underlying SPI peripheral in software slave select mode and
// setups csn as slave select pin. If reconf is true the SPI peripheral is
// reconfigured by any Cmd call so it can be shared with other applications
// (exclusive acces is required until End call).
func (dci *SPI) UseCSN(csn gpio.Pin, reconf bool) {
	dci.csn = csn
	dci.reconf = reconf
	dci.wmode |= spi.SoftSS | spi.ISSHigh
	dci.rmode |= spi.SoftSS | spi.ISSHigh
	csn.Set()
	csn.Setup(gpioOut)
	dci.spi.Periph().SetConfig(dci.wmode, 0)
}

func (dci *SPI) Driver() *spi.Driver  { return dci.spi }
func (dci *SPI) Err(clear bool) error { return dci.spi.Err(clear) }
func (dci *SPI) DC() gpio.Pin         { return dci.dc }

func start(dci *SPI) {
	dci.started = true
	if dci.csn.IsValid() {
		dci.csn.Clear()
		if dci.reconf {
			dci.spi.Periph().SetConfig(dci.wmode, 0)
		}
	}
	dci.spi.Enable()
}

func (dci *SPI) Cmd(cmd byte) {
	if !dci.started {
		start(dci)
	}
	dci.dc.Clear()
	dci.spi.SetWordSize(8)
	dci.spi.WriteReadByte(cmd)
	dci.dc.Set()
}

func (dci *SPI) End() {
	dci.started = false
	if dci.csn.IsValid() {
		dci.csn.Set()
	}
	dci.spi.Disable()
}

func (dci *SPI) WriteBytes(p []uint8) {
	if !dci.started {
		start(dci)
	}
	dci.spi.SetWordSize(8)
	dci.spi.WriteRead(p, nil)
}

func (dci *SPI) WriteString(s string) {
	if !dci.started {
		start(dci)
	}
	dci.spi.SetWordSize(8)
	dci.spi.WriteStringRead(s, nil)
}

func (dci *SPI) WriteByteN(b byte, n int) {
	if !dci.started {
		start(dci)
	}
	dci.spi.SetWordSize(8)
	dci.spi.WriteByteN(b, n)
}

func (dci *SPI) WriteWords(p []uint16) {
	if !dci.started {
		start(dci)
	}
	dci.spi.SetWordSize(16)
	dci.spi.WriteRead16(p, nil)
}

func (dci *SPI) WriteWordN(w uint16, n int) {
	if !dci.started {
		start(dci)
	}
	dci.spi.SetWordSize(16)
	dci.spi.WriteWord16N(w, n)
}

func (dci *SPI) ReadBytes(p []byte) {
	if !dci.started {
		start(dci)
	}
	dci.spi.SetWordSize(8)
	dci.spi.Periph().SetConfig(dci.rmode, 0)
	dci.spi.WriteRead(nil, p)
	dci.spi.Periph().SetConfig(dci.wmode, 0)
}
