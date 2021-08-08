// Copyright 2021 The Embedded Go authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tftdci

import (
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
)

// SPI is an implementation of the display/tft.DCI (Display Controller
// Interface) that uses an SPI peripheral to communicate with the display in
// what is known as 4-line serial mode. As the DCI does not require to read
// any data from a display the only required SPI signals are SCK and MOSI.
type SPI struct {
	spi *spi.Driver
	dc  gpio.Pin
}

// NewSPI returns new SPI based implementation of display/tft.DCI that uses the
// configured SPI driver and the configred data/command pin.
func NewSPI(drv *spi.Driver, dc gpio.Pin) *SPI {
	return &SPI{drv, dc}
}

func (dci *SPI) Driver() *spi.Driver { return dci.spi }
func (dci *SPI) DC() gpio.Pin        { return dci.dc }

func (dci *SPI) Cmd(cmd byte) {
	dci.dc.Clear()
	dci.spi.SetWordSize(8)
	dci.spi.WriteReadByte(cmd)
	dci.spi.SetWordSize(16)
	dci.dc.Set()
}

func (dci *SPI) WriteBytes(p ...uint8) {
	dci.spi.SetWordSize(8)
	dci.spi.WriteRead(p, nil)
	dci.spi.SetWordSize(16)
}

func (dci *SPI) WriteWords(p ...uint16) {
	dci.spi.WriteRead16(p, nil)
}

func (dci *SPI) WriteWordN(w uint16, n int) {
	dci.spi.WriteWord16N(w, n)
}

func (dci *SPI) Err(clear bool) error {
	return dci.spi.Err(clear)
}
