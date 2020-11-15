// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spi2

import "github.com/embeddedgo/stm32/hal/spi"

var driver *spi.Driver

// Driver returns ready to use driver for SPI2 peripheral.
func Driver() *spi.Driver {
	if driver == nil {
		setupDriver()
	}
	return driver
}
