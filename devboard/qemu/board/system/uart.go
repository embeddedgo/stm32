// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"unsafe"

	"github.com/embeddedgo/stm32/hal/usart"
)

type uartDriver struct {
	p *usart.Periph
}

func newUARTDriver(p *usart.Periph) *uartDriver {
	p.EnableTx()
	return &uartDriver{p}
}

func (d *uartDriver) Write(s []byte) (int, error) {
	p := d.p
	for _, b := range s {
		for {
			if ev, _ := p.Status(); ev&usart.TxEmpty != 0 {
				break
			}
		}
		p.Store(uint32(b))
	}
	return len(s), nil
}

func (d *uartDriver) WriteString(s string) (int, error) {
	return d.Write(unsafe.Slice(unsafe.StringData(s), len(s)))
}

func (d *uartDriver) Read(s []byte) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}
	p := d.p
	for {
		if ev, _ := p.Status(); ev&usart.RxNotEmpty != 0 {
			break
		}
	}
	s[0] = byte(p.Load())
	return 1, nil
}
