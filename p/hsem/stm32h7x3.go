// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package hsem provides access to the registers of the HSEM peripheral.
//
// Instances:
//
//	HSEM  HSEM_BASE  AHB4  HSEM0
//
// Registers:
//
//	0x000 32  HSEM_R0     HSEM register HSEM_R0 HSEM_R31
//	0x004 32  HSEM_R1     HSEM register HSEM_R0 HSEM_R31
//	0x008 32  HSEM_R2     HSEM register HSEM_R0 HSEM_R31
//	0x00C 32  HSEM_R3     HSEM register HSEM_R0 HSEM_R31
//	0x010 32  HSEM_R4     HSEM register HSEM_R0 HSEM_R31
//	0x014 32  HSEM_R5     HSEM register HSEM_R0 HSEM_R31
//	0x018 32  HSEM_R6     HSEM register HSEM_R0 HSEM_R31
//	0x01C 32  HSEM_R7     HSEM register HSEM_R0 HSEM_R31
//	0x020 32  HSEM_R8     HSEM register HSEM_R0 HSEM_R31
//	0x024 32  HSEM_R9     HSEM register HSEM_R0 HSEM_R31
//	0x028 32  HSEM_R10    HSEM register HSEM_R0 HSEM_R31
//	0x02C 32  HSEM_R11    HSEM register HSEM_R0 HSEM_R31
//	0x030 32  HSEM_R12    HSEM register HSEM_R0 HSEM_R31
//	0x034 32  HSEM_R13    HSEM register HSEM_R0 HSEM_R31
//	0x038 32  HSEM_R14    HSEM register HSEM_R0 HSEM_R31
//	0x03C 32  HSEM_R15    HSEM register HSEM_R0 HSEM_R31
//	0x040 32  HSEM_R16    HSEM register HSEM_R0 HSEM_R31
//	0x044 32  HSEM_R17    HSEM register HSEM_R0 HSEM_R31
//	0x048 32  HSEM_R18    HSEM register HSEM_R0 HSEM_R31
//	0x04C 32  HSEM_R19    HSEM register HSEM_R0 HSEM_R31
//	0x050 32  HSEM_R20    HSEM register HSEM_R0 HSEM_R31
//	0x054 32  HSEM_R21    HSEM register HSEM_R0 HSEM_R31
//	0x058 32  HSEM_R22    HSEM register HSEM_R0 HSEM_R31
//	0x05C 32  HSEM_R23    HSEM register HSEM_R0 HSEM_R31
//	0x060 32  HSEM_R24    HSEM register HSEM_R0 HSEM_R31
//	0x064 32  HSEM_R25    HSEM register HSEM_R0 HSEM_R31
//	0x068 32  HSEM_R26    HSEM register HSEM_R0 HSEM_R31
//	0x06C 32  HSEM_R27    HSEM register HSEM_R0 HSEM_R31
//	0x070 32  HSEM_R28    HSEM register HSEM_R0 HSEM_R31
//	0x074 32  HSEM_R29    HSEM register HSEM_R0 HSEM_R31
//	0x078 32  HSEM_R30    HSEM register HSEM_R0 HSEM_R31
//	0x07C 32  HSEM_R31    HSEM register HSEM_R0 HSEM_R31
//	0x080 32  HSEM_RLR0   HSEM Read lock register
//	0x084 32  HSEM_RLR1   HSEM Read lock register
//	0x088 32  HSEM_RLR2   HSEM Read lock register
//	0x08C 32  HSEM_RLR3   HSEM Read lock register
//	0x090 32  HSEM_RLR4   HSEM Read lock register
//	0x094 32  HSEM_RLR5   HSEM Read lock register
//	0x098 32  HSEM_RLR6   HSEM Read lock register
//	0x09C 32  HSEM_RLR7   HSEM Read lock register
//	0x0A0 32  HSEM_RLR8   HSEM Read lock register
//	0x0A4 32  HSEM_RLR9   HSEM Read lock register
//	0x0A8 32  HSEM_RLR10  HSEM Read lock register
//	0x0AC 32  HSEM_RLR11  HSEM Read lock register
//	0x0B0 32  HSEM_RLR12  HSEM Read lock register
//	0x0B4 32  HSEM_RLR13  HSEM Read lock register
//	0x0B8 32  HSEM_RLR14  HSEM Read lock register
//	0x0BC 32  HSEM_RLR15  HSEM Read lock register
//	0x0C0 32  HSEM_RLR16  HSEM Read lock register
//	0x0C4 32  HSEM_RLR17  HSEM Read lock register
//	0x0C8 32  HSEM_RLR18  HSEM Read lock register
//	0x0CC 32  HSEM_RLR19  HSEM Read lock register
//	0x0D0 32  HSEM_RLR20  HSEM Read lock register
//	0x0D4 32  HSEM_RLR21  HSEM Read lock register
//	0x0D8 32  HSEM_RLR22  HSEM Read lock register
//	0x0DC 32  HSEM_RLR23  HSEM Read lock register
//	0x0E0 32  HSEM_RLR24  HSEM Read lock register
//	0x0E4 32  HSEM_RLR25  HSEM Read lock register
//	0x0E8 32  HSEM_RLR26  HSEM Read lock register
//	0x0EC 32  HSEM_RLR27  HSEM Read lock register
//	0x0F0 32  HSEM_RLR28  HSEM Read lock register
//	0x0F4 32  HSEM_RLR29  HSEM Read lock register
//	0x0F8 32  HSEM_RLR30  HSEM Read lock register
//	0x0FC 32  HSEM_RLR31  HSEM Read lock register
//	0x100 32  HSEM_IER    HSEM Interrupt enable register
//	0x104 32  HSEM_ICR    HSEM Interrupt clear register
//	0x108 32  HSEM_ISR    HSEM Interrupt status register
//	0x10C 32  HSEM_MISR   HSEM Masked interrupt status register
//	0x140 32  HSEM_CR     HSEM Clear register
//	0x144 32  HSEM_KEYR   HSEM Interrupt clear register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package hsem

const (
	PROCID   HSEM_R0 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R0 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R0 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R1 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R1 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R1 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R2 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R2 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R2 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R3 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R3 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R3 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R4 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R4 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R4 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R5 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R5 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R5 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R6 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R6 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R6 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R7 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R7 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R7 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R8 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R8 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R8 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R9 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R9 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R9 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R10 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R10 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R10 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R11 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R11 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R11 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R12 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R12 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R12 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R13 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R13 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R13 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R14 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R14 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R14 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R15 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R15 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R15 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R16 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R16 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R16 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R17 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R17 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R17 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R18 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R18 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R18 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R19 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R19 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R19 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R20 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R20 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R20 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R21 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R21 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R21 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R22 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R22 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R22 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R23 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R23 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R23 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R24 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R24 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R24 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R25 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R25 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R25 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R26 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R26 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R26 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R27 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R27 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R27 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R28 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R28 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R28 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R29 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R29 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R29 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R30 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R30 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R30 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_R31 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_R31 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_R31 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR0 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR0 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR0 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR1 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR1 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR1 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR2 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR2 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR2 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR3 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR3 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR3 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR4 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR4 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR4 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR5 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR5 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR5 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR6 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR6 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR6 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR7 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR7 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR7 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR8 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR8 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR8 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR9 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR9 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR9 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR10 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR10 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR10 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR11 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR11 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR11 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR12 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR12 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR12 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR13 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR13 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR13 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR14 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR14 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR14 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR15 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR15 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR15 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR16 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR16 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR16 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR17 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR17 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR17 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR18 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR18 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR18 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR19 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR19 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR19 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR20 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR20 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR20 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR21 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR21 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR21 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR22 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR22 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR22 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR23 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR23 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR23 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR24 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR24 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR24 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR25 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR25 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR25 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR26 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR26 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR26 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR27 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR27 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR27 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR28 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR28 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR28 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR29 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR29 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR29 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR30 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR30 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR30 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	PROCID   HSEM_RLR31 = 0xFF << 0  //+ Semaphore ProcessID
	MASTERID HSEM_RLR31 = 0xFF << 8  //+ Semaphore MasterID
	LOCK     HSEM_RLR31 = 0x01 << 31 //+ Lock indication
)

const (
	PROCIDn   = 0
	MASTERIDn = 8
	LOCKn     = 31
)

const (
	ISEM0  HSEM_IER = 0x01 << 0  //+ Interrupt semaphore n enable bit
	ISEM1  HSEM_IER = 0x01 << 1  //+ Interrupt semaphore n enable bit
	ISEM2  HSEM_IER = 0x01 << 2  //+ Interrupt semaphore n enable bit
	ISEM3  HSEM_IER = 0x01 << 3  //+ Interrupt semaphore n enable bit
	ISEM4  HSEM_IER = 0x01 << 4  //+ Interrupt semaphore n enable bit
	ISEM5  HSEM_IER = 0x01 << 5  //+ Interrupt semaphore n enable bit
	ISEM6  HSEM_IER = 0x01 << 6  //+ Interrupt semaphore n enable bit
	ISEM7  HSEM_IER = 0x01 << 7  //+ Interrupt semaphore n enable bit
	ISEM8  HSEM_IER = 0x01 << 8  //+ Interrupt semaphore n enable bit
	ISEM9  HSEM_IER = 0x01 << 9  //+ Interrupt semaphore n enable bit
	ISEM10 HSEM_IER = 0x01 << 10 //+ Interrupt semaphore n enable bit
	ISEM11 HSEM_IER = 0x01 << 11 //+ Interrupt semaphore n enable bit
	ISEM12 HSEM_IER = 0x01 << 12 //+ Interrupt semaphore n enable bit
	ISEM13 HSEM_IER = 0x01 << 13 //+ Interrupt semaphore n enable bit
	ISEM14 HSEM_IER = 0x01 << 14 //+ Interrupt semaphore n enable bit
	ISEM15 HSEM_IER = 0x01 << 15 //+ Interrupt semaphore n enable bit
	ISEM16 HSEM_IER = 0x01 << 16 //+ Interrupt semaphore n enable bit
	ISEM17 HSEM_IER = 0x01 << 17 //+ Interrupt semaphore n enable bit
	ISEM18 HSEM_IER = 0x01 << 18 //+ Interrupt semaphore n enable bit
	ISEM19 HSEM_IER = 0x01 << 19 //+ Interrupt semaphore n enable bit
	ISEM20 HSEM_IER = 0x01 << 20 //+ Interrupt semaphore n enable bit
	ISEM21 HSEM_IER = 0x01 << 21 //+ Interrupt semaphore n enable bit
	ISEM22 HSEM_IER = 0x01 << 22 //+ Interrupt semaphore n enable bit
	ISEM23 HSEM_IER = 0x01 << 23 //+ Interrupt semaphore n enable bit
	ISEM24 HSEM_IER = 0x01 << 24 //+ Interrupt semaphore n enable bit
	ISEM25 HSEM_IER = 0x01 << 25 //+ Interrupt semaphore n enable bit
	ISEM26 HSEM_IER = 0x01 << 26 //+ Interrupt semaphore n enable bit
	ISEM27 HSEM_IER = 0x01 << 27 //+ Interrupt semaphore n enable bit
	ISEM28 HSEM_IER = 0x01 << 28 //+ Interrupt semaphore n enable bit
	ISEM29 HSEM_IER = 0x01 << 29 //+ Interrupt semaphore n enable bit
	ISEM30 HSEM_IER = 0x01 << 30 //+ Interrupt semaphore n enable bit
	ISEM31 HSEM_IER = 0x01 << 31 //+ Interrupt(N) semaphore n enable bit.
)

const (
	ISEM0n  = 0
	ISEM1n  = 1
	ISEM2n  = 2
	ISEM3n  = 3
	ISEM4n  = 4
	ISEM5n  = 5
	ISEM6n  = 6
	ISEM7n  = 7
	ISEM8n  = 8
	ISEM9n  = 9
	ISEM10n = 10
	ISEM11n = 11
	ISEM12n = 12
	ISEM13n = 13
	ISEM14n = 14
	ISEM15n = 15
	ISEM16n = 16
	ISEM17n = 17
	ISEM18n = 18
	ISEM19n = 19
	ISEM20n = 20
	ISEM21n = 21
	ISEM22n = 22
	ISEM23n = 23
	ISEM24n = 24
	ISEM25n = 25
	ISEM26n = 26
	ISEM27n = 27
	ISEM28n = 28
	ISEM29n = 29
	ISEM30n = 30
	ISEM31n = 31
)

const (
	ISEM0  HSEM_ICR = 0x01 << 0  //+ Interrupt(N) semaphore n clear bit
	ISEM1  HSEM_ICR = 0x01 << 1  //+ Interrupt(N) semaphore n clear bit
	ISEM2  HSEM_ICR = 0x01 << 2  //+ Interrupt(N) semaphore n clear bit
	ISEM3  HSEM_ICR = 0x01 << 3  //+ Interrupt(N) semaphore n clear bit
	ISEM4  HSEM_ICR = 0x01 << 4  //+ Interrupt(N) semaphore n clear bit
	ISEM5  HSEM_ICR = 0x01 << 5  //+ Interrupt(N) semaphore n clear bit
	ISEM6  HSEM_ICR = 0x01 << 6  //+ Interrupt(N) semaphore n clear bit
	ISEM7  HSEM_ICR = 0x01 << 7  //+ Interrupt(N) semaphore n clear bit
	ISEM8  HSEM_ICR = 0x01 << 8  //+ Interrupt(N) semaphore n clear bit
	ISEM9  HSEM_ICR = 0x01 << 9  //+ Interrupt(N) semaphore n clear bit
	ISEM10 HSEM_ICR = 0x01 << 10 //+ Interrupt(N) semaphore n clear bit
	ISEM11 HSEM_ICR = 0x01 << 11 //+ Interrupt(N) semaphore n clear bit
	ISEM12 HSEM_ICR = 0x01 << 12 //+ Interrupt(N) semaphore n clear bit
	ISEM13 HSEM_ICR = 0x01 << 13 //+ Interrupt(N) semaphore n clear bit
	ISEM14 HSEM_ICR = 0x01 << 14 //+ Interrupt(N) semaphore n clear bit
	ISEM15 HSEM_ICR = 0x01 << 15 //+ Interrupt(N) semaphore n clear bit
	ISEM16 HSEM_ICR = 0x01 << 16 //+ Interrupt(N) semaphore n clear bit
	ISEM17 HSEM_ICR = 0x01 << 17 //+ Interrupt(N) semaphore n clear bit
	ISEM18 HSEM_ICR = 0x01 << 18 //+ Interrupt(N) semaphore n clear bit
	ISEM19 HSEM_ICR = 0x01 << 19 //+ Interrupt(N) semaphore n clear bit
	ISEM20 HSEM_ICR = 0x01 << 20 //+ Interrupt(N) semaphore n clear bit
	ISEM21 HSEM_ICR = 0x01 << 21 //+ Interrupt(N) semaphore n clear bit
	ISEM22 HSEM_ICR = 0x01 << 22 //+ Interrupt(N) semaphore n clear bit
	ISEM23 HSEM_ICR = 0x01 << 23 //+ Interrupt(N) semaphore n clear bit
	ISEM24 HSEM_ICR = 0x01 << 24 //+ Interrupt(N) semaphore n clear bit
	ISEM25 HSEM_ICR = 0x01 << 25 //+ Interrupt(N) semaphore n clear bit
	ISEM26 HSEM_ICR = 0x01 << 26 //+ Interrupt(N) semaphore n clear bit
	ISEM27 HSEM_ICR = 0x01 << 27 //+ Interrupt(N) semaphore n clear bit
	ISEM28 HSEM_ICR = 0x01 << 28 //+ Interrupt(N) semaphore n clear bit
	ISEM29 HSEM_ICR = 0x01 << 29 //+ Interrupt(N) semaphore n clear bit
	ISEM30 HSEM_ICR = 0x01 << 30 //+ Interrupt(N) semaphore n clear bit
	ISEM31 HSEM_ICR = 0x01 << 31 //+ Interrupt(N) semaphore n clear bit
)

const (
	ISEM0n  = 0
	ISEM1n  = 1
	ISEM2n  = 2
	ISEM3n  = 3
	ISEM4n  = 4
	ISEM5n  = 5
	ISEM6n  = 6
	ISEM7n  = 7
	ISEM8n  = 8
	ISEM9n  = 9
	ISEM10n = 10
	ISEM11n = 11
	ISEM12n = 12
	ISEM13n = 13
	ISEM14n = 14
	ISEM15n = 15
	ISEM16n = 16
	ISEM17n = 17
	ISEM18n = 18
	ISEM19n = 19
	ISEM20n = 20
	ISEM21n = 21
	ISEM22n = 22
	ISEM23n = 23
	ISEM24n = 24
	ISEM25n = 25
	ISEM26n = 26
	ISEM27n = 27
	ISEM28n = 28
	ISEM29n = 29
	ISEM30n = 30
	ISEM31n = 31
)

const (
	ISEM0  HSEM_ISR = 0x01 << 0  //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM1  HSEM_ISR = 0x01 << 1  //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM2  HSEM_ISR = 0x01 << 2  //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM3  HSEM_ISR = 0x01 << 3  //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM4  HSEM_ISR = 0x01 << 4  //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM5  HSEM_ISR = 0x01 << 5  //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM6  HSEM_ISR = 0x01 << 6  //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM7  HSEM_ISR = 0x01 << 7  //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM8  HSEM_ISR = 0x01 << 8  //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM9  HSEM_ISR = 0x01 << 9  //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM10 HSEM_ISR = 0x01 << 10 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM11 HSEM_ISR = 0x01 << 11 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM12 HSEM_ISR = 0x01 << 12 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM13 HSEM_ISR = 0x01 << 13 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM14 HSEM_ISR = 0x01 << 14 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM15 HSEM_ISR = 0x01 << 15 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM16 HSEM_ISR = 0x01 << 16 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM17 HSEM_ISR = 0x01 << 17 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM18 HSEM_ISR = 0x01 << 18 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM19 HSEM_ISR = 0x01 << 19 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM20 HSEM_ISR = 0x01 << 20 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM21 HSEM_ISR = 0x01 << 21 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM22 HSEM_ISR = 0x01 << 22 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM23 HSEM_ISR = 0x01 << 23 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM24 HSEM_ISR = 0x01 << 24 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM25 HSEM_ISR = 0x01 << 25 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM26 HSEM_ISR = 0x01 << 26 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM27 HSEM_ISR = 0x01 << 27 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM28 HSEM_ISR = 0x01 << 28 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM29 HSEM_ISR = 0x01 << 29 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM30 HSEM_ISR = 0x01 << 30 //+ Interrupt(N) semaphore n status bit before enable (mask)
	ISEM31 HSEM_ISR = 0x01 << 31 //+ Interrupt(N) semaphore n status bit before enable (mask)
)

const (
	ISEM0n  = 0
	ISEM1n  = 1
	ISEM2n  = 2
	ISEM3n  = 3
	ISEM4n  = 4
	ISEM5n  = 5
	ISEM6n  = 6
	ISEM7n  = 7
	ISEM8n  = 8
	ISEM9n  = 9
	ISEM10n = 10
	ISEM11n = 11
	ISEM12n = 12
	ISEM13n = 13
	ISEM14n = 14
	ISEM15n = 15
	ISEM16n = 16
	ISEM17n = 17
	ISEM18n = 18
	ISEM19n = 19
	ISEM20n = 20
	ISEM21n = 21
	ISEM22n = 22
	ISEM23n = 23
	ISEM24n = 24
	ISEM25n = 25
	ISEM26n = 26
	ISEM27n = 27
	ISEM28n = 28
	ISEM29n = 29
	ISEM30n = 30
	ISEM31n = 31
)

const (
	ISEM0  HSEM_MISR = 0x01 << 0  //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM1  HSEM_MISR = 0x01 << 1  //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM2  HSEM_MISR = 0x01 << 2  //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM3  HSEM_MISR = 0x01 << 3  //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM4  HSEM_MISR = 0x01 << 4  //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM5  HSEM_MISR = 0x01 << 5  //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM6  HSEM_MISR = 0x01 << 6  //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM7  HSEM_MISR = 0x01 << 7  //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM8  HSEM_MISR = 0x01 << 8  //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM9  HSEM_MISR = 0x01 << 9  //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM10 HSEM_MISR = 0x01 << 10 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM11 HSEM_MISR = 0x01 << 11 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM12 HSEM_MISR = 0x01 << 12 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM13 HSEM_MISR = 0x01 << 13 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM14 HSEM_MISR = 0x01 << 14 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM15 HSEM_MISR = 0x01 << 15 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM16 HSEM_MISR = 0x01 << 16 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM17 HSEM_MISR = 0x01 << 17 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM18 HSEM_MISR = 0x01 << 18 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM19 HSEM_MISR = 0x01 << 19 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM20 HSEM_MISR = 0x01 << 20 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM21 HSEM_MISR = 0x01 << 21 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM22 HSEM_MISR = 0x01 << 22 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM23 HSEM_MISR = 0x01 << 23 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM24 HSEM_MISR = 0x01 << 24 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM25 HSEM_MISR = 0x01 << 25 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM26 HSEM_MISR = 0x01 << 26 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM27 HSEM_MISR = 0x01 << 27 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM28 HSEM_MISR = 0x01 << 28 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM29 HSEM_MISR = 0x01 << 29 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM30 HSEM_MISR = 0x01 << 30 //+ masked interrupt(N) semaphore n status bit after enable (mask)
	ISEM31 HSEM_MISR = 0x01 << 31 //+ masked interrupt(N) semaphore n status bit after enable (mask)
)

const (
	ISEM0n  = 0
	ISEM1n  = 1
	ISEM2n  = 2
	ISEM3n  = 3
	ISEM4n  = 4
	ISEM5n  = 5
	ISEM6n  = 6
	ISEM7n  = 7
	ISEM8n  = 8
	ISEM9n  = 9
	ISEM10n = 10
	ISEM11n = 11
	ISEM12n = 12
	ISEM13n = 13
	ISEM14n = 14
	ISEM15n = 15
	ISEM16n = 16
	ISEM17n = 17
	ISEM18n = 18
	ISEM19n = 19
	ISEM20n = 20
	ISEM21n = 21
	ISEM22n = 22
	ISEM23n = 23
	ISEM24n = 24
	ISEM25n = 25
	ISEM26n = 26
	ISEM27n = 27
	ISEM28n = 28
	ISEM29n = 29
	ISEM30n = 30
	ISEM31n = 31
)

const (
	MASTERID HSEM_CR = 0xFF << 8    //+ MasterID of semaphores to be cleared
	KEY      HSEM_CR = 0xFFFF << 16 //+ Semaphore clear Key
)

const (
	MASTERIDn = 8
	KEYn      = 16
)

const (
	KEY HSEM_KEYR = 0xFFFF << 16 //+ Semaphore Clear Key
)

const (
	KEYn = 16
)
