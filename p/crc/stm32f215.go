// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f215

// Package crc provides access to the registers of the CRC peripheral.
//
// Instances:
//
//	CRC  CRC_BASE  AHB1  -  cyclic redundancy check calculation unit
//
// Registers:
//
//	0x000 32  DR   Data register
//	0x004 32  IDR  Independent data register
//	0x008 32  CR   Control register
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
	RESET CR = 0x01 << 0 //+ reset bit
)

const (
	RESETn = 0
)
