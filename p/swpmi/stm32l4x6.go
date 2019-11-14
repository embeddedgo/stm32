// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32l4x6

// Package swpmi provides access to the registers of the SWPMI1 peripheral.
//
// Instances:
//  SWPMI1  SWPMI1_BASE  APB1  SWPMI1  Single Wire Protocol Master Interface
// Registers:
//  0x000 32  CR   SWPMI Configuration/Control register
//  0x004 32  BRR  SWPMI Bitrate register
//  0x00C 32  ISR  SWPMI Interrupt and Status register
//  0x010 32  ICR  SWPMI Interrupt Flag Clear register
//  0x014 32  IER  SWPMI Interrupt Enable register
//  0x018 32  RFL  SWPMI Receive Frame Length register
//  0x01C 32  TDR  SWPMI Transmit data register
//  0x020 32  RDR  SWPMI Receive data register
//  0x024 32  OR   SWPMI Option register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package swpmi

const (
	RXDMA  CR = 0x01 << 0  //+ Reception DMA enable
	TXDMA  CR = 0x01 << 1  //+ Transmission DMA enable
	RXMODE CR = 0x01 << 2  //+ Reception buffering mode
	TXMODE CR = 0x01 << 3  //+ Transmission buffering mode
	LPBK   CR = 0x01 << 4  //+ Loopback mode enable
	SWPACT CR = 0x01 << 5  //+ Single wire protocol master interface activate
	DEACT  CR = 0x01 << 10 //+ Single wire protocol master interface deactivate
)

const (
	RXDMAn  = 0
	TXDMAn  = 1
	RXMODEn = 2
	TXMODEn = 3
	LPBKn   = 4
	SWPACTn = 5
	DEACTn  = 10
)

const (
	BR BRR = 0x3F << 0 //+ Bitrate prescaler
)

const (
	BRn = 0
)

const (
	RXBFF  ISR = 0x01 << 0  //+ Receive buffer full flag
	TXBEF  ISR = 0x01 << 1  //+ Transmit buffer empty flag
	RXBERF ISR = 0x01 << 2  //+ Receive CRC error flag
	RXOVRF ISR = 0x01 << 3  //+ Receive overrun error flag
	TXUNRF ISR = 0x01 << 4  //+ Transmit underrun error flag
	RXNE   ISR = 0x01 << 5  //+ Receive data register not empty
	TXE    ISR = 0x01 << 6  //+ Transmit data register empty
	TCF    ISR = 0x01 << 7  //+ Transfer complete flag
	SRF    ISR = 0x01 << 8  //+ Slave resume flag
	SUSP   ISR = 0x01 << 9  //+ SUSPEND flag
	DEACTF ISR = 0x01 << 10 //+ DEACTIVATED flag
)

const (
	RXBFFn  = 0
	TXBEFn  = 1
	RXBERFn = 2
	RXOVRFn = 3
	TXUNRFn = 4
	RXNEn   = 5
	TXEn    = 6
	TCFn    = 7
	SRFn    = 8
	SUSPn   = 9
	DEACTFn = 10
)

const (
	CRXBFF  ICR = 0x01 << 0 //+ Clear receive buffer full flag
	CTXBEF  ICR = 0x01 << 1 //+ Clear transmit buffer empty flag
	CRXBERF ICR = 0x01 << 2 //+ Clear receive CRC error flag
	CRXOVRF ICR = 0x01 << 3 //+ Clear receive overrun error flag
	CTXUNRF ICR = 0x01 << 4 //+ Clear transmit underrun error flag
	CTCF    ICR = 0x01 << 7 //+ Clear transfer complete flag
	CSRF    ICR = 0x01 << 8 //+ Clear slave resume flag
)

const (
	CRXBFFn  = 0
	CTXBEFn  = 1
	CRXBERFn = 2
	CRXOVRFn = 3
	CTXUNRFn = 4
	CTCFn    = 7
	CSRFn    = 8
)

const (
	RXBFIE  IER = 0x01 << 0 //+ Receive buffer full interrupt enable
	TXBEIE  IER = 0x01 << 1 //+ Transmit buffer empty interrupt enable
	RXBERIE IER = 0x01 << 2 //+ Receive CRC error interrupt enable
	RXOVRIE IER = 0x01 << 3 //+ Receive overrun error interrupt enable
	TXUNRIE IER = 0x01 << 4 //+ Transmit underrun error interrupt enable
	RIE     IER = 0x01 << 5 //+ Receive interrupt enable
	TIE     IER = 0x01 << 6 //+ Transmit interrupt enable
	TCIE    IER = 0x01 << 7 //+ Transmit complete interrupt enable
	SRIE    IER = 0x01 << 8 //+ Slave resume interrupt enable
)

const (
	RXBFIEn  = 0
	TXBEIEn  = 1
	RXBERIEn = 2
	RXOVRIEn = 3
	TXUNRIEn = 4
	RIEn     = 5
	TIEn     = 6
	TCIEn    = 7
	SRIEn    = 8
)

const (
	RFL RFL = 0x1F << 0 //+ Receive frame length
)

const (
	RFLn = 0
)

const (
	TD TDR = 0xFFFFFFFF << 0 //+ Transmit data
)

const (
	TDn = 0
)

const (
	RD RDR = 0xFFFFFFFF << 0 //+ received data
)

const (
	RDn = 0
)

const (
	SWP_TBYP  OR = 0x01 << 0 //+ SWP transceiver bypass
	SWP_CLASS OR = 0x01 << 1 //+ SWP class selection
)

const (
	SWP_TBYPn  = 0
	SWP_CLASSn = 1
)
