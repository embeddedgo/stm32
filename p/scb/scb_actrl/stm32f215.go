// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f215

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
	DISFOLD        ACTRL = 0x01 << 2  //+ DISFOLD
	FPEXCODIS      ACTRL = 0x01 << 10 //+ FPEXCODIS
	DISRAMODE      ACTRL = 0x01 << 11 //+ DISRAMODE
	DISITMATBFLUSH ACTRL = 0x01 << 12 //+ DISITMATBFLUSH
)

const (
	DISFOLDn        = 2
	FPEXCODISn      = 10
	DISRAMODEn      = 11
	DISITMATBFLUSHn = 12
)
