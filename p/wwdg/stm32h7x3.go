// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package wwdg provides access to the registers of the WWDG peripheral.
//
// Instances:
//
//	WWDG  WWDG_BASE  -  WWDG1,WWDG1_RST
//
// Registers:
//
//	0x000 32  CR   Control register
//	0x004 32  CFR  Configuration register
//	0x008 32  SR   Status register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package wwdg

const (
	T    CR = 0x7F << 0 //+ 7-bit counter (MSB to LSB) These bits contain the value of the watchdog counter. It is decremented every (4096 x 2WDGTB[1:0]) PCLK cycles. A reset is produced when it is decremented from 0x40 to 0x3F (T6 becomes cleared).
	WDGA CR = 0x01 << 7 //+ Activation bit This bit is set by software and only cleared by hardware after a reset. When WDGA=1, the watchdog can generate a reset.
)

const (
	Tn    = 0
	WDGAn = 7
)

const (
	W     CFR = 0x7F << 0  //+ 7-bit window value These bits contain the window value to be compared to the downcounter.
	EWI   CFR = 0x01 << 9  //+ Early wakeup interrupt When set, an interrupt occurs whenever the counter reaches the value 0x40. This interrupt is only cleared by hardware after a reset.
	WDGTB CFR = 0x03 << 11 //+ Timer base The time base of the prescaler can be modified as follows:
)

const (
	Wn     = 0
	EWIn   = 9
	WDGTBn = 11
)

const (
	EWIF SR = 0x01 << 0 //+ Early wakeup interrupt flag This bit is set by hardware when the counter has reached the value 0x40. It must be cleared by software by writing 0. A write of 1 has no effect. This bit is also set if the interrupt is not enabled.
)

const (
	EWIFn = 0
)
