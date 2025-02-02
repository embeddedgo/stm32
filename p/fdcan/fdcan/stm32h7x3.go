// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package fdcan provides access to the registers of the FDCAN peripheral.
//
// Instances:
//
//	FDCAN1  FDCAN1_BASE  -  FDCAN1_IT0,FDCAN1_IT1,FDCAN_CAL  FDCAN1
//	FDCAN2  FDCAN2_BASE  -  FDCAN2_IT0,FDCAN2_IT1            FDCAN1
//
// Registers:
//
//	0x000 32  FDCAN_CREL    FDCAN Core Release Register
//	0x004 32  FDCAN_ENDN    FDCAN Core Release Register
//	0x00C 32  FDCAN_DBTP    FDCAN Data Bit Timing and Prescaler Register
//	0x010 32  FDCAN_TEST    FDCAN Test Register
//	0x014 32  FDCAN_RWD     FDCAN RAM Watchdog Register
//	0x018 32  FDCAN_CCCR    FDCAN CC Control Register
//	0x01C 32  FDCAN_NBTP    FDCAN Nominal Bit Timing and Prescaler Register
//	0x020 32  FDCAN_TSCC    FDCAN Timestamp Counter Configuration Register
//	0x024 32  FDCAN_TSCV    FDCAN Timestamp Counter Value Register
//	0x028 32  FDCAN_TOCC    FDCAN Timeout Counter Configuration Register
//	0x02C 32  FDCAN_TOCV    FDCAN Timeout Counter Value Register
//	0x040 32  FDCAN_ECR     FDCAN Error Counter Register
//	0x044 32  FDCAN_PSR     FDCAN Protocol Status Register
//	0x048 32  FDCAN_TDCR    FDCAN Transmitter Delay Compensation Register
//	0x050 32  FDCAN_IR      FDCAN Interrupt Register
//	0x054 32  FDCAN_IE      FDCAN Interrupt Enable Register
//	0x058 32  FDCAN_ILS     FDCAN Interrupt Line Select Register
//	0x05C 32  FDCAN_ILE     FDCAN Interrupt Line Enable Register
//	0x080 32  FDCAN_GFC     FDCAN Global Filter Configuration Register
//	0x084 32  FDCAN_SIDFC   FDCAN Standard ID Filter Configuration Register
//	0x088 32  FDCAN_XIDFC   FDCAN Extended ID Filter Configuration Register
//	0x090 32  FDCAN_XIDAM   FDCAN Extended ID and Mask Register
//	0x094 32  FDCAN_HPMS    FDCAN High Priority Message Status Register
//	0x098 32  FDCAN_NDAT1   FDCAN New Data 1 Register
//	0x09C 32  FDCAN_NDAT2   FDCAN New Data 2 Register
//	0x0A0 32  FDCAN_RXF0C   FDCAN Rx FIFO 0 Configuration Register
//	0x0A4 32  FDCAN_RXF0S   FDCAN Rx FIFO 0 Status Register
//	0x0A8 32  FDCAN_RXF0A   CAN Rx FIFO 0 Acknowledge Register
//	0x0AC 32  FDCAN_RXBC    FDCAN Rx Buffer Configuration Register
//	0x0B0 32  FDCAN_RXF1C   FDCAN Rx FIFO 1 Configuration Register
//	0x0B4 32  FDCAN_RXF1S   FDCAN Rx FIFO 1 Status Register
//	0x0B8 32  FDCAN_RXF1A   FDCAN Rx FIFO 1 Acknowledge Register
//	0x0BC 32  FDCAN_RXESC   FDCAN Rx Buffer Element Size Configuration Register
//	0x0C0 32  FDCAN_TXBC    FDCAN Tx Buffer Configuration Register
//	0x0C4 32  FDCAN_TXFQS   FDCAN Tx FIFO/Queue Status Register
//	0x0C8 32  FDCAN_TXESC   FDCAN Tx Buffer Element Size Configuration Register
//	0x0CC 32  FDCAN_TXBRP   FDCAN Tx Buffer Request Pending Register
//	0x0D0 32  FDCAN_TXBAR   FDCAN Tx Buffer Add Request Register
//	0x0D4 32  FDCAN_TXBCR   FDCAN Tx Buffer Cancellation Request Register
//	0x0D8 32  FDCAN_TXBTO   FDCAN Tx Buffer Transmission Occurred Register
//	0x0DC 32  FDCAN_TXBCF   FDCAN Tx Buffer Cancellation Finished Register
//	0x0E0 32  FDCAN_TXBTIE  FDCAN Tx Buffer Transmission Interrupt Enable Register
//	0x0E4 32  FDCAN_TXBCIE  FDCAN Tx Buffer Cancellation Finished Interrupt Enable Register
//	0x0F0 32  FDCAN_TXEFC   FDCAN Tx Event FIFO Configuration Register
//	0x0F4 32  FDCAN_TXEFS   FDCAN Tx Event FIFO Status Register
//	0x0F8 32  FDCAN_TXEFA   FDCAN Tx Event FIFO Acknowledge Register
//	0x100 32  FDCAN_TTTMC   FDCAN TT Trigger Memory Configuration Register
//	0x104 32  FDCAN_TTRMC   FDCAN TT Reference Message Configuration Register
//	0x108 32  FDCAN_TTOCF   FDCAN TT Operation Configuration Register
//	0x10C 32  FDCAN_TTMLM   FDCAN TT Matrix Limits Register
//	0x110 32  FDCAN_TURCF   FDCAN TUR Configuration Register
//	0x114 32  FDCAN_TTOCN   FDCAN TT Operation Control Register
//	0x118 32  CAN_TTGTP     FDCAN TT Global Time Preset Register
//	0x11C 32  FDCAN_TTTMK   FDCAN TT Time Mark Register
//	0x120 32  FDCAN_TTIR    FDCAN TT Interrupt Register
//	0x124 32  FDCAN_TTIE    FDCAN TT Interrupt Enable Register
//	0x128 32  FDCAN_TTILS   FDCAN TT Interrupt Line Select Register
//	0x12C 32  FDCAN_TTOST   FDCAN TT Operation Status Register
//	0x130 32  FDCAN_TURNA   FDCAN TUR Numerator Actual Register
//	0x134 32  FDCAN_TTLGT   FDCAN TT Local and Global Time Register
//	0x138 32  FDCAN_TTCTC   FDCAN TT Cycle Time and Count Register
//	0x13C 32  FDCAN_TTCPT   FDCAN TT Capture Time Register
//	0x140 32  FDCAN_TTCSM   FDCAN TT Cycle Sync Mark Register
//	0x300 32  FDCAN_TTTS    FDCAN TT Trigger Select Register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package fdcan

const (
	DAY     FDCAN_CREL = 0xFF << 0  //+ Timestamp Day
	MON     FDCAN_CREL = 0xFF << 8  //+ Timestamp Month
	YEAR    FDCAN_CREL = 0x0F << 16 //+ Timestamp Year
	SUBSTEP FDCAN_CREL = 0x0F << 20 //+ Sub-step of Core release
	STEP    FDCAN_CREL = 0x0F << 24 //+ Step of Core release
	REL     FDCAN_CREL = 0x0F << 28 //+ Core release
)

const (
	DAYn     = 0
	MONn     = 8
	YEARn    = 16
	SUBSTEPn = 20
	STEPn    = 24
	RELn     = 28
)

const (
	ETV FDCAN_ENDN = 0xFFFFFFFF << 0 //+ Endiannes Test Value
)

const (
	ETVn = 0
)

const (
	DSJW   FDCAN_DBTP = 0x0F << 0  //+ Synchronization Jump Width
	DTSEG2 FDCAN_DBTP = 0x0F << 4  //+ Data time segment after sample point
	DTSEG1 FDCAN_DBTP = 0x1F << 8  //+ Data time segment after sample point
	DBRP   FDCAN_DBTP = 0x1F << 16 //+ Data BIt Rate Prescaler
	TDC    FDCAN_DBTP = 0x01 << 23 //+ Transceiver Delay Compensation
)

const (
	DSJWn   = 0
	DTSEG2n = 4
	DTSEG1n = 8
	DBRPn   = 16
	TDCn    = 23
)

const (
	LBCK FDCAN_TEST = 0x01 << 4 //+ Loop Back mode
	TX   FDCAN_TEST = 0x03 << 5 //+ Loop Back mode
	RX   FDCAN_TEST = 0x01 << 7 //+ Control of Transmit Pin
)

const (
	LBCKn = 4
	TXn   = 5
	RXn   = 7
)

const (
	WDC FDCAN_RWD = 0xFF << 0 //+ Watchdog configuration
	WDV FDCAN_RWD = 0xFF << 8 //+ Watchdog value
)

const (
	WDCn = 0
	WDVn = 8
)

const (
	INIT FDCAN_CCCR = 0x01 << 0  //+ Initialization
	CCE  FDCAN_CCCR = 0x01 << 1  //+ Configuration Change Enable
	ASM  FDCAN_CCCR = 0x01 << 2  //+ ASM Restricted Operation Mode
	CSA  FDCAN_CCCR = 0x01 << 3  //+ Clock Stop Acknowledge
	CSR  FDCAN_CCCR = 0x01 << 4  //+ Clock Stop Request
	MON  FDCAN_CCCR = 0x01 << 5  //+ Bus Monitoring Mode
	DAR  FDCAN_CCCR = 0x01 << 6  //+ Disable Automatic Retransmission
	TEST FDCAN_CCCR = 0x01 << 7  //+ Test Mode Enable
	FDOE FDCAN_CCCR = 0x01 << 8  //+ FD Operation Enable
	BSE  FDCAN_CCCR = 0x01 << 9  //+ FDCAN Bit Rate Switching
	PXHD FDCAN_CCCR = 0x01 << 12 //+ Protocol Exception Handling Disable
	EFBI FDCAN_CCCR = 0x01 << 13 //+ Edge Filtering during Bus Integration
	TXP  FDCAN_CCCR = 0x01 << 14 //+ TXP
	NISO FDCAN_CCCR = 0x01 << 15 //+ Non ISO Operation
)

const (
	INITn = 0
	CCEn  = 1
	ASMn  = 2
	CSAn  = 3
	CSRn  = 4
	MONn  = 5
	DARn  = 6
	TESTn = 7
	FDOEn = 8
	BSEn  = 9
	PXHDn = 12
	EFBIn = 13
	TXPn  = 14
	NISOn = 15
)

const (
	TSEG2  FDCAN_NBTP = 0x7F << 0   //+ Nominal Time segment after sample point
	NTSEG1 FDCAN_NBTP = 0xFF << 8   //+ Nominal Time segment before sample point
	NBRP   FDCAN_NBTP = 0x1FF << 16 //+ Bit Rate Prescaler
	NSJW   FDCAN_NBTP = 0x7F << 25  //+ NSJW: Nominal (Re)Synchronization Jump Width
)

const (
	TSEG2n  = 0
	NTSEG1n = 8
	NBRPn   = 16
	NSJWn   = 25
)

const (
	TSS FDCAN_TSCC = 0x03 << 0  //+ Timestamp Select
	TCP FDCAN_TSCC = 0x0F << 16 //+ Timestamp Counter Prescaler
)

const (
	TSSn = 0
	TCPn = 16
)

const (
	TSC FDCAN_TSCV = 0xFFFF << 0 //+ Timestamp Counter
)

const (
	TSCn = 0
)

const (
	ETOC FDCAN_TOCC = 0x01 << 0    //+ Enable Timeout Counter
	TOS  FDCAN_TOCC = 0x03 << 1    //+ Timeout Select
	TOP  FDCAN_TOCC = 0xFFFF << 16 //+ Timeout Period
)

const (
	ETOCn = 0
	TOSn  = 1
	TOPn  = 16
)

const (
	TOC FDCAN_TOCV = 0xFFFF << 0 //+ Timeout Counter
)

const (
	TOCn = 0
)

const (
	TEC  FDCAN_ECR = 0xFF << 0  //+ Transmit Error Counter
	TREC FDCAN_ECR = 0x7F << 8  //+ Receive Error Counter
	RP   FDCAN_ECR = 0x01 << 15 //+ Receive Error Passive
	CEL  FDCAN_ECR = 0xFF << 16 //+ AN Error Logging
)

const (
	TECn  = 0
	TRECn = 8
	RPn   = 15
	CELn  = 16
)

const (
	LEC  FDCAN_PSR = 0x07 << 0  //+ Last Error Code
	ACT  FDCAN_PSR = 0x03 << 3  //+ Activity
	EP   FDCAN_PSR = 0x01 << 5  //+ Error Passive
	EW   FDCAN_PSR = 0x01 << 6  //+ Warning Status
	BO   FDCAN_PSR = 0x01 << 7  //+ Bus_Off Status
	DLEC FDCAN_PSR = 0x07 << 8  //+ Data Last Error Code
	RESI FDCAN_PSR = 0x01 << 11 //+ ESI flag of last received FDCAN Message
	RBRS FDCAN_PSR = 0x01 << 12 //+ BRS flag of last received FDCAN Message
	REDL FDCAN_PSR = 0x01 << 13 //+ Received FDCAN Message
	PXE  FDCAN_PSR = 0x01 << 14 //+ Protocol Exception Event
	TDCV FDCAN_PSR = 0x7F << 16 //+ Transmitter Delay Compensation Value
)

const (
	LECn  = 0
	ACTn  = 3
	EPn   = 5
	EWn   = 6
	BOn   = 7
	DLECn = 8
	RESIn = 11
	RBRSn = 12
	REDLn = 13
	PXEn  = 14
	TDCVn = 16
)

const (
	TDCF FDCAN_TDCR = 0x7F << 0 //+ Transmitter Delay Compensation Filter Window Length
	TDCO FDCAN_TDCR = 0x7F << 8 //+ Transmitter Delay Compensation Offset
)

const (
	TDCFn = 0
	TDCOn = 8
)

const (
	RF0N FDCAN_IR = 0x01 << 0  //+ Rx FIFO 0 New Message
	RF0W FDCAN_IR = 0x01 << 1  //+ Rx FIFO 0 Full
	RF0F FDCAN_IR = 0x01 << 2  //+ Rx FIFO 0 Full
	RF0L FDCAN_IR = 0x01 << 3  //+ Rx FIFO 0 Message Lost
	RF1N FDCAN_IR = 0x01 << 4  //+ Rx FIFO 1 New Message
	RF1W FDCAN_IR = 0x01 << 5  //+ Rx FIFO 1 Watermark Reached
	RF1F FDCAN_IR = 0x01 << 6  //+ Rx FIFO 1 Watermark Reached
	RF1L FDCAN_IR = 0x01 << 7  //+ Rx FIFO 1 Message Lost
	HPM  FDCAN_IR = 0x01 << 8  //+ High Priority Message
	TC   FDCAN_IR = 0x01 << 9  //+ Transmission Completed
	TCF  FDCAN_IR = 0x01 << 10 //+ Transmission Cancellation Finished
	TEF  FDCAN_IR = 0x01 << 11 //+ Tx FIFO Empty
	TEFN FDCAN_IR = 0x01 << 12 //+ Tx Event FIFO New Entry
	TEFW FDCAN_IR = 0x01 << 13 //+ Tx Event FIFO Watermark Reached
	TEFF FDCAN_IR = 0x01 << 14 //+ Tx Event FIFO Full
	TEFL FDCAN_IR = 0x01 << 15 //+ Tx Event FIFO Element Lost
	TSW  FDCAN_IR = 0x01 << 16 //+ Timestamp Wraparound
	MRAF FDCAN_IR = 0x01 << 17 //+ Message RAM Access Failure
	TOO  FDCAN_IR = 0x01 << 18 //+ Timeout Occurred
	DRX  FDCAN_IR = 0x01 << 19 //+ Message stored to Dedicated Rx Buffer
	ELO  FDCAN_IR = 0x01 << 22 //+ Error Logging Overflow
	EP   FDCAN_IR = 0x01 << 23 //+ Error Passive
	EW   FDCAN_IR = 0x01 << 24 //+ Warning Status
	BO   FDCAN_IR = 0x01 << 25 //+ Bus_Off Status
	WDI  FDCAN_IR = 0x01 << 26 //+ Watchdog Interrupt
	PEA  FDCAN_IR = 0x01 << 27 //+ Protocol Error in Arbitration Phase (Nominal Bit Time is used)
	PED  FDCAN_IR = 0x01 << 28 //+ Protocol Error in Data Phase (Data Bit Time is used)
	ARA  FDCAN_IR = 0x01 << 29 //+ Access to Reserved Address
)

const (
	RF0Nn = 0
	RF0Wn = 1
	RF0Fn = 2
	RF0Ln = 3
	RF1Nn = 4
	RF1Wn = 5
	RF1Fn = 6
	RF1Ln = 7
	HPMn  = 8
	TCn   = 9
	TCFn  = 10
	TEFn  = 11
	TEFNn = 12
	TEFWn = 13
	TEFFn = 14
	TEFLn = 15
	TSWn  = 16
	MRAFn = 17
	TOOn  = 18
	DRXn  = 19
	ELOn  = 22
	EPn   = 23
	EWn   = 24
	BOn   = 25
	WDIn  = 26
	PEAn  = 27
	PEDn  = 28
	ARAn  = 29
)

const (
	RF0NE FDCAN_IE = 0x01 << 0  //+ Rx FIFO 0 New Message Enable
	RF0WE FDCAN_IE = 0x01 << 1  //+ Rx FIFO 0 Full Enable
	RF0FE FDCAN_IE = 0x01 << 2  //+ Rx FIFO 0 Full Enable
	RF0LE FDCAN_IE = 0x01 << 3  //+ Rx FIFO 0 Message Lost Enable
	RF1NE FDCAN_IE = 0x01 << 4  //+ Rx FIFO 1 New Message Enable
	RF1WE FDCAN_IE = 0x01 << 5  //+ Rx FIFO 1 Watermark Reached Enable
	RF1FE FDCAN_IE = 0x01 << 6  //+ Rx FIFO 1 Watermark Reached Enable
	RF1LE FDCAN_IE = 0x01 << 7  //+ Rx FIFO 1 Message Lost Enable
	HPME  FDCAN_IE = 0x01 << 8  //+ High Priority Message Enable
	TCE   FDCAN_IE = 0x01 << 9  //+ Transmission Completed Enable
	TCFE  FDCAN_IE = 0x01 << 10 //+ Transmission Cancellation Finished Enable
	TEFE  FDCAN_IE = 0x01 << 11 //+ Tx FIFO Empty Enable
	TEFNE FDCAN_IE = 0x01 << 12 //+ Tx Event FIFO New Entry Enable
	TEFWE FDCAN_IE = 0x01 << 13 //+ Tx Event FIFO Watermark Reached Enable
	TEFFE FDCAN_IE = 0x01 << 14 //+ Tx Event FIFO Full Enable
	TEFLE FDCAN_IE = 0x01 << 15 //+ Tx Event FIFO Element Lost Enable
	TSWE  FDCAN_IE = 0x01 << 16 //+ Timestamp Wraparound Enable
	MRAFE FDCAN_IE = 0x01 << 17 //+ Message RAM Access Failure Enable
	TOOE  FDCAN_IE = 0x01 << 18 //+ Timeout Occurred Enable
	DRXE  FDCAN_IE = 0x01 << 19 //+ Message stored to Dedicated Rx Buffer Enable
	BECE  FDCAN_IE = 0x01 << 20 //+ Bit Error Corrected Interrupt Enable
	BEUE  FDCAN_IE = 0x01 << 21 //+ Bit Error Uncorrected Interrupt Enable
	ELOE  FDCAN_IE = 0x01 << 22 //+ Error Logging Overflow Enable
	EPE   FDCAN_IE = 0x01 << 23 //+ Error Passive Enable
	EWE   FDCAN_IE = 0x01 << 24 //+ Warning Status Enable
	BOE   FDCAN_IE = 0x01 << 25 //+ Bus_Off Status Enable
	WDIE  FDCAN_IE = 0x01 << 26 //+ Watchdog Interrupt Enable
	PEAE  FDCAN_IE = 0x01 << 27 //+ Protocol Error in Arbitration Phase Enable
	PEDE  FDCAN_IE = 0x01 << 28 //+ Protocol Error in Data Phase Enable
	ARAE  FDCAN_IE = 0x01 << 29 //+ Access to Reserved Address Enable
)

const (
	RF0NEn = 0
	RF0WEn = 1
	RF0FEn = 2
	RF0LEn = 3
	RF1NEn = 4
	RF1WEn = 5
	RF1FEn = 6
	RF1LEn = 7
	HPMEn  = 8
	TCEn   = 9
	TCFEn  = 10
	TEFEn  = 11
	TEFNEn = 12
	TEFWEn = 13
	TEFFEn = 14
	TEFLEn = 15
	TSWEn  = 16
	MRAFEn = 17
	TOOEn  = 18
	DRXEn  = 19
	BECEn  = 20
	BEUEn  = 21
	ELOEn  = 22
	EPEn   = 23
	EWEn   = 24
	BOEn   = 25
	WDIEn  = 26
	PEAEn  = 27
	PEDEn  = 28
	ARAEn  = 29
)

const (
	RF0NL FDCAN_ILS = 0x01 << 0  //+ Rx FIFO 0 New Message Interrupt Line
	RF0WL FDCAN_ILS = 0x01 << 1  //+ Rx FIFO 0 Watermark Reached Interrupt Line
	RF0FL FDCAN_ILS = 0x01 << 2  //+ Rx FIFO 0 Full Interrupt Line
	RF0LL FDCAN_ILS = 0x01 << 3  //+ Rx FIFO 0 Message Lost Interrupt Line
	RF1NL FDCAN_ILS = 0x01 << 4  //+ Rx FIFO 1 New Message Interrupt Line
	RF1WL FDCAN_ILS = 0x01 << 5  //+ Rx FIFO 1 Watermark Reached Interrupt Line
	RF1FL FDCAN_ILS = 0x01 << 6  //+ Rx FIFO 1 Full Interrupt Line
	RF1LL FDCAN_ILS = 0x01 << 7  //+ Rx FIFO 1 Message Lost Interrupt Line
	HPML  FDCAN_ILS = 0x01 << 8  //+ High Priority Message Interrupt Line
	TCL   FDCAN_ILS = 0x01 << 9  //+ Transmission Completed Interrupt Line
	TCFL  FDCAN_ILS = 0x01 << 10 //+ Transmission Cancellation Finished Interrupt Line
	TEFL  FDCAN_ILS = 0x01 << 11 //+ Tx FIFO Empty Interrupt Line
	TEFNL FDCAN_ILS = 0x01 << 12 //+ Tx Event FIFO New Entry Interrupt Line
	TEFWL FDCAN_ILS = 0x01 << 13 //+ Tx Event FIFO Watermark Reached Interrupt Line
	TEFFL FDCAN_ILS = 0x01 << 14 //+ Tx Event FIFO Full Interrupt Line
	TEFLL FDCAN_ILS = 0x01 << 15 //+ Tx Event FIFO Element Lost Interrupt Line
	TSWL  FDCAN_ILS = 0x01 << 16 //+ Timestamp Wraparound Interrupt Line
	MRAFL FDCAN_ILS = 0x01 << 17 //+ Message RAM Access Failure Interrupt Line
	TOOL  FDCAN_ILS = 0x01 << 18 //+ Timeout Occurred Interrupt Line
	DRXL  FDCAN_ILS = 0x01 << 19 //+ Message stored to Dedicated Rx Buffer Interrupt Line
	BECL  FDCAN_ILS = 0x01 << 20 //+ Bit Error Corrected Interrupt Line
	BEUL  FDCAN_ILS = 0x01 << 21 //+ Bit Error Uncorrected Interrupt Line
	ELOL  FDCAN_ILS = 0x01 << 22 //+ Error Logging Overflow Interrupt Line
	EPL   FDCAN_ILS = 0x01 << 23 //+ Error Passive Interrupt Line
	EWL   FDCAN_ILS = 0x01 << 24 //+ Warning Status Interrupt Line
	BOL   FDCAN_ILS = 0x01 << 25 //+ Bus_Off Status
	WDIL  FDCAN_ILS = 0x01 << 26 //+ Watchdog Interrupt Line
	PEAL  FDCAN_ILS = 0x01 << 27 //+ Protocol Error in Arbitration Phase Line
	PEDL  FDCAN_ILS = 0x01 << 28 //+ Protocol Error in Data Phase Line
	ARAL  FDCAN_ILS = 0x01 << 29 //+ Access to Reserved Address Line
)

const (
	RF0NLn = 0
	RF0WLn = 1
	RF0FLn = 2
	RF0LLn = 3
	RF1NLn = 4
	RF1WLn = 5
	RF1FLn = 6
	RF1LLn = 7
	HPMLn  = 8
	TCLn   = 9
	TCFLn  = 10
	TEFLn  = 11
	TEFNLn = 12
	TEFWLn = 13
	TEFFLn = 14
	TEFLLn = 15
	TSWLn  = 16
	MRAFLn = 17
	TOOLn  = 18
	DRXLn  = 19
	BECLn  = 20
	BEULn  = 21
	ELOLn  = 22
	EPLn   = 23
	EWLn   = 24
	BOLn   = 25
	WDILn  = 26
	PEALn  = 27
	PEDLn  = 28
	ARALn  = 29
)

const (
	EINT0 FDCAN_ILE = 0x01 << 0 //+ Enable Interrupt Line 0
	EINT1 FDCAN_ILE = 0x01 << 1 //+ Enable Interrupt Line 1
)

const (
	EINT0n = 0
	EINT1n = 1
)

const (
	RRFE FDCAN_GFC = 0x01 << 0 //+ Reject Remote Frames Extended
	RRFS FDCAN_GFC = 0x01 << 1 //+ Reject Remote Frames Standard
	ANFE FDCAN_GFC = 0x03 << 2 //+ Accept Non-matching Frames Extended
	ANFS FDCAN_GFC = 0x03 << 4 //+ Accept Non-matching Frames Standard
)

const (
	RRFEn = 0
	RRFSn = 1
	ANFEn = 2
	ANFSn = 4
)

const (
	FLSSA FDCAN_SIDFC = 0x3FFF << 2 //+ Filter List Standard Start Address
	LSS   FDCAN_SIDFC = 0xFF << 16  //+ List Size Standard
)

const (
	FLSSAn = 2
	LSSn   = 16
)

const (
	FLESA FDCAN_XIDFC = 0x3FFF << 2 //+ Filter List Standard Start Address
	LSE   FDCAN_XIDFC = 0xFF << 16  //+ List Size Extended
)

const (
	FLESAn = 2
	LSEn   = 16
)

const (
	EIDM FDCAN_XIDAM = 0x1FFFFFFF << 0 //+ Extended ID Mask
)

const (
	EIDMn = 0
)

const (
	BIDX FDCAN_HPMS = 0x3F << 0  //+ Buffer Index
	MSI  FDCAN_HPMS = 0x03 << 6  //+ Message Storage Indicator
	FIDX FDCAN_HPMS = 0x7F << 8  //+ Filter Index
	FLST FDCAN_HPMS = 0x01 << 15 //+ Filter List
)

const (
	BIDXn = 0
	MSIn  = 6
	FIDXn = 8
	FLSTn = 15
)

const (
	ND0  FDCAN_NDAT1 = 0x01 << 0  //+ New data
	ND1  FDCAN_NDAT1 = 0x01 << 1  //+ New data
	ND2  FDCAN_NDAT1 = 0x01 << 2  //+ New data
	ND3  FDCAN_NDAT1 = 0x01 << 3  //+ New data
	ND4  FDCAN_NDAT1 = 0x01 << 4  //+ New data
	ND5  FDCAN_NDAT1 = 0x01 << 5  //+ New data
	ND6  FDCAN_NDAT1 = 0x01 << 6  //+ New data
	ND7  FDCAN_NDAT1 = 0x01 << 7  //+ New data
	ND8  FDCAN_NDAT1 = 0x01 << 8  //+ New data
	ND9  FDCAN_NDAT1 = 0x01 << 9  //+ New data
	ND10 FDCAN_NDAT1 = 0x01 << 10 //+ New data
	ND11 FDCAN_NDAT1 = 0x01 << 11 //+ New data
	ND12 FDCAN_NDAT1 = 0x01 << 12 //+ New data
	ND13 FDCAN_NDAT1 = 0x01 << 13 //+ New data
	ND14 FDCAN_NDAT1 = 0x01 << 14 //+ New data
	ND15 FDCAN_NDAT1 = 0x01 << 15 //+ New data
	ND16 FDCAN_NDAT1 = 0x01 << 16 //+ New data
	ND17 FDCAN_NDAT1 = 0x01 << 17 //+ New data
	ND18 FDCAN_NDAT1 = 0x01 << 18 //+ New data
	ND19 FDCAN_NDAT1 = 0x01 << 19 //+ New data
	ND20 FDCAN_NDAT1 = 0x01 << 20 //+ New data
	ND21 FDCAN_NDAT1 = 0x01 << 21 //+ New data
	ND22 FDCAN_NDAT1 = 0x01 << 22 //+ New data
	ND23 FDCAN_NDAT1 = 0x01 << 23 //+ New data
	ND24 FDCAN_NDAT1 = 0x01 << 24 //+ New data
	ND25 FDCAN_NDAT1 = 0x01 << 25 //+ New data
	ND26 FDCAN_NDAT1 = 0x01 << 26 //+ New data
	ND27 FDCAN_NDAT1 = 0x01 << 27 //+ New data
	ND28 FDCAN_NDAT1 = 0x01 << 28 //+ New data
	ND29 FDCAN_NDAT1 = 0x01 << 29 //+ New data
	ND30 FDCAN_NDAT1 = 0x01 << 30 //+ New data
	ND31 FDCAN_NDAT1 = 0x01 << 31 //+ New data
)

const (
	ND0n  = 0
	ND1n  = 1
	ND2n  = 2
	ND3n  = 3
	ND4n  = 4
	ND5n  = 5
	ND6n  = 6
	ND7n  = 7
	ND8n  = 8
	ND9n  = 9
	ND10n = 10
	ND11n = 11
	ND12n = 12
	ND13n = 13
	ND14n = 14
	ND15n = 15
	ND16n = 16
	ND17n = 17
	ND18n = 18
	ND19n = 19
	ND20n = 20
	ND21n = 21
	ND22n = 22
	ND23n = 23
	ND24n = 24
	ND25n = 25
	ND26n = 26
	ND27n = 27
	ND28n = 28
	ND29n = 29
	ND30n = 30
	ND31n = 31
)

const (
	ND32 FDCAN_NDAT2 = 0x01 << 0  //+ New data
	ND33 FDCAN_NDAT2 = 0x01 << 1  //+ New data
	ND34 FDCAN_NDAT2 = 0x01 << 2  //+ New data
	ND35 FDCAN_NDAT2 = 0x01 << 3  //+ New data
	ND36 FDCAN_NDAT2 = 0x01 << 4  //+ New data
	ND37 FDCAN_NDAT2 = 0x01 << 5  //+ New data
	ND38 FDCAN_NDAT2 = 0x01 << 6  //+ New data
	ND39 FDCAN_NDAT2 = 0x01 << 7  //+ New data
	ND40 FDCAN_NDAT2 = 0x01 << 8  //+ New data
	ND41 FDCAN_NDAT2 = 0x01 << 9  //+ New data
	ND42 FDCAN_NDAT2 = 0x01 << 10 //+ New data
	ND43 FDCAN_NDAT2 = 0x01 << 11 //+ New data
	ND44 FDCAN_NDAT2 = 0x01 << 12 //+ New data
	ND45 FDCAN_NDAT2 = 0x01 << 13 //+ New data
	ND46 FDCAN_NDAT2 = 0x01 << 14 //+ New data
	ND47 FDCAN_NDAT2 = 0x01 << 15 //+ New data
	ND48 FDCAN_NDAT2 = 0x01 << 16 //+ New data
	ND49 FDCAN_NDAT2 = 0x01 << 17 //+ New data
	ND50 FDCAN_NDAT2 = 0x01 << 18 //+ New data
	ND51 FDCAN_NDAT2 = 0x01 << 19 //+ New data
	ND52 FDCAN_NDAT2 = 0x01 << 20 //+ New data
	ND53 FDCAN_NDAT2 = 0x01 << 21 //+ New data
	ND54 FDCAN_NDAT2 = 0x01 << 22 //+ New data
	ND55 FDCAN_NDAT2 = 0x01 << 23 //+ New data
	ND56 FDCAN_NDAT2 = 0x01 << 24 //+ New data
	ND57 FDCAN_NDAT2 = 0x01 << 25 //+ New data
	ND58 FDCAN_NDAT2 = 0x01 << 26 //+ New data
	ND59 FDCAN_NDAT2 = 0x01 << 27 //+ New data
	ND60 FDCAN_NDAT2 = 0x01 << 28 //+ New data
	ND61 FDCAN_NDAT2 = 0x01 << 29 //+ New data
	ND62 FDCAN_NDAT2 = 0x01 << 30 //+ New data
	ND63 FDCAN_NDAT2 = 0x01 << 31 //+ New data
)

const (
	ND32n = 0
	ND33n = 1
	ND34n = 2
	ND35n = 3
	ND36n = 4
	ND37n = 5
	ND38n = 6
	ND39n = 7
	ND40n = 8
	ND41n = 9
	ND42n = 10
	ND43n = 11
	ND44n = 12
	ND45n = 13
	ND46n = 14
	ND47n = 15
	ND48n = 16
	ND49n = 17
	ND50n = 18
	ND51n = 19
	ND52n = 20
	ND53n = 21
	ND54n = 22
	ND55n = 23
	ND56n = 24
	ND57n = 25
	ND58n = 26
	ND59n = 27
	ND60n = 28
	ND61n = 29
	ND62n = 30
	ND63n = 31
)

const (
	F0SA FDCAN_RXF0C = 0x3FFF << 2 //+ Rx FIFO 0 Start Address
	F0S  FDCAN_RXF0C = 0xFF << 16  //+ Rx FIFO 0 Size
	F0WM FDCAN_RXF0C = 0xFF << 24  //+ FIFO 0 Watermark
)

const (
	F0SAn = 2
	F0Sn  = 16
	F0WMn = 24
)

const (
	F0FL FDCAN_RXF0S = 0x7F << 0  //+ Rx FIFO 0 Fill Level
	F0G  FDCAN_RXF0S = 0x3F << 8  //+ Rx FIFO 0 Get Index
	F0P  FDCAN_RXF0S = 0x3F << 16 //+ Rx FIFO 0 Put Index
	F0F  FDCAN_RXF0S = 0x01 << 24 //+ Rx FIFO 0 Full
	RF0L FDCAN_RXF0S = 0x01 << 25 //+ Rx FIFO 0 Message Lost
)

const (
	F0FLn = 0
	F0Gn  = 8
	F0Pn  = 16
	F0Fn  = 24
	RF0Ln = 25
)

const (
	FA01 FDCAN_RXF0A = 0x3F << 0 //+ Rx FIFO 0 Acknowledge Index
)

const (
	FA01n = 0
)

const (
	RBSA FDCAN_RXBC = 0x3FFF << 2 //+ Rx Buffer Start Address
)

const (
	RBSAn = 2
)

const (
	F1SA FDCAN_RXF1C = 0x3FFF << 2 //+ Rx FIFO 1 Start Address
	F1S  FDCAN_RXF1C = 0x7F << 16  //+ Rx FIFO 1 Size
	F1WM FDCAN_RXF1C = 0x7F << 24  //+ Rx FIFO 1 Watermark
)

const (
	F1SAn = 2
	F1Sn  = 16
	F1WMn = 24
)

const (
	F1FL FDCAN_RXF1S = 0x7F << 0  //+ Rx FIFO 1 Fill Level
	F1GI FDCAN_RXF1S = 0x7F << 8  //+ Rx FIFO 1 Get Index
	F1PI FDCAN_RXF1S = 0x7F << 16 //+ Rx FIFO 1 Put Index
	F1F  FDCAN_RXF1S = 0x01 << 24 //+ Rx FIFO 1 Full
	RF1L FDCAN_RXF1S = 0x01 << 25 //+ Rx FIFO 1 Message Lost
	DMS  FDCAN_RXF1S = 0x03 << 30 //+ Debug Message Status
)

const (
	F1FLn = 0
	F1GIn = 8
	F1PIn = 16
	F1Fn  = 24
	RF1Ln = 25
	DMSn  = 30
)

const (
	F1AI FDCAN_RXF1A = 0x3F << 0 //+ Rx FIFO 1 Acknowledge Index
)

const (
	F1AIn = 0
)

const (
	F0DS FDCAN_RXESC = 0x07 << 0 //+ Rx FIFO 1 Data Field Size:
	F1DS FDCAN_RXESC = 0x07 << 4 //+ Rx FIFO 0 Data Field Size:
	RBDS FDCAN_RXESC = 0x07 << 8 //+ Rx Buffer Data Field Size:
)

const (
	F0DSn = 0
	F1DSn = 4
	RBDSn = 8
)

const (
	TBSA FDCAN_TXBC = 0x3FFF << 2 //+ Tx Buffers Start Address
	NDTB FDCAN_TXBC = 0x3F << 16  //+ Number of Dedicated Transmit Buffers
	TFQS FDCAN_TXBC = 0x3F << 24  //+ Transmit FIFO/Queue Size
	TFQM FDCAN_TXBC = 0x01 << 30  //+ Tx FIFO/Queue Mode
)

const (
	TBSAn = 2
	NDTBn = 16
	TFQSn = 24
	TFQMn = 30
)

const (
	TFFL  FDCAN_TXFQS = 0x3F << 0  //+ Tx FIFO Free Level
	TFGI  FDCAN_TXFQS = 0x1F << 8  //+ TFGI
	TFQPI FDCAN_TXFQS = 0x1F << 16 //+ Tx FIFO/Queue Put Index
	TFQF  FDCAN_TXFQS = 0x01 << 21 //+ Tx FIFO/Queue Full
)

const (
	TFFLn  = 0
	TFGIn  = 8
	TFQPIn = 16
	TFQFn  = 21
)

const (
	TBDS FDCAN_TXESC = 0x07 << 0 //+ Tx Buffer Data Field Size:
)

const (
	TBDSn = 0
)

const (
	TRP FDCAN_TXBRP = 0xFFFFFFFF << 0 //+ Transmission Request Pending
)

const (
	TRPn = 0
)

const (
	AR FDCAN_TXBAR = 0xFFFFFFFF << 0 //+ Add Request
)

const (
	ARn = 0
)

const (
	CR FDCAN_TXBCR = 0xFFFFFFFF << 0 //+ Cancellation Request
)

const (
	CRn = 0
)

const (
	TO FDCAN_TXBTO = 0xFFFFFFFF << 0 //+ Transmission Occurred.
)

const (
	TOn = 0
)

const (
	CF FDCAN_TXBCF = 0xFFFFFFFF << 0 //+ Cancellation Finished
)

const (
	CFn = 0
)

const (
	TIE FDCAN_TXBTIE = 0xFFFFFFFF << 0 //+ Transmission Interrupt Enable
)

const (
	TIEn = 0
)

const (
	CF FDCAN_TXBCIE = 0xFFFFFFFF << 0 //+ Cancellation Finished Interrupt Enable
)

const (
	CFn = 0
)

const (
	EFSA FDCAN_TXEFC = 0x3FFF << 2 //+ Event FIFO Start Address
	EFS  FDCAN_TXEFC = 0x3F << 16  //+ Event FIFO Size
	EFWM FDCAN_TXEFC = 0x3F << 24  //+ Event FIFO Watermark
)

const (
	EFSAn = 2
	EFSn  = 16
	EFWMn = 24
)

const (
	EFFL FDCAN_TXEFS = 0x3F << 0  //+ Event FIFO Fill Level
	EFGI FDCAN_TXEFS = 0x1F << 8  //+ Event FIFO Get Index.
	EFPI FDCAN_TXEFS = 0x1F << 16 //+ Event FIFO put index.
	EFF  FDCAN_TXEFS = 0x01 << 24 //+ Event FIFO Full.
	TEFL FDCAN_TXEFS = 0x01 << 25 //+ Tx Event FIFO Element Lost.
)

const (
	EFFLn = 0
	EFGIn = 8
	EFPIn = 16
	EFFn  = 24
	TEFLn = 25
)

const (
	EFAI FDCAN_TXEFA = 0x1F << 0 //+ Event FIFO Acknowledge Index
)

const (
	EFAIn = 0
)

const (
	TMSA FDCAN_TTTMC = 0x3FFF << 2 //+ Trigger Memory Start Address
	TME  FDCAN_TTTMC = 0x7F << 16  //+ Trigger Memory Elements
)

const (
	TMSAn = 2
	TMEn  = 16
)

const (
	RID  FDCAN_TTRMC = 0x1FFFFFFF << 0 //+ Reference Identifier.
	XTD  FDCAN_TTRMC = 0x01 << 30      //+ Extended Identifier
	RMPS FDCAN_TTRMC = 0x01 << 31      //+ Reference Message Payload Select
)

const (
	RIDn  = 0
	XTDn  = 30
	RMPSn = 31
)

const (
	OM    FDCAN_TTOCF = 0x03 << 0  //+ Operation Mode
	GEN   FDCAN_TTOCF = 0x01 << 3  //+ Gap Enable
	TM    FDCAN_TTOCF = 0x01 << 4  //+ Time Master
	LDSDL FDCAN_TTOCF = 0x07 << 5  //+ LD of Synchronization Deviation Limit
	IRTO  FDCAN_TTOCF = 0x7F << 8  //+ Initial Reference Trigger Offset
	EECS  FDCAN_TTOCF = 0x01 << 15 //+ Enable External Clock Synchronization
	AWL   FDCAN_TTOCF = 0xFF << 16 //+ Application Watchdog Limit
	EGTF  FDCAN_TTOCF = 0x01 << 24 //+ Enable Global Time Filtering
	ECC   FDCAN_TTOCF = 0x01 << 25 //+ Enable Clock Calibration
	EVTP  FDCAN_TTOCF = 0x01 << 26 //+ Event Trigger Polarity
)

const (
	OMn    = 0
	GENn   = 3
	TMn    = 4
	LDSDLn = 5
	IRTOn  = 8
	EECSn  = 15
	AWLn   = 16
	EGTFn  = 24
	ECCn   = 25
	EVTPn  = 26
)

const (
	CCM  FDCAN_TTMLM = 0x3F << 0   //+ Cycle Count Max
	CSS  FDCAN_TTMLM = 0x03 << 6   //+ Cycle Start Synchronization
	TXEW FDCAN_TTMLM = 0x0F << 8   //+ Tx Enable Window
	ENTT FDCAN_TTMLM = 0xFFF << 16 //+ Expected Number of Tx Triggers
)

const (
	CCMn  = 0
	CSSn  = 6
	TXEWn = 8
	ENTTn = 16
)

const (
	NCL FDCAN_TURCF = 0xFFFF << 0  //+ Numerator Configuration Low.
	DC  FDCAN_TURCF = 0x3FFF << 16 //+ Denominator Configuration.
	ELT FDCAN_TURCF = 0x01 << 31   //+ Enable Local Time
)

const (
	NCLn = 0
	DCn  = 16
	ELTn = 31
)

const (
	SGT  FDCAN_TTOCN = 0x01 << 0  //+ Set Global time
	ECS  FDCAN_TTOCN = 0x01 << 1  //+ External Clock Synchronization
	SWP  FDCAN_TTOCN = 0x01 << 2  //+ Stop Watch Polarity
	SWS  FDCAN_TTOCN = 0x03 << 3  //+ Stop Watch Source.
	RTIE FDCAN_TTOCN = 0x01 << 5  //+ Register Time Mark Interrupt Pulse Enable
	TMC  FDCAN_TTOCN = 0x03 << 6  //+ Register Time Mark Compare
	TTIE FDCAN_TTOCN = 0x01 << 8  //+ Trigger Time Mark Interrupt Pulse Enable
	GCS  FDCAN_TTOCN = 0x01 << 9  //+ Gap Control Select
	FGP  FDCAN_TTOCN = 0x01 << 10 //+ Finish Gap.
	TMG  FDCAN_TTOCN = 0x01 << 11 //+ Time Mark Gap
	NIG  FDCAN_TTOCN = 0x01 << 12 //+ Next is Gap
	ESCN FDCAN_TTOCN = 0x01 << 13 //+ External Synchronization Control
	LCKC FDCAN_TTOCN = 0x01 << 15 //+ TT Operation Control Register Locked
)

const (
	SGTn  = 0
	ECSn  = 1
	SWPn  = 2
	SWSn  = 3
	RTIEn = 5
	TMCn  = 6
	TTIEn = 8
	GCSn  = 9
	FGPn  = 10
	TMGn  = 11
	NIGn  = 12
	ESCNn = 13
	LCKCn = 15
)

const (
	NCL CAN_TTGTP = 0xFFFF << 0  //+ Time Preset
	CTP CAN_TTGTP = 0xFFFF << 16 //+ Cycle Time Target Phase
)

const (
	NCLn = 0
	CTPn = 16
)

const (
	TM   FDCAN_TTTMK = 0xFFFF << 0 //+ Time Mark
	TICC FDCAN_TTTMK = 0x7F << 16  //+ Time Mark Cycle Code
	LCKM FDCAN_TTTMK = 0x01 << 31  //+ TT Time Mark Register Locked
)

const (
	TMn   = 0
	TICCn = 16
	LCKMn = 31
)

const (
	SBC  FDCAN_TTIR = 0x01 << 0  //+ Start of Basic Cycle
	SMC  FDCAN_TTIR = 0x01 << 1  //+ Start of Matrix Cycle
	CSM  FDCAN_TTIR = 0x01 << 2  //+ Change of Synchronization Mode
	SOG  FDCAN_TTIR = 0x01 << 3  //+ Start of Gap
	RTMI FDCAN_TTIR = 0x01 << 4  //+ Register Time Mark Interrupt.
	TTMI FDCAN_TTIR = 0x01 << 5  //+ Trigger Time Mark Event Internal
	SWE  FDCAN_TTIR = 0x01 << 6  //+ Stop Watch Event
	GTW  FDCAN_TTIR = 0x01 << 7  //+ Global Time Wrap
	GTD  FDCAN_TTIR = 0x01 << 8  //+ Global Time Discontinuity
	GTE  FDCAN_TTIR = 0x01 << 9  //+ Global Time Error
	TXU  FDCAN_TTIR = 0x01 << 10 //+ Tx Count Underflow
	TXO  FDCAN_TTIR = 0x01 << 11 //+ Tx Count Overflow
	SE1  FDCAN_TTIR = 0x01 << 12 //+ Scheduling Error 1
	SE2  FDCAN_TTIR = 0x01 << 13 //+ Scheduling Error 2
	ELC  FDCAN_TTIR = 0x01 << 14 //+ Error Level Changed.
	IWTG FDCAN_TTIR = 0x01 << 15 //+ Initialization Watch Trigger
	WT   FDCAN_TTIR = 0x01 << 16 //+ Watch Trigger
	AW   FDCAN_TTIR = 0x01 << 17 //+ Application Watchdog
	CER  FDCAN_TTIR = 0x01 << 18 //+ Configuration Error
)

const (
	SBCn  = 0
	SMCn  = 1
	CSMn  = 2
	SOGn  = 3
	RTMIn = 4
	TTMIn = 5
	SWEn  = 6
	GTWn  = 7
	GTDn  = 8
	GTEn  = 9
	TXUn  = 10
	TXOn  = 11
	SE1n  = 12
	SE2n  = 13
	ELCn  = 14
	IWTGn = 15
	WTn   = 16
	AWn   = 17
	CERn  = 18
)

const (
	SBCE  FDCAN_TTIE = 0x01 << 0  //+ Start of Basic Cycle Interrupt Enable
	SMCE  FDCAN_TTIE = 0x01 << 1  //+ Start of Matrix Cycle Interrupt Enable
	CSME  FDCAN_TTIE = 0x01 << 2  //+ Change of Synchronization Mode Interrupt Enable
	SOGE  FDCAN_TTIE = 0x01 << 3  //+ Start of Gap Interrupt Enable
	RTMIE FDCAN_TTIE = 0x01 << 4  //+ Register Time Mark Interrupt Enable
	TTMIE FDCAN_TTIE = 0x01 << 5  //+ Trigger Time Mark Event Internal Interrupt Enable
	SWEE  FDCAN_TTIE = 0x01 << 6  //+ Stop Watch Event Interrupt Enable
	GTWE  FDCAN_TTIE = 0x01 << 7  //+ Global Time Wrap Interrupt Enable
	GTDE  FDCAN_TTIE = 0x01 << 8  //+ Global Time Discontinuity Interrupt Enable
	GTEE  FDCAN_TTIE = 0x01 << 9  //+ Global Time Error Interrupt Enable
	TXUE  FDCAN_TTIE = 0x01 << 10 //+ Tx Count Underflow Interrupt Enable
	TXOE  FDCAN_TTIE = 0x01 << 11 //+ Tx Count Overflow Interrupt Enable
	SE1E  FDCAN_TTIE = 0x01 << 12 //+ Scheduling Error 1 Interrupt Enable
	SE2E  FDCAN_TTIE = 0x01 << 13 //+ Scheduling Error 2 Interrupt Enable
	ELCE  FDCAN_TTIE = 0x01 << 14 //+ Change Error Level Interrupt Enable
	IWTGE FDCAN_TTIE = 0x01 << 15 //+ Initialization Watch Trigger Interrupt Enable
	WTE   FDCAN_TTIE = 0x01 << 16 //+ Watch Trigger Interrupt Enable
	AWE   FDCAN_TTIE = 0x01 << 17 //+ Application Watchdog Interrupt Enable
	CERE  FDCAN_TTIE = 0x01 << 18 //+ Configuration Error Interrupt Enable
)

const (
	SBCEn  = 0
	SMCEn  = 1
	CSMEn  = 2
	SOGEn  = 3
	RTMIEn = 4
	TTMIEn = 5
	SWEEn  = 6
	GTWEn  = 7
	GTDEn  = 8
	GTEEn  = 9
	TXUEn  = 10
	TXOEn  = 11
	SE1En  = 12
	SE2En  = 13
	ELCEn  = 14
	IWTGEn = 15
	WTEn   = 16
	AWEn   = 17
	CEREn  = 18
)

const (
	SBCL  FDCAN_TTILS = 0x01 << 0  //+ Start of Basic Cycle Interrupt Line
	SMCL  FDCAN_TTILS = 0x01 << 1  //+ Start of Matrix Cycle Interrupt Line
	CSML  FDCAN_TTILS = 0x01 << 2  //+ Change of Synchronization Mode Interrupt Line
	SOGL  FDCAN_TTILS = 0x01 << 3  //+ Start of Gap Interrupt Line
	RTMIL FDCAN_TTILS = 0x01 << 4  //+ Register Time Mark Interrupt Line
	TTMIL FDCAN_TTILS = 0x01 << 5  //+ Trigger Time Mark Event Internal Interrupt Line
	SWEL  FDCAN_TTILS = 0x01 << 6  //+ Stop Watch Event Interrupt Line
	GTWL  FDCAN_TTILS = 0x01 << 7  //+ Global Time Wrap Interrupt Line
	GTDL  FDCAN_TTILS = 0x01 << 8  //+ Global Time Discontinuity Interrupt Line
	GTEL  FDCAN_TTILS = 0x01 << 9  //+ Global Time Error Interrupt Line
	TXUL  FDCAN_TTILS = 0x01 << 10 //+ Tx Count Underflow Interrupt Line
	TXOL  FDCAN_TTILS = 0x01 << 11 //+ Tx Count Overflow Interrupt Line
	SE1L  FDCAN_TTILS = 0x01 << 12 //+ Scheduling Error 1 Interrupt Line
	SE2L  FDCAN_TTILS = 0x01 << 13 //+ Scheduling Error 2 Interrupt Line
	ELCL  FDCAN_TTILS = 0x01 << 14 //+ Change Error Level Interrupt Line
	IWTGL FDCAN_TTILS = 0x01 << 15 //+ Initialization Watch Trigger Interrupt Line
	WTL   FDCAN_TTILS = 0x01 << 16 //+ Watch Trigger Interrupt Line
	AWL   FDCAN_TTILS = 0x01 << 17 //+ Application Watchdog Interrupt Line
	CERL  FDCAN_TTILS = 0x01 << 18 //+ Configuration Error Interrupt Line
)

const (
	SBCLn  = 0
	SMCLn  = 1
	CSMLn  = 2
	SOGLn  = 3
	RTMILn = 4
	TTMILn = 5
	SWELn  = 6
	GTWLn  = 7
	GTDLn  = 8
	GTELn  = 9
	TXULn  = 10
	TXOLn  = 11
	SE1Ln  = 12
	SE2Ln  = 13
	ELCLn  = 14
	IWTGLn = 15
	WTLn   = 16
	AWLn   = 17
	CERLn  = 18
)

const (
	EL   FDCAN_TTOST = 0x03 << 0  //+ Error Level
	MS   FDCAN_TTOST = 0x03 << 2  //+ Master State.
	SYS  FDCAN_TTOST = 0x03 << 4  //+ Synchronization State
	GTP  FDCAN_TTOST = 0x01 << 6  //+ Quality of Global Time Phase
	QCS  FDCAN_TTOST = 0x01 << 7  //+ Quality of Clock Speed
	RTO  FDCAN_TTOST = 0xFF << 8  //+ Reference Trigger Offset
	WGTD FDCAN_TTOST = 0x01 << 22 //+ Wait for Global Time Discontinuity
	GFI  FDCAN_TTOST = 0x01 << 23 //+ Gap Finished Indicator.
	TMP  FDCAN_TTOST = 0x07 << 24 //+ Time Master Priority
	GSI  FDCAN_TTOST = 0x01 << 27 //+ Gap Started Indicator.
	WFE  FDCAN_TTOST = 0x01 << 28 //+ Wait for Event
	AWE  FDCAN_TTOST = 0x01 << 29 //+ Application Watchdog Event
	WECS FDCAN_TTOST = 0x01 << 30 //+ Wait for External Clock Synchronization
	SPL  FDCAN_TTOST = 0x01 << 31 //+ Schedule Phase Lock
)

const (
	ELn   = 0
	MSn   = 2
	SYSn  = 4
	GTPn  = 6
	QCSn  = 7
	RTOn  = 8
	WGTDn = 22
	GFIn  = 23
	TMPn  = 24
	GSIn  = 27
	WFEn  = 28
	AWEn  = 29
	WECSn = 30
	SPLn  = 31
)

const (
	NAV FDCAN_TURNA = 0x3FFFF << 0 //+ Numerator Actual Value
)

const (
	NAVn = 0
)

const (
	LT FDCAN_TTLGT = 0xFFFF << 0  //+ Local Time
	GT FDCAN_TTLGT = 0xFFFF << 16 //+ Global Time
)

const (
	LTn = 0
	GTn = 16
)

const (
	CT FDCAN_TTCTC = 0xFFFF << 0 //+ Cycle Time
	CC FDCAN_TTCTC = 0x3F << 16  //+ Cycle Count
)

const (
	CTn = 0
	CCn = 16
)

const (
	CT  FDCAN_TTCPT = 0x3F << 0    //+ Cycle Count Value
	SWV FDCAN_TTCPT = 0xFFFF << 16 //+ Stop Watch Value
)

const (
	CTn  = 0
	SWVn = 16
)

const (
	CSM FDCAN_TTCSM = 0xFFFF << 0 //+ Cycle Sync Mark
)

const (
	CSMn = 0
)

const (
	SWTDEL FDCAN_TTTS = 0x03 << 0 //+ Stop watch trigger input selection
	EVTSEL FDCAN_TTTS = 0x03 << 4 //+ Event trigger input selection
)

const (
	SWTDELn = 0
	EVTSELn = 4
)
