// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32h7x3

// Package iwdg provides access to the registers of the IWDG peripheral.
//
// Instances:
//  IWDG  IWDG_BASE  -  -  IWDG
// Registers:
//  0x000 32  KR    Key register
//  0x004 32  PR    Prescaler register
//  0x008 32  RLR   Reload register
//  0x00C 32  SR    Status register
//  0x010 32  WINR  Window register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package iwdg

const (
	KEY KR = 0xFFFF << 0 //+ Key value (write only, read 0x0000) These bits must be written by software at regular intervals with the key value 0xAAAA, otherwise the watchdog generates a reset when the counter reaches 0. Writing the key value 0x5555 to enable access to the IWDG_PR, IWDG_RLR and IWDG_WINR registers (see Section23.3.6: Register access protection) Writing the key value CCCCh starts the watchdog (except if the hardware watchdog option is selected)
)

const (
	KEYn = 0
)

const (
	PR PR = 0x07 << 0 //+ Prescaler divider These bits are write access protected see Section23.3.6: Register access protection. They are written by software to select the prescaler divider feeding the counter clock. PVU bit of IWDG_SR must be reset in order to be able to change the prescaler divider. Note: Reading this register returns the prescaler value from the VDD voltage domain. This value may not be up to date/valid if a write operation to this register is ongoing. For this reason the value read from this register is valid only when the PVU bit in the IWDG_SR register is reset.
)

const (
	PRn = 0
)

const (
	RL RLR = 0xFFF << 0 //+ Watchdog counter reload value These bits are write access protected see Section23.3.6. They are written by software to define the value to be loaded in the watchdog counter each time the value 0xAAAA is written in the IWDG_KR register. The watchdog counter counts down from this value. The timeout period is a function of this value and the clock prescaler. Refer to the datasheet for the timeout information. The RVU bit in the IWDG_SR register must be reset in order to be able to change the reload value. Note: Reading this register returns the reload value from the VDD voltage domain. This value may not be up to date/valid if a write operation to this register is ongoing on this register. For this reason the value read from this register is valid only when the RVU bit in the IWDG_SR register is reset.
)

const (
	RLn = 0
)

const (
	PVU SR = 0x01 << 0 //+ Watchdog prescaler value update This bit is set by hardware to indicate that an update of the prescaler value is ongoing. It is reset by hardware when the prescaler update operation is completed in the VDD voltage domain (takes up to 5 RC 40 kHz cycles). Prescaler value can be updated only when PVU bit is reset.
	RVU SR = 0x01 << 1 //+ Watchdog counter reload value update This bit is set by hardware to indicate that an update of the reload value is ongoing. It is reset by hardware when the reload value update operation is completed in the VDD voltage domain (takes up to 5 RC 40 kHz cycles). Reload value can be updated only when RVU bit is reset.
	WVU SR = 0x01 << 2 //+ Watchdog counter window value update This bit is set by hardware to indicate that an update of the window value is ongoing. It is reset by hardware when the reload value update operation is completed in the VDD voltage domain (takes up to 5 RC 40 kHz cycles). Window value can be updated only when WVU bit is reset. This bit is generated only if generic window = 1
)

const (
	PVUn = 0
	RVUn = 1
	WVUn = 2
)

const (
	WIN WINR = 0xFFF << 0 //+ Watchdog counter window value These bits are write access protected see Section23.3.6. These bits contain the high limit of the window value to be compared to the downcounter. To prevent a reset, the downcounter must be reloaded when its value is lower than the window register value and greater than 0x0 The WVU bit in the IWDG_SR register must be reset in order to be able to change the reload value. Note: Reading this register returns the reload value from the VDD voltage domain. This value may not be valid if a write operation to this register is ongoing. For this reason the value read from this register is valid only when the WVU bit in the IWDG_SR register is reset.
)

const (
	WINn = 0
)