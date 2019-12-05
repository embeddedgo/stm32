// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spi

import "github.com/embeddedgo/stm32/hal/gpio"

// UsePinsMaster is helper function that can be used to configure GPIO pins as
// required by SPI master if they all can use the same alternate function.
func UsePinsMaster(af gpio.AltFunc, sck, mosi gpio.Pin, miso_nss ...gpio.Pin) {
	cfg := &gpio.Config{Mode: gpio.Alt, Speed: gpio.VeryHigh}
	if sck.IsValid() {
		sck.Setup(cfg)
		sck.SetAltFunc(af)
	}
	if mosi.IsValid() {
		mosi.Setup(cfg)
		mosi.SetAltFunc(af)
	}
	if len(miso_nss) > 0 && miso_nss[0].IsValid() {
		miso_nss[0].Setup(&gpio.Config{Mode: gpio.AltIn})
		miso_nss[0].SetAltFunc(af)
	}
	if len(miso_nss) > 1 {
		miso_nss[1].Setup(cfg)
		miso_nss[1].SetAltFunc(af)
	}
}

// UsePinsSlave is helper function that can be used to configure GPIO pins as
// required by SPI slave if they all can use the same alternate function.
func UsePinsSlave(af gpio.AltFunc, sck, miso gpio.Pin, mosi_nss ...gpio.Pin) {
	cfg := &gpio.Config{Mode: gpio.AltIn}
	if sck.IsValid() {
		sck.Setup(cfg)
		sck.SetAltFunc(af)
	}
	if miso.IsValid() {
		miso.Setup(cfg)
		miso.SetAltFunc(af)
	}
	if len(mosi_nss) > 0 && mosi_nss[0].IsValid() {
		mosi_nss[0].Setup(&gpio.Config{Mode: gpio.Alt, Speed: gpio.VeryHigh})
		mosi_nss[0].SetAltFunc(af)
	}
	if len(mosi_nss) > 1 && mosi_nss[1].IsValid() {
		mosi_nss[1].Setup(cfg)
		mosi_nss[1].SetAltFunc(af)
	}
}
