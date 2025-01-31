// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package tim15 provides access to the registers of the TIM peripheral.
//
// Instances:
//
//	TIM15  TIM15_BASE  APB2  TIM15  General purpose timers
//
// Registers:
//
//	0x000 32  CR1           control register 1
//	0x004 32  CR2           control register 2
//	0x008 32  SMCR          slave mode control register
//	0x00C 32  DIER          DMA/Interrupt enable register
//	0x010 32  SR            status register
//	0x014 32  EGR           event generation register
//	0x018 32  CCMR1_Output  capture/compare mode register (output mode)
//	0x018 32  CCMR1_Input   capture/compare mode register 1 (input mode)
//	0x020 32  CCER          capture/compare enable register
//	0x024 32  CNT           counter
//	0x028 32  PSC           prescaler
//	0x02C 32  ARR           auto-reload register
//	0x030 32  RCR           repetition counter register
//	0x034 32  CCR1          capture/compare register 1
//	0x038 32  CCR2          capture/compare register 2
//	0x044 32  BDTR          break and dead-time register
//	0x048 32  DCR           DMA control register
//	0x04C 32  DMAR          DMA address for full transfer
//	0x060 32  AF1           TIM15 alternate fdfsdm1_breakon register 1
//	0x068 32  TISEL         TIM15 input selection register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package tim15

const (
	CEN      CR1 = 0x01 << 0  //+ Counter enable
	UDIS     CR1 = 0x01 << 1  //+ Update disable
	URS      CR1 = 0x01 << 2  //+ Update request source
	OPM      CR1 = 0x01 << 3  //+ One-pulse mode
	ARPE     CR1 = 0x01 << 7  //+ Auto-reload preload enable
	CKD      CR1 = 0x03 << 8  //+ Clock division
	UIFREMAP CR1 = 0x01 << 11 //+ UIF status bit remapping
)

const (
	CENn      = 0
	UDISn     = 1
	URSn      = 2
	OPMn      = 3
	ARPEn     = 7
	CKDn      = 8
	UIFREMAPn = 11
)

const (
	CCPC  CR2 = 0x01 << 0  //+ Capture/compare preloaded control
	CCUS  CR2 = 0x01 << 2  //+ Capture/compare control update selection
	CCDS  CR2 = 0x01 << 3  //+ Capture/compare DMA selection
	MMS   CR2 = 0x07 << 4  //+ Master mode selection
	TI1S  CR2 = 0x01 << 7  //+ TI1 selection
	OIS1  CR2 = 0x01 << 8  //+ Output Idle state 1
	OIS1N CR2 = 0x01 << 9  //+ Output Idle state 1
	OIS2  CR2 = 0x01 << 10 //+ Output Idle state 2
)

const (
	CCPCn  = 0
	CCUSn  = 2
	CCDSn  = 3
	MMSn   = 4
	TI1Sn  = 7
	OIS1n  = 8
	OIS1Nn = 9
	OIS2n  = 10
)

const (
	SMS    SMCR = 0x07 << 0  //+ Slave mode selection
	TS_2_0 SMCR = 0x07 << 4  //+ Trigger selection
	MSM    SMCR = 0x01 << 7  //+ Master/Slave mode
	SMS_3  SMCR = 0x01 << 16 //+ Slave mode selection bit 3
	TS_4_3 SMCR = 0x03 << 20 //+ Trigger selection - bit 4:3
)

const (
	SMSn    = 0
	TS_2_0n = 4
	MSMn    = 7
	SMS_3n  = 16
	TS_4_3n = 20
)

const (
	UIE   DIER = 0x01 << 0  //+ Update interrupt enable
	CC1IE DIER = 0x01 << 1  //+ Capture/Compare 1 interrupt enable
	CC2IE DIER = 0x01 << 2  //+ Capture/Compare 2 interrupt enable
	COMIE DIER = 0x01 << 5  //+ COM interrupt enable
	TIE   DIER = 0x01 << 6  //+ Trigger interrupt enable
	BIE   DIER = 0x01 << 7  //+ Break interrupt enable
	UDE   DIER = 0x01 << 8  //+ Update DMA request enable
	CC1DE DIER = 0x01 << 9  //+ Capture/Compare 1 DMA request enable
	CC2DE DIER = 0x01 << 10 //+ Capture/Compare 2 DMA request enable
	COMDE DIER = 0x01 << 13 //+ COM DMA request enable
	TDE   DIER = 0x01 << 14 //+ Trigger DMA request enable
)

const (
	UIEn   = 0
	CC1IEn = 1
	CC2IEn = 2
	COMIEn = 5
	TIEn   = 6
	BIEn   = 7
	UDEn   = 8
	CC1DEn = 9
	CC2DEn = 10
	COMDEn = 13
	TDEn   = 14
)

const (
	UIF   SR = 0x01 << 0  //+ Update interrupt flag
	CC1IF SR = 0x01 << 1  //+ Capture/compare 1 interrupt flag
	CC2IF SR = 0x01 << 2  //+ Capture/Compare 2 interrupt flag
	COMIF SR = 0x01 << 5  //+ COM interrupt flag
	TIF   SR = 0x01 << 6  //+ Trigger interrupt flag
	BIF   SR = 0x01 << 7  //+ Break interrupt flag
	CC1OF SR = 0x01 << 9  //+ Capture/Compare 1 overcapture flag
	CC2OF SR = 0x01 << 10 //+ Capture/compare 2 overcapture flag
)

const (
	UIFn   = 0
	CC1IFn = 1
	CC2IFn = 2
	COMIFn = 5
	TIFn   = 6
	BIFn   = 7
	CC1OFn = 9
	CC2OFn = 10
)

const (
	UG   EGR = 0x01 << 0 //+ Update generation
	CC1G EGR = 0x01 << 1 //+ Capture/compare 1 generation
	CC2G EGR = 0x01 << 2 //+ Capture/compare 2 generation
	COMG EGR = 0x01 << 5 //+ Capture/Compare control update generation
	TG   EGR = 0x01 << 6 //+ Trigger generation
	BG   EGR = 0x01 << 7 //+ Break generation
)

const (
	UGn   = 0
	CC1Gn = 1
	CC2Gn = 2
	COMGn = 5
	TGn   = 6
	BGn   = 7
)

const (
	CC1S   CCMR1_Output = 0x03 << 0  //+ Capture/Compare 1 selection
	OC1FE  CCMR1_Output = 0x01 << 2  //+ Output Compare 1 fast enable
	OC1PE  CCMR1_Output = 0x01 << 3  //+ Output Compare 1 preload enable
	OC1M   CCMR1_Output = 0x07 << 4  //+ Output Compare 1 mode
	CC2S   CCMR1_Output = 0x03 << 8  //+ Capture/Compare 2 selection
	OC2FE  CCMR1_Output = 0x01 << 10 //+ Output Compare 2 fast enable
	OC2PE  CCMR1_Output = 0x01 << 11 //+ Output Compare 2 preload enable
	OC2M   CCMR1_Output = 0x07 << 12 //+ Output Compare 2 mode
	OC1M_3 CCMR1_Output = 0x01 << 16 //+ Output Compare 1 mode bit 3
	OC2M_3 CCMR1_Output = 0x01 << 24 //+ Output Compare 2 mode bit 3
)

const (
	CC1Sn   = 0
	OC1FEn  = 2
	OC1PEn  = 3
	OC1Mn   = 4
	CC2Sn   = 8
	OC2FEn  = 10
	OC2PEn  = 11
	OC2Mn   = 12
	OC1M_3n = 16
	OC2M_3n = 24
)

const (
	CC1S   CCMR1_Input = 0x03 << 0  //+ Capture/Compare 1 selection
	IC1PSC CCMR1_Input = 0x03 << 2  //+ Input capture 1 prescaler
	IC1F   CCMR1_Input = 0x0F << 4  //+ Input capture 1 filter
	CC2S   CCMR1_Input = 0x03 << 8  //+ Capture/Compare 2 selection
	IC2PSC CCMR1_Input = 0x03 << 10 //+ Input capture 2 prescaler
	IC2F   CCMR1_Input = 0x0F << 12 //+ Input capture 2 filter
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
	CC1NE CCER = 0x01 << 2 //+ Capture/Compare 1 complementary output enable
	CC1NP CCER = 0x01 << 3 //+ Capture/Compare 1 output Polarity
	CC2E  CCER = 0x01 << 4 //+ Capture/Compare 2 output enable
	CC2P  CCER = 0x01 << 5 //+ Capture/Compare 2 output Polarity
	CC2NP CCER = 0x01 << 7 //+ Capture/Compare 2 output Polarity
)

const (
	CC1En  = 0
	CC1Pn  = 1
	CC1NEn = 2
	CC1NPn = 3
	CC2En  = 4
	CC2Pn  = 5
	CC2NPn = 7
)

const (
	CNT    CNT = 0xFFFF << 0 //+ counter value
	UIFCPY CNT = 0x01 << 31  //+ UIF copy
)

const (
	CNTn    = 0
	UIFCPYn = 31
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
	REP RCR = 0xFF << 0 //+ Repetition counter value
)

const (
	REPn = 0
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

const (
	DTG  BDTR = 0xFF << 0  //+ Dead-time generator setup
	LOCK BDTR = 0x03 << 8  //+ Lock configuration
	OSSI BDTR = 0x01 << 10 //+ Off-state selection for Idle mode
	OSSR BDTR = 0x01 << 11 //+ Off-state selection for Run mode
	BKE  BDTR = 0x01 << 12 //+ Break enable
	BKP  BDTR = 0x01 << 13 //+ Break polarity
	AOE  BDTR = 0x01 << 14 //+ Automatic output enable
	MOE  BDTR = 0x01 << 15 //+ Main output enable
	BKF  BDTR = 0x0F << 16 //+ Break filter
)

const (
	DTGn  = 0
	LOCKn = 8
	OSSIn = 10
	OSSRn = 11
	BKEn  = 12
	BKPn  = 13
	AOEn  = 14
	MOEn  = 15
	BKFn  = 16
)

const (
	DBA DCR = 0x1F << 0 //+ DMA base address
	DBL DCR = 0x1F << 8 //+ DMA burst length
)

const (
	DBAn = 0
	DBLn = 8
)

const (
	DMAB DMAR = 0xFFFF << 0 //+ DMA register for burst accesses
)

const (
	DMABn = 0
)

const (
	BKINE     AF1 = 0x01 << 0  //+ BRK BKIN input enable
	BKCMP1E   AF1 = 0x01 << 1  //+ BRK COMP1 enable
	BKCMP2E   AF1 = 0x01 << 2  //+ BRK COMP2 enable
	BKDF1BK0E AF1 = 0x01 << 8  //+ BRK dfsdm1_break[0] enable
	BKINP     AF1 = 0x01 << 9  //+ BRK BKIN input polarity
	BKCMP1P   AF1 = 0x01 << 10 //+ BRK COMP1 input polarity
	BKCMP2P   AF1 = 0x01 << 11 //+ BRK COMP2 input polarity
)

const (
	BKINEn     = 0
	BKCMP1En   = 1
	BKCMP2En   = 2
	BKDF1BK0En = 8
	BKINPn     = 9
	BKCMP1Pn   = 10
	BKCMP2Pn   = 11
)

const (
	TI1SEL TISEL = 0x0F << 0 //+ selects TI1[0] to TI1[15] input
	TI2SEL TISEL = 0x0F << 8 //+ selects TI2[0] to TI2[15] input
)

const (
	TI1SELn = 0
	TI2SELn = 8
)
