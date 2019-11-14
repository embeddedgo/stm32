// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f215

// Package tim9 provides access to the registers of the TIM9 peripheral.
//
// Instances:
//  TIM9   TIM9_BASE   APB2  TIM1_BRK_TIM9*   General-purpose-timers
//  TIM12  TIM12_BASE  APB1  TIM8_BRK_TIM12*  General-purpose-timers
// Registers:
//  0x000 32  CR1           control register 1
//  0x004 32  CR2           control register 2
//  0x008 32  SMCR          slave mode control register
//  0x00C 32  DIER          DMA/Interrupt enable register
//  0x010 32  SR            status register
//  0x014 32  EGR           event generation register
//  0x018 32  CCMR1_Output  capture/compare mode register 1 (output mode)
//  0x018 32  CCMR1_Input   capture/compare mode register 1 (input mode)
//  0x020 32  CCER          capture/compare enable register
//  0x024 32  CNT           counter
//  0x028 32  PSC           prescaler
//  0x02C 32  ARR           auto-reload register
//  0x034 32  CCR1          capture/compare register 1
//  0x038 32  CCR2          capture/compare register 2
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package tim9

const (
	CEN  CR1 = 0x01 << 0 //+ Counter enable
	UDIS CR1 = 0x01 << 1 //+ Update disable
	URS  CR1 = 0x01 << 2 //+ Update request source
	OPM  CR1 = 0x01 << 3 //+ One-pulse mode
	ARPE CR1 = 0x01 << 7 //+ Auto-reload preload enable
	CKD  CR1 = 0x03 << 8 //+ Clock division
)

const (
	CENn  = 0
	UDISn = 1
	URSn  = 2
	OPMn  = 3
	ARPEn = 7
	CKDn  = 8
)

const (
	MMS CR2 = 0x07 << 4 //+ Master mode selection
)

const (
	MMSn = 4
)

const (
	SMS SMCR = 0x07 << 0 //+ Slave mode selection
	TS  SMCR = 0x07 << 4 //+ Trigger selection
	MSM SMCR = 0x01 << 7 //+ Master/Slave mode
)

const (
	SMSn = 0
	TSn  = 4
	MSMn = 7
)

const (
	UIE   DIER = 0x01 << 0 //+ Update interrupt enable
	CC1IE DIER = 0x01 << 1 //+ Capture/Compare 1 interrupt enable
	CC2IE DIER = 0x01 << 2 //+ Capture/Compare 2 interrupt enable
	TIE   DIER = 0x01 << 6 //+ Trigger interrupt enable
)

const (
	UIEn   = 0
	CC1IEn = 1
	CC2IEn = 2
	TIEn   = 6
)

const (
	UIF   SR = 0x01 << 0  //+ Update interrupt flag
	CC1IF SR = 0x01 << 1  //+ Capture/compare 1 interrupt flag
	CC2IF SR = 0x01 << 2  //+ Capture/Compare 2 interrupt flag
	TIF   SR = 0x01 << 6  //+ Trigger interrupt flag
	CC1OF SR = 0x01 << 9  //+ Capture/Compare 1 overcapture flag
	CC2OF SR = 0x01 << 10 //+ Capture/compare 2 overcapture flag
)

const (
	UIFn   = 0
	CC1IFn = 1
	CC2IFn = 2
	TIFn   = 6
	CC1OFn = 9
	CC2OFn = 10
)

const (
	UG   EGR = 0x01 << 0 //+ Update generation
	CC1G EGR = 0x01 << 1 //+ Capture/compare 1 generation
	CC2G EGR = 0x01 << 2 //+ Capture/compare 2 generation
	TG   EGR = 0x01 << 6 //+ Trigger generation
)

const (
	UGn   = 0
	CC1Gn = 1
	CC2Gn = 2
	TGn   = 6
)

const (
	CC1S  CCMR1_Output = 0x03 << 0  //+ Capture/Compare 1 selection
	OC1FE CCMR1_Output = 0x01 << 2  //+ Output Compare 1 fast enable
	OC1PE CCMR1_Output = 0x01 << 3  //+ Output Compare 1 preload enable
	OC1M  CCMR1_Output = 0x07 << 4  //+ Output Compare 1 mode
	CC2S  CCMR1_Output = 0x03 << 8  //+ Capture/Compare 2 selection
	OC2FE CCMR1_Output = 0x01 << 10 //+ Output Compare 2 fast enable
	OC2PE CCMR1_Output = 0x01 << 11 //+ Output Compare 2 preload enable
	OC2M  CCMR1_Output = 0x07 << 12 //+ Output Compare 2 mode
)

const (
	CC1Sn  = 0
	OC1FEn = 2
	OC1PEn = 3
	OC1Mn  = 4
	CC2Sn  = 8
	OC2FEn = 10
	OC2PEn = 11
	OC2Mn  = 12
)

const (
	CC1S   CCMR1_Input = 0x03 << 0  //+ Capture/Compare 1 selection
	IC1PSC CCMR1_Input = 0x03 << 2  //+ Input capture 1 prescaler
	IC1F   CCMR1_Input = 0x07 << 4  //+ Input capture 1 filter
	CC2S   CCMR1_Input = 0x03 << 8  //+ Capture/Compare 2 selection
	IC2PSC CCMR1_Input = 0x03 << 10 //+ Input capture 2 prescaler
	IC2F   CCMR1_Input = 0x07 << 12 //+ Input capture 2 filter
)

const (
	CC1Sn   = 0
	IC1PSCn = 2
	IC1Fn   = 4
	CC2Sn   = 8
	IC2PSCn = 10
	IC2Fn   = 12
)

const (
	CC1E  CCER = 0x01 << 0 //+ Capture/Compare 1 output enable
	CC1P  CCER = 0x01 << 1 //+ Capture/Compare 1 output Polarity
	CC1NP CCER = 0x01 << 3 //+ Capture/Compare 1 output Polarity
	CC2E  CCER = 0x01 << 4 //+ Capture/Compare 2 output enable
	CC2P  CCER = 0x01 << 5 //+ Capture/Compare 2 output Polarity
	CC2NP CCER = 0x01 << 7 //+ Capture/Compare 2 output Polarity
)

const (
	CC1En  = 0
	CC1Pn  = 1
	CC1NPn = 3
	CC2En  = 4
	CC2Pn  = 5
	CC2NPn = 7
)

const (
	CNT CNT = 0xFFFF << 0 //+ counter value
)

const (
	CNTn = 0
)

const (
	PSC PSC = 0xFFFF << 0 //+ Prescaler value
)

const (
	PSCn = 0
)

const (
	ARR ARR = 0xFFFF << 0 //+ Auto-reload value
)

const (
	ARRn = 0
)

const (
	CCR1 CCR1 = 0xFFFF << 0 //+ Capture/Compare 1 value
)

const (
	CCR1n = 0
)

const (
	CCR2 CCR2 = 0xFFFF << 0 //+ Capture/Compare 2 value
)

const (
	CCR2n = 0
)
