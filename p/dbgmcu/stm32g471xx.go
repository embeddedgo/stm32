// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32g471xx

// Package dbgmcu provides access to the registers of the DBGMCU peripheral.
//
// Instances:
//
//	DBGMCU  DBGMCU_BASE  -  -  Debug support
//
// Registers:
//
//	0x000 32  IDCODE    MCU Device ID Code Register
//	0x004 32  CR        Debug MCU Configuration Register
//	0x008 32  APB1L_FZ  APB Low Freeze Register 1
//	0x00C 32  APB1H_FZ  APB Low Freeze Register 2
//	0x010 32  APB2_FZ   APB High Freeze Register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package dbgmcu

const (
	DEV_ID IDCODE = 0xFFFF << 0  //+ Device Identifier
	REV_ID IDCODE = 0xFFFF << 16 //+ Revision Identifier
)

const (
	DEV_IDn = 0
	REV_IDn = 16
)

const (
	DBG_SLEEP   CR = 0x01 << 0 //+ Debug Sleep Mode
	DBG_STOP    CR = 0x01 << 1 //+ Debug Stop Mode
	DBG_STANDBY CR = 0x01 << 2 //+ Debug Standby Mode
	TRACE_IOEN  CR = 0x01 << 5 //+ Trace pin assignment control
	TRACE_MODE  CR = 0x03 << 6 //+ Trace pin assignment control
)

const (
	DBG_SLEEPn   = 0
	DBG_STOPn    = 1
	DBG_STANDBYn = 2
	TRACE_IOENn  = 5
	TRACE_MODEn  = 6
)

const (
	DBG_TIMER2_STOP  APB1L_FZ = 0x01 << 0  //+ Debug Timer 2 stopped when Core is halted
	DBG_TIM3_STOP    APB1L_FZ = 0x01 << 1  //+ TIM3 counter stopped when core is halted
	DBG_TIM4_STOP    APB1L_FZ = 0x01 << 2  //+ TIM4 counter stopped when core is halted
	DBG_TIM5_STOP    APB1L_FZ = 0x01 << 3  //+ TIM5 counter stopped when core is halted
	DBG_TIMER6_STOP  APB1L_FZ = 0x01 << 4  //+ Debug Timer 6 stopped when Core is halted
	DBG_TIM7_STOP    APB1L_FZ = 0x01 << 5  //+ TIM7 counter stopped when core is halted
	DBG_RTC_STOP     APB1L_FZ = 0x01 << 10 //+ Debug RTC stopped when Core is halted
	DBG_WWDG_STOP    APB1L_FZ = 0x01 << 11 //+ Debug Window Wachdog stopped when Core is halted
	DBG_IWDG_STOP    APB1L_FZ = 0x01 << 12 //+ Debug Independent Wachdog stopped when Core is halted
	DBG_I2C1_STOP    APB1L_FZ = 0x01 << 21 //+ I2C1 SMBUS timeout mode stopped when core is halted
	DBG_I2C2_STOP    APB1L_FZ = 0x01 << 22 //+ I2C2 SMBUS timeout mode stopped when core is halted
	DBG_I2C3_STOP    APB1L_FZ = 0x01 << 30 //+ I2C3 SMBUS timeout mode stopped when core is halted
	DBG_LPTIMER_STOP APB1L_FZ = 0x01 << 31 //+ LPTIM1 counter stopped when core is halted
)

const (
	DBG_TIMER2_STOPn  = 0
	DBG_TIM3_STOPn    = 1
	DBG_TIM4_STOPn    = 2
	DBG_TIM5_STOPn    = 3
	DBG_TIMER6_STOPn  = 4
	DBG_TIM7_STOPn    = 5
	DBG_RTC_STOPn     = 10
	DBG_WWDG_STOPn    = 11
	DBG_IWDG_STOPn    = 12
	DBG_I2C1_STOPn    = 21
	DBG_I2C2_STOPn    = 22
	DBG_I2C3_STOPn    = 30
	DBG_LPTIMER_STOPn = 31
)

const (
	DBG_I2C4_STOP APB1H_FZ = 0x01 << 1 //+ DBG_I2C4_STOP
)

const (
	DBG_I2C4_STOPn = 1
)

const (
	DBG_TIM1_STOP   APB2_FZ = 0x01 << 11 //+ TIM1 counter stopped when core is halted
	DBG_TIM8_STOP   APB2_FZ = 0x01 << 13 //+ TIM8 counter stopped when core is halted
	DBG_TIM15_STOP  APB2_FZ = 0x01 << 16 //+ TIM15 counter stopped when core is halted
	DBG_TIM16_STOP  APB2_FZ = 0x01 << 17 //+ TIM16 counter stopped when core is halted
	DBG_TIM17_STOP  APB2_FZ = 0x01 << 18 //+ TIM17 counter stopped when core is halted
	DBG_TIM20_STOP  APB2_FZ = 0x01 << 20 //+ TIM20counter stopped when core is halted
	DBG_HRTIM0_STOP APB2_FZ = 0x01 << 26 //+ DBG_HRTIM0_STOP
	DBG_HRTIM1_STOP APB2_FZ = 0x01 << 27 //+ DBG_HRTIM0_STOP
	DBG_HRTIM2_STOP APB2_FZ = 0x01 << 28 //+ DBG_HRTIM0_STOP
	DBG_HRTIM3_STOP APB2_FZ = 0x01 << 29 //+ DBG_HRTIM0_STOP
)

const (
	DBG_TIM1_STOPn   = 11
	DBG_TIM8_STOPn   = 13
	DBG_TIM15_STOPn  = 16
	DBG_TIM16_STOPn  = 17
	DBG_TIM17_STOPn  = 18
	DBG_TIM20_STOPn  = 20
	DBG_HRTIM0_STOPn = 26
	DBG_HRTIM1_STOPn = 27
	DBG_HRTIM2_STOPn = 28
	DBG_HRTIM3_STOPn = 29
)
