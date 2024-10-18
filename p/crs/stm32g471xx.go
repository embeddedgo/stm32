// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32g471xx

// Package crs provides access to the registers of the CRS peripheral.
//
// Instances:
//
//	CRS  CRS_BASE  APB1  -
//
// Registers:
//
//	0x000 32  CR    CRS control register
//	0x004 32  CFGR  This register can be written only when the frequency error counter is disabled (CEN bit is cleared in CRS_CR). When the counter is enabled, this register is write-protected.
//	0x008 32  ISR   CRS interrupt and status register
//	0x00C 32  ICR   CRS interrupt flag clear register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package crs

const (
	SYNCOKIE   CR = 0x01 << 0 //+ SYNC event OK interrupt enable
	SYNCWARNIE CR = 0x01 << 1 //+ SYNC warning interrupt enable
	ERRIE      CR = 0x01 << 2 //+ Synchronization or trimming error interrupt enable
	ESYNCIE    CR = 0x01 << 3 //+ Expected SYNC interrupt enable
	CEN        CR = 0x01 << 5 //+ Frequency error counter enable This bit enables the oscillator clock for the frequency error counter. When this bit is set, the CRS_CFGR register is write-protected and cannot be modified.
	AUTOTRIMEN CR = 0x01 << 6 //+ Automatic trimming enable This bit enables the automatic hardware adjustment of TRIM bits according to the measured frequency error between two SYNC events. If this bit is set, the TRIM bits are read-only. The TRIM value can be adjusted by hardware by one or two steps at a time, depending on the measured frequency error value. Refer to Section7.3.4: Frequency error evaluation and automatic trimming for more details.
	SWSYNC     CR = 0x01 << 7 //+ Generate software SYNC event This bit is set by software in order to generate a software SYNC event. It is automatically cleared by hardware.
	TRIM       CR = 0x7F << 8 //+ HSI48 oscillator smooth trimming These bits provide a user-programmable trimming value to the HSI48 oscillator. They can be programmed to adjust to variations in voltage and temperature that influence the frequency of the HSI48. The default value is 32, which corresponds to the middle of the trimming interval. The trimming step is around 67 kHz between two consecutive TRIM steps. A higher TRIM value corresponds to a higher output frequency. When the AUTOTRIMEN bit is set, this field is controlled by hardware and is read-only.
)

const (
	SYNCOKIEn   = 0
	SYNCWARNIEn = 1
	ERRIEn      = 2
	ESYNCIEn    = 3
	CENn        = 5
	AUTOTRIMENn = 6
	SWSYNCn     = 7
	TRIMn       = 8
)

const (
	RELOAD  CFGR = 0xFFFF << 0 //+ Counter reload value RELOAD is the value to be loaded in the frequency error counter with each SYNC event. Refer to Section7.3.3: Frequency error measurement for more details about counter behavior.
	FELIM   CFGR = 0xFF << 16  //+ Frequency error limit FELIM contains the value to be used to evaluate the captured frequency error value latched in the FECAP[15:0] bits of the CRS_ISR register. Refer to Section7.3.4: Frequency error evaluation and automatic trimming for more details about FECAP evaluation.
	SYNCDIV CFGR = 0x07 << 24  //+ SYNC divider These bits are set and cleared by software to control the division factor of the SYNC signal.
	SYNCSRC CFGR = 0x03 << 28  //+ SYNC signal source selection These bits are set and cleared by software to select the SYNC signal source. Note: When using USB LPM (Link Power Management) and the device is in Sleep mode, the periodic USB SOF will not be generated by the host. No SYNC signal will therefore be provided to the CRS to calibrate the HSI48 on the run. To guarantee the required clock precision after waking up from Sleep mode, the LSE or reference clock on the GPIOs should be used as SYNC signal.
	SYNCPOL CFGR = 0x01 << 31  //+ SYNC polarity selection This bit is set and cleared by software to select the input polarity for the SYNC signal source.
)

const (
	RELOADn  = 0
	FELIMn   = 16
	SYNCDIVn = 24
	SYNCSRCn = 28
	SYNCPOLn = 31
)

const (
	SYNCOKF   ISR = 0x01 << 0    //+ SYNC event OK flag This flag is set by hardware when the measured frequency error is smaller than FELIM * 3. This means that either no adjustment of the TRIM value is needed or that an adjustment by one trimming step is enough to compensate the frequency error. An interrupt is generated if the SYNCOKIE bit is set in the CRS_CR register. It is cleared by software by setting the SYNCOKC bit in the CRS_ICR register.
	SYNCWARNF ISR = 0x01 << 1    //+ SYNC warning flag This flag is set by hardware when the measured frequency error is greater than or equal to FELIM * 3, but smaller than FELIM * 128. This means that to compensate the frequency error, the TRIM value must be adjusted by two steps or more. An interrupt is generated if the SYNCWARNIE bit is set in the CRS_CR register. It is cleared by software by setting the SYNCWARNC bit in the CRS_ICR register.
	ERRF      ISR = 0x01 << 2    //+ Error flag This flag is set by hardware in case of any synchronization or trimming error. It is the logical OR of the TRIMOVF, SYNCMISS and SYNCERR bits. An interrupt is generated if the ERRIE bit is set in the CRS_CR register. It is cleared by software in reaction to setting the ERRC bit in the CRS_ICR register, which clears the TRIMOVF, SYNCMISS and SYNCERR bits.
	ESYNCF    ISR = 0x01 << 3    //+ Expected SYNC flag This flag is set by hardware when the frequency error counter reached a zero value. An interrupt is generated if the ESYNCIE bit is set in the CRS_CR register. It is cleared by software by setting the ESYNCC bit in the CRS_ICR register.
	SYNCERR   ISR = 0x01 << 8    //+ SYNC error This flag is set by hardware when the SYNC pulse arrives before the ESYNC event and the measured frequency error is greater than or equal to FELIM * 128. This means that the frequency error is too big (internal frequency too low) to be compensated by adjusting the TRIM value, and that some other action should be taken. An interrupt is generated if the ERRIE bit is set in the CRS_CR register. It is cleared by software by setting the ERRC bit in the CRS_ICR register.
	SYNCMISS  ISR = 0x01 << 9    //+ SYNC missed This flag is set by hardware when the frequency error counter reached value FELIM * 128 and no SYNC was detected, meaning either that a SYNC pulse was missed or that the frequency error is too big (internal frequency too high) to be compensated by adjusting the TRIM value, and that some other action should be taken. At this point, the frequency error counter is stopped (waiting for a next SYNC) and an interrupt is generated if the ERRIE bit is set in the CRS_CR register. It is cleared by software by setting the ERRC bit in the CRS_ICR register.
	TRIMOVF   ISR = 0x01 << 10   //+ Trimming overflow or underflow This flag is set by hardware when the automatic trimming tries to over- or under-flow the TRIM value. An interrupt is generated if the ERRIE bit is set in the CRS_CR register. It is cleared by software by setting the ERRC bit in the CRS_ICR register.
	FEDIR     ISR = 0x01 << 15   //+ Frequency error direction FEDIR is the counting direction of the frequency error counter latched in the time of the last SYNC event. It shows whether the actual frequency is below or above the target.
	FECAP     ISR = 0xFFFF << 16 //+ Frequency error capture FECAP is the frequency error counter value latched in the time ofthe last SYNC event. Refer to Section7.3.4: Frequency error evaluation and automatic trimming for more details about FECAP usage.
)

const (
	SYNCOKFn   = 0
	SYNCWARNFn = 1
	ERRFn      = 2
	ESYNCFn    = 3
	SYNCERRn   = 8
	SYNCMISSn  = 9
	TRIMOVFn   = 10
	FEDIRn     = 15
	FECAPn     = 16
)

const (
	SYNCOKC   ICR = 0x01 << 0 //+ SYNC event OK clear flag Writing 1 to this bit clears the SYNCOKF flag in the CRS_ISR register.
	SYNCWARNC ICR = 0x01 << 1 //+ SYNC warning clear flag Writing 1 to this bit clears the SYNCWARNF flag in the CRS_ISR register.
	ERRC      ICR = 0x01 << 2 //+ Error clear flag Writing 1 to this bit clears TRIMOVF, SYNCMISS and SYNCERR bits and consequently also the ERRF flag in the CRS_ISR register.
	ESYNCC    ICR = 0x01 << 3 //+ Expected SYNC clear flag Writing 1 to this bit clears the ESYNCF flag in the CRS_ISR register.
)

const (
	SYNCOKCn   = 0
	SYNCWARNCn = 1
	ERRCn      = 2
	ESYNCCn    = 3
)
