// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32l4x6

// Package pwr provides access to the registers of the PWR peripheral.
//
// Instances:
//
//	PWR  PWR_BASE  APB1  -  Power control
//
// Registers:
//
//	0x000 32  CR1                 Power control register 1
//	0x004 32  CR2                 Power control register 2
//	0x008 32  CR3                 Power control register 3
//	0x00C 32  CR4                 Power control register 4
//	0x010 32  SR1                 Power status register 1
//	0x014 32  SR2                 Power status register 2
//	0x018 32  SCR                 Power status clear register
//	0x020 32  PUDC{PUCR,PDCR}[8]  Power Port x pull-up/pull-down control registers
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package pwr

const (
	LPMS CR1 = 0x07 << 0  //+ Low-power mode selection
	DBP  CR1 = 0x01 << 8  //+ Disable backup domain write protection
	VOS  CR1 = 0x03 << 9  //+ Voltage scaling range selection
	LPR  CR1 = 0x01 << 14 //+ Low-power run
)

const (
	LPMSn = 0
	DBPn  = 8
	VOSn  = 9
	LPRn  = 14
)

const (
	PVDE  CR2 = 0x01 << 0  //+ Power voltage detector enable
	PLS   CR2 = 0x07 << 1  //+ Power voltage detector level selection
	PVME1 CR2 = 0x01 << 4  //+ Peripheral voltage monitoring 1 enable: VDDUSB vs. 1.2V
	PVME2 CR2 = 0x01 << 5  //+ Peripheral voltage monitoring 2 enable: VDDIO2 vs. 0.9V
	PVME3 CR2 = 0x01 << 6  //+ Peripheral voltage monitoring 3 enable: VDDA vs. 1.62V
	PVME4 CR2 = 0x01 << 7  //+ Peripheral voltage monitoring 4 enable: VDDA vs. 2.2V
	IOSV  CR2 = 0x01 << 9  //+ VDDIO2 Independent I/Os supply valid
	USV   CR2 = 0x01 << 10 //+ VDDUSB USB supply valid
)

const (
	PVDEn  = 0
	PLSn   = 1
	PVME1n = 4
	PVME2n = 5
	PVME3n = 6
	PVME4n = 7
	IOSVn  = 9
	USVn   = 10
)

const (
	EWUP1 CR3 = 0x01 << 0  //+ Enable Wakeup pin WKUP1
	EWUP2 CR3 = 0x01 << 1  //+ Enable Wakeup pin WKUP2
	EWUP3 CR3 = 0x01 << 2  //+ Enable Wakeup pin WKUP3
	EWUP4 CR3 = 0x01 << 3  //+ Enable Wakeup pin WKUP4
	EWUP5 CR3 = 0x01 << 4  //+ Enable Wakeup pin WKUP5
	RRS   CR3 = 0x01 << 8  //+ SRAM2 retention in Standby mode
	APC   CR3 = 0x01 << 10 //+ Apply pull-up and pull-down configuration
	EWF   CR3 = 0x01 << 15 //+ Enable internal wakeup line
)

const (
	EWUP1n = 0
	EWUP2n = 1
	EWUP3n = 2
	EWUP4n = 3
	EWUP5n = 4
	RRSn   = 8
	APCn   = 10
	EWFn   = 15
)

const (
	WP1  CR4 = 0x01 << 0 //+ Wakeup pin WKUP1 polarity
	WP2  CR4 = 0x01 << 1 //+ Wakeup pin WKUP2 polarity
	WP3  CR4 = 0x01 << 2 //+ Wakeup pin WKUP3 polarity
	WP4  CR4 = 0x01 << 3 //+ Wakeup pin WKUP4 polarity
	WP5  CR4 = 0x01 << 4 //+ Wakeup pin WKUP5 polarity
	VBE  CR4 = 0x01 << 8 //+ VBAT battery charging enable
	VBRS CR4 = 0x01 << 9 //+ VBAT battery charging resistor selection
)

const (
	WP1n  = 0
	WP2n  = 1
	WP3n  = 2
	WP4n  = 3
	WP5n  = 4
	VBEn  = 8
	VBRSn = 9
)

const (
	CWUF1 SR1 = 0x01 << 0  //+ Wakeup flag 1
	CWUF2 SR1 = 0x01 << 1  //+ Wakeup flag 2
	CWUF3 SR1 = 0x01 << 2  //+ Wakeup flag 3
	CWUF4 SR1 = 0x01 << 3  //+ Wakeup flag 4
	CWUF5 SR1 = 0x01 << 4  //+ Wakeup flag 5
	CSBF  SR1 = 0x01 << 8  //+ Standby flag
	WUFI  SR1 = 0x01 << 15 //+ Wakeup flag internal
)

const (
	CWUF1n = 0
	CWUF2n = 1
	CWUF3n = 2
	CWUF4n = 3
	CWUF5n = 4
	CSBFn  = 8
	WUFIn  = 15
)

const (
	REGLPS SR2 = 0x01 << 8  //+ Low-power regulator started
	REGLPF SR2 = 0x01 << 9  //+ Low-power regulator flag
	VOSF   SR2 = 0x01 << 10 //+ Voltage scaling flag
	PVDO   SR2 = 0x01 << 11 //+ Power voltage detector output
	PVMO1  SR2 = 0x01 << 12 //+ Peripheral voltage monitoring output: VDDUSB vs. 1.2 V
	PVMO2  SR2 = 0x01 << 13 //+ Peripheral voltage monitoring output: VDDIO2 vs. 0.9 V
	PVMO3  SR2 = 0x01 << 14 //+ Peripheral voltage monitoring output: VDDA vs. 1.62 V
	PVMO4  SR2 = 0x01 << 15 //+ Peripheral voltage monitoring output: VDDA vs. 2.2 V
)

const (
	REGLPSn = 8
	REGLPFn = 9
	VOSFn   = 10
	PVDOn   = 11
	PVMO1n  = 12
	PVMO2n  = 13
	PVMO3n  = 14
	PVMO4n  = 15
)

const (
	WUF1 SCR = 0x01 << 0 //+ Clear wakeup flag 1
	WUF2 SCR = 0x01 << 1 //+ Clear wakeup flag 2
	WUF3 SCR = 0x01 << 2 //+ Clear wakeup flag 3
	WUF4 SCR = 0x01 << 3 //+ Clear wakeup flag 4
	WUF5 SCR = 0x01 << 4 //+ Clear wakeup flag 5
	SBF  SCR = 0x01 << 8 //+ Clear standby flag
)

const (
	WUF1n = 0
	WUF2n = 1
	WUF3n = 2
	WUF4n = 3
	WUF5n = 4
	SBFn  = 8
)

const (
	PU0  PUCR = 0x01 << 0  //+ Port A pull-up bit y (y=0..15)
	PU1  PUCR = 0x01 << 1  //+ Port A pull-up bit y (y=0..15)
	PU2  PUCR = 0x01 << 2  //+ Port A pull-up bit y (y=0..15)
	PU3  PUCR = 0x01 << 3  //+ Port A pull-up bit y (y=0..15)
	PU4  PUCR = 0x01 << 4  //+ Port A pull-up bit y (y=0..15)
	PU5  PUCR = 0x01 << 5  //+ Port A pull-up bit y (y=0..15)
	PU6  PUCR = 0x01 << 6  //+ Port A pull-up bit y (y=0..15)
	PU7  PUCR = 0x01 << 7  //+ Port A pull-up bit y (y=0..15)
	PU8  PUCR = 0x01 << 8  //+ Port A pull-up bit y (y=0..15)
	PU9  PUCR = 0x01 << 9  //+ Port A pull-up bit y (y=0..15)
	PU10 PUCR = 0x01 << 10 //+ Port A pull-up bit y (y=0..15)
	PU11 PUCR = 0x01 << 11 //+ Port A pull-up bit y (y=0..15)
	PU12 PUCR = 0x01 << 12 //+ Port A pull-up bit y (y=0..15)
	PU13 PUCR = 0x01 << 13 //+ Port A pull-up bit y (y=0..15)
	PU14 PUCR = 0x01 << 14 //+ Port A pull-up bit y (y=0..15)
	PU15 PUCR = 0x01 << 15 //+ Port A pull-up bit y (y=0..15)
)

const (
	PU0n  = 0
	PU1n  = 1
	PU2n  = 2
	PU3n  = 3
	PU4n  = 4
	PU5n  = 5
	PU6n  = 6
	PU7n  = 7
	PU8n  = 8
	PU9n  = 9
	PU10n = 10
	PU11n = 11
	PU12n = 12
	PU13n = 13
	PU14n = 14
	PU15n = 15
)

const (
	PD0  PDCR = 0x01 << 0  //+ Port A pull-down bit y (y=0..15)
	PD1  PDCR = 0x01 << 1  //+ Port A pull-down bit y (y=0..15)
	PD2  PDCR = 0x01 << 2  //+ Port A pull-down bit y (y=0..15)
	PD3  PDCR = 0x01 << 3  //+ Port A pull-down bit y (y=0..15)
	PD4  PDCR = 0x01 << 4  //+ Port A pull-down bit y (y=0..15)
	PD5  PDCR = 0x01 << 5  //+ Port A pull-down bit y (y=0..15)
	PD6  PDCR = 0x01 << 6  //+ Port A pull-down bit y (y=0..15)
	PD7  PDCR = 0x01 << 7  //+ Port A pull-down bit y (y=0..15)
	PD8  PDCR = 0x01 << 8  //+ Port A pull-down bit y (y=0..15)
	PD9  PDCR = 0x01 << 9  //+ Port A pull-down bit y (y=0..15)
	PD10 PDCR = 0x01 << 10 //+ Port A pull-down bit y (y=0..15)
	PD11 PDCR = 0x01 << 11 //+ Port A pull-down bit y (y=0..15)
	PD12 PDCR = 0x01 << 12 //+ Port A pull-down bit y (y=0..15)
	PD13 PDCR = 0x01 << 13 //+ Port A pull-down bit y (y=0..15)
	PD14 PDCR = 0x01 << 14 //+ Port A pull-down bit y (y=0..15)
	PD15 PDCR = 0x01 << 15 //+ Port A pull-down bit y (y=0..15)
)

const (
	PD0n  = 0
	PD1n  = 1
	PD2n  = 2
	PD3n  = 3
	PD4n  = 4
	PD5n  = 5
	PD6n  = 6
	PD7n  = 7
	PD8n  = 8
	PD9n  = 9
	PD10n = 10
	PD11n = 11
	PD12n = 12
	PD13n = 13
	PD14n = 14
	PD15n = 15
)
