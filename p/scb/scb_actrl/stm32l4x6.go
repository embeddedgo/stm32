// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32l4x6

// Package scb_actrl provides access to the registers of the SCB_ACTRL peripheral.
//
// Instances:
//
//	SCB_ACTRL  SCB_ACTRL_BASE  -  -  System control block ACTLR
//
// Registers:
//
//	0x000 32  ACTRL  Auxiliary control register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package scb_actrl

const (
	DISMCYCINT ACTRL = 0x01 << 0 //+ DISMCYCINT
	DISDEFWBUF ACTRL = 0x01 << 1 //+ DISDEFWBUF
	DISFOLD    ACTRL = 0x01 << 2 //+ DISFOLD
	DISFPCA    ACTRL = 0x01 << 8 //+ DISFPCA
	DISOOFP    ACTRL = 0x01 << 9 //+ DISOOFP
)

const (
	DISMCYCINTn = 0
	DISDEFWBUFn = 1
	DISFOLDn    = 2
	DISFPCAn    = 8
	DISOOFPn    = 9
)
