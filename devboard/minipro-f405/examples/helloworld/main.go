// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The helloworld program prints text in a loop to the standard output. The
// standard output is configured to use UART1 peripheral conected to the onboard
// CH340 USB<->UART converter and available on the left micro USB port (when you
// are looking at the MCU chip and the USB ports are above it).
package main

import (
	"fmt"

	_ "github.com/embeddedgo/stm32/devboard/minipro-f405/board/system"
)

func main() {
	for {
		fmt.Println("Hello, World!")
	}
}
