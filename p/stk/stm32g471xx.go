// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32g471xx

// Package stk provides access to the registers of the STK peripheral.
//
// Instances:
//
//	STK  STK_BASE  -  -  SysTick timer
//
// Registers:
//
//	0x000 32  CTRL   SysTick control and status register
//	0x004 32  LOAD   SysTick reload value register
//	0x008 32  VAL    SysTick current value register
//	0x00C 32  CALIB  SysTick calibration value register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package stk

const (
	ENABLE    CTRL = 0x01 << 0  //+ Counter enable
	TICKINT   CTRL = 0x01 << 1  //+ SysTick exception request enable
	CLKSOURCE CTRL = 0x01 << 2  //+ Clock source selection
	COUNTFLAG CTRL = 0x01 << 16 //+ COUNTFLAG
)

const (
	ENABLEn    = 0
	TICKINTn   = 1
	CLKSOURCEn = 2
	COUNTFLAGn = 16
)

const (
	RELOAD LOAD = 0xFFFFFF << 0 //+ RELOAD value
)

const (
	RELOADn = 0
)

const (
	CURRENT VAL = 0xFFFFFF << 0 //+ Current counter value
)

const (
	CURRENTn = 0
)

const (
	TENMS CALIB = 0xFFFFFF << 0 //+ Calibration value
	SKEW  CALIB = 0x01 << 30    //+ SKEW flag: Indicates whether the TENMS value is exact
	NOREF CALIB = 0x01 << 31    //+ NOREF flag. Reads as zero
)

const (
	TENMSn = 0
	SKEWn  = 30
	NOREFn = 31
)
