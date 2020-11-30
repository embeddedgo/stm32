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

func checkErr(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
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
	con := termfs.New(name, d, d)
	con.SetReplaceIn(termfs.CRtoLF)
	con.SetReplaceOut(termfs.LFtoCRLF)
	rtos.Mount(con, "/dev/console")
	var err error
	os.Stdin, err = os.OpenFile("/dev/console", syscall.O_RDONLY, 0)
	checkErr(err)
	os.Stdout, err = os.OpenFile("/dev/console", syscall.O_WRONLY, 0)
	checkErr(err)
	os.Stderr = os.Stdout
}
