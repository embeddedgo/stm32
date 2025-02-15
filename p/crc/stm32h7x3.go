// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package crc provides access to the registers of the CRC peripheral.
//
// Instances:
//
//	CRC  CRC_BASE  AHB4  -  Cryptographic processor
//
// Registers:
//
//	0x000 32  DR    Data register
//	0x004 32  IDR   Independent Data register
//	0x008 32  CR    Control register
//	0x010 32  INIT  Initial CRC value
//	0x014 32  POL   CRC polynomial
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package crc

const (
	DR DR = 0xFFFFFFFF << 0 //+ Data Register
)

const (
	DRn = 0
)

const (
	IDR IDR = 0xFFFFFFFF << 0 //+ Independent Data register
)

const (
	IDRn = 0
)

const (
	RESET    CR = 0x01 << 0 //+ RESET bit
	POLYSIZE CR = 0x03 << 3 //+ Polynomial size
	REV_IN   CR = 0x03 << 5 //+ Reverse input data
	REV_OUT  CR = 0x01 << 7 //+ Reverse output data
)

const (
	RESETn    = 0
	POLYSIZEn = 3
	REV_INn   = 5
	REV_OUTn  = 7
)

const (
	CRC_INIT INIT = 0xFFFFFFFF << 0 //+ Programmable initial CRC value
)

const (
	CRC_INITn = 0
)

const (
	POL POL = 0xFFFFFFFF << 0 //+ Programmable polynomial
)

const (
	POLn = 0
)
