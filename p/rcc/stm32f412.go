// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f412

// Package rcc provides access to the registers of the RCC peripheral.
//
// Instances:
//
//	RCC  RCC_BASE  -  RCC  Reset and clock control
//
// Registers:
//
//	0x000 32  CR          clock control register
//	0x004 32  PLLCFGR     PLL configuration register
//	0x008 32  CFGR        clock configuration register
//	0x00C 32  CIR         clock interrupt register
//	0x010 32  AHB1RSTR    AHB1 peripheral reset register
//	0x014 32  AHB2RSTR    AHB2 peripheral reset register
//	0x020 32  APB1RSTR    APB1 peripheral reset register
//	0x024 32  APB2RSTR    APB2 peripheral reset register
//	0x030 32  AHB1ENR     AHB1 peripheral clock register
//	0x034 32  AHB2ENR     AHB2 peripheral clock enable register
//	0x040 32  APB1ENR     APB1 peripheral clock enable register
//	0x044 32  APB2ENR     APB2 peripheral clock enable register
//	0x050 32  AHB1LPENR   AHB1 peripheral clock enable in low power mode register
//	0x054 32  AHB2LPENR   AHB2 peripheral clock enable in low power mode register
//	0x060 32  APB1LPENR   APB1 peripheral clock enable in low power mode register
//	0x064 32  APB2LPENR   APB2 peripheral clock enabled in low power mode register
//	0x070 32  BDCR        Backup domain control register
//	0x074 32  CSR         clock control & status register
//	0x080 32  SSCGR       spread spectrum clock generation register
//	0x084 32  PLLI2SCFGR  PLLI2S configuration register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package rcc

const (
	HSION     CR = 0x01 << 0  //+ Internal high-speed clock enable
	HSIRDY    CR = 0x01 << 1  //+ Internal high-speed clock ready flag
	HSITRIM   CR = 0x1F << 3  //+ Internal high-speed clock trimming
	HSICAL    CR = 0xFF << 8  //+ Internal high-speed clock calibration
	HSEON     CR = 0x01 << 16 //+ HSE clock enable
	HSERDY    CR = 0x01 << 17 //+ HSE clock ready flag
	HSEBYP    CR = 0x01 << 18 //+ HSE clock bypass
	CSSON     CR = 0x01 << 19 //+ Clock security system enable
	PLLON     CR = 0x01 << 24 //+ Main PLL (PLL) enable
	PLLRDY    CR = 0x01 << 25 //+ Main PLL (PLL) clock ready flag
	PLLI2SON  CR = 0x01 << 26 //+ PLLI2S enable
	PLLI2SRDY CR = 0x01 << 27 //+ PLLI2S clock ready flag
)

const (
	HSIONn     = 0
	HSIRDYn    = 1
	HSITRIMn   = 3
	HSICALn    = 8
	HSEONn     = 16
	HSERDYn    = 17
	HSEBYPn    = 18
	CSSONn     = 19
	PLLONn     = 24
	PLLRDYn    = 25
	PLLI2SONn  = 26
	PLLI2SRDYn = 27
)

const (
	PLLM       PLLCFGR = 0x3F << 0  //+ Division factor for the main PLL (PLL) and audio PLL (PLLI2S) input clock
	PLLN       PLLCFGR = 0x1FF << 6 //+ Main PLL (PLL) multiplication factor for VCO
	PLLP       PLLCFGR = 0x03 << 16 //+ Main PLL (PLL) division factor for main system clock
	PLLSRC     PLLCFGR = 0x01 << 22 //+ Main PLL(PLL) and audio PLL (PLLI2S) entry clock source
	PLLSRC_HSI PLLCFGR = 0x00 << 22 //  HSI usead as clock source
	PLLSRC_HSE PLLCFGR = 0x01 << 22 //  HSE usead as clock source
	PLLQ       PLLCFGR = 0x0F << 24 //+ Main PLL (PLL) division factor for USB OTG FS, SDIO and random number generator clocks
)

const (
	PLLMn   = 0
	PLLNn   = 6
	PLLPn   = 16
	PLLSRCn = 22
	PLLQn   = 24
)

const (
	SW          CFGR = 0x03 << 0  //+ System clock switch
	SW_HSI      CFGR = 0x00 << 0  //  HSI selected as system clock
	SW_HSE      CFGR = 0x01 << 0  //  HSE selected as system clock
	SW_PLL      CFGR = 0x02 << 0  //  PLL/PLLP selected as system clock
	SWS         CFGR = 0x03 << 2  //+ System clock switch status
	SWS_HSI     CFGR = 0x00 << 2  //  HSI oscillator used as system clock
	SWS_HSE     CFGR = 0x01 << 2  //  HSE oscillator used as system clock
	SWS_PLL     CFGR = 0x02 << 2  //  PLL/PLLP used as system clock
	HPRE        CFGR = 0x0F << 4  //+ AHB prescaler
	HPRE_DIV1   CFGR = 0x00 << 4  //  SYSCLK not divided
	HPRE_DIV2   CFGR = 0x08 << 4  //  SYSCLK divided by 2
	HPRE_DIV4   CFGR = 0x09 << 4  //  SYSCLK divided by 4
	HPRE_DIV8   CFGR = 0x0A << 4  //  SYSCLK divided by 8
	HPRE_DIV16  CFGR = 0x0B << 4  //  SYSCLK divided by 16
	HPRE_DIV64  CFGR = 0x0C << 4  //  SYSCLK divided by 64
	HPRE_DIV128 CFGR = 0x0D << 4  //  SYSCLK divided by 128
	HPRE_DIV256 CFGR = 0x0E << 4  //  SYSCLK divided by 256
	HPRE_DIV512 CFGR = 0x0F << 4  //  SYSCLK divided by 512
	PPRE1       CFGR = 0x07 << 10 //+ APB Low speed prescaler (APB1)
	PPRE1_DIV1  CFGR = 0x00 << 10 //  HCLK not divided
	PPRE1_DIV2  CFGR = 0x04 << 10 //  HCLK divided by 2
	PPRE1_DIV4  CFGR = 0x05 << 10 //  HCLK divided by 4
	PPRE1_DIV8  CFGR = 0x06 << 10 //  HCLK divided by 8
	PPRE1_DIV16 CFGR = 0x07 << 10 //  HCLK divided by 16
	PPRE2       CFGR = 0x07 << 13 //+ APB high-speed prescaler (APB2)
	PPRE2_DIV1  CFGR = 0x00 << 13 //  HCLK not divided
	PPRE2_DIV2  CFGR = 0x04 << 13 //  HCLK divided by 2
	PPRE2_DIV4  CFGR = 0x05 << 13 //  HCLK divided by 4
	PPRE2_DIV8  CFGR = 0x06 << 13 //  HCLK divided by 8
	PPRE2_DIV16 CFGR = 0x07 << 13 //  HCLK divided by 16
	RTCPRE      CFGR = 0x1F << 16 //+ HSE division factor for RTC clock
	MCO1        CFGR = 0x03 << 21 //+ Microcontroller clock output 1
	I2SSRC      CFGR = 0x01 << 23 //+ I2S clock selection
	MCO1PRE     CFGR = 0x07 << 24 //+ MCO1 prescaler
	MCO2PRE     CFGR = 0x07 << 27 //+ MCO2 prescaler
	MCO2        CFGR = 0x03 << 30 //+ Microcontroller clock output 2
)

const (
	SWn      = 0
	SWSn     = 2
	HPREn    = 4
	PPRE1n   = 10
	PPRE2n   = 13
	RTCPREn  = 16
	MCO1n    = 21
	I2SSRCn  = 23
	MCO1PREn = 24
	MCO2PREn = 27
	MCO2n    = 30
)

const (
	LSIRDYF     CIR = 0x01 << 0  //+ LSI ready interrupt flag
	LSERDYF     CIR = 0x01 << 1  //+ LSE ready interrupt flag
	HSIRDYF     CIR = 0x01 << 2  //+ HSI ready interrupt flag
	HSERDYF     CIR = 0x01 << 3  //+ HSE ready interrupt flag
	PLLRDYF     CIR = 0x01 << 4  //+ Main PLL (PLL) ready interrupt flag
	PLLI2SRDYF  CIR = 0x01 << 5  //+ PLLI2S ready interrupt flag
	CSSF        CIR = 0x01 << 7  //+ Clock security system interrupt flag
	LSIRDYIE    CIR = 0x01 << 8  //+ LSI ready interrupt enable
	LSERDYIE    CIR = 0x01 << 9  //+ LSE ready interrupt enable
	HSIRDYIE    CIR = 0x01 << 10 //+ HSI ready interrupt enable
	HSERDYIE    CIR = 0x01 << 11 //+ HSE ready interrupt enable
	PLLRDYIE    CIR = 0x01 << 12 //+ Main PLL (PLL) ready interrupt enable
	PLLI2SRDYIE CIR = 0x01 << 13 //+ PLLI2S ready interrupt enable
	LSIRDYC     CIR = 0x01 << 16 //+ LSI ready interrupt clear
	LSERDYC     CIR = 0x01 << 17 //+ LSE ready interrupt clear
	HSIRDYC     CIR = 0x01 << 18 //+ HSI ready interrupt clear
	HSERDYC     CIR = 0x01 << 19 //+ HSE ready interrupt clear
	PLLRDYC     CIR = 0x01 << 20 //+ Main PLL(PLL) ready interrupt clear
	PLLI2SRDYC  CIR = 0x01 << 21 //+ PLLI2S ready interrupt clear
	CSSC        CIR = 0x01 << 23 //+ Clock security system interrupt clear
)

const (
	LSIRDYFn     = 0
	LSERDYFn     = 1
	HSIRDYFn     = 2
	HSERDYFn     = 3
	PLLRDYFn     = 4
	PLLI2SRDYFn  = 5
	CSSFn        = 7
	LSIRDYIEn    = 8
	LSERDYIEn    = 9
	HSIRDYIEn    = 10
	HSERDYIEn    = 11
	PLLRDYIEn    = 12
	PLLI2SRDYIEn = 13
	LSIRDYCn     = 16
	LSERDYCn     = 17
	HSIRDYCn     = 18
	HSERDYCn     = 19
	PLLRDYCn     = 20
	PLLI2SRDYCn  = 21
	CSSCn        = 23
)

const (
	GPIOARST AHB1RSTR = 0x01 << 0  //+ IO port A reset
	GPIOBRST AHB1RSTR = 0x01 << 1  //+ IO port B reset
	GPIOCRST AHB1RSTR = 0x01 << 2  //+ IO port C reset
	GPIODRST AHB1RSTR = 0x01 << 3  //+ IO port D reset
	GPIOERST AHB1RSTR = 0x01 << 4  //+ IO port E reset
	GPIOFRST AHB1RSTR = 0x01 << 5  //+ IO port F reset
	GPIOGRST AHB1RSTR = 0x01 << 6  //+ IO port G reset
	GPIOHRST AHB1RSTR = 0x01 << 7  //+ IO port H reset
	CRCRST   AHB1RSTR = 0x01 << 12 //+ CRC reset
	DMA1RST  AHB1RSTR = 0x01 << 21 //+ DMA2 reset
	DMA2RST  AHB1RSTR = 0x01 << 22 //+ DMA2 reset
)

const (
	GPIOARSTn = 0
	GPIOBRSTn = 1
	GPIOCRSTn = 2
	GPIODRSTn = 3
	GPIOERSTn = 4
	GPIOFRSTn = 5
	GPIOGRSTn = 6
	GPIOHRSTn = 7
	CRCRSTn   = 12
	DMA1RSTn  = 21
	DMA2RSTn  = 22
)

const (
	RNGRST   AHB2RSTR = 0x01 << 6 //+ RNGRST
	OTGFSRST AHB2RSTR = 0x01 << 7 //+ USB OTG FS module reset
)

const (
	RNGRSTn   = 6
	OTGFSRSTn = 7
)

const (
	TIM2RST   APB1RSTR = 0x01 << 0  //+ TIM2 reset
	TIM3RST   APB1RSTR = 0x01 << 1  //+ TIM3 reset
	TIM4RST   APB1RSTR = 0x01 << 2  //+ TIM4 reset
	TIM5RST   APB1RSTR = 0x01 << 3  //+ TIM5 reset
	TIM6RST   APB1RSTR = 0x01 << 4  //+ TIM6RST
	TIM7RST   APB1RSTR = 0x01 << 5  //+ TIM7RST
	TIM12RST  APB1RSTR = 0x01 << 6  //+ TIM12RST
	TIM13RST  APB1RSTR = 0x01 << 7  //+ TIM13RST
	TIM14RST  APB1RSTR = 0x01 << 8  //+ TIM14RST
	WWDGRST   APB1RSTR = 0x01 << 11 //+ Window watchdog reset
	SPI2RST   APB1RSTR = 0x01 << 14 //+ SPI 2 reset
	SPI3RST   APB1RSTR = 0x01 << 15 //+ SPI 3 reset
	UART2RST  APB1RSTR = 0x01 << 17 //+ USART 2 reset
	USART3RST APB1RSTR = 0x01 << 18 //+ USART3RST
	I2C1RST   APB1RSTR = 0x01 << 21 //+ I2C 1 reset
	I2C2RST   APB1RSTR = 0x01 << 22 //+ I2C 2 reset
	I2C3RST   APB1RSTR = 0x01 << 23 //+ I2C3 reset
	I2C4RST   APB1RSTR = 0x01 << 24 //+ I2C4RST
	CAN1RST   APB1RSTR = 0x01 << 25 //+ CAN1RST
	CAN2RST   APB1RSTR = 0x01 << 26 //+ CAN2RST
	PWRRST    APB1RSTR = 0x01 << 28 //+ Power interface reset
)

const (
	TIM2RSTn   = 0
	TIM3RSTn   = 1
	TIM4RSTn   = 2
	TIM5RSTn   = 3
	TIM6RSTn   = 4
	TIM7RSTn   = 5
	TIM12RSTn  = 6
	TIM13RSTn  = 7
	TIM14RSTn  = 8
	WWDGRSTn   = 11
	SPI2RSTn   = 14
	SPI3RSTn   = 15
	UART2RSTn  = 17
	USART3RSTn = 18
	I2C1RSTn   = 21
	I2C2RSTn   = 22
	I2C3RSTn   = 23
	I2C4RSTn   = 24
	CAN1RSTn   = 25
	CAN2RSTn   = 26
	PWRRSTn    = 28
)

const (
	TIM1RST   APB2RSTR = 0x01 << 0  //+ TIM1 reset
	TIM8RST   APB2RSTR = 0x01 << 1  //+ TIM8RST
	USART1RST APB2RSTR = 0x01 << 4  //+ USART1 reset
	USART6RST APB2RSTR = 0x01 << 5  //+ USART6 reset
	ADCRST    APB2RSTR = 0x01 << 8  //+ ADC interface reset (common to all ADCs)
	SDIORST   APB2RSTR = 0x01 << 11 //+ SDIO reset
	SPI1RST   APB2RSTR = 0x01 << 12 //+ SPI 1 reset
	SYSCFGRST APB2RSTR = 0x01 << 14 //+ System configuration controller reset
	TIM9RST   APB2RSTR = 0x01 << 16 //+ TIM9 reset
	TIM10RST  APB2RSTR = 0x01 << 17 //+ TIM10 reset
	TIM11RST  APB2RSTR = 0x01 << 18 //+ TIM11 reset
	DFSDMRST  APB2RSTR = 0x01 << 24 //+ DFSDMRST
)

const (
	TIM1RSTn   = 0
	TIM8RSTn   = 1
	USART1RSTn = 4
	USART6RSTn = 5
	ADCRSTn    = 8
	SDIORSTn   = 11
	SPI1RSTn   = 12
	SYSCFGRSTn = 14
	TIM9RSTn   = 16
	TIM10RSTn  = 17
	TIM11RSTn  = 18
	DFSDMRSTn  = 24
)

const (
	GPIOAEN AHB1ENR = 0x01 << 0  //+ IO port A clock enable
	GPIOBEN AHB1ENR = 0x01 << 1  //+ IO port B clock enable
	GPIOCEN AHB1ENR = 0x01 << 2  //+ IO port C clock enable
	GPIODEN AHB1ENR = 0x01 << 3  //+ IO port D clock enable
	GPIOEEN AHB1ENR = 0x01 << 4  //+ IO port E clock enable
	GPIOFEN AHB1ENR = 0x01 << 5  //+ IO port F clock enable
	GPIOGEN AHB1ENR = 0x01 << 6  //+ IO port G clock enable
	GPIOHEN AHB1ENR = 0x01 << 7  //+ IO port H clock enable
	CRCEN   AHB1ENR = 0x01 << 12 //+ CRC clock enable
	DMA1EN  AHB1ENR = 0x01 << 21 //+ DMA1 clock enable
	DMA2EN  AHB1ENR = 0x01 << 22 //+ DMA2 clock enable
)

const (
	GPIOAENn = 0
	GPIOBENn = 1
	GPIOCENn = 2
	GPIODENn = 3
	GPIOEENn = 4
	GPIOFENn = 5
	GPIOGENn = 6
	GPIOHENn = 7
	CRCENn   = 12
	DMA1ENn  = 21
	DMA2ENn  = 22
)

const (
	RNGEN   AHB2ENR = 0x01 << 6 //+ RNGEN
	OTGFSEN AHB2ENR = 0x01 << 7 //+ USB OTG FS clock enable
)

const (
	RNGENn   = 6
	OTGFSENn = 7
)

const (
	TIM2EN   APB1ENR = 0x01 << 0  //+ TIM2 clock enable
	TIM3EN   APB1ENR = 0x01 << 1  //+ TIM3 clock enable
	TIM4EN   APB1ENR = 0x01 << 2  //+ TIM4 clock enable
	TIM5EN   APB1ENR = 0x01 << 3  //+ TIM5 clock enable
	TIM6EN   APB1ENR = 0x01 << 4  //+ TIM6EN
	TIM7EN   APB1ENR = 0x01 << 5  //+ TIM7EN
	TIM12EN  APB1ENR = 0x01 << 6  //+ TIM12EN
	TIM13EN  APB1ENR = 0x01 << 7  //+ TIM13EN
	TIM14EN  APB1ENR = 0x01 << 8  //+ TIM14EN
	WWDGEN   APB1ENR = 0x01 << 11 //+ Window watchdog clock enable
	SPI2EN   APB1ENR = 0x01 << 14 //+ SPI2 clock enable
	SPI3EN   APB1ENR = 0x01 << 15 //+ SPI3 clock enable
	USART2EN APB1ENR = 0x01 << 17 //+ USART 2 clock enable
	USART3EN APB1ENR = 0x01 << 18 //+ USART3EN
	I2C1EN   APB1ENR = 0x01 << 21 //+ I2C1 clock enable
	I2C2EN   APB1ENR = 0x01 << 22 //+ I2C2 clock enable
	I2C3EN   APB1ENR = 0x01 << 23 //+ I2C3 clock enable
	I2C4EN   APB1ENR = 0x01 << 24 //+ I2C4EN
	CAN1EN   APB1ENR = 0x01 << 25 //+ CAN1EN
	CAN2EN   APB1ENR = 0x01 << 26 //+ CAN2EN
	PWREN    APB1ENR = 0x01 << 28 //+ Power interface clock enable
)

const (
	TIM2ENn   = 0
	TIM3ENn   = 1
	TIM4ENn   = 2
	TIM5ENn   = 3
	TIM6ENn   = 4
	TIM7ENn   = 5
	TIM12ENn  = 6
	TIM13ENn  = 7
	TIM14ENn  = 8
	WWDGENn   = 11
	SPI2ENn   = 14
	SPI3ENn   = 15
	USART2ENn = 17
	USART3ENn = 18
	I2C1ENn   = 21
	I2C2ENn   = 22
	I2C3ENn   = 23
	I2C4ENn   = 24
	CAN1ENn   = 25
	CAN2ENn   = 26
	PWRENn    = 28
)

const (
	TIM1EN   APB2ENR = 0x01 << 0  //+ TIM1 clock enable
	TIM8EN   APB2ENR = 0x01 << 1  //+ TIM8EN
	USART1EN APB2ENR = 0x01 << 4  //+ USART1 clock enable
	USART6EN APB2ENR = 0x01 << 5  //+ USART6 clock enable
	ADC1EN   APB2ENR = 0x01 << 8  //+ ADC1 clock enable
	SDIOEN   APB2ENR = 0x01 << 11 //+ SDIO clock enable
	SPI1EN   APB2ENR = 0x01 << 12 //+ SPI1 clock enable
	SPI4EN   APB2ENR = 0x01 << 13 //+ SPI4 clock enable
	SYSCFGEN APB2ENR = 0x01 << 14 //+ System configuration controller clock enable
	TIM9EN   APB2ENR = 0x01 << 16 //+ TIM9 clock enable
	TIM10EN  APB2ENR = 0x01 << 17 //+ TIM10 clock enable
	TIM11EN  APB2ENR = 0x01 << 18 //+ TIM11 clock enable
	DFSDMEN  APB2ENR = 0x01 << 24 //+ DFSDMEN
)

const (
	TIM1ENn   = 0
	TIM8ENn   = 1
	USART1ENn = 4
	USART6ENn = 5
	ADC1ENn   = 8
	SDIOENn   = 11
	SPI1ENn   = 12
	SPI4ENn   = 13
	SYSCFGENn = 14
	TIM9ENn   = 16
	TIM10ENn  = 17
	TIM11ENn  = 18
	DFSDMENn  = 24
)

const (
	GPIOALPEN AHB1LPENR = 0x01 << 0  //+ IO port A clock enable during sleep mode
	GPIOBLPEN AHB1LPENR = 0x01 << 1  //+ IO port B clock enable during Sleep mode
	GPIOCLPEN AHB1LPENR = 0x01 << 2  //+ IO port C clock enable during Sleep mode
	GPIODLPEN AHB1LPENR = 0x01 << 3  //+ IO port D clock enable during Sleep mode
	GPIOELPEN AHB1LPENR = 0x01 << 4  //+ IO port E clock enable during Sleep mode
	GPIOFLPEN AHB1LPENR = 0x01 << 5  //+ IO port F clock enable during sleep mode
	GPIOGLPEN AHB1LPENR = 0x01 << 6  //+ IO port G clock enable during sleep mode
	GPIOHLPEN AHB1LPENR = 0x01 << 7  //+ IO port H clock enable during Sleep mode
	CRCLPEN   AHB1LPENR = 0x01 << 12 //+ CRC clock enable during Sleep mode
	FLITFLPEN AHB1LPENR = 0x01 << 15 //+ Flash interface clock enable during Sleep mode
	SRAM1LPEN AHB1LPENR = 0x01 << 16 //+ SRAM 1interface clock enable during Sleep mode
	DMA1LPEN  AHB1LPENR = 0x01 << 21 //+ DMA1 clock enable during Sleep mode
	DMA2LPEN  AHB1LPENR = 0x01 << 22 //+ DMA2 clock enable during Sleep mode
)

const (
	GPIOALPENn = 0
	GPIOBLPENn = 1
	GPIOCLPENn = 2
	GPIODLPENn = 3
	GPIOELPENn = 4
	GPIOFLPENn = 5
	GPIOGLPENn = 6
	GPIOHLPENn = 7
	CRCLPENn   = 12
	FLITFLPENn = 15
	SRAM1LPENn = 16
	DMA1LPENn  = 21
	DMA2LPENn  = 22
)

const (
	RNGLPEN   AHB2LPENR = 0x01 << 6 //+ RNGLPEN
	OTGFSLPEN AHB2LPENR = 0x01 << 7 //+ USB OTG FS clock enable during Sleep mode
)

const (
	RNGLPENn   = 6
	OTGFSLPENn = 7
)

const (
	TIM2LPEN   APB1LPENR = 0x01 << 0  //+ TIM2 clock enable during Sleep mode
	TIM3LPEN   APB1LPENR = 0x01 << 1  //+ TIM3 clock enable during Sleep mode
	TIM4LPEN   APB1LPENR = 0x01 << 2  //+ TIM4 clock enable during Sleep mode
	TIM5LPEN   APB1LPENR = 0x01 << 3  //+ TIM5 clock enable during Sleep mode
	TIM6LPEN   APB1LPENR = 0x01 << 4  //+ TIM6LPEN
	TIM7LPEN   APB1LPENR = 0x01 << 5  //+ TIM7LPEN
	TIM12LPEN  APB1LPENR = 0x01 << 6  //+ TIM12LPEN
	TIM13LPEN  APB1LPENR = 0x01 << 7  //+ TIM13LPEN
	TIM14LPEN  APB1LPENR = 0x01 << 8  //+ TIM14LPEN
	WWDGLPEN   APB1LPENR = 0x01 << 11 //+ Window watchdog clock enable during Sleep mode
	SPI2LPEN   APB1LPENR = 0x01 << 14 //+ SPI2 clock enable during Sleep mode
	SPI3LPEN   APB1LPENR = 0x01 << 15 //+ SPI3 clock enable during Sleep mode
	USART2LPEN APB1LPENR = 0x01 << 17 //+ USART2 clock enable during Sleep mode
	USART3LPEN APB1LPENR = 0x01 << 18 //+ USART3LPEN
	I2C1LPEN   APB1LPENR = 0x01 << 21 //+ I2C1 clock enable during Sleep mode
	I2C2LPEN   APB1LPENR = 0x01 << 22 //+ I2C2 clock enable during Sleep mode
	I2C3LPEN   APB1LPENR = 0x01 << 23 //+ I2C3 clock enable during Sleep mode
	I2C4LPEN   APB1LPENR = 0x01 << 24 //+ I2C4LPEN
	CAN1LPEN   APB1LPENR = 0x01 << 25 //+ CAN1LPEN
	CAN2LPEN   APB1LPENR = 0x01 << 26 //+ CAN2LPEN
	PWRLPEN    APB1LPENR = 0x01 << 28 //+ Power interface clock enable during Sleep mode
)

const (
	TIM2LPENn   = 0
	TIM3LPENn   = 1
	TIM4LPENn   = 2
	TIM5LPENn   = 3
	TIM6LPENn   = 4
	TIM7LPENn   = 5
	TIM12LPENn  = 6
	TIM13LPENn  = 7
	TIM14LPENn  = 8
	WWDGLPENn   = 11
	SPI2LPENn   = 14
	SPI3LPENn   = 15
	USART2LPENn = 17
	USART3LPENn = 18
	I2C1LPENn   = 21
	I2C2LPENn   = 22
	I2C3LPENn   = 23
	I2C4LPENn   = 24
	CAN1LPENn   = 25
	CAN2LPENn   = 26
	PWRLPENn    = 28
)

const (
	TIM1LPEN   APB2LPENR = 0x01 << 0  //+ TIM1 clock enable during Sleep mode
	TIM8LPEN   APB2LPENR = 0x01 << 1  //+ TIM8LPEN
	USART1LPEN APB2LPENR = 0x01 << 4  //+ USART1 clock enable during Sleep mode
	USART6LPEN APB2LPENR = 0x01 << 5  //+ USART6 clock enable during Sleep mode
	ADC1LPEN   APB2LPENR = 0x01 << 8  //+ ADC1 clock enable during Sleep mode
	SDIOLPEN   APB2LPENR = 0x01 << 11 //+ SDIO clock enable during Sleep mode
	SPI1LPEN   APB2LPENR = 0x01 << 12 //+ SPI 1 clock enable during Sleep mode
	SPI4LPEN   APB2LPENR = 0x01 << 13 //+ SPI4 clock enable during Sleep mode
	SYSCFGLPEN APB2LPENR = 0x01 << 14 //+ System configuration controller clock enable during Sleep mode
	TIM9LPEN   APB2LPENR = 0x01 << 16 //+ TIM9 clock enable during sleep mode
	TIM10LPEN  APB2LPENR = 0x01 << 17 //+ TIM10 clock enable during Sleep mode
	TIM11LPEN  APB2LPENR = 0x01 << 18 //+ TIM11 clock enable during Sleep mode
	DFSDMLPEN  APB2LPENR = 0x01 << 24 //+ DFSDMLPEN
)

const (
	TIM1LPENn   = 0
	TIM8LPENn   = 1
	USART1LPENn = 4
	USART6LPENn = 5
	ADC1LPENn   = 8
	SDIOLPENn   = 11
	SPI1LPENn   = 12
	SPI4LPENn   = 13
	SYSCFGLPENn = 14
	TIM9LPENn   = 16
	TIM10LPENn  = 17
	TIM11LPENn  = 18
	DFSDMLPENn  = 24
)

const (
	LSEON       BDCR = 0x01 << 0  //+ External low-speed oscillator enable
	LSERDY      BDCR = 0x01 << 1  //+ External low-speed oscillator ready
	LSEBYP      BDCR = 0x01 << 2  //+ External low-speed oscillator bypass
	RTCSEL      BDCR = 0x03 << 8  //+ RTC clock source selection
	RTCSEL_NONE BDCR = 0x00 << 8  //  no clock
	RTCSEL_LSE  BDCR = 0x01 << 8  //  LSE oscillator clock used as RTC clock
	RTCSEL_LSI  BDCR = 0x02 << 8  //  LSI oscillator clock used as RTC clock
	RTCSEL_HSE  BDCR = 0x03 << 8  //  HSE clock divided by RTCPRE value is used as RTC clock
	RTCEN       BDCR = 0x01 << 15 //+ RTC clock enable
	BDRST       BDCR = 0x01 << 16 //+ Backup domain software reset
)

const (
	LSEONn  = 0
	LSERDYn = 1
	LSEBYPn = 2
	RTCSELn = 8
	RTCENn  = 15
	BDRSTn  = 16
)

const (
	LSION    CSR = 0x01 << 0  //+ Internal low-speed oscillator enable
	LSIRDY   CSR = 0x01 << 1  //+ Internal low-speed oscillator ready
	RMVF     CSR = 0x01 << 24 //+ Remove reset flag
	BORRSTF  CSR = 0x01 << 25 //+ BOR reset flag
	PADRSTF  CSR = 0x01 << 26 //+ PIN reset flag
	PORRSTF  CSR = 0x01 << 27 //+ POR/PDR reset flag
	SFTRSTF  CSR = 0x01 << 28 //+ Software reset flag
	WDGRSTF  CSR = 0x01 << 29 //+ Independent watchdog reset flag
	WWDGRSTF CSR = 0x01 << 30 //+ Window watchdog reset flag
	LPWRRSTF CSR = 0x01 << 31 //+ Low-power reset flag
)

const (
	LSIONn    = 0
	LSIRDYn   = 1
	RMVFn     = 24
	BORRSTFn  = 25
	PADRSTFn  = 26
	PORRSTFn  = 27
	SFTRSTFn  = 28
	WDGRSTFn  = 29
	WWDGRSTFn = 30
	LPWRRSTFn = 31
)

const (
	MODPER    SSCGR = 0x1FFF << 0  //+ Modulation period
	INCSTEP   SSCGR = 0x7FFF << 13 //+ Incrementation step
	SPREADSEL SSCGR = 0x01 << 30   //+ Spread Select
	SSCGEN    SSCGR = 0x01 << 31   //+ Spread spectrum modulation enable
)

const (
	MODPERn    = 0
	INCSTEPn   = 13
	SPREADSELn = 30
	SSCGENn    = 31
)

const (
	PLLI2SNx PLLI2SCFGR = 0x1FF << 6 //+ PLLI2S multiplication factor for VCO
	PLLI2SRx PLLI2SCFGR = 0x07 << 28 //+ PLLI2S division factor for I2S clocks
)

const (
	PLLI2SNxn = 6
	PLLI2SRxn = 28
)
