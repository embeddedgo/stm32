// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package spih

type CR1 uint32

// CR1 bits
const (
	SPE      CR1 = 0x01 << 0  //+ Serial Peripheral Enable
	MASRX    CR1 = 0x01 << 8  //+ Master automatic SUSP in Receive mode
	CSTART   CR1 = 0x01 << 9  //+ Master transfer start
	CSUSP    CR1 = 0x01 << 10 //+ Master SUSPend request
	HDDIR    CR1 = 0x01 << 11 //+ Rx/Tx direction at Half-duplex mode
	SSI      CR1 = 0x01 << 12 //+ Internal SS signal input level
	CRC33_17 CR1 = 0x01 << 13 //+ 32-bit CRC polynomial configuration
	RCRCI    CR1 = 0x01 << 14 //+ CRC calculation initialization pattern control for receiver
	TCRCI    CR1 = 0x01 << 15 //+ CRC calculation initialization pattern control for transmitter
	IOLOCK   CR1 = 0x01 << 16 //+ Locking the AF configuration of associated IOs
)

const (
	SPEn      = 0
	MASRXn    = 8
	CSTARTn   = 9
	CSUSPn    = 10
	HDDIRn    = 11
	SSIn      = 12
	CRC33_17n = 13
	RCRCIn    = 14
	TCRCIn    = 15
	IOLOCKn   = 16
)

// CR2 bits
const (
	TSIZE uint32 = 0xFFFF << 0  //+ Number of data at current transfer
	TSER  uint32 = 0xFFFF << 16 //+ Number of data transfer extension to be reload into TSIZE just when a previous
)

const (
	TSIZEn = 0
	TSERn  = 16
)

type CFG1 uint32

// CFG1 bits
const (
	DSIZE   CFG1 = 0x1F << 0  //+ Number of bits in at single SPI data frame
	FTHVL   CFG1 = 0x0F << 5  //+ threshold level
	UDRCFG  CFG1 = 0x03 << 9  //+ Behavior of slave transmitter at underrun condition
	UDRDET  CFG1 = 0x03 << 11 //+ Detection of underrun condition at slave transmitter
	RXDMAEN CFG1 = 0x01 << 14 //+ Rx DMA stream enable
	TXDMAEN CFG1 = 0x01 << 15 //+ Tx DMA stream enable
	CRCSIZE CFG1 = 0x1F << 16 //+ Length of CRC frame to be transacted and compared
	CRCEN   CFG1 = 0x01 << 22 //+ Hardware CRC computation enable
	MBR     CFG1 = 0x07 << 28 //+ Master baud rate
)

const (
	DSIZEn   = 0
	FTHVLn   = 5
	UDRCFGn  = 9
	UDRDETn  = 11
	RXDMAENn = 14
	TXDMAENn = 15
	CRCSIZEn = 16
	CRCENn   = 22
	MBRn     = 28
)

type CFG2 uint32

// CFG2 bits
const (
	MSSI    CFG2 = 0x0F << 0  //+ Master SS Idleness
	MIDI    CFG2 = 0x0F << 4  //+ Master Inter-Data Idleness
	IOSWP   CFG2 = 0x01 << 15 //+ Swap functionality of MISO and MOSI pins
	COMM    CFG2 = 0x03 << 17 //+ SPI Communication Mode
	SP      CFG2 = 0x07 << 19 //+ Serial Protocol
	MASTER  CFG2 = 0x01 << 22 //+ SPI Master
	LSBFRST CFG2 = 0x01 << 23 //+ Data frame format
	CPHA    CFG2 = 0x01 << 24 //+ Clock phase
	CPOL    CFG2 = 0x01 << 25 //+ Clock polarity
	SSM     CFG2 = 0x01 << 26 //+ Software management of SS signal input
	SSIOP   CFG2 = 0x01 << 28 //+ SS input/output polarity
	SSOE    CFG2 = 0x01 << 29 //+ SS output enable
	SSOM    CFG2 = 0x01 << 30 //+ SS output management in master mode
	AFCNTR  CFG2 = 0x01 << 31 //+ Alternate function GPIOs control
)

const (
	MSSIn    = 0
	MIDIn    = 4
	IOSWPn   = 15
	COMMn    = 17
	SPn      = 19
	MASTERn  = 22
	LSBFRSTn = 23
	CPHAn    = 24
	CPOLn    = 25
	SSMn     = 26
	SSIOPn   = 28
	SSOEn    = 29
	SSOMn    = 30
	AFCNTRn  = 31
)

type SR uint32

// SR bits
const (
	RXP    SR = 0x01 << 0    //+ Rx-Packet available
	TXP    SR = 0x01 << 1    //+ Tx-Packet space available
	DXP    SR = 0x01 << 2    //+ Duplex Packet
	EOT    SR = 0x01 << 3    //+ End Of Transfer
	TXTF   SR = 0x01 << 4    //+ Transmission Transfer Filled
	UDR    SR = 0x01 << 5    //+ Underrun at slave transmission mode
	OVR    SR = 0x01 << 6    //+ Overrun
	CRCE   SR = 0x01 << 7    //+ CRC Error
	TIFRE  SR = 0x01 << 8    //+ TI frame format error
	MODF   SR = 0x01 << 9    //+ Mode Fault
	TSERF  SR = 0x01 << 10   //+ Additional number of SPI data to be transacted was reload
	SUSP   SR = 0x01 << 11   //+ SUSPend
	TXC    SR = 0x01 << 12   //+ TxFIFO transmission complete
	RXPLVL SR = 0x03 << 13   //+ RxFIFO Packing LeVeL
	RXWNE  SR = 0x01 << 15   //+ RxFIFO Word Not Empty
	CTSIZE SR = 0xFFFF << 16 //+ Number of data frames remaining in current TSIZE session
)

const (
	RXPn    = 0
	TXPn    = 1
	DXPn    = 2
	EOTn    = 3
	TXTFn   = 4
	UDRn    = 5
	OVRn    = 6
	CRCEn   = 7
	TIFREn  = 8
	MODFn   = 9
	TSERFn  = 10
	SUSPn   = 11
	TXCn    = 12
	RXPLVLn = 13
	RXWNEn  = 15
	CTSIZEn = 16
)

type CGFR uint32

// CGFR bits
const (
	I2SMOD  CGFR = 0x01 << 0  //+ I2S mode selection
	I2SCFG  CGFR = 0x07 << 1  //+ I2S configuration mode
	I2SSTD  CGFR = 0x03 << 4  //+ I2S standard selection
	PCMSYNC CGFR = 0x01 << 7  //+ PCM frame synchronization
	DATLEN  CGFR = 0x03 << 8  //+ Data length to be transferred
	CHLEN   CGFR = 0x01 << 10 //+ Channel length (number of bits per audio channel)
	CKPOL   CGFR = 0x01 << 11 //+ Serial audio clock polarity
	FIXCH   CGFR = 0x01 << 12 //+ Word select inversion
	WSINV   CGFR = 0x01 << 13 //+ Fixed channel length in SLAVE
	DATFMT  CGFR = 0x01 << 14 //+ Data format
	I2SDIV  CGFR = 0xFF << 16 //+ I2S linear prescaler
	ODD     CGFR = 0x01 << 24 //+ Odd factor for the prescaler
	MCKOE   CGFR = 0x01 << 25 //+ Master clock output enable
)

const (
	I2SMODn  = 0
	I2SCFGn  = 1
	I2SSTDn  = 4
	PCMSYNCn = 7
	DATLENn  = 8
	CHLENn   = 10
	CKPOLn   = 11
	FIXCHn   = 12
	WSINVn   = 13
	DATFMTn  = 14
	I2SDIVn  = 16
	ODDn     = 24
	MCKOEn   = 25
)
