// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32g471xx

// Package dma provides access to the registers of the DMA peripheral.
//
// Instances:
//
//	DMA1  DMA1_BASE  AHB1  DMA1_CH1,DMA1_CH2,DMA1_CH3,DMA1_CH4,DMA1_CH5,DMA1_CH6,DMA1_CH8                    DMA controller
//	DMA2  DMA2_BASE  AHB1  DMA1_CH7,DMA2_CH1,DMA2_CH2,DMA2_CH3,DMA2_CH4,DMA2_CH5,DMA2_CH6,DMA2_CH7,DMA2_CH8  DMA controller
//
// Registers:
//
//	0x000 32  ISR                      interrupt status register
//	0x004 32  IFCR                     DMA interrupt flag clear register
//	0x008 32  C{CR,NDTR,PAR,MAR,_}[8]  channel configuration and controll registers
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package dma

const (
	GIF1  ISR = 0x01 << 0  //+ GIF1
	TCIF1 ISR = 0x01 << 1  //+ TCIF1
	HTIF1 ISR = 0x01 << 2  //+ HTIF1
	TEIF1 ISR = 0x01 << 3  //+ TEIF1
	GIF2  ISR = 0x01 << 4  //+ GIF2
	TCIF2 ISR = 0x01 << 5  //+ TCIF2
	HTIF2 ISR = 0x01 << 6  //+ HTIF2
	TEIF2 ISR = 0x01 << 7  //+ TEIF2
	GIF3  ISR = 0x01 << 8  //+ GIF3
	TCIF3 ISR = 0x01 << 9  //+ TCIF3
	HTIF3 ISR = 0x01 << 10 //+ HTIF3
	TEIF3 ISR = 0x01 << 11 //+ TEIF3
	GIF4  ISR = 0x01 << 12 //+ GIF4
	TCIF4 ISR = 0x01 << 13 //+ TCIF4
	HTIF4 ISR = 0x01 << 14 //+ HTIF4
	TEIF4 ISR = 0x01 << 15 //+ TEIF4
	GIF5  ISR = 0x01 << 16 //+ GIF5
	TCIF5 ISR = 0x01 << 17 //+ TCIF5
	HTIF5 ISR = 0x01 << 18 //+ HTIF5
	TEIF5 ISR = 0x01 << 19 //+ TEIF5
	GIF6  ISR = 0x01 << 20 //+ GIF6
	TCIF6 ISR = 0x01 << 21 //+ TCIF6
	HTIF6 ISR = 0x01 << 22 //+ HTIF6
	TEIF6 ISR = 0x01 << 23 //+ TEIF6
	GIF7  ISR = 0x01 << 24 //+ GIF7
	TCIF7 ISR = 0x01 << 25 //+ TCIF7
	HTIF7 ISR = 0x01 << 26 //+ HTIF7
	TEIF7 ISR = 0x01 << 27 //+ TEIF7
	GIF8  ISR = 0x01 << 28 //+ GIF8
	TCIF8 ISR = 0x01 << 29 //+ TCIF8
	HTIF8 ISR = 0x01 << 30 //+ HTIF8
	TEIF8 ISR = 0x01 << 31 //+ TEIF8
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
	GIF8n  = 28
	TCIF8n = 29
	HTIF8n = 30
	TEIF8n = 31
)

const (
	GIF1  IFCR = 0x01 << 0  //+ GIF1
	TCIF1 IFCR = 0x01 << 1  //+ TCIF1
	HTIF1 IFCR = 0x01 << 2  //+ HTIF1
	TEIF1 IFCR = 0x01 << 3  //+ TEIF1
	GIF2  IFCR = 0x01 << 4  //+ GIF2
	TCIF2 IFCR = 0x01 << 5  //+ TCIF2
	HTIF2 IFCR = 0x01 << 6  //+ HTIF2
	TEIF2 IFCR = 0x01 << 7  //+ TEIF2
	GIF3  IFCR = 0x01 << 8  //+ GIF3
	TCIF3 IFCR = 0x01 << 9  //+ TCIF3
	HTIF3 IFCR = 0x01 << 10 //+ HTIF3
	TEIF3 IFCR = 0x01 << 11 //+ TEIF3
	GIF4  IFCR = 0x01 << 12 //+ GIF4
	TCIF4 IFCR = 0x01 << 13 //+ TCIF4
	HTIF4 IFCR = 0x01 << 14 //+ HTIF4
	TEIF4 IFCR = 0x01 << 15 //+ TEIF4
	GIF5  IFCR = 0x01 << 16 //+ GIF5
	TCIF5 IFCR = 0x01 << 17 //+ TCIF5
	HTIF5 IFCR = 0x01 << 18 //+ HTIF5
	TEIF5 IFCR = 0x01 << 19 //+ TEIF5
	GIF6  IFCR = 0x01 << 20 //+ GIF6
	TCIF6 IFCR = 0x01 << 21 //+ TCIF6
	HTIF6 IFCR = 0x01 << 22 //+ HTIF6
	TEIF6 IFCR = 0x01 << 23 //+ TEIF6
	GIF7  IFCR = 0x01 << 24 //+ GIF7
	TCIF7 IFCR = 0x01 << 25 //+ TCIF7
	HTIF7 IFCR = 0x01 << 26 //+ HTIF7
	TEIF7 IFCR = 0x01 << 27 //+ TEIF7
	GIF8  IFCR = 0x01 << 28 //+ GIF8
	TCIF8 IFCR = 0x01 << 29 //+ TCIF8
	HTIF8 IFCR = 0x01 << 30 //+ HTIF8
	TEIF8 IFCR = 0x01 << 31 //+ TEIF8
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
	GIF8n  = 28
	TCIF8n = 29
	HTIF8n = 30
	TEIF8n = 31
)

const (
	EN      CR = 0x01 << 0  //+ channel enable
	TCIE    CR = 0x01 << 1  //+ TCIE
	HTIE    CR = 0x01 << 2  //+ HTIE
	TEIE    CR = 0x01 << 3  //+ TEIE
	DIR     CR = 0x01 << 4  //+ DIR
	CIRC    CR = 0x01 << 5  //+ CIRC
	PINC    CR = 0x01 << 6  //+ PINC
	MINC    CR = 0x01 << 7  //+ MINC
	PSIZE   CR = 0x03 << 8  //+ PSIZE
	MSIZE   CR = 0x03 << 10 //+ MSIZE
	PL      CR = 0x03 << 12 //+ PL
	MEM2MEM CR = 0x01 << 14 //+ MEM2MEM
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
	NDT NDTR = 0xFFFF << 0 //+ Number of data items to transfer
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
	MA MAR = 0xFFFFFFFF << 0 //+ Memory 1 address (used in case of Double buffer mode)
)

const (
	MAn = 0
)
