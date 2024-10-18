// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f215

// Package adc_common provides access to the registers of the ADC_Common peripheral.
//
// Instances:
//
//	ADC_Common  ADC_Common_BASE  -  -  Common ADC registers
//
// Registers:
//
//	0x000 32  CSR  ADC Common status register
//	0x004 32  CCR  ADC common control register
//	0x008 32  CDR  ADC common regular data register for dual and triple modes
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package adc_common

const (
	AWD1   CSR = 0x01 << 0  //+ Analog watchdog flag of ADC 1
	EOC1   CSR = 0x01 << 1  //+ End of conversion of ADC 1
	JEOC1  CSR = 0x01 << 2  //+ Injected channel end of conversion of ADC 1
	JSTRT1 CSR = 0x01 << 3  //+ Injected channel Start flag of ADC 1
	STRT1  CSR = 0x01 << 4  //+ Regular channel Start flag of ADC 1
	OVR1   CSR = 0x01 << 5  //+ Overrun flag of ADC 1
	AWD2   CSR = 0x01 << 8  //+ Analog watchdog flag of ADC 2
	EOC2   CSR = 0x01 << 9  //+ End of conversion of ADC 2
	JEOC2  CSR = 0x01 << 10 //+ Injected channel end of conversion of ADC 2
	JSTRT2 CSR = 0x01 << 11 //+ Injected channel Start flag of ADC 2
	STRT2  CSR = 0x01 << 12 //+ Regular channel Start flag of ADC 2
	OVR2   CSR = 0x01 << 13 //+ Overrun flag of ADC 2
	AWD3   CSR = 0x01 << 16 //+ Analog watchdog flag of ADC 3
	EOC3   CSR = 0x01 << 17 //+ End of conversion of ADC 3
	JEOC3  CSR = 0x01 << 18 //+ Injected channel end of conversion of ADC 3
	JSTRT3 CSR = 0x01 << 19 //+ Injected channel Start flag of ADC 3
	STRT3  CSR = 0x01 << 20 //+ Regular channel Start flag of ADC 3
	OVR3   CSR = 0x01 << 21 //+ Overrun flag of ADC3
)

const (
	AWD1n   = 0
	EOC1n   = 1
	JEOC1n  = 2
	JSTRT1n = 3
	STRT1n  = 4
	OVR1n   = 5
	AWD2n   = 8
	EOC2n   = 9
	JEOC2n  = 10
	JSTRT2n = 11
	STRT2n  = 12
	OVR2n   = 13
	AWD3n   = 16
	EOC3n   = 17
	JEOC3n  = 18
	JSTRT3n = 19
	STRT3n  = 20
	OVR3n   = 21
)

const (
	MULT    CCR = 0x1F << 0  //+ Multi ADC mode selection
	DELAY   CCR = 0x0F << 8  //+ Delay between 2 sampling phases
	DDS     CCR = 0x01 << 13 //+ DMA disable selection for multi-ADC mode
	DMA     CCR = 0x03 << 14 //+ Direct memory access mode for multi ADC mode
	ADCPRE  CCR = 0x03 << 16 //+ ADC prescaler
	VBATE   CCR = 0x01 << 22 //+ VBAT enable
	TSVREFE CCR = 0x01 << 23 //+ Temperature sensor and VREFINT enable
)

const (
	MULTn    = 0
	DELAYn   = 8
	DDSn     = 13
	DMAn     = 14
	ADCPREn  = 16
	VBATEn   = 22
	TSVREFEn = 23
)

const (
	DATA1 CDR = 0xFFFF << 0  //+ 1st data item of a pair of regular conversions
	DATA2 CDR = 0xFFFF << 16 //+ 2nd data item of a pair of regular conversions
)

const (
	DATA1n = 0
	DATA2n = 16
)
