// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32g471xx

// Package dma provides access to the registers of the DMA1 peripheral.
//
// Instances:
//  DMA1  DMA1_BASE  AHB1  DMA1_CH1,DMA1_CH2,DMA1_CH3,DMA1_CH4,DMA1_CH5,DMA1_CH6,DMA1_CH8                    DMA controller
//  DMA2  DMA2_BASE  AHB1  DMA1_CH7,DMA2_CH1,DMA2_CH2,DMA2_CH3,DMA2_CH4,DMA2_CH5,DMA2_CH6,DMA2_CH7,DMA2_CH8  DMA controller
// Registers:
//  0x000 32  ISR     interrupt status register
//  0x004 32  IFCR    DMA interrupt flag clear register
//  0x008 32  CCR1    DMA channel 1 configuration register
//  0x01C 32  CCR2    DMA channel 2 configuration register
//  0x030 32  CCR3    DMA channel 3 configuration register
//  0x044 32  CCR4    DMA channel 3 configuration register
//  0x058 32  CCR5    DMA channel 4 configuration register
//  0x06C 32  CCR6    DMA channel 5 configuration register
//  0x080 32  CCR7    DMA channel 6 configuration register
//  0x094 32  CCR8    DMA channel 7 configuration register
//  0x00C 32  CNDTR1  channel x number of data to transfer register
//  0x020 32  CNDTR2  channel x number of data to transfer register
//  0x034 32  CNDTR3  channel x number of data to transfer register
//  0x048 32  CNDTR4  channel x number of data to transfer register
//  0x05C 32  CNDTR5  channel x number of data to transfer register
//  0x070 32  CNDTR6  channel x number of data to transfer register
//  0x084 32  CNDTR7  channel x number of data to transfer register
//  0x098 32  CNDTR8  channel x number of data to transfer register
//  0x010 32  CPAR1   DMA channel x peripheral address register
//  0x024 32  CPAR2   DMA channel x peripheral address register
//  0x038 32  CPAR3   DMA channel x peripheral address register
//  0x04C 32  CPAR4   DMA channel x peripheral address register
//  0x060 32  CPAR5   DMA channel x peripheral address register
//  0x074 32  CPAR6   DMA channel x peripheral address register
//  0x088 32  CPAR7   DMA channel x peripheral address register
//  0x09C 32  CPAR8   DMA channel x peripheral address register
//  0x014 32  CMAR1   DMA channel x memory address register
//  0x028 32  CMAR2   DMA channel x memory address register
//  0x03C 32  CMAR3   DMA channel x memory address register
//  0x050 32  CMAR4   DMA channel x memory address register
//  0x064 32  CMAR5   DMA channel x memory address register
//  0x078 32  CMAR6   DMA channel x memory address register
//  0x08C 32  CMAR7   DMA channel x memory address register
//  0x0A0 32  CMAR8   DMA channel x memory address register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
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
	EN      CCR1 = 0x01 << 0  //+ channel enable
	TCIE    CCR1 = 0x01 << 1  //+ TCIE
	HTIE    CCR1 = 0x01 << 2  //+ HTIE
	TEIE    CCR1 = 0x01 << 3  //+ TEIE
	DIR     CCR1 = 0x01 << 4  //+ DIR
	CIRC    CCR1 = 0x01 << 5  //+ CIRC
	PINC    CCR1 = 0x01 << 6  //+ PINC
	MINC    CCR1 = 0x01 << 7  //+ MINC
	PSIZE   CCR1 = 0x03 << 8  //+ PSIZE
	MSIZE   CCR1 = 0x03 << 10 //+ MSIZE
	PL      CCR1 = 0x03 << 12 //+ PL
	MEM2MEM CCR1 = 0x01 << 14 //+ MEM2MEM
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
	EN      CCR2 = 0x01 << 0  //+ channel enable
	TCIE    CCR2 = 0x01 << 1  //+ TCIE
	HTIE    CCR2 = 0x01 << 2  //+ HTIE
	TEIE    CCR2 = 0x01 << 3  //+ TEIE
	DIR     CCR2 = 0x01 << 4  //+ DIR
	CIRC    CCR2 = 0x01 << 5  //+ CIRC
	PINC    CCR2 = 0x01 << 6  //+ PINC
	MINC    CCR2 = 0x01 << 7  //+ MINC
	PSIZE   CCR2 = 0x03 << 8  //+ PSIZE
	MSIZE   CCR2 = 0x03 << 10 //+ MSIZE
	PL      CCR2 = 0x03 << 12 //+ PL
	MEM2MEM CCR2 = 0x01 << 14 //+ MEM2MEM
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
	EN      CCR3 = 0x01 << 0  //+ channel enable
	TCIE    CCR3 = 0x01 << 1  //+ TCIE
	HTIE    CCR3 = 0x01 << 2  //+ HTIE
	TEIE    CCR3 = 0x01 << 3  //+ TEIE
	DIR     CCR3 = 0x01 << 4  //+ DIR
	CIRC    CCR3 = 0x01 << 5  //+ CIRC
	PINC    CCR3 = 0x01 << 6  //+ PINC
	MINC    CCR3 = 0x01 << 7  //+ MINC
	PSIZE   CCR3 = 0x03 << 8  //+ PSIZE
	MSIZE   CCR3 = 0x03 << 10 //+ MSIZE
	PL      CCR3 = 0x03 << 12 //+ PL
	MEM2MEM CCR3 = 0x01 << 14 //+ MEM2MEM
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
	EN      CCR4 = 0x01 << 0  //+ channel enable
	TCIE    CCR4 = 0x01 << 1  //+ TCIE
	HTIE    CCR4 = 0x01 << 2  //+ HTIE
	TEIE    CCR4 = 0x01 << 3  //+ TEIE
	DIR     CCR4 = 0x01 << 4  //+ DIR
	CIRC    CCR4 = 0x01 << 5  //+ CIRC
	PINC    CCR4 = 0x01 << 6  //+ PINC
	MINC    CCR4 = 0x01 << 7  //+ MINC
	PSIZE   CCR4 = 0x03 << 8  //+ PSIZE
	MSIZE   CCR4 = 0x03 << 10 //+ MSIZE
	PL      CCR4 = 0x03 << 12 //+ PL
	MEM2MEM CCR4 = 0x01 << 14 //+ MEM2MEM
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
	EN      CCR5 = 0x01 << 0  //+ channel enable
	TCIE    CCR5 = 0x01 << 1  //+ TCIE
	HTIE    CCR5 = 0x01 << 2  //+ HTIE
	TEIE    CCR5 = 0x01 << 3  //+ TEIE
	DIR     CCR5 = 0x01 << 4  //+ DIR
	CIRC    CCR5 = 0x01 << 5  //+ CIRC
	PINC    CCR5 = 0x01 << 6  //+ PINC
	MINC    CCR5 = 0x01 << 7  //+ MINC
	PSIZE   CCR5 = 0x03 << 8  //+ PSIZE
	MSIZE   CCR5 = 0x03 << 10 //+ MSIZE
	PL      CCR5 = 0x03 << 12 //+ PL
	MEM2MEM CCR5 = 0x01 << 14 //+ MEM2MEM
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
	EN      CCR6 = 0x01 << 0  //+ channel enable
	TCIE    CCR6 = 0x01 << 1  //+ TCIE
	HTIE    CCR6 = 0x01 << 2  //+ HTIE
	TEIE    CCR6 = 0x01 << 3  //+ TEIE
	DIR     CCR6 = 0x01 << 4  //+ DIR
	CIRC    CCR6 = 0x01 << 5  //+ CIRC
	PINC    CCR6 = 0x01 << 6  //+ PINC
	MINC    CCR6 = 0x01 << 7  //+ MINC
	PSIZE   CCR6 = 0x03 << 8  //+ PSIZE
	MSIZE   CCR6 = 0x03 << 10 //+ MSIZE
	PL      CCR6 = 0x03 << 12 //+ PL
	MEM2MEM CCR6 = 0x01 << 14 //+ MEM2MEM
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
	EN      CCR7 = 0x01 << 0  //+ channel enable
	TCIE    CCR7 = 0x01 << 1  //+ TCIE
	HTIE    CCR7 = 0x01 << 2  //+ HTIE
	TEIE    CCR7 = 0x01 << 3  //+ TEIE
	DIR     CCR7 = 0x01 << 4  //+ DIR
	CIRC    CCR7 = 0x01 << 5  //+ CIRC
	PINC    CCR7 = 0x01 << 6  //+ PINC
	MINC    CCR7 = 0x01 << 7  //+ MINC
	PSIZE   CCR7 = 0x03 << 8  //+ PSIZE
	MSIZE   CCR7 = 0x03 << 10 //+ MSIZE
	PL      CCR7 = 0x03 << 12 //+ PL
	MEM2MEM CCR7 = 0x01 << 14 //+ MEM2MEM
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
	EN      CCR8 = 0x01 << 0  //+ channel enable
	TCIE    CCR8 = 0x01 << 1  //+ TCIE
	HTIE    CCR8 = 0x01 << 2  //+ HTIE
	TEIE    CCR8 = 0x01 << 3  //+ TEIE
	DIR     CCR8 = 0x01 << 4  //+ DIR
	CIRC    CCR8 = 0x01 << 5  //+ CIRC
	PINC    CCR8 = 0x01 << 6  //+ PINC
	MINC    CCR8 = 0x01 << 7  //+ MINC
	PSIZE   CCR8 = 0x03 << 8  //+ PSIZE
	MSIZE   CCR8 = 0x03 << 10 //+ MSIZE
	PL      CCR8 = 0x03 << 12 //+ PL
	MEM2MEM CCR8 = 0x01 << 14 //+ MEM2MEM
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
	NDT CNDTR1 = 0xFFFF << 0 //+ Number of data items to transfer
)

const (
	NDTn = 0
)

const (
	NDT CNDTR2 = 0xFFFF << 0 //+ Number of data items to transfer
)

const (
	NDTn = 0
)

const (
	NDT CNDTR3 = 0xFFFF << 0 //+ Number of data items to transfer
)

const (
	NDTn = 0
)

const (
	NDT CNDTR4 = 0xFFFF << 0 //+ Number of data items to transfer
)

const (
	NDTn = 0
)

const (
	NDT CNDTR5 = 0xFFFF << 0 //+ Number of data items to transfer
)

const (
	NDTn = 0
)

const (
	NDT CNDTR6 = 0xFFFF << 0 //+ Number of data items to transfer
)

const (
	NDTn = 0
)

const (
	NDT CNDTR7 = 0xFFFF << 0 //+ Number of data items to transfer
)

const (
	NDTn = 0
)

const (
	NDT CNDTR8 = 0xFFFF << 0 //+ Number of data items to transfer
)

const (
	NDTn = 0
)

const (
	PA CPAR1 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	PA CPAR2 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	PA CPAR3 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	PA CPAR4 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	PA CPAR5 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	PA CPAR6 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	PA CPAR7 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	PA CPAR8 = 0xFFFFFFFF << 0 //+ Peripheral address
)

const (
	PAn = 0
)

const (
	MA CMAR1 = 0xFFFFFFFF << 0 //+ Memory 1 address (used in case of Double buffer mode)
)

const (
	MAn = 0
)

const (
	MA CMAR2 = 0xFFFFFFFF << 0 //+ Memory 1 address (used in case of Double buffer mode)
)

const (
	MAn = 0
)

const (
	MA CMAR3 = 0xFFFFFFFF << 0 //+ Memory 1 address (used in case of Double buffer mode)
)

const (
	MAn = 0
)

const (
	MA CMAR4 = 0xFFFFFFFF << 0 //+ Memory 1 address (used in case of Double buffer mode)
)

const (
	MAn = 0
)

const (
	MA CMAR5 = 0xFFFFFFFF << 0 //+ Memory 1 address (used in case of Double buffer mode)
)

const (
	MAn = 0
)

const (
	MA CMAR6 = 0xFFFFFFFF << 0 //+ Memory 1 address (used in case of Double buffer mode)
)

const (
	MAn = 0
)

const (
	MA CMAR7 = 0xFFFFFFFF << 0 //+ Memory 1 address (used in case of Double buffer mode)
)

const (
	MAn = 0
)

const (
	MA CMAR8 = 0xFFFFFFFF << 0 //+ Memory 1 address (used in case of Double buffer mode)
)

const (
	MAn = 0
)
