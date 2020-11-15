// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usart2

import (
	"github.com/embeddedgo/stm32/hal/usart"
)

var driver *usart.Driver

// Driver returns a ready to use driver for USART2 peripheral.
func Driver() *usart.Driver {
	if driver == nil {
		setupDriver()
	}
	return driver
}

