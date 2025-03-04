// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f407

// Package sai provides access to the registers of the SAI peripheral.
//
// Instances:
//
//	SAI1  SAI1_BASE  -  -  Serial audio interface
//
// Registers:
//
//	0x004 32  SAI_ACR1    SAI AConfiguration register 1
//	0x008 32  SAI_ACR2    SAI AConfiguration register 2
//	0x00C 32  SAI_AFRCR   SAI AFrame configuration register
//	0x010 32  SAI_ASLOTR  SAI ASlot register
//	0x014 32  SAI_AIM     SAI AInterrupt mask register2
//	0x018 32  SAI_ASR     SAI AStatus register
//	0x01C 32  SAI_ACLRFR  SAI AClear flag register
//	0x020 32  SAI_ADR     SAI AData register
//	0x024 32  SAI_BCR1    SAI BConfiguration register 1
//	0x028 32  SAI_BCR2    SAI BConfiguration register 2
//	0x02C 32  SAI_BFRCR   SAI BFrame configuration register
//	0x030 32  SAI_BSLOTR  SAI BSlot register
//	0x034 32  SAI_BIM     SAI BInterrupt mask register2
//	0x038 32  SAI_BSR     SAI BStatus register
//	0x03C 32  SAI_BCLRFR  SAI BClear flag register
//	0x040 32  SAI_BDR     SAI BData register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package sai

const (
	MODE     SAI_ACR1 = 0x03 << 0  //+ Audio block mode
	PRTCFG   SAI_ACR1 = 0x03 << 2  //+ Protocol configuration
	DS       SAI_ACR1 = 0x07 << 5  //+ Data size
	LSBFIRST SAI_ACR1 = 0x01 << 8  //+ Least significant bit first
	CKSTR    SAI_ACR1 = 0x01 << 9  //+ Clock strobing edge
	SYNCEN   SAI_ACR1 = 0x03 << 10 //+ Synchronization enable
	MONO     SAI_ACR1 = 0x01 << 12 //+ Mono mode
	OUTDRIV  SAI_ACR1 = 0x01 << 13 //+ Output drive
	SAIAEN   SAI_ACR1 = 0x01 << 16 //+ Audio block enable
	DMAEN    SAI_ACR1 = 0x01 << 17 //+ DMA enable
	NODIV    SAI_ACR1 = 0x01 << 19 //+ No divider
	MCKDIV   SAI_ACR1 = 0x0F << 20 //+ Master clock divider
)

const (
	MODEn     = 0
	PRTCFGn   = 2
	DSn       = 5
	LSBFIRSTn = 8
	CKSTRn    = 9
	SYNCENn   = 10
	MONOn     = 12
	OUTDRIVn  = 13
	SAIAENn   = 16
	DMAENn    = 17
	NODIVn    = 19
	MCKDIVn   = 20
)

const (
	FTH     SAI_ACR2 = 0x07 << 0  //+ FIFO threshold
	FFLUSH  SAI_ACR2 = 0x01 << 3  //+ FIFO flush
	TRIS    SAI_ACR2 = 0x01 << 4  //+ Tristate management on data line
	MUTE    SAI_ACR2 = 0x01 << 5  //+ Mute
	MUTEVAL SAI_ACR2 = 0x01 << 6  //+ Mute value
	MUTECNT SAI_ACR2 = 0x3F << 7  //+ Mute counter
	CPL     SAI_ACR2 = 0x01 << 13 //+ Complement bit
	COMP    SAI_ACR2 = 0x03 << 14 //+ Companding mode
)

const (
	FTHn     = 0
	FFLUSHn  = 3
	TRISn    = 4
	MUTEn    = 5
	MUTEVALn = 6
	MUTECNTn = 7
	CPLn     = 13
	COMPn    = 14
)

const (
	FRL   SAI_AFRCR = 0xFF << 0  //+ Frame length
	FSALL SAI_AFRCR = 0x7F << 8  //+ Frame synchronization active level length
	FSDEF SAI_AFRCR = 0x01 << 16 //+ Frame synchronization definition
	FSPOL SAI_AFRCR = 0x01 << 17 //+ Frame synchronization polarity
	FSOFF SAI_AFRCR = 0x01 << 18 //+ Frame synchronization offset
)

const (
	FRLn   = 0
	FSALLn = 8
	FSDEFn = 16
	FSPOLn = 17
	FSOFFn = 18
)

const (
	FBOFF  SAI_ASLOTR = 0x1F << 0    //+ First bit offset
	SLOTSZ SAI_ASLOTR = 0x03 << 6    //+ Slot size
	NBSLOT SAI_ASLOTR = 0x0F << 8    //+ Number of slots in an audio frame
	SLOTEN SAI_ASLOTR = 0xFFFF << 16 //+ Slot enable
)

const (
	FBOFFn  = 0
	SLOTSZn = 6
	NBSLOTn = 8
	SLOTENn = 16
)

const (
	OVRUDRIE  SAI_AIM = 0x01 << 0 //+ Overrun/underrun interrupt enable
	MUTEDETIE SAI_AIM = 0x01 << 1 //+ Mute detection interrupt enable
	WCKCFGIE  SAI_AIM = 0x01 << 2 //+ Wrong clock configuration interrupt enable
	FREQIE    SAI_AIM = 0x01 << 3 //+ FIFO request interrupt enable
	CNRDYIE   SAI_AIM = 0x01 << 4 //+ Codec not ready interrupt enable
	AFSDETIE  SAI_AIM = 0x01 << 5 //+ Anticipated frame synchronization detection interrupt enable
	LFSDETIE  SAI_AIM = 0x01 << 6 //+ Late frame synchronization detection interrupt enable
)

const (
	OVRUDRIEn  = 0
	MUTEDETIEn = 1
	WCKCFGIEn  = 2
	FREQIEn    = 3
	CNRDYIEn   = 4
	AFSDETIEn  = 5
	LFSDETIEn  = 6
)

const (
	OVRUDR  SAI_ASR = 0x01 << 0  //+ Overrun / underrun
	MUTEDET SAI_ASR = 0x01 << 1  //+ Mute detection
	WCKCFG  SAI_ASR = 0x01 << 2  //+ Wrong clock configuration flag
	FREQ    SAI_ASR = 0x01 << 3  //+ FIFO request
	CNRDY   SAI_ASR = 0x01 << 4  //+ Codec not ready
	AFSDET  SAI_ASR = 0x01 << 5  //+ Anticipated frame synchronization detection
	LFSDET  SAI_ASR = 0x01 << 6  //+ Late frame synchronization detection
	FLTH    SAI_ASR = 0x07 << 16 //+ FIFO level threshold
)

const (
	OVRUDRn  = 0
	MUTEDETn = 1
	WCKCFGn  = 2
	FREQn    = 3
	CNRDYn   = 4
	AFSDETn  = 5
	LFSDETn  = 6
	FLTHn    = 16
)

const (
	COVRUDR  SAI_ACLRFR = 0x01 << 0 //+ Clear overrun / underrun
	CMUTEDET SAI_ACLRFR = 0x01 << 1 //+ Mute detection flag
	CWCKCFG  SAI_ACLRFR = 0x01 << 2 //+ Clear wrong clock configuration flag
	CCNRDY   SAI_ACLRFR = 0x01 << 4 //+ Clear codec not ready flag
	CAFSDET  SAI_ACLRFR = 0x01 << 5 //+ Clear anticipated frame synchronization detection flag
	CLFSDET  SAI_ACLRFR = 0x01 << 6 //+ Clear late frame synchronization detection flag
)

const (
	COVRUDRn  = 0
	CMUTEDETn = 1
	CWCKCFGn  = 2
	CCNRDYn   = 4
	CAFSDETn  = 5
	CLFSDETn  = 6
)

const (
	DATA SAI_ADR = 0xFFFFFFFF << 0 //+ Data
)

const (
	DATAn = 0
)

const (
	MODE     SAI_BCR1 = 0x03 << 0  //+ Audio block mode
	PRTCFG   SAI_BCR1 = 0x03 << 2  //+ Protocol configuration
	DS       SAI_BCR1 = 0x07 << 5  //+ Data size
	LSBFIRST SAI_BCR1 = 0x01 << 8  //+ Least significant bit first
	CKSTR    SAI_BCR1 = 0x01 << 9  //+ Clock strobing edge
	SYNCEN   SAI_BCR1 = 0x03 << 10 //+ Synchronization enable
	MONO     SAI_BCR1 = 0x01 << 12 //+ Mono mode
	OUTDRIV  SAI_BCR1 = 0x01 << 13 //+ Output drive
	SAIBEN   SAI_BCR1 = 0x01 << 16 //+ Audio block enable
	DMAEN    SAI_BCR1 = 0x01 << 17 //+ DMA enable
	NODIV    SAI_BCR1 = 0x01 << 19 //+ No divider
	MCKDIV   SAI_BCR1 = 0x0F << 20 //+ Master clock divider
)

const (
	MODEn     = 0
	PRTCFGn   = 2
	DSn       = 5
	LSBFIRSTn = 8
	CKSTRn    = 9
	SYNCENn   = 10
	MONOn     = 12
	OUTDRIVn  = 13
	SAIBENn   = 16
	DMAENn    = 17
	NODIVn    = 19
	MCKDIVn   = 20
)

const (
	FTH     SAI_BCR2 = 0x07 << 0  //+ FIFO threshold
	FFLUSH  SAI_BCR2 = 0x01 << 3  //+ FIFO flush
	TRIS    SAI_BCR2 = 0x01 << 4  //+ Tristate management on data line
	MUTE    SAI_BCR2 = 0x01 << 5  //+ Mute
	MUTEVAL SAI_BCR2 = 0x01 << 6  //+ Mute value
	MUTECNT SAI_BCR2 = 0x3F << 7  //+ Mute counter
	CPL     SAI_BCR2 = 0x01 << 13 //+ Complement bit
	COMP    SAI_BCR2 = 0x03 << 14 //+ Companding mode
)

const (
	FTHn     = 0
	FFLUSHn  = 3
	TRISn    = 4
	MUTEn    = 5
	MUTEVALn = 6
	MUTECNTn = 7
	CPLn     = 13
	COMPn    = 14
)

const (
	FRL   SAI_BFRCR = 0xFF << 0  //+ Frame length
	FSALL SAI_BFRCR = 0x7F << 8  //+ Frame synchronization active level length
	FSDEF SAI_BFRCR = 0x01 << 16 //+ Frame synchronization definition
	FSPOL SAI_BFRCR = 0x01 << 17 //+ Frame synchronization polarity
	FSOFF SAI_BFRCR = 0x01 << 18 //+ Frame synchronization offset
)

const (
	FRLn   = 0
	FSALLn = 8
	FSDEFn = 16
	FSPOLn = 17
	FSOFFn = 18
)

const (
	FBOFF  SAI_BSLOTR = 0x1F << 0    //+ First bit offset
	SLOTSZ SAI_BSLOTR = 0x03 << 6    //+ Slot size
	NBSLOT SAI_BSLOTR = 0x0F << 8    //+ Number of slots in an audio frame
	SLOTEN SAI_BSLOTR = 0xFFFF << 16 //+ Slot enable
)

const (
	FBOFFn  = 0
	SLOTSZn = 6
	NBSLOTn = 8
	SLOTENn = 16
)

const (
	OVRUDRIE  SAI_BIM = 0x01 << 0 //+ Overrun/underrun interrupt enable
	MUTEDETIE SAI_BIM = 0x01 << 1 //+ Mute detection interrupt enable
	WCKCFGIE  SAI_BIM = 0x01 << 2 //+ Wrong clock configuration interrupt enable
	FREQIE    SAI_BIM = 0x01 << 3 //+ FIFO request interrupt enable
	CNRDYIE   SAI_BIM = 0x01 << 4 //+ Codec not ready interrupt enable
	AFSDETIE  SAI_BIM = 0x01 << 5 //+ Anticipated frame synchronization detection interrupt enable
	LFSDETIE  SAI_BIM = 0x01 << 6 //+ Late frame synchronization detection interrupt enable
)

const (
	OVRUDRIEn  = 0
	MUTEDETIEn = 1
	WCKCFGIEn  = 2
	FREQIEn    = 3
	CNRDYIEn   = 4
	AFSDETIEn  = 5
	LFSDETIEn  = 6
)

const (
	OVRUDR  SAI_BSR = 0x01 << 0  //+ Overrun / underrun
	MUTEDET SAI_BSR = 0x01 << 1  //+ Mute detection
	WCKCFG  SAI_BSR = 0x01 << 2  //+ Wrong clock configuration flag
	FREQ    SAI_BSR = 0x01 << 3  //+ FIFO request
	CNRDY   SAI_BSR = 0x01 << 4  //+ Codec not ready
	AFSDET  SAI_BSR = 0x01 << 5  //+ Anticipated frame synchronization detection
	LFSDET  SAI_BSR = 0x01 << 6  //+ Late frame synchronization detection
	FLTH    SAI_BSR = 0x07 << 16 //+ FIFO level threshold
)

const (
	OVRUDRn  = 0
	MUTEDETn = 1
	WCKCFGn  = 2
	FREQn    = 3
	CNRDYn   = 4
	AFSDETn  = 5
	LFSDETn  = 6
	FLTHn    = 16
)

const (
	COVRUDR  SAI_BCLRFR = 0x01 << 0 //+ Clear overrun / underrun
	CMUTEDET SAI_BCLRFR = 0x01 << 1 //+ Mute detection flag
	CWCKCFG  SAI_BCLRFR = 0x01 << 2 //+ Clear wrong clock configuration flag
	CCNRDY   SAI_BCLRFR = 0x01 << 4 //+ Clear codec not ready flag
	CAFSDET  SAI_BCLRFR = 0x01 << 5 //+ Clear anticipated frame synchronization detection flag
	CLFSDET  SAI_BCLRFR = 0x01 << 6 //+ Clear late frame synchronization detection flag
)

const (
	COVRUDRn  = 0
	CMUTEDETn = 1
	CWCKCFGn  = 2
	CCNRDYn   = 4
	CAFSDETn  = 5
	CLFSDETn  = 6
)

const (
	DATA SAI_BDR = 0xFFFFFFFF << 0 //+ Data
)

const (
	DATAn = 0
)
