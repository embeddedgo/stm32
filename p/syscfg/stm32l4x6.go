// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32l4x6

// Package syscfg provides access to the registers of the SYSCFG peripheral.
//
// Instances:
//
//	SYSCFG  SYSCFG_BASE  APB2  -  System configuration controller
//
// Registers:
//
//	0x000 32  MEMRMP     memory remap register
//	0x004 32  CFGR1      configuration register 1
//	0x008 32  EXTICR[4]  select GPIO port for EXTI line (4 x 4bit)
//	0x018 32  SCSR       SCSR
//	0x01C 32  CFGR2      CFGR2
//	0x020 32  SWPR       SWPR
//	0x024 32  SKR        SKR
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package syscfg

const (
	MEM_MODE MEMRMP = 0x07 << 0 //+ Memory mapping selection
	QFS      MEMRMP = 0x01 << 3 //+ QUADSPI memory mapping swap
	FB_MODE  MEMRMP = 0x01 << 8 //+ Flash Bank mode selection
)

const (
	MEM_MODEn = 0
	QFSn      = 3
	FB_MODEn  = 8
)

const (
	FWDIS       CFGR1 = 0x01 << 0  //+ Firewall disable
	BOOSTEN     CFGR1 = 0x01 << 8  //+ I/O analog switch voltage booster enable
	I2C_PB6_FMP CFGR1 = 0x01 << 16 //+ Fast-mode Plus (Fm+) driving capability activation on PB6
	I2C_PB7_FMP CFGR1 = 0x01 << 17 //+ Fast-mode Plus (Fm+) driving capability activation on PB7
	I2C_PB8_FMP CFGR1 = 0x01 << 18 //+ Fast-mode Plus (Fm+) driving capability activation on PB8
	I2C_PB9_FMP CFGR1 = 0x01 << 19 //+ Fast-mode Plus (Fm+) driving capability activation on PB9
	I2C1_FMP    CFGR1 = 0x01 << 20 //+ I2C1 Fast-mode Plus driving capability activation
	I2C2_FMP    CFGR1 = 0x01 << 21 //+ I2C2 Fast-mode Plus driving capability activation
	I2C3_FMP    CFGR1 = 0x01 << 22 //+ I2C3 Fast-mode Plus driving capability activation
	FPU_IE      CFGR1 = 0x3F << 26 //+ Floating Point Unit interrupts enable bits
)

const (
	FWDISn       = 0
	BOOSTENn     = 8
	I2C_PB6_FMPn = 16
	I2C_PB7_FMPn = 17
	I2C_PB8_FMPn = 18
	I2C_PB9_FMPn = 19
	I2C1_FMPn    = 20
	I2C2_FMPn    = 21
	I2C3_FMPn    = 22
	FPU_IEn      = 26
)

const (
	SRAM2ER  SCSR = 0x01 << 0 //+ SRAM2 Erase
	SRAM2BSY SCSR = 0x01 << 1 //+ SRAM2 busy by erase operation
)

const (
	SRAM2ERn  = 0
	SRAM2BSYn = 1
)

const (
	CLL  CFGR2 = 0x01 << 0 //+ Cortex-M4 LOCKUP (Hardfault) output enable bit
	SPL  CFGR2 = 0x01 << 1 //+ SRAM2 parity lock bit
	PVDL CFGR2 = 0x01 << 2 //+ PVD lock enable bit
	ECCL CFGR2 = 0x01 << 3 //+ ECC Lock
	SPF  CFGR2 = 0x01 << 8 //+ SRAM2 parity error flag
)

const (
	CLLn  = 0
	SPLn  = 1
	PVDLn = 2
	ECCLn = 3
	SPFn  = 8
)

const (
	P0WP  SWPR = 0x01 << 0  //+ P0WP
	P1WP  SWPR = 0x01 << 1  //+ P1WP
	P2WP  SWPR = 0x01 << 2  //+ P2WP
	P3WP  SWPR = 0x01 << 3  //+ P3WP
	P4WP  SWPR = 0x01 << 4  //+ P4WP
	P5WP  SWPR = 0x01 << 5  //+ P5WP
	P6WP  SWPR = 0x01 << 6  //+ P6WP
	P7WP  SWPR = 0x01 << 7  //+ P7WP
	P8WP  SWPR = 0x01 << 8  //+ P8WP
	P9WP  SWPR = 0x01 << 9  //+ P9WP
	P10WP SWPR = 0x01 << 10 //+ P10WP
	P11WP SWPR = 0x01 << 11 //+ P11WP
	P12WP SWPR = 0x01 << 12 //+ P12WP
	P13WP SWPR = 0x01 << 13 //+ P13WP
	P14WP SWPR = 0x01 << 14 //+ P14WP
	P15WP SWPR = 0x01 << 15 //+ P15WP
	P16WP SWPR = 0x01 << 16 //+ P16WP
	P17WP SWPR = 0x01 << 17 //+ P17WP
	P18WP SWPR = 0x01 << 18 //+ P18WP
	P19WP SWPR = 0x01 << 19 //+ P19WP
	P20WP SWPR = 0x01 << 20 //+ P20WP
	P21WP SWPR = 0x01 << 21 //+ P21WP
	P22WP SWPR = 0x01 << 22 //+ P22WP
	P23WP SWPR = 0x01 << 23 //+ P23WP
	P24WP SWPR = 0x01 << 24 //+ P24WP
	P25WP SWPR = 0x01 << 25 //+ P25WP
	P26WP SWPR = 0x01 << 26 //+ P26WP
	P27WP SWPR = 0x01 << 27 //+ P27WP
	P28WP SWPR = 0x01 << 28 //+ P28WP
	P29WP SWPR = 0x01 << 29 //+ P29WP
	P30WP SWPR = 0x01 << 30 //+ P30WP
	P31WP SWPR = 0x01 << 31 //+ SRAM2 page 31 write protection
)

const (
	P0WPn  = 0
	P1WPn  = 1
	P2WPn  = 2
	P3WPn  = 3
	P4WPn  = 4
	P5WPn  = 5
	P6WPn  = 6
	P7WPn  = 7
	P8WPn  = 8
	P9WPn  = 9
	P10WPn = 10
	P11WPn = 11
	P12WPn = 12
	P13WPn = 13
	P14WPn = 14
	P15WPn = 15
	P16WPn = 16
	P17WPn = 17
	P18WPn = 18
	P19WPn = 19
	P20WPn = 20
	P21WPn = 21
	P22WPn = 22
	P23WPn = 23
	P24WPn = 24
	P25WPn = 25
	P26WPn = 26
	P27WPn = 27
	P28WPn = 28
	P29WPn = 29
	P30WPn = 30
	P31WPn = 31
)

const (
	KEY SKR = 0xFF << 0 //+ SRAM2 write protection key for software erase
)

const (
	KEYn = 0
)
