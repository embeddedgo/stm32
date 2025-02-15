// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f215

// Package syscfg provides access to the registers of the SYSCFG peripheral.
//
// Instances:
//
//	SYSCFG  SYSCFG_BASE  APB2  -  System configuration controller
//
// Registers:
//
//	0x000 32  MEMRM      memory remap register
//	0x004 32  PMC        peripheral mode configuration register
//	0x008 32  EXTICR[4]  select GPIO port for EXTI line (4 x 4bit)
//	0x020 32  CMPCR      Compensation cell control register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package syscfg

const (
	MEM_MODE MEMRM = 0x03 << 0 //+ MEM_MODE
)

const (
	MEM_MODEn = 0
)

const (
	MII_RMII_SEL PMC = 0x01 << 23 //+ Ethernet PHY interface selection
)

const (
	MII_RMII_SELn = 23
)

const (
	CMP_PD CMPCR = 0x01 << 0 //+ Compensation cell power-down
	READY  CMPCR = 0x01 << 7 //+ READY
)

const (
	CMP_PDn = 0
	READYn  = 7
)
