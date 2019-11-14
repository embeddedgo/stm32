// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32l4x6

// Package aes provides access to the registers of the AES peripheral.
//
// Instances:
//  AES  AES_BASE  AHB2  AES  Advanced encryption standard hardware accelerator
// Registers:
//  0x000 32  CR     control register
//  0x004 32  SR     status register
//  0x008 32  DINR   data input register
//  0x00C 32  DOUTR  data output register
//  0x010 32  KEYR0  key register 0
//  0x014 32  KEYR1  key register 1
//  0x018 32  KEYR2  key register 2
//  0x01C 32  KEYR3  key register 3
//  0x020 32  IVR0   initialization vector register 0
//  0x024 32  IVR1   initialization vector register 1
//  0x028 32  IVR2   initialization vector register 2
//  0x02C 32  IVR3   initialization vector register 3
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package aes

const (
	EN       CR = 0x01 << 0  //+ AES enable
	DATATYPE CR = 0x03 << 1  //+ Data type selection (for data in and data out to/from the cryptographic block)
	MODE     CR = 0x03 << 3  //+ AES operating mode
	CHMOD    CR = 0x03 << 5  //+ AES chaining mode
	CCFC     CR = 0x01 << 7  //+ Computation Complete Flag Clear
	ERRC     CR = 0x01 << 8  //+ Error clear
	CCFIE    CR = 0x01 << 9  //+ CCF flag interrupt enable
	ERRIE    CR = 0x01 << 10 //+ Error interrupt enable
	DMAINEN  CR = 0x01 << 11 //+ Enable DMA management of data input phase
	DMAOUTEN CR = 0x01 << 12 //+ Enable DMA management of data output phase
)

const (
	ENn       = 0
	DATATYPEn = 1
	MODEn     = 3
	CHMODn    = 5
	CCFCn     = 7
	ERRCn     = 8
	CCFIEn    = 9
	ERRIEn    = 10
	DMAINENn  = 11
	DMAOUTENn = 12
)

const (
	CCF   SR = 0x01 << 0 //+ Computation complete flag
	RDERR SR = 0x01 << 1 //+ Read error flag
	WRERR SR = 0x01 << 2 //+ Write error flag
)

const (
	CCFn   = 0
	RDERRn = 1
	WRERRn = 2
)

const (
	AES_DINR DINR = 0xFFFFFFFF << 0 //+ Data Input Register
)

const (
	AES_DINRn = 0
)

const (
	AES_DOUTR DOUTR = 0xFFFFFFFF << 0 //+ Data output register
)

const (
	AES_DOUTRn = 0
)

const (
	AES_KEYR0 KEYR0 = 0xFFFFFFFF << 0 //+ Data Output Register (LSB key [31:0])
)

const (
	AES_KEYR0n = 0
)

const (
	AES_KEYR1 KEYR1 = 0xFFFFFFFF << 0 //+ AES key register (key [63:32])
)

const (
	AES_KEYR1n = 0
)

const (
	AES_KEYR2 KEYR2 = 0xFFFFFFFF << 0 //+ AES key register (key [95:64])
)

const (
	AES_KEYR2n = 0
)

const (
	AES_KEYR3 KEYR3 = 0xFFFFFFFF << 0 //+ AES key register (MSB key [127:96])
)

const (
	AES_KEYR3n = 0
)

const (
	AES_IVR0 IVR0 = 0xFFFFFFFF << 0 //+ initialization vector register (LSB IVR [31:0])
)

const (
	AES_IVR0n = 0
)

const (
	AES_IVR1 IVR1 = 0xFFFFFFFF << 0 //+ Initialization Vector Register (IVR [63:32])
)

const (
	AES_IVR1n = 0
)

const (
	AES_IVR2 IVR2 = 0xFFFFFFFF << 0 //+ Initialization Vector Register (IVR [95:64])
)

const (
	AES_IVR2n = 0
)

const (
	AES_IVR3 IVR3 = 0xFFFFFFFF << 0 //+ Initialization Vector Register (MSB IVR [127:96])
)

const (
	AES_IVR3n = 0
)
