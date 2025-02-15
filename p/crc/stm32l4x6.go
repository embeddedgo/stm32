// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32l4x6

// Package crc provides access to the registers of the CRC peripheral.
//
// Instances:
//
//	CRC  CRC_BASE  AHB1  -  Cyclic redundancy check calculation unit
//
// Registers:
//
//	0x000 32  DR    Data register
//	0x004 32  IDR   Independent data register
//	0x008 32  CR    Control register
//	0x010 32  INIT  Initial CRC value
//	0x014 32  POL   polynomial
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package crc

const (
	DR DR = 0xFFFFFFFF << 0 //+ Data register bits
)

const (
	DRn = 0
)

const (
	IDR IDR = 0xFF << 0 //+ General-purpose 8-bit data register bits
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
	Polynomialcoefficients POL = 0xFFFFFFFF << 0 //+ Programmable polynomial
)

const (
	Polynomialcoefficientsn = 0
)
