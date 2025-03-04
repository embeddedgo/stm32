// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f412

// Package dfsdm provides access to the registers of the DFSDM peripheral.
//
// Instances:
//
//	DFSDM  DFSDM_BASE  APB2  -  Digital filter for sigma delta modulators
//
// Registers:
//
//	0x000 32  CHCFG0R1        channel configuration y register
//	0x004 32  CHCFG0R2        channel configuration y register
//	0x008 32  AWSCD0R         analog watchdog and short-circuit detector register
//	0x00C 32  CHWDAT0R        channel watchdog filter data register
//	0x010 32  CHDATIN0R       channel data input register
//	0x020 32  CHCFG1R1        CHCFG1R1
//	0x024 32  CHCFG1R2        CHCFG1R2
//	0x028 32  AWSCD1R         AWSCD1R
//	0x02C 32  CHWDAT1R        CHWDAT1R
//	0x030 32  CHDATIN1R       CHDATIN1R
//	0x040 32  CHCFG2R1        CHCFG2R1
//	0x044 32  CHCFG2R2        CHCFG2R2
//	0x048 32  AWSCD2R         AWSCD2R
//	0x04C 32  CHWDAT2R        CHWDAT2R
//	0x050 32  CHDATIN2R       CHDATIN2R
//	0x060 32  CHCFG3R1        CHCFG3R1
//	0x064 32  CHCFG3R2        CHCFG3R2
//	0x068 32  AWSCD3R         AWSCD3R
//	0x06C 32  CHWDAT3R        CHWDAT3R
//	0x070 32  CHDATIN3R       CHDATIN3R
//	0x080 32  CHCFG4R1        CHCFG4R1
//	0x084 32  CHCFG4R2        CHCFG4R2
//	0x088 32  AWSCD4R         AWSCD4R
//	0x08C 32  CHWDAT4R        CHWDAT4R
//	0x090 32  CHDATIN4R       CHDATIN4R
//	0x0A0 32  CHCFG5R1        CHCFG5R1
//	0x0A4 32  CHCFG5R2        CHCFG5R2
//	0x0A8 32  AWSCD5R         AWSCD5R
//	0x0AC 32  CHWDAT5R        CHWDAT5R
//	0x0B0 32  CHDATIN5R       CHDATIN5R
//	0x0C0 32  CHCFG6R1        CHCFG6R1
//	0x0C4 32  CHCFG6R2        CHCFG6R2
//	0x0C8 32  AWSCD6R         AWSCD6R
//	0x0CC 32  CHWDAT6R        CHWDAT6R
//	0x0D0 32  CHDATIN6R       CHDATIN6R
//	0x0E0 32  CHCFG7R1        CHCFG7R1
//	0x0E4 32  CHCFG7R2        CHCFG7R2
//	0x0E8 32  AWSCD7R         AWSCD7R
//	0x0EC 32  CHWDAT7R        CHWDAT7R
//	0x0F0 32  CHDATIN7R       CHDATIN7R
//	0x100 32  DFSDM0_CR1      control register 1
//	0x104 32  DFSDM0_CR2      control register 2
//	0x108 32  DFSDM0_ISR      interrupt and status register
//	0x10C 32  DFSDM0_ICR      interrupt flag clear register
//	0x110 32  DFSDM0_JCHGR    injected channel group selection register
//	0x114 32  DFSDM0_FCR      filter control register
//	0x118 32  DFSDM0_JDATAR   data register for injected group
//	0x11C 32  DFSDM0_RDATAR   data register for the regular channel
//	0x120 32  DFSDM0_AWHTR    analog watchdog high threshold register
//	0x124 32  DFSDM0_AWLTR    analog watchdog low threshold register
//	0x128 32  DFSDM0_AWSR     analog watchdog status register
//	0x12C 32  DFSDM0_AWCFR    analog watchdog clear flag register
//	0x130 32  DFSDM0_EXMAX    Extremes detector maximum register
//	0x134 32  DFSDM0_EXMIN    Extremes detector minimum register
//	0x138 32  DFSDM0_CNVTIMR  conversion timer register
//	0x200 32  DFSDM1_CR1      control register 1
//	0x204 32  DFSDM1_CR2      control register 2
//	0x208 32  DFSDM1_ISR      interrupt and status register
//	0x20C 32  DFSDM1_ICR      interrupt flag clear register
//	0x210 32  DFSDM1_JCHGR    injected channel group selection register
//	0x214 32  DFSDM1_FCR      filter control register
//	0x218 32  DFSDM1_JDATAR   data register for injected group
//	0x21C 32  DFSDM1_RDATAR   data register for the regular channel
//	0x220 32  DFSDM1_AWHTR    analog watchdog high threshold register
//	0x224 32  DFSDM1_AWLTR    analog watchdog low threshold register
//	0x228 32  DFSDM1_AWSR     analog watchdog status register
//	0x22C 32  DFSDM1_AWCFR    analog watchdog clear flag register
//	0x230 32  DFSDM1_EXMAX    Extremes detector maximum register
//	0x234 32  DFSDM1_EXMIN    Extremes detector minimum register
//	0x238 32  DFSDM1_CNVTIMR  conversion timer register
//	0x300 32  DFSDM2_CR1      control register 1
//	0x304 32  DFSDM2_CR2      control register 2
//	0x308 32  DFSDM2_ISR      interrupt and status register
//	0x30C 32  DFSDM2_ICR      interrupt flag clear register
//	0x310 32  DFSDM2_JCHGR    injected channel group selection register
//	0x314 32  DFSDM2_FCR      filter control register
//	0x318 32  DFSDM2_JDATAR   data register for injected group
//	0x31C 32  DFSDM2_RDATAR   data register for the regular channel
//	0x320 32  DFSDM2_AWHTR    analog watchdog high threshold register
//	0x324 32  DFSDM2_AWLTR    analog watchdog low threshold register
//	0x328 32  DFSDM2_AWSR     analog watchdog status register
//	0x32C 32  DFSDM2_AWCFR    analog watchdog clear flag register
//	0x330 32  DFSDM2_EXMAX    Extremes detector maximum register
//	0x334 32  DFSDM2_EXMIN    Extremes detector minimum register
//	0x338 32  DFSDM2_CNVTIMR  conversion timer register
//	0x400 32  DFSDM3_CR1      control register 1
//	0x404 32  DFSDM3_CR2      control register 2
//	0x408 32  DFSDM3_ISR      interrupt and status register
//	0x40C 32  DFSDM3_ICR      interrupt flag clear register
//	0x410 32  DFSDM3_JCHGR    injected channel group selection register
//	0x414 32  DFSDM3_FCR      filter control register
//	0x418 32  DFSDM3_JDATAR   data register for injected group
//	0x41C 32  DFSDM3_RDATAR   data register for the regular channel
//	0x420 32  DFSDM3_AWHTR    analog watchdog high threshold register
//	0x424 32  DFSDM3_AWLTR    analog watchdog low threshold register
//	0x428 32  DFSDM3_AWSR     analog watchdog status register
//	0x42C 32  DFSDM3_AWCFR    analog watchdog clear flag register
//	0x430 32  DFSDM3_EXMAX    Extremes detector maximum register
//	0x434 32  DFSDM3_EXMIN    Extremes detector minimum register
//	0x438 32  DFSDM3_CNVTIMR  conversion timer register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package dfsdm

const (
	SITP     CHCFG0R1 = 0x03 << 0  //+ SITP
	SPICKSEL CHCFG0R1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CHCFG0R1 = 0x01 << 5  //+ SCDEN
	CKABEN   CHCFG0R1 = 0x01 << 6  //+ CKABEN
	CHEN     CHCFG0R1 = 0x01 << 7  //+ CHEN
	CHINSEL  CHCFG0R1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CHCFG0R1 = 0x03 << 12 //+ DATMPX
	DATPACK  CHCFG0R1 = 0x03 << 14 //+ DATPACK
	CKOUTDIV CHCFG0R1 = 0xFF << 16 //+ CKOUTDIV
	CKOUTSRC CHCFG0R1 = 0x01 << 30 //+ CKOUTSRC
	DFSDMEN  CHCFG0R1 = 0x01 << 31 //+ DFSDMEN
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
	DTRBS  CHCFG0R2 = 0x1F << 3     //+ DTRBS
	OFFSET CHCFG0R2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   AWSCD0R = 0xFF << 0  //+ SCDT
	BKSCD  AWSCD0R = 0x0F << 12 //+ BKSCD
	AWFOSR AWSCD0R = 0x1F << 16 //+ AWFOSR
	AWFORD AWSCD0R = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CHWDAT0R = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CHDATIN0R = 0xFFFF << 0  //+ INDAT0
	INDAT1 CHDATIN0R = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	SITP     CHCFG1R1 = 0x03 << 0  //+ SITP
	SPICKSEL CHCFG1R1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CHCFG1R1 = 0x01 << 5  //+ SCDEN
	CKABEN   CHCFG1R1 = 0x01 << 6  //+ CKABEN
	CHEN     CHCFG1R1 = 0x01 << 7  //+ CHEN
	CHINSEL  CHCFG1R1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CHCFG1R1 = 0x03 << 12 //+ DATMPX
	DATPACK  CHCFG1R1 = 0x03 << 14 //+ DATPACK
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
	DTRBS  CHCFG1R2 = 0x1F << 3     //+ DTRBS
	OFFSET CHCFG1R2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   AWSCD1R = 0xFF << 0  //+ SCDT
	BKSCD  AWSCD1R = 0x0F << 12 //+ BKSCD
	AWFOSR AWSCD1R = 0x1F << 16 //+ AWFOSR
	AWFORD AWSCD1R = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CHWDAT1R = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CHDATIN1R = 0xFFFF << 0  //+ INDAT0
	INDAT1 CHDATIN1R = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	SITP     CHCFG2R1 = 0x03 << 0  //+ SITP
	SPICKSEL CHCFG2R1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CHCFG2R1 = 0x01 << 5  //+ SCDEN
	CKABEN   CHCFG2R1 = 0x01 << 6  //+ CKABEN
	CHEN     CHCFG2R1 = 0x01 << 7  //+ CHEN
	CHINSEL  CHCFG2R1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CHCFG2R1 = 0x03 << 12 //+ DATMPX
	DATPACK  CHCFG2R1 = 0x03 << 14 //+ DATPACK
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
	DTRBS  CHCFG2R2 = 0x1F << 3     //+ DTRBS
	OFFSET CHCFG2R2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   AWSCD2R = 0xFF << 0  //+ SCDT
	BKSCD  AWSCD2R = 0x0F << 12 //+ BKSCD
	AWFOSR AWSCD2R = 0x1F << 16 //+ AWFOSR
	AWFORD AWSCD2R = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CHWDAT2R = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CHDATIN2R = 0xFFFF << 0  //+ INDAT0
	INDAT1 CHDATIN2R = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	SITP     CHCFG3R1 = 0x03 << 0  //+ SITP
	SPICKSEL CHCFG3R1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CHCFG3R1 = 0x01 << 5  //+ SCDEN
	CKABEN   CHCFG3R1 = 0x01 << 6  //+ CKABEN
	CHEN     CHCFG3R1 = 0x01 << 7  //+ CHEN
	CHINSEL  CHCFG3R1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CHCFG3R1 = 0x03 << 12 //+ DATMPX
	DATPACK  CHCFG3R1 = 0x03 << 14 //+ DATPACK
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
	DTRBS  CHCFG3R2 = 0x1F << 3     //+ DTRBS
	OFFSET CHCFG3R2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   AWSCD3R = 0xFF << 0  //+ SCDT
	BKSCD  AWSCD3R = 0x0F << 12 //+ BKSCD
	AWFOSR AWSCD3R = 0x1F << 16 //+ AWFOSR
	AWFORD AWSCD3R = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CHWDAT3R = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CHDATIN3R = 0xFFFF << 0  //+ INDAT0
	INDAT1 CHDATIN3R = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	SITP     CHCFG4R1 = 0x03 << 0  //+ SITP
	SPICKSEL CHCFG4R1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CHCFG4R1 = 0x01 << 5  //+ SCDEN
	CKABEN   CHCFG4R1 = 0x01 << 6  //+ CKABEN
	CHEN     CHCFG4R1 = 0x01 << 7  //+ CHEN
	CHINSEL  CHCFG4R1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CHCFG4R1 = 0x03 << 12 //+ DATMPX
	DATPACK  CHCFG4R1 = 0x03 << 14 //+ DATPACK
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
	DTRBS  CHCFG4R2 = 0x1F << 3     //+ DTRBS
	OFFSET CHCFG4R2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   AWSCD4R = 0xFF << 0  //+ SCDT
	BKSCD  AWSCD4R = 0x0F << 12 //+ BKSCD
	AWFOSR AWSCD4R = 0x1F << 16 //+ AWFOSR
	AWFORD AWSCD4R = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CHWDAT4R = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CHDATIN4R = 0xFFFF << 0  //+ INDAT0
	INDAT1 CHDATIN4R = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	SITP     CHCFG5R1 = 0x03 << 0  //+ SITP
	SPICKSEL CHCFG5R1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CHCFG5R1 = 0x01 << 5  //+ SCDEN
	CKABEN   CHCFG5R1 = 0x01 << 6  //+ CKABEN
	CHEN     CHCFG5R1 = 0x01 << 7  //+ CHEN
	CHINSEL  CHCFG5R1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CHCFG5R1 = 0x03 << 12 //+ DATMPX
	DATPACK  CHCFG5R1 = 0x03 << 14 //+ DATPACK
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
	DTRBS  CHCFG5R2 = 0x1F << 3     //+ DTRBS
	OFFSET CHCFG5R2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   AWSCD5R = 0xFF << 0  //+ SCDT
	BKSCD  AWSCD5R = 0x0F << 12 //+ BKSCD
	AWFOSR AWSCD5R = 0x1F << 16 //+ AWFOSR
	AWFORD AWSCD5R = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CHWDAT5R = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CHDATIN5R = 0xFFFF << 0  //+ INDAT0
	INDAT1 CHDATIN5R = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	SITP     CHCFG6R1 = 0x03 << 0  //+ SITP
	SPICKSEL CHCFG6R1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CHCFG6R1 = 0x01 << 5  //+ SCDEN
	CKABEN   CHCFG6R1 = 0x01 << 6  //+ CKABEN
	CHEN     CHCFG6R1 = 0x01 << 7  //+ CHEN
	CHINSEL  CHCFG6R1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CHCFG6R1 = 0x03 << 12 //+ DATMPX
	DATPACK  CHCFG6R1 = 0x03 << 14 //+ DATPACK
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
	DTRBS  CHCFG6R2 = 0x1F << 3     //+ DTRBS
	OFFSET CHCFG6R2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   AWSCD6R = 0xFF << 0  //+ SCDT
	BKSCD  AWSCD6R = 0x0F << 12 //+ BKSCD
	AWFOSR AWSCD6R = 0x1F << 16 //+ AWFOSR
	AWFORD AWSCD6R = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CHWDAT6R = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CHDATIN6R = 0xFFFF << 0  //+ INDAT0
	INDAT1 CHDATIN6R = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	SITP     CHCFG7R1 = 0x03 << 0  //+ SITP
	SPICKSEL CHCFG7R1 = 0x03 << 2  //+ SPICKSEL
	SCDEN    CHCFG7R1 = 0x01 << 5  //+ SCDEN
	CKABEN   CHCFG7R1 = 0x01 << 6  //+ CKABEN
	CHEN     CHCFG7R1 = 0x01 << 7  //+ CHEN
	CHINSEL  CHCFG7R1 = 0x01 << 8  //+ CHINSEL
	DATMPX   CHCFG7R1 = 0x03 << 12 //+ DATMPX
	DATPACK  CHCFG7R1 = 0x03 << 14 //+ DATPACK
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
	DTRBS  CHCFG7R2 = 0x1F << 3     //+ DTRBS
	OFFSET CHCFG7R2 = 0xFFFFFF << 8 //+ OFFSET
)

const (
	DTRBSn  = 3
	OFFSETn = 8
)

const (
	SCDT   AWSCD7R = 0xFF << 0  //+ SCDT
	BKSCD  AWSCD7R = 0x0F << 12 //+ BKSCD
	AWFOSR AWSCD7R = 0x1F << 16 //+ AWFOSR
	AWFORD AWSCD7R = 0x03 << 22 //+ AWFORD
)

const (
	SCDTn   = 0
	BKSCDn  = 12
	AWFOSRn = 16
	AWFORDn = 22
)

const (
	WDATA CHWDAT7R = 0xFFFF << 0 //+ WDATA
)

const (
	WDATAn = 0
)

const (
	INDAT0 CHDATIN7R = 0xFFFF << 0  //+ INDAT0
	INDAT1 CHDATIN7R = 0xFFFF << 16 //+ INDAT1
)

const (
	INDAT0n = 0
	INDAT1n = 16
)

const (
	DFEN     DFSDM0_CR1 = 0x01 << 0  //+ DFSDM enable
	JSWSTART DFSDM0_CR1 = 0x01 << 1  //+ Start a conversion of the injected group of channels
	JSYNC    DFSDM0_CR1 = 0x01 << 3  //+ Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
	JSCAN    DFSDM0_CR1 = 0x01 << 4  //+ Scanning conversion mode for injected conversions
	JDMAEN   DFSDM0_CR1 = 0x01 << 5  //+ DMA channel enabled to read data for the injected channel group
	JEXTSEL  DFSDM0_CR1 = 0x07 << 8  //+ Trigger signal selection for launching injected conversions
	JEXTEN   DFSDM0_CR1 = 0x03 << 13 //+ Trigger enable and trigger edge selection for injected conversions
	RSWSTART DFSDM0_CR1 = 0x01 << 17 //+ Software start of a conversion on the regular channel
	RCONT    DFSDM0_CR1 = 0x01 << 18 //+ Continuous mode selection for regular conversions
	RSYNC    DFSDM0_CR1 = 0x01 << 19 //+ Launch regular conversion synchronously with DFSDM0
	RDMAEN   DFSDM0_CR1 = 0x01 << 21 //+ DMA channel enabled to read data for the regular conversion
	RCH      DFSDM0_CR1 = 0x07 << 24 //+ Regular channel selection
	FAST     DFSDM0_CR1 = 0x01 << 29 //+ Fast conversion mode selection for regular conversions
	AWFSEL   DFSDM0_CR1 = 0x01 << 30 //+ Analog watchdog fast mode select
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
	JEOCIE DFSDM0_CR2 = 0x01 << 0  //+ Injected end of conversion interrupt enable
	REOCIE DFSDM0_CR2 = 0x01 << 1  //+ Regular end of conversion interrupt enable
	JOVRIE DFSDM0_CR2 = 0x01 << 2  //+ Injected data overrun interrupt enable
	ROVRIE DFSDM0_CR2 = 0x01 << 3  //+ Regular data overrun interrupt enable
	AWDIE  DFSDM0_CR2 = 0x01 << 4  //+ Analog watchdog interrupt enable
	SCDIE  DFSDM0_CR2 = 0x01 << 5  //+ Short-circuit detector interrupt enable
	CKABIE DFSDM0_CR2 = 0x01 << 6  //+ Clock absence interrupt enable
	EXCH   DFSDM0_CR2 = 0xFF << 8  //+ Extremes detector channel selection
	AWDCH  DFSDM0_CR2 = 0xFF << 16 //+ Analog watchdog channel selection
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
	JEOCF DFSDM0_ISR = 0x01 << 0  //+ End of injected conversion flag
	REOCF DFSDM0_ISR = 0x01 << 1  //+ End of regular conversion flag
	JOVRF DFSDM0_ISR = 0x01 << 2  //+ Injected conversion overrun flag
	ROVRF DFSDM0_ISR = 0x01 << 3  //+ Regular conversion overrun flag
	AWDF  DFSDM0_ISR = 0x01 << 4  //+ Analog watchdog
	JCIP  DFSDM0_ISR = 0x01 << 13 //+ Injected conversion in progress status
	RCIP  DFSDM0_ISR = 0x01 << 14 //+ Regular conversion in progress status
	CKABF DFSDM0_ISR = 0xFF << 16 //+ Clock absence flag
	SCDF  DFSDM0_ISR = 0xFF << 24 //+ short-circuit detector flag
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
	CLRJOVRF DFSDM0_ICR = 0x01 << 2  //+ Clear the injected conversion overrun flag
	CLRROVRF DFSDM0_ICR = 0x01 << 3  //+ Clear the regular conversion overrun flag
	CLRCKABF DFSDM0_ICR = 0xFF << 16 //+ Clear the clock absence flag
	CLRSCDF  DFSDM0_ICR = 0xFF << 24 //+ Clear the short-circuit detector flag
)

const (
	CLRJOVRFn = 2
	CLRROVRFn = 3
	CLRCKABFn = 16
	CLRSCDFn  = 24
)

const (
	JCHG DFSDM0_JCHGR = 0xFF << 0 //+ Injected channel group selection
)

const (
	JCHGn = 0
)

const (
	IOSR DFSDM0_FCR = 0xFF << 0   //+ Integrator oversampling ratio (averaging length)
	FOSR DFSDM0_FCR = 0x3FF << 16 //+ Sinc filter oversampling ratio (decimation rate)
	FORD DFSDM0_FCR = 0x07 << 29  //+ Sinc filter order
)

const (
	IOSRn = 0
	FOSRn = 16
	FORDn = 29
)

const (
	JDATACH DFSDM0_JDATAR = 0x07 << 0     //+ Injected channel most recently converted
	JDATA   DFSDM0_JDATAR = 0xFFFFFF << 8 //+ Injected group conversion data
)

const (
	JDATACHn = 0
	JDATAn   = 8
)

const (
	RDATACH DFSDM0_RDATAR = 0x07 << 0     //+ Regular channel most recently converted
	RPEND   DFSDM0_RDATAR = 0x01 << 4     //+ Regular channel pending data
	RDATA   DFSDM0_RDATAR = 0xFFFFFF << 8 //+ Regular channel conversion data
)

const (
	RDATACHn = 0
	RPENDn   = 4
	RDATAn   = 8
)

const (
	BKAWH DFSDM0_AWHTR = 0x0F << 0     //+ Break signal assignment to analog watchdog high threshold event
	AWHT  DFSDM0_AWHTR = 0xFFFFFF << 8 //+ Analog watchdog high threshold
)

const (
	BKAWHn = 0
	AWHTn  = 8
)

const (
	BKAWL DFSDM0_AWLTR = 0x0F << 0     //+ Break signal assignment to analog watchdog low threshold event
	AWLT  DFSDM0_AWLTR = 0xFFFFFF << 8 //+ Analog watchdog low threshold
)

const (
	BKAWLn = 0
	AWLTn  = 8
)

const (
	AWLTF DFSDM0_AWSR = 0xFF << 0 //+ Analog watchdog low threshold flag
	AWHTF DFSDM0_AWSR = 0xFF << 8 //+ Analog watchdog high threshold flag
)

const (
	AWLTFn = 0
	AWHTFn = 8
)

const (
	CLRAWLTF DFSDM0_AWCFR = 0xFF << 0 //+ Clear the analog watchdog low threshold flag
	CLRAWHTF DFSDM0_AWCFR = 0xFF << 8 //+ Clear the analog watchdog high threshold flag
)

const (
	CLRAWLTFn = 0
	CLRAWHTFn = 8
)

const (
	EXMAXCH DFSDM0_EXMAX = 0x07 << 0     //+ Extremes detector maximum data channel
	EXMAX   DFSDM0_EXMAX = 0xFFFFFF << 8 //+ Extremes detector maximum value
)

const (
	EXMAXCHn = 0
	EXMAXn   = 8
)

const (
	EXMINCH DFSDM0_EXMIN = 0x07 << 0     //+ Extremes detector minimum data channel
	EXMIN   DFSDM0_EXMIN = 0xFFFFFF << 8 //+ EXMIN
)

const (
	EXMINCHn = 0
	EXMINn   = 8
)

const (
	CNVCNT DFSDM0_CNVTIMR = 0xFFFFFFF << 4 //+ 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
)

const (
	CNVCNTn = 4
)

const (
	DFEN     DFSDM1_CR1 = 0x01 << 0  //+ DFSDM enable
	JSWSTART DFSDM1_CR1 = 0x01 << 1  //+ Start a conversion of the injected group of channels
	JSYNC    DFSDM1_CR1 = 0x01 << 3  //+ Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
	JSCAN    DFSDM1_CR1 = 0x01 << 4  //+ Scanning conversion mode for injected conversions
	JDMAEN   DFSDM1_CR1 = 0x01 << 5  //+ DMA channel enabled to read data for the injected channel group
	JEXTSEL  DFSDM1_CR1 = 0x07 << 8  //+ Trigger signal selection for launching injected conversions
	JEXTEN   DFSDM1_CR1 = 0x03 << 13 //+ Trigger enable and trigger edge selection for injected conversions
	RSWSTART DFSDM1_CR1 = 0x01 << 17 //+ Software start of a conversion on the regular channel
	RCONT    DFSDM1_CR1 = 0x01 << 18 //+ Continuous mode selection for regular conversions
	RSYNC    DFSDM1_CR1 = 0x01 << 19 //+ Launch regular conversion synchronously with DFSDM0
	RDMAEN   DFSDM1_CR1 = 0x01 << 21 //+ DMA channel enabled to read data for the regular conversion
	RCH      DFSDM1_CR1 = 0x07 << 24 //+ Regular channel selection
	FAST     DFSDM1_CR1 = 0x01 << 29 //+ Fast conversion mode selection for regular conversions
	AWFSEL   DFSDM1_CR1 = 0x01 << 30 //+ Analog watchdog fast mode select
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
	JEOCIE DFSDM1_CR2 = 0x01 << 0  //+ Injected end of conversion interrupt enable
	REOCIE DFSDM1_CR2 = 0x01 << 1  //+ Regular end of conversion interrupt enable
	JOVRIE DFSDM1_CR2 = 0x01 << 2  //+ Injected data overrun interrupt enable
	ROVRIE DFSDM1_CR2 = 0x01 << 3  //+ Regular data overrun interrupt enable
	AWDIE  DFSDM1_CR2 = 0x01 << 4  //+ Analog watchdog interrupt enable
	SCDIE  DFSDM1_CR2 = 0x01 << 5  //+ Short-circuit detector interrupt enable
	CKABIE DFSDM1_CR2 = 0x01 << 6  //+ Clock absence interrupt enable
	EXCH   DFSDM1_CR2 = 0xFF << 8  //+ Extremes detector channel selection
	AWDCH  DFSDM1_CR2 = 0xFF << 16 //+ Analog watchdog channel selection
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
	JEOCF DFSDM1_ISR = 0x01 << 0  //+ End of injected conversion flag
	REOCF DFSDM1_ISR = 0x01 << 1  //+ End of regular conversion flag
	JOVRF DFSDM1_ISR = 0x01 << 2  //+ Injected conversion overrun flag
	ROVRF DFSDM1_ISR = 0x01 << 3  //+ Regular conversion overrun flag
	AWDF  DFSDM1_ISR = 0x01 << 4  //+ Analog watchdog
	JCIP  DFSDM1_ISR = 0x01 << 13 //+ Injected conversion in progress status
	RCIP  DFSDM1_ISR = 0x01 << 14 //+ Regular conversion in progress status
	CKABF DFSDM1_ISR = 0xFF << 16 //+ Clock absence flag
	SCDF  DFSDM1_ISR = 0xFF << 24 //+ short-circuit detector flag
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
	CLRJOVRF DFSDM1_ICR = 0x01 << 2  //+ Clear the injected conversion overrun flag
	CLRROVRF DFSDM1_ICR = 0x01 << 3  //+ Clear the regular conversion overrun flag
	CLRCKABF DFSDM1_ICR = 0xFF << 16 //+ Clear the clock absence flag
	CLRSCDF  DFSDM1_ICR = 0xFF << 24 //+ Clear the short-circuit detector flag
)

const (
	CLRJOVRFn = 2
	CLRROVRFn = 3
	CLRCKABFn = 16
	CLRSCDFn  = 24
)

const (
	JCHG DFSDM1_JCHGR = 0xFF << 0 //+ Injected channel group selection
)

const (
	JCHGn = 0
)

const (
	IOSR DFSDM1_FCR = 0xFF << 0   //+ Integrator oversampling ratio (averaging length)
	FOSR DFSDM1_FCR = 0x3FF << 16 //+ Sinc filter oversampling ratio (decimation rate)
	FORD DFSDM1_FCR = 0x07 << 29  //+ Sinc filter order
)

const (
	IOSRn = 0
	FOSRn = 16
	FORDn = 29
)

const (
	JDATACH DFSDM1_JDATAR = 0x07 << 0     //+ Injected channel most recently converted
	JDATA   DFSDM1_JDATAR = 0xFFFFFF << 8 //+ Injected group conversion data
)

const (
	JDATACHn = 0
	JDATAn   = 8
)

const (
	RDATACH DFSDM1_RDATAR = 0x07 << 0     //+ Regular channel most recently converted
	RPEND   DFSDM1_RDATAR = 0x01 << 4     //+ Regular channel pending data
	RDATA   DFSDM1_RDATAR = 0xFFFFFF << 8 //+ Regular channel conversion data
)

const (
	RDATACHn = 0
	RPENDn   = 4
	RDATAn   = 8
)

const (
	BKAWH DFSDM1_AWHTR = 0x0F << 0     //+ Break signal assignment to analog watchdog high threshold event
	AWHT  DFSDM1_AWHTR = 0xFFFFFF << 8 //+ Analog watchdog high threshold
)

const (
	BKAWHn = 0
	AWHTn  = 8
)

const (
	BKAWL DFSDM1_AWLTR = 0x0F << 0     //+ Break signal assignment to analog watchdog low threshold event
	AWLT  DFSDM1_AWLTR = 0xFFFFFF << 8 //+ Analog watchdog low threshold
)

const (
	BKAWLn = 0
	AWLTn  = 8
)

const (
	AWLTF DFSDM1_AWSR = 0xFF << 0 //+ Analog watchdog low threshold flag
	AWHTF DFSDM1_AWSR = 0xFF << 8 //+ Analog watchdog high threshold flag
)

const (
	AWLTFn = 0
	AWHTFn = 8
)

const (
	CLRAWLTF DFSDM1_AWCFR = 0xFF << 0 //+ Clear the analog watchdog low threshold flag
	CLRAWHTF DFSDM1_AWCFR = 0xFF << 8 //+ Clear the analog watchdog high threshold flag
)

const (
	CLRAWLTFn = 0
	CLRAWHTFn = 8
)

const (
	EXMAXCH DFSDM1_EXMAX = 0x07 << 0     //+ Extremes detector maximum data channel
	EXMAX   DFSDM1_EXMAX = 0xFFFFFF << 8 //+ Extremes detector maximum value
)

const (
	EXMAXCHn = 0
	EXMAXn   = 8
)

const (
	EXMINCH DFSDM1_EXMIN = 0x07 << 0     //+ Extremes detector minimum data channel
	EXMIN   DFSDM1_EXMIN = 0xFFFFFF << 8 //+ EXMIN
)

const (
	EXMINCHn = 0
	EXMINn   = 8
)

const (
	CNVCNT DFSDM1_CNVTIMR = 0xFFFFFFF << 4 //+ 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
)

const (
	CNVCNTn = 4
)

const (
	DFEN     DFSDM2_CR1 = 0x01 << 0  //+ DFSDM enable
	JSWSTART DFSDM2_CR1 = 0x01 << 1  //+ Start a conversion of the injected group of channels
	JSYNC    DFSDM2_CR1 = 0x01 << 3  //+ Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
	JSCAN    DFSDM2_CR1 = 0x01 << 4  //+ Scanning conversion mode for injected conversions
	JDMAEN   DFSDM2_CR1 = 0x01 << 5  //+ DMA channel enabled to read data for the injected channel group
	JEXTSEL  DFSDM2_CR1 = 0x07 << 8  //+ Trigger signal selection for launching injected conversions
	JEXTEN   DFSDM2_CR1 = 0x03 << 13 //+ Trigger enable and trigger edge selection for injected conversions
	RSWSTART DFSDM2_CR1 = 0x01 << 17 //+ Software start of a conversion on the regular channel
	RCONT    DFSDM2_CR1 = 0x01 << 18 //+ Continuous mode selection for regular conversions
	RSYNC    DFSDM2_CR1 = 0x01 << 19 //+ Launch regular conversion synchronously with DFSDM0
	RDMAEN   DFSDM2_CR1 = 0x01 << 21 //+ DMA channel enabled to read data for the regular conversion
	RCH      DFSDM2_CR1 = 0x07 << 24 //+ Regular channel selection
	FAST     DFSDM2_CR1 = 0x01 << 29 //+ Fast conversion mode selection for regular conversions
	AWFSEL   DFSDM2_CR1 = 0x01 << 30 //+ Analog watchdog fast mode select
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
	JEOCIE DFSDM2_CR2 = 0x01 << 0  //+ Injected end of conversion interrupt enable
	REOCIE DFSDM2_CR2 = 0x01 << 1  //+ Regular end of conversion interrupt enable
	JOVRIE DFSDM2_CR2 = 0x01 << 2  //+ Injected data overrun interrupt enable
	ROVRIE DFSDM2_CR2 = 0x01 << 3  //+ Regular data overrun interrupt enable
	AWDIE  DFSDM2_CR2 = 0x01 << 4  //+ Analog watchdog interrupt enable
	SCDIE  DFSDM2_CR2 = 0x01 << 5  //+ Short-circuit detector interrupt enable
	CKABIE DFSDM2_CR2 = 0x01 << 6  //+ Clock absence interrupt enable
	EXCH   DFSDM2_CR2 = 0xFF << 8  //+ Extremes detector channel selection
	AWDCH  DFSDM2_CR2 = 0xFF << 16 //+ Analog watchdog channel selection
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
	JEOCF DFSDM2_ISR = 0x01 << 0  //+ End of injected conversion flag
	REOCF DFSDM2_ISR = 0x01 << 1  //+ End of regular conversion flag
	JOVRF DFSDM2_ISR = 0x01 << 2  //+ Injected conversion overrun flag
	ROVRF DFSDM2_ISR = 0x01 << 3  //+ Regular conversion overrun flag
	AWDF  DFSDM2_ISR = 0x01 << 4  //+ Analog watchdog
	JCIP  DFSDM2_ISR = 0x01 << 13 //+ Injected conversion in progress status
	RCIP  DFSDM2_ISR = 0x01 << 14 //+ Regular conversion in progress status
	CKABF DFSDM2_ISR = 0xFF << 16 //+ Clock absence flag
	SCDF  DFSDM2_ISR = 0xFF << 24 //+ short-circuit detector flag
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
	CLRJOVRF DFSDM2_ICR = 0x01 << 2  //+ Clear the injected conversion overrun flag
	CLRROVRF DFSDM2_ICR = 0x01 << 3  //+ Clear the regular conversion overrun flag
	CLRCKABF DFSDM2_ICR = 0xFF << 16 //+ Clear the clock absence flag
	CLRSCDF  DFSDM2_ICR = 0xFF << 24 //+ Clear the short-circuit detector flag
)

const (
	CLRJOVRFn = 2
	CLRROVRFn = 3
	CLRCKABFn = 16
	CLRSCDFn  = 24
)

const (
	JCHG DFSDM2_JCHGR = 0xFF << 0 //+ Injected channel group selection
)

const (
	JCHGn = 0
)

const (
	IOSR DFSDM2_FCR = 0xFF << 0   //+ Integrator oversampling ratio (averaging length)
	FOSR DFSDM2_FCR = 0x3FF << 16 //+ Sinc filter oversampling ratio (decimation rate)
	FORD DFSDM2_FCR = 0x07 << 29  //+ Sinc filter order
)

const (
	IOSRn = 0
	FOSRn = 16
	FORDn = 29
)

const (
	JDATACH DFSDM2_JDATAR = 0x07 << 0     //+ Injected channel most recently converted
	JDATA   DFSDM2_JDATAR = 0xFFFFFF << 8 //+ Injected group conversion data
)

const (
	JDATACHn = 0
	JDATAn   = 8
)

const (
	RDATACH DFSDM2_RDATAR = 0x07 << 0     //+ Regular channel most recently converted
	RPEND   DFSDM2_RDATAR = 0x01 << 4     //+ Regular channel pending data
	RDATA   DFSDM2_RDATAR = 0xFFFFFF << 8 //+ Regular channel conversion data
)

const (
	RDATACHn = 0
	RPENDn   = 4
	RDATAn   = 8
)

const (
	BKAWH DFSDM2_AWHTR = 0x0F << 0     //+ Break signal assignment to analog watchdog high threshold event
	AWHT  DFSDM2_AWHTR = 0xFFFFFF << 8 //+ Analog watchdog high threshold
)

const (
	BKAWHn = 0
	AWHTn  = 8
)

const (
	BKAWL DFSDM2_AWLTR = 0x0F << 0     //+ Break signal assignment to analog watchdog low threshold event
	AWLT  DFSDM2_AWLTR = 0xFFFFFF << 8 //+ Analog watchdog low threshold
)

const (
	BKAWLn = 0
	AWLTn  = 8
)

const (
	AWLTF DFSDM2_AWSR = 0xFF << 0 //+ Analog watchdog low threshold flag
	AWHTF DFSDM2_AWSR = 0xFF << 8 //+ Analog watchdog high threshold flag
)

const (
	AWLTFn = 0
	AWHTFn = 8
)

const (
	CLRAWLTF DFSDM2_AWCFR = 0xFF << 0 //+ Clear the analog watchdog low threshold flag
	CLRAWHTF DFSDM2_AWCFR = 0xFF << 8 //+ Clear the analog watchdog high threshold flag
)

const (
	CLRAWLTFn = 0
	CLRAWHTFn = 8
)

const (
	EXMAXCH DFSDM2_EXMAX = 0x07 << 0     //+ Extremes detector maximum data channel
	EXMAX   DFSDM2_EXMAX = 0xFFFFFF << 8 //+ Extremes detector maximum value
)

const (
	EXMAXCHn = 0
	EXMAXn   = 8
)

const (
	EXMINCH DFSDM2_EXMIN = 0x07 << 0     //+ Extremes detector minimum data channel
	EXMIN   DFSDM2_EXMIN = 0xFFFFFF << 8 //+ EXMIN
)

const (
	EXMINCHn = 0
	EXMINn   = 8
)

const (
	CNVCNT DFSDM2_CNVTIMR = 0xFFFFFFF << 4 //+ 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
)

const (
	CNVCNTn = 4
)

const (
	DFEN     DFSDM3_CR1 = 0x01 << 0  //+ DFSDM enable
	JSWSTART DFSDM3_CR1 = 0x01 << 1  //+ Start a conversion of the injected group of channels
	JSYNC    DFSDM3_CR1 = 0x01 << 3  //+ Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger
	JSCAN    DFSDM3_CR1 = 0x01 << 4  //+ Scanning conversion mode for injected conversions
	JDMAEN   DFSDM3_CR1 = 0x01 << 5  //+ DMA channel enabled to read data for the injected channel group
	JEXTSEL  DFSDM3_CR1 = 0x07 << 8  //+ Trigger signal selection for launching injected conversions
	JEXTEN   DFSDM3_CR1 = 0x03 << 13 //+ Trigger enable and trigger edge selection for injected conversions
	RSWSTART DFSDM3_CR1 = 0x01 << 17 //+ Software start of a conversion on the regular channel
	RCONT    DFSDM3_CR1 = 0x01 << 18 //+ Continuous mode selection for regular conversions
	RSYNC    DFSDM3_CR1 = 0x01 << 19 //+ Launch regular conversion synchronously with DFSDM0
	RDMAEN   DFSDM3_CR1 = 0x01 << 21 //+ DMA channel enabled to read data for the regular conversion
	RCH      DFSDM3_CR1 = 0x07 << 24 //+ Regular channel selection
	FAST     DFSDM3_CR1 = 0x01 << 29 //+ Fast conversion mode selection for regular conversions
	AWFSEL   DFSDM3_CR1 = 0x01 << 30 //+ Analog watchdog fast mode select
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
	JEOCIE DFSDM3_CR2 = 0x01 << 0  //+ Injected end of conversion interrupt enable
	REOCIE DFSDM3_CR2 = 0x01 << 1  //+ Regular end of conversion interrupt enable
	JOVRIE DFSDM3_CR2 = 0x01 << 2  //+ Injected data overrun interrupt enable
	ROVRIE DFSDM3_CR2 = 0x01 << 3  //+ Regular data overrun interrupt enable
	AWDIE  DFSDM3_CR2 = 0x01 << 4  //+ Analog watchdog interrupt enable
	SCDIE  DFSDM3_CR2 = 0x01 << 5  //+ Short-circuit detector interrupt enable
	CKABIE DFSDM3_CR2 = 0x01 << 6  //+ Clock absence interrupt enable
	EXCH   DFSDM3_CR2 = 0xFF << 8  //+ Extremes detector channel selection
	AWDCH  DFSDM3_CR2 = 0xFF << 16 //+ Analog watchdog channel selection
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
	JEOCF DFSDM3_ISR = 0x01 << 0  //+ End of injected conversion flag
	REOCF DFSDM3_ISR = 0x01 << 1  //+ End of regular conversion flag
	JOVRF DFSDM3_ISR = 0x01 << 2  //+ Injected conversion overrun flag
	ROVRF DFSDM3_ISR = 0x01 << 3  //+ Regular conversion overrun flag
	AWDF  DFSDM3_ISR = 0x01 << 4  //+ Analog watchdog
	JCIP  DFSDM3_ISR = 0x01 << 13 //+ Injected conversion in progress status
	RCIP  DFSDM3_ISR = 0x01 << 14 //+ Regular conversion in progress status
	CKABF DFSDM3_ISR = 0xFF << 16 //+ Clock absence flag
	SCDF  DFSDM3_ISR = 0xFF << 24 //+ short-circuit detector flag
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
	CLRJOVRF DFSDM3_ICR = 0x01 << 2  //+ Clear the injected conversion overrun flag
	CLRROVRF DFSDM3_ICR = 0x01 << 3  //+ Clear the regular conversion overrun flag
	CLRCKABF DFSDM3_ICR = 0xFF << 16 //+ Clear the clock absence flag
	CLRSCDF  DFSDM3_ICR = 0xFF << 24 //+ Clear the short-circuit detector flag
)

const (
	CLRJOVRFn = 2
	CLRROVRFn = 3
	CLRCKABFn = 16
	CLRSCDFn  = 24
)

const (
	JCHG DFSDM3_JCHGR = 0xFF << 0 //+ Injected channel group selection
)

const (
	JCHGn = 0
)

const (
	IOSR DFSDM3_FCR = 0xFF << 0   //+ Integrator oversampling ratio (averaging length)
	FOSR DFSDM3_FCR = 0x3FF << 16 //+ Sinc filter oversampling ratio (decimation rate)
	FORD DFSDM3_FCR = 0x07 << 29  //+ Sinc filter order
)

const (
	IOSRn = 0
	FOSRn = 16
	FORDn = 29
)

const (
	JDATACH DFSDM3_JDATAR = 0x07 << 0     //+ Injected channel most recently converted
	JDATA   DFSDM3_JDATAR = 0xFFFFFF << 8 //+ Injected group conversion data
)

const (
	JDATACHn = 0
	JDATAn   = 8
)

const (
	RDATACH DFSDM3_RDATAR = 0x07 << 0     //+ Regular channel most recently converted
	RPEND   DFSDM3_RDATAR = 0x01 << 4     //+ Regular channel pending data
	RDATA   DFSDM3_RDATAR = 0xFFFFFF << 8 //+ Regular channel conversion data
)

const (
	RDATACHn = 0
	RPENDn   = 4
	RDATAn   = 8
)

const (
	BKAWH DFSDM3_AWHTR = 0x0F << 0     //+ Break signal assignment to analog watchdog high threshold event
	AWHT  DFSDM3_AWHTR = 0xFFFFFF << 8 //+ Analog watchdog high threshold
)

const (
	BKAWHn = 0
	AWHTn  = 8
)

const (
	BKAWL DFSDM3_AWLTR = 0x0F << 0     //+ Break signal assignment to analog watchdog low threshold event
	AWLT  DFSDM3_AWLTR = 0xFFFFFF << 8 //+ Analog watchdog low threshold
)

const (
	BKAWLn = 0
	AWLTn  = 8
)

const (
	AWLTF DFSDM3_AWSR = 0xFF << 0 //+ Analog watchdog low threshold flag
	AWHTF DFSDM3_AWSR = 0xFF << 8 //+ Analog watchdog high threshold flag
)

const (
	AWLTFn = 0
	AWHTFn = 8
)

const (
	CLRAWLTF DFSDM3_AWCFR = 0xFF << 0 //+ Clear the analog watchdog low threshold flag
	CLRAWHTF DFSDM3_AWCFR = 0xFF << 8 //+ Clear the analog watchdog high threshold flag
)

const (
	CLRAWLTFn = 0
	CLRAWHTFn = 8
)

const (
	EXMAXCH DFSDM3_EXMAX = 0x07 << 0     //+ Extremes detector maximum data channel
	EXMAX   DFSDM3_EXMAX = 0xFFFFFF << 8 //+ Extremes detector maximum value
)

const (
	EXMAXCHn = 0
	EXMAXn   = 8
)

const (
	EXMINCH DFSDM3_EXMIN = 0x07 << 0     //+ Extremes detector minimum data channel
	EXMIN   DFSDM3_EXMIN = 0xFFFFFF << 8 //+ EXMIN
)

const (
	EXMINCHn = 0
	EXMINn   = 8
)

const (
	CNVCNT DFSDM3_CNVTIMR = 0xFFFFFFF << 4 //+ 28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN
)

const (
	CNVCNTn = 4
)
