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
	clkHz   int
	mode    spi.Config
	started bool
	reconf  bool
}

// NewSPI returns new SPI based implementation of display/tft.DCI that uses the
// configured SPI driver and the configred data/command pin. Mode can be 0, 1,
// 2, 3 (see https://en.wikipedia.org/wiki/Serial_Peripheral_Interface)
func NewSPI(drv *spi.Driver, dc gpio.Pin, mode, clkHz int) *SPI {
	dci := new(SPI)
	dci.spi = drv
	dci.dc = dc
	dci.clkHz = clkHz
	switch mode {
	case 0:
		dci.mode = spi.CPOL0 | spi.CPHA0
	case 1:
		dci.mode = spi.CPOL0 | spi.CPHA1
	case 2:
		dci.mode = spi.CPOL1 | spi.CPHA0
	default: // 3
		dci.mode = spi.CPOL1 | spi.CPHA1
	}
	drv.Setup(spi.Master|spi.HardSS|spi.SSOut|dci.mode, clkHz)
	return dci
}

func (dci *SPI) UseCSN(csn gpio.Pin, reconfSPI bool) {
	dci.csn = csn
	dci.reconf = reconfSPI
	csn.Set()
	csn.Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.High})
	dci.spi.Setup(spi.Master|spi.SoftSS|spi.ISSHigh, dci.clkHz)
}

func (dci *SPI) Driver() *spi.Driver  { return dci.spi }
func (dci *SPI) Err(clear bool) error { return dci.spi.Err(clear) }
func (dci *SPI) DC() gpio.Pin         { return dci.dc }

func (dci *SPI) Cmd(cmd byte) {
	if !dci.started {
		dci.started = true
		if dci.csn.IsValid() {
			dci.csn.Clear()
			if dci.reconf {
				dci.spi.Setup(spi.Master|spi.SoftSS|spi.ISSHigh, dci.clkHz)
			}
		}
		dci.spi.Enable()
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
	dci.spi.SetWordSize(8)
	dci.spi.WriteRead(p, nil)
}

func (dci *SPI) WriteString(s string) {
	dci.spi.SetWordSize(8)
	dci.spi.WriteStringRead(s, nil)
}

func (dci *SPI) WriteByteN(b byte, n int) {
	dci.spi.SetWordSize(8)
	dci.spi.WriteByteN(b, n)
}

func (dci *SPI) WriteWords(p []uint16) {
	dci.spi.SetWordSize(16)
	dci.spi.WriteRead16(p, nil)
}

func (dci *SPI) WriteWordN(w uint16, n int) {
	dci.spi.SetWordSize(16)
	dci.spi.WriteWord16N(w, n)
}

func (dci *SPI) ReadBytes(p []byte) {
	dci.spi.SetWordSize(8)
	dci.spi.WriteRead(nil, p)
}
