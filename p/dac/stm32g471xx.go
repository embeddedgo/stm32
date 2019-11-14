// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32g471xx

// Package dac provides access to the registers of the DAC1 peripheral.
//
// Instances:
//  DAC1  DAC1_BASE  AHB2  TIM6_DACUNDER*  Digital-to-analog converter
//  DAC2  DAC2_BASE  AHB2  TIM7_DACUNDER*  Digital-to-analog converter
//  DAC3  DAC3_BASE  AHB2  TIM6_DACUNDER*  Digital-to-analog converter
//  DAC4  DAC4_BASE  AHB2  TIM7_DACUNDER*  Digital-to-analog converter
// Registers:
//  0x000 32  DAC_CR       DAC control register
//  0x004 32  DAC_SWTRGR   DAC software trigger register
//  0x008 32  DAC_DHR12R1  DAC channel1 12-bit right-aligned data holding register
//  0x00C 32  DAC_DHR12L1  DAC channel1 12-bit left aligned data holding register
//  0x010 32  DAC_DHR8R1   DAC channel1 8-bit right aligned data holding register
//  0x014 32  DAC_DHR12R2  DAC channel2 12-bit right aligned data holding register
//  0x018 32  DAC_DHR12L2  DAC channel2 12-bit left aligned data holding register
//  0x01C 32  DAC_DHR8R2   DAC channel2 8-bit right-aligned data holding register
//  0x020 32  DAC_DHR12RD  Dual DAC 12-bit right-aligned data holding register
//  0x024 32  DAC_DHR12LD  DUAL DAC 12-bit left aligned data holding register
//  0x028 32  DAC_DHR8RD   DUAL DAC 8-bit right aligned data holding register
//  0x02C 32  DAC_DOR1     DAC channel1 data output register
//  0x030 32  DAC_DOR2     DAC channel2 data output register
//  0x034 32  DAC_SR       DAC status register
//  0x038 32  DAC_CCR      DAC calibration control register
//  0x03C 32  DAC_MCR      DAC mode control register
//  0x040 32  DAC_SHSR1    DAC Sample and Hold sample time register 1
//  0x044 32  DAC_SHSR2    DAC Sample and Hold sample time register 2
//  0x048 32  DAC_SHHR     DAC Sample and Hold hold time register
//  0x04C 32  DAC_SHRR     DAC Sample and Hold refresh time register
//  0x058 32  DAC_STR1     Sawtooth register
//  0x05C 32  DAC_STR2     Sawtooth register
//  0x060 32  DAC_STMODR   Sawtooth Mode register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package dac

const (
	EN1       DAC_CR = 0x01 << 0  //+ DAC channel1 enable This bit is set and cleared by software to enable/disable DAC channel1.
	TEN1      DAC_CR = 0x01 << 1  //+ DAC channel1 trigger enable
	TSEL1     DAC_CR = 0x0F << 2  //+ DAC channel1 trigger selection These bits select the external event used to trigger DAC channel1. Note: Only used if bit TEN1 = 1 (DAC channel1 trigger enabled).
	WAVE1     DAC_CR = 0x03 << 6  //+ DAC channel1 noise/triangle wave generation enable These bits are set and cleared by software. Note: Only used if bit TEN1 = 1 (DAC channel1 trigger enabled).
	MAMP1     DAC_CR = 0x0F << 8  //+ DAC channel1 mask/amplitude selector These bits are written by software to select mask in wave generation mode or amplitude in triangle generation mode. = 1011: Unmask bits[11:0] of LFSR/ triangle amplitude equal to 4095
	DMAEN1    DAC_CR = 0x01 << 12 //+ DAC channel1 DMA enable This bit is set and cleared by software.
	DMAUDRIE1 DAC_CR = 0x01 << 13 //+ DAC channel1 DMA Underrun Interrupt enable This bit is set and cleared by software.
	CEN1      DAC_CR = 0x01 << 14 //+ DAC Channel 1 calibration enable This bit is set and cleared by software to enable/disable DAC channel 1 calibration, it can be written only if bit EN1=0 into DAC_CR (the calibration mode can be entered/exit only when the DAC channel is disabled) Otherwise, the write operation is ignored.
	EN2       DAC_CR = 0x01 << 16 //+ DAC channel2 enable This bit is set and cleared by software to enable/disable DAC channel2.
	TEN2      DAC_CR = 0x01 << 17 //+ DAC channel2 trigger enable
	TSEL2     DAC_CR = 0x0F << 18 //+ DAC channel2 trigger selection These bits select the external event used to trigger DAC channel2 Note: Only used if bit TEN2 = 1 (DAC channel2 trigger enabled).
	WAVE2     DAC_CR = 0x03 << 22 //+ DAC channel2 noise/triangle wave generation enable These bits are set/reset by software. 1x: Triangle wave generation enabled Note: Only used if bit TEN2 = 1 (DAC channel2 trigger enabled)
	MAMP2     DAC_CR = 0x0F << 24 //+ DAC channel2 mask/amplitude selector These bits are written by software to select mask in wave generation mode or amplitude in triangle generation mode. = 1011: Unmask bits[11:0] of LFSR/ triangle amplitude equal to 4095
	DMAEN2    DAC_CR = 0x01 << 28 //+ DAC channel2 DMA enable This bit is set and cleared by software.
	DMAUDRIE2 DAC_CR = 0x01 << 29 //+ DAC channel2 DMA underrun interrupt enable This bit is set and cleared by software.
	CEN2      DAC_CR = 0x01 << 30 //+ DAC Channel 2 calibration enable This bit is set and cleared by software to enable/disable DAC channel 2 calibration, it can be written only if bit EN2=0 into DAC_CR (the calibration mode can be entered/exit only when the DAC channel is disabled) Otherwise, the write operation is ignored.
)

const (
	EN1n       = 0
	TEN1n      = 1
	TSEL1n     = 2
	WAVE1n     = 6
	MAMP1n     = 8
	DMAEN1n    = 12
	DMAUDRIE1n = 13
	CEN1n      = 14
	EN2n       = 16
	TEN2n      = 17
	TSEL2n     = 18
	WAVE2n     = 22
	MAMP2n     = 24
	DMAEN2n    = 28
	DMAUDRIE2n = 29
	CEN2n      = 30
)

const (
	SWTRIG1  DAC_SWTRGR = 0x01 << 0  //+ DAC channel1 software trigger This bit is set by software to trigger the DAC in software trigger mode. Note: This bit is cleared by hardware (one APB1 clock cycle later) once the DAC_DHR1 register value has been loaded into the DAC_DOR1 register.
	SWTRIG2  DAC_SWTRGR = 0x01 << 1  //+ DAC channel2 software trigger This bit is set by software to trigger the DAC in software trigger mode. Note: This bit is cleared by hardware (one APB1 clock cycle later) once the DAC_DHR2 register value has been loaded into the DAC_DOR2 register.
	SWTRIGB1 DAC_SWTRGR = 0x01 << 16 //+ DAC channel1 software trigger B
	SWTRIGB2 DAC_SWTRGR = 0x01 << 17 //+ DAC channel2 software trigger B
)

const (
	SWTRIG1n  = 0
	SWTRIG2n  = 1
	SWTRIGB1n = 16
	SWTRIGB2n = 17
)

const (
	DACC1DHR  DAC_DHR12R1 = 0xFFF << 0  //+ DAC channel1 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
	DACC1DHRB DAC_DHR12R1 = 0xFFF << 16 //+ DAC channel1 12-bit right-aligned data B
)

const (
	DACC1DHRn  = 0
	DACC1DHRBn = 16
)

const (
	DACC1DHR  DAC_DHR12L1 = 0xFFF << 4  //+ DAC channel1 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
	DACC1DHRB DAC_DHR12L1 = 0xFFF << 20 //+ DAC channel1 12-bit left-aligned data B
)

const (
	DACC1DHRn  = 4
	DACC1DHRBn = 20
)

const (
	DACC1DHR  DAC_DHR8R1 = 0xFF << 0 //+ DAC channel1 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel1.
	DACC1DHRB DAC_DHR8R1 = 0xFF << 8 //+ DAC channel1 8-bit right-aligned data
)

const (
	DACC1DHRn  = 0
	DACC1DHRBn = 8
)

const (
	DACC2DHR  DAC_DHR12R2 = 0xFFF << 0  //+ DAC channel2 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.
	DACC2DHRB DAC_DHR12R2 = 0xFFF << 16 //+ DAC channel2 12-bit right-aligned data
)

const (
	DACC2DHRn  = 0
	DACC2DHRBn = 16
)

const (
	DACC2DHR  DAC_DHR12L2 = 0xFFF << 4  //+ DAC channel2 12-bit left-aligned data These bits are written by software which specify 12-bit data for DAC channel2.
	DACC2DHRB DAC_DHR12L2 = 0xFFF << 20 //+ DAC channel2 12-bit left-aligned data B
)

const (
	DACC2DHRn  = 4
	DACC2DHRBn = 20
)

const (
	DACC2DHR  DAC_DHR8R2 = 0xFF << 0 //+ DAC channel2 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel2.
	DACC2DHRB DAC_DHR8R2 = 0xFF << 8 //+ DAC channel2 8-bit right-aligned data
)

const (
	DACC2DHRn  = 0
	DACC2DHRBn = 8
)

const (
	DACC1DHR DAC_DHR12RD = 0xFFF << 0  //+ DAC channel1 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
	DACC2DHR DAC_DHR12RD = 0xFFF << 16 //+ DAC channel2 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.
)

const (
	DACC1DHRn = 0
	DACC2DHRn = 16
)

const (
	DACC1DHR DAC_DHR12LD = 0xFFF << 4  //+ DAC channel1 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.
	DACC2DHR DAC_DHR12LD = 0xFFF << 20 //+ DAC channel2 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.
)

const (
	DACC1DHRn = 4
	DACC2DHRn = 20
)

const (
	DACC1DHR DAC_DHR8RD = 0xFF << 0 //+ DAC channel1 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel1.
	DACC2DHR DAC_DHR8RD = 0xFF << 8 //+ DAC channel2 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel2.
)

const (
	DACC1DHRn = 0
	DACC2DHRn = 8
)

const (
	DACC1DOR  DAC_DOR1 = 0xFFF << 0  //+ DAC channel1 data output These bits are read-only, they contain data output for DAC channel1.
	DACC1DORB DAC_DOR1 = 0xFFF << 16 //+ DAC channel1 data output
)

const (
	DACC1DORn  = 0
	DACC1DORBn = 16
)

const (
	DACC2DOR  DAC_DOR2 = 0xFFF << 0  //+ DAC channel2 data output These bits are read-only, they contain data output for DAC channel2.
	DACC2DORB DAC_DOR2 = 0xFFF << 16 //+ DAC channel2 data output
)

const (
	DACC2DORn  = 0
	DACC2DORBn = 16
)

const (
	DAC1RDY   DAC_SR = 0x01 << 11 //+ DAC channel1 ready status bit
	DORSTAT1  DAC_SR = 0x01 << 12 //+ DAC channel1 output register status bit
	DMAUDR1   DAC_SR = 0x01 << 13 //+ DAC channel1 DMA underrun flag This bit is set by hardware and cleared by software (by writing it to 1).
	CAL_FLAG1 DAC_SR = 0x01 << 14 //+ DAC Channel 1 calibration offset status This bit is set and cleared by hardware
	BWST1     DAC_SR = 0x01 << 15 //+ DAC Channel 1 busy writing sample time flag This bit is systematically set just after Sample & Hold mode enable and is set each time the software writes the register DAC_SHSR1, It is cleared by hardware when the write operation of DAC_SHSR1 is complete. (It takes about 3LSI periods of synchronization).
	DAC2RDY   DAC_SR = 0x01 << 27 //+ DAC channel 2 ready status bit
	DORSTAT2  DAC_SR = 0x01 << 28 //+ DAC channel 2 output register status bit
	DMAUDR2   DAC_SR = 0x01 << 29 //+ DAC channel2 DMA underrun flag This bit is set by hardware and cleared by software (by writing it to 1).
	CAL_FLAG2 DAC_SR = 0x01 << 30 //+ DAC Channel 2 calibration offset status This bit is set and cleared by hardware
	BWST2     DAC_SR = 0x01 << 31 //+ DAC Channel 2 busy writing sample time flag This bit is systematically set just after Sample & Hold mode enable and is set each time the software writes the register DAC_SHSR2, It is cleared by hardware when the write operation of DAC_SHSR2 is complete. (It takes about 3 LSI periods of synchronization).
)

const (
	DAC1RDYn   = 11
	DORSTAT1n  = 12
	DMAUDR1n   = 13
	CAL_FLAG1n = 14
	BWST1n     = 15
	DAC2RDYn   = 27
	DORSTAT2n  = 28
	DMAUDR2n   = 29
	CAL_FLAG2n = 30
	BWST2n     = 31
)

const (
	OTRIM1 DAC_CCR = 0x1F << 0  //+ DAC Channel 1 offset trimming value
	OTRIM2 DAC_CCR = 0x1F << 16 //+ DAC Channel 2 offset trimming value
)

const (
	OTRIM1n = 0
	OTRIM2n = 16
)

const (
	MODE1      DAC_MCR = 0x07 << 0  //+ DAC Channel 1 mode These bits can be written only when the DAC is disabled and not in the calibration mode (when bit EN1=0 and bit CEN1 =0 in the DAC_CR register). If EN1=1 or CEN1 =1 the write operation is ignored. They can be set and cleared by software to select the DAC Channel 1 mode: DAC Channel 1 in normal Mode DAC Channel 1 in sample &amp; hold mode
	DMADOUBLE1 DAC_MCR = 0x01 << 8  //+ DAC Channel1 DMA double data mode
	SINFORMAT1 DAC_MCR = 0x01 << 9  //+ Enable signed format for DAC channel1
	HFSEL      DAC_MCR = 0x03 << 14 //+ High frequency interface mode selection
	MODE2      DAC_MCR = 0x07 << 16 //+ DAC Channel 2 mode These bits can be written only when the DAC is disabled and not in the calibration mode (when bit EN2=0 and bit CEN2 =0 in the DAC_CR register). If EN2=1 or CEN2 =1 the write operation is ignored. They can be set and cleared by software to select the DAC Channel 2 mode: DAC Channel 2 in normal Mode DAC Channel 2 in sample &amp; hold mode
	DMADOUBLE2 DAC_MCR = 0x01 << 24 //+ DAC Channel2 DMA double data mode
	SINFORMAT2 DAC_MCR = 0x01 << 25 //+ Enable signed format for DAC channel2
)

const (
	MODE1n      = 0
	DMADOUBLE1n = 8
	SINFORMAT1n = 9
	HFSELn      = 14
	MODE2n      = 16
	DMADOUBLE2n = 24
	SINFORMAT2n = 25
)

const (
	TSAMPLE1 DAC_SHSR1 = 0x3FF << 0 //+ DAC Channel 1 sample Time (only valid in sample &amp; hold mode) These bits can be written when the DAC channel1 is disabled or also during normal operation. in the latter case, the write can be done only when BWSTx of DAC_SR register is low, If BWSTx=1, the write operation is ignored.
)

const (
	TSAMPLE1n = 0
)

const (
	TSAMPLE2 DAC_SHSR2 = 0x3FF << 0 //+ DAC Channel 2 sample Time (only valid in sample &amp; hold mode) These bits can be written when the DAC channel2 is disabled or also during normal operation. in the latter case, the write can be done only when BWSTx of DAC_SR register is low, if BWSTx=1, the write operation is ignored.
)

const (
	TSAMPLE2n = 0
)

const (
	THOLD1 DAC_SHHR = 0x3FF << 0  //+ DAC Channel 1 hold Time (only valid in sample &amp; hold mode) Hold time= (THOLD[9:0]) x T LSI
	THOLD2 DAC_SHHR = 0x3FF << 16 //+ DAC Channel 2 hold time (only valid in sample &amp; hold mode). Hold time= (THOLD[9:0]) x T LSI
)

const (
	THOLD1n = 0
	THOLD2n = 16
)

const (
	TREFRESH1 DAC_SHRR = 0xFF << 0  //+ DAC Channel 1 refresh Time (only valid in sample &amp; hold mode) Refresh time= (TREFRESH[7:0]) x T LSI
	TREFRESH2 DAC_SHRR = 0xFF << 16 //+ DAC Channel 2 refresh Time (only valid in sample &amp; hold mode) Refresh time= (TREFRESH[7:0]) x T LSI
)

const (
	TREFRESH1n = 0
	TREFRESH2n = 16
)

const (
	STRSTDATA1 DAC_STR1 = 0xFFF << 0   //+ DAC Channel 1 Sawtooth reset value
	STDIR1     DAC_STR1 = 0x01 << 12   //+ DAC Channel1 Sawtooth direction setting
	STINCDATA1 DAC_STR1 = 0xFFFF << 16 //+ DAC CH1 Sawtooth increment value (12.4 bit format)
)

const (
	STRSTDATA1n = 0
	STDIR1n     = 12
	STINCDATA1n = 16
)

const (
	STRSTDATA2 DAC_STR2 = 0xFFF << 0   //+ DAC Channel 2 Sawtooth reset value
	STDIR2     DAC_STR2 = 0x01 << 12   //+ DAC Channel2 Sawtooth direction setting
	STINCDATA2 DAC_STR2 = 0xFFFF << 16 //+ DAC CH2 Sawtooth increment value (12.4 bit format)
)

const (
	STRSTDATA2n = 0
	STDIR2n     = 12
	STINCDATA2n = 16
)

const (
	STRSTTRIGSEL1 DAC_STMODR = 0x0F << 0  //+ DAC Channel 1 Sawtooth Reset trigger selection
	STINCTRIGSEL1 DAC_STMODR = 0x0F << 8  //+ DAC Channel 1 Sawtooth Increment trigger selection
	STRSTTRIGSEL2 DAC_STMODR = 0x0F << 16 //+ DAC Channel 1 Sawtooth Reset trigger selection
	STINCTRIGSEL2 DAC_STMODR = 0x0F << 24 //+ DAC Channel 2 Sawtooth Increment trigger selection
)

const (
	STRSTTRIGSEL1n = 0
	STINCTRIGSEL1n = 8
	STRSTTRIGSEL2n = 16
	STINCTRIGSEL2n = 24
)
