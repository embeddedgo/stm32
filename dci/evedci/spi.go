// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package evedci

import (
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
)

// SPI implements eve.DCI using Serial Peripheral Interface.
type SPI struct {
	spi      *spi.Driver
	pdn, csn gpio.Pin
	clkHz    int
	maxClkHz int
	reconf   bool
}

// NewSPI returns a new SPI based eve.DCI implementation. By default it expects
// the SPI NSS signal is used to drive EVE CSN. Use UseCSN method to move to
// software control of CSN. The pdn pin is configured as high-speed output.
func NewSPI(spidrv *spi.Driver, pdn gpio.Pin) *SPI {
	pdn.Clear()
	pdn.Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.High})
	return &SPI{spi: spidrv, pdn: pdn, maxClkHz: 30e6}
}

// UseCSN specifies the GPIO pin to be used as CSN signal. The pin is configured
// as high-speed output. If reconfSPI is true the SPI peripheral will be
// reconfigured before start any new transaction. This allows to share SPI bus
// with other external peripherals that use different SPI configuration.
func (dci *SPI) UseCSN(csn gpio.Pin, reconfSPI bool) {
	csn.Set()
	csn.Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.High})
	dci.csn = csn
	dci.reconf = reconfSPI
}

// SetMaxClk allow to reduce the clock speed set by SetClk. The eve.Driver.Init
// method uses SetClk to set 30 MHz clock speed. Use this function to reduce
// this speed as required.
func (dci *SPI) SetMaxClk(clkHz int) {
	dci.maxClkHz = clkHz
}

// SetClk setups SPI peripheral to use clkHz clock speed.
func (dci *SPI) SetClk(clkHz int) {
	if clkHz > dci.maxClkHz {
		clkHz = dci.maxClkHz
	}
	dci.clkHz = clkHz
	dci.spi.Disable()
	if dci.csn.IsValid() {
		dci.csn.Set()
		dci.spi.Setup(spi.Master|spi.SoftSS|spi.ISSHigh, clkHz)
	} else {
		dci.spi.Setup(spi.Master|spi.HardSS|spi.SSOut, clkHz)
	}
	dci.spi.SetWordSize(8) // default settings are wrong in case of F0, F3, L4.
}

func (dci *SPI) Err(clear bool) error {
	return dci.spi.Err(clear)
}

func (dci *SPI) SetPDN(pdn int) {
	dci.pdn.Store(pdn)
}

func (dci *SPI) Begin() {
	if dci.csn.IsValid() {
		dci.csn.Clear()
		if dci.reconf {
			dci.spi.Setup(spi.Master|spi.SoftSS|spi.ISSHigh, dci.clkHz)
		}
	}
	dci.spi.Enable()
}

func (dci *SPI) End() {
	if dci.csn.IsValid() {
		dci.csn.Set()
	}
	dci.spi.Disable()
}

func (dci *SPI) Read(s []byte) {
	dci.spi.WriteRead(nil, s)
}

func (dci *SPI) Write(s []byte) {
	dci.spi.WriteRead(s, nil)
}
