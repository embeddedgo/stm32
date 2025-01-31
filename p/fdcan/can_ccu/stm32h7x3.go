// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package can_ccu provides access to the registers of the CAN_CCU peripheral.
//
// Instances:
//
//	CAN_CCU  CAN_CCU_BASE  -  -  CCU registers
//
// Registers:
//
//	0x000 32  CREL   Clock Calibration Unit Core Release Register
//	0x004 32  CCFG   Calibration Configuration Register
//	0x008 32  CSTAT  Calibration Status Register
//	0x00C 32  CWD    Calibration Watchdog Register
//	0x010 32  IR     Clock Calibration Unit Interrupt Register
//	0x014 32  IE     Clock Calibration Unit Interrupt Enable Register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package can_ccu

const (
	DAY     CREL = 0xFF << 0  //+ Time Stamp Day
	MON     CREL = 0xFF << 8  //+ Time Stamp Month
	YEAR    CREL = 0x0F << 16 //+ Time Stamp Year
	SUBSTEP CREL = 0x0F << 20 //+ Sub-step of Core Release
	STEP    CREL = 0x0F << 24 //+ Step of Core Release
	REL     CREL = 0x0F << 28 //+ Core Release
)

const (
	DAYn     = 0
	MONn     = 8
	YEARn    = 16
	SUBSTEPn = 20
	STEPn    = 24
	RELn     = 28
)

const (
	TQBT CCFG = 0x1F << 0  //+ Time Quanta per Bit Time
	BCC  CCFG = 0x01 << 6  //+ Bypass Clock Calibration
	CFL  CCFG = 0x01 << 7  //+ Calibration Field Length
	OCPM CCFG = 0xFF << 8  //+ Oscillator Clock Periods Minimum
	CDIV CCFG = 0x0F << 16 //+ Clock Divider
	SWR  CCFG = 0x01 << 31 //+ Software Reset
)

const (
	TQBTn = 0
	BCCn  = 6
	CFLn  = 7
	OCPMn = 8
	CDIVn = 16
	SWRn  = 31
)

const (
	OCPC CSTAT = 0x3FFFF << 0 //+ Oscillator Clock Period Counter
	TQC  CSTAT = 0x7FF << 18  //+ Time Quanta Counter
	CALS CSTAT = 0x03 << 30   //+ Calibration State
)

const (
	OCPCn = 0
	TQCn  = 18
	CALSn = 30
)

const (
	WDC CWD = 0xFFFF << 0  //+ WDC
	WDV CWD = 0xFFFF << 16 //+ WDV
)

const (
	WDCn = 0
	WDVn = 16
)

const (
	CWE IR = 0x01 << 0 //+ Calibration Watchdog Event
	CSC IR = 0x01 << 1 //+ Calibration State Changed
)

const (
	CWEn = 0
	CSCn = 1
)

const (
	CWEE IE = 0x01 << 0 //+ Calibration Watchdog Event Enable
	CSCE IE = 0x01 << 1 //+ Calibration State Changed Enable
)

const (
	CWEEn = 0
	CSCEn = 1
)
