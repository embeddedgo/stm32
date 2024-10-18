// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package dfsdm provides access to the registers of the DFSDM peripheral.
//
// Instances:
//
//	DFSDM  DFSDM_BASE  -  DFSDM1_FLT0,DFSDM1_FLT1,DFSDM1_FLT2,DFSDM1_FLT3  Digital filter for sigma delta modulators
//
// Registers:
//
//	0x000 32  CH0CFGR1           channel configuration y register
//	0x004 32  CH0CFGR2           channel configuration y register
//	0x008 32  CH0AWSCDR          analog watchdog and short-circuit detector register
//	0x00C 32  CH0WDATR           channel watchdog filter data register
//	0x010 32  CH0DATINR          channel data input register
//	0x014 32  CH0DLYR            channel y delay register
//	0x020 32  CH1CFGR1           CH1CFGR1
//	0x024 32  CH1CFGR2           CH1CFGR2
//	0x028 32  CH1AWSCDR          CH1AWSCDR
//	0x02C 32  CH1WDATR           CH1WDATR
//	0x030 32  CH1DATINR          CH1DATINR
//	0x034 32  CH1DLYR            channel y delay register
//	0x040 32  CH2CFGR1           CH2CFGR1
//	0x044 32  CH2CFGR2           CH2CFGR2
//	0x048 32  CH2AWSCDR          CH2AWSCDR
//	0x04C 32  CH2WDATR           CH2WDATR
//	0x050 32  CH2DATINR          CH2DATINR
//	0x054 32  CH2DLYR            channel y delay register
//	0x060 32  CH3CFGR1           CH3CFGR1
//	0x064 32  CH3CFGR2           CH3CFGR2
//	0x068 32  CH3AWSCDR          CH3AWSCDR
//	0x06C 32  CH3WDATR           CH3WDATR
//	0x070 32  CH3DATINR          CH3DATINR
//	0x074 32  CH3DLYR            channel y delay register
//	0x080 32  CH4CFGR1           CH4CFGR1
//	0x084 32  CH4CFGR2           CH4CFGR2
//	0x088 32  CH4AWSCDR          CH4AWSCDR
//	0x08C 32  CH4WDATR           CH4WDATR
//	0x090 32  CH4DATINR          CH4DATINR
//	0x094 32  CH4DLYR            channel y delay register
//	0x0A0 32  CH5CFGR1           CH5CFGR1
//	0x0A4 32  CH5CFGR2           CH5CFGR2
//	0x0A8 32  CH5AWSCDR          CH5AWSCDR
//	0x0AC 32  CH5WDATR           CH5WDATR
//	0x0B0 32  CH5DATINR          CH5DATINR
//	0x0B4 32  CH5DLYR            channel y delay register
//	0x0C0 32  CH6CFGR1           CH6CFGR1
//	0x0C4 32  CH6CFGR2           CH6CFGR2
//	0x0C8 32  CH6AWSCDR          CH6AWSCDR
//	0x0CC 32  CH6WDATR           CH6WDATR
//	0x0D0 32  CH6DATINR          CH6DATINR
//	0x0D4 32  CH6DLYR            channel y delay register
//	0x0E0 32  CH7CFGR1           CH7CFGR1
//	0x0E4 32  CH7CFGR2           CH7CFGR2
//	0x0E8 32  CH7AWSCDR          CH7AWSCDR
//	0x0EC 32  CH7WDATR           CH7WDATR
//	0x0F0 32  CH7DATINR          CH7DATINR
//	0x0F4 32  CH7DLYR            channel y delay register
//	0x100 32  DFSDM_FLT0CR1      control register 1
//	0x104 32  DFSDM_FLT0CR2      control register 2
//	0x108 32  DFSDM_FLT0ISR      interrupt and status register
//	0x10C 32  DFSDM_FLT0ICR      interrupt flag clear register
//	0x110 32  DFSDM_FLT0JCHGR    injected channel group selection register
//	0x114 32  DFSDM_FLT0FCR      filter control register
//	0x118 32  DFSDM_FLT0JDATAR   data register for injected group
//	0x11C 32  DFSDM_FLT0RDATAR   data register for the regular channel
//	0x120 32  DFSDM_FLT0AWHTR    analog watchdog high threshold register
//	0x124 32  DFSDM_FLT0AWLTR    analog watchdog low threshold register
//	0x128 32  DFSDM_FLT0AWSR     analog watchdog status register
//	0x12C 32  DFSDM_FLT0AWCFR    analog watchdog clear flag register
//	0x130 32  DFSDM_FLT0EXMAX    Extremes detector maximum register
//	0x134 32  DFSDM_FLT0EXMIN    Extremes detector minimum register
//	0x138 32  DFSDM_FLT0CNVTIMR  conversion timer register
//	0x180 32  DFSDM_FLT1CR1      control register 1
//	0x184 32  DFSDM_FLT1CR2      control register 2
//	0x188 32  DFSDM_FLT1ISR      interrupt and status register
//	0x18C 32  DFSDM_FLT1ICR      interrupt flag clear register
//	0x190 32  DFSDM_FLT1CHGR     injected channel group selection register
//	0x194 32  DFSDM_FLT1FCR      filter control register
//	0x198 32  DFSDM_FLT1JDATAR   data register for injected group
//	0x19C 32  DFSDM_FLT1RDATAR   data register for the regular channel
//	0x1A0 32  DFSDM_FLT1AWHTR    analog watchdog high threshold register
//	0x1A4 32  DFSDM_FLT1AWLTR    analog watchdog low threshold register
//	0x1A8 32  DFSDM_FLT1AWSR     analog watchdog status register
//	0x1AC 32  DFSDM_FLT1AWCFR    analog watchdog clear flag register
//	0x1B0 32  DFSDM_FLT1EXMAX    Extremes detector maximum register
//	0x1B4 32  DFSDM_FLT1EXMIN    Extremes detector minimum register
//	0x1B8 32  DFSDM_FLT1CNVTIMR  conversion timer register
//	0x200 32  DFSDM_FLT2CR1      control register 1
//	0x204 32  DFSDM_FLT2CR2      control register 2
//	0x208 32  DFSDM_FLT2ISR      interrupt and status register
//	0x20C 32  DFSDM_FLT2ICR      interrupt flag clear register
//	0x210 32  DFSDM_FLT2JCHGR    injected channel group selection register
//	0x214 32  DFSDM_FLT2FCR      filter control register
//	0x218 32  DFSDM_FLT2JDATAR   data register for injected group
//	0x21C 32  DFSDM_FLT2RDATAR   data register for the regular channel
//	0x220 32  DFSDM_FLT2AWHTR    analog watchdog high threshold register
//	0x224 32  DFSDM_FLT2AWLTR    analog watchdog low threshold register
//	0x228 32  DFSDM_FLT2AWSR     analog watchdog status register
//	0x22C 32  DFSDM_FLT2AWCFR    analog watchdog clear flag register
//	0x230 32  DFSDM_FLT2EXMAX    Extremes detector maximum register
//	0x234 32  DFSDM_FLT2EXMIN    Extremes detector minimum register
//	0x238 32  DFSDM_FLT2CNVTIMR  conversion timer register
//	0x280 32  DFSDM_FLT3CR1      control register 1
//	0x284 32  DFSDM_FLT3CR2      control register 2
//	0x288 32  DFSDM_FLT3ISR      interrupt and status register
//	0x28C 32  DFSDM_FLT3ICR      interrupt flag clear register
//	0x290 32  DFSDM_FLT3JCHGR    injected channel group selection register
//	0x294 32  DFSDM_FLT3FCR      filter control register
//	0x298 32  DFSDM_FLT3JDATAR   data register for injected group
//	0x29C 32  DFSDM_FLT3RDATAR   data register for the regular channel
//	0x2A0 32  DFSDM_FLT3AWHTR    analog watchdog high threshold register
//	0x2A4 32  DFSDM_FLT3AWLTR    analog watchdog low threshold register
//	0x2A8 32  DFSDM_FLT3AWSR     analog watchdog status register
//	0x2AC 32  DFSDM_FLT3AWCFR    analog watchdog clear flag register
//	0x2B0 32  DFSDM_FLT3EXMAX    Extremes detector maximum register
//	0x2B4 32  DFSDM_FLT3EXMIN    Extremes detector minimum register
//	0x2B8 32  DFSDM_FLT3CNVTIMR  conversion timer register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package dfsdm

const (
	SITP     CH0CFGR1 = 0x03 << 0  //+ SITP
	SPICKSEL CH0CFGR1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CH0CFGR1 = 0x01 << 5  //+ SCDEN
	CKABEN   CH0CFGR1 = 0x01 << 6  //+ CKABEN
	CHEN     CH0CFGR1 = 0x01 << 7  //+ CHEN
	CHINSEL  CH0CFGR1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CH0CFGR1 = 0x03 << 12 //+ DATMPX
	DATPACK  CH0CFGR1 = 0x03 << 14 //+ DATPACK
	CKOUTDIV CH0CFGR1 = 0xFF << 16 //+ CKOUTDIV
	CKOUTSRC CH0CFGR1 = 0x01 << 30 //+ CKOUTSRC
	DFSDMEN  CH0CFGR1 = 0x01 << 31 //+ DFSDMEN
)

const (
	SITPn     = 0
	SPICKSELn = 2
	SCDENn    = 5
	CKABENn   = 6
	CHENn     = 7
	CHINSELn  = 8
	DATMPXn   = 12
	DATPACKn  = 14
	CKOUTDIVn = 16
	CKOUTSRCn = 30
	DFSDMENn  = 31
)

const (
	DTRBS  CH0CFGR2 = 0x1F << 3     //+ DTRBS
	OFFSET CH0CFGR2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   CH0AWSCDR = 0xFF << 0  //+ SCDT
	BKSCD  CH0AWSCDR = 0x0F << 12 //+ BKSCD
	AWFOSR CH0AWSCDR = 0x1F << 16 //+ AWFOSR
	AWFORD CH0AWSCDR = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CH0WDATR = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CH0DATINR = 0xFFFF << 0  //+ INDAT0
	INDAT1 CH0DATINR = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	PLSSKP CH0DLYR = 0x3F << 0 //+ PLSSKP
)

const (
	PLSSKPn = 0
)

const (
	SITP     CH1CFGR1 = 0x03 << 0  //+ SITP
	SPICKSEL CH1CFGR1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CH1CFGR1 = 0x01 << 5  //+ SCDEN
	CKABEN   CH1CFGR1 = 0x01 << 6  //+ CKABEN
	CHEN     CH1CFGR1 = 0x01 << 7  //+ CHEN
	CHINSEL  CH1CFGR1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CH1CFGR1 = 0x03 << 12 //+ DATMPX
	DATPACK  CH1CFGR1 = 0x03 << 14 //+ DATPACK
)

const (
	SITPn     = 0
	SPICKSELn = 2
	SCDENn    = 5
	CKABENn   = 6
	CHENn     = 7
	CHINSELn  = 8
	DATMPXn   = 12
	DATPACKn  = 14
)

const (
	DTRBS  CH1CFGR2 = 0x1F << 3     //+ DTRBS
	OFFSET CH1CFGR2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   CH1AWSCDR = 0xFF << 0  //+ SCDT
	BKSCD  CH1AWSCDR = 0x0F << 12 //+ BKSCD
	AWFOSR CH1AWSCDR = 0x1F << 16 //+ AWFOSR
	AWFORD CH1AWSCDR = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CH1WDATR = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CH1DATINR = 0xFFFF << 0  //+ INDAT0
	INDAT1 CH1DATINR = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	PLSSKP CH1DLYR = 0x3F << 0 //+ PLSSKP
)

const (
	PLSSKPn = 0
)

const (
	SITP     CH2CFGR1 = 0x03 << 0  //+ SITP
	SPICKSEL CH2CFGR1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CH2CFGR1 = 0x01 << 5  //+ SCDEN
	CKABEN   CH2CFGR1 = 0x01 << 6  //+ CKABEN
	CHEN     CH2CFGR1 = 0x01 << 7  //+ CHEN
	CHINSEL  CH2CFGR1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CH2CFGR1 = 0x03 << 12 //+ DATMPX
	DATPACK  CH2CFGR1 = 0x03 << 14 //+ DATPACK
)

const (
	SITPn     = 0
	SPICKSELn = 2
	SCDENn    = 5
	CKABENn   = 6
	CHENn     = 7
	CHINSELn  = 8
	DATMPXn   = 12
	DATPACKn  = 14
)

const (
	DTRBS  CH2CFGR2 = 0x1F << 3     //+ DTRBS
	OFFSET CH2CFGR2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   CH2AWSCDR = 0xFF << 0  //+ SCDT
	BKSCD  CH2AWSCDR = 0x0F << 12 //+ BKSCD
	AWFOSR CH2AWSCDR = 0x1F << 16 //+ AWFOSR
	AWFORD CH2AWSCDR = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CH2WDATR = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CH2DATINR = 0xFFFF << 0  //+ INDAT0
	INDAT1 CH2DATINR = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	PLSSKP CH2DLYR = 0x3F << 0 //+ PLSSKP
)

const (
	PLSSKPn = 0
)

const (
	SITP     CH3CFGR1 = 0x03 << 0  //+ SITP
	SPICKSEL CH3CFGR1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CH3CFGR1 = 0x01 << 5  //+ SCDEN
	CKABEN   CH3CFGR1 = 0x01 << 6  //+ CKABEN
	CHEN     CH3CFGR1 = 0x01 << 7  //+ CHEN
	CHINSEL  CH3CFGR1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CH3CFGR1 = 0x03 << 12 //+ DATMPX
	DATPACK  CH3CFGR1 = 0x03 << 14 //+ DATPACK
)

const (
	SITPn     = 0
	SPICKSELn = 2
	SCDENn    = 5
	CKABENn   = 6
	CHENn     = 7
	CHINSELn  = 8
	DATMPXn   = 12
	DATPACKn  = 14
)

const (
	DTRBS  CH3CFGR2 = 0x1F << 3     //+ DTRBS
	OFFSET CH3CFGR2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   CH3AWSCDR = 0xFF << 0  //+ SCDT
	BKSCD  CH3AWSCDR = 0x0F << 12 //+ BKSCD
	AWFOSR CH3AWSCDR = 0x1F << 16 //+ AWFOSR
	AWFORD CH3AWSCDR = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CH3WDATR = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CH3DATINR = 0xFFFF << 0  //+ INDAT0
	INDAT1 CH3DATINR = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	PLSSKP CH3DLYR = 0x3F << 0 //+ PLSSKP
)

const (
	PLSSKPn = 0
)

const (
	SITP     CH4CFGR1 = 0x03 << 0  //+ SITP
	SPICKSEL CH4CFGR1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CH4CFGR1 = 0x01 << 5  //+ SCDEN
	CKABEN   CH4CFGR1 = 0x01 << 6  //+ CKABEN
	CHEN     CH4CFGR1 = 0x01 << 7  //+ CHEN
	CHINSEL  CH4CFGR1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CH4CFGR1 = 0x03 << 12 //+ DATMPX
	DATPACK  CH4CFGR1 = 0x03 << 14 //+ DATPACK
)

const (
	SITPn     = 0
	SPICKSELn = 2
	SCDENn    = 5
	CKABENn   = 6
	CHENn     = 7
	CHINSELn  = 8
	DATMPXn   = 12
	DATPACKn  = 14
)

const (
	DTRBS  CH4CFGR2 = 0x1F << 3     //+ DTRBS
	OFFSET CH4CFGR2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   CH4AWSCDR = 0xFF << 0  //+ SCDT
	BKSCD  CH4AWSCDR = 0x0F << 12 //+ BKSCD
	AWFOSR CH4AWSCDR = 0x1F << 16 //+ AWFOSR
	AWFORD CH4AWSCDR = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CH4WDATR = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CH4DATINR = 0xFFFF << 0  //+ INDAT0
	INDAT1 CH4DATINR = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	PLSSKP CH4DLYR = 0x3F << 0 //+ PLSSKP
)

const (
	PLSSKPn = 0
)

const (
	SITP     CH5CFGR1 = 0x03 << 0  //+ SITP
	SPICKSEL CH5CFGR1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CH5CFGR1 = 0x01 << 5  //+ SCDEN
	CKABEN   CH5CFGR1 = 0x01 << 6  //+ CKABEN
	CHEN     CH5CFGR1 = 0x01 << 7  //+ CHEN
	CHINSEL  CH5CFGR1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CH5CFGR1 = 0x03 << 12 //+ DATMPX
	DATPACK  CH5CFGR1 = 0x03 << 14 //+ DATPACK
)

const (
	SITPn     = 0
	SPICKSELn = 2
	SCDENn    = 5
	CKABENn   = 6
	CHENn     = 7
	CHINSELn  = 8
	DATMPXn   = 12
	DATPACKn  = 14
)

const (
	DTRBS  CH5CFGR2 = 0x1F << 3     //+ DTRBS
	OFFSET CH5CFGR2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   CH5AWSCDR = 0xFF << 0  //+ SCDT
	BKSCD  CH5AWSCDR = 0x0F << 12 //+ BKSCD
	AWFOSR CH5AWSCDR = 0x1F << 16 //+ AWFOSR
	AWFORD CH5AWSCDR = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CH5WDATR = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CH5DATINR = 0xFFFF << 0  //+ INDAT0
	INDAT1 CH5DATINR = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	PLSSKP CH5DLYR = 0x3F << 0 //+ PLSSKP
)

const (
	PLSSKPn = 0
)

const (
	SITP     CH6CFGR1 = 0x03 << 0  //+ SITP
	SPICKSEL CH6CFGR1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CH6CFGR1 = 0x01 << 5  //+ SCDEN
	CKABEN   CH6CFGR1 = 0x01 << 6  //+ CKABEN
	CHEN     CH6CFGR1 = 0x01 << 7  //+ CHEN
	CHINSEL  CH6CFGR1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CH6CFGR1 = 0x03 << 12 //+ DATMPX
	DATPACK  CH6CFGR1 = 0x03 << 14 //+ DATPACK
)

const (
	SITPn     = 0
	SPICKSELn = 2
	SCDENn    = 5
	CKABENn   = 6
	CHENn     = 7
	CHINSELn  = 8
	DATMPXn   = 12
	DATPACKn  = 14
)

const (
	DTRBS  CH6CFGR2 = 0x1F << 3     //+ DTRBS
	OFFSET CH6CFGR2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   CH6AWSCDR = 0xFF << 0  //+ SCDT
	BKSCD  CH6AWSCDR = 0x0F << 12 //+ BKSCD
	AWFOSR CH6AWSCDR = 0x1F << 16 //+ AWFOSR
	AWFORD CH6AWSCDR = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CH6WDATR = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CH6DATINR = 0xFFFF << 0  //+ INDAT0
	INDAT1 CH6DATINR = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	PLSSKP CH6DLYR = 0x3F << 0 //+ PLSSKP
)

const (
	PLSSKPn = 0
)

const (
	SITP     CH7CFGR1 = 0x03 << 0  //+ SITP
	SPICKSEL CH7CFGR1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CH7CFGR1 = 0x01 << 5  //+ SCDEN
	CKABEN   CH7CFGR1 = 0x01 << 6  //+ CKABEN
	CHEN     CH7CFGR1 = 0x01 << 7  //+ CHEN
	CHINSEL  CH7CFGR1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CH7CFGR1 = 0x03 << 12 //+ DATMPX
	DATPACK  CH7CFGR1 = 0x03 << 14 //+ DATPACK
)

const (
	SITPn     = 0
	SPICKSELn = 2
	SCDENn    = 5
	CKABENn   = 6
	CHENn     = 7
	CHINSELn  = 8
	DATMPXn   = 12
	DATPACKn  = 14
)

const (
	DTRBS  CH7CFGR2 = 0x1F << 3     //+ DTRBS
	OFFSET CH7CFGR2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   CH7AWSCDR = 0xFF << 0  //+ SCDT
	BKSCD  CH7AWSCDR = 0x0F << 12 //+ BKSCD
	AWFOSR CH7AWSCDR = 0x1F << 16 //+ AWFOSR
	AWFORD CH7AWSCDR = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CH7WDATR = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CH7DATINR = 0xFFFF << 0  //+ INDAT0
	INDAT1 CH7DATINR = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	PLSSKP CH7DLYR = 0x3F << 0 //+ PLSSKP
)

const (
	PLSSKPn = 0
)

const (
	DFEN     DFSDM_FLT0CR1 = 0x01 << 0  //+ DFSDM enable
	JSWSTART DFSDM_FLT0CR1 = 0x01 << 1  //+ Start a conversion of the injected group of channels
	JSYNC    DFSDM_FLT0CR1 = 0x01 << 3  //+ Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
	JSCAN    DFSDM_FLT0CR1 = 0x01 << 4  //+ Scanning conversion mode for injected conversions
	JDMAEN   DFSDM_FLT0CR1 = 0x01 << 5  //+ DMA channel enabled to read data for the injected channel group
	JEXTSEL  DFSDM_FLT0CR1 = 0x07 << 8  //+ Trigger signal selection for launching injected conversions
	JEXTEN   DFSDM_FLT0CR1 = 0x03 << 13 //+ Trigger enable and trigger edge selection for injected conversions
	RSWSTART DFSDM_FLT0CR1 = 0x01 << 17 //+ Software start of a conversion on the regular channel
	RCONT    DFSDM_FLT0CR1 = 0x01 << 18 //+ Continuous mode selection for regular conversions
	RSYNC    DFSDM_FLT0CR1 = 0x01 << 19 //+ Launch regular conversion synchronously with DFSDM0
	RDMAEN   DFSDM_FLT0CR1 = 0x01 << 21 //+ DMA channel enabled to read data for the regular conversion
	RCH      DFSDM_FLT0CR1 = 0x07 << 24 //+ Regular channel selection
	FAST     DFSDM_FLT0CR1 = 0x01 << 29 //+ Fast conversion mode selection for regular conversions
	AWFSEL   DFSDM_FLT0CR1 = 0x01 << 30 //+ Analog watchdog fast mode select
)

const (
	DFENn     = 0
	JSWSTARTn = 1
	JSYNCn    = 3
	JSCANn    = 4
	JDMAENn   = 5
	JEXTSELn  = 8
	JEXTENn   = 13
	RSWSTARTn = 17
	RCONTn    = 18
	RSYNCn    = 19
	RDMAENn   = 21
	RCHn      = 24
	FASTn     = 29
	AWFSELn   = 30
)

const (
	JEOCIE DFSDM_FLT0CR2 = 0x01 << 0  //+ Injected end of conversion interrupt enable
	REOCIE DFSDM_FLT0CR2 = 0x01 << 1  //+ Regular end of conversion interrupt enable
	JOVRIE DFSDM_FLT0CR2 = 0x01 << 2  //+ Injected data overrun interrupt enable
	ROVRIE DFSDM_FLT0CR2 = 0x01 << 3  //+ Regular data overrun interrupt enable
	AWDIE  DFSDM_FLT0CR2 = 0x01 << 4  //+ Analog watchdog interrupt enable
	SCDIE  DFSDM_FLT0CR2 = 0x01 << 5  //+ Short-circuit detector interrupt enable
	CKABIE DFSDM_FLT0CR2 = 0x01 << 6  //+ Clock absence interrupt enable
	EXCH   DFSDM_FLT0CR2 = 0xFF << 8  //+ Extremes detector channel selection
	AWDCH  DFSDM_FLT0CR2 = 0xFF << 16 //+ Analog watchdog channel selection
)

const (
	JEOCIEn = 0
	REOCIEn = 1
	JOVRIEn = 2
	ROVRIEn = 3
	AWDIEn  = 4
	SCDIEn  = 5
	CKABIEn = 6
	EXCHn   = 8
	AWDCHn  = 16
)

const (
	JEOCF DFSDM_FLT0ISR = 0x01 << 0  //+ End of injected conversion flag
	REOCF DFSDM_FLT0ISR = 0x01 << 1  //+ End of regular conversion flag
	JOVRF DFSDM_FLT0ISR = 0x01 << 2  //+ Injected conversion overrun flag
	ROVRF DFSDM_FLT0ISR = 0x01 << 3  //+ Regular conversion overrun flag
	AWDF  DFSDM_FLT0ISR = 0x01 << 4  //+ Analog watchdog
	JCIP  DFSDM_FLT0ISR = 0x01 << 13 //+ Injected conversion in progress status
	RCIP  DFSDM_FLT0ISR = 0x01 << 14 //+ Regular conversion in progress status
	CKABF DFSDM_FLT0ISR = 0xFF << 16 //+ Clock absence flag
	SCDF  DFSDM_FLT0ISR = 0xFF << 24 //+ short-circuit detector flag
)

const (
	JEOCFn = 0
	REOCFn = 1
	JOVRFn = 2
	ROVRFn = 3
	AWDFn  = 4
	JCIPn  = 13
	RCIPn  = 14
	CKABFn = 16
	SCDFn  = 24
)

const (
	CLRJOVRF DFSDM_FLT0ICR = 0x01 << 2  //+ Clear the injected conversion overrun flag
	CLRROVRF DFSDM_FLT0ICR = 0x01 << 3  //+ Clear the regular conversion overrun flag
	CLRCKABF DFSDM_FLT0ICR = 0xFF << 16 //+ Clear the clock absence flag
	CLRSCDF  DFSDM_FLT0ICR = 0xFF << 24 //+ Clear the short-circuit detector flag
)

const (
	CLRJOVRFn = 2
	CLRROVRFn = 3
	CLRCKABFn = 16
	CLRSCDFn  = 24
)

const (
	JCHG DFSDM_FLT0JCHGR = 0xFF << 0 //+ Injected channel group selection
)

const (
	JCHGn = 0
)

const (
	IOSR DFSDM_FLT0FCR = 0xFF << 0   //+ Integrator oversampling ratio (averaging length)
	FOSR DFSDM_FLT0FCR = 0x3FF << 16 //+ Sinc filter oversampling ratio (decimation rate)
	FORD DFSDM_FLT0FCR = 0x07 << 29  //+ Sinc filter order
)

const (
	IOSRn = 0
	FOSRn = 16
	FORDn = 29
)

const (
	JDATACH DFSDM_FLT0JDATAR = 0x07 << 0     //+ Injected channel most recently converted
	JDATA   DFSDM_FLT0JDATAR = 0xFFFFFF << 8 //+ Injected group conversion data
)

const (
	JDATACHn = 0
	JDATAn   = 8
)

const (
	RDATACH DFSDM_FLT0RDATAR = 0x07 << 0     //+ Regular channel most recently converted
	RPEND   DFSDM_FLT0RDATAR = 0x01 << 4     //+ Regular channel pending data
	RDATA   DFSDM_FLT0RDATAR = 0xFFFFFF << 8 //+ Regular channel conversion data
)

const (
	RDATACHn = 0
	RPENDn   = 4
	RDATAn   = 8
)

const (
	BKAWH DFSDM_FLT0AWHTR = 0x0F << 0     //+ Break signal assignment to analog watchdog high threshold event
	AWHT  DFSDM_FLT0AWHTR = 0xFFFFFF << 8 //+ Analog watchdog high threshold
)

const (
	BKAWHn = 0
	AWHTn  = 8
)

const (
	BKAWL DFSDM_FLT0AWLTR = 0x0F << 0     //+ Break signal assignment to analog watchdog low threshold event
	AWLT  DFSDM_FLT0AWLTR = 0xFFFFFF << 8 //+ Analog watchdog low threshold
)

const (
	BKAWLn = 0
	AWLTn  = 8
)

const (
	AWLTF DFSDM_FLT0AWSR = 0xFF << 0 //+ Analog watchdog low threshold flag
	AWHTF DFSDM_FLT0AWSR = 0xFF << 8 //+ Analog watchdog high threshold flag
)

const (
	AWLTFn = 0
	AWHTFn = 8
)

const (
	CLRAWLTF DFSDM_FLT0AWCFR = 0xFF << 0 //+ Clear the analog watchdog low threshold flag
	CLRAWHTF DFSDM_FLT0AWCFR = 0xFF << 8 //+ Clear the analog watchdog high threshold flag
)

const (
	CLRAWLTFn = 0
	CLRAWHTFn = 8
)

const (
	EXMAXCH DFSDM_FLT0EXMAX = 0x07 << 0     //+ Extremes detector maximum data channel
	EXMAX   DFSDM_FLT0EXMAX = 0xFFFFFF << 8 //+ Extremes detector maximum value
)

const (
	EXMAXCHn = 0
	EXMAXn   = 8
)

const (
	EXMINCH DFSDM_FLT0EXMIN = 0x07 << 0     //+ Extremes detector minimum data channel
	EXMIN   DFSDM_FLT0EXMIN = 0xFFFFFF << 8 //+ EXMIN
)

const (
	EXMINCHn = 0
	EXMINn   = 8
)

const (
	CNVCNT DFSDM_FLT0CNVTIMR = 0xFFFFFFF << 4 //+ 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
)

const (
	CNVCNTn = 4
)

const (
	DFEN     DFSDM_FLT1CR1 = 0x01 << 0  //+ DFSDM enable
	JSWSTART DFSDM_FLT1CR1 = 0x01 << 1  //+ Start a conversion of the injected group of channels
	JSYNC    DFSDM_FLT1CR1 = 0x01 << 3  //+ Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
	JSCAN    DFSDM_FLT1CR1 = 0x01 << 4  //+ Scanning conversion mode for injected conversions
	JDMAEN   DFSDM_FLT1CR1 = 0x01 << 5  //+ DMA channel enabled to read data for the injected channel group
	JEXTSEL  DFSDM_FLT1CR1 = 0x07 << 8  //+ Trigger signal selection for launching injected conversions
	JEXTEN   DFSDM_FLT1CR1 = 0x03 << 13 //+ Trigger enable and trigger edge selection for injected conversions
	RSWSTART DFSDM_FLT1CR1 = 0x01 << 17 //+ Software start of a conversion on the regular channel
	RCONT    DFSDM_FLT1CR1 = 0x01 << 18 //+ Continuous mode selection for regular conversions
	RSYNC    DFSDM_FLT1CR1 = 0x01 << 19 //+ Launch regular conversion synchronously with DFSDM0
	RDMAEN   DFSDM_FLT1CR1 = 0x01 << 21 //+ DMA channel enabled to read data for the regular conversion
	RCH      DFSDM_FLT1CR1 = 0x07 << 24 //+ Regular channel selection
	FAST     DFSDM_FLT1CR1 = 0x01 << 29 //+ Fast conversion mode selection for regular conversions
	AWFSEL   DFSDM_FLT1CR1 = 0x01 << 30 //+ Analog watchdog fast mode select
)

const (
	DFENn     = 0
	JSWSTARTn = 1
	JSYNCn    = 3
	JSCANn    = 4
	JDMAENn   = 5
	JEXTSELn  = 8
	JEXTENn   = 13
	RSWSTARTn = 17
	RCONTn    = 18
	RSYNCn    = 19
	RDMAENn   = 21
	RCHn      = 24
	FASTn     = 29
	AWFSELn   = 30
)

const (
	JEOCIE DFSDM_FLT1CR2 = 0x01 << 0  //+ Injected end of conversion interrupt enable
	REOCIE DFSDM_FLT1CR2 = 0x01 << 1  //+ Regular end of conversion interrupt enable
	JOVRIE DFSDM_FLT1CR2 = 0x01 << 2  //+ Injected data overrun interrupt enable
	ROVRIE DFSDM_FLT1CR2 = 0x01 << 3  //+ Regular data overrun interrupt enable
	AWDIE  DFSDM_FLT1CR2 = 0x01 << 4  //+ Analog watchdog interrupt enable
	SCDIE  DFSDM_FLT1CR2 = 0x01 << 5  //+ Short-circuit detector interrupt enable
	CKABIE DFSDM_FLT1CR2 = 0x01 << 6  //+ Clock absence interrupt enable
	EXCH   DFSDM_FLT1CR2 = 0xFF << 8  //+ Extremes detector channel selection
	AWDCH  DFSDM_FLT1CR2 = 0xFF << 16 //+ Analog watchdog channel selection
)

const (
	JEOCIEn = 0
	REOCIEn = 1
	JOVRIEn = 2
	ROVRIEn = 3
	AWDIEn  = 4
	SCDIEn  = 5
	CKABIEn = 6
	EXCHn   = 8
	AWDCHn  = 16
)

const (
	JEOCF DFSDM_FLT1ISR = 0x01 << 0  //+ End of injected conversion flag
	REOCF DFSDM_FLT1ISR = 0x01 << 1  //+ End of regular conversion flag
	JOVRF DFSDM_FLT1ISR = 0x01 << 2  //+ Injected conversion overrun flag
	ROVRF DFSDM_FLT1ISR = 0x01 << 3  //+ Regular conversion overrun flag
	AWDF  DFSDM_FLT1ISR = 0x01 << 4  //+ Analog watchdog
	JCIP  DFSDM_FLT1ISR = 0x01 << 13 //+ Injected conversion in progress status
	RCIP  DFSDM_FLT1ISR = 0x01 << 14 //+ Regular conversion in progress status
	CKABF DFSDM_FLT1ISR = 0xFF << 16 //+ Clock absence flag
	SCDF  DFSDM_FLT1ISR = 0xFF << 24 //+ short-circuit detector flag
)

const (
	JEOCFn = 0
	REOCFn = 1
	JOVRFn = 2
	ROVRFn = 3
	AWDFn  = 4
	JCIPn  = 13
	RCIPn  = 14
	CKABFn = 16
	SCDFn  = 24
)

const (
	CLRJOVRF DFSDM_FLT1ICR = 0x01 << 2  //+ Clear the injected conversion overrun flag
	CLRROVRF DFSDM_FLT1ICR = 0x01 << 3  //+ Clear the regular conversion overrun flag
	CLRCKABF DFSDM_FLT1ICR = 0xFF << 16 //+ Clear the clock absence flag
	CLRSCDF  DFSDM_FLT1ICR = 0xFF << 24 //+ Clear the short-circuit detector flag
)

const (
	CLRJOVRFn = 2
	CLRROVRFn = 3
	CLRCKABFn = 16
	CLRSCDFn  = 24
)

const (
	JCHG DFSDM_FLT1CHGR = 0xFF << 0 //+ Injected channel group selection
)

const (
	JCHGn = 0
)

const (
	IOSR DFSDM_FLT1FCR = 0xFF << 0   //+ Integrator oversampling ratio (averaging length)
	FOSR DFSDM_FLT1FCR = 0x3FF << 16 //+ Sinc filter oversampling ratio (decimation rate)
	FORD DFSDM_FLT1FCR = 0x07 << 29  //+ Sinc filter order
)

const (
	IOSRn = 0
	FOSRn = 16
	FORDn = 29
)

const (
	JDATACH DFSDM_FLT1JDATAR = 0x07 << 0     //+ Injected channel most recently converted
	JDATA   DFSDM_FLT1JDATAR = 0xFFFFFF << 8 //+ Injected group conversion data
)

const (
	JDATACHn = 0
	JDATAn   = 8
)

const (
	RDATACH DFSDM_FLT1RDATAR = 0x07 << 0     //+ Regular channel most recently converted
	RPEND   DFSDM_FLT1RDATAR = 0x01 << 4     //+ Regular channel pending data
	RDATA   DFSDM_FLT1RDATAR = 0xFFFFFF << 8 //+ Regular channel conversion data
)

const (
	RDATACHn = 0
	RPENDn   = 4
	RDATAn   = 8
)

const (
	BKAWH DFSDM_FLT1AWHTR = 0x0F << 0     //+ Break signal assignment to analog watchdog high threshold event
	AWHT  DFSDM_FLT1AWHTR = 0xFFFFFF << 8 //+ Analog watchdog high threshold
)

const (
	BKAWHn = 0
	AWHTn  = 8
)

const (
	BKAWL DFSDM_FLT1AWLTR = 0x0F << 0     //+ Break signal assignment to analog watchdog low threshold event
	AWLT  DFSDM_FLT1AWLTR = 0xFFFFFF << 8 //+ Analog watchdog low threshold
)

const (
	BKAWLn = 0
	AWLTn  = 8
)

const (
	AWLTF DFSDM_FLT1AWSR = 0xFF << 0 //+ Analog watchdog low threshold flag
	AWHTF DFSDM_FLT1AWSR = 0xFF << 8 //+ Analog watchdog high threshold flag
)

const (
	AWLTFn = 0
	AWHTFn = 8
)

const (
	CLRAWLTF DFSDM_FLT1AWCFR = 0xFF << 0 //+ Clear the analog watchdog low threshold flag
	CLRAWHTF DFSDM_FLT1AWCFR = 0xFF << 8 //+ Clear the analog watchdog high threshold flag
)

const (
	CLRAWLTFn = 0
	CLRAWHTFn = 8
)

const (
	EXMAXCH DFSDM_FLT1EXMAX = 0x07 << 0     //+ Extremes detector maximum data channel
	EXMAX   DFSDM_FLT1EXMAX = 0xFFFFFF << 8 //+ Extremes detector maximum value
)

const (
	EXMAXCHn = 0
	EXMAXn   = 8
)

const (
	EXMINCH DFSDM_FLT1EXMIN = 0x07 << 0     //+ Extremes detector minimum data channel
	EXMIN   DFSDM_FLT1EXMIN = 0xFFFFFF << 8 //+ EXMIN
)

const (
	EXMINCHn = 0
	EXMINn   = 8
)

const (
	CNVCNT DFSDM_FLT1CNVTIMR = 0xFFFFFFF << 4 //+ 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
)

const (
	CNVCNTn = 4
)

const (
	DFEN     DFSDM_FLT2CR1 = 0x01 << 0  //+ DFSDM enable
	JSWSTART DFSDM_FLT2CR1 = 0x01 << 1  //+ Start a conversion of the injected group of channels
	JSYNC    DFSDM_FLT2CR1 = 0x01 << 3  //+ Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
	JSCAN    DFSDM_FLT2CR1 = 0x01 << 4  //+ Scanning conversion mode for injected conversions
	JDMAEN   DFSDM_FLT2CR1 = 0x01 << 5  //+ DMA channel enabled to read data for the injected channel group
	JEXTSEL  DFSDM_FLT2CR1 = 0x07 << 8  //+ Trigger signal selection for launching injected conversions
	JEXTEN   DFSDM_FLT2CR1 = 0x03 << 13 //+ Trigger enable and trigger edge selection for injected conversions
	RSWSTART DFSDM_FLT2CR1 = 0x01 << 17 //+ Software start of a conversion on the regular channel
	RCONT    DFSDM_FLT2CR1 = 0x01 << 18 //+ Continuous mode selection for regular conversions
	RSYNC    DFSDM_FLT2CR1 = 0x01 << 19 //+ Launch regular conversion synchronously with DFSDM0
	RDMAEN   DFSDM_FLT2CR1 = 0x01 << 21 //+ DMA channel enabled to read data for the regular conversion
	RCH      DFSDM_FLT2CR1 = 0x07 << 24 //+ Regular channel selection
	FAST     DFSDM_FLT2CR1 = 0x01 << 29 //+ Fast conversion mode selection for regular conversions
	AWFSEL   DFSDM_FLT2CR1 = 0x01 << 30 //+ Analog watchdog fast mode select
)

const (
	DFENn     = 0
	JSWSTARTn = 1
	JSYNCn    = 3
	JSCANn    = 4
	JDMAENn   = 5
	JEXTSELn  = 8
	JEXTENn   = 13
	RSWSTARTn = 17
	RCONTn    = 18
	RSYNCn    = 19
	RDMAENn   = 21
	RCHn      = 24
	FASTn     = 29
	AWFSELn   = 30
)

const (
	JEOCIE DFSDM_FLT2CR2 = 0x01 << 0  //+ Injected end of conversion interrupt enable
	REOCIE DFSDM_FLT2CR2 = 0x01 << 1  //+ Regular end of conversion interrupt enable
	JOVRIE DFSDM_FLT2CR2 = 0x01 << 2  //+ Injected data overrun interrupt enable
	ROVRIE DFSDM_FLT2CR2 = 0x01 << 3  //+ Regular data overrun interrupt enable
	AWDIE  DFSDM_FLT2CR2 = 0x01 << 4  //+ Analog watchdog interrupt enable
	SCDIE  DFSDM_FLT2CR2 = 0x01 << 5  //+ Short-circuit detector interrupt enable
	CKABIE DFSDM_FLT2CR2 = 0x01 << 6  //+ Clock absence interrupt enable
	EXCH   DFSDM_FLT2CR2 = 0xFF << 8  //+ Extremes detector channel selection
	AWDCH  DFSDM_FLT2CR2 = 0xFF << 16 //+ Analog watchdog channel selection
)

const (
	JEOCIEn = 0
	REOCIEn = 1
	JOVRIEn = 2
	ROVRIEn = 3
	AWDIEn  = 4
	SCDIEn  = 5
	CKABIEn = 6
	EXCHn   = 8
	AWDCHn  = 16
)

const (
	JEOCF DFSDM_FLT2ISR = 0x01 << 0  //+ End of injected conversion flag
	REOCF DFSDM_FLT2ISR = 0x01 << 1  //+ End of regular conversion flag
	JOVRF DFSDM_FLT2ISR = 0x01 << 2  //+ Injected conversion overrun flag
	ROVRF DFSDM_FLT2ISR = 0x01 << 3  //+ Regular conversion overrun flag
	AWDF  DFSDM_FLT2ISR = 0x01 << 4  //+ Analog watchdog
	JCIP  DFSDM_FLT2ISR = 0x01 << 13 //+ Injected conversion in progress status
	RCIP  DFSDM_FLT2ISR = 0x01 << 14 //+ Regular conversion in progress status
	CKABF DFSDM_FLT2ISR = 0xFF << 16 //+ Clock absence flag
	SCDF  DFSDM_FLT2ISR = 0xFF << 24 //+ short-circuit detector flag
)

const (
	JEOCFn = 0
	REOCFn = 1
	JOVRFn = 2
	ROVRFn = 3
	AWDFn  = 4
	JCIPn  = 13
	RCIPn  = 14
	CKABFn = 16
	SCDFn  = 24
)

const (
	CLRJOVRF DFSDM_FLT2ICR = 0x01 << 2  //+ Clear the injected conversion overrun flag
	CLRROVRF DFSDM_FLT2ICR = 0x01 << 3  //+ Clear the regular conversion overrun flag
	CLRCKABF DFSDM_FLT2ICR = 0xFF << 16 //+ Clear the clock absence flag
	CLRSCDF  DFSDM_FLT2ICR = 0xFF << 24 //+ Clear the short-circuit detector flag
)

const (
	CLRJOVRFn = 2
	CLRROVRFn = 3
	CLRCKABFn = 16
	CLRSCDFn  = 24
)

const (
	JCHG DFSDM_FLT2JCHGR = 0xFF << 0 //+ Injected channel group selection
)

const (
	JCHGn = 0
)

const (
	IOSR DFSDM_FLT2FCR = 0xFF << 0   //+ Integrator oversampling ratio (averaging length)
	FOSR DFSDM_FLT2FCR = 0x3FF << 16 //+ Sinc filter oversampling ratio (decimation rate)
	FORD DFSDM_FLT2FCR = 0x07 << 29  //+ Sinc filter order
)

const (
	IOSRn = 0
	FOSRn = 16
	FORDn = 29
)

const (
	JDATACH DFSDM_FLT2JDATAR = 0x07 << 0     //+ Injected channel most recently converted
	JDATA   DFSDM_FLT2JDATAR = 0xFFFFFF << 8 //+ Injected group conversion data
)

const (
	JDATACHn = 0
	JDATAn   = 8
)

const (
	RDATACH DFSDM_FLT2RDATAR = 0x07 << 0     //+ Regular channel most recently converted
	RPEND   DFSDM_FLT2RDATAR = 0x01 << 4     //+ Regular channel pending data
	RDATA   DFSDM_FLT2RDATAR = 0xFFFFFF << 8 //+ Regular channel conversion data
)

const (
	RDATACHn = 0
	RPENDn   = 4
	RDATAn   = 8
)

const (
	BKAWH DFSDM_FLT2AWHTR = 0x0F << 0     //+ Break signal assignment to analog watchdog high threshold event
	AWHT  DFSDM_FLT2AWHTR = 0xFFFFFF << 8 //+ Analog watchdog high threshold
)

const (
	BKAWHn = 0
	AWHTn  = 8
)

const (
	BKAWL DFSDM_FLT2AWLTR = 0x0F << 0     //+ Break signal assignment to analog watchdog low threshold event
	AWLT  DFSDM_FLT2AWLTR = 0xFFFFFF << 8 //+ Analog watchdog low threshold
)

const (
	BKAWLn = 0
	AWLTn  = 8
)

const (
	AWLTF DFSDM_FLT2AWSR = 0xFF << 0 //+ Analog watchdog low threshold flag
	AWHTF DFSDM_FLT2AWSR = 0xFF << 8 //+ Analog watchdog high threshold flag
)

const (
	AWLTFn = 0
	AWHTFn = 8
)

const (
	CLRAWLTF DFSDM_FLT2AWCFR = 0xFF << 0 //+ Clear the analog watchdog low threshold flag
	CLRAWHTF DFSDM_FLT2AWCFR = 0xFF << 8 //+ Clear the analog watchdog high threshold flag
)

const (
	CLRAWLTFn = 0
	CLRAWHTFn = 8
)

const (
	EXMAXCH DFSDM_FLT2EXMAX = 0x07 << 0     //+ Extremes detector maximum data channel
	EXMAX   DFSDM_FLT2EXMAX = 0xFFFFFF << 8 //+ Extremes detector maximum value
)

const (
	EXMAXCHn = 0
	EXMAXn   = 8
)

const (
	EXMINCH DFSDM_FLT2EXMIN = 0x07 << 0     //+ Extremes detector minimum data channel
	EXMIN   DFSDM_FLT2EXMIN = 0xFFFFFF << 8 //+ EXMIN
)

const (
	EXMINCHn = 0
	EXMINn   = 8
)

const (
	CNVCNT DFSDM_FLT2CNVTIMR = 0xFFFFFFF << 4 //+ 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
)

const (
	CNVCNTn = 4
)

const (
	DFEN     DFSDM_FLT3CR1 = 0x01 << 0  //+ DFSDM enable
	JSWSTART DFSDM_FLT3CR1 = 0x01 << 1  //+ Start a conversion of the injected group of channels
	JSYNC    DFSDM_FLT3CR1 = 0x01 << 3  //+ Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
	JSCAN    DFSDM_FLT3CR1 = 0x01 << 4  //+ Scanning conversion mode for injected conversions
	JDMAEN   DFSDM_FLT3CR1 = 0x01 << 5  //+ DMA channel enabled to read data for the injected channel group
	JEXTSEL  DFSDM_FLT3CR1 = 0x07 << 8  //+ Trigger signal selection for launching injected conversions
	JEXTEN   DFSDM_FLT3CR1 = 0x03 << 13 //+ Trigger enable and trigger edge selection for injected conversions
	RSWSTART DFSDM_FLT3CR1 = 0x01 << 17 //+ Software start of a conversion on the regular channel
	RCONT    DFSDM_FLT3CR1 = 0x01 << 18 //+ Continuous mode selection for regular conversions
	RSYNC    DFSDM_FLT3CR1 = 0x01 << 19 //+ Launch regular conversion synchronously with DFSDM0
	RDMAEN   DFSDM_FLT3CR1 = 0x01 << 21 //+ DMA channel enabled to read data for the regular conversion
	RCH      DFSDM_FLT3CR1 = 0x07 << 24 //+ Regular channel selection
	FAST     DFSDM_FLT3CR1 = 0x01 << 29 //+ Fast conversion mode selection for regular conversions
	AWFSEL   DFSDM_FLT3CR1 = 0x01 << 30 //+ Analog watchdog fast mode select
)

const (
	DFENn     = 0
	JSWSTARTn = 1
	JSYNCn    = 3
	JSCANn    = 4
	JDMAENn   = 5
	JEXTSELn  = 8
	JEXTENn   = 13
	RSWSTARTn = 17
	RCONTn    = 18
	RSYNCn    = 19
	RDMAENn   = 21
	RCHn      = 24
	FASTn     = 29
	AWFSELn   = 30
)

const (
	JEOCIE DFSDM_FLT3CR2 = 0x01 << 0  //+ Injected end of conversion interrupt enable
	REOCIE DFSDM_FLT3CR2 = 0x01 << 1  //+ Regular end of conversion interrupt enable
	JOVRIE DFSDM_FLT3CR2 = 0x01 << 2  //+ Injected data overrun interrupt enable
	ROVRIE DFSDM_FLT3CR2 = 0x01 << 3  //+ Regular data overrun interrupt enable
	AWDIE  DFSDM_FLT3CR2 = 0x01 << 4  //+ Analog watchdog interrupt enable
	SCDIE  DFSDM_FLT3CR2 = 0x01 << 5  //+ Short-circuit detector interrupt enable
	CKABIE DFSDM_FLT3CR2 = 0x01 << 6  //+ Clock absence interrupt enable
	EXCH   DFSDM_FLT3CR2 = 0xFF << 8  //+ Extremes detector channel selection
	AWDCH  DFSDM_FLT3CR2 = 0xFF << 16 //+ Analog watchdog channel selection
)

const (
	JEOCIEn = 0
	REOCIEn = 1
	JOVRIEn = 2
	ROVRIEn = 3
	AWDIEn  = 4
	SCDIEn  = 5
	CKABIEn = 6
	EXCHn   = 8
	AWDCHn  = 16
)

const (
	JEOCF DFSDM_FLT3ISR = 0x01 << 0  //+ End of injected conversion flag
	REOCF DFSDM_FLT3ISR = 0x01 << 1  //+ End of regular conversion flag
	JOVRF DFSDM_FLT3ISR = 0x01 << 2  //+ Injected conversion overrun flag
	ROVRF DFSDM_FLT3ISR = 0x01 << 3  //+ Regular conversion overrun flag
	AWDF  DFSDM_FLT3ISR = 0x01 << 4  //+ Analog watchdog
	JCIP  DFSDM_FLT3ISR = 0x01 << 13 //+ Injected conversion in progress status
	RCIP  DFSDM_FLT3ISR = 0x01 << 14 //+ Regular conversion in progress status
	CKABF DFSDM_FLT3ISR = 0xFF << 16 //+ Clock absence flag
	SCDF  DFSDM_FLT3ISR = 0xFF << 24 //+ short-circuit detector flag
)

const (
	JEOCFn = 0
	REOCFn = 1
	JOVRFn = 2
	ROVRFn = 3
	AWDFn  = 4
	JCIPn  = 13
	RCIPn  = 14
	CKABFn = 16
	SCDFn  = 24
)

const (
	CLRJOVRF DFSDM_FLT3ICR = 0x01 << 2  //+ Clear the injected conversion overrun flag
	CLRROVRF DFSDM_FLT3ICR = 0x01 << 3  //+ Clear the regular conversion overrun flag
	CLRCKABF DFSDM_FLT3ICR = 0xFF << 16 //+ Clear the clock absence flag
	CLRSCDF  DFSDM_FLT3ICR = 0xFF << 24 //+ Clear the short-circuit detector flag
)

const (
	CLRJOVRFn = 2
	CLRROVRFn = 3
	CLRCKABFn = 16
	CLRSCDFn  = 24
)

const (
	JCHG DFSDM_FLT3JCHGR = 0xFF << 0 //+ Injected channel group selection
)

const (
	JCHGn = 0
)

const (
	IOSR DFSDM_FLT3FCR = 0xFF << 0   //+ Integrator oversampling ratio (averaging length)
	FOSR DFSDM_FLT3FCR = 0x3FF << 16 //+ Sinc filter oversampling ratio (decimation rate)
	FORD DFSDM_FLT3FCR = 0x07 << 29  //+ Sinc filter order
)

const (
	IOSRn = 0
	FOSRn = 16
	FORDn = 29
)

const (
	JDATACH DFSDM_FLT3JDATAR = 0x07 << 0     //+ Injected channel most recently converted
	JDATA   DFSDM_FLT3JDATAR = 0xFFFFFF << 8 //+ Injected group conversion data
)

const (
	JDATACHn = 0
	JDATAn   = 8
)

const (
	RDATACH DFSDM_FLT3RDATAR = 0x07 << 0     //+ Regular channel most recently converted
	RPEND   DFSDM_FLT3RDATAR = 0x01 << 4     //+ Regular channel pending data
	RDATA   DFSDM_FLT3RDATAR = 0xFFFFFF << 8 //+ Regular channel conversion data
)

const (
	RDATACHn = 0
	RPENDn   = 4
	RDATAn   = 8
)

const (
	BKAWH DFSDM_FLT3AWHTR = 0x0F << 0     //+ Break signal assignment to analog watchdog high threshold event
	AWHT  DFSDM_FLT3AWHTR = 0xFFFFFF << 8 //+ Analog watchdog high threshold
)

const (
	BKAWHn = 0
	AWHTn  = 8
)

const (
	BKAWL DFSDM_FLT3AWLTR = 0x0F << 0     //+ Break signal assignment to analog watchdog low threshold event
	AWLT  DFSDM_FLT3AWLTR = 0xFFFFFF << 8 //+ Analog watchdog low threshold
)

const (
	BKAWLn = 0
	AWLTn  = 8
)

const (
	AWLTF DFSDM_FLT3AWSR = 0xFF << 0 //+ Analog watchdog low threshold flag
	AWHTF DFSDM_FLT3AWSR = 0xFF << 8 //+ Analog watchdog high threshold flag
)

const (
	AWLTFn = 0
	AWHTFn = 8
)

const (
	CLRAWLTF DFSDM_FLT3AWCFR = 0xFF << 0 //+ Clear the analog watchdog low threshold flag
	CLRAWHTF DFSDM_FLT3AWCFR = 0xFF << 8 //+ Clear the analog watchdog high threshold flag
)

const (
	CLRAWLTFn = 0
	CLRAWHTFn = 8
)

const (
	EXMAXCH DFSDM_FLT3EXMAX = 0x07 << 0     //+ Extremes detector maximum data channel
	EXMAX   DFSDM_FLT3EXMAX = 0xFFFFFF << 8 //+ Extremes detector maximum value
)

const (
	EXMAXCHn = 0
	EXMAXn   = 8
)

const (
	EXMINCH DFSDM_FLT3EXMIN = 0x07 << 0     //+ Extremes detector minimum data channel
	EXMIN   DFSDM_FLT3EXMIN = 0xFFFFFF << 8 //+ EXMIN
)

const (
	EXMINCHn = 0
	EXMINn   = 8
)

const (
	CNVCNT DFSDM_FLT3CNVTIMR = 0xFFFFFFF << 4 //+ 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
)

const (
	CNVCNTn = 4
)
