// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32g471xx

// Package adc3 provides access to the registers of the ADC peripheral.
//
// Instances:
//
//	ADC3  ADC3_BASE  -  ADC3  Analog-to-Digital Converter
//
// Registers:
//
//	0x000 32  ISR      interrupt and status register
//	0x004 32  IER      interrupt enable register
//	0x008 32  CR       control register
//	0x00C 32  CFGR     configuration register
//	0x010 32  CFGR2    configuration register
//	0x014 32  SMPR1    sample time register 1
//	0x018 32  SMPR2    sample time register 2
//	0x020 32  TR1      watchdog threshold register 1
//	0x024 32  TR2      watchdog threshold register
//	0x028 32  TR3      watchdog threshold register 3
//	0x030 32  SQR1     regular sequence register 1
//	0x034 32  SQR2     regular sequence register 2
//	0x038 32  SQR3     regular sequence register 3
//	0x03C 32  SQR4     regular sequence register 4
//	0x040 32  DR       regular Data Register
//	0x04C 32  JSQR     injected sequence register
//	0x060 32  OFR1     offset register 1
//	0x064 32  OFR2     offset register 2
//	0x068 32  OFR3     offset register 3
//	0x06C 32  OFR4     offset register 4
//	0x080 32  JDR1     injected data register 1
//	0x084 32  JDR2     injected data register 2
//	0x088 32  JDR3     injected data register 3
//	0x08C 32  JDR4     injected data register 4
//	0x0A0 32  AWD2CR   Analog Watchdog 2 Configuration Register
//	0x0A4 32  AWD3CR   Analog Watchdog 3 Configuration Register
//	0x0B0 32  DIFSEL   Differential Mode Selection Register 2
//	0x0B4 32  CALFACT  Calibration Factors
//	0x0C0 32  GCOMP    Gain compensation Register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package adc3

const (
	ADRDY ISR = 0x01 << 0  //+ ADRDY
	EOSMP ISR = 0x01 << 1  //+ EOSMP
	EOC   ISR = 0x01 << 2  //+ EOC
	EOS   ISR = 0x01 << 3  //+ EOS
	OVR   ISR = 0x01 << 4  //+ OVR
	JEOC  ISR = 0x01 << 5  //+ JEOC
	JEOS  ISR = 0x01 << 6  //+ JEOS
	AWD1  ISR = 0x01 << 7  //+ AWD1
	AWD2  ISR = 0x01 << 8  //+ AWD2
	AWD3  ISR = 0x01 << 9  //+ AWD3
	JQOVF ISR = 0x01 << 10 //+ JQOVF
)

const (
	ADRDYn = 0
	EOSMPn = 1
	EOCn   = 2
	EOSn   = 3
	OVRn   = 4
	JEOCn  = 5
	JEOSn  = 6
	AWD1n  = 7
	AWD2n  = 8
	AWD3n  = 9
	JQOVFn = 10
)

const (
	ADRDYIE IER = 0x01 << 0  //+ ADRDYIE
	EOSMPIE IER = 0x01 << 1  //+ EOSMPIE
	EOCIE   IER = 0x01 << 2  //+ EOCIE
	EOSIE   IER = 0x01 << 3  //+ EOSIE
	OVRIE   IER = 0x01 << 4  //+ OVRIE
	JEOCIE  IER = 0x01 << 5  //+ JEOCIE
	JEOSIE  IER = 0x01 << 6  //+ JEOSIE
	AWD1IE  IER = 0x01 << 7  //+ AWD1IE
	AWD2IE  IER = 0x01 << 8  //+ AWD2IE
	AWD3IE  IER = 0x01 << 9  //+ AWD3IE
	JQOVFIE IER = 0x01 << 10 //+ JQOVFIE
)

const (
	ADRDYIEn = 0
	EOSMPIEn = 1
	EOCIEn   = 2
	EOSIEn   = 3
	OVRIEn   = 4
	JEOCIEn  = 5
	JEOSIEn  = 6
	AWD1IEn  = 7
	AWD2IEn  = 8
	AWD3IEn  = 9
	JQOVFIEn = 10
)

const (
	ADEN     CR = 0x01 << 0  //+ ADEN
	ADDIS    CR = 0x01 << 1  //+ ADDIS
	ADSTART  CR = 0x01 << 2  //+ ADSTART
	JADSTART CR = 0x01 << 3  //+ JADSTART
	ADSTP    CR = 0x01 << 4  //+ ADSTP
	JADSTP   CR = 0x01 << 5  //+ JADSTP
	ADVREGEN CR = 0x01 << 28 //+ ADVREGEN
	DEEPPWD  CR = 0x01 << 29 //+ DEEPPWD
	ADCALDIF CR = 0x01 << 30 //+ ADCALDIF
	ADCAL    CR = 0x01 << 31 //+ ADCAL
)

const (
	ADENn     = 0
	ADDISn    = 1
	ADSTARTn  = 2
	JADSTARTn = 3
	ADSTPn    = 4
	JADSTPn   = 5
	ADVREGENn = 28
	DEEPPWDn  = 29
	ADCALDIFn = 30
	ADCALn    = 31
)

const (
	DMAEN    CFGR = 0x01 << 0  //+ DMAEN
	DMACFG   CFGR = 0x01 << 1  //+ DMACFG
	RES      CFGR = 0x03 << 3  //+ RES
	ALIGN_5  CFGR = 0x01 << 5  //+ ALIGN_5
	EXTSEL   CFGR = 0x0F << 6  //+ EXTSEL
	EXTEN    CFGR = 0x03 << 10 //+ EXTEN
	OVRMOD   CFGR = 0x01 << 12 //+ OVRMOD
	CONT     CFGR = 0x01 << 13 //+ CONT
	AUTDLY   CFGR = 0x01 << 14 //+ AUTDLY
	ALIGN    CFGR = 0x01 << 15 //+ ALIGN
	DISCEN   CFGR = 0x01 << 16 //+ DISCEN
	DISCNUM  CFGR = 0x07 << 17 //+ DISCNUM
	JDISCEN  CFGR = 0x01 << 20 //+ JDISCEN
	JQM      CFGR = 0x01 << 21 //+ JQM
	AWD1SGL  CFGR = 0x01 << 22 //+ AWD1SGL
	AWD1EN   CFGR = 0x01 << 23 //+ AWD1EN
	JAWD1EN  CFGR = 0x01 << 24 //+ JAWD1EN
	JAUTO    CFGR = 0x01 << 25 //+ JAUTO
	AWDCH1CH CFGR = 0x1F << 26 //+ AWDCH1CH
	JQDIS    CFGR = 0x01 << 31 //+ Injected Queue disable
)

const (
	DMAENn    = 0
	DMACFGn   = 1
	RESn      = 3
	ALIGN_5n  = 5
	EXTSELn   = 6
	EXTENn    = 10
	OVRMODn   = 12
	CONTn     = 13
	AUTDLYn   = 14
	ALIGNn    = 15
	DISCENn   = 16
	DISCNUMn  = 17
	JDISCENn  = 20
	JQMn      = 21
	AWD1SGLn  = 22
	AWD1ENn   = 23
	JAWD1ENn  = 24
	JAUTOn    = 25
	AWDCH1CHn = 26
	JQDISn    = 31
)

const (
	ROVSE   CFGR2 = 0x01 << 0  //+ DMAEN
	JOVSE   CFGR2 = 0x01 << 1  //+ DMACFG
	OVSR    CFGR2 = 0x07 << 2  //+ RES
	OVSS    CFGR2 = 0x0F << 5  //+ ALIGN
	TROVS   CFGR2 = 0x01 << 9  //+ Triggered Regular Oversampling
	ROVSM   CFGR2 = 0x01 << 10 //+ EXTEN
	GCOMP   CFGR2 = 0x01 << 16 //+ GCOMP
	SWTRIG  CFGR2 = 0x01 << 25 //+ SWTRIG
	BULB    CFGR2 = 0x01 << 26 //+ BULB
	SMPTRIG CFGR2 = 0x01 << 27 //+ SMPTRIG
)

const (
	ROVSEn   = 0
	JOVSEn   = 1
	OVSRn    = 2
	OVSSn    = 5
	TROVSn   = 9
	ROVSMn   = 10
	GCOMPn   = 16
	SWTRIGn  = 25
	BULBn    = 26
	SMPTRIGn = 27
)

const (
	SMP0    SMPR1 = 0x07 << 0  //+ SMP0
	SMP1    SMPR1 = 0x07 << 3  //+ SMP1
	SMP2    SMPR1 = 0x07 << 6  //+ SMP2
	SMP3    SMPR1 = 0x07 << 9  //+ SMP3
	SMP4    SMPR1 = 0x07 << 12 //+ SMP4
	SMP5    SMPR1 = 0x07 << 15 //+ SMP5
	SMP6    SMPR1 = 0x07 << 18 //+ SMP6
	SMP7    SMPR1 = 0x07 << 21 //+ SMP7
	SMP8    SMPR1 = 0x07 << 24 //+ SMP8
	SMP9    SMPR1 = 0x07 << 27 //+ SMP9
	SMPPLUS SMPR1 = 0x01 << 31 //+ Addition of one clock cycle to the sampling time
)

const (
	SMP0n    = 0
	SMP1n    = 3
	SMP2n    = 6
	SMP3n    = 9
	SMP4n    = 12
	SMP5n    = 15
	SMP6n    = 18
	SMP7n    = 21
	SMP8n    = 24
	SMP9n    = 27
	SMPPLUSn = 31
)

const (
	SMP10 SMPR2 = 0x07 << 0  //+ SMP10
	SMP11 SMPR2 = 0x07 << 3  //+ SMP11
	SMP12 SMPR2 = 0x07 << 6  //+ SMP12
	SMP13 SMPR2 = 0x07 << 9  //+ SMP13
	SMP14 SMPR2 = 0x07 << 12 //+ SMP14
	SMP15 SMPR2 = 0x07 << 15 //+ SMP15
	SMP16 SMPR2 = 0x07 << 18 //+ SMP16
	SMP17 SMPR2 = 0x07 << 21 //+ SMP17
	SMP18 SMPR2 = 0x07 << 24 //+ SMP18
)

const (
	SMP10n = 0
	SMP11n = 3
	SMP12n = 6
	SMP13n = 9
	SMP14n = 12
	SMP15n = 15
	SMP16n = 18
	SMP17n = 21
	SMP18n = 24
)

const (
	LT1     TR1 = 0xFFF << 0  //+ LT1
	AWDFILT TR1 = 0x07 << 12  //+ AWDFILT
	HT1     TR1 = 0xFFF << 16 //+ HT1
)

const (
	LT1n     = 0
	AWDFILTn = 12
	HT1n     = 16
)

const (
	LT2 TR2 = 0xFF << 0  //+ LT2
	HT2 TR2 = 0xFF << 16 //+ HT2
)

const (
	LT2n = 0
	HT2n = 16
)

const (
	LT3 TR3 = 0xFF << 0  //+ LT3
	HT3 TR3 = 0xFF << 16 //+ HT3
)

const (
	LT3n = 0
	HT3n = 16
)

const (
	L   SQR1 = 0x0F << 0  //+ Regular channel sequence length
	SQ1 SQR1 = 0x1F << 6  //+ SQ1
	SQ2 SQR1 = 0x1F << 12 //+ SQ2
	SQ3 SQR1 = 0x1F << 18 //+ SQ3
	SQ4 SQR1 = 0x1F << 24 //+ SQ4
)

const (
	Ln   = 0
	SQ1n = 6
	SQ2n = 12
	SQ3n = 18
	SQ4n = 24
)

const (
	SQ5 SQR2 = 0x1F << 0  //+ SQ5
	SQ6 SQR2 = 0x1F << 6  //+ SQ6
	SQ7 SQR2 = 0x1F << 12 //+ SQ7
	SQ8 SQR2 = 0x1F << 18 //+ SQ8
	SQ9 SQR2 = 0x1F << 24 //+ SQ9
)

const (
	SQ5n = 0
	SQ6n = 6
	SQ7n = 12
	SQ8n = 18
	SQ9n = 24
)

const (
	SQ10 SQR3 = 0x1F << 0  //+ SQ10
	SQ11 SQR3 = 0x1F << 6  //+ SQ11
	SQ12 SQR3 = 0x1F << 12 //+ SQ12
	SQ13 SQR3 = 0x1F << 18 //+ SQ13
	SQ14 SQR3 = 0x1F << 24 //+ SQ14
)

const (
	SQ10n = 0
	SQ11n = 6
	SQ12n = 12
	SQ13n = 18
	SQ14n = 24
)

const (
	SQ15 SQR4 = 0x1F << 0 //+ SQ15
	SQ16 SQR4 = 0x1F << 6 //+ SQ16
)

const (
	SQ15n = 0
	SQ16n = 6
)

const (
	RDATA DR = 0xFFFF << 0 //+ Regular Data converted
)

const (
	RDATAn = 0
)

const (
	JL      JSQR = 0x03 << 0  //+ JL
	JEXTSEL JSQR = 0x1F << 2  //+ JEXTSEL
	JEXTEN  JSQR = 0x03 << 7  //+ JEXTEN
	JSQ1    JSQR = 0x1F << 9  //+ JSQ1
	JSQ2    JSQR = 0x1F << 15 //+ JSQ2
	JSQ3    JSQR = 0x1F << 21 //+ JSQ3
	JSQ4    JSQR = 0x1F << 27 //+ JSQ4
)

const (
	JLn      = 0
	JEXTSELn = 2
	JEXTENn  = 7
	JSQ1n    = 9
	JSQ2n    = 15
	JSQ3n    = 21
	JSQ4n    = 27
)

const (
	OFFSET1    OFR1 = 0xFFF << 0 //+ OFFSET1
	OFFSETPOS  OFR1 = 0x01 << 24 //+ OFFSETPOS
	SATEN      OFR1 = 0x01 << 25 //+ SATEN
	OFFSET1_CH OFR1 = 0x1F << 26 //+ OFFSET1_CH
	OFFSET1_EN OFR1 = 0x01 << 31 //+ OFFSET1_EN
)

const (
	OFFSET1n    = 0
	OFFSETPOSn  = 24
	SATENn      = 25
	OFFSET1_CHn = 26
	OFFSET1_ENn = 31
)

const (
	OFFSET1    OFR2 = 0xFFF << 0 //+ OFFSET1
	OFFSETPOS  OFR2 = 0x01 << 24 //+ OFFSETPOS
	SATEN      OFR2 = 0x01 << 25 //+ SATEN
	OFFSET1_CH OFR2 = 0x1F << 26 //+ OFFSET1_CH
	OFFSET1_EN OFR2 = 0x01 << 31 //+ OFFSET1_EN
)

const (
	OFFSET1n    = 0
	OFFSETPOSn  = 24
	SATENn      = 25
	OFFSET1_CHn = 26
	OFFSET1_ENn = 31
)

const (
	OFFSET1    OFR3 = 0xFFF << 0 //+ OFFSET1
	OFFSETPOS  OFR3 = 0x01 << 24 //+ OFFSETPOS
	SATEN      OFR3 = 0x01 << 25 //+ SATEN
	OFFSET1_CH OFR3 = 0x1F << 26 //+ OFFSET1_CH
	OFFSET1_EN OFR3 = 0x01 << 31 //+ OFFSET1_EN
)

const (
	OFFSET1n    = 0
	OFFSETPOSn  = 24
	SATENn      = 25
	OFFSET1_CHn = 26
	OFFSET1_ENn = 31
)

const (
	OFFSET1    OFR4 = 0xFFF << 0 //+ OFFSET1
	OFFSETPOS  OFR4 = 0x01 << 24 //+ OFFSETPOS
	SATEN      OFR4 = 0x01 << 25 //+ SATEN
	OFFSET1_CH OFR4 = 0x1F << 26 //+ OFFSET1_CH
	OFFSET1_EN OFR4 = 0x01 << 31 //+ OFFSET1_EN
)

const (
	OFFSET1n    = 0
	OFFSETPOSn  = 24
	SATENn      = 25
	OFFSET1_CHn = 26
	OFFSET1_ENn = 31
)

const (
	JDATA1 JDR1 = 0xFFFF << 0 //+ JDATA1
)

const (
	JDATA1n = 0
)

const (
	JDATA2 JDR2 = 0xFFFF << 0 //+ JDATA2
)

const (
	JDATA2n = 0
)

const (
	JDATA3 JDR3 = 0xFFFF << 0 //+ JDATA3
)

const (
	JDATA3n = 0
)

const (
	JDATA4 JDR4 = 0xFFFF << 0 //+ JDATA4
)

const (
	JDATA4n = 0
)

const (
	AWD2CH AWD2CR = 0x7FFFF << 0 //+ AWD2CH
)

const (
	AWD2CHn = 0
)

const (
	AWD3CH AWD3CR = 0x7FFFF << 0 //+ AWD3CH
)

const (
	AWD3CHn = 0
)

const (
	DIFSEL_0    DIFSEL = 0x01 << 0    //+ Differential mode for channels 0
	DIFSEL_1_18 DIFSEL = 0x3FFFF << 1 //+ Differential mode for channels 15 to 1
)

const (
	DIFSEL_0n    = 0
	DIFSEL_1_18n = 1
)

const (
	CALFACT_S CALFACT = 0x7F << 0  //+ CALFACT_S
	CALFACT_D CALFACT = 0x7F << 16 //+ CALFACT_D
)

const (
	CALFACT_Sn = 0
	CALFACT_Dn = 16
)

const (
	GCOMPCOEFF GCOMP = 0x3FFF << 0 //+ GCOMPCOEFF
)

const (
	GCOMPCOEFFn = 0
)
