// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package spih

import (
	"embedded/rtos"
	"math/bits"
	"runtime"
)

type master struct {
	cfg1 CFG1
	cfg2 CFG2
	mode int8 // 1: wite, 2: read, 3: write and read
	slow bool
	done rtos.Note
}

func masterConfig(d *Master) Config {
	return (Config(d.cfg1) | Config(d.cfg2<<32)) & cfgMask
}

func masterSetConfig(d *Master, cfg Config) {
	cfg1 := CFG1(cfg)
	cfg2 := CFG2(cfg>>32) |
		MASTER | // master mode
		AFCNTR | // keep pins under controll when disabled
		SSOE //     enable hardware Slave Select (CS) output

	if d.cfg1 != cfg1 {
		d.cfg1 = cfg1
		d.mode = 0 // force reconfiguration
	}
	if d.cfg2 != cfg2 {
		d.cfg2 = cfg2
		d.mode = 0 // force reconfiguration
	}
}

func kernHz(p *Periph) uint {
	return 100e6 // TODO: determine at runtime
}

func masterBaudrate(d *Master) int {
	log2div := uint(d.cfg1>>MBRn) + 1
	return int(kernHz(d.p) >> log2div)
}

func masterClkDiv(p *Periph, baudrate int) Config {
	if baudrate <= 0 {
		panic("spi: baudrate<=0")
	}
	div := (kernHz(p) + uint(baudrate) - 1) / uint(baudrate)
	log2div := uint(bits.Len(div) - 1)
	if div&^(1<<log2div) != 0 {
		log2div++
	}
	if log2div != 0 {
		log2div--
	}
	return Config(log2div << MBRn)
}

func masterWaitTxDone(d *Master) {
	p, slow := d.p, d.slow
	for p.SR.LoadBits(TXC) == 0 {
		if slow {
			runtime.Gosched()
		}
	}
}

const masterEnable = SPE | MASRX

func masterDisable(d *Master) {
	masterWaitTxDone(d)
	d.p.CR1.Store(0)
}

type dataWord interface{ ~uint8 | ~uint16 | ~uint32 }
