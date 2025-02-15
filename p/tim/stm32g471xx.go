// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32g471xx

// Package tim provides access to the registers of the TIM peripheral.
//
// Instances:
//
//	TIM1   TIM1_BASE   APB2  TIM1_BRK_TIM15,TIM1_UP_TIM16,TIM1_TRG_COM,TIM1_CC,TIM8_CC  Advanced-timers
//	TIM2   TIM2_BASE   APB1  TIM2                                                       Advanced-timers
//	TIM3   TIM3_BASE   APB1  TIM3                                                       Advanced-timers
//	TIM4   TIM4_BASE   APB1  TIM4                                                       Advanced-timers
//	TIM5   TIM5_BASE   APB1  TIM5                                                       Advanced-timers
//	TIM6   TIM6_BASE   APB1  TIM6_DACUNDER*                                             Basic-timers
//	TIM7   TIM7_BASE   APB1  TIM7                                                       Basic-timers
//	TIM8   TIM8_BASE   APB2  TIM8_BRK,TIM8_UP,TIM8_TRG_COM                              Advanced-timers
//	TIM15  TIM15_BASE  APB2  -                                                          General purpose timers
//	TIM16  TIM16_BASE  APB2  -                                                          General purpose timers
//	TIM17  TIM17_BASE  APB2  -                                                          General purpose timers
//
// Registers:
//
//	0x000 32  CR1    control register 1
//	0x004 32  CR2    control register 2
//	0x008 32  SMCR   slave mode control register
//	0x00C 32  DIER   DMA/Interrupt enable register
//	0x010 32  SR     status register
//	0x014 32  EGR    event generation register
//	0x018 32  CCMR1  capture/compare mode register 1
//	0x01C 32  CCMR2  capture/compare mode register 2
//	0x020 32  CCER   capture/compare enable register
//	0x024 32  CNT    counter
//	0x028 32  PSC    prescaler
//	0x02C 32  ARR    auto-reload register
//	0x030 32  RCR    repetition counter register
//	0x034 32  CCR1   capture/compare register 1
//	0x038 32  CCR2   capture/compare register 2
//	0x03C 32  CCR3   capture/compare register 3
//	0x040 32  CCR4   capture/compare register 4
//	0x044 32  BDTR   break and dead-time register
//	0x048 32  CCR5   capture/compare register 4
//	0x04C 32  CCR6   capture/compare register 4
//	0x050 32  CCMR3  capture/compare mode register 2
//	0x054 32  DTR2   timer Deadtime Register 2
//	0x058 32  ECR    DMA control register
//	0x05C 32  TISEL  TIM timer input selection register
//	0x060 32  AF1    TIM alternate function option register 1
//	0x064 32  AF2    TIM alternate function option register 2
//	0x3DC 32  DCR    control register
//	0x3E0 32  DMAR   DMA address for full transfer
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package tim

const (
	CEN      CR1 = 0x01 << 0  //+ Counter enable
	UDIS     CR1 = 0x01 << 1  //+ Update disable
	URS      CR1 = 0x01 << 2  //+ Update request source
	OPM      CR1 = 0x01 << 3  //+ One-pulse mode
	DIR      CR1 = 0x01 << 4  //+ Direction
	CMS      CR1 = 0x03 << 5  //+ Center-aligned mode selection
	ARPE     CR1 = 0x01 << 7  //+ Auto-reload preload enable
	CKD      CR1 = 0x03 << 8  //+ Clock division
	UIFREMAP CR1 = 0x01 << 11 //+ UIF status bit remapping
	DITHEN   CR1 = 0x01 << 12 //+ Dithering Enable
)

const (
	CENn      = 0
	UDISn     = 1
	URSn      = 2
	OPMn      = 3
	DIRn      = 4
	CMSn      = 5
	ARPEn     = 7
	CKDn      = 8
	UIFREMAPn = 11
	DITHENn   = 12
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
	OIS2N CR2 = 0x01 << 11 //+ Output Idle state 2
	OIS3  CR2 = 0x01 << 12 //+ Output Idle state 3
	OIS3N CR2 = 0x01 << 13 //+ Output Idle state 3
	OIS4  CR2 = 0x01 << 14 //+ Output Idle state 4
	OIS4N CR2 = 0x01 << 15 //+ Output Idle state 4 (OC4N output)
	OIS5  CR2 = 0x01 << 16 //+ Output Idle state 5 (OC5 output)
	OIS6  CR2 = 0x01 << 18 //+ Output Idle state 6 (OC6 output)
	MMS2  CR2 = 0x0F << 20 //+ Master mode selection 2
	MMS_3 CR2 = 0x01 << 25 //+ Master mode selection - bit 3
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
	OIS2Nn = 11
	OIS3n  = 12
	OIS3Nn = 13
	OIS4n  = 14
	OIS4Nn = 15
	OIS5n  = 16
	OIS6n  = 18
	MMS2n  = 20
	MMS_3n = 25
)

const (
	SMS    SMCR = 0x07 << 0  //+ Slave mode selection
	OCCS   SMCR = 0x01 << 3  //+ OCREF clear selection
	TS     SMCR = 0x07 << 4  //+ Trigger selection
	MSM    SMCR = 0x01 << 7  //+ Master/Slave mode
	ETF    SMCR = 0x0F << 8  //+ External trigger filter
	ETPS   SMCR = 0x03 << 12 //+ External trigger prescaler
	ECE    SMCR = 0x01 << 14 //+ External clock enable
	ETP    SMCR = 0x01 << 15 //+ External trigger polarity
	SMS_3  SMCR = 0x01 << 16 //+ Slave mode selection - bit 3
	TS_4_3 SMCR = 0x03 << 20 //+ Trigger selection - bit 4:3
	SMSPE  SMCR = 0x01 << 24 //+ SMS Preload Enable
	SMSPS  SMCR = 0x01 << 25 //+ SMS Preload Source
)

const (
	SMSn    = 0
	OCCSn   = 3
	TSn     = 4
	MSMn    = 7
	ETFn    = 8
	ETPSn   = 12
	ECEn    = 14
	ETPn    = 15
	SMS_3n  = 16
	TS_4_3n = 20
	SMSPEn  = 24
	SMSPSn  = 25
)

const (
	UIE    DIER = 0x01 << 0  //+ Update interrupt enable
	CC1IE  DIER = 0x01 << 1  //+ Capture/Compare 1 interrupt enable
	CC2IE  DIER = 0x01 << 2  //+ Capture/Compare 2 interrupt enable
	CC3IE  DIER = 0x01 << 3  //+ Capture/Compare 3 interrupt enable
	CC4IE  DIER = 0x01 << 4  //+ Capture/Compare 4 interrupt enable
	COMIE  DIER = 0x01 << 5  //+ COM interrupt enable
	TIE    DIER = 0x01 << 6  //+ Trigger interrupt enable
	BIE    DIER = 0x01 << 7  //+ Break interrupt enable
	UDE    DIER = 0x01 << 8  //+ Update DMA request enable
	CC1DE  DIER = 0x01 << 9  //+ Capture/Compare 1 DMA request enable
	CC2DE  DIER = 0x01 << 10 //+ Capture/Compare 2 DMA request enable
	CC3DE  DIER = 0x01 << 11 //+ Capture/Compare 3 DMA request enable
	CC4DE  DIER = 0x01 << 12 //+ Capture/Compare 4 DMA request enable
	COMDE  DIER = 0x01 << 13 //+ COM DMA request enable
	TDE    DIER = 0x01 << 14 //+ Trigger DMA request enable
	IDXIE  DIER = 0x01 << 20 //+ Index interrupt enable
	DIRIE  DIER = 0x01 << 21 //+ Direction Change interrupt enable
	IERRIE DIER = 0x01 << 22 //+ Index Error interrupt enable
	TERRIE DIER = 0x01 << 23 //+ Transition Error interrupt enable
)

const (
	UIEn    = 0
	CC1IEn  = 1
	CC2IEn  = 2
	CC3IEn  = 3
	CC4IEn  = 4
	COMIEn  = 5
	TIEn    = 6
	BIEn    = 7
	UDEn    = 8
	CC1DEn  = 9
	CC2DEn  = 10
	CC3DEn  = 11
	CC4DEn  = 12
	COMDEn  = 13
	TDEn    = 14
	IDXIEn  = 20
	DIRIEn  = 21
	IERRIEn = 22
	TERRIEn = 23
)

const (
	UIF   SR = 0x01 << 0  //+ Update interrupt flag
	CC1IF SR = 0x01 << 1  //+ Capture/compare 1 interrupt flag
	CC2IF SR = 0x01 << 2  //+ Capture/Compare 2 interrupt flag
	CC3IF SR = 0x01 << 3  //+ Capture/Compare 3 interrupt flag
	CC4IF SR = 0x01 << 4  //+ Capture/Compare 4 interrupt flag
	COMIF SR = 0x01 << 5  //+ COM interrupt flag
	TIF   SR = 0x01 << 6  //+ Trigger interrupt flag
	BIF   SR = 0x01 << 7  //+ Break interrupt flag
	B2IF  SR = 0x01 << 8  //+ Break 2 interrupt flag
	CC1OF SR = 0x01 << 9  //+ Capture/Compare 1 overcapture flag
	CC2OF SR = 0x01 << 10 //+ Capture/compare 2 overcapture flag
	CC3OF SR = 0x01 << 11 //+ Capture/Compare 3 overcapture flag
	CC4OF SR = 0x01 << 12 //+ Capture/Compare 4 overcapture flag
	SBIF  SR = 0x01 << 13 //+ System Break interrupt flag
	CC5IF SR = 0x01 << 16 //+ Compare 5 interrupt flag
	CC6IF SR = 0x01 << 17 //+ Compare 6 interrupt flag
	IDXF  SR = 0x01 << 20 //+ Index interrupt flag
	DIRF  SR = 0x01 << 21 //+ Direction Change interrupt flag
	IERRF SR = 0x01 << 22 //+ Index Error interrupt flag
	TERRF SR = 0x01 << 23 //+ Transition Error interrupt flag
)

const (
	UIFn   = 0
	CC1IFn = 1
	CC2IFn = 2
	CC3IFn = 3
	CC4IFn = 4
	COMIFn = 5
	TIFn   = 6
	BIFn   = 7
	B2IFn  = 8
	CC1OFn = 9
	CC2OFn = 10
	CC3OFn = 11
	CC4OFn = 12
	SBIFn  = 13
	CC5IFn = 16
	CC6IFn = 17
	IDXFn  = 20
	DIRFn  = 21
	IERRFn = 22
	TERRFn = 23
)

const (
	UG   EGR = 0x01 << 0 //+ Update generation
	CC1G EGR = 0x01 << 1 //+ Capture/compare 1 generation
	CC2G EGR = 0x01 << 2 //+ Capture/compare 2 generation
	CC3G EGR = 0x01 << 3 //+ Capture/compare 3 generation
	CC4G EGR = 0x01 << 4 //+ Capture/compare 4 generation
	COMG EGR = 0x01 << 5 //+ Capture/Compare control update generation
	TG   EGR = 0x01 << 6 //+ Trigger generation
	BG   EGR = 0x01 << 7 //+ Break generation
	B2G  EGR = 0x01 << 8 //+ Break 2 generation
)

const (
	UGn   = 0
	CC1Gn = 1
	CC2Gn = 2
	CC3Gn = 3
	CC4Gn = 4
	COMGn = 5
	TGn   = 6
	BGn   = 7
	B2Gn  = 8
)

const (
	CC1S   CCMR1 = 0x03 << 0  //+ Capture/Compare 1 selection
	OC1FE  CCMR1 = 0x01 << 2  //+ Output Compare 1 fast enable
	OC1PE  CCMR1 = 0x01 << 3  //+ Output Compare 1 preload enable
	OC1M   CCMR1 = 0x07 << 4  //+ Output Compare 1 mode
	OC1CE  CCMR1 = 0x01 << 7  //+ Output Compare 1 clear enable
	CC2S   CCMR1 = 0x03 << 8  //+ Capture/Compare 2 selection
	OC2FE  CCMR1 = 0x01 << 10 //+ Output Compare 2 fast enable
	OC2PE  CCMR1 = 0x01 << 11 //+ Output Compare 2 preload enable
	OC2M   CCMR1 = 0x07 << 12 //+ Output Compare 2 mode
	OC2CE  CCMR1 = 0x01 << 15 //+ Output Compare 2 clear enable
	OC1M_3 CCMR1 = 0x01 << 16 //+ Output Compare 1 mode - bit 3
	OC2M_3 CCMR1 = 0x01 << 24 //+ Output Compare 2 mode - bit 3
	ICPCS  CCMR1 = 0x03 << 2  //+ Input capture 1 prescaler
	IC1F   CCMR1 = 0x0F << 4  //+ Input capture 1 filter
	IC2PSC CCMR1 = 0x03 << 10 //+ Input capture 2 prescaler
	IC2F   CCMR1 = 0x0F << 12 //+ Input capture 2 filter
)

const (
	CC1Sn   = 0
	OC1FEn  = 2
	OC1PEn  = 3
	OC1Mn   = 4
	OC1CEn  = 7
	CC2Sn   = 8
	OC2FEn  = 10
	OC2PEn  = 11
	OC2Mn   = 12
	OC2CEn  = 15
	OC1M_3n = 16
	OC2M_3n = 24
	ICPCSn  = 2
	IC1Fn   = 4
	IC2PSCn = 10
	IC2Fn   = 12
)

const (
	CC3S   CCMR2 = 0x03 << 0  //+ Capture/Compare 3 selection
	OC3FE  CCMR2 = 0x01 << 2  //+ Output compare 3 fast enable
	OC3PE  CCMR2 = 0x01 << 3  //+ Output compare 3 preload enable
	OC3M   CCMR2 = 0x07 << 4  //+ Output compare 3 mode
	OC3CE  CCMR2 = 0x01 << 7  //+ Output compare 3 clear enable
	CC4S   CCMR2 = 0x03 << 8  //+ Capture/Compare 4 selection
	OC4FE  CCMR2 = 0x01 << 10 //+ Output compare 4 fast enable
	OC4PE  CCMR2 = 0x01 << 11 //+ Output compare 4 preload enable
	OC4M   CCMR2 = 0x07 << 12 //+ Output compare 4 mode
	OC4CE  CCMR2 = 0x01 << 15 //+ Output compare 4 clear enable
	OC3M_3 CCMR2 = 0x01 << 16 //+ Output Compare 3 mode - bit 3
	OC4M_3 CCMR2 = 0x01 << 24 //+ Output Compare 4 mode - bit 3
	IC3PSC CCMR2 = 0x03 << 2  //+ Input capture 3 prescaler
	IC3F   CCMR2 = 0x0F << 4  //+ Input capture 3 filter
	IC4PSC CCMR2 = 0x03 << 10 //+ Input capture 4 prescaler
	IC4F   CCMR2 = 0x0F << 12 //+ Input capture 4 filter
)

const (
	CC3Sn   = 0
	OC3FEn  = 2
	OC3PEn  = 3
	OC3Mn   = 4
	OC3CEn  = 7
	CC4Sn   = 8
	OC4FEn  = 10
	OC4PEn  = 11
	OC4Mn   = 12
	OC4CEn  = 15
	OC3M_3n = 16
	OC4M_3n = 24
	IC3PSCn = 2
	IC3Fn   = 4
	IC4PSCn = 10
	IC4Fn   = 12
)

const (
	CC1E  CCER = 0x01 << 0  //+ Capture/Compare 1 output enable
	CC1P  CCER = 0x01 << 1  //+ Capture/Compare 1 output Polarity
	CC1NE CCER = 0x01 << 2  //+ Capture/Compare 1 complementary output enable
	CC1NP CCER = 0x01 << 3  //+ Capture/Compare 1 output Polarity
	CC2E  CCER = 0x01 << 4  //+ Capture/Compare 2 output enable
	CC2P  CCER = 0x01 << 5  //+ Capture/Compare 2 output Polarity
	CC2NE CCER = 0x01 << 6  //+ Capture/Compare 2 complementary output enable
	CC2NP CCER = 0x01 << 7  //+ Capture/Compare 2 output Polarity
	CC3E  CCER = 0x01 << 8  //+ Capture/Compare 3 output enable
	CC3P  CCER = 0x01 << 9  //+ Capture/Compare 3 output Polarity
	CC3NE CCER = 0x01 << 10 //+ Capture/Compare 3 complementary output enable
	CC3NP CCER = 0x01 << 11 //+ Capture/Compare 3 output Polarity
	CC4E  CCER = 0x01 << 12 //+ Capture/Compare 4 output enable
	CC4P  CCER = 0x01 << 13 //+ Capture/Compare 3 output Polarity
	CC4NE CCER = 0x01 << 14 //+ Capture/Compare 4 complementary output enable
	CC4NP CCER = 0x01 << 15 //+ Capture/Compare 4 complementary output polarity
	CC5E  CCER = 0x01 << 16 //+ Capture/Compare 5 output enable
	CC5P  CCER = 0x01 << 17 //+ Capture/Compare 5 output polarity
	CC6E  CCER = 0x01 << 20 //+ Capture/Compare 6 output enable
	CC6P  CCER = 0x01 << 21 //+ Capture/Compare 6 output polarity
)

const (
	CC1En  = 0
	CC1Pn  = 1
	CC1NEn = 2
	CC1NPn = 3
	CC2En  = 4
	CC2Pn  = 5
	CC2NEn = 6
	CC2NPn = 7
	CC3En  = 8
	CC3Pn  = 9
	CC3NEn = 10
	CC3NPn = 11
	CC4En  = 12
	CC4Pn  = 13
	CC4NEn = 14
	CC4NPn = 15
	CC5En  = 16
	CC5Pn  = 17
	CC6En  = 20
	CC6Pn  = 21
)

const (
	DTG     BDTR = 0xFF << 0  //+ Dead-time generator setup
	LOCK    BDTR = 0x03 << 8  //+ Lock configuration
	OSSI    BDTR = 0x01 << 10 //+ Off-state selection for Idle mode
	OSSR    BDTR = 0x01 << 11 //+ Off-state selection for Run mode
	BKE     BDTR = 0x01 << 12 //+ Break enable
	BKP     BDTR = 0x01 << 13 //+ Break polarity
	AOE     BDTR = 0x01 << 14 //+ Automatic output enable
	MOE     BDTR = 0x01 << 15 //+ Main output enable
	BKF     BDTR = 0x0F << 16 //+ Break filter
	BK2F    BDTR = 0x0F << 20 //+ Break 2 filter
	BK2E    BDTR = 0x01 << 24 //+ Break 2 Enable
	BK2P    BDTR = 0x01 << 25 //+ Break 2 polarity
	BKDSRM  BDTR = 0x01 << 26 //+ BKDSRM
	BK2DSRM BDTR = 0x01 << 27 //+ BK2DSRM
	BKBID   BDTR = 0x01 << 28 //+ BKBID
	BK2ID   BDTR = 0x01 << 29 //+ BK2ID
)

const (
	DTGn     = 0
	LOCKn    = 8
	OSSIn    = 10
	OSSRn    = 11
	BKEn     = 12
	BKPn     = 13
	AOEn     = 14
	MOEn     = 15
	BKFn     = 16
	BK2Fn    = 20
	BK2En    = 24
	BK2Pn    = 25
	BKDSRMn  = 26
	BK2DSRMn = 27
	BKBIDn   = 28
	BK2IDn   = 29
)

const (
	OC5FE     CCMR3 = 0x01 << 2  //+ Output compare 5 fast enable
	OC5PE     CCMR3 = 0x01 << 3  //+ Output compare 5 preload enable
	OC5M      CCMR3 = 0x07 << 4  //+ Output compare 5 mode
	OC5CE     CCMR3 = 0x01 << 7  //+ Output compare 5 clear enable
	OC6FE     CCMR3 = 0x01 << 10 //+ Output compare 6 fast enable
	OC6PE     CCMR3 = 0x01 << 11 //+ Output compare 6 preload enable
	OC6M      CCMR3 = 0x07 << 12 //+ Output compare 6 mode
	OC6CE     CCMR3 = 0x01 << 15 //+ Output compare 6 clear enable
	OC5M_bit3 CCMR3 = 0x07 << 16 //+ Output Compare 5 mode bit 3
	OC6M_bit3 CCMR3 = 0x01 << 24 //+ Output Compare 6 mode bit 3
)

const (
	OC5FEn     = 2
	OC5PEn     = 3
	OC5Mn      = 4
	OC5CEn     = 7
	OC6FEn     = 10
	OC6PEn     = 11
	OC6Mn      = 12
	OC6CEn     = 15
	OC5M_bit3n = 16
	OC6M_bit3n = 24
)

const (
	DTGF DTR2 = 0xFF << 0  //+ Dead-time falling edge generator setup
	DTAE DTR2 = 0x01 << 16 //+ Deadtime Asymmetric Enable
	DTPE DTR2 = 0x01 << 17 //+ Deadtime Preload Enable
)

const (
	DTGFn = 0
	DTAEn = 16
	DTPEn = 17
)

const (
	IE     ECR = 0x01 << 0  //+ Index Enable
	IDIR   ECR = 0x03 << 1  //+ Index Direction
	IBLK   ECR = 0x03 << 3  //+ Index Blanking
	FIDX   ECR = 0x01 << 5  //+ First Index
	IPOS   ECR = 0x03 << 6  //+ Index Positioning
	PW     ECR = 0xFF << 16 //+ Pulse width
	PWPRSC ECR = 0x07 << 24 //+ Pulse Width prescaler
)

const (
	IEn     = 0
	IDIRn   = 1
	IBLKn   = 3
	FIDXn   = 5
	IPOSn   = 6
	PWn     = 16
	PWPRSCn = 24
)

const (
	TI1SEL TISEL = 0x0F << 0  //+ TI1[0] to TI1[15] input selection
	TI2SEL TISEL = 0x0F << 8  //+ TI2[0] to TI2[15] input selection
	TI3SEL TISEL = 0x0F << 16 //+ TI3[0] to TI3[15] input selection
	TI4SEL TISEL = 0x0F << 24 //+ TI4[0] to TI4[15] input selection
)

const (
	TI1SELn = 0
	TI2SELn = 8
	TI3SELn = 16
	TI4SELn = 24
)

const (
	BKINE   AF1 = 0x01 << 0  //+ BRK BKIN input enable
	BKCMP1E AF1 = 0x01 << 1  //+ BRK COMP1 enable
	BKCMP2E AF1 = 0x01 << 2  //+ BRK COMP2 enable
	BKCMP3E AF1 = 0x01 << 3  //+ BRK COMP3 enable
	BKCMP4E AF1 = 0x01 << 4  //+ BRK COMP4 enable
	BKCMP5E AF1 = 0x01 << 5  //+ BRK COMP5 enable
	BKCMP6E AF1 = 0x01 << 6  //+ BRK COMP6 enable
	BKCMP7E AF1 = 0x01 << 7  //+ BRK COMP7 enable
	BKINP   AF1 = 0x01 << 9  //+ BRK BKIN input polarity
	BKCMP1P AF1 = 0x01 << 10 //+ BRK COMP1 input polarity
	BKCMP2P AF1 = 0x01 << 11 //+ BRK COMP2 input polarity
	BKCMP3P AF1 = 0x01 << 12 //+ BRK COMP3 input polarity
	BKCMP4P AF1 = 0x01 << 13 //+ BRK COMP4 input polarity
	ETRSEL  AF1 = 0x0F << 14 //+ ETR source selection
)

const (
	BKINEn   = 0
	BKCMP1En = 1
	BKCMP2En = 2
	BKCMP3En = 3
	BKCMP4En = 4
	BKCMP5En = 5
	BKCMP6En = 6
	BKCMP7En = 7
	BKINPn   = 9
	BKCMP1Pn = 10
	BKCMP2Pn = 11
	BKCMP3Pn = 12
	BKCMP4Pn = 13
	ETRSELn  = 14
)

const (
	BKINE    AF2 = 0x01 << 0  //+ BRK BKIN input enable
	BK2CMP1E AF2 = 0x01 << 1  //+ BRK2 COMP1 enable
	BK2CMP2E AF2 = 0x01 << 2  //+ BRK2 COMP2 enable
	BK2CMP3E AF2 = 0x01 << 3  //+ BRK2 COMP3 enable
	BK2CMP4E AF2 = 0x01 << 4  //+ BRK2 COMP4 enable
	BK2CMP5E AF2 = 0x01 << 5  //+ BRK2 COMP5 enable
	BK2CMP6E AF2 = 0x01 << 6  //+ BRK2 COMP6 enable
	BK2CMP7E AF2 = 0x01 << 7  //+ BRK2 COMP7 enable
	BK2INP   AF2 = 0x01 << 9  //+ BRK2 BKIN input polarity
	BK2CMP1P AF2 = 0x01 << 10 //+ BRK2 COMP1 input polarity
	BK2CMP2P AF2 = 0x01 << 11 //+ BRK2 COMP2 input polarity
	BK2CMP3P AF2 = 0x01 << 12 //+ BRK2 COMP3 input polarity
	BK2CMP4P AF2 = 0x01 << 13 //+ BRK2 COMP4 input polarity
	OCRSEL   AF2 = 0x07 << 16 //+ OCREF_CLR source selection
)

const (
	BKINEn    = 0
	BK2CMP1En = 1
	BK2CMP2En = 2
	BK2CMP3En = 3
	BK2CMP4En = 4
	BK2CMP5En = 5
	BK2CMP6En = 6
	BK2CMP7En = 7
	BK2INPn   = 9
	BK2CMP1Pn = 10
	BK2CMP2Pn = 11
	BK2CMP3Pn = 12
	BK2CMP4Pn = 13
	OCRSELn   = 16
)

const (
	DBA DCR = 0x1F << 0 //+ DMA base address
	DBL DCR = 0x1F << 8 //+ DMA burst length
)

const (
	DBAn = 0
	DBLn = 8
)
