// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"embedded/arch/cortexm/systim"
	"embedded/rtos"
	"os"
	"runtime"
	"syscall"

	"github.com/embeddedgo/fs/termfs"
	"github.com/embeddedgo/stm32/hal/usart"
)

// QEMU supported devics/peripherals:
//
// ARM Cortex-M3, Cortex M4F, ADC, EXTI, USART, SPI, SYSCFG, TIMER,
//
// Missing devices:
//
// RCC, DCMI, CAN, CRC, DAC, DMA, Ethernet, Flash interface, GPIO controller,
// I2C, I2S, PWR, RNG, RTC, SDIO, USB, IWDG, WWDG

func systickSetup(periodns int64) {
	runtime.LockOSThread()
	pl, _ := rtos.SetPrivLevel(0)
	systim.Setup(periodns, 1e3, false)
	rtos.SetPrivLevel(pl)
	runtime.UnlockOSThread()
	rtos.SetSystemTimer(systim.Nanotime, nil)
}

var uart *uartDriver

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

func termfsSetupLight(p *usart.Periph, name string) {
	uart = newUARTDriver(p)
	rtos.SetSystemWriter(write)

	con := termfs.NewLight(name, uart, uart)
	rtos.Mount(con, "/dev/console")
	var err error
	os.Stdin, err = os.OpenFile("/dev/console", syscall.O_RDONLY, 0)
	checkErr(err)
	os.Stdout, err = os.OpenFile("/dev/console", syscall.O_WRONLY, 0)
	checkErr(err)
	os.Stderr = os.Stdout
}

func init() {
	//systickSetup(2e6)
	termfsSetupLight(usart.USART1(), "USART1")
}
