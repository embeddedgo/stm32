// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spi

// CR1
const (
	cpha     = 1 << 0
	cpol     = 1 << 1
	mstr     = 1 << 2
	br       = 7 << 3
	brn      = 3
	spe      = 1 << 6
	lsbfirst = 1 << 7
	ssi      = 1 << 8
	ssm      = 1 << 9
	rxonly   = 1 << 10
	dff      = 1 << 11 // SPIv1
	crcl     = 1 << 11 // SPIv2
	crcnext  = 1 << 12
	crcen    = 1 << 13
	bidioe   = 1 << 14
	bidimode = 1 << 15
)

// CR2
const (
	rxdmaen = 1 << 0
	txdmaen = 1 << 1
	ssoe    = 1 << 2
	nssp    = 1 << 3 // SPIv2
	frf     = 1 << 4
	errie   = 1 << 5
	errien  = 5
	rxneie  = 1 << 6
	txeie   = 1 << 7
	ds      = 15 << 8 // SPIv2
	dsn     = 8       // SPIv2
	frxth   = 1 << 12 // SPIv2
	ldmarx  = 1 << 13 // SPIv2
	ldmatx  = 1 << 14 // SPIv2
)

// SR
const (
	rxne   = 1 << 0
	txe    = 1 << 1
	chside = 1 << 2 // SPIv1
	udr    = 1 << 3 // SPIv1
	crcerr = 1 << 4
	modf   = 1 << 5
	ovr    = 1 << 6
	bsy    = 1 << 7
	tifrfe = 1 << 8
	frlvl  = 3 << 9  // SPIv2
	ftlvl  = 3 << 11 // SPIv2
)
