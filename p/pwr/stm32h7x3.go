// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package pwr provides access to the registers of the PWR peripheral.
//
// Instances:
//
//	PWR  PWR_BASE  -  -
//
// Registers:
//
//	0x000 32  CR1      PWR control register 1
//	0x004 32  CSR1     PWR control status register 1
//	0x008 32  CR2      This register is not reset by wakeup from Standby mode, RESET signal and VDD POR. It is only reset by VSW POR and VSWRST reset. This register shall not be accessed when VSWRST bit in RCC_BDCR register resets the VSW domain.After reset, PWR_CR2 register is write-protected. Prior to modifying its content, the DBP bit in PWR_CR1 register must be set to disable the write protection.
//	0x00C 32  CR3      Reset only by POR only, not reset by wakeup from Standby mode and RESET pad. The lower byte of this register is written once after POR and shall be written before changing VOS level or ck_sys clock frequency. No limitation applies to the upper bytes.Programming data corresponding to an invalid combination of SDLEVEL, SDEXTHP, SDEN, LDOEN and BYPASS bits (see Table9) will be ignored: data will not be written, the written-once mechanism will lock the register and any further write access will be ignored. The default supply configuration will be kept and the ACTVOSRDY bit in PWR control status register 1 (PWR_CSR1) will go on indicating invalid voltage levels. The system shall be power cycled before writing a new value.
//	0x010 32  CPUCR    This register allows controlling CPU1 power.
//	0x018 32  D3CR     This register allows controlling D3 domain power.Following reset VOSRDY will be read 1 by software
//	0x020 32  WKUPCR   reset only by system reset, not reset by wakeup from Standby mode5 wait states are required when writing this register (when clearing a WKUPF bit in PWR_WKUPFR, the AHB write access will complete after the WKUPF has been cleared).
//	0x024 32  WKUPFR   reset only by system reset, not reset by wakeup from Standby mode
//	0x028 32  WKUPEPR  Reset only by system reset, not reset by wakeup from Standby mode
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package pwr

const (
	LPDS  CR1 = 0x01 << 0  //+ Low-power Deepsleep with SVOS3 (SVOS4 and SVOS5 always use low-power, regardless of the setting of this bit)
	PVDE  CR1 = 0x01 << 4  //+ Programmable voltage detector enable
	PLS   CR1 = 0x07 << 5  //+ Programmable voltage detector level selection These bits select the voltage threshold detected by the PVD. Note: Refer to Section Electrical characteristics of the product datasheet for more details.
	DBP   CR1 = 0x01 << 8  //+ Disable backup domain write protection In reset state, the RCC_BDCR register, the RTC registers (including the backup registers), BREN and MOEN bits in PWR_CR2 register, are protected against parasitic write access. This bit must be set to enable write access to these registers.
	FLPS  CR1 = 0x01 << 9  //+ Flash low-power mode in DStop mode This bit allows to obtain the best trade-off between low-power consumption and restart time when exiting from DStop mode. When it is set, the Flash memory enters low-power mode when D1 domain is in DStop mode.
	SVOS  CR1 = 0x03 << 14 //+ System Stop mode voltage scaling selection These bits control the VCORE voltage level in system Stop mode, to obtain the best trade-off between power consumption and performance.
	AVDEN CR1 = 0x01 << 16 //+ Peripheral voltage monitor on VDDA enable
	ALS   CR1 = 0x03 << 17 //+ Analog voltage detector level selection These bits select the voltage threshold detected by the AVD.
)

const (
	LPDSn  = 0
	PVDEn  = 4
	PLSn   = 5
	DBPn   = 8
	FLPSn  = 9
	SVOSn  = 14
	AVDENn = 16
	ALSn   = 17
)

const (
	PVDO      CSR1 = 0x01 << 4  //+ Programmable voltage detect output This bit is set and cleared by hardware. It is valid only if the PVD has been enabled by the PVDE bit. Note: since the PVD is disabled in Standby mode, this bit is equal to 0 after Standby or reset until the PVDE bit is set.
	ACTVOSRDY CSR1 = 0x01 << 13 //+ Voltage levels ready bit for currently used VOS and SDLEVEL This bit is set to 1 by hardware when the voltage regulator and the SD converter are both disabled and Bypass mode is selected in PWR control register 3 (PWR_CR3).
	ACTVOS    CSR1 = 0x03 << 14 //+ VOS currently applied for VCORE voltage scaling selection. These bits reflect the last VOS value applied to the PMU.
	AVDO      CSR1 = 0x01 << 16 //+ Analog voltage detector output on VDDA This bit is set and cleared by hardware. It is valid only if AVD on VDDA is enabled by the AVDEN bit. Note: Since the AVD is disabled in Standby mode, this bit is equal to 0 after Standby or reset until the AVDEN bit is set.
)

const (
	PVDOn      = 4
	ACTVOSRDYn = 13
	ACTVOSn    = 14
	AVDOn      = 16
)

const (
	BREN  CR2 = 0x01 << 0  //+ Backup regulator enable When set, the Backup regulator (used to maintain the backup RAM content in Standby and VBAT modes) is enabled. If BREN is reset, the backup regulator is switched off. The backup RAM can still be used in Run and Stop modes. However, its content will be lost in Standby and VBAT modes. If BREN is set, the application must wait till the Backup Regulator Ready flag (BRRDY) is set to indicate that the data written into the SRAM will be maintained in Standby and VBAT modes.
	MONEN CR2 = 0x01 << 4  //+ VBAT and temperature monitoring enable When set, the VBAT supply and temperature monitoring is enabled.
	BRRDY CR2 = 0x01 << 16 //+ Backup regulator ready This bit is set by hardware to indicate that the Backup regulator is ready.
	VBATL CR2 = 0x01 << 20 //+ VBAT level monitoring versus low threshold
	VBATH CR2 = 0x01 << 21 //+ VBAT level monitoring versus high threshold
	TEMPL CR2 = 0x01 << 22 //+ Temperature level monitoring versus low threshold
	TEMPH CR2 = 0x01 << 23 //+ Temperature level monitoring versus high threshold
)

const (
	BRENn  = 0
	MONENn = 4
	BRRDYn = 16
	VBATLn = 20
	VBATHn = 21
	TEMPLn = 22
	TEMPHn = 23
)

const (
	BYPASS   CR3 = 0x01 << 0  //+ Power management unit bypass
	LDOEN    CR3 = 0x01 << 1  //+ Low drop-out regulator enable
	SCUEN    CR3 = 0x01 << 2  //+ SD converter Enable
	VBE      CR3 = 0x01 << 8  //+ VBAT charging enable
	VBRS     CR3 = 0x01 << 9  //+ VBAT charging resistor selection
	USB33DEN CR3 = 0x01 << 24 //+ VDD33USB voltage level detector enable.
	USBREGEN CR3 = 0x01 << 25 //+ USB regulator enable.
	USB33RDY CR3 = 0x01 << 26 //+ USB supply ready.
)

const (
	BYPASSn   = 0
	LDOENn    = 1
	SCUENn    = 2
	VBEn      = 8
	VBRSn     = 9
	USB33DENn = 24
	USBREGENn = 25
	USB33RDYn = 26
)

const (
	PDDS_D1 CPUCR = 0x01 << 0  //+ D1 domain Power Down Deepsleep selection. This bit allows CPU1 to define the Deepsleep mode for D1 domain.
	PDDS_D2 CPUCR = 0x01 << 1  //+ D2 domain Power Down Deepsleep. This bit allows CPU1 to define the Deepsleep mode for D2 domain.
	PDDS_D3 CPUCR = 0x01 << 2  //+ System D3 domain Power Down Deepsleep. This bit allows CPU1 to define the Deepsleep mode for System D3 domain.
	STOPF   CPUCR = 0x01 << 5  //+ STOP flag This bit is set by hardware and cleared only by any reset or by setting the CPU1 CSSF bit.
	SBF     CPUCR = 0x01 << 6  //+ System Standby flag This bit is set by hardware and cleared only by a POR (Power-on Reset) or by setting the CPU1 CSSF bit
	SBF_D1  CPUCR = 0x01 << 7  //+ D1 domain DStandby flag This bit is set by hardware and cleared by any system reset or by setting the CPU1 CSSF bit. Once set, this bit can be cleared only when the D1 domain is no longer in DStandby mode.
	SBF_D2  CPUCR = 0x01 << 8  //+ D2 domain DStandby flag This bit is set by hardware and cleared by any system reset or by setting the CPU1 CSSF bit. Once set, this bit can be cleared only when the D2 domain is no longer in DStandby mode.
	CSSF    CPUCR = 0x01 << 9  //+ Clear D1 domain CPU1 Standby, Stop and HOLD flags (always read as 0) This bit is cleared to 0 by hardware.
	RUN_D3  CPUCR = 0x01 << 11 //+ Keep system D3 domain in Run mode regardless of the CPU sub-systems modes
)

const (
	PDDS_D1n = 0
	PDDS_D2n = 1
	PDDS_D3n = 2
	STOPFn   = 5
	SBFn     = 6
	SBF_D1n  = 7
	SBF_D2n  = 8
	CSSFn    = 9
	RUN_D3n  = 11
)

const (
	VOSRDY D3CR = 0x01 << 13 //+ VOS Ready bit for VCORE voltage scaling output selection. This bit is set to 1 by hardware when Bypass mode is selected in PWR control register 3 (PWR_CR3).
	VOS    D3CR = 0x03 << 14 //+ Voltage scaling selection according to performance These bits control the VCORE voltage level and allow to obtains the best trade-off between power consumption and performance: When increasing the performance, the voltage scaling shall be changed before increasing the system frequency. When decreasing performance, the system frequency shall first be decreased before changing the voltage scaling.
)

const (
	VOSRDYn = 13
	VOSn    = 14
)

const (
	WKUPC WKUPCR = 0x3F << 0 //+ Clear Wakeup pin flag for WKUP. These bits are always read as 0.
)

const (
	WKUPCn = 0
)

const (
	WKUPF1 WKUPFR = 0x01 << 0 //+ Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
	WKUPF2 WKUPFR = 0x01 << 1 //+ Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
	WKUPF3 WKUPFR = 0x01 << 2 //+ Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
	WKUPF4 WKUPFR = 0x01 << 3 //+ Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
	WKUPF5 WKUPFR = 0x01 << 4 //+ Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
	WKUPF6 WKUPFR = 0x01 << 5 //+ Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
)

const (
	WKUPF1n = 0
	WKUPF2n = 1
	WKUPF3n = 2
	WKUPF4n = 3
	WKUPF5n = 4
	WKUPF6n = 5
)

const (
	WKUPEN1   WKUPEPR = 0x01 << 0  //+ Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
	WKUPEN2   WKUPEPR = 0x01 << 1  //+ Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
	WKUPEN3   WKUPEPR = 0x01 << 2  //+ Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
	WKUPEN4   WKUPEPR = 0x01 << 3  //+ Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
	WKUPEN5   WKUPEPR = 0x01 << 4  //+ Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
	WKUPEN6   WKUPEPR = 0x01 << 5  //+ Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
	WKUPP1    WKUPEPR = 0x01 << 8  //+ Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
	WKUPP2    WKUPEPR = 0x01 << 9  //+ Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
	WKUPP3    WKUPEPR = 0x01 << 10 //+ Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
	WKUPP4    WKUPEPR = 0x01 << 11 //+ Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
	WKUPP5    WKUPEPR = 0x01 << 12 //+ Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
	WKUPP6    WKUPEPR = 0x01 << 13 //+ Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
	WKUPPUPD1 WKUPEPR = 0x03 << 16 //+ Wakeup pin pull configuration
	WKUPPUPD2 WKUPEPR = 0x03 << 18 //+ Wakeup pin pull configuration
	WKUPPUPD3 WKUPEPR = 0x03 << 20 //+ Wakeup pin pull configuration
	WKUPPUPD4 WKUPEPR = 0x03 << 22 //+ Wakeup pin pull configuration
	WKUPPUPD5 WKUPEPR = 0x03 << 24 //+ Wakeup pin pull configuration
	WKUPPUPD6 WKUPEPR = 0x03 << 26 //+ Wakeup pin pull configuration for WKUP(truncate(n/2)-7) These bits define the I/O pad pull configuration used when WKUPEN(truncate(n/2)-7) = 1. The associated GPIO port pull configuration shall be set to the same value or to 00. The Wakeup pin pull configuration is kept in Standby mode.
)

const (
	WKUPEN1n   = 0
	WKUPEN2n   = 1
	WKUPEN3n   = 2
	WKUPEN4n   = 3
	WKUPEN5n   = 4
	WKUPEN6n   = 5
	WKUPP1n    = 8
	WKUPP2n    = 9
	WKUPP3n    = 10
	WKUPP4n    = 11
	WKUPP5n    = 12
	WKUPP6n    = 13
	WKUPPUPD1n = 16
	WKUPPUPD2n = 18
	WKUPPUPD3n = 20
	WKUPPUPD4n = 22
	WKUPPUPD5n = 24
	WKUPPUPD6n = 26
)
