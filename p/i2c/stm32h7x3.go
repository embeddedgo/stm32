// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package i2c provides access to the registers of the I2C peripheral.
//
// Instances:
//
//	I2C1  I2C1_BASE  APB1  I2C1_EV,I2C1_ER
//	I2C2  I2C2_BASE  APB1  I2C2_EV,I2C2_ER
//	I2C3  I2C3_BASE  APB1  I2C3_EV,I2C3_ER
//	I2C4  I2C4_BASE  APB4  I2C4_EV,I2C4_ER
//
// Registers:
//
//	0x000 32  CR1       Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.
//	0x004 32  CR2       Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.
//	0x008 32  OAR1      Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.
//	0x00C 32  OAR2      Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.
//	0x010 32  TIMINGR   Access: No wait states
//	0x014 32  TIMEOUTR  Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.
//	0x018 32  ISR       Access: No wait states
//	0x01C 32  ICR       Access: No wait states
//	0x020 32  PECR      Access: No wait states
//	0x024 32  RXDR      Access: No wait states
//	0x028 32  TXDR      Access: No wait states
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package i2c

const (
	PE        CR1 = 0x01 << 0  //+ Peripheral enable Note: When PE=0, the I2C SCL and SDA lines are released. Internal state machines and status bits are put back to their reset value. When cleared, PE must be kept low for at least 3 APB clock cycles.
	TXIE      CR1 = 0x01 << 1  //+ TX Interrupt enable
	RXIE      CR1 = 0x01 << 2  //+ RX Interrupt enable
	ADDRIE    CR1 = 0x01 << 3  //+ Address match Interrupt enable (slave only)
	NACKIE    CR1 = 0x01 << 4  //+ Not acknowledge received Interrupt enable
	STOPIE    CR1 = 0x01 << 5  //+ STOP detection Interrupt enable
	TCIE      CR1 = 0x01 << 6  //+ Transfer Complete interrupt enable Note: Any of these events will generate an interrupt: Transfer Complete (TC) Transfer Complete Reload (TCR)
	ERRIE     CR1 = 0x01 << 7  //+ Error interrupts enable Note: Any of these errors generate an interrupt: Arbitration Loss (ARLO) Bus Error detection (BERR) Overrun/Underrun (OVR) Timeout detection (TIMEOUT) PEC error detection (PECERR) Alert pin event detection (ALERT)
	DNF       CR1 = 0x0F << 8  //+ Digital noise filter These bits are used to configure the digital noise filter on SDA and SCL input. The digital filter will filter spikes with a length of up to DNF[3:0] * tI2CCLK ... Note: If the analog filter is also enabled, the digital filter is added to the analog filter. This filter can only be programmed when the I2C is disabled (PE = 0).
	ANFOFF    CR1 = 0x01 << 12 //+ Analog noise filter OFF Note: This bit can only be programmed when the I2C is disabled (PE = 0).
	TXDMAEN   CR1 = 0x01 << 14 //+ DMA transmission requests enable
	RXDMAEN   CR1 = 0x01 << 15 //+ DMA reception requests enable
	SBC       CR1 = 0x01 << 16 //+ Slave byte control This bit is used to enable hardware byte control in slave mode.
	NOSTRETCH CR1 = 0x01 << 17 //+ Clock stretching disable This bit is used to disable clock stretching in slave mode. It must be kept cleared in master mode. Note: This bit can only be programmed when the I2C is disabled (PE = 0).
	WUPEN     CR1 = 0x01 << 18 //+ Wakeup from Stop mode enable Note: If the Wakeup from Stop mode feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation. Note: WUPEN can be set only when DNF = 0000
	GCEN      CR1 = 0x01 << 19 //+ General call enable
	SMBHEN    CR1 = 0x01 << 20 //+ SMBus Host address enable Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
	SMBDEN    CR1 = 0x01 << 21 //+ SMBus Device Default address enable Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
	ALERTEN   CR1 = 0x01 << 22 //+ SMBus alert enable Device mode (SMBHEN=0): Host mode (SMBHEN=1): Note: When ALERTEN=0, the SMBA pin can be used as a standard GPIO. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
	PECEN     CR1 = 0x01 << 23 //+ PEC enable Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
)

const (
	PEn        = 0
	TXIEn      = 1
	RXIEn      = 2
	ADDRIEn    = 3
	NACKIEn    = 4
	STOPIEn    = 5
	TCIEn      = 6
	ERRIEn     = 7
	DNFn       = 8
	ANFOFFn    = 12
	TXDMAENn   = 14
	RXDMAENn   = 15
	SBCn       = 16
	NOSTRETCHn = 17
	WUPENn     = 18
	GCENn      = 19
	SMBHENn    = 20
	SMBDENn    = 21
	ALERTENn   = 22
	PECENn     = 23
)

const (
	SADD0   CR2 = 0x01 << 0  //+ Slave address bit 0 (master mode) In 7-bit addressing mode (ADD10 = 0): This bit is dont care In 10-bit addressing mode (ADD10 = 1): This bit should be written with bit 0 of the slave address to be sent Note: Changing these bits when the START bit is set is not allowed.
	SADD1   CR2 = 0x01 << 1  //+ Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
	SADD2   CR2 = 0x01 << 2  //+ Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
	SADD3   CR2 = 0x01 << 3  //+ Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
	SADD4   CR2 = 0x01 << 4  //+ Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
	SADD5   CR2 = 0x01 << 5  //+ Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
	SADD6   CR2 = 0x01 << 6  //+ Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
	SADD7   CR2 = 0x01 << 7  //+ Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
	SADD8   CR2 = 0x01 << 8  //+ Slave address bit 9:8 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits are dont care In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 9:8 of the slave address to be sent Note: Changing these bits when the START bit is set is not allowed.
	SADD9   CR2 = 0x01 << 9  //+ Slave address bit 9:8 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits are dont care In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 9:8 of the slave address to be sent Note: Changing these bits when the START bit is set is not allowed.
	RD_WRN  CR2 = 0x01 << 10 //+ Transfer direction (master mode) Note: Changing this bit when the START bit is set is not allowed.
	ADD10   CR2 = 0x01 << 11 //+ 10-bit addressing mode (master mode) Note: Changing this bit when the START bit is set is not allowed.
	HEAD10R CR2 = 0x01 << 12 //+ 10-bit address header only read direction (master receiver mode) Note: Changing this bit when the START bit is set is not allowed.
	START   CR2 = 0x01 << 13 //+ Start generation This bit is set by software, and cleared by hardware after the Start followed by the address sequence is sent, by an arbitration loss, by a timeout error detection, or when PE = 0. It can also be cleared by software by writing 1 to the ADDRCF bit in the I2C_ICR register. If the I2C is already in master mode with AUTOEND = 0, setting this bit generates a Repeated Start condition when RELOAD=0, after the end of the NBYTES transfer. Otherwise setting this bit will generate a START condition once the bus is free. Note: Writing 0 to this bit has no effect. The START bit can be set even if the bus is BUSY or I2C is in slave mode. This bit has no effect when RELOAD is set.
	STOP    CR2 = 0x01 << 14 //+ Stop generation (master mode) The bit is set by software, cleared by hardware when a Stop condition is detected, or when PE = 0. In Master Mode: Note: Writing 0 to this bit has no effect.
	NACK    CR2 = 0x01 << 15 //+ NACK generation (slave mode) The bit is set by software, cleared by hardware when the NACK is sent, or when a STOP condition or an Address matched is received, or when PE=0. Note: Writing 0 to this bit has no effect. This bit is used in slave mode only: in master receiver mode, NACK is automatically generated after last byte preceding STOP or RESTART condition, whatever the NACK bit value. When an overrun occurs in slave receiver NOSTRETCH mode, a NACK is automatically generated whatever the NACK bit value. When hardware PEC checking is enabled (PECBYTE=1), the PEC acknowledge value does not depend on the NACK value.
	NBYTES  CR2 = 0xFF << 16 //+ Number of bytes The number of bytes to be transmitted/received is programmed there. This field is dont care in slave mode with SBC=0. Note: Changing these bits when the START bit is set is not allowed.
	RELOAD  CR2 = 0x01 << 24 //+ NBYTES reload mode This bit is set and cleared by software.
	AUTOEND CR2 = 0x01 << 25 //+ Automatic end mode (master mode) This bit is set and cleared by software. Note: This bit has no effect in slave mode or when the RELOAD bit is set.
	PECBYTE CR2 = 0x01 << 26 //+ Packet error checking byte This bit is set by software, and cleared by hardware when the PEC is transferred, or when a STOP condition or an Address matched is received, also when PE=0. Note: Writing 0 to this bit has no effect. This bit has no effect when RELOAD is set. This bit has no effect is slave mode when SBC=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
)

const (
	SADD0n   = 0
	SADD1n   = 1
	SADD2n   = 2
	SADD3n   = 3
	SADD4n   = 4
	SADD5n   = 5
	SADD6n   = 6
	SADD7n   = 7
	SADD8n   = 8
	SADD9n   = 9
	RD_WRNn  = 10
	ADD10n   = 11
	HEAD10Rn = 12
	STARTn   = 13
	STOPn    = 14
	NACKn    = 15
	NBYTESn  = 16
	RELOADn  = 24
	AUTOENDn = 25
	PECBYTEn = 26
)

const (
	OA1     OAR1 = 0x3FF << 0 //+ Interface address 7-bit addressing mode: dont care 10-bit addressing mode: bits 9:8 of address Note: These bits can be written only when OA1EN=0. OA1[7:1]: Interface address Bits 7:1 of address Note: These bits can be written only when OA1EN=0. OA1[0]: Interface address 7-bit addressing mode: dont care 10-bit addressing mode: bit 0 of address Note: This bit can be written only when OA1EN=0.
	OA1MODE OAR1 = 0x01 << 10 //+ Own Address 1 10-bit mode Note: This bit can be written only when OA1EN=0.
	OA1EN   OAR1 = 0x01 << 15 //+ Own Address 1 enable
)

const (
	OA1n     = 0
	OA1MODEn = 10
	OA1ENn   = 15
)

const (
	OA2    OAR2 = 0x7F << 1  //+ Interface address bits 7:1 of address Note: These bits can be written only when OA2EN=0.
	OA2MSK OAR2 = 0x07 << 8  //+ Own Address 2 masks Note: These bits can be written only when OA2EN=0. As soon as OA2MSK is not equal to 0, the reserved I2C addresses (0b0000xxx and 0b1111xxx) are not acknowledged even if the comparison matches.
	OA2EN  OAR2 = 0x01 << 15 //+ Own Address 2 enable
)

const (
	OA2n    = 1
	OA2MSKn = 8
	OA2ENn  = 15
)

const (
	SCLL   TIMINGR = 0xFF << 0  //+ SCL low period (master mode) This field is used to generate the SCL low period in master mode. tSCLL = (SCLL+1) x tPRESC Note: SCLL is also used to generate tBUF and tSU:STA timings.
	SCLH   TIMINGR = 0xFF << 8  //+ SCL high period (master mode) This field is used to generate the SCL high period in master mode. tSCLH = (SCLH+1) x tPRESC Note: SCLH is also used to generate tSU:STO and tHD:STA timing.
	SDADEL TIMINGR = 0x0F << 16 //+ Data hold time This field is used to generate the delay tSDADEL between SCL falling edge and SDA edge. In master mode and in slave mode with NOSTRETCH = 0, the SCL line is stretched low during tSDADEL. tSDADEL= SDADEL x tPRESC Note: SDADEL is used to generate tHD:DAT timing.
	SCLDEL TIMINGR = 0x0F << 20 //+ Data setup time This field is used to generate a delay tSCLDEL between SDA edge and SCL rising edge. In master mode and in slave mode with NOSTRETCH = 0, the SCL line is stretched low during tSCLDEL. tSCLDEL = (SCLDEL+1) x tPRESC Note: tSCLDEL is used to generate tSU:DAT timing.
	PRESC  TIMINGR = 0x0F << 28 //+ Timing prescaler This field is used to prescale I2CCLK in order to generate the clock period tPRESC used for data setup and hold counters (refer to I2C timings on page9) and for SCL high and low level counters (refer to I2C master initialization on page24). tPRESC = (PRESC+1) x tI2CCLK
)

const (
	SCLLn   = 0
	SCLHn   = 8
	SDADELn = 16
	SCLDELn = 20
	PRESCn  = 28
)

const (
	TIMEOUTA TIMEOUTR = 0xFFF << 0  //+ Bus Timeout A This field is used to configure: The SCL low timeout condition tTIMEOUT when TIDLE=0 tTIMEOUT= (TIMEOUTA+1) x 2048 x tI2CCLK The bus idle condition (both SCL and SDA high) when TIDLE=1 tIDLE= (TIMEOUTA+1) x 4 x tI2CCLK Note: These bits can be written only when TIMOUTEN=0.
	TIDLE    TIMEOUTR = 0x01 << 12  //+ Idle clock timeout detection Note: This bit can be written only when TIMOUTEN=0.
	TIMOUTEN TIMEOUTR = 0x01 << 15  //+ Clock timeout enable
	TIMEOUTB TIMEOUTR = 0xFFF << 16 //+ Bus timeout B This field is used to configure the cumulative clock extension timeout: In master mode, the master cumulative clock low extend time (tLOW:MEXT) is detected In slave mode, the slave cumulative clock low extend time (tLOW:SEXT) is detected tLOW:EXT= (TIMEOUTB+1) x 2048 x tI2CCLK Note: These bits can be written only when TEXTEN=0.
	TEXTEN   TIMEOUTR = 0x01 << 31  //+ Extended clock timeout enable
)

const (
	TIMEOUTAn = 0
	TIDLEn    = 12
	TIMOUTENn = 15
	TIMEOUTBn = 16
	TEXTENn   = 31
)

const (
	TXE     ISR = 0x01 << 0  //+ Transmit data register empty (transmitters) This bit is set by hardware when the I2C_TXDR register is empty. It is cleared when the next data to be sent is written in the I2C_TXDR register. This bit can be written to 1 by software in order to flush the transmit data register I2C_TXDR. Note: This bit is set by hardware when PE=0.
	TXIS    ISR = 0x01 << 1  //+ Transmit interrupt status (transmitters) This bit is set by hardware when the I2C_TXDR register is empty and the data to be transmitted must be written in the I2C_TXDR register. It is cleared when the next data to be sent is written in the I2C_TXDR register. This bit can be written to 1 by software when NOSTRETCH=1 only, in order to generate a TXIS event (interrupt if TXIE=1 or DMA request if TXDMAEN=1). Note: This bit is cleared by hardware when PE=0.
	RXNE    ISR = 0x01 << 2  //+ Receive data register not empty (receivers) This bit is set by hardware when the received data is copied into the I2C_RXDR register, and is ready to be read. It is cleared when I2C_RXDR is read. Note: This bit is cleared by hardware when PE=0.
	ADDR    ISR = 0x01 << 3  //+ Address matched (slave mode) This bit is set by hardware as soon as the received slave address matched with one of the enabled slave addresses. It is cleared by software by setting ADDRCF bit. Note: This bit is cleared by hardware when PE=0.
	NACKF   ISR = 0x01 << 4  //+ Not Acknowledge received flag This flag is set by hardware when a NACK is received after a byte transmission. It is cleared by software by setting the NACKCF bit. Note: This bit is cleared by hardware when PE=0.
	STOPF   ISR = 0x01 << 5  //+ Stop detection flag This flag is set by hardware when a Stop condition is detected on the bus and the peripheral is involved in this transfer: either as a master, provided that the STOP condition is generated by the peripheral. or as a slave, provided that the peripheral has been addressed previously during this transfer. It is cleared by software by setting the STOPCF bit. Note: This bit is cleared by hardware when PE=0.
	TC      ISR = 0x01 << 6  //+ Transfer Complete (master mode) This flag is set by hardware when RELOAD=0, AUTOEND=0 and NBYTES data have been transferred. It is cleared by software when START bit or STOP bit is set. Note: This bit is cleared by hardware when PE=0.
	TCR     ISR = 0x01 << 7  //+ Transfer Complete Reload This flag is set by hardware when RELOAD=1 and NBYTES data have been transferred. It is cleared by software when NBYTES is written to a non-zero value. Note: This bit is cleared by hardware when PE=0. This flag is only for master mode, or for slave mode when the SBC bit is set.
	BERR    ISR = 0x01 << 8  //+ Bus error This flag is set by hardware when a misplaced Start or Stop condition is detected whereas the peripheral is involved in the transfer. The flag is not set during the address phase in slave mode. It is cleared by software by setting BERRCF bit. Note: This bit is cleared by hardware when PE=0.
	ARLO    ISR = 0x01 << 9  //+ Arbitration lost This flag is set by hardware in case of arbitration loss. It is cleared by software by setting the ARLOCF bit. Note: This bit is cleared by hardware when PE=0.
	OVR     ISR = 0x01 << 10 //+ Overrun/Underrun (slave mode) This flag is set by hardware in slave mode with NOSTRETCH=1, when an overrun/underrun error occurs. It is cleared by software by setting the OVRCF bit. Note: This bit is cleared by hardware when PE=0.
	PECERR  ISR = 0x01 << 11 //+ PEC Error in reception This flag is set by hardware when the received PEC does not match with the PEC register content. A NACK is automatically sent after the wrong PEC reception. It is cleared by software by setting the PECCF bit. Note: This bit is cleared by hardware when PE=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
	TIMEOUT ISR = 0x01 << 12 //+ Timeout or tLOW detection flag This flag is set by hardware when a timeout or extended clock timeout occurred. It is cleared by software by setting the TIMEOUTCF bit. Note: This bit is cleared by hardware when PE=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
	ALERT   ISR = 0x01 << 13 //+ SMBus alert This flag is set by hardware when SMBHEN=1 (SMBus host configuration), ALERTEN=1 and a SMBALERT event (falling edge) is detected on SMBA pin. It is cleared by software by setting the ALERTCF bit. Note: This bit is cleared by hardware when PE=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
	BUSY    ISR = 0x01 << 15 //+ Bus busy This flag indicates that a communication is in progress on the bus. It is set by hardware when a START condition is detected. It is cleared by hardware when a Stop condition is detected, or when PE=0.
	DIR     ISR = 0x01 << 16 //+ Transfer direction (Slave mode) This flag is updated when an address match event occurs (ADDR=1).
	ADDCODE ISR = 0x7F << 17 //+ Address match code (Slave mode) These bits are updated with the received address when an address match event occurs (ADDR = 1). In the case of a 10-bit address, ADDCODE provides the 10-bit header followed by the 2 MSBs of the address.
)

const (
	TXEn     = 0
	TXISn    = 1
	RXNEn    = 2
	ADDRn    = 3
	NACKFn   = 4
	STOPFn   = 5
	TCn      = 6
	TCRn     = 7
	BERRn    = 8
	ARLOn    = 9
	OVRn     = 10
	PECERRn  = 11
	TIMEOUTn = 12
	ALERTn   = 13
	BUSYn    = 15
	DIRn     = 16
	ADDCODEn = 17
)

const (
	ADDRCF   ICR = 0x01 << 3  //+ Address matched flag clear Writing 1 to this bit clears the ADDR flag in the I2C_ISR register. Writing 1 to this bit also clears the START bit in the I2C_CR2 register.
	NACKCF   ICR = 0x01 << 4  //+ Not Acknowledge flag clear Writing 1 to this bit clears the ACKF flag in I2C_ISR register.
	STOPCF   ICR = 0x01 << 5  //+ Stop detection flag clear Writing 1 to this bit clears the STOPF flag in the I2C_ISR register.
	BERRCF   ICR = 0x01 << 8  //+ Bus error flag clear Writing 1 to this bit clears the BERRF flag in the I2C_ISR register.
	ARLOCF   ICR = 0x01 << 9  //+ Arbitration Lost flag clear Writing 1 to this bit clears the ARLO flag in the I2C_ISR register.
	OVRCF    ICR = 0x01 << 10 //+ Overrun/Underrun flag clear Writing 1 to this bit clears the OVR flag in the I2C_ISR register.
	PECCF    ICR = 0x01 << 11 //+ PEC Error flag clear Writing 1 to this bit clears the PECERR flag in the I2C_ISR register. Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
	TIMOUTCF ICR = 0x01 << 12 //+ Timeout detection flag clear Writing 1 to this bit clears the TIMEOUT flag in the I2C_ISR register. Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
	ALERTCF  ICR = 0x01 << 13 //+ Alert flag clear Writing 1 to this bit clears the ALERT flag in the I2C_ISR register. Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
)

const (
	ADDRCFn   = 3
	NACKCFn   = 4
	STOPCFn   = 5
	BERRCFn   = 8
	ARLOCFn   = 9
	OVRCFn    = 10
	PECCFn    = 11
	TIMOUTCFn = 12
	ALERTCFn  = 13
)

const (
	PEC PECR = 0xFF << 0 //+ Packet error checking register This field contains the internal PEC when PECEN=1. The PEC is cleared by hardware when PE=0.
)

const (
	PECn = 0
)

const (
	RXDATA RXDR = 0xFF << 0 //+ 8-bit receive data Data byte received from the I2C bus.
)

const (
	RXDATAn = 0
)

const (
	TXDATA TXDR = 0xFF << 0 //+ 8-bit transmit data Data byte to be transmitted to the I2C bus. Note: These bits can be written only when TXE=1.
)

const (
	TXDATAn = 0
)
