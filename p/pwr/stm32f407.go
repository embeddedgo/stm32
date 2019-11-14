// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f407

// Package pwr provides access to the registers of the PWR peripheral.
//
// Instances:
//  PWR  PWR_BASE  APB1  PVD  Power control
// Registers:
//  0x000 32  CR   power control register
//  0x004 32  CSR  power control/status register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package pwr

const (
	LPDS CR = 0x01 << 0 //+ Low-power deep sleep
	PDDS CR = 0x01 << 1 //+ Power down deepsleep
	CWUF CR = 0x01 << 2 //+ Clear wakeup flag
	CSBF CR = 0x01 << 3 //+ Clear standby flag
	PVDE CR = 0x01 << 4 //+ Power voltage detector enable
	PLS  CR = 0x07 << 5 //+ PVD level selection
	DBP  CR = 0x01 << 8 //+ Disable backup domain write protection
	FPDS CR = 0x01 << 9 //+ Flash power down in Stop mode
)

const (
	LPDSn = 0
	PDDSn = 1
	CWUFn = 2
	CSBFn = 3
	PVDEn = 4
	PLSn  = 5
	DBPn  = 8
	FPDSn = 9
)

const (
	WUF    CSR = 0x01 << 0  //+ Wakeup flag
	SBF    CSR = 0x01 << 1  //+ Standby flag
	PVDO   CSR = 0x01 << 2  //+ PVD output
	BRR    CSR = 0x01 << 3  //+ Backup regulator ready
	EWUP   CSR = 0x01 << 8  //+ Enable WKUP pin
	BRE    CSR = 0x01 << 9  //+ Backup regulator enable
	VOSRDY CSR = 0x01 << 14 //+ Regulator voltage scaling output selection ready bit
)

const (
	WUFn    = 0
	SBFn    = 1
	PVDOn   = 2
	BRRn    = 3
	EWUPn   = 8
	BREn    = 9
	VOSRDYn = 14
)
