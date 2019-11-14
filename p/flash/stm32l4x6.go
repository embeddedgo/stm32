// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32l4x6

// Package flash provides access to the registers of the FLASH peripheral.
//
// Instances:
//  FLASH  FLASH_BASE  AHB1  FLASH  Flash
// Registers:
//  0x000 32  ACR       Access control register
//  0x004 32  PDKEYR    Power down key register
//  0x008 32  KEYR      Flash key register
//  0x00C 32  OPTKEYR   Option byte key register
//  0x010 32  SR        Status register
//  0x014 32  CR        Flash control register
//  0x018 32  ECCR      Flash ECC register
//  0x020 32  OPTR      Flash option register
//  0x024 32  PCROP1SR  Flash Bank 1 PCROP Start address register
//  0x028 32  PCROP1ER  Flash Bank 1 PCROP End address register
//  0x02C 32  WRP1AR    Flash Bank 1 WRP area A address register
//  0x030 32  WRP1BR    Flash Bank 1 WRP area B address register
//  0x044 32  PCROP2SR  Flash Bank 2 PCROP Start address register
//  0x048 32  PCROP2ER  Flash Bank 2 PCROP End address register
//  0x04C 32  WRP2AR    Flash Bank 2 WRP area A address register
//  0x050 32  WRP2BR    Flash Bank 2 WRP area B address register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package flash

const (
	LATENCY  ACR = 0x07 << 0  //+ Latency
	PRFTEN   ACR = 0x01 << 8  //+ Prefetch enable
	ICEN     ACR = 0x01 << 9  //+ Instruction cache enable
	DCEN     ACR = 0x01 << 10 //+ Data cache enable
	ICRST    ACR = 0x01 << 11 //+ Instruction cache reset
	DCRST    ACR = 0x01 << 12 //+ Data cache reset
	RUN_PD   ACR = 0x01 << 13 //+ Flash Power-down mode during Low-power run mode
	SLEEP_PD ACR = 0x01 << 14 //+ Flash Power-down mode during Low-power sleep mode
)

const (
	LATENCYn  = 0
	PRFTENn   = 8
	ICENn     = 9
	DCENn     = 10
	ICRSTn    = 11
	DCRSTn    = 12
	RUN_PDn   = 13
	SLEEP_PDn = 14
)

const (
	EOP     SR = 0x01 << 0  //+ End of operation
	OPERR   SR = 0x01 << 1  //+ Operation error
	PROGERR SR = 0x01 << 3  //+ Programming error
	WRPERR  SR = 0x01 << 4  //+ Write protected error
	PGAERR  SR = 0x01 << 5  //+ Programming alignment error
	SIZERR  SR = 0x01 << 6  //+ Size error
	PGSERR  SR = 0x01 << 7  //+ Programming sequence error
	MISERR  SR = 0x01 << 8  //+ Fast programming data miss error
	FASTERR SR = 0x01 << 9  //+ Fast programming error
	RDERR   SR = 0x01 << 14 //+ PCROP read error
	OPTVERR SR = 0x01 << 15 //+ Option validity error
	BSY     SR = 0x01 << 16 //+ Busy
)

const (
	EOPn     = 0
	OPERRn   = 1
	PROGERRn = 3
	WRPERRn  = 4
	PGAERRn  = 5
	SIZERRn  = 6
	PGSERRn  = 7
	MISERRn  = 8
	FASTERRn = 9
	RDERRn   = 14
	OPTVERRn = 15
	BSYn     = 16
)

const (
	PG         CR = 0x01 << 0  //+ Programming
	PER        CR = 0x01 << 1  //+ Page erase
	MER1       CR = 0x01 << 2  //+ Bank 1 Mass erase
	PNB        CR = 0xFF << 3  //+ Page number
	BKER       CR = 0x01 << 11 //+ Bank erase
	MER2       CR = 0x01 << 15 //+ Bank 2 Mass erase
	START      CR = 0x01 << 16 //+ Start
	OPTSTRT    CR = 0x01 << 17 //+ Options modification start
	FSTPG      CR = 0x01 << 18 //+ Fast programming
	EOPIE      CR = 0x01 << 24 //+ End of operation interrupt enable
	ERRIE      CR = 0x01 << 25 //+ Error interrupt enable
	RDERRIE    CR = 0x01 << 26 //+ PCROP read error interrupt enable
	OBL_LAUNCH CR = 0x01 << 27 //+ Force the option byte loading
	OPTLOCK    CR = 0x01 << 30 //+ Options Lock
	LOCK       CR = 0x01 << 31 //+ FLASH_CR Lock
)

const (
	PGn         = 0
	PERn        = 1
	MER1n       = 2
	PNBn        = 3
	BKERn       = 11
	MER2n       = 15
	STARTn      = 16
	OPTSTRTn    = 17
	FSTPGn      = 18
	EOPIEn      = 24
	ERRIEn      = 25
	RDERRIEn    = 26
	OBL_LAUNCHn = 27
	OPTLOCKn    = 30
	LOCKn       = 31
)

const (
	ADDR_ECC ECCR = 0x7FFFF << 0 //+ ECC fail address
	BK_ECC   ECCR = 0x01 << 19   //+ ECC fail bank
	SYSF_ECC ECCR = 0x01 << 20   //+ System Flash ECC fail
	ECCIE    ECCR = 0x01 << 24   //+ ECC correction interrupt enable
	ECCC     ECCR = 0x01 << 30   //+ ECC correction
	ECCD     ECCR = 0x01 << 31   //+ ECC detection
)

const (
	ADDR_ECCn = 0
	BK_ECCn   = 19
	SYSF_ECCn = 20
	ECCIEn    = 24
	ECCCn     = 30
	ECCDn     = 31
)

const (
	RDP        OPTR = 0xFF << 0  //+ Read protection level
	BOR_LEV    OPTR = 0x07 << 8  //+ BOR reset Level
	nRST_STOP  OPTR = 0x01 << 12 //+ nRST_STOP
	nRST_STDBY OPTR = 0x01 << 13 //+ nRST_STDBY
	IDWG_SW    OPTR = 0x01 << 16 //+ Independent watchdog selection
	IWDG_STOP  OPTR = 0x01 << 17 //+ Independent watchdog counter freeze in Stop mode
	IWDG_STDBY OPTR = 0x01 << 18 //+ Independent watchdog counter freeze in Standby mode
	WWDG_SW    OPTR = 0x01 << 19 //+ Window watchdog selection
	BFB2       OPTR = 0x01 << 20 //+ Dual-bank boot
	DUALBANK   OPTR = 0x01 << 21 //+ Dual-Bank on 512 KB or 256 KB Flash memory devices
	nBOOT1     OPTR = 0x01 << 23 //+ Boot configuration
	SRAM2_PE   OPTR = 0x01 << 24 //+ SRAM2 parity check enable
	SRAM2_RST  OPTR = 0x01 << 25 //+ SRAM2 Erase when system reset
)

const (
	RDPn        = 0
	BOR_LEVn    = 8
	nRST_STOPn  = 12
	nRST_STDBYn = 13
	IDWG_SWn    = 16
	IWDG_STOPn  = 17
	IWDG_STDBYn = 18
	WWDG_SWn    = 19
	BFB2n       = 20
	DUALBANKn   = 21
	nBOOT1n     = 23
	SRAM2_PEn   = 24
	SRAM2_RSTn  = 25
)

const (
	PCROP1_STRT PCROP1SR = 0xFFFF << 0 //+ Bank 1 PCROP area start offset
)

const (
	PCROP1_STRTn = 0
)

const (
	PCROP1_END PCROP1ER = 0xFFFF << 0 //+ Bank 1 PCROP area end offset
	PCROP_RDP  PCROP1ER = 0x01 << 31  //+ PCROP area preserved when RDP level decreased
)

const (
	PCROP1_ENDn = 0
	PCROP_RDPn  = 31
)

const (
	WRP1A_STRT WRP1AR = 0xFF << 0  //+ Bank 1 WRP first area start offset
	WRP1A_END  WRP1AR = 0xFF << 16 //+ Bank 1 WRP first area A end offset
)

const (
	WRP1A_STRTn = 0
	WRP1A_ENDn  = 16
)

const (
	WRP1B_END  WRP1BR = 0xFF << 0  //+ Bank 1 WRP second area B start offset
	WRP1B_STRT WRP1BR = 0xFF << 16 //+ Bank 1 WRP second area B end offset
)

const (
	WRP1B_ENDn  = 0
	WRP1B_STRTn = 16
)

const (
	PCROP2_STRT PCROP2SR = 0xFFFF << 0 //+ Bank 2 PCROP area start offset
)

const (
	PCROP2_STRTn = 0
)

const (
	PCROP2_END PCROP2ER = 0xFFFF << 0 //+ Bank 2 PCROP area end offset
)

const (
	PCROP2_ENDn = 0
)

const (
	WRP2A_STRT WRP2AR = 0xFF << 0  //+ Bank 2 WRP first area A start offset
	WRP2A_END  WRP2AR = 0xFF << 16 //+ Bank 2 WRP first area A end offset
)

const (
	WRP2A_STRTn = 0
	WRP2A_ENDn  = 16
)

const (
	WRP2B_STRT WRP2BR = 0xFF << 0  //+ Bank 2 WRP second area B start offset
	WRP2B_END  WRP2BR = 0xFF << 16 //+ Bank 2 WRP second area B end offset
)

const (
	WRP2B_STRTn = 0
	WRP2B_ENDn  = 16
)
