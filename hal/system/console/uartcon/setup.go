// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uartcon

import (
	"embedded/rtos"
	"os"
	"syscall"

	"github.com/embeddedgo/fs/termfs"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/usart"
)

var uart *usart.Driver

func write(_ int, p []byte) int {
	n, _ := uart.Write(p)
	return n
}

// Setup setpus an USART peripheral represented by d to work as system console.
func Setup(d *usart.Driver, rx, tx gpio.Pin, conf usart.Config, baudrate int, lf string, name string) {
	rxport := rx.Port()
	rxport.EnableClock(true)
	if txport := tx.Port(); txport != rxport {
		txport.EnableClock(true)
	}
	d.UsePin(rx, usart.RXD)
	d.UsePin(tx, usart.TXD)
	d.Setup(conf, baudrate)
	d.EnableTx()
	d.EnableRx(nil)
	uart = d
	rtos.SetSystemWriter(write)
	rtos.Mount("/dev/console", termfs.New(name, d, d, "\r\n"))
	os.Stdin, _ = os.OpenFile("/dev/console", syscall.O_RDONLY, 0)
	os.Stdout, _ = os.OpenFile("/dev/console", syscall.O_RDWR, 0)
	os.Stderr = os.Stdout
}
