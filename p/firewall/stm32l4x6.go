// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32l4x6

// Package firewall provides access to the registers of the FIREWALL peripheral.
//
// Instances:
//  FIREWALL  FIREWALL_BASE  APB2  -  Firewall
// Registers:
//  0x000 32  CSSA    Code segment start address
//  0x004 32  CSL     Code segment length
//  0x008 32  NVDSSA  Non-volatile data segment start address
//  0x00C 32  NVDSL   Non-volatile data segment length
//  0x010 32  VDSSA   Volatile data segment start address
//  0x014 32  VDSL    Volatile data segment length
//  0x020 32  CR      Configuration register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package firewall

const (
	ADD CSSA = 0xFFFF << 8 //+ code segment start address
)

const (
	ADDn = 8
)

const (
	LENG CSL = 0x3FFF << 8 //+ code segment length
)

const (
	LENGn = 8
)

const (
	ADD NVDSSA = 0xFFFF << 8 //+ Non-volatile data segment start address
)

const (
	ADDn = 8
)

const (
	LENG NVDSL = 0x3FFF << 8 //+ Non-volatile data segment length
)

const (
	LENGn = 8
)

const (
	ADD VDSSA = 0x3FF << 6 //+ Volatile data segment start address
)

const (
	ADDn = 6
)

const (
	LENG VDSL = 0x3FF << 6 //+ Non-volatile data segment length
)

const (
	LENGn = 6
)

const (
	FPA CR = 0x01 << 0 //+ Firewall pre alarm
	VDS CR = 0x01 << 1 //+ Volatile data shared
	VDE CR = 0x01 << 2 //+ Volatile data execution
)

const (
	FPAn = 0
	VDSn = 1
	VDEn = 2
)
