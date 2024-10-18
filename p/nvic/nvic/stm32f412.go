// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f412

// Package nvic provides access to the registers of the NVIC peripheral.
//
// Instances:
//
//	NVIC  NVIC_BASE  -  -  Nested Vectored Interrupt Controller
//
// Registers:
//
//	0x000 32  ISER0  Interrupt Set-Enable Register
//	0x004 32  ISER1  Interrupt Set-Enable Register
//	0x008 32  ISER2  Interrupt Set-Enable Register
//	0x080 32  ICER0  Interrupt Clear-Enable Register
//	0x084 32  ICER1  Interrupt Clear-Enable Register
//	0x088 32  ICER2  Interrupt Clear-Enable Register
//	0x100 32  ISPR0  Interrupt Set-Pending Register
//	0x104 32  ISPR1  Interrupt Set-Pending Register
//	0x108 32  ISPR2  Interrupt Set-Pending Register
//	0x180 32  ICPR0  Interrupt Clear-Pending Register
//	0x184 32  ICPR1  Interrupt Clear-Pending Register
//	0x188 32  ICPR2  Interrupt Clear-Pending Register
//	0x200 32  IABR0  Interrupt Active Bit Register
//	0x204 32  IABR1  Interrupt Active Bit Register
//	0x208 32  IABR2  Interrupt Active Bit Register
//	0x300 32  IPR0   Interrupt Priority Register
//	0x304 32  IPR1   Interrupt Priority Register
//	0x308 32  IPR2   Interrupt Priority Register
//	0x30C 32  IPR3   Interrupt Priority Register
//	0x310 32  IPR4   Interrupt Priority Register
//	0x314 32  IPR5   Interrupt Priority Register
//	0x318 32  IPR6   Interrupt Priority Register
//	0x31C 32  IPR7   Interrupt Priority Register
//	0x320 32  IPR8   Interrupt Priority Register
//	0x324 32  IPR9   Interrupt Priority Register
//	0x328 32  IPR10  Interrupt Priority Register
//	0x32C 32  IPR11  Interrupt Priority Register
//	0x330 32  IPR12  Interrupt Priority Register
//	0x334 32  IPR13  Interrupt Priority Register
//	0x338 32  IPR14  Interrupt Priority Register
//	0x33C 32  IPR15  Interrupt Priority Register
//	0x340 32  IPR16  Interrupt Priority Register
//	0x344 32  IPR17  Interrupt Priority Register
//	0x348 32  IPR18  Interrupt Priority Register
//	0x34C 32  IPR19  Interrupt Priority Register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package nvic

const (
	SETENA ISER0 = 0xFFFFFFFF << 0 //+ SETENA
)

const (
	SETENAn = 0
)

const (
	SETENA ISER1 = 0xFFFFFFFF << 0 //+ SETENA
)

const (
	SETENAn = 0
)

const (
	SETENA ISER2 = 0xFFFFFFFF << 0 //+ SETENA
)

const (
	SETENAn = 0
)

const (
	CLRENA ICER0 = 0xFFFFFFFF << 0 //+ CLRENA
)

const (
	CLRENAn = 0
)

const (
	CLRENA ICER1 = 0xFFFFFFFF << 0 //+ CLRENA
)

const (
	CLRENAn = 0
)

const (
	CLRENA ICER2 = 0xFFFFFFFF << 0 //+ CLRENA
)

const (
	CLRENAn = 0
)

const (
	SETPEND ISPR0 = 0xFFFFFFFF << 0 //+ SETPEND
)

const (
	SETPENDn = 0
)

const (
	SETPEND ISPR1 = 0xFFFFFFFF << 0 //+ SETPEND
)

const (
	SETPENDn = 0
)

const (
	SETPEND ISPR2 = 0xFFFFFFFF << 0 //+ SETPEND
)

const (
	SETPENDn = 0
)

const (
	CLRPEND ICPR0 = 0xFFFFFFFF << 0 //+ CLRPEND
)

const (
	CLRPENDn = 0
)

const (
	CLRPEND ICPR1 = 0xFFFFFFFF << 0 //+ CLRPEND
)

const (
	CLRPENDn = 0
)

const (
	CLRPEND ICPR2 = 0xFFFFFFFF << 0 //+ CLRPEND
)

const (
	CLRPENDn = 0
)

const (
	ACTIVE IABR0 = 0xFFFFFFFF << 0 //+ ACTIVE
)

const (
	ACTIVEn = 0
)

const (
	ACTIVE IABR1 = 0xFFFFFFFF << 0 //+ ACTIVE
)

const (
	ACTIVEn = 0
)

const (
	ACTIVE IABR2 = 0xFFFFFFFF << 0 //+ ACTIVE
)

const (
	ACTIVEn = 0
)

const (
	IPR_N0 IPR0 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR0 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR0 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR0 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR1 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR1 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR1 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR1 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR2 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR2 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR2 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR2 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR3 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR3 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR3 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR3 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR4 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR4 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR4 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR4 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR5 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR5 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR5 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR5 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR6 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR6 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR6 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR6 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR7 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR7 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR7 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR7 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR8 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR8 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR8 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR8 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR9 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR9 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR9 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR9 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR10 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR10 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR10 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR10 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR11 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR11 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR11 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR11 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR12 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR12 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR12 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR12 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR13 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR13 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR13 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR13 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR14 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR14 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR14 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR14 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR15 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR15 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR15 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR15 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR16 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR16 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR16 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR16 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR17 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR17 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR17 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR17 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR18 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR18 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR18 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR18 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)

const (
	IPR_N0 IPR19 = 0xFF << 0  //+ IPR_N0
	IPR_N1 IPR19 = 0xFF << 8  //+ IPR_N1
	IPR_N2 IPR19 = 0xFF << 16 //+ IPR_N2
	IPR_N3 IPR19 = 0xFF << 24 //+ IPR_N3
)

const (
	IPR_N0n = 0
	IPR_N1n = 8
	IPR_N2n = 16
	IPR_N3n = 24
)
