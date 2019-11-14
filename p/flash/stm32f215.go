// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f215

// Package flash provides access to the registers of the FLASH peripheral.
//
// Instances:
//  FLASH  FLASH_BASE  -  FLASH  FLASH
// Registers:
//  0x000 32  ACR      Flash access control register
//  0x004 32  KEYR     Flash key register
//  0x008 32  OPTKEYR  Flash option key register
//  0x00C 32  SR       Status register
//  0x010 32  CR       Control register
//  0x014 32  OPTCR    Flash option control register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package flash

const (
	LATENCY ACR = 0x07 << 0  //+ Latency
	PRFTEN  ACR = 0x01 << 8  //+ Prefetch enable
	ICEN    ACR = 0x01 << 9  //+ Instruction cache enable
	DCEN    ACR = 0x01 << 10 //+ Data cache enable
	ICRST   ACR = 0x01 << 11 //+ Instruction cache reset
	DCRST   ACR = 0x01 << 12 //+ Data cache reset
)

const (
	LATENCYn = 0
	PRFTENn  = 8
	ICENn    = 9
	DCENn    = 10
	ICRSTn   = 11
	DCRSTn   = 12
)

const (
	EOP    SR = 0x01 << 0  //+ End of operation
	OPERR  SR = 0x01 << 1  //+ Operation error
	WRPERR SR = 0x01 << 4  //+ Write protection error
	PGAERR SR = 0x01 << 5  //+ Programming alignment error
	PGPERR SR = 0x01 << 6  //+ Programming parallelism error
	PGSERR SR = 0x01 << 7  //+ Programming sequence error
	BSY    SR = 0x01 << 16 //+ Busy
)

const (
	EOPn    = 0
	OPERRn  = 1
	WRPERRn = 4
	PGAERRn = 5
	PGPERRn = 6
	PGSERRn = 7
	BSYn    = 16
)

const (
	PG    CR = 0x01 << 0  //+ Programming
	SER   CR = 0x01 << 1  //+ Sector Erase
	MER   CR = 0x01 << 2  //+ Mass Erase
	SNB   CR = 0x0F << 3  //+ Sector number
	PSIZE CR = 0x03 << 8  //+ Program size
	STRT  CR = 0x01 << 16 //+ Start
	EOPIE CR = 0x01 << 24 //+ End of operation interrupt enable
	ERRIE CR = 0x01 << 25 //+ Error interrupt enable
	LOCK  CR = 0x01 << 31 //+ Lock
)

const (
	PGn    = 0
	SERn   = 1
	MERn   = 2
	SNBn   = 3
	PSIZEn = 8
	STRTn  = 16
	EOPIEn = 24
	ERRIEn = 25
	LOCKn  = 31
)

const (
	OPTLOCK    OPTCR = 0x01 << 0   //+ Option lock
	OPTSTRT    OPTCR = 0x01 << 1   //+ Option start
	BOR_LEV    OPTCR = 0x03 << 2   //+ BOR reset Level
	WDG_SW     OPTCR = 0x01 << 5   //+ WDG_SW User option bytes
	nRST_STOP  OPTCR = 0x01 << 6   //+ nRST_STOP User option bytes
	nRST_STDBY OPTCR = 0x01 << 7   //+ nRST_STDBY User option bytes
	RDP        OPTCR = 0xFF << 8   //+ Read protect
	nWRP       OPTCR = 0xFFF << 16 //+ Not write protect
)

const (
	OPTLOCKn    = 0
	OPTSTRTn    = 1
	BOR_LEVn    = 2
	WDG_SWn     = 5
	nRST_STOPn  = 6
	nRST_STDBYn = 7
	RDPn        = 8
	nWRPn       = 16
)
