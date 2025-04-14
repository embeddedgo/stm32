// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spih


type Config uint64

const (
	FrameFormat = Config(sp) // Frame format:
	MSPI        = Config(0)  // - Motorola SPI
	SyncSerial  = Config(ti) // - Texas Instruments Synchronous Serial

	CPHA0 = Config(0)    // Sample on leading edge.
	CPHA1 = Config(cpha) // Sample on trailing edge.
	CPOL0 = Config(0)    // Clock idle state is 0.
	CPOL1 = Config(cpol) // Clock idle state is 1.

	WordLen = Config(wlen) // Data word length (no support for SPIv3 17-24 bit):
	Word4b  = Config(w4)   // - 4 bit (SPI v1.3+)
	Word5b  = Config(w5)   // - 5 bit (SPI v1.3+)
	Word6b  = Config(w6)   // - 6 bit (SPI v1.3+)
	Word7b  = Config(w7)   // - 7 bit (SPI v1.3+)
	Word8b  = Config(w8)   // - 8 bit
	Word9b  = Config(w9)   // - 9 bit (SPI v1.3+)
	Word10b = Config(w10)  // - 10 bit (SPI v1.3+)
	Word11b = Config(w11)  // - 11 bit (SPI v1.3+)
	Word12b = Config(w12)  // - 12 bit (SPI v1.3+)
	Word13b = Config(w13)  // - 13 bit (SPI v1.3+)
	Word14b = Config(w14)  // - 14 bit (SPI v1.3+)
	Word15b = Config(w15)  // - 15 bit (SPI v1.3+)
	Word16b = Config(w16)  // - 16 bit
	Word17b = Config(w17)  // - 17 bit (SPI v2+)
	Word18b = Config(w18)  // - 18 bit (SPI v2+)
	Word19b = Config(w19)  // - 19 bit (SPI v2+)
	Word20b = Config(w20)  // - 20 bit (SPI v2+)
	Word21b = Config(w21)  // - 21 bit (SPI v2+)
	Word22b = Config(w22)  // - 22 bit (SPI v2+)
	Word23b = Config(w23)  // - 23 bit (SPI v2+)
	Word24b = Config(w24)  // - 24 bit (SPI v2+)
	Word25b = Config(w25)  // - 25 bit (SPI v2+)
	Word26b = Config(w26)  // - 26 bit (SPI v2+)
	Word27b = Config(w27)  // - 27 bit (SPI v2+)
	Word28b = Config(w28)  // - 28 bit (SPI v2+)
	Word29b = Config(w29)  // - 29 bit (SPI v2+)
	Word30b = Config(w30)  // - 30 bit (SPI v2+)
	Word31b = Config(w31)  // - 31 bit (SPI v2+)
	Word32b = Config(w32)  // - 32 bit (SPI v2+)

	ClkDiv    = Config(clkDiv)      // Clock divider (see also SetBaudrate):
	ClkDiv2   = Config(0 * clkDiv4) // - baud rate = ker_ck/2
	ClkDiv4   = Config(1 * clkDiv4) // - baud rate = ker_ck/4
	ClkDiv8   = Config(2 * clkDiv4) // - baud rate = ker_ck/8
	ClkDiv16  = Config(3 * clkDiv4) // - baud rate = ker_ck/16
	ClkDiv32  = Config(4 * clkDiv4) // - baud rate = ker_ck/32
	ClkDiv64  = Config(5 * clkDiv4) // - baud rate = ker_ck/64
	ClkDiv128 = Config(6 * clkDiv4) // - baud rate = ker_ck/128
	ClkDiv256 = Config(7 * clkDiv4) // - baud rate = ker_ck/256

	MSBF = Config(0)        // Most significant bit first.
	LSBF = Config(lsbfirst) // Least significant bit first.

	HalfDuplex = Config(hd) // MOSI is used for Tx and Rx
)

