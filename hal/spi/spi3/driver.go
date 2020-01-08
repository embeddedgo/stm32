// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spi3

import (
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
)

var driver *spi.Driver

// Driver returns ready to use driver for SPI3 peripheral.
func Driver() *spi.Driver {
	if driver == nil {
		setupDriver()
	}
	return driver
}

// UsePinsMaster is helper function that can be used to configure GPIO pins as
// required by SPI master device
func UsePinsMaster(sck, mosi gpio.Pin, miso_nss ...gpio.Pin) {
	spi.UsePinsMaster(altFunc, sck, mosi, miso_nss...)
}

// UsePinsSlave is helper function that can be used to configure GPIO pins as
// required by SPI slave device.
func UsePinsSlave(sck, miso gpio.Pin, mosi_nss ...gpio.Pin) {
	spi.UsePinsSlave(altFunc, sck, miso, mosi_nss...)
}
