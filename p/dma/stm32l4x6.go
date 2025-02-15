// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32l4x6

// Package dma provides access to the registers of the DMA peripheral.
//
// Instances:
//
//	DMA1  DMA1_BASE  AHB1  DMA1_CH1,DMA1_CH2,DMA1_CH3,DMA1_CH4,DMA1_CH5,DMA1_CH6,DMA1_CH7  Direct memory access controller
//	DMA2  DMA2_BASE  AHB1  DMA2_CH1,DMA2_CH2,DMA2_CH3,DMA2_CH4,DMA2_CH5,DMA2_CH6,DMA2_CH7  Direct memory access controller
//
// Registers:
//
//	0x000 32  ISR                      interrupt status register
//	0x004 32  IFCR                     interrupt flag clear register
//	0x008 32  C{CR,NDTR,PAR,MAR,_}[7]  channel configuration and controll registers
//	0x0A8 32  CSELR                    channel selection register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package dma

const (
	GIF1  ISR = 0x01 << 0  //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF1 ISR = 0x01 << 1  //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF1 ISR = 0x01 << 2  //+ Channel x half transfer flag (x = 1 ..7)
	TEIF1 ISR = 0x01 << 3  //+ Channel x transfer error flag (x = 1 ..7)
	GIF2  ISR = 0x01 << 4  //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF2 ISR = 0x01 << 5  //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF2 ISR = 0x01 << 6  //+ Channel x half transfer flag (x = 1 ..7)
	TEIF2 ISR = 0x01 << 7  //+ Channel x transfer error flag (x = 1 ..7)
	GIF3  ISR = 0x01 << 8  //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF3 ISR = 0x01 << 9  //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF3 ISR = 0x01 << 10 //+ Channel x half transfer flag (x = 1 ..7)
	TEIF3 ISR = 0x01 << 11 //+ Channel x transfer error flag (x = 1 ..7)
	GIF4  ISR = 0x01 << 12 //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF4 ISR = 0x01 << 13 //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF4 ISR = 0x01 << 14 //+ Channel x half transfer flag (x = 1 ..7)
	TEIF4 ISR = 0x01 << 15 //+ Channel x transfer error flag (x = 1 ..7)
	GIF5  ISR = 0x01 << 16 //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF5 ISR = 0x01 << 17 //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF5 ISR = 0x01 << 18 //+ Channel x half transfer flag (x = 1 ..7)
	TEIF5 ISR = 0x01 << 19 //+ Channel x transfer error flag (x = 1 ..7)
	GIF6  ISR = 0x01 << 20 //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF6 ISR = 0x01 << 21 //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF6 ISR = 0x01 << 22 //+ Channel x half transfer flag (x = 1 ..7)
	TEIF6 ISR = 0x01 << 23 //+ Channel x transfer error flag (x = 1 ..7)
	GIF7  ISR = 0x01 << 24 //+ Channel x global interrupt flag (x = 1 ..7)
	TCIF7 ISR = 0x01 << 25 //+ Channel x transfer complete flag (x = 1 ..7)
	HTIF7 ISR = 0x01 << 26 //+ Channel x half transfer flag (x = 1 ..7)
	TEIF7 ISR = 0x01 << 27 //+ Channel x transfer error flag (x = 1 ..7)
)

const (
	GIF1n  = 0
	TCIF1n = 1
	HTIF1n = 2
	TEIF1n = 3
	GIF2n  = 4
	TCIF2n = 5
	HTIF2n = 6
	TEIF2n = 7
	GIF3n  = 8
	TCIF3n = 9
	HTIF3n = 10
	TEIF3n = 11
	GIF4n  = 12
	TCIF4n = 13
	HTIF4n = 14
	TEIF4n = 15
	GIF5n  = 16
	TCIF5n = 17
	HTIF5n = 18
	TEIF5n = 19
	GIF6n  = 20
	TCIF6n = 21
	HTIF6n = 22
	TEIF6n = 23
	GIF7n  = 24
	TCIF7n = 25
	HTIF7n = 26
	TEIF7n = 27
)

const (
	CGIF1  IFCR = 0x01 << 0  //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF1 IFCR = 0x01 << 1  //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF1 IFCR = 0x01 << 2  //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF1 IFCR = 0x01 << 3  //+ Channel x transfer error clear (x = 1 ..7)
	CGIF2  IFCR = 0x01 << 4  //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF2 IFCR = 0x01 << 5  //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF2 IFCR = 0x01 << 6  //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF2 IFCR = 0x01 << 7  //+ Channel x transfer error clear (x = 1 ..7)
	CGIF3  IFCR = 0x01 << 8  //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF3 IFCR = 0x01 << 9  //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF3 IFCR = 0x01 << 10 //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF3 IFCR = 0x01 << 11 //+ Channel x transfer error clear (x = 1 ..7)
	CGIF4  IFCR = 0x01 << 12 //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF4 IFCR = 0x01 << 13 //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF4 IFCR = 0x01 << 14 //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF4 IFCR = 0x01 << 15 //+ Channel x transfer error clear (x = 1 ..7)
	CGIF5  IFCR = 0x01 << 16 //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF5 IFCR = 0x01 << 17 //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF5 IFCR = 0x01 << 18 //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF5 IFCR = 0x01 << 19 //+ Channel x transfer error clear (x = 1 ..7)
	CGIF6  IFCR = 0x01 << 20 //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF6 IFCR = 0x01 << 21 //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF6 IFCR = 0x01 << 22 //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF6 IFCR = 0x01 << 23 //+ Channel x transfer error clear (x = 1 ..7)
	CGIF7  IFCR = 0x01 << 24 //+ Channel x global interrupt clear (x = 1 ..7)
	CTCIF7 IFCR = 0x01 << 25 //+ Channel x transfer complete clear (x = 1 ..7)
	CHTIF7 IFCR = 0x01 << 26 //+ Channel x half transfer clear (x = 1 ..7)
	CTEIF7 IFCR = 0x01 << 27 //+ Channel x transfer error clear (x = 1 ..7)
)

const (
	CGIF1n  = 0
	CTCIF1n = 1
	CHTIF1n = 2
	CTEIF1n = 3
	CGIF2n  = 4
	CTCIF2n = 5
	CHTIF2n = 6
	CTEIF2n = 7
	CGIF3n  = 8
	CTCIF3n = 9
	CHTIF3n = 10
	CTEIF3n = 11
	CGIF4n  = 12
	CTCIF4n = 13
	CHTIF4n = 14
	CTEIF4n = 15
	CGIF5n  = 16
	CTCIF5n = 17
	CHTIF5n = 18
	CTEIF5n = 19
	CGIF6n  = 20
	CTCIF6n = 21
	CHTIF6n = 22
	CTEIF6n = 23
	CGIF7n  = 24
	CTCIF7n = 25
	CHTIF7n = 26
	CTEIF7n = 27
)

const (
	EN      CR = 0x01 << 0  //+ Channel enable
	TCIE    CR = 0x01 << 1  //+ Transfer complete interrupt enable
	HTIE    CR = 0x01 << 2  //+ Half transfer interrupt enable
	TEIE    CR = 0x01 << 3  //+ Transfer error interrupt enable
	DIR     CR = 0x01 << 4  //+ Data transfer direction
	CIRC    CR = 0x01 << 5  //+ Circular mode
	PINC    CR = 0x01 << 6  //+ Peripheral increment mode
	MINC    CR = 0x01 << 7  //+ Memory increment mode
	PSIZE   CR = 0x03 << 8  //+ Peripheral size
	MSIZE   CR = 0x03 << 10 //+ Memory size
	PL      CR = 0x03 << 12 //+ Channel priority level
	MEM2MEM CR = 0x01 << 14 //+ Memory to memory mode
)

const (
	ENn      = 0
	TCIEn    = 1
	HTIEn    = 2
	TEIEn    = 3
	DIRn     = 4
	CIRCn    = 5
	PINCn    = 6
	MINCn    = 7
	PSIZEn   = 8
	MSIZEn   = 10
	PLn      = 12
	MEM2MEMn = 14
)

const (
	NDT NDTR = 0xFFFF << 0 //+ Number of data to transfer
)

const (
	NDTn = 0
)

const (
	PA PAR = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	MA MAR = 0xFFFFFFFF << 0 //+ Memory address
)

const (
	MAn = 0
)

const (
	C1S CSELR = 0x0F << 0  //+ DMA channel 1 selection
	C2S CSELR = 0x0F << 4  //+ DMA channel 2 selection
	C3S CSELR = 0x0F << 8  //+ DMA channel 3 selection
	C4S CSELR = 0x0F << 12 //+ DMA channel 4 selection
	C5S CSELR = 0x0F << 16 //+ DMA channel 5 selection
	C6S CSELR = 0x0F << 20 //+ DMA channel 6 selection
	C7S CSELR = 0x0F << 24 //+ DMA channel 7 selection
)

const (
	C1Sn = 0
	C2Sn = 4
	C3Sn = 8
	C4Sn = 12
	C5Sn = 16
	C6Sn = 20
	C7Sn = 24
)
