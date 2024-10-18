// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package gpio provides access to the registers of the GPIOA peripheral.
//
// Instances:
//
//	GPIOA  GPIOA_BASE  AHB4  -
//	GPIOB  GPIOB_BASE  AHB4  -
//	GPIOC  GPIOC_BASE  AHB4  -
//	GPIOD  GPIOD_BASE  AHB4  -
//	GPIOE  GPIOE_BASE  AHB4  -
//	GPIOF  GPIOF_BASE  AHB4  -
//	GPIOG  GPIOG_BASE  AHB4  -
//	GPIOH  GPIOH_BASE  AHB4  -
//	GPIOI  GPIOI_BASE  AHB4  -
//	GPIOJ  GPIOJ_BASE  AHB4  -
//	GPIOK  GPIOK_BASE  AHB4  -
//
// Registers:
//
//	0x000 32  MODER    GPIO port mode register
//	0x004 32  OTYPER   GPIO port output type register
//	0x008 32  OSPEEDR  GPIO port output speed register
//	0x00C 32  PUPDR    GPIO port pull-up/pull-down register
//	0x010 32  IDR      GPIO port input data register
//	0x014 32  ODR      GPIO port output data register
//	0x018 32  BSRR     GPIO port bit set/reset register
//	0x01C 32  LCKR     This register is used to lock the configuration of the port bits when a correct write sequence is applied to bit 16 (LCKK). The value of bits [15:0] is used to lock the configuration of the GPIO. During the write sequence, the value of LCKR[15:0] must not change. When the LOCK sequence has been applied on a port bit, the value of this port bit can no longer be modified until the next MCU reset or peripheral reset.A specific write sequence is used to write to the GPIOx_LCKR register. Only word access (32-bit long) is allowed during this locking sequence.Each lock bit freezes a specific configuration register (control and alternate function registers).
//	0x020 32  AFRL     GPIO alternate function low register
//	0x024 32  AFRH     GPIO alternate function high register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package gpio

const (
	MODE0  MODER = 0x03 << 0  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE1  MODER = 0x03 << 2  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE2  MODER = 0x03 << 4  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE3  MODER = 0x03 << 6  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE4  MODER = 0x03 << 8  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE5  MODER = 0x03 << 10 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE6  MODER = 0x03 << 12 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE7  MODER = 0x03 << 14 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE8  MODER = 0x03 << 16 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE9  MODER = 0x03 << 18 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE10 MODER = 0x03 << 20 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE11 MODER = 0x03 << 22 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE12 MODER = 0x03 << 24 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE13 MODER = 0x03 << 26 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE14 MODER = 0x03 << 28 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
	MODE15 MODER = 0x03 << 30 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O mode.
)

const (
	MODE0n  = 0
	MODE1n  = 2
	MODE2n  = 4
	MODE3n  = 6
	MODE4n  = 8
	MODE5n  = 10
	MODE6n  = 12
	MODE7n  = 14
	MODE8n  = 16
	MODE9n  = 18
	MODE10n = 20
	MODE11n = 22
	MODE12n = 24
	MODE13n = 26
	MODE14n = 28
	MODE15n = 30
)

const (
	OT0  OTYPER = 0x01 << 0  //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT1  OTYPER = 0x01 << 1  //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT2  OTYPER = 0x01 << 2  //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT3  OTYPER = 0x01 << 3  //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT4  OTYPER = 0x01 << 4  //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT5  OTYPER = 0x01 << 5  //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT6  OTYPER = 0x01 << 6  //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT7  OTYPER = 0x01 << 7  //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT8  OTYPER = 0x01 << 8  //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT9  OTYPER = 0x01 << 9  //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT10 OTYPER = 0x01 << 10 //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT11 OTYPER = 0x01 << 11 //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT12 OTYPER = 0x01 << 12 //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT13 OTYPER = 0x01 << 13 //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT14 OTYPER = 0x01 << 14 //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
	OT15 OTYPER = 0x01 << 15 //+ Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output type.
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
	OSPEED0  OSPEEDR = 0x03 << 0  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED1  OSPEEDR = 0x03 << 2  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED2  OSPEEDR = 0x03 << 4  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED3  OSPEEDR = 0x03 << 6  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED4  OSPEEDR = 0x03 << 8  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED5  OSPEEDR = 0x03 << 10 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED6  OSPEEDR = 0x03 << 12 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED7  OSPEEDR = 0x03 << 14 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED8  OSPEEDR = 0x03 << 16 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED9  OSPEEDR = 0x03 << 18 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED10 OSPEEDR = 0x03 << 20 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED11 OSPEEDR = 0x03 << 22 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED12 OSPEEDR = 0x03 << 24 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED13 OSPEEDR = 0x03 << 26 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED14 OSPEEDR = 0x03 << 28 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
	OSPEED15 OSPEEDR = 0x03 << 30 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O output speed. Note: Refer to the device datasheet for the frequency specifications and the power supply and load conditions for each speed.
)

const (
	OSPEED0n  = 0
	OSPEED1n  = 2
	OSPEED2n  = 4
	OSPEED3n  = 6
	OSPEED4n  = 8
	OSPEED5n  = 10
	OSPEED6n  = 12
	OSPEED7n  = 14
	OSPEED8n  = 16
	OSPEED9n  = 18
	OSPEED10n = 20
	OSPEED11n = 22
	OSPEED12n = 24
	OSPEED13n = 26
	OSPEED14n = 28
	OSPEED15n = 30
)

const (
	PUPD0  PUPDR = 0x03 << 0  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD1  PUPDR = 0x03 << 2  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD2  PUPDR = 0x03 << 4  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD3  PUPDR = 0x03 << 6  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD4  PUPDR = 0x03 << 8  //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD5  PUPDR = 0x03 << 10 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD6  PUPDR = 0x03 << 12 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD7  PUPDR = 0x03 << 14 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD8  PUPDR = 0x03 << 16 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD9  PUPDR = 0x03 << 18 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD10 PUPDR = 0x03 << 20 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD11 PUPDR = 0x03 << 22 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD12 PUPDR = 0x03 << 24 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD13 PUPDR = 0x03 << 26 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD14 PUPDR = 0x03 << 28 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
	PUPD15 PUPDR = 0x03 << 30 //+ [1:0]: Port x configuration bits (y = 0..15) These bits are written by software to configure the I/O pull-up or pull-down
)

const (
	PUPD0n  = 0
	PUPD1n  = 2
	PUPD2n  = 4
	PUPD3n  = 6
	PUPD4n  = 8
	PUPD5n  = 10
	PUPD6n  = 12
	PUPD7n  = 14
	PUPD8n  = 16
	PUPD9n  = 18
	PUPD10n = 20
	PUPD11n = 22
	PUPD12n = 24
	PUPD13n = 26
	PUPD14n = 28
	PUPD15n = 30
)

const (
	ID0  IDR = 0x01 << 0  //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID1  IDR = 0x01 << 1  //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID2  IDR = 0x01 << 2  //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID3  IDR = 0x01 << 3  //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID4  IDR = 0x01 << 4  //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID5  IDR = 0x01 << 5  //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID6  IDR = 0x01 << 6  //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID7  IDR = 0x01 << 7  //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID8  IDR = 0x01 << 8  //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID9  IDR = 0x01 << 9  //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID10 IDR = 0x01 << 10 //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID11 IDR = 0x01 << 11 //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID12 IDR = 0x01 << 12 //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID13 IDR = 0x01 << 13 //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID14 IDR = 0x01 << 14 //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
	ID15 IDR = 0x01 << 15 //+ Port input data bit (y = 0..15) These bits are read-only. They contain the input value of the corresponding I/O port.
)

const (
	ID0n  = 0
	ID1n  = 1
	ID2n  = 2
	ID3n  = 3
	ID4n  = 4
	ID5n  = 5
	ID6n  = 6
	ID7n  = 7
	ID8n  = 8
	ID9n  = 9
	ID10n = 10
	ID11n = 11
	ID12n = 12
	ID13n = 13
	ID14n = 14
	ID15n = 15
)

const (
	OD0  ODR = 0x01 << 0  //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD1  ODR = 0x01 << 1  //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD2  ODR = 0x01 << 2  //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD3  ODR = 0x01 << 3  //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD4  ODR = 0x01 << 4  //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD5  ODR = 0x01 << 5  //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD6  ODR = 0x01 << 6  //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD7  ODR = 0x01 << 7  //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD8  ODR = 0x01 << 8  //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD9  ODR = 0x01 << 9  //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD10 ODR = 0x01 << 10 //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD11 ODR = 0x01 << 11 //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD12 ODR = 0x01 << 12 //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD13 ODR = 0x01 << 13 //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD14 ODR = 0x01 << 14 //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
	OD15 ODR = 0x01 << 15 //+ Port output data bit These bits can be read and written by software. Note: For atomic bit set/reset, the OD bits can be individually set and/or reset by writing to the GPIOx_BSRR or GPIOx_BRR registers (x = A..F).
)

const (
	OD0n  = 0
	OD1n  = 1
	OD2n  = 2
	OD3n  = 3
	OD4n  = 4
	OD5n  = 5
	OD6n  = 6
	OD7n  = 7
	OD8n  = 8
	OD9n  = 9
	OD10n = 10
	OD11n = 11
	OD12n = 12
	OD13n = 13
	OD14n = 14
	OD15n = 15
)

const (
	BS0  BSRR = 0x01 << 0  //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS1  BSRR = 0x01 << 1  //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS2  BSRR = 0x01 << 2  //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS3  BSRR = 0x01 << 3  //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS4  BSRR = 0x01 << 4  //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS5  BSRR = 0x01 << 5  //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS6  BSRR = 0x01 << 6  //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS7  BSRR = 0x01 << 7  //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS8  BSRR = 0x01 << 8  //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS9  BSRR = 0x01 << 9  //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS10 BSRR = 0x01 << 10 //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS11 BSRR = 0x01 << 11 //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS12 BSRR = 0x01 << 12 //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS13 BSRR = 0x01 << 13 //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS14 BSRR = 0x01 << 14 //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BS15 BSRR = 0x01 << 15 //+ Port x set bit y (y= 0..15) These bits are write-only. A read to these bits returns the value 0x0000.
	BR0  BSRR = 0x01 << 16 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR1  BSRR = 0x01 << 17 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR2  BSRR = 0x01 << 18 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR3  BSRR = 0x01 << 19 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR4  BSRR = 0x01 << 20 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR5  BSRR = 0x01 << 21 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR6  BSRR = 0x01 << 22 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR7  BSRR = 0x01 << 23 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR8  BSRR = 0x01 << 24 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR9  BSRR = 0x01 << 25 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR10 BSRR = 0x01 << 26 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR11 BSRR = 0x01 << 27 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR12 BSRR = 0x01 << 28 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR13 BSRR = 0x01 << 29 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR14 BSRR = 0x01 << 30 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
	BR15 BSRR = 0x01 << 31 //+ Port x reset bit y (y = 0..15) These bits are write-only. A read to these bits returns the value 0x0000. Note: If both BSx and BRx are set, BSx has priority.
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
	LCK0  LCKR = 0x01 << 0  //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK1  LCKR = 0x01 << 1  //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK2  LCKR = 0x01 << 2  //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK3  LCKR = 0x01 << 3  //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK4  LCKR = 0x01 << 4  //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK5  LCKR = 0x01 << 5  //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK6  LCKR = 0x01 << 6  //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK7  LCKR = 0x01 << 7  //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK8  LCKR = 0x01 << 8  //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK9  LCKR = 0x01 << 9  //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK10 LCKR = 0x01 << 10 //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK11 LCKR = 0x01 << 11 //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK12 LCKR = 0x01 << 12 //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK13 LCKR = 0x01 << 13 //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK14 LCKR = 0x01 << 14 //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCK15 LCKR = 0x01 << 15 //+ Port x lock bit y (y= 0..15) These bits are read/write but can only be written when the LCKK bit is 0.
	LCKK  LCKR = 0x01 << 16 //+ Lock key This bit can be read any time. It can only be modified using the lock key write sequence. LOCK key write sequence: WR LCKR[16] = 1 + LCKR[15:0] WR LCKR[16] = 0 + LCKR[15:0] WR LCKR[16] = 1 + LCKR[15:0] RD LCKR RD LCKR[16] = 1 (this read operation is optional but it confirms that the lock is active) Note: During the LOCK key write sequence, the value of LCK[15:0] must not change. Any error in the lock sequence aborts the lock. After the first lock sequence on any bit of the port, any read access on the LCKK bit will return 1 until the next MCU reset or peripheral reset.
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
	AFSEL0 AFRL = 0x0F << 0  //+ [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
	AFSEL1 AFRL = 0x0F << 4  //+ [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
	AFSEL2 AFRL = 0x0F << 8  //+ [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
	AFSEL3 AFRL = 0x0F << 12 //+ [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
	AFSEL4 AFRL = 0x0F << 16 //+ [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
	AFSEL5 AFRL = 0x0F << 20 //+ [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
	AFSEL6 AFRL = 0x0F << 24 //+ [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
	AFSEL7 AFRL = 0x0F << 28 //+ [3:0]: Alternate function selection for port x pin y (y = 0..7) These bits are written by software to configure alternate function I/Os AFSELy selection:
)

const (
	AFSEL0n = 0
	AFSEL1n = 4
	AFSEL2n = 8
	AFSEL3n = 12
	AFSEL4n = 16
	AFSEL5n = 20
	AFSEL6n = 24
	AFSEL7n = 28
)

const (
	AFSEL8  AFRH = 0x0F << 0  //+ [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
	AFSEL9  AFRH = 0x0F << 4  //+ [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
	AFSEL10 AFRH = 0x0F << 8  //+ [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
	AFSEL11 AFRH = 0x0F << 12 //+ [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
	AFSEL12 AFRH = 0x0F << 16 //+ [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
	AFSEL13 AFRH = 0x0F << 20 //+ [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
	AFSEL14 AFRH = 0x0F << 24 //+ [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
	AFSEL15 AFRH = 0x0F << 28 //+ [3:0]: Alternate function selection for port x pin y (y = 8..15) These bits are written by software to configure alternate function I/Os
)

const (
	AFSEL8n  = 0
	AFSEL9n  = 4
	AFSEL10n = 8
	AFSEL11n = 12
	AFSEL12n = 16
	AFSEL13n = 20
	AFSEL14n = 24
	AFSEL15n = 28
)
