// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package owdci provides implementation of onewire.DCI interface.
package owdci

import (
	"io"
	"time"

	"github.com/embeddedgo/onewire"
	"github.com/embeddedgo/stm32/hal/usart"
)

// USART wraps an usart.Driver to implements the onewire.DCI.
type USART usart.Driver

func usartDrv(dci *USART) *usart.Driver { return (*usart.Driver)(dci) }

// SetupUSART configures d to be used as onewire.DCI.
func SetupUSART(d *usart.Driver) *USART {
	d.Setup(usart.Word8b|usart.HalfDuplex, 115200)
	p := d.Periph()
	p.SetConf3(p.Conf3() | usart.OneBit) // disables noise detection
	d.EnableRx(32)
	d.EnableTx()
	return (*USART)(d)
}

func checkErr(d *usart.Driver, err error) error {
	if err == nil {
		return nil
	}
	d.DiscardRx()
	return err
}

func (dci *USART) Reset() error {
	d := usartDrv(dci)
	d.SetBaudrate(9600)
	err := d.WriteByte(0xf0)
	if err != nil {
		return err
	}
	d.SetReadTimeout(time.Second)
	r, err := d.ReadByte()
	if err != nil {
		return err
	}
	if r == 0xf0 {
		return onewire.ErrNoResponse
	}
	d.SetBaudrate(115200)
	return nil
}

func sendRecvSlot(d *usart.Driver, slot byte) (byte, error) {
	if err := d.WriteByte(slot); err != nil {
		d.DiscardRx()
		return 0, err
	}
	d.SetReadTimeout(time.Second)
	b, err := d.ReadByte()
	return b, checkErr(d, err)
}

func sendRecv(d *usart.Driver, slots *[8]byte) error {
	if _, err := d.Write(slots[:]); err != nil {
		d.DiscardRx()
		return err
	}
	d.SetReadTimeout(time.Second)
	_, err := io.ReadFull(d, slots[:])
	return checkErr(d, err)
}

func (dci *USART) ReadBit() (int, error) {
	slot, err := sendRecvSlot(usartDrv(dci), 0xff)
	if err != nil {
		return 0, err
	}
	return int(slot) & 1, nil
}

func (dci *USART) WriteBit(bit int) error {
	var b byte
	if bit&1 != 0 {
		b = 0xff
	}
	d := usartDrv(dci)
	slot, err := sendRecvSlot(d, b)
	if err != nil {
		return err
	}
	if slot != b {
		d.DiscardRx()
		return onewire.ErrBusFault
	}
	return nil
}

func (dci *USART) ReadByte() (byte, error) {
	slots := [8]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	err := sendRecv(usartDrv(dci), &slots)
	var v int
	for i, slot := range slots {
		v += int(slot&1) << uint(i)
	}
	return byte(v), err
}

func (dci *USART) WriteByte(b byte) error {
	var slots [8]byte
	v := int(b)
	for i := range slots {
		if v&1 != 0 {
			slots[i] = 0xff
		}
		v >>= 1
	}
	d := usartDrv(dci)
	if err := sendRecv(d, &slots); err != nil {
		return err
	}
	v = int(b)
	for i, slot := range slots {
		r := v & (1 << uint(i))
		if r != 0 {
			r = 0xff
		}
		if int(slot) != r {
			d.DiscardRx()
			return onewire.ErrBusFault
		}
	}
	return nil
}
