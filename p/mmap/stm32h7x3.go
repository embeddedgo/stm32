// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32h7x3

// Package mmap provides base memory adresses for all peripherals.
package mmap

// AC
const (
	AC_BASE uintptr = 0xE000EF90 // Access control
)

// ADC
const (
	ADC1_BASE         uintptr = 0x40022000 // Analog to Digital Converter
	ADC12_Common_BASE uintptr = 0x40022300 // Analog-to-Digital Converter
	ADC2_BASE         uintptr = 0x40022100 // Analog to Digital Converter
	ADC3_BASE         uintptr = 0x58026000 // Analog to Digital Converter
	ADC3_Common_BASE  uintptr = 0x58026300 // Analog-to-Digital Converter
)

// AXI
const (
	AXI_BASE uintptr = 0x51000000 // AXI interconnect registers
)

// BDMA
const (
	BDMA_BASE uintptr = 0x58025400 // BDMA
)

// CEC
const (
	CEC_BASE uintptr = 0x40006C00 // CEC
)

// COMP1
const (
	COMP1_BASE uintptr = 0x58003800 // COMP1
)

// CRC
const (
	CRC_BASE uintptr = 0x58024C00 // Cryptographic processor
)

// CRS
const (
	CRS_BASE uintptr = 0x40008400 // CRS
)

// CRYP
const (
	CRYP_BASE uintptr = 0x48021000 // Cryptographic processor
)

// DAC
const (
	DAC_BASE uintptr = 0x40007400 // DAC
)

// DCMI
const (
	DCMI_BASE uintptr = 0x48020000 // Digital camera interface
)

// DFSDM
const (
	DFSDM_BASE uintptr = 0x40017000 // Digital filter for sigma delta modulators
)

// DLYB
const (
	DELAY_Block_QUADSPI_BASE uintptr = 0x52006000 // DELAY_Block_SDMMC1
	DELAY_Block_SDMMC1_BASE  uintptr = 0x52008000 // DELAY_Block_SDMMC1
	DELAY_Block_SDMMC2_BASE  uintptr = 0x48022800 // DELAY_Block_SDMMC1
)

// DMA
const (
	DMA1_BASE uintptr = 0x40020000 // DMA controller
	DMA2_BASE uintptr = 0x40020400 // DMA controller
)

// DMA2D
const (
	DMA2D_BASE uintptr = 0x52001000 // DMA2D
)

// DMAMUX
const (
	DMAMUX1_BASE uintptr = 0x40020800 // DMAMUX
	DMAMUX2_BASE uintptr = 0x58025800 // DMAMUX
)

// EXTI
const (
	EXTI_BASE uintptr = 0x58000000 // External interrupt/event controller
)

// Ethernet
const (
	Ethernet_DMA_BASE uintptr = 0x40029000 // Ethernet: DMA mode register (DMA)
	Ethernet_MAC_BASE uintptr = 0x40028000 // Ethernet: media access control (MAC)
	Ethernet_MTL_BASE uintptr = 0x40028C00 // Ethernet: MTL mode register (MTL)
)

// FDCAN
const (
	CAN_CCU_BASE uintptr = 0x4000A800 // CCU registers
	FDCAN1_BASE  uintptr = 0x4000A000 // FDCAN1
	FDCAN2_BASE  uintptr = 0x4000A400 // FDCAN1
)

// FMC
const (
	FMC_BASE uintptr = 0x52004000 // FMC
)

// FPU
const (
	FPU_BASE       uintptr = 0xE000EF34 // Floting point unit
	FPU_CPACR_BASE uintptr = 0xE000ED88 // Floating point unit CPACR
)

// Flash
const (
	Flash_BASE uintptr = 0x52002000 // Flash
)

// GPIO
const (
	GPIOA_BASE uintptr = 0x58020000 // GPIO
	GPIOB_BASE uintptr = 0x58020400 // GPIO
	GPIOC_BASE uintptr = 0x58020800 // GPIO
	GPIOD_BASE uintptr = 0x58020C00 // GPIO
	GPIOE_BASE uintptr = 0x58021000 // GPIO
	GPIOF_BASE uintptr = 0x58021400 // GPIO
	GPIOG_BASE uintptr = 0x58021800 // GPIO
	GPIOH_BASE uintptr = 0x58021C00 // GPIO
	GPIOI_BASE uintptr = 0x58022000 // GPIO
	GPIOJ_BASE uintptr = 0x58022400 // GPIO
	GPIOK_BASE uintptr = 0x58022800 // GPIO
)

// HASH
const (
	HASH_BASE uintptr = 0x48021400 // Hash processor
)

// HRTIM
const (
	HRTIM_Common_BASE uintptr = 0x40017780 // High Resolution Timer: Common functions
	HRTIM_Master_BASE uintptr = 0x40017400 // High Resolution Timer: Master Timers
	HRTIM_TIMA_BASE   uintptr = 0x40017480 // High Resolution Timer: TIMA
	HRTIM_TIMB_BASE   uintptr = 0x40017500 // High Resolution Timer: TIMB
	HRTIM_TIMC_BASE   uintptr = 0x40017580 // High Resolution Timer: TIMC
	HRTIM_TIMD_BASE   uintptr = 0x40017600 // High Resolution Timer: TIMD
	HRTIM_TIME_BASE   uintptr = 0x40017680 // High Resolution Timer: TIME
)

// HSEM
const (
	HSEM_BASE uintptr = 0x58026400 // HSEM
)

// I2C
const (
	I2C1_BASE uintptr = 0x40005400 // I2C
	I2C2_BASE uintptr = 0x40005800 // I2C
	I2C3_BASE uintptr = 0x40005C00 // I2C
	I2C4_BASE uintptr = 0x58001C00 // I2C
)

// IWDG
const (
	IWDG_BASE uintptr = 0x58004800 // IWDG
)

// JPEG
const (
	JPEG_BASE uintptr = 0x52003000 // JPEG
)

// LPTIM
const (
	LPTIM1_BASE uintptr = 0x40002400 // Low power timer
	LPTIM2_BASE uintptr = 0x58002400 // Low power timer
	LPTIM3_BASE uintptr = 0x58002800 // Low power timer
	LPTIM4_BASE uintptr = 0x58002C00 // Low power timer
	LPTIM5_BASE uintptr = 0x58003000 // Low power timer
)

// LPUART
const (
	LPUART1_BASE uintptr = 0x58000C00 // LPUART1
)

// LTDC
const (
	LTDC_BASE uintptr = 0x50001000 // LCD-TFT Controller
)

// MDIOS
const (
	MDIOS_BASE uintptr = 0x40009400 // Management data input/output slave
)

// MDMA
const (
	MDMA_BASE uintptr = 0x52000000 // MDMA
)

// MPU
const (
	MPU_BASE uintptr = 0xE000ED90 // Memory protection unit
)

// NVIC
const (
	NVIC_BASE      uintptr = 0xE000E100 // Nested Vectored Interrupt Controller
	NVIC_STIR_BASE uintptr = 0xE000EF00 // Nested vectored interrupt controller
)

// OPAMP
const (
	OPAMP_BASE uintptr = 0x40009000 // Operational amplifiers
)

// PF
const (
	PF_BASE uintptr = 0xE000ED78 // Processor features
)

// PWR
const (
	PWR_BASE uintptr = 0x58024800 // PWR
)

// QUADSPI
const (
	QUADSPI_BASE uintptr = 0x52005000 // QUADSPI
)

// RCC
const (
	RCC_BASE uintptr = 0x58024400 // Reset and clock control
)

// RNG
const (
	RNG_BASE uintptr = 0x48021800 // RNG
)

// RTC
const (
	RTC_BASE uintptr = 0x58004000 // RTC
)

// SAI
const (
	SAI1_BASE uintptr = 0x40015800 // SAI
	SAI2_BASE uintptr = 0x40015C00 // SAI
	SAI3_BASE uintptr = 0x40016000 // SAI
	SAI4_BASE uintptr = 0x58005400 // SAI
)

// SCB
const (
	SCB_BASE       uintptr = 0xE000ED00 // System control block
	SCB_ACTRL_BASE uintptr = 0xE000E008 // System control block ACTLR
)

// SDMMC
const (
	SDMMC1_BASE uintptr = 0x52007000 // SDMMC1
	SDMMC2_BASE uintptr = 0x48022400 // SDMMC1
)

// SPDIFRX
const (
	SPDIFRX_BASE uintptr = 0x40004000 // Receiver Interface
)

// SPI
const (
	SPI1_BASE uintptr = 0x40013000 // Serial peripheral interface
	SPI2_BASE uintptr = 0x40003800 // Serial peripheral interface
	SPI3_BASE uintptr = 0x40003C00 // Serial peripheral interface
	SPI4_BASE uintptr = 0x40013400 // Serial peripheral interface
	SPI5_BASE uintptr = 0x40015000 // Serial peripheral interface
	SPI6_BASE uintptr = 0x58001400 // Serial peripheral interface
)

// STK
const (
	STK_BASE uintptr = 0xE000E010 // SysTick timer
)

// SWPMI
const (
	SWPMI_BASE uintptr = 0x40008800 // Single Wire Protocol Master Interface
)

// SYSCFG
const (
	SYSCFG_BASE uintptr = 0x58000400 // System configuration controller
)

// TIM
const (
	TIM1_BASE  uintptr = 0x40010000 // Advanced-timers
	TIM2_BASE  uintptr = 0x40000000 // General purpose timers
	TIM3_BASE  uintptr = 0x40000400 // General purpose timers
	TIM4_BASE  uintptr = 0x40000800 // General purpose timers
	TIM5_BASE  uintptr = 0x40000C00 // General purpose timers
	TIM6_BASE  uintptr = 0x40001000 // Basic timers
	TIM7_BASE  uintptr = 0x40001400 // Basic timers
	TIM8_BASE  uintptr = 0x40010400 // Advanced-timers
	TIM12_BASE uintptr = 0x40001800 // General purpose timers
	TIM13_BASE uintptr = 0x40001C00 // General purpose timers
	TIM14_BASE uintptr = 0x40002000 // General purpose timers
)

// TIMs
const (
	TIM15_BASE uintptr = 0x40014000 // General purpose timers
	TIM16_BASE uintptr = 0x40014400 // General-purpose-timers
	TIM17_BASE uintptr = 0x40014800 // General-purpose-timers
)

// USART
const (
	UART4_BASE  uintptr = 0x40004C00 // Universal synchronous asynchronous receiver transmitter
	UART5_BASE  uintptr = 0x40005000 // Universal synchronous asynchronous receiver transmitter
	UART7_BASE  uintptr = 0x40007800 // Universal synchronous asynchronous receiver transmitter
	UART8_BASE  uintptr = 0x40007C00 // Universal synchronous asynchronous receiver transmitter
	USART1_BASE uintptr = 0x40011000 // Universal synchronous asynchronous receiver transmitter
	USART2_BASE uintptr = 0x40004400 // Universal synchronous asynchronous receiver transmitter
	USART3_BASE uintptr = 0x40004800 // Universal synchronous asynchronous receiver transmitter
	USART6_BASE uintptr = 0x40011400 // Universal synchronous asynchronous receiver transmitter
)

// USB_OTG_HS
const (
	OTG1_HS_DEVICE_BASE uintptr = 0x40040800 // USB 1 on the go high speed
	OTG1_HS_GLOBAL_BASE uintptr = 0x40040000 // USB 1 on the go high speed
	OTG1_HS_HOST_BASE   uintptr = 0x40040400 // USB 1 on the go high speed
	OTG1_HS_PWRCLK_BASE uintptr = 0x40040E00 // USB 1 on the go high speed
	OTG2_HS_DEVICE_BASE uintptr = 0x40080800 // USB 1 on the go high speed
	OTG2_HS_GLOBAL_BASE uintptr = 0x40080000 // USB 1 on the go high speed
	OTG2_HS_HOST_BASE   uintptr = 0x40080400 // USB 1 on the go high speed
	OTG2_HS_PWRCLK_BASE uintptr = 0x40080E00 // USB 1 on the go high speed
)

// VREFBUF
const (
	VREFBUF_BASE uintptr = 0x58003C00 // VREFBUF
)

// WWDG
const (
	WWDG_BASE uintptr = 0x50003000 // WWDG
)