// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package spih

// CFG1 (FIFO threshold 4 bytes)
const (
	wlen = uint64(DSIZE | FTHVL)
	w4   = 3 | 3<<FTHVLn
	w5   = 4 | 3<<FTHVLn
	w6   = 5 | 3<<FTHVLn
	w7   = 6 | 3<<FTHVLn
	w8   = 7 | 3<<FTHVLn
	w9   = 8 | 1<<FTHVLn
	w10  = 9 | 1<<FTHVLn
	w11  = 10 | 1<<FTHVLn
	w12  = 11 | 1<<FTHVLn
	w13  = 12 | 1<<FTHVLn
	w14  = 13 | 1<<FTHVLn
	w15  = 14 | 1<<FTHVLn
	w16  = 15 | 1<<FTHVLn
	w17  = 16 | 0<<FTHVLn
	w18  = 17 | 0<<FTHVLn
	w19  = 18 | 0<<FTHVLn
	w20  = 19 | 0<<FTHVLn
	w21  = 20 | 0<<FTHVLn
	w22  = 21 | 0<<FTHVLn
	w23  = 22 | 0<<FTHVLn
	w24  = 23 | 0<<FTHVLn
	w25  = 24 | 0<<FTHVLn
	w26  = 25 | 0<<FTHVLn
	w27  = 26 | 0<<FTHVLn
	w28  = 27 | 0<<FTHVLn
	w29  = 28 | 0<<FTHVLn
	w30  = 29 | 0<<FTHVLn
	w31  = 30 | 0<<FTHVLn
	w32  = 31 | 0<<FTHVLn

	clkDiv  = uint64(MBR)
	clkDiv4 = 1 << MBRn
)

// CFG2
const (
	sp = uint64(SP) << 32
	ti = uint64(1) << (SPn + 32)

	cpha = uint64(CPHA) << 32
	cpol = uint64(CPOL) << 32

	lsbfirst = uint64(LSBFRST) << 32

	dir = uint64(COMM) << 32
	tx  = uint64(1) << (COMMn + 3)
	rx  = uint64(2) << (COMMn + 3)
	hd  = uint64(3) << (COMMn + 3)
)

const cfgMask = Config(wlen | clkDiv | sp | cpha | cpol | lsbfirst | dir)
