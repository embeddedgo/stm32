// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/onewire"
	"github.com/embeddedgo/stm32/dci/owdci"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/system/console/uartcon"
	"github.com/embeddedgo/stm32/hal/usart"
	"github.com/embeddedgo/stm32/hal/usart/usart1"
	"github.com/embeddedgo/stm32/hal/usart/usart6"

	_ "github.com/embeddedgo/stm32/devboard/devebox-h743/board/system"
)

func main() {
	// GPIO pins assignment
	pa := gpio.PA()
	pa.EnableClock(true)
	conTx := pa.Pin(9)
	conRx := pa.Pin(10)
	pc := gpio.PC()
	pc.EnableClock(true)
	owTxRx := pc.Pin(6)

	// Serial console
	uartcon.Setup(usart1.Driver(), conRx, conTx, usart.Word8b, 115200, "USART1")

	// 1-Wire driver
	owm := onewire.Master{DCI: owdci.SetupUSART(usart6.Driver(), owTxRx)}

	// Support temperature sensor types.
	dtypes := []onewire.Type{onewire.DS18S20, onewire.DS18B20, onewire.DS1822}

start:
	for {
		fmt.Print("\nConfigure all DS18x20, DS1822 to 10bit resolution: ")
		if printErr(owm.SkipROM()) {
			continue start
		}
		if printErr(owm.WriteScratchpad(127, -128, onewire.T12bit)) {
			continue start
		}
		printOK()

		fmt.Print("Sending ConvertT command (SkipROM addressing): ")
		if printErr(owm.SkipROM()) {
			continue start
		}
		if printErr(owm.ConvertT()) {
			continue start
		}
		printOK()

		fmt.Print("Waiting until all devices finish the conversion: ")
		for {
			time.Sleep(50 * time.Millisecond)
			b, err := owm.ReadBit()
			if printErr(err) {
				continue start
			}
			fmt.Print(". ")
			if b != 0 {
				printOK()
				break
			}
		}
		fmt.Print("Searching for temperature sensors: ")
		for _, typ := range dtypes {
			s := onewire.NewSearch(typ, false)
			for owm.SearchNext(s) {
				d := s.Dev()
				fmt.Printf("\n %v: ", d)
				if printErr(owm.MatchROM(d)) {
					continue start
				}
				s, err := owm.ReadScratchpad()
				if printErr(err) {
					continue start
				}
				t, err := s.Temp(typ)
				if printErr(err) {
					continue start
				}
				fmt.Printf("%6.2f °C", t)
			}
			if printErr(s.Err()) {
				continue start
			}
		}
		fmt.Print("\nDone.\n\n")
		time.Sleep(2 * time.Second)
	}
}

func printErr(err error) bool {
	if err == nil {
		return false
	}
	fmt.Printf("Error: %v\n", err)
	time.Sleep(2 * time.Second)
	return true
}

func printOK() {
	fmt.Print("OK.\n")
}
