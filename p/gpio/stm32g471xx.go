// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32g471xx

// Package gpio provides access to the registers of the GPIO peripheral.
//
// Instances:
//  GPIOA  GPIOA_BASE  AHB2  -  General-purpose I/Os
//  GPIOB  GPIOB_BASE  AHB2  -  General-purpose I/Os
//  GPIOC  GPIOC_BASE  AHB2  -  General-purpose I/Os
//  GPIOD  GPIOD_BASE  AHB2  -  General-purpose I/Os
//  GPIOE  GPIOE_BASE  AHB2  -  General-purpose I/Os
//  GPIOF  GPIOF_BASE  AHB2  -  General-purpose I/Os
//  GPIOG  GPIOG_BASE  AHB2  -  General-purpose I/Os
// Registers:
//  0x000 32  MODER    GPIO port mode register
//  0x004 32  OTYPER   GPIO port output type register
//  0x008 32  OSPEEDR  GPIO port output speed register
//  0x00C 32  PUPDR    GPIO port pull-up/pull-down register
//  0x010 32  IDR      GPIO port input data register
//  0x014 32  ODR      GPIO port output data register
//  0x018 32  BSRR     GPIO port bit set/reset register
//  0x01C 32  LCKR     GPIO port configuration lock register
//  0x020 32  AFRL     GPIO alternate function low register
//  0x024 32  AFRH     GPIO alternate function high register
//  0x028 32  BRR      GPIO port bit reset register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package gpio

const (
	MODER0  MODER = 0x03 << 0  //+ Port x configuration bits (y = 0..15)
	MODER1  MODER = 0x03 << 2  //+ Port x configuration bits (y = 0..15)
	MODER2  MODER = 0x03 << 4  //+ Port x configuration bits (y = 0..15)
	MODER3  MODER = 0x03 << 6  //+ Port x configuration bits (y = 0..15)
	MODER4  MODER = 0x03 << 8  //+ Port x configuration bits (y = 0..15)
	MODER5  MODER = 0x03 << 10 //+ Port x configuration bits (y = 0..15)
	MODER6  MODER = 0x03 << 12 //+ Port x configuration bits (y = 0..15)
	MODER7  MODER = 0x03 << 14 //+ Port x configuration bits (y = 0..15)
	MODER8  MODER = 0x03 << 16 //+ Port x configuration bits (y = 0..15)
	MODER9  MODER = 0x03 << 18 //+ Port x configuration bits (y = 0..15)
	MODER10 MODER = 0x03 << 20 //+ Port x configuration bits (y = 0..15)
	MODER11 MODER = 0x03 << 22 //+ Port x configuration bits (y = 0..15)
	MODER12 MODER = 0x03 << 24 //+ Port x configuration bits (y = 0..15)
	MODER13 MODER = 0x03 << 26 //+ Port x configuration bits (y = 0..15)
	MODER14 MODER = 0x03 << 28 //+ Port x configuration bits (y = 0..15)
	MODER15 MODER = 0x03 << 30 //+ Port x configuration bits (y = 0..15)
)

const (
	MODER0n  = 0
	MODER1n  = 2
	MODER2n  = 4
	MODER3n  = 6
	MODER4n  = 8
	MODER5n  = 10
	MODER6n  = 12
	MODER7n  = 14
	MODER8n  = 16
	MODER9n  = 18
	MODER10n = 20
	MODER11n = 22
	MODER12n = 24
	MODER13n = 26
	MODER14n = 28
	MODER15n = 30
)

const (
	OT0  OTYPER = 0x01 << 0  //+ Port x configuration bits (y = 0..15)
	OT1  OTYPER = 0x01 << 1  //+ Port x configuration bits (y = 0..15)
	OT2  OTYPER = 0x01 << 2  //+ Port x configuration bits (y = 0..15)
	OT3  OTYPER = 0x01 << 3  //+ Port x configuration bits (y = 0..15)
	OT4  OTYPER = 0x01 << 4  //+ Port x configuration bits (y = 0..15)
	OT5  OTYPER = 0x01 << 5  //+ Port x configuration bits (y = 0..15)
	OT6  OTYPER = 0x01 << 6  //+ Port x configuration bits (y = 0..15)
	OT7  OTYPER = 0x01 << 7  //+ Port x configuration bits (y = 0..15)
	OT8  OTYPER = 0x01 << 8  //+ Port x configuration bits (y = 0..15)
	OT9  OTYPER = 0x01 << 9  //+ Port x configuration bits (y = 0..15)
	OT10 OTYPER = 0x01 << 10 //+ Port x configuration bits (y = 0..15)
	OT11 OTYPER = 0x01 << 11 //+ Port x configuration bits (y = 0..15)
	OT12 OTYPER = 0x01 << 12 //+ Port x configuration bits (y = 0..15)
	OT13 OTYPER = 0x01 << 13 //+ Port x configuration bits (y = 0..15)
	OT14 OTYPER = 0x01 << 14 //+ Port x configuration bits (y = 0..15)
	OT15 OTYPER = 0x01 << 15 //+ Port x configuration bits (y = 0..15)
)

const (
	OT0n  = 0
	OT1n  = 1
	OT2n  = 2
	OT3n  = 3
	OT4n  = 4
	OT5n  = 5
	OT6n  = 6
	OT7n  = 7
	OT8n  = 8
	OT9n  = 9
	OT10n = 10
	OT11n = 11
	OT12n = 12
	OT13n = 13
	OT14n = 14
	OT15n = 15
)

const (
	OSPEEDR0  OSPEEDR = 0x03 << 0  //+ Port x configuration bits (y = 0..15)
	OSPEEDR1  OSPEEDR = 0x03 << 2  //+ Port x configuration bits (y = 0..15)
	OSPEEDR2  OSPEEDR = 0x03 << 4  //+ Port x configuration bits (y = 0..15)
	OSPEEDR3  OSPEEDR = 0x03 << 6  //+ Port x configuration bits (y = 0..15)
	OSPEEDR4  OSPEEDR = 0x03 << 8  //+ Port x configuration bits (y = 0..15)
	OSPEEDR5  OSPEEDR = 0x03 << 10 //+ Port x configuration bits (y = 0..15)
	OSPEEDR6  OSPEEDR = 0x03 << 12 //+ Port x configuration bits (y = 0..15)
	OSPEEDR7  OSPEEDR = 0x03 << 14 //+ Port x configuration bits (y = 0..15)
	OSPEEDR8  OSPEEDR = 0x03 << 16 //+ Port x configuration bits (y = 0..15)
	OSPEEDR9  OSPEEDR = 0x03 << 18 //+ Port x configuration bits (y = 0..15)
	OSPEEDR10 OSPEEDR = 0x03 << 20 //+ Port x configuration bits (y = 0..15)
	OSPEEDR11 OSPEEDR = 0x03 << 22 //+ Port x configuration bits (y = 0..15)
	OSPEEDR12 OSPEEDR = 0x03 << 24 //+ Port x configuration bits (y = 0..15)
	OSPEEDR13 OSPEEDR = 0x03 << 26 //+ Port x configuration bits (y = 0..15)
	OSPEEDR14 OSPEEDR = 0x03 << 28 //+ Port x configuration bits (y = 0..15)
	OSPEEDR15 OSPEEDR = 0x03 << 30 //+ Port x configuration bits (y = 0..15)
)

const (
	OSPEEDR0n  = 0
	OSPEEDR1n  = 2
	OSPEEDR2n  = 4
	OSPEEDR3n  = 6
	OSPEEDR4n  = 8
	OSPEEDR5n  = 10
	OSPEEDR6n  = 12
	OSPEEDR7n  = 14
	OSPEEDR8n  = 16
	OSPEEDR9n  = 18
	OSPEEDR10n = 20
	OSPEEDR11n = 22
	OSPEEDR12n = 24
	OSPEEDR13n = 26
	OSPEEDR14n = 28
	OSPEEDR15n = 30
)

const (
	PUPDR0  PUPDR = 0x03 << 0  //+ Port x configuration bits (y = 0..15)
	PUPDR1  PUPDR = 0x03 << 2  //+ Port x configuration bits (y = 0..15)
	PUPDR2  PUPDR = 0x03 << 4  //+ Port x configuration bits (y = 0..15)
	PUPDR3  PUPDR = 0x03 << 6  //+ Port x configuration bits (y = 0..15)
	PUPDR4  PUPDR = 0x03 << 8  //+ Port x configuration bits (y = 0..15)
	PUPDR5  PUPDR = 0x03 << 10 //+ Port x configuration bits (y = 0..15)
	PUPDR6  PUPDR = 0x03 << 12 //+ Port x configuration bits (y = 0..15)
	PUPDR7  PUPDR = 0x03 << 14 //+ Port x configuration bits (y = 0..15)
	PUPDR8  PUPDR = 0x03 << 16 //+ Port x configuration bits (y = 0..15)
	PUPDR9  PUPDR = 0x03 << 18 //+ Port x configuration bits (y = 0..15)
	PUPDR10 PUPDR = 0x03 << 20 //+ Port x configuration bits (y = 0..15)
	PUPDR11 PUPDR = 0x03 << 22 //+ Port x configuration bits (y = 0..15)
	PUPDR12 PUPDR = 0x03 << 24 //+ Port x configuration bits (y = 0..15)
	PUPDR13 PUPDR = 0x03 << 26 //+ Port x configuration bits (y = 0..15)
	PUPDR14 PUPDR = 0x03 << 28 //+ Port x configuration bits (y = 0..15)
	PUPDR15 PUPDR = 0x03 << 30 //+ Port x configuration bits (y = 0..15)
)

const (
	PUPDR0n  = 0
	PUPDR1n  = 2
	PUPDR2n  = 4
	PUPDR3n  = 6
	PUPDR4n  = 8
	PUPDR5n  = 10
	PUPDR6n  = 12
	PUPDR7n  = 14
	PUPDR8n  = 16
	PUPDR9n  = 18
	PUPDR10n = 20
	PUPDR11n = 22
	PUPDR12n = 24
	PUPDR13n = 26
	PUPDR14n = 28
	PUPDR15n = 30
)

const (
	IDR0  IDR = 0x01 << 0  //+ Port input data (y = 0..15)
	IDR1  IDR = 0x01 << 1  //+ Port input data (y = 0..15)
	IDR2  IDR = 0x01 << 2  //+ Port input data (y = 0..15)
	IDR3  IDR = 0x01 << 3  //+ Port input data (y = 0..15)
	IDR4  IDR = 0x01 << 4  //+ Port input data (y = 0..15)
	IDR5  IDR = 0x01 << 5  //+ Port input data (y = 0..15)
	IDR6  IDR = 0x01 << 6  //+ Port input data (y = 0..15)
	IDR7  IDR = 0x01 << 7  //+ Port input data (y = 0..15)
	IDR8  IDR = 0x01 << 8  //+ Port input data (y = 0..15)
	IDR9  IDR = 0x01 << 9  //+ Port input data (y = 0..15)
	IDR10 IDR = 0x01 << 10 //+ Port input data (y = 0..15)
	IDR11 IDR = 0x01 << 11 //+ Port input data (y = 0..15)
	IDR12 IDR = 0x01 << 12 //+ Port input data (y = 0..15)
	IDR13 IDR = 0x01 << 13 //+ Port input data (y = 0..15)
	IDR14 IDR = 0x01 << 14 //+ Port input data (y = 0..15)
	IDR15 IDR = 0x01 << 15 //+ Port input data (y = 0..15)
)

const (
	IDR0n  = 0
	IDR1n  = 1
	IDR2n  = 2
	IDR3n  = 3
	IDR4n  = 4
	IDR5n  = 5
	IDR6n  = 6
	IDR7n  = 7
	IDR8n  = 8
	IDR9n  = 9
	IDR10n = 10
	IDR11n = 11
	IDR12n = 12
	IDR13n = 13
	IDR14n = 14
	IDR15n = 15
)

const (
	ODR0  ODR = 0x01 << 0  //+ Port output data (y = 0..15)
	ODR1  ODR = 0x01 << 1  //+ Port output data (y = 0..15)
	ODR2  ODR = 0x01 << 2  //+ Port output data (y = 0..15)
	ODR3  ODR = 0x01 << 3  //+ Port output data (y = 0..15)
	ODR4  ODR = 0x01 << 4  //+ Port output data (y = 0..15)
	ODR5  ODR = 0x01 << 5  //+ Port output data (y = 0..15)
	ODR6  ODR = 0x01 << 6  //+ Port output data (y = 0..15)
	ODR7  ODR = 0x01 << 7  //+ Port output data (y = 0..15)
	ODR8  ODR = 0x01 << 8  //+ Port output data (y = 0..15)
	ODR9  ODR = 0x01 << 9  //+ Port output data (y = 0..15)
	ODR10 ODR = 0x01 << 10 //+ Port output data (y = 0..15)
	ODR11 ODR = 0x01 << 11 //+ Port output data (y = 0..15)
	ODR12 ODR = 0x01 << 12 //+ Port output data (y = 0..15)
	ODR13 ODR = 0x01 << 13 //+ Port output data (y = 0..15)
	ODR14 ODR = 0x01 << 14 //+ Port output data (y = 0..15)
	ODR15 ODR = 0x01 << 15 //+ Port output data (y = 0..15)
)

const (
	ODR0n  = 0
	ODR1n  = 1
	ODR2n  = 2
	ODR3n  = 3
	ODR4n  = 4
	ODR5n  = 5
	ODR6n  = 6
	ODR7n  = 7
	ODR8n  = 8
	ODR9n  = 9
	ODR10n = 10
	ODR11n = 11
	ODR12n = 12
	ODR13n = 13
	ODR14n = 14
	ODR15n = 15
)

const (
	BS0  BSRR = 0x01 << 0  //+ Port x set bit y (y= 0..15)
	BS1  BSRR = 0x01 << 1  //+ Port x set bit y (y= 0..15)
	BS2  BSRR = 0x01 << 2  //+ Port x set bit y (y= 0..15)
	BS3  BSRR = 0x01 << 3  //+ Port x set bit y (y= 0..15)
	BS4  BSRR = 0x01 << 4  //+ Port x set bit y (y= 0..15)
	BS5  BSRR = 0x01 << 5  //+ Port x set bit y (y= 0..15)
	BS6  BSRR = 0x01 << 6  //+ Port x set bit y (y= 0..15)
	BS7  BSRR = 0x01 << 7  //+ Port x set bit y (y= 0..15)
	BS8  BSRR = 0x01 << 8  //+ Port x set bit y (y= 0..15)
	BS9  BSRR = 0x01 << 9  //+ Port x set bit y (y= 0..15)
	BS10 BSRR = 0x01 << 10 //+ Port x set bit y (y= 0..15)
	BS11 BSRR = 0x01 << 11 //+ Port x set bit y (y= 0..15)
	BS12 BSRR = 0x01 << 12 //+ Port x set bit y (y= 0..15)
	BS13 BSRR = 0x01 << 13 //+ Port x set bit y (y= 0..15)
	BS14 BSRR = 0x01 << 14 //+ Port x set bit y (y= 0..15)
	BS15 BSRR = 0x01 << 15 //+ Port x set bit y (y= 0..15)
	BR0  BSRR = 0x01 << 16 //+ Port x set bit y (y= 0..15)
	BR1  BSRR = 0x01 << 17 //+ Port x reset bit y (y = 0..15)
	BR2  BSRR = 0x01 << 18 //+ Port x reset bit y (y = 0..15)
	BR3  BSRR = 0x01 << 19 //+ Port x reset bit y (y = 0..15)
	BR4  BSRR = 0x01 << 20 //+ Port x reset bit y (y = 0..15)
	BR5  BSRR = 0x01 << 21 //+ Port x reset bit y (y = 0..15)
	BR6  BSRR = 0x01 << 22 //+ Port x reset bit y (y = 0..15)
	BR7  BSRR = 0x01 << 23 //+ Port x reset bit y (y = 0..15)
	BR8  BSRR = 0x01 << 24 //+ Port x reset bit y (y = 0..15)
	BR9  BSRR = 0x01 << 25 //+ Port x reset bit y (y = 0..15)
	BR10 BSRR = 0x01 << 26 //+ Port x reset bit y (y = 0..15)
	BR11 BSRR = 0x01 << 27 //+ Port x reset bit y (y = 0..15)
	BR12 BSRR = 0x01 << 28 //+ Port x reset bit y (y = 0..15)
	BR13 BSRR = 0x01 << 29 //+ Port x reset bit y (y = 0..15)
	BR14 BSRR = 0x01 << 30 //+ Port x reset bit y (y = 0..15)
	BR15 BSRR = 0x01 << 31 //+ Port x reset bit y (y = 0..15)
)

const (
	BS0n  = 0
	BS1n  = 1
	BS2n  = 2
	BS3n  = 3
	BS4n  = 4
	BS5n  = 5
	BS6n  = 6
	BS7n  = 7
	BS8n  = 8
	BS9n  = 9
	BS10n = 10
	BS11n = 11
	BS12n = 12
	BS13n = 13
	BS14n = 14
	BS15n = 15
	BR0n  = 16
	BR1n  = 17
	BR2n  = 18
	BR3n  = 19
	BR4n  = 20
	BR5n  = 21
	BR6n  = 22
	BR7n  = 23
	BR8n  = 24
	BR9n  = 25
	BR10n = 26
	BR11n = 27
	BR12n = 28
	BR13n = 29
	BR14n = 30
	BR15n = 31
)

const (
	LCK0  LCKR = 0x01 << 0  //+ Port x lock bit y (y= 0..15)
	LCK1  LCKR = 0x01 << 1  //+ Port x lock bit y (y= 0..15)
	LCK2  LCKR = 0x01 << 2  //+ Port x lock bit y (y= 0..15)
	LCK3  LCKR = 0x01 << 3  //+ Port x lock bit y (y= 0..15)
	LCK4  LCKR = 0x01 << 4  //+ Port x lock bit y (y= 0..15)
	LCK5  LCKR = 0x01 << 5  //+ Port x lock bit y (y= 0..15)
	LCK6  LCKR = 0x01 << 6  //+ Port x lock bit y (y= 0..15)
	LCK7  LCKR = 0x01 << 7  //+ Port x lock bit y (y= 0..15)
	LCK8  LCKR = 0x01 << 8  //+ Port x lock bit y (y= 0..15)
	LCK9  LCKR = 0x01 << 9  //+ Port x lock bit y (y= 0..15)
	LCK10 LCKR = 0x01 << 10 //+ Port x lock bit y (y= 0..15)
	LCK11 LCKR = 0x01 << 11 //+ Port x lock bit y (y= 0..15)
	LCK12 LCKR = 0x01 << 12 //+ Port x lock bit y (y= 0..15)
	LCK13 LCKR = 0x01 << 13 //+ Port x lock bit y (y= 0..15)
	LCK14 LCKR = 0x01 << 14 //+ Port x lock bit y (y= 0..15)
	LCK15 LCKR = 0x01 << 15 //+ Port x lock bit y (y= 0..15)
	LCKK  LCKR = 0x01 << 16 //+ Port x lock bit y (y= 0..15)
)

const (
	LCK0n  = 0
	LCK1n  = 1
	LCK2n  = 2
	LCK3n  = 3
	LCK4n  = 4
	LCK5n  = 5
	LCK6n  = 6
	LCK7n  = 7
	LCK8n  = 8
	LCK9n  = 9
	LCK10n = 10
	LCK11n = 11
	LCK12n = 12
	LCK13n = 13
	LCK14n = 14
	LCK15n = 15
	LCKKn  = 16
)

const (
	AFRL0 AFRL = 0x0F << 0  //+ Alternate function selection for port x bit y (y = 0..7)
	AFRL1 AFRL = 0x0F << 4  //+ Alternate function selection for port x bit y (y = 0..7)
	AFRL2 AFRL = 0x0F << 8  //+ Alternate function selection for port x bit y (y = 0..7)
	AFRL3 AFRL = 0x0F << 12 //+ Alternate function selection for port x bit y (y = 0..7)
	AFRL4 AFRL = 0x0F << 16 //+ Alternate function selection for port x bit y (y = 0..7)
	AFRL5 AFRL = 0x0F << 20 //+ Alternate function selection for port x bit y (y = 0..7)
	AFRL6 AFRL = 0x0F << 24 //+ Alternate function selection for port x bit y (y = 0..7)
	AFRL7 AFRL = 0x0F << 28 //+ Alternate function selection for port x bit y (y = 0..7)
)

const (
	AFRL0n = 0
	AFRL1n = 4
	AFRL2n = 8
	AFRL3n = 12
	AFRL4n = 16
	AFRL5n = 20
	AFRL6n = 24
	AFRL7n = 28
)

const (
	AFRH8  AFRH = 0x0F << 0  //+ Alternate function selection for port x bit y (y = 8..15)
	AFRH9  AFRH = 0x0F << 4  //+ Alternate function selection for port x bit y (y = 8..15)
	AFRH10 AFRH = 0x0F << 8  //+ Alternate function selection for port x bit y (y = 8..15)
	AFRH11 AFRH = 0x0F << 12 //+ Alternate function selection for port x bit y (y = 8..15)
	AFRH12 AFRH = 0x0F << 16 //+ Alternate function selection for port x bit y (y = 8..15)
	AFRH13 AFRH = 0x0F << 20 //+ Alternate function selection for port x bit y (y = 8..15)
	AFRH14 AFRH = 0x0F << 24 //+ Alternate function selection for port x bit y (y = 8..15)
	AFRH15 AFRH = 0x0F << 28 //+ Alternate function selection for port x bit y (y = 8..15)
)

const (
	AFRH8n  = 0
	AFRH9n  = 4
	AFRH10n = 8
	AFRH11n = 12
	AFRH12n = 16
	AFRH13n = 20
	AFRH14n = 24
	AFRH15n = 28
)

const (
	BC0  BRR = 0x01 << 0  //+ Port x reset (clear) bit y
	BC1  BRR = 0x01 << 1  //+ Port x reset (clear) bit y
	BC2  BRR = 0x01 << 2  //+ Port x reset (clear) bit y
	BC3  BRR = 0x01 << 3  //+ Port x reset (clear) bit y
	BC4  BRR = 0x01 << 4  //+ Port x reset (clear) bit y
	BC5  BRR = 0x01 << 5  //+ Port x reset (clear) bit y
	BC6  BRR = 0x01 << 6  //+ Port x reset (clear) bit y
	BC7  BRR = 0x01 << 7  //+ Port x reset (clear) bit y
	BC8  BRR = 0x01 << 8  //+ Port x reset (clear) bit y
	BC9  BRR = 0x01 << 9  //+ Port x reset (clear) bit y
	BC10 BRR = 0x01 << 10 //+ Port x reset (clear) bit y
	BC11 BRR = 0x01 << 11 //+ Port x reset (clear) bit y
	BC12 BRR = 0x01 << 12 //+ Port x reset (clear) bit y
	BC13 BRR = 0x01 << 13 //+ Port x reset (clear) bit y
	BC14 BRR = 0x01 << 14 //+ Port x reset (clear) bit y
	BC15 BRR = 0x01 << 15 //+ Port x reset (clear) bit y
)

const (
	BC0n  = 0
	BC1n  = 1
	BC2n  = 2
	BC3n  = 3
	BC4n  = 4
	BC5n  = 5
	BC6n  = 6
	BC7n  = 7
	BC8n  = 8
	BC9n  = 9
	BC10n = 10
	BC11n = 11
	BC12n = 12
	BC13n = 13
	BC14n = 14
	BC15n = 15
)
