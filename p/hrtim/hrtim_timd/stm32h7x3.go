// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package hrtim_timd provides access to the registers of the HRTIM_TIMD peripheral.
//
// Instances:
//
//	HRTIM_TIMD  HRTIM_TIMD_BASE  -  HRTIM1_TIMC  High Resolution Timer: TIMD
//
// Registers:
//
//	0x000 32  TIMDCR     Timerx Control Register
//	0x004 32  TIMDISR    Timerx Interrupt Status Register
//	0x008 32  TIMDICR    Timerx Interrupt Clear Register
//	0x00C 32  TIMDDIER5  TIMxDIER5
//	0x010 32  CNTDR      Timerx Counter Register
//	0x014 32  PERDR      Timerx Period Register
//	0x018 32  REPDR      Timerx Repetition Register
//	0x01C 32  CMP1DR     Timerx Compare 1 Register
//	0x020 32  CMP1CDR    Timerx Compare 1 Compound Register
//	0x024 32  CMP2DR     Timerx Compare 2 Register
//	0x028 32  CMP3DR     Timerx Compare 3 Register
//	0x02C 32  CMP4DR     Timerx Compare 4 Register
//	0x030 32  CPT1DR     Timerx Capture 1 Register
//	0x034 32  CPT2DR     Timerx Capture 2 Register
//	0x038 32  DTDR       Timerx Deadtime Register
//	0x03C 32  SETD1R     Timerx Output1 Set Register
//	0x040 32  RSTD1R     Timerx Output1 Reset Register
//	0x044 32  SETD2R     Timerx Output2 Set Register
//	0x048 32  RSTD2R     Timerx Output2 Reset Register
//	0x04C 32  EEFDR1     Timerx External Event Filtering Register 1
//	0x050 32  EEFDR2     Timerx External Event Filtering Register 2
//	0x054 32  RSTDR      TimerA Reset Register
//	0x058 32  CHPDR      Timerx Chopper Register
//	0x05C 32  CPT1DCR    Timerx Capture 2 Control Register
//	0x060 32  CPT2DCR    CPT2xCR
//	0x064 32  OUTDR      Timerx Output Register
//	0x068 32  FLTDR      Timerx Fault Register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package hrtim_timd

const (
	CK_PSCx   TIMDCR = 0x07 << 0  //+ HRTIM Timer x Clock prescaler
	CONT      TIMDCR = 0x01 << 3  //+ Continuous mode
	RETRIG    TIMDCR = 0x01 << 4  //+ Re-triggerable mode
	HALF      TIMDCR = 0x01 << 5  //+ Half mode enable
	PSHPLL    TIMDCR = 0x01 << 6  //+ Push-Pull mode enable
	SYNCRSTx  TIMDCR = 0x01 << 10 //+ Synchronization Resets Timer x
	SYNCSTRTx TIMDCR = 0x01 << 11 //+ Synchronization Starts Timer x
	DELCMP2   TIMDCR = 0x03 << 12 //+ Delayed CMP2 mode
	DELCMP4   TIMDCR = 0x03 << 14 //+ Delayed CMP4 mode
	TxREPU    TIMDCR = 0x01 << 17 //+ Timer x Repetition update
	TxRSTU    TIMDCR = 0x01 << 18 //+ Timerx reset update
	TBU       TIMDCR = 0x01 << 20 //+ TBU
	TCU       TIMDCR = 0x01 << 21 //+ TCU
	TDU       TIMDCR = 0x01 << 22 //+ TDU
	TEU       TIMDCR = 0x01 << 23 //+ TEU
	MSTU      TIMDCR = 0x01 << 24 //+ Master Timer update
	DACSYNC   TIMDCR = 0x03 << 25 //+ AC Synchronization
	PREEN     TIMDCR = 0x01 << 27 //+ Preload enable
	UPDGAT    TIMDCR = 0x0F << 28 //+ Update Gating
)

const (
	CK_PSCxn   = 0
	CONTn      = 3
	RETRIGn    = 4
	HALFn      = 5
	PSHPLLn    = 6
	SYNCRSTxn  = 10
	SYNCSTRTxn = 11
	DELCMP2n   = 12
	DELCMP4n   = 14
	TxREPUn    = 17
	TxRSTUn    = 18
	TBUn       = 20
	TCUn       = 21
	TDUn       = 22
	TEUn       = 23
	MSTUn      = 24
	DACSYNCn   = 25
	PREENn     = 27
	UPDGATn    = 28
)

const (
	CMP1    TIMDISR = 0x01 << 0  //+ Compare 1 Interrupt Flag
	CMP2    TIMDISR = 0x01 << 1  //+ Compare 2 Interrupt Flag
	CMP3    TIMDISR = 0x01 << 2  //+ Compare 3 Interrupt Flag
	CMP4    TIMDISR = 0x01 << 3  //+ Compare 4 Interrupt Flag
	REP     TIMDISR = 0x01 << 4  //+ Repetition Interrupt Flag
	UPD     TIMDISR = 0x01 << 6  //+ Update Interrupt Flag
	CPT1    TIMDISR = 0x01 << 7  //+ Capture1 Interrupt Flag
	CPT2    TIMDISR = 0x01 << 8  //+ Capture2 Interrupt Flag
	SETx1   TIMDISR = 0x01 << 9  //+ Output 1 Set Interrupt Flag
	RSTx1   TIMDISR = 0x01 << 10 //+ Output 1 Reset Interrupt Flag
	SETx2   TIMDISR = 0x01 << 11 //+ Output 2 Set Interrupt Flag
	RSTx2   TIMDISR = 0x01 << 12 //+ Output 2 Reset Interrupt Flag
	RST     TIMDISR = 0x01 << 13 //+ Reset Interrupt Flag
	DLYPRT  TIMDISR = 0x01 << 14 //+ Delayed Protection Flag
	CPPSTAT TIMDISR = 0x01 << 16 //+ Current Push Pull Status
	IPPSTAT TIMDISR = 0x01 << 17 //+ Idle Push Pull Status
	O1STAT  TIMDISR = 0x01 << 18 //+ Output 1 State
	O2STAT  TIMDISR = 0x01 << 19 //+ Output 2 State
)

const (
	CMP1n    = 0
	CMP2n    = 1
	CMP3n    = 2
	CMP4n    = 3
	REPn     = 4
	UPDn     = 6
	CPT1n    = 7
	CPT2n    = 8
	SETx1n   = 9
	RSTx1n   = 10
	SETx2n   = 11
	RSTx2n   = 12
	RSTn     = 13
	DLYPRTn  = 14
	CPPSTATn = 16
	IPPSTATn = 17
	O1STATn  = 18
	O2STATn  = 19
)

const (
	CMP1C   TIMDICR = 0x01 << 0  //+ Compare 1 Interrupt flag Clear
	CMP2C   TIMDICR = 0x01 << 1  //+ Compare 2 Interrupt flag Clear
	CMP3C   TIMDICR = 0x01 << 2  //+ Compare 3 Interrupt flag Clear
	CMP4C   TIMDICR = 0x01 << 3  //+ Compare 4 Interrupt flag Clear
	REPC    TIMDICR = 0x01 << 4  //+ Repetition Interrupt flag Clear
	UPDC    TIMDICR = 0x01 << 6  //+ Update Interrupt flag Clear
	CPT1C   TIMDICR = 0x01 << 7  //+ Capture1 Interrupt flag Clear
	CPT2C   TIMDICR = 0x01 << 8  //+ Capture2 Interrupt flag Clear
	SET1xC  TIMDICR = 0x01 << 9  //+ Output 1 Set flag Clear
	RSTx1C  TIMDICR = 0x01 << 10 //+ Output 1 Reset flag Clear
	SET2xC  TIMDICR = 0x01 << 11 //+ Output 2 Set flag Clear
	RSTx2C  TIMDICR = 0x01 << 12 //+ Output 2 Reset flag Clear
	RSTC    TIMDICR = 0x01 << 13 //+ Reset Interrupt flag Clear
	DLYPRTC TIMDICR = 0x01 << 14 //+ Delayed Protection Flag Clear
)

const (
	CMP1Cn   = 0
	CMP2Cn   = 1
	CMP3Cn   = 2
	CMP4Cn   = 3
	REPCn    = 4
	UPDCn    = 6
	CPT1Cn   = 7
	CPT2Cn   = 8
	SET1xCn  = 9
	RSTx1Cn  = 10
	SET2xCn  = 11
	RSTx2Cn  = 12
	RSTCn    = 13
	DLYPRTCn = 14
)

const (
	CMP1IE   TIMDDIER5 = 0x01 << 0  //+ CMP1IE
	CMP2IE   TIMDDIER5 = 0x01 << 1  //+ CMP2IE
	CMP3IE   TIMDDIER5 = 0x01 << 2  //+ CMP3IE
	CMP4IE   TIMDDIER5 = 0x01 << 3  //+ CMP4IE
	REPIE    TIMDDIER5 = 0x01 << 4  //+ REPIE
	UPDIE    TIMDDIER5 = 0x01 << 6  //+ UPDIE
	CPT1IE   TIMDDIER5 = 0x01 << 7  //+ CPT1IE
	CPT2IE   TIMDDIER5 = 0x01 << 8  //+ CPT2IE
	SET1xIE  TIMDDIER5 = 0x01 << 9  //+ SET1xIE
	RSTx1IE  TIMDDIER5 = 0x01 << 10 //+ RSTx1IE
	SETx2IE  TIMDDIER5 = 0x01 << 11 //+ SETx2IE
	RSTx2IE  TIMDDIER5 = 0x01 << 12 //+ RSTx2IE
	RSTIE    TIMDDIER5 = 0x01 << 13 //+ RSTIE
	DLYPRTIE TIMDDIER5 = 0x01 << 14 //+ DLYPRTIE
	CMP1DE   TIMDDIER5 = 0x01 << 16 //+ CMP1DE
	CMP2DE   TIMDDIER5 = 0x01 << 17 //+ CMP2DE
	CMP3DE   TIMDDIER5 = 0x01 << 18 //+ CMP3DE
	CMP4DE   TIMDDIER5 = 0x01 << 19 //+ CMP4DE
	REPDE    TIMDDIER5 = 0x01 << 20 //+ REPDE
	UPDDE    TIMDDIER5 = 0x01 << 22 //+ UPDDE
	CPT1DE   TIMDDIER5 = 0x01 << 23 //+ CPT1DE
	CPT2DE   TIMDDIER5 = 0x01 << 24 //+ CPT2DE
	SET1xDE  TIMDDIER5 = 0x01 << 25 //+ SET1xDE
	RSTx1DE  TIMDDIER5 = 0x01 << 26 //+ RSTx1DE
	SETx2DE  TIMDDIER5 = 0x01 << 27 //+ SETx2DE
	RSTx2DE  TIMDDIER5 = 0x01 << 28 //+ RSTx2DE
	RSTDE    TIMDDIER5 = 0x01 << 29 //+ RSTDE
	DLYPRTDE TIMDDIER5 = 0x01 << 30 //+ DLYPRTDE
)

const (
	CMP1IEn   = 0
	CMP2IEn   = 1
	CMP3IEn   = 2
	CMP4IEn   = 3
	REPIEn    = 4
	UPDIEn    = 6
	CPT1IEn   = 7
	CPT2IEn   = 8
	SET1xIEn  = 9
	RSTx1IEn  = 10
	SETx2IEn  = 11
	RSTx2IEn  = 12
	RSTIEn    = 13
	DLYPRTIEn = 14
	CMP1DEn   = 16
	CMP2DEn   = 17
	CMP3DEn   = 18
	CMP4DEn   = 19
	REPDEn    = 20
	UPDDEn    = 22
	CPT1DEn   = 23
	CPT2DEn   = 24
	SET1xDEn  = 25
	RSTx1DEn  = 26
	SETx2DEn  = 27
	RSTx2DEn  = 28
	RSTDEn    = 29
	DLYPRTDEn = 30
)

const (
	CNTx CNTDR = 0xFFFF << 0 //+ Timerx Counter value
)

const (
	CNTxn = 0
)

const (
	PERx PERDR = 0xFFFF << 0 //+ Timerx Period value
)

const (
	PERxn = 0
)

const (
	REPx REPDR = 0xFF << 0 //+ Timerx Repetition counter value
)

const (
	REPxn = 0
)

const (
	CMP1x CMP1DR = 0xFFFF << 0 //+ Timerx Compare 1 value
)

const (
	CMP1xn = 0
)

const (
	CMP1x CMP1CDR = 0xFFFF << 0 //+ Timerx Compare 1 value
	REPx  CMP1CDR = 0xFF << 16  //+ Timerx Repetition value (aliased from HRTIM_REPx register)
)

const (
	CMP1xn = 0
	REPxn  = 16
)

const (
	CMP2x CMP2DR = 0xFFFF << 0 //+ Timerx Compare 2 value
)

const (
	CMP2xn = 0
)

const (
	CMP3x CMP3DR = 0xFFFF << 0 //+ Timerx Compare 3 value
)

const (
	CMP3xn = 0
)

const (
	CMP4x CMP4DR = 0xFFFF << 0 //+ Timerx Compare 4 value
)

const (
	CMP4xn = 0
)

const (
	CPT1x CPT1DR = 0xFFFF << 0 //+ Timerx Capture 1 value
)

const (
	CPT1xn = 0
)

const (
	CPT2x CPT2DR = 0xFFFF << 0 //+ Timerx Capture 2 value
)

const (
	CPT2xn = 0
)

const (
	DTRx    DTDR = 0x1FF << 0  //+ Deadtime Rising value
	SDTRx   DTDR = 0x01 << 9   //+ Sign Deadtime Rising value
	DTPRSC  DTDR = 0x07 << 10  //+ Deadtime Prescaler
	DTRSLKx DTDR = 0x01 << 14  //+ Deadtime Rising Sign Lock
	DTRLKx  DTDR = 0x01 << 15  //+ Deadtime Rising Lock
	DTFx    DTDR = 0x1FF << 16 //+ Deadtime Falling value
	SDTFx   DTDR = 0x01 << 25  //+ Sign Deadtime Falling value
	DTFSLKx DTDR = 0x01 << 30  //+ Deadtime Falling Sign Lock
	DTFLKx  DTDR = 0x01 << 31  //+ Deadtime Falling Lock
)

const (
	DTRxn    = 0
	SDTRxn   = 9
	DTPRSCn  = 10
	DTRSLKxn = 14
	DTRLKxn  = 15
	DTFxn    = 16
	SDTFxn   = 25
	DTFSLKxn = 30
	DTFLKxn  = 31
)

const (
	SST       SETD1R = 0x01 << 0  //+ Software Set trigger
	RESYNC    SETD1R = 0x01 << 1  //+ Timer A resynchronizaton
	PER       SETD1R = 0x01 << 2  //+ Timer A Period
	CMP1      SETD1R = 0x01 << 3  //+ Timer A compare 1
	CMP2      SETD1R = 0x01 << 4  //+ Timer A compare 2
	CMP3      SETD1R = 0x01 << 5  //+ Timer A compare 3
	CMP4      SETD1R = 0x01 << 6  //+ Timer A compare 4
	MSTPER    SETD1R = 0x01 << 7  //+ Master Period
	MSTCMP1   SETD1R = 0x01 << 8  //+ Master Compare 1
	MSTCMP2   SETD1R = 0x01 << 9  //+ Master Compare 2
	MSTCMP3   SETD1R = 0x01 << 10 //+ Master Compare 3
	MSTCMP4   SETD1R = 0x01 << 11 //+ Master Compare 4
	TIMEVNT1  SETD1R = 0x01 << 12 //+ Timer Event 1
	TIMEVNT2  SETD1R = 0x01 << 13 //+ Timer Event 2
	TIMEVNT3  SETD1R = 0x01 << 14 //+ Timer Event 3
	TIMEVNT4  SETD1R = 0x01 << 15 //+ Timer Event 4
	TIMEVNT5  SETD1R = 0x01 << 16 //+ Timer Event 5
	TIMEVNT6  SETD1R = 0x01 << 17 //+ Timer Event 6
	TIMEVNT7  SETD1R = 0x01 << 18 //+ Timer Event 7
	TIMEVNT8  SETD1R = 0x01 << 19 //+ Timer Event 8
	TIMEVNT9  SETD1R = 0x01 << 20 //+ Timer Event 9
	EXTEVNT1  SETD1R = 0x01 << 21 //+ External Event 1
	EXTEVNT2  SETD1R = 0x01 << 22 //+ External Event 2
	EXTEVNT3  SETD1R = 0x01 << 23 //+ External Event 3
	EXTEVNT4  SETD1R = 0x01 << 24 //+ External Event 4
	EXTEVNT5  SETD1R = 0x01 << 25 //+ External Event 5
	EXTEVNT6  SETD1R = 0x01 << 26 //+ External Event 6
	EXTEVNT7  SETD1R = 0x01 << 27 //+ External Event 7
	EXTEVNT8  SETD1R = 0x01 << 28 //+ External Event 8
	EXTEVNT9  SETD1R = 0x01 << 29 //+ External Event 9
	EXTEVNT10 SETD1R = 0x01 << 30 //+ External Event 10
	UPDATE    SETD1R = 0x01 << 31 //+ Registers update (transfer preload to active)
)

const (
	SSTn       = 0
	RESYNCn    = 1
	PERn       = 2
	CMP1n      = 3
	CMP2n      = 4
	CMP3n      = 5
	CMP4n      = 6
	MSTPERn    = 7
	MSTCMP1n   = 8
	MSTCMP2n   = 9
	MSTCMP3n   = 10
	MSTCMP4n   = 11
	TIMEVNT1n  = 12
	TIMEVNT2n  = 13
	TIMEVNT3n  = 14
	TIMEVNT4n  = 15
	TIMEVNT5n  = 16
	TIMEVNT6n  = 17
	TIMEVNT7n  = 18
	TIMEVNT8n  = 19
	TIMEVNT9n  = 20
	EXTEVNT1n  = 21
	EXTEVNT2n  = 22
	EXTEVNT3n  = 23
	EXTEVNT4n  = 24
	EXTEVNT5n  = 25
	EXTEVNT6n  = 26
	EXTEVNT7n  = 27
	EXTEVNT8n  = 28
	EXTEVNT9n  = 29
	EXTEVNT10n = 30
	UPDATEn    = 31
)

const (
	SRT       RSTD1R = 0x01 << 0  //+ SRT
	RESYNC    RSTD1R = 0x01 << 1  //+ RESYNC
	PER       RSTD1R = 0x01 << 2  //+ PER
	CMP1      RSTD1R = 0x01 << 3  //+ CMP1
	CMP2      RSTD1R = 0x01 << 4  //+ CMP2
	CMP3      RSTD1R = 0x01 << 5  //+ CMP3
	CMP4      RSTD1R = 0x01 << 6  //+ CMP4
	MSTPER    RSTD1R = 0x01 << 7  //+ MSTPER
	MSTCMP1   RSTD1R = 0x01 << 8  //+ MSTCMP1
	MSTCMP2   RSTD1R = 0x01 << 9  //+ MSTCMP2
	MSTCMP3   RSTD1R = 0x01 << 10 //+ MSTCMP3
	MSTCMP4   RSTD1R = 0x01 << 11 //+ MSTCMP4
	TIMEVNT1  RSTD1R = 0x01 << 12 //+ TIMEVNT1
	TIMEVNT2  RSTD1R = 0x01 << 13 //+ TIMEVNT2
	TIMEVNT3  RSTD1R = 0x01 << 14 //+ TIMEVNT3
	TIMEVNT4  RSTD1R = 0x01 << 15 //+ TIMEVNT4
	TIMEVNT5  RSTD1R = 0x01 << 16 //+ TIMEVNT5
	TIMEVNT6  RSTD1R = 0x01 << 17 //+ TIMEVNT6
	TIMEVNT7  RSTD1R = 0x01 << 18 //+ TIMEVNT7
	TIMEVNT8  RSTD1R = 0x01 << 19 //+ TIMEVNT8
	TIMEVNT9  RSTD1R = 0x01 << 20 //+ TIMEVNT9
	EXTEVNT1  RSTD1R = 0x01 << 21 //+ EXTEVNT1
	EXTEVNT2  RSTD1R = 0x01 << 22 //+ EXTEVNT2
	EXTEVNT3  RSTD1R = 0x01 << 23 //+ EXTEVNT3
	EXTEVNT4  RSTD1R = 0x01 << 24 //+ EXTEVNT4
	EXTEVNT5  RSTD1R = 0x01 << 25 //+ EXTEVNT5
	EXTEVNT6  RSTD1R = 0x01 << 26 //+ EXTEVNT6
	EXTEVNT7  RSTD1R = 0x01 << 27 //+ EXTEVNT7
	EXTEVNT8  RSTD1R = 0x01 << 28 //+ EXTEVNT8
	EXTEVNT9  RSTD1R = 0x01 << 29 //+ EXTEVNT9
	EXTEVNT10 RSTD1R = 0x01 << 30 //+ EXTEVNT10
	UPDATE    RSTD1R = 0x01 << 31 //+ UPDATE
)

const (
	SRTn       = 0
	RESYNCn    = 1
	PERn       = 2
	CMP1n      = 3
	CMP2n      = 4
	CMP3n      = 5
	CMP4n      = 6
	MSTPERn    = 7
	MSTCMP1n   = 8
	MSTCMP2n   = 9
	MSTCMP3n   = 10
	MSTCMP4n   = 11
	TIMEVNT1n  = 12
	TIMEVNT2n  = 13
	TIMEVNT3n  = 14
	TIMEVNT4n  = 15
	TIMEVNT5n  = 16
	TIMEVNT6n  = 17
	TIMEVNT7n  = 18
	TIMEVNT8n  = 19
	TIMEVNT9n  = 20
	EXTEVNT1n  = 21
	EXTEVNT2n  = 22
	EXTEVNT3n  = 23
	EXTEVNT4n  = 24
	EXTEVNT5n  = 25
	EXTEVNT6n  = 26
	EXTEVNT7n  = 27
	EXTEVNT8n  = 28
	EXTEVNT9n  = 29
	EXTEVNT10n = 30
	UPDATEn    = 31
)

const (
	SST       SETD2R = 0x01 << 0  //+ SST
	RESYNC    SETD2R = 0x01 << 1  //+ RESYNC
	PER       SETD2R = 0x01 << 2  //+ PER
	CMP1      SETD2R = 0x01 << 3  //+ CMP1
	CMP2      SETD2R = 0x01 << 4  //+ CMP2
	CMP3      SETD2R = 0x01 << 5  //+ CMP3
	CMP4      SETD2R = 0x01 << 6  //+ CMP4
	MSTPER    SETD2R = 0x01 << 7  //+ MSTPER
	MSTCMP1   SETD2R = 0x01 << 8  //+ MSTCMP1
	MSTCMP2   SETD2R = 0x01 << 9  //+ MSTCMP2
	MSTCMP3   SETD2R = 0x01 << 10 //+ MSTCMP3
	MSTCMP4   SETD2R = 0x01 << 11 //+ MSTCMP4
	TIMEVNT1  SETD2R = 0x01 << 12 //+ TIMEVNT1
	TIMEVNT2  SETD2R = 0x01 << 13 //+ TIMEVNT2
	TIMEVNT3  SETD2R = 0x01 << 14 //+ TIMEVNT3
	TIMEVNT4  SETD2R = 0x01 << 15 //+ TIMEVNT4
	TIMEVNT5  SETD2R = 0x01 << 16 //+ TIMEVNT5
	TIMEVNT6  SETD2R = 0x01 << 17 //+ TIMEVNT6
	TIMEVNT7  SETD2R = 0x01 << 18 //+ TIMEVNT7
	TIMEVNT8  SETD2R = 0x01 << 19 //+ TIMEVNT8
	TIMEVNT9  SETD2R = 0x01 << 20 //+ TIMEVNT9
	EXTEVNT1  SETD2R = 0x01 << 21 //+ EXTEVNT1
	EXTEVNT2  SETD2R = 0x01 << 22 //+ EXTEVNT2
	EXTEVNT3  SETD2R = 0x01 << 23 //+ EXTEVNT3
	EXTEVNT4  SETD2R = 0x01 << 24 //+ EXTEVNT4
	EXTEVNT5  SETD2R = 0x01 << 25 //+ EXTEVNT5
	EXTEVNT6  SETD2R = 0x01 << 26 //+ EXTEVNT6
	EXTEVNT7  SETD2R = 0x01 << 27 //+ EXTEVNT7
	EXTEVNT8  SETD2R = 0x01 << 28 //+ EXTEVNT8
	EXTEVNT9  SETD2R = 0x01 << 29 //+ EXTEVNT9
	EXTEVNT10 SETD2R = 0x01 << 30 //+ EXTEVNT10
	UPDATE    SETD2R = 0x01 << 31 //+ UPDATE
)

const (
	SSTn       = 0
	RESYNCn    = 1
	PERn       = 2
	CMP1n      = 3
	CMP2n      = 4
	CMP3n      = 5
	CMP4n      = 6
	MSTPERn    = 7
	MSTCMP1n   = 8
	MSTCMP2n   = 9
	MSTCMP3n   = 10
	MSTCMP4n   = 11
	TIMEVNT1n  = 12
	TIMEVNT2n  = 13
	TIMEVNT3n  = 14
	TIMEVNT4n  = 15
	TIMEVNT5n  = 16
	TIMEVNT6n  = 17
	TIMEVNT7n  = 18
	TIMEVNT8n  = 19
	TIMEVNT9n  = 20
	EXTEVNT1n  = 21
	EXTEVNT2n  = 22
	EXTEVNT3n  = 23
	EXTEVNT4n  = 24
	EXTEVNT5n  = 25
	EXTEVNT6n  = 26
	EXTEVNT7n  = 27
	EXTEVNT8n  = 28
	EXTEVNT9n  = 29
	EXTEVNT10n = 30
	UPDATEn    = 31
)

const (
	SRT       RSTD2R = 0x01 << 0  //+ SRT
	RESYNC    RSTD2R = 0x01 << 1  //+ RESYNC
	PER       RSTD2R = 0x01 << 2  //+ PER
	CMP1      RSTD2R = 0x01 << 3  //+ CMP1
	CMP2      RSTD2R = 0x01 << 4  //+ CMP2
	CMP3      RSTD2R = 0x01 << 5  //+ CMP3
	CMP4      RSTD2R = 0x01 << 6  //+ CMP4
	MSTPER    RSTD2R = 0x01 << 7  //+ MSTPER
	MSTCMP1   RSTD2R = 0x01 << 8  //+ MSTCMP1
	MSTCMP2   RSTD2R = 0x01 << 9  //+ MSTCMP2
	MSTCMP3   RSTD2R = 0x01 << 10 //+ MSTCMP3
	MSTCMP4   RSTD2R = 0x01 << 11 //+ MSTCMP4
	TIMEVNT1  RSTD2R = 0x01 << 12 //+ TIMEVNT1
	TIMEVNT2  RSTD2R = 0x01 << 13 //+ TIMEVNT2
	TIMEVNT3  RSTD2R = 0x01 << 14 //+ TIMEVNT3
	TIMEVNT4  RSTD2R = 0x01 << 15 //+ TIMEVNT4
	TIMEVNT5  RSTD2R = 0x01 << 16 //+ TIMEVNT5
	TIMEVNT6  RSTD2R = 0x01 << 17 //+ TIMEVNT6
	TIMEVNT7  RSTD2R = 0x01 << 18 //+ TIMEVNT7
	TIMEVNT8  RSTD2R = 0x01 << 19 //+ TIMEVNT8
	TIMEVNT9  RSTD2R = 0x01 << 20 //+ TIMEVNT9
	EXTEVNT1  RSTD2R = 0x01 << 21 //+ EXTEVNT1
	EXTEVNT2  RSTD2R = 0x01 << 22 //+ EXTEVNT2
	EXTEVNT3  RSTD2R = 0x01 << 23 //+ EXTEVNT3
	EXTEVNT4  RSTD2R = 0x01 << 24 //+ EXTEVNT4
	EXTEVNT5  RSTD2R = 0x01 << 25 //+ EXTEVNT5
	EXTEVNT6  RSTD2R = 0x01 << 26 //+ EXTEVNT6
	EXTEVNT7  RSTD2R = 0x01 << 27 //+ EXTEVNT7
	EXTEVNT8  RSTD2R = 0x01 << 28 //+ EXTEVNT8
	EXTEVNT9  RSTD2R = 0x01 << 29 //+ EXTEVNT9
	EXTEVNT10 RSTD2R = 0x01 << 30 //+ EXTEVNT10
	UPDATE    RSTD2R = 0x01 << 31 //+ UPDATE
)

const (
	SRTn       = 0
	RESYNCn    = 1
	PERn       = 2
	CMP1n      = 3
	CMP2n      = 4
	CMP3n      = 5
	CMP4n      = 6
	MSTPERn    = 7
	MSTCMP1n   = 8
	MSTCMP2n   = 9
	MSTCMP3n   = 10
	MSTCMP4n   = 11
	TIMEVNT1n  = 12
	TIMEVNT2n  = 13
	TIMEVNT3n  = 14
	TIMEVNT4n  = 15
	TIMEVNT5n  = 16
	TIMEVNT6n  = 17
	TIMEVNT7n  = 18
	TIMEVNT8n  = 19
	TIMEVNT9n  = 20
	EXTEVNT1n  = 21
	EXTEVNT2n  = 22
	EXTEVNT3n  = 23
	EXTEVNT4n  = 24
	EXTEVNT5n  = 25
	EXTEVNT6n  = 26
	EXTEVNT7n  = 27
	EXTEVNT8n  = 28
	EXTEVNT9n  = 29
	EXTEVNT10n = 30
	UPDATEn    = 31
)

const (
	EE1LTCH EEFDR1 = 0x01 << 0  //+ External Event 1 latch
	EE1FLTR EEFDR1 = 0x0F << 1  //+ External Event 1 filter
	EE2LTCH EEFDR1 = 0x01 << 6  //+ External Event 2 latch
	EE2FLTR EEFDR1 = 0x0F << 7  //+ External Event 2 filter
	EE3LTCH EEFDR1 = 0x01 << 12 //+ External Event 3 latch
	EE3FLTR EEFDR1 = 0x0F << 13 //+ External Event 3 filter
	EE4LTCH EEFDR1 = 0x01 << 18 //+ External Event 4 latch
	EE4FLTR EEFDR1 = 0x0F << 19 //+ External Event 4 filter
	EE5LTCH EEFDR1 = 0x01 << 24 //+ External Event 5 latch
	EE5FLTR EEFDR1 = 0x0F << 25 //+ External Event 5 filter
)

const (
	EE1LTCHn = 0
	EE1FLTRn = 1
	EE2LTCHn = 6
	EE2FLTRn = 7
	EE3LTCHn = 12
	EE3FLTRn = 13
	EE4LTCHn = 18
	EE4FLTRn = 19
	EE5LTCHn = 24
	EE5FLTRn = 25
)

const (
	EE6LTCH  EEFDR2 = 0x01 << 0  //+ External Event 6 latch
	EE6FLTR  EEFDR2 = 0x0F << 1  //+ External Event 6 filter
	EE7LTCH  EEFDR2 = 0x01 << 6  //+ External Event 7 latch
	EE7FLTR  EEFDR2 = 0x0F << 7  //+ External Event 7 filter
	EE8LTCH  EEFDR2 = 0x01 << 12 //+ External Event 8 latch
	EE8FLTR  EEFDR2 = 0x0F << 13 //+ External Event 8 filter
	EE9LTCH  EEFDR2 = 0x01 << 18 //+ External Event 9 latch
	EE9FLTR  EEFDR2 = 0x0F << 19 //+ External Event 9 filter
	EE10LTCH EEFDR2 = 0x01 << 24 //+ External Event 10 latch
	EE10FLTR EEFDR2 = 0x0F << 25 //+ External Event 10 filter
)

const (
	EE6LTCHn  = 0
	EE6FLTRn  = 1
	EE7LTCHn  = 6
	EE7FLTRn  = 7
	EE8LTCHn  = 12
	EE8FLTRn  = 13
	EE9LTCHn  = 18
	EE9FLTRn  = 19
	EE10LTCHn = 24
	EE10FLTRn = 25
)

const (
	UPDT      RSTDR = 0x01 << 1  //+ Timer A Update reset
	CMP2      RSTDR = 0x01 << 2  //+ Timer A compare 2 reset
	CMP4      RSTDR = 0x01 << 3  //+ Timer A compare 4 reset
	MSTPER    RSTDR = 0x01 << 4  //+ Master timer Period
	MSTCMP1   RSTDR = 0x01 << 5  //+ Master compare 1
	MSTCMP2   RSTDR = 0x01 << 6  //+ Master compare 2
	MSTCMP3   RSTDR = 0x01 << 7  //+ Master compare 3
	MSTCMP4   RSTDR = 0x01 << 8  //+ Master compare 4
	EXTEVNT1  RSTDR = 0x01 << 9  //+ External Event 1
	EXTEVNT2  RSTDR = 0x01 << 10 //+ External Event 2
	EXTEVNT3  RSTDR = 0x01 << 11 //+ External Event 3
	EXTEVNT4  RSTDR = 0x01 << 12 //+ External Event 4
	EXTEVNT5  RSTDR = 0x01 << 13 //+ External Event 5
	EXTEVNT6  RSTDR = 0x01 << 14 //+ External Event 6
	EXTEVNT7  RSTDR = 0x01 << 15 //+ External Event 7
	EXTEVNT8  RSTDR = 0x01 << 16 //+ External Event 8
	EXTEVNT9  RSTDR = 0x01 << 17 //+ External Event 9
	EXTEVNT10 RSTDR = 0x01 << 18 //+ External Event 10
	TIMACMP1  RSTDR = 0x01 << 19 //+ Timer A Compare 1
	TIMACMP2  RSTDR = 0x01 << 20 //+ Timer A Compare 2
	TIMACMP4  RSTDR = 0x01 << 21 //+ Timer A Compare 4
	TIMBCMP1  RSTDR = 0x01 << 22 //+ Timer B Compare 1
	TIMBCMP2  RSTDR = 0x01 << 23 //+ Timer B Compare 2
	TIMBCMP4  RSTDR = 0x01 << 24 //+ Timer B Compare 4
	TIMCCMP1  RSTDR = 0x01 << 25 //+ Timer C Compare 1
	TIMCCMP2  RSTDR = 0x01 << 26 //+ Timer C Compare 2
	TIMCCMP4  RSTDR = 0x01 << 27 //+ Timer C Compare 4
	TIMECMP1  RSTDR = 0x01 << 28 //+ Timer E Compare 1
	TIMECMP2  RSTDR = 0x01 << 29 //+ Timer E Compare 2
	TIMECMP4  RSTDR = 0x01 << 30 //+ Timer E Compare 4
)

const (
	UPDTn      = 1
	CMP2n      = 2
	CMP4n      = 3
	MSTPERn    = 4
	MSTCMP1n   = 5
	MSTCMP2n   = 6
	MSTCMP3n   = 7
	MSTCMP4n   = 8
	EXTEVNT1n  = 9
	EXTEVNT2n  = 10
	EXTEVNT3n  = 11
	EXTEVNT4n  = 12
	EXTEVNT5n  = 13
	EXTEVNT6n  = 14
	EXTEVNT7n  = 15
	EXTEVNT8n  = 16
	EXTEVNT9n  = 17
	EXTEVNT10n = 18
	TIMACMP1n  = 19
	TIMACMP2n  = 20
	TIMACMP4n  = 21
	TIMBCMP1n  = 22
	TIMBCMP2n  = 23
	TIMBCMP4n  = 24
	TIMCCMP1n  = 25
	TIMCCMP2n  = 26
	TIMCCMP4n  = 27
	TIMECMP1n  = 28
	TIMECMP2n  = 29
	TIMECMP4n  = 30
)

const (
	CHPFRQ CHPDR = 0x0F << 0 //+ Timerx carrier frequency value
	CHPDTY CHPDR = 0x07 << 4 //+ Timerx chopper duty cycle value
	STRTPW CHPDR = 0x0F << 7 //+ STRTPW
)

const (
	CHPFRQn = 0
	CHPDTYn = 4
	STRTPWn = 7
)

const (
	SWCPT     CPT1DCR = 0x01 << 0  //+ Software Capture
	UDPCPT    CPT1DCR = 0x01 << 1  //+ Update Capture
	EXEV1CPT  CPT1DCR = 0x01 << 2  //+ External Event 1 Capture
	EXEV2CPT  CPT1DCR = 0x01 << 3  //+ External Event 2 Capture
	EXEV3CPT  CPT1DCR = 0x01 << 4  //+ External Event 3 Capture
	EXEV4CPT  CPT1DCR = 0x01 << 5  //+ External Event 4 Capture
	EXEV5CPT  CPT1DCR = 0x01 << 6  //+ External Event 5 Capture
	EXEV6CPT  CPT1DCR = 0x01 << 7  //+ External Event 6 Capture
	EXEV7CPT  CPT1DCR = 0x01 << 8  //+ External Event 7 Capture
	EXEV8CPT  CPT1DCR = 0x01 << 9  //+ External Event 8 Capture
	EXEV9CPT  CPT1DCR = 0x01 << 10 //+ External Event 9 Capture
	EXEV10CPT CPT1DCR = 0x01 << 11 //+ External Event 10 Capture
	TA1SET    CPT1DCR = 0x01 << 12 //+ Timer A output 1 Set
	TA1RST    CPT1DCR = 0x01 << 13 //+ Timer A output 1 Reset
	TACMP1    CPT1DCR = 0x01 << 14 //+ Timer A Compare 1
	TACMP2    CPT1DCR = 0x01 << 15 //+ Timer A Compare 2
	TB1SET    CPT1DCR = 0x01 << 16 //+ Timer B output 1 Set
	TB1RST    CPT1DCR = 0x01 << 17 //+ Timer B output 1 Reset
	TBCMP1    CPT1DCR = 0x01 << 18 //+ Timer B Compare 1
	TBCMP2    CPT1DCR = 0x01 << 19 //+ Timer B Compare 2
	TC1SET    CPT1DCR = 0x01 << 20 //+ Timer C output 1 Set
	TC1RST    CPT1DCR = 0x01 << 21 //+ Timer C output 1 Reset
	TCCMP1    CPT1DCR = 0x01 << 22 //+ Timer C Compare 1
	TCCMP2    CPT1DCR = 0x01 << 23 //+ Timer C Compare 2
	TE1SET    CPT1DCR = 0x01 << 28 //+ Timer E output 1 Set
	TE1RST    CPT1DCR = 0x01 << 29 //+ Timer E output 1 Reset
	TECMP1    CPT1DCR = 0x01 << 30 //+ Timer E Compare 1
	TECMP2    CPT1DCR = 0x01 << 31 //+ Timer E Compare 2
)

const (
	SWCPTn     = 0
	UDPCPTn    = 1
	EXEV1CPTn  = 2
	EXEV2CPTn  = 3
	EXEV3CPTn  = 4
	EXEV4CPTn  = 5
	EXEV5CPTn  = 6
	EXEV6CPTn  = 7
	EXEV7CPTn  = 8
	EXEV8CPTn  = 9
	EXEV9CPTn  = 10
	EXEV10CPTn = 11
	TA1SETn    = 12
	TA1RSTn    = 13
	TACMP1n    = 14
	TACMP2n    = 15
	TB1SETn    = 16
	TB1RSTn    = 17
	TBCMP1n    = 18
	TBCMP2n    = 19
	TC1SETn    = 20
	TC1RSTn    = 21
	TCCMP1n    = 22
	TCCMP2n    = 23
	TE1SETn    = 28
	TE1RSTn    = 29
	TECMP1n    = 30
	TECMP2n    = 31
)

const (
	SWCPT     CPT2DCR = 0x01 << 0  //+ Software Capture
	UDPCPT    CPT2DCR = 0x01 << 1  //+ Update Capture
	EXEV1CPT  CPT2DCR = 0x01 << 2  //+ External Event 1 Capture
	EXEV2CPT  CPT2DCR = 0x01 << 3  //+ External Event 2 Capture
	EXEV3CPT  CPT2DCR = 0x01 << 4  //+ External Event 3 Capture
	EXEV4CPT  CPT2DCR = 0x01 << 5  //+ External Event 4 Capture
	EXEV5CPT  CPT2DCR = 0x01 << 6  //+ External Event 5 Capture
	EXEV6CPT  CPT2DCR = 0x01 << 7  //+ External Event 6 Capture
	EXEV7CPT  CPT2DCR = 0x01 << 8  //+ External Event 7 Capture
	EXEV8CPT  CPT2DCR = 0x01 << 9  //+ External Event 8 Capture
	EXEV9CPT  CPT2DCR = 0x01 << 10 //+ External Event 9 Capture
	EXEV10CPT CPT2DCR = 0x01 << 11 //+ External Event 10 Capture
	TA1SET    CPT2DCR = 0x01 << 12 //+ Timer A output 1 Set
	TA1RST    CPT2DCR = 0x01 << 13 //+ Timer A output 1 Reset
	TACMP1    CPT2DCR = 0x01 << 14 //+ Timer A Compare 1
	TACMP2    CPT2DCR = 0x01 << 15 //+ Timer A Compare 2
	TB1SET    CPT2DCR = 0x01 << 16 //+ Timer B output 1 Set
	TB1RST    CPT2DCR = 0x01 << 17 //+ Timer B output 1 Reset
	TBCMP1    CPT2DCR = 0x01 << 18 //+ Timer B Compare 1
	TBCMP2    CPT2DCR = 0x01 << 19 //+ Timer B Compare 2
	TC1SET    CPT2DCR = 0x01 << 20 //+ Timer C output 1 Set
	TC1RST    CPT2DCR = 0x01 << 21 //+ Timer C output 1 Reset
	TCCMP1    CPT2DCR = 0x01 << 22 //+ Timer C Compare 1
	TCCMP2    CPT2DCR = 0x01 << 23 //+ Timer C Compare 2
	TE1SET    CPT2DCR = 0x01 << 28 //+ Timer E output 1 Set
	TE1RST    CPT2DCR = 0x01 << 29 //+ Timer E output 1 Reset
	TECMP1    CPT2DCR = 0x01 << 30 //+ Timer E Compare 1
	TECMP2    CPT2DCR = 0x01 << 31 //+ Timer E Compare 2
)

const (
	SWCPTn     = 0
	UDPCPTn    = 1
	EXEV1CPTn  = 2
	EXEV2CPTn  = 3
	EXEV3CPTn  = 4
	EXEV4CPTn  = 5
	EXEV5CPTn  = 6
	EXEV6CPTn  = 7
	EXEV7CPTn  = 8
	EXEV8CPTn  = 9
	EXEV9CPTn  = 10
	EXEV10CPTn = 11
	TA1SETn    = 12
	TA1RSTn    = 13
	TACMP1n    = 14
	TACMP2n    = 15
	TB1SETn    = 16
	TB1RSTn    = 17
	TBCMP1n    = 18
	TBCMP2n    = 19
	TC1SETn    = 20
	TC1RSTn    = 21
	TCCMP1n    = 22
	TCCMP2n    = 23
	TE1SETn    = 28
	TE1RSTn    = 29
	TECMP1n    = 30
	TECMP2n    = 31
)

const (
	POL1     OUTDR = 0x01 << 1  //+ Output 1 polarity
	IDLEM1   OUTDR = 0x01 << 2  //+ Output 1 Idle mode
	IDLES1   OUTDR = 0x01 << 3  //+ Output 1 Idle State
	FAULT1   OUTDR = 0x03 << 4  //+ Output 1 Fault state
	CHP1     OUTDR = 0x01 << 6  //+ Output 1 Chopper enable
	DIDL1    OUTDR = 0x01 << 7  //+ Output 1 Deadtime upon burst mode Idle entry
	DTEN     OUTDR = 0x01 << 8  //+ Deadtime enable
	DLYPRTEN OUTDR = 0x01 << 9  //+ Delayed Protection Enable
	DLYPRT   OUTDR = 0x07 << 10 //+ Delayed Protection
	POL2     OUTDR = 0x01 << 17 //+ Output 2 polarity
	IDLEM2   OUTDR = 0x01 << 18 //+ Output 2 Idle mode
	IDLES2   OUTDR = 0x01 << 19 //+ Output 2 Idle State
	FAULT2   OUTDR = 0x03 << 20 //+ Output 2 Fault state
	CHP2     OUTDR = 0x01 << 22 //+ Output 2 Chopper enable
	DIDL2    OUTDR = 0x01 << 23 //+ Output 2 Deadtime upon burst mode Idle entry
)

const (
	POL1n     = 1
	IDLEM1n   = 2
	IDLES1n   = 3
	FAULT1n   = 4
	CHP1n     = 6
	DIDL1n    = 7
	DTENn     = 8
	DLYPRTENn = 9
	DLYPRTn   = 10
	POL2n     = 17
	IDLEM2n   = 18
	IDLES2n   = 19
	FAULT2n   = 20
	CHP2n     = 22
	DIDL2n    = 23
)

const (
	FLT1EN FLTDR = 0x01 << 0  //+ Fault 1 enable
	FLT2EN FLTDR = 0x01 << 1  //+ Fault 2 enable
	FLT3EN FLTDR = 0x01 << 2  //+ Fault 3 enable
	FLT4EN FLTDR = 0x01 << 3  //+ Fault 4 enable
	FLT5EN FLTDR = 0x01 << 4  //+ Fault 5 enable
	FLTLCK FLTDR = 0x01 << 31 //+ Fault sources Lock
)

const (
	FLT1ENn = 0
	FLT2ENn = 1
	FLT3ENn = 2
	FLT4ENn = 3
	FLT5ENn = 4
	FLTLCKn = 31
)
