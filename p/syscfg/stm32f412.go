// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f412

// Package syscfg provides access to the registers of the SYSCFG peripheral.
//
// Instances:
//
//	SYSCFG  SYSCFG_BASE  APB2  -  System configuration controller
//
// Registers:
//
//	0x000 32  MEMRM       memory remap register
//	0x004 32  PMC         peripheral mode configuration register
//	0x008 32  EXTICR[4]   select GPIO port for EXTI line (4 x 4bit)
//	0x020 32  CMPCR       Compensation cell control register
//	0x02C 32  I2C_BUFOUT  I2C_BUFOUT
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
	ADC1DC2 PMC = 0x01 << 16 //+ ADC1DC2
)

const (
	ADC1DC2n = 16
)

const (
	CMP_PD CMPCR = 0x01 << 0 //+ Compensation cell power-down
	READY  CMPCR = 0x01 << 8 //+ READY
)

const (
	CMP_PDn = 0
	READYn  = 8
)

const (
	I2C4SCL I2C_BUFOUT = 0x01 << 0 //+ I2C4SCL
	I2C4SDA I2C_BUFOUT = 0x01 << 1 //+ I2C4SDA
)

const (
	I2C4SCLn = 0
	I2C4SDAn = 1
)
