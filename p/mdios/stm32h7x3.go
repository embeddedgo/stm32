// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package mdios provides access to the registers of the MDIOS peripheral.
//
// Instances:
//
//	MDIOS  MDIOS_BASE  APB1  MDIOS_WKUP,MDIOS  Management data input/output slave
//
// Registers:
//
//	0x000 32  CR       MDIOS configuration register
//	0x004 32  WRFR     MDIOS write flag register
//	0x008 32  CWRFR    MDIOS clear write flag register
//	0x00C 32  RDFR     MDIOS read flag register
//	0x010 32  CRDFR    MDIOS clear read flag register
//	0x014 32  SR       MDIOS status register
//	0x018 32  CLRFR    MDIOS clear flag register
//	0x01C 32  DINR0    MDIOS input data register 0
//	0x020 32  DINR1    MDIOS input data register 1
//	0x024 32  DINR2    MDIOS input data register 2
//	0x028 32  DINR3    MDIOS input data register 3
//	0x02C 32  DINR4    MDIOS input data register 4
//	0x030 32  DINR5    MDIOS input data register 5
//	0x034 32  DINR6    MDIOS input data register 6
//	0x038 32  DINR7    MDIOS input data register 7
//	0x03C 32  DINR8    MDIOS input data register 8
//	0x040 32  DINR9    MDIOS input data register 9
//	0x044 32  DINR10   MDIOS input data register 10
//	0x048 32  DINR11   MDIOS input data register 11
//	0x04C 32  DINR12   MDIOS input data register 12
//	0x050 32  DINR13   MDIOS input data register 13
//	0x054 32  DINR14   MDIOS input data register 14
//	0x058 32  DINR15   MDIOS input data register 15
//	0x05C 32  DINR16   MDIOS input data register 16
//	0x060 32  DINR17   MDIOS input data register 17
//	0x064 32  DINR18   MDIOS input data register 18
//	0x068 32  DINR19   MDIOS input data register 19
//	0x06C 32  DINR20   MDIOS input data register 20
//	0x070 32  DINR21   MDIOS input data register 21
//	0x074 32  DINR22   MDIOS input data register 22
//	0x078 32  DINR23   MDIOS input data register 23
//	0x07C 32  DINR24   MDIOS input data register 24
//	0x080 32  DINR25   MDIOS input data register 25
//	0x084 32  DINR26   MDIOS input data register 26
//	0x088 32  DINR27   MDIOS input data register 27
//	0x08C 32  DINR28   MDIOS input data register 28
//	0x090 32  DINR29   MDIOS input data register 29
//	0x094 32  DINR30   MDIOS input data register 30
//	0x098 32  DINR31   MDIOS input data register 31
//	0x09C 32  DOUTR0   MDIOS output data register 0
//	0x0A0 32  DOUTR1   MDIOS output data register 1
//	0x0A4 32  DOUTR2   MDIOS output data register 2
//	0x0A8 32  DOUTR3   MDIOS output data register 3
//	0x0AC 32  DOUTR4   MDIOS output data register 4
//	0x0B0 32  DOUTR5   MDIOS output data register 5
//	0x0B4 32  DOUTR6   MDIOS output data register 6
//	0x0B8 32  DOUTR7   MDIOS output data register 7
//	0x0BC 32  DOUTR8   MDIOS output data register 8
//	0x0C0 32  DOUTR9   MDIOS output data register 9
//	0x0C4 32  DOUTR10  MDIOS output data register 10
//	0x0C8 32  DOUTR11  MDIOS output data register 11
//	0x0CC 32  DOUTR12  MDIOS output data register 12
//	0x0D0 32  DOUTR13  MDIOS output data register 13
//	0x0D4 32  DOUTR14  MDIOS output data register 14
//	0x0D8 32  DOUTR15  MDIOS output data register 15
//	0x0DC 32  DOUTR16  MDIOS output data register 16
//	0x0E0 32  DOUTR17  MDIOS output data register 17
//	0x0E4 32  DOUTR18  MDIOS output data register 18
//	0x0E8 32  DOUTR19  MDIOS output data register 19
//	0x0EC 32  DOUTR20  MDIOS output data register 20
//	0x0F0 32  DOUTR21  MDIOS output data register 21
//	0x0F4 32  DOUTR22  MDIOS output data register 22
//	0x0F8 32  DOUTR23  MDIOS output data register 23
//	0x0FC 32  DOUTR24  MDIOS output data register 24
//	0x100 32  DOUTR25  MDIOS output data register 25
//	0x104 32  DOUTR26  MDIOS output data register 26
//	0x108 32  DOUTR27  MDIOS output data register 27
//	0x10C 32  DOUTR28  MDIOS output data register 28
//	0x110 32  DOUTR29  MDIOS output data register 29
//	0x114 32  DOUTR30  MDIOS output data register 30
//	0x118 32  DOUTR31  MDIOS output data register 31
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package mdios

const (
	EN           CR = 0x01 << 0 //+ Peripheral enable
	WRIE         CR = 0x01 << 1 //+ Register write interrupt enable
	RDIE         CR = 0x01 << 2 //+ Register Read Interrupt Enable
	EIE          CR = 0x01 << 3 //+ Error interrupt enable
	DPC          CR = 0x01 << 7 //+ Disable Preamble Check
	PORT_ADDRESS CR = 0x1F << 8 //+ Slaves's address
)

const (
	ENn           = 0
	WRIEn         = 1
	RDIEn         = 2
	EIEn          = 3
	DPCn          = 7
	PORT_ADDRESSn = 8
)

const (
	WRF WRFR = 0xFFFFFFFF << 0 //+ Write flags for MDIO registers 0 to 31
)

const (
	WRFn = 0
)

const (
	CWRF CWRFR = 0xFFFFFFFF << 0 //+ Clear the write flag
)

const (
	CWRFn = 0
)

const (
	RDF RDFR = 0xFFFFFFFF << 0 //+ Read flags for MDIO registers 0 to 31
)

const (
	RDFn = 0
)

const (
	CRDF CRDFR = 0xFFFFFFFF << 0 //+ Clear the read flag
)

const (
	CRDFn = 0
)

const (
	PERF SR = 0x01 << 0 //+ Preamble error flag
	SERF SR = 0x01 << 1 //+ Start error flag
	TERF SR = 0x01 << 2 //+ Turnaround error flag
)

const (
	PERFn = 0
	SERFn = 1
	TERFn = 2
)

const (
	CPERF CLRFR = 0x01 << 0 //+ Clear the preamble error flag
	CSERF CLRFR = 0x01 << 1 //+ Clear the start error flag
	CTERF CLRFR = 0x01 << 2 //+ Clear the turnaround error flag
)

const (
	CPERFn = 0
	CSERFn = 1
	CTERFn = 2
)

const (
	DIN0 DINR0 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN0n = 0
)

const (
	DIN1 DINR1 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN1n = 0
)

const (
	DIN2 DINR2 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN2n = 0
)

const (
	DIN3 DINR3 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN3n = 0
)

const (
	DIN4 DINR4 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN4n = 0
)

const (
	DIN5 DINR5 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN5n = 0
)

const (
	DIN6 DINR6 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN6n = 0
)

const (
	DIN7 DINR7 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN7n = 0
)

const (
	DIN8 DINR8 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN8n = 0
)

const (
	DIN9 DINR9 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN9n = 0
)

const (
	DIN10 DINR10 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN10n = 0
)

const (
	DIN11 DINR11 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN11n = 0
)

const (
	DIN12 DINR12 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN12n = 0
)

const (
	DIN13 DINR13 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN13n = 0
)

const (
	DIN14 DINR14 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN14n = 0
)

const (
	DIN15 DINR15 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN15n = 0
)

const (
	DIN16 DINR16 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN16n = 0
)

const (
	DIN17 DINR17 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN17n = 0
)

const (
	DIN18 DINR18 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN18n = 0
)

const (
	DIN19 DINR19 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN19n = 0
)

const (
	DIN20 DINR20 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN20n = 0
)

const (
	DIN21 DINR21 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN21n = 0
)

const (
	DIN22 DINR22 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN22n = 0
)

const (
	DIN23 DINR23 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN23n = 0
)

const (
	DIN24 DINR24 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN24n = 0
)

const (
	DIN25 DINR25 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN25n = 0
)

const (
	DIN26 DINR26 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN26n = 0
)

const (
	DIN27 DINR27 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN27n = 0
)

const (
	DIN28 DINR28 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN28n = 0
)

const (
	DIN29 DINR29 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN29n = 0
)

const (
	DIN30 DINR30 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN30n = 0
)

const (
	DIN31 DINR31 = 0xFFFF << 0 //+ Input data received from MDIO Master during write frames
)

const (
	DIN31n = 0
)

const (
	DOUT0 DOUTR0 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT0n = 0
)

const (
	DOUT1 DOUTR1 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT1n = 0
)

const (
	DOUT2 DOUTR2 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT2n = 0
)

const (
	DOUT3 DOUTR3 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT3n = 0
)

const (
	DOUT4 DOUTR4 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT4n = 0
)

const (
	DOUT5 DOUTR5 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT5n = 0
)

const (
	DOUT6 DOUTR6 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT6n = 0
)

const (
	DOUT7 DOUTR7 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT7n = 0
)

const (
	DOUT8 DOUTR8 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT8n = 0
)

const (
	DOUT9 DOUTR9 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT9n = 0
)

const (
	DOUT10 DOUTR10 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT10n = 0
)

const (
	DOUT11 DOUTR11 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT11n = 0
)

const (
	DOUT12 DOUTR12 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT12n = 0
)

const (
	DOUT13 DOUTR13 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT13n = 0
)

const (
	DOUT14 DOUTR14 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT14n = 0
)

const (
	DOUT15 DOUTR15 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT15n = 0
)

const (
	DOUT16 DOUTR16 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT16n = 0
)

const (
	DOUT17 DOUTR17 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT17n = 0
)

const (
	DOUT18 DOUTR18 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT18n = 0
)

const (
	DOUT19 DOUTR19 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT19n = 0
)

const (
	DOUT20 DOUTR20 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT20n = 0
)

const (
	DOUT21 DOUTR21 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT21n = 0
)

const (
	DOUT22 DOUTR22 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT22n = 0
)

const (
	DOUT23 DOUTR23 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT23n = 0
)

const (
	DOUT24 DOUTR24 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT24n = 0
)

const (
	DOUT25 DOUTR25 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT25n = 0
)

const (
	DOUT26 DOUTR26 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT26n = 0
)

const (
	DOUT27 DOUTR27 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT27n = 0
)

const (
	DOUT28 DOUTR28 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT28n = 0
)

const (
	DOUT29 DOUTR29 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT29n = 0
)

const (
	DOUT30 DOUTR30 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT30n = 0
)

const (
	DOUT31 DOUTR31 = 0xFFFF << 0 //+ Output data sent to MDIO Master during read frames
)

const (
	DOUT31n = 0
)
