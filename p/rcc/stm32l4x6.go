// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32l4x6

// Package rcc provides access to the registers of the RCC peripheral.
//
// Instances:
//  RCC  RCC_BASE  -  RCC  Reset and clock control
// Registers:
//  0x000 32  CR           Clock control register
//  0x004 32  ICSCR        Internal clock sources calibration register
//  0x008 32  CFGR         Clock configuration register
//  0x00C 32  PLLCFGR      PLL configuration register
//  0x010 32  PLLSAI1CFGR  PLLSAI1 configuration register
//  0x014 32  PLLSAI2CFGR  PLLSAI2 configuration register
//  0x018 32  CIER         Clock interrupt enable register
//  0x01C 32  CIFR         Clock interrupt flag register
//  0x020 32  CICR         Clock interrupt clear register
//  0x028 32  AHB1RSTR     AHB1 peripheral reset register
//  0x02C 32  AHB2RSTR     AHB2 peripheral reset register
//  0x030 32  AHB3RSTR     AHB3 peripheral reset register
//  0x038 32  APB1RSTR1    APB1 peripheral reset register 1
//  0x03C 32  APB1RSTR2    APB1 peripheral reset register 2
//  0x040 32  APB2RSTR     APB2 peripheral reset register
//  0x048 32  AHB1ENR      AHB1 peripheral clock enable register
//  0x04C 32  AHB2ENR      AHB2 peripheral clock enable register
//  0x050 32  AHB3ENR      AHB3 peripheral clock enable register
//  0x058 32  APB1ENR1     APB1ENR1
//  0x05C 32  APB1ENR2     APB1 peripheral clock enable register 2
//  0x060 32  APB2ENR      APB2ENR
//  0x068 32  AHB1SMENR    AHB1 peripheral clocks enable in Sleep and Stop modes register
//  0x06C 32  AHB2SMENR    AHB2 peripheral clocks enable in Sleep and Stop modes register
//  0x070 32  AHB3SMENR    AHB3 peripheral clocks enable in Sleep and Stop modes register
//  0x078 32  APB1SMENR1   APB1SMENR1
//  0x07C 32  APB1SMENR2   APB1 peripheral clocks enable in Sleep and Stop modes register 2
//  0x080 32  APB2SMENR    APB2SMENR
//  0x088 32  CCIPR        CCIPR
//  0x090 32  BDCR         BDCR
//  0x094 32  CSR          CSR
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package rcc

const (
	MSION      CR = 0x01 << 0  //+ MSI clock enable
	MSIRDY     CR = 0x01 << 1  //+ MSI clock ready flag
	MSIPLLEN   CR = 0x01 << 2  //+ MSI clock PLL enable
	MSIRGSEL   CR = 0x01 << 3  //+ MSI clock range selection
	MSIRANGE   CR = 0x0F << 4  //+ MSI clock ranges
	HSION      CR = 0x01 << 8  //+ HSI clock enable
	HSIKERON   CR = 0x01 << 9  //+ HSI always enable for peripheral kernels
	HSIRDY     CR = 0x01 << 10 //+ HSI clock ready flag
	HSIASFS    CR = 0x01 << 11 //+ HSI automatic start from Stop
	HSEON      CR = 0x01 << 16 //+ HSE clock enable
	HSERDY     CR = 0x01 << 17 //+ HSE clock ready flag
	HSEBYP     CR = 0x01 << 18 //+ HSE crystal oscillator bypass
	CSSON      CR = 0x01 << 19 //+ Clock security system enable
	PLLON      CR = 0x01 << 24 //+ Main PLL enable
	PLLRDY     CR = 0x01 << 25 //+ Main PLL clock ready flag
	PLLSAI1ON  CR = 0x01 << 26 //+ SAI1 PLL enable
	PLLSAI1RDY CR = 0x01 << 27 //+ SAI1 PLL clock ready flag
	PLLSAI2ON  CR = 0x01 << 28 //+ SAI2 PLL enable
	PLLSAI2RDY CR = 0x01 << 29 //+ SAI2 PLL clock ready flag
)

const (
	MSIONn      = 0
	MSIRDYn     = 1
	MSIPLLENn   = 2
	MSIRGSELn   = 3
	MSIRANGEn   = 4
	HSIONn      = 8
	HSIKERONn   = 9
	HSIRDYn     = 10
	HSIASFSn    = 11
	HSEONn      = 16
	HSERDYn     = 17
	HSEBYPn     = 18
	CSSONn      = 19
	PLLONn      = 24
	PLLRDYn     = 25
	PLLSAI1ONn  = 26
	PLLSAI1RDYn = 27
	PLLSAI2ONn  = 28
	PLLSAI2RDYn = 29
)

const (
	MSICAL  ICSCR = 0xFF << 0  //+ MSI clock calibration
	MSITRIM ICSCR = 0xFF << 8  //+ MSI clock trimming
	HSICAL  ICSCR = 0xFF << 16 //+ HSI clock calibration
	HSITRIM ICSCR = 0x7F << 24 //+ HSI clock trimming
)

const (
	MSICALn  = 0
	MSITRIMn = 8
	HSICALn  = 16
	HSITRIMn = 24
)

const (
	SW          CFGR = 0x03 << 0  //+ System clock switch
	SW_MSI      CFGR = 0x00 << 0  //  MSI selected as system clock
	SW_HSI      CFGR = 0x01 << 0  //  HSI selected as system clock
	SW_HSE      CFGR = 0x02 << 0  //  HSE selected as system clock
	SW_PLL      CFGR = 0x03 << 0  //  PLL selected as system clock
	SWS         CFGR = 0x03 << 2  //+ System clock switch status
	SWS_MSI     CFGR = 0x00 << 2  //  MSI oscillator used as system clock
	SWS_HSI     CFGR = 0x01 << 2  //  HSI oscillator used as system clock
	SWS_HSE     CFGR = 0x02 << 2  //  HSE oscillator used as system clock
	SWS_PLL     CFGR = 0x03 << 2  //  PLL used as system clock
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
	PPRE1       CFGR = 0x07 << 8  //+ PB low-speed prescaler (APB1)
	PPRE1_DIV1  CFGR = 0x00 << 8  //  HCLK not divided
	PPRE1_DIV2  CFGR = 0x04 << 8  //  HCLK divided by 2
	PPRE1_DIV4  CFGR = 0x05 << 8  //  HCLK divided by 4
	PPRE1_DIV8  CFGR = 0x06 << 8  //  HCLK divided by 8
	PPRE1_DIV16 CFGR = 0x07 << 8  //  HCLK divided by 16
	PPRE2       CFGR = 0x07 << 11 //+ APB high-speed prescaler (APB2)
	PPRE2_DIV1  CFGR = 0x00 << 11 //  HCLK not divided
	PPRE2_DIV2  CFGR = 0x04 << 11 //  HCLK divided by 2
	PPRE2_DIV4  CFGR = 0x05 << 11 //  HCLK divided by 4
	PPRE2_DIV8  CFGR = 0x06 << 11 //  HCLK divided by 8
	PPRE2_DIV16 CFGR = 0x07 << 11 //  HCLK divided by 16
	STOPWUCK    CFGR = 0x01 << 15 //+ Wakeup from Stop and CSS backup clock selection
	MCOSEL      CFGR = 0x07 << 24 //+ Microcontroller clock output
	MCOPRE      CFGR = 0x07 << 28 //+ Microcontroller clock output prescaler
)

const (
	SWn       = 0
	SWSn      = 2
	HPREn     = 4
	PPRE1n    = 8
	PPRE2n    = 11
	STOPWUCKn = 15
	MCOSELn   = 24
	MCOPREn   = 28
)

const (
	PLLSRC      PLLCFGR = 0x03 << 0  //+ Main PLL, PLLSAI1 and PLLSAI2 entry clock source
	PLLSRC_NONE PLLCFGR = 0x00 << 0  //  No clock source selected
	PLLSRC_MSI  PLLCFGR = 0x01 << 0  //  MSI oscillator source clock selected
	PLLSRC_HSI  PLLCFGR = 0x02 << 0  //  HSI16 oscillator source clock selected
	PLLSRC_HSE  PLLCFGR = 0x03 << 0  //  HSE oscillator source clock selected
	PLLM        PLLCFGR = 0x07 << 4  //+ Division factor for the main PLL and audio PLL (PLLSAI1 and PLLSAI2) input clock
	PLLN        PLLCFGR = 0x7F << 8  //+ Main PLL multiplication factor for VCO
	PLLPEN      PLLCFGR = 0x01 << 16 //+ Main PLL PLLSAI3CLK output enable
	PLLP        PLLCFGR = 0x01 << 17 //+ Main PLL division factor for PLLSAI3CLK (SAI1 and SAI2 clock)
	PLLQEN      PLLCFGR = 0x01 << 20 //+ Main PLL PLLUSB1CLK output enable
	PLLQ        PLLCFGR = 0x03 << 21 //+ Main PLL division factor for PLLUSB1CLK(48 MHz clock)
	PLLREN      PLLCFGR = 0x01 << 24 //+ Main PLL PLLCLK output enable
	PLLR        PLLCFGR = 0x03 << 25 //+ Main PLL division factor for PLLCLK (system clock)
	PLLPDIV     PLLCFGR = 0x1F << 27 //+ Main PLL division factor for PLLSAI2CLK
)

const (
	PLLSRCn  = 0
	PLLMn    = 4
	PLLNn    = 8
	PLLPENn  = 16
	PLLPn    = 17
	PLLQENn  = 20
	PLLQn    = 21
	PLLRENn  = 24
	PLLRn    = 25
	PLLPDIVn = 27
)

const (
	PLLSAI1N    PLLSAI1CFGR = 0x7F << 8  //+ SAI1PLL multiplication factor for VCO
	PLLSAI1PEN  PLLSAI1CFGR = 0x01 << 16 //+ SAI1PLL PLLSAI1CLK output enable
	PLLSAI1P    PLLSAI1CFGR = 0x01 << 17 //+ SAI1PLL division factor for PLLSAI1CLK (SAI1 or SAI2 clock)
	PLLSAI1QEN  PLLSAI1CFGR = 0x01 << 20 //+ SAI1PLL PLLUSB2CLK output enable
	PLLSAI1Q    PLLSAI1CFGR = 0x03 << 21 //+ SAI1PLL division factor for PLLUSB2CLK (48 MHz clock)
	PLLSAI1REN  PLLSAI1CFGR = 0x01 << 24 //+ PLLSAI1 PLLADC1CLK output enable
	PLLSAI1R    PLLSAI1CFGR = 0x03 << 25 //+ PLLSAI1 division factor for PLLADC1CLK (ADC clock)
	PLLSAI1PDIV PLLSAI1CFGR = 0x1F << 27 //+ PLLSAI1 division factor for PLLSAI1CLK
)

const (
	PLLSAI1Nn    = 8
	PLLSAI1PENn  = 16
	PLLSAI1Pn    = 17
	PLLSAI1QENn  = 20
	PLLSAI1Qn    = 21
	PLLSAI1RENn  = 24
	PLLSAI1Rn    = 25
	PLLSAI1PDIVn = 27
)

const (
	PLLSAI2N    PLLSAI2CFGR = 0x7F << 8  //+ SAI2PLL multiplication factor for VCO
	PLLSAI2PEN  PLLSAI2CFGR = 0x01 << 16 //+ SAI2PLL PLLSAI2CLK output enable
	PLLSAI2P    PLLSAI2CFGR = 0x01 << 17 //+ SAI1PLL division factor for PLLSAI2CLK (SAI1 or SAI2 clock)
	PLLSAI2REN  PLLSAI2CFGR = 0x01 << 24 //+ PLLSAI2 PLLADC2CLK output enable
	PLLSAI2R    PLLSAI2CFGR = 0x03 << 25 //+ PLLSAI2 division factor for PLLADC2CLK (ADC clock)
	PLLSAI2PDIV PLLSAI2CFGR = 0x1F << 27 //+ PLLSAI2 division factor for PLLSAI2CLK
)

const (
	PLLSAI2Nn    = 8
	PLLSAI2PENn  = 16
	PLLSAI2Pn    = 17
	PLLSAI2RENn  = 24
	PLLSAI2Rn    = 25
	PLLSAI2PDIVn = 27
)

const (
	LSIRDYIE     CIER = 0x01 << 0  //+ LSI ready interrupt enable
	LSERDYIE     CIER = 0x01 << 1  //+ LSE ready interrupt enable
	MSIRDYIE     CIER = 0x01 << 2  //+ MSI ready interrupt enable
	HSIRDYIE     CIER = 0x01 << 3  //+ HSI ready interrupt enable
	HSERDYIE     CIER = 0x01 << 4  //+ HSE ready interrupt enable
	PLLRDYIE     CIER = 0x01 << 5  //+ PLL ready interrupt enable
	PLLSAI1RDYIE CIER = 0x01 << 6  //+ PLLSAI1 ready interrupt enable
	PLLSAI2RDYIE CIER = 0x01 << 7  //+ PLLSAI2 ready interrupt enable
	LSECSSIE     CIER = 0x01 << 9  //+ LSE clock security system interrupt enable
	HSI48RDYIE   CIER = 0x01 << 10 //+ HSI48 ready interrupt enable
)

const (
	LSIRDYIEn     = 0
	LSERDYIEn     = 1
	MSIRDYIEn     = 2
	HSIRDYIEn     = 3
	HSERDYIEn     = 4
	PLLRDYIEn     = 5
	PLLSAI1RDYIEn = 6
	PLLSAI2RDYIEn = 7
	LSECSSIEn     = 9
	HSI48RDYIEn   = 10
)

const (
	LSIRDYF     CIFR = 0x01 << 0  //+ LSI ready interrupt flag
	LSERDYF     CIFR = 0x01 << 1  //+ LSE ready interrupt flag
	MSIRDYF     CIFR = 0x01 << 2  //+ MSI ready interrupt flag
	HSIRDYF     CIFR = 0x01 << 3  //+ HSI ready interrupt flag
	HSERDYF     CIFR = 0x01 << 4  //+ HSE ready interrupt flag
	PLLRDYF     CIFR = 0x01 << 5  //+ PLL ready interrupt flag
	PLLSAI1RDYF CIFR = 0x01 << 6  //+ PLLSAI1 ready interrupt flag
	PLLSAI2RDYF CIFR = 0x01 << 7  //+ PLLSAI2 ready interrupt flag
	CSSF        CIFR = 0x01 << 8  //+ Clock security system interrupt flag
	LSECSSF     CIFR = 0x01 << 9  //+ LSE Clock security system interrupt flag
	HSI48RDYF   CIFR = 0x01 << 10 //+ HSI48 ready interrupt flag
)

const (
	LSIRDYFn     = 0
	LSERDYFn     = 1
	MSIRDYFn     = 2
	HSIRDYFn     = 3
	HSERDYFn     = 4
	PLLRDYFn     = 5
	PLLSAI1RDYFn = 6
	PLLSAI2RDYFn = 7
	CSSFn        = 8
	LSECSSFn     = 9
	HSI48RDYFn   = 10
)

const (
	LSIRDYC     CICR = 0x01 << 0  //+ LSI ready interrupt clear
	LSERDYC     CICR = 0x01 << 1  //+ LSE ready interrupt clear
	MSIRDYC     CICR = 0x01 << 2  //+ MSI ready interrupt clear
	HSIRDYC     CICR = 0x01 << 3  //+ HSI ready interrupt clear
	HSERDYC     CICR = 0x01 << 4  //+ HSE ready interrupt clear
	PLLRDYC     CICR = 0x01 << 5  //+ PLL ready interrupt clear
	PLLSAI1RDYC CICR = 0x01 << 6  //+ PLLSAI1 ready interrupt clear
	PLLSAI2RDYC CICR = 0x01 << 7  //+ PLLSAI2 ready interrupt clear
	CSSC        CICR = 0x01 << 8  //+ Clock security system interrupt clear
	LSECSSC     CICR = 0x01 << 9  //+ LSE Clock security system interrupt clear
	HSI48RDYC   CICR = 0x01 << 10 //+ HSI48 oscillator ready interrupt clear
)

const (
	LSIRDYCn     = 0
	LSERDYCn     = 1
	MSIRDYCn     = 2
	HSIRDYCn     = 3
	HSERDYCn     = 4
	PLLRDYCn     = 5
	PLLSAI1RDYCn = 6
	PLLSAI2RDYCn = 7
	CSSCn        = 8
	LSECSSCn     = 9
	HSI48RDYCn   = 10
)

const (
	DMA1RST  AHB1RSTR = 0x01 << 0  //+ DMA1 reset
	DMA2RST  AHB1RSTR = 0x01 << 1  //+ DMA2 reset
	FLASHRST AHB1RSTR = 0x01 << 8  //+ Flash memory interface reset
	CRCRST   AHB1RSTR = 0x01 << 12 //+ CRC reset
	TSCRST   AHB1RSTR = 0x01 << 16 //+ Touch Sensing Controller reset
	DMA2DRST AHB1RSTR = 0x01 << 17 //+ DMA2D reset
)

const (
	DMA1RSTn  = 0
	DMA2RSTn  = 1
	FLASHRSTn = 8
	CRCRSTn   = 12
	TSCRSTn   = 16
	DMA2DRSTn = 17
)

const (
	GPIOARST AHB2RSTR = 0x01 << 0  //+ IO port A reset
	GPIOBRST AHB2RSTR = 0x01 << 1  //+ IO port B reset
	GPIOCRST AHB2RSTR = 0x01 << 2  //+ IO port C reset
	GPIODRST AHB2RSTR = 0x01 << 3  //+ IO port D reset
	GPIOERST AHB2RSTR = 0x01 << 4  //+ IO port E reset
	GPIOFRST AHB2RSTR = 0x01 << 5  //+ IO port F reset
	GPIOGRST AHB2RSTR = 0x01 << 6  //+ IO port G reset
	GPIOHRST AHB2RSTR = 0x01 << 7  //+ IO port H reset
	GPIOIRST AHB2RSTR = 0x01 << 8  //+ IO port I reset
	OTGFSRST AHB2RSTR = 0x01 << 12 //+ USB OTG FS reset
	ADCRST   AHB2RSTR = 0x01 << 13 //+ ADC reset
	DCMIRST  AHB2RSTR = 0x01 << 14 //+ Digital Camera Interface reset
	AESRST   AHB2RSTR = 0x01 << 16 //+ AES hardware accelerator reset
	HASH1RST AHB2RSTR = 0x01 << 17 //+ Hash reset
	RNGRST   AHB2RSTR = 0x01 << 18 //+ Random number generator reset
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
	GPIOIRSTn = 8
	OTGFSRSTn = 12
	ADCRSTn   = 13
	DCMIRSTn  = 14
	AESRSTn   = 16
	HASH1RSTn = 17
	RNGRSTn   = 18
)

const (
	FMCRST  AHB3RSTR = 0x01 << 0 //+ Flexible memory controller reset
	QSPIRST AHB3RSTR = 0x01 << 8 //+ Quad SPI memory interface reset
)

const (
	FMCRSTn  = 0
	QSPIRSTn = 8
)

const (
	TIM2RST   APB1RSTR1 = 0x01 << 0  //+ TIM2 timer reset
	TIM3RST   APB1RSTR1 = 0x01 << 1  //+ TIM3 timer reset
	TIM4RST   APB1RSTR1 = 0x01 << 2  //+ TIM3 timer reset
	TIM5RST   APB1RSTR1 = 0x01 << 3  //+ TIM5 timer reset
	TIM6RST   APB1RSTR1 = 0x01 << 4  //+ TIM6 timer reset
	TIM7RST   APB1RSTR1 = 0x01 << 5  //+ TIM7 timer reset
	LCDRST    APB1RSTR1 = 0x01 << 9  //+ LCD interface reset
	SPI2RST   APB1RSTR1 = 0x01 << 14 //+ SPI2 reset
	SPI3RST   APB1RSTR1 = 0x01 << 15 //+ SPI3 reset
	USART2RST APB1RSTR1 = 0x01 << 17 //+ USART2 reset
	USART3RST APB1RSTR1 = 0x01 << 18 //+ USART3 reset
	UART4RST  APB1RSTR1 = 0x01 << 19 //+ UART4 reset
	UART5RST  APB1RSTR1 = 0x01 << 20 //+ UART5 reset
	I2C1RST   APB1RSTR1 = 0x01 << 21 //+ I2C1 reset
	I2C2RST   APB1RSTR1 = 0x01 << 22 //+ I2C2 reset
	I2C3RST   APB1RSTR1 = 0x01 << 23 //+ I2C3 reset
	CRSRST    APB1RSTR1 = 0x01 << 24 //+ CRS reset
	CAN1RST   APB1RSTR1 = 0x01 << 25 //+ CAN1 reset
	CAN2RST   APB1RSTR1 = 0x01 << 26 //+ CAN2 reset
	PWRRST    APB1RSTR1 = 0x01 << 28 //+ Power interface reset
	DAC1RST   APB1RSTR1 = 0x01 << 29 //+ DAC1 interface reset
	OPAMPRST  APB1RSTR1 = 0x01 << 30 //+ OPAMP interface reset
	LPTIM1RST APB1RSTR1 = 0x01 << 31 //+ Low Power Timer 1 reset
)

const (
	TIM2RSTn   = 0
	TIM3RSTn   = 1
	TIM4RSTn   = 2
	TIM5RSTn   = 3
	TIM6RSTn   = 4
	TIM7RSTn   = 5
	LCDRSTn    = 9
	SPI2RSTn   = 14
	SPI3RSTn   = 15
	USART2RSTn = 17
	USART3RSTn = 18
	UART4RSTn  = 19
	UART5RSTn  = 20
	I2C1RSTn   = 21
	I2C2RSTn   = 22
	I2C3RSTn   = 23
	CRSRSTn    = 24
	CAN1RSTn   = 25
	CAN2RSTn   = 26
	PWRRSTn    = 28
	DAC1RSTn   = 29
	OPAMPRSTn  = 30
	LPTIM1RSTn = 31
)

const (
	LPUART1RST APB1RSTR2 = 0x01 << 0 //+ Low-power UART 1 reset
	I2C4RST    APB1RSTR2 = 0x01 << 1 //+ I2C4 reset
	SWPMI1RST  APB1RSTR2 = 0x01 << 2 //+ Single wire protocol reset
	LPTIM2RST  APB1RSTR2 = 0x01 << 5 //+ Low-power timer 2 reset
)

const (
	LPUART1RSTn = 0
	I2C4RSTn    = 1
	SWPMI1RSTn  = 2
	LPTIM2RSTn  = 5
)

const (
	SYSCFGRST APB2RSTR = 0x01 << 0  //+ System configuration (SYSCFG) reset
	SDMMCRST  APB2RSTR = 0x01 << 10 //+ SDMMC reset
	TIM1RST   APB2RSTR = 0x01 << 11 //+ TIM1 timer reset
	SPI1RST   APB2RSTR = 0x01 << 12 //+ SPI1 reset
	TIM8RST   APB2RSTR = 0x01 << 13 //+ TIM8 timer reset
	USART1RST APB2RSTR = 0x01 << 14 //+ USART1 reset
	TIM15RST  APB2RSTR = 0x01 << 16 //+ TIM15 timer reset
	TIM16RST  APB2RSTR = 0x01 << 17 //+ TIM16 timer reset
	TIM17RST  APB2RSTR = 0x01 << 18 //+ TIM17 timer reset
	SAI1RST   APB2RSTR = 0x01 << 21 //+ Serial audio interface 1 (SAI1) reset
	SAI2RST   APB2RSTR = 0x01 << 22 //+ Serial audio interface 2 (SAI2) reset
	DFSDMRST  APB2RSTR = 0x01 << 24 //+ Digital filters for sigma-delata modulators (DFSDM) reset
)

const (
	SYSCFGRSTn = 0
	SDMMCRSTn  = 10
	TIM1RSTn   = 11
	SPI1RSTn   = 12
	TIM8RSTn   = 13
	USART1RSTn = 14
	TIM15RSTn  = 16
	TIM16RSTn  = 17
	TIM17RSTn  = 18
	SAI1RSTn   = 21
	SAI2RSTn   = 22
	DFSDMRSTn  = 24
)

const (
	DMA1EN  AHB1ENR = 0x01 << 0  //+ DMA1 clock enable
	DMA2EN  AHB1ENR = 0x01 << 1  //+ DMA2 clock enable
	FLASHEN AHB1ENR = 0x01 << 8  //+ Flash memory interface clock enable
	CRCEN   AHB1ENR = 0x01 << 12 //+ CRC clock enable
	TSCEN   AHB1ENR = 0x01 << 16 //+ Touch Sensing Controller clock enable
	DMA2DEN AHB1ENR = 0x01 << 17 //+ DMA2D clock enable
)

const (
	DMA1ENn  = 0
	DMA2ENn  = 1
	FLASHENn = 8
	CRCENn   = 12
	TSCENn   = 16
	DMA2DENn = 17
)

const (
	GPIOAEN AHB2ENR = 0x01 << 0  //+ IO port A clock enable
	GPIOBEN AHB2ENR = 0x01 << 1  //+ IO port B clock enable
	GPIOCEN AHB2ENR = 0x01 << 2  //+ IO port C clock enable
	GPIODEN AHB2ENR = 0x01 << 3  //+ IO port D clock enable
	GPIOEEN AHB2ENR = 0x01 << 4  //+ IO port E clock enable
	GPIOFEN AHB2ENR = 0x01 << 5  //+ IO port F clock enable
	GPIOGEN AHB2ENR = 0x01 << 6  //+ IO port G clock enable
	GPIOHEN AHB2ENR = 0x01 << 7  //+ IO port H clock enable
	GPIOIEN AHB2ENR = 0x01 << 8  //+ IO port I clock enable
	OTGFSEN AHB2ENR = 0x01 << 12 //+ OTG full speed clock enable
	ADCEN   AHB2ENR = 0x01 << 13 //+ ADC clock enable
	DCMIEN  AHB2ENR = 0x01 << 14 //+ DCMI clock enable
	AESEN   AHB2ENR = 0x01 << 16 //+ AES accelerator clock enable
	HASH1EN AHB2ENR = 0x01 << 17 //+ HASH clock enable
	RNGEN   AHB2ENR = 0x01 << 18 //+ Random Number Generator clock enable
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
	GPIOIENn = 8
	OTGFSENn = 12
	ADCENn   = 13
	DCMIENn  = 14
	AESENn   = 16
	HASH1ENn = 17
	RNGENn   = 18
)

const (
	FMCEN  AHB3ENR = 0x01 << 0 //+ Flexible memory controller clock enable
	QSPIEN AHB3ENR = 0x01 << 8 //+ QSPIEN
)

const (
	FMCENn  = 0
	QSPIENn = 8
)

const (
	TIM2EN   APB1ENR1 = 0x01 << 0  //+ TIM2 timer clock enable
	TIM3EN   APB1ENR1 = 0x01 << 1  //+ TIM3 timer clock enable
	TIM4EN   APB1ENR1 = 0x01 << 2  //+ TIM4 timer clock enable
	TIM5EN   APB1ENR1 = 0x01 << 3  //+ TIM5 timer clock enable
	TIM6EN   APB1ENR1 = 0x01 << 4  //+ TIM6 timer clock enable
	TIM7EN   APB1ENR1 = 0x01 << 5  //+ TIM7 timer clock enable
	LCDEN    APB1ENR1 = 0x01 << 9  //+ LCD clock enable
	RTCAPBEN APB1ENR1 = 0x01 << 10 //+ RTC APB clock enable
	WWDGEN   APB1ENR1 = 0x01 << 11 //+ Window watchdog clock enable
	SPI2EN   APB1ENR1 = 0x01 << 14 //+ SPI2 clock enable
	SP3EN    APB1ENR1 = 0x01 << 15 //+ SPI3 clock enable
	USART2EN APB1ENR1 = 0x01 << 17 //+ USART2 clock enable
	USART3EN APB1ENR1 = 0x01 << 18 //+ USART3 clock enable
	UART4EN  APB1ENR1 = 0x01 << 19 //+ UART4 clock enable
	UART5EN  APB1ENR1 = 0x01 << 20 //+ UART5 clock enable
	I2C1EN   APB1ENR1 = 0x01 << 21 //+ I2C1 clock enable
	I2C2EN   APB1ENR1 = 0x01 << 22 //+ I2C2 clock enable
	I2C3EN   APB1ENR1 = 0x01 << 23 //+ I2C3 clock enable
	CRSEN    APB1ENR1 = 0x01 << 24 //+ Clock Recovery System clock enable
	CAN1EN   APB1ENR1 = 0x01 << 25 //+ CAN1 clock enable
	CAN2EN   APB1ENR1 = 0x01 << 26 //+ CAN2 clock enable
	PWREN    APB1ENR1 = 0x01 << 28 //+ Power interface clock enable
	DAC1EN   APB1ENR1 = 0x01 << 29 //+ DAC1 interface clock enable
	OPAMPEN  APB1ENR1 = 0x01 << 30 //+ OPAMP interface clock enable
	LPTIM1EN APB1ENR1 = 0x01 << 31 //+ Low power timer 1 clock enable
)

const (
	TIM2ENn   = 0
	TIM3ENn   = 1
	TIM4ENn   = 2
	TIM5ENn   = 3
	TIM6ENn   = 4
	TIM7ENn   = 5
	LCDENn    = 9
	RTCAPBENn = 10
	WWDGENn   = 11
	SPI2ENn   = 14
	SP3ENn    = 15
	USART2ENn = 17
	USART3ENn = 18
	UART4ENn  = 19
	UART5ENn  = 20
	I2C1ENn   = 21
	I2C2ENn   = 22
	I2C3ENn   = 23
	CRSENn    = 24
	CAN1ENn   = 25
	CAN2ENn   = 26
	PWRENn    = 28
	DAC1ENn   = 29
	OPAMPENn  = 30
	LPTIM1ENn = 31
)

const (
	LPUART1EN APB1ENR2 = 0x01 << 0 //+ Low power UART 1 clock enable
	I2C4EN    APB1ENR2 = 0x01 << 1 //+ I2C4 clock enable
	SWPMI1EN  APB1ENR2 = 0x01 << 2 //+ Single wire protocol clock enable
	LPTIM2EN  APB1ENR2 = 0x01 << 5 //+ LPTIM2EN
)

const (
	LPUART1ENn = 0
	I2C4ENn    = 1
	SWPMI1ENn  = 2
	LPTIM2ENn  = 5
)

const (
	SYSCFGEN   APB2ENR = 0x01 << 0  //+ SYSCFG clock enable
	FIREWALLEN APB2ENR = 0x01 << 7  //+ Firewall clock enable
	SDMMCEN    APB2ENR = 0x01 << 10 //+ SDMMC clock enable
	TIM1EN     APB2ENR = 0x01 << 11 //+ TIM1 timer clock enable
	SPI1EN     APB2ENR = 0x01 << 12 //+ SPI1 clock enable
	TIM8EN     APB2ENR = 0x01 << 13 //+ TIM8 timer clock enable
	USART1EN   APB2ENR = 0x01 << 14 //+ USART1clock enable
	TIM15EN    APB2ENR = 0x01 << 16 //+ TIM15 timer clock enable
	TIM16EN    APB2ENR = 0x01 << 17 //+ TIM16 timer clock enable
	TIM17EN    APB2ENR = 0x01 << 18 //+ TIM17 timer clock enable
	SAI1EN     APB2ENR = 0x01 << 21 //+ SAI1 clock enable
	SAI2EN     APB2ENR = 0x01 << 22 //+ SAI2 clock enable
	DFSDMEN    APB2ENR = 0x01 << 24 //+ DFSDM timer clock enable
)

const (
	SYSCFGENn   = 0
	FIREWALLENn = 7
	SDMMCENn    = 10
	TIM1ENn     = 11
	SPI1ENn     = 12
	TIM8ENn     = 13
	USART1ENn   = 14
	TIM15ENn    = 16
	TIM16ENn    = 17
	TIM17ENn    = 18
	SAI1ENn     = 21
	SAI2ENn     = 22
	DFSDMENn    = 24
)

const (
	DMA1SMEN  AHB1SMENR = 0x01 << 0  //+ DMA1 clocks enable during Sleep and Stop modes
	DMA2SMEN  AHB1SMENR = 0x01 << 1  //+ DMA2 clocks enable during Sleep and Stop modes
	FLASHSMEN AHB1SMENR = 0x01 << 8  //+ Flash memory interface clocks enable during Sleep and Stop modes
	SRAM1SMEN AHB1SMENR = 0x01 << 9  //+ SRAM1 interface clocks enable during Sleep and Stop modes
	CRCSMEN   AHB1SMENR = 0x01 << 12 //+ CRCSMEN
	TSCSMEN   AHB1SMENR = 0x01 << 16 //+ Touch Sensing Controller clocks enable during Sleep and Stop modes
	DMA2DSMEN AHB1SMENR = 0x01 << 17 //+ DMA2D clock enable during Sleep and Stop modes
)

const (
	DMA1SMENn  = 0
	DMA2SMENn  = 1
	FLASHSMENn = 8
	SRAM1SMENn = 9
	CRCSMENn   = 12
	TSCSMENn   = 16
	DMA2DSMENn = 17
)

const (
	GPIOASMEN AHB2SMENR = 0x01 << 0  //+ IO port A clocks enable during Sleep and Stop modes
	GPIOBSMEN AHB2SMENR = 0x01 << 1  //+ IO port B clocks enable during Sleep and Stop modes
	GPIOCSMEN AHB2SMENR = 0x01 << 2  //+ IO port C clocks enable during Sleep and Stop modes
	GPIODSMEN AHB2SMENR = 0x01 << 3  //+ IO port D clocks enable during Sleep and Stop modes
	GPIOESMEN AHB2SMENR = 0x01 << 4  //+ IO port E clocks enable during Sleep and Stop modes
	GPIOFSMEN AHB2SMENR = 0x01 << 5  //+ IO port F clocks enable during Sleep and Stop modes
	GPIOGSMEN AHB2SMENR = 0x01 << 6  //+ IO port G clocks enable during Sleep and Stop modes
	GPIOHSMEN AHB2SMENR = 0x01 << 7  //+ IO port H clocks enable during Sleep and Stop modes
	GPIOISMEN AHB2SMENR = 0x01 << 8  //+ IO port I clocks enable during Sleep and Stop modes
	SRAM2SMEN AHB2SMENR = 0x01 << 9  //+ SRAM2 interface clocks enable during Sleep and Stop modes
	OTGFSSMEN AHB2SMENR = 0x01 << 12 //+ OTG full speed clocks enable during Sleep and Stop modes
	ADCFSSMEN AHB2SMENR = 0x01 << 13 //+ ADC clocks enable during Sleep and Stop modes
	DCMISMEN  AHB2SMENR = 0x01 << 14 //+ DCMI clock enable during Sleep and Stop modes
	AESSMEN   AHB2SMENR = 0x01 << 16 //+ AES accelerator clocks enable during Sleep and Stop modes
	HASH1SMEN AHB2SMENR = 0x01 << 17 //+ HASH clock enable during Sleep and Stop modes
	RNGSMEN   AHB2SMENR = 0x01 << 18 //+ Random Number Generator clocks enable during Sleep and Stop modes
)

const (
	GPIOASMENn = 0
	GPIOBSMENn = 1
	GPIOCSMENn = 2
	GPIODSMENn = 3
	GPIOESMENn = 4
	GPIOFSMENn = 5
	GPIOGSMENn = 6
	GPIOHSMENn = 7
	GPIOISMENn = 8
	SRAM2SMENn = 9
	OTGFSSMENn = 12
	ADCFSSMENn = 13
	DCMISMENn  = 14
	AESSMENn   = 16
	HASH1SMENn = 17
	RNGSMENn   = 18
)

const (
	FMCSMEN  AHB3SMENR = 0x01 << 0 //+ Flexible memory controller clocks enable during Sleep and Stop modes
	QSPISMEN AHB3SMENR = 0x01 << 8 //+ QSPISMEN
)

const (
	FMCSMENn  = 0
	QSPISMENn = 8
)

const (
	TIM2SMEN   APB1SMENR1 = 0x01 << 0  //+ TIM2 timer clocks enable during Sleep and Stop modes
	TIM3SMEN   APB1SMENR1 = 0x01 << 1  //+ TIM3 timer clocks enable during Sleep and Stop modes
	TIM4SMEN   APB1SMENR1 = 0x01 << 2  //+ TIM4 timer clocks enable during Sleep and Stop modes
	TIM5SMEN   APB1SMENR1 = 0x01 << 3  //+ TIM5 timer clocks enable during Sleep and Stop modes
	TIM6SMEN   APB1SMENR1 = 0x01 << 4  //+ TIM6 timer clocks enable during Sleep and Stop modes
	TIM7SMEN   APB1SMENR1 = 0x01 << 5  //+ TIM7 timer clocks enable during Sleep and Stop modes
	LCDSMEN    APB1SMENR1 = 0x01 << 9  //+ LCD clocks enable during Sleep and Stop modes
	RTCAPBSMEN APB1SMENR1 = 0x01 << 10 //+ RTC APB clock enable during Sleep and Stop modes
	WWDGSMEN   APB1SMENR1 = 0x01 << 11 //+ Window watchdog clocks enable during Sleep and Stop modes
	SPI2SMEN   APB1SMENR1 = 0x01 << 14 //+ SPI2 clocks enable during Sleep and Stop modes
	SP3SMEN    APB1SMENR1 = 0x01 << 15 //+ SPI3 clocks enable during Sleep and Stop modes
	USART2SMEN APB1SMENR1 = 0x01 << 17 //+ USART2 clocks enable during Sleep and Stop modes
	USART3SMEN APB1SMENR1 = 0x01 << 18 //+ USART3 clocks enable during Sleep and Stop modes
	UART4SMEN  APB1SMENR1 = 0x01 << 19 //+ UART4 clocks enable during Sleep and Stop modes
	UART5SMEN  APB1SMENR1 = 0x01 << 20 //+ UART5 clocks enable during Sleep and Stop modes
	I2C1SMEN   APB1SMENR1 = 0x01 << 21 //+ I2C1 clocks enable during Sleep and Stop modes
	I2C2SMEN   APB1SMENR1 = 0x01 << 22 //+ I2C2 clocks enable during Sleep and Stop modes
	I2C3SMEN   APB1SMENR1 = 0x01 << 23 //+ I2C3 clocks enable during Sleep and Stop modes
	CAN1SMEN   APB1SMENR1 = 0x01 << 25 //+ CAN1 clocks enable during Sleep and Stop modes
	CAN2SMEN   APB1SMENR1 = 0x01 << 26 //+ CAN2 clocks enable during Sleep and Stop modes
	PWRSMEN    APB1SMENR1 = 0x01 << 28 //+ Power interface clocks enable during Sleep and Stop modes
	DAC1SMEN   APB1SMENR1 = 0x01 << 29 //+ DAC1 interface clocks enable during Sleep and Stop modes
	OPAMPSMEN  APB1SMENR1 = 0x01 << 30 //+ OPAMP interface clocks enable during Sleep and Stop modes
	LPTIM1SMEN APB1SMENR1 = 0x01 << 31 //+ Low power timer 1 clocks enable during Sleep and Stop modes
)

const (
	TIM2SMENn   = 0
	TIM3SMENn   = 1
	TIM4SMENn   = 2
	TIM5SMENn   = 3
	TIM6SMENn   = 4
	TIM7SMENn   = 5
	LCDSMENn    = 9
	RTCAPBSMENn = 10
	WWDGSMENn   = 11
	SPI2SMENn   = 14
	SP3SMENn    = 15
	USART2SMENn = 17
	USART3SMENn = 18
	UART4SMENn  = 19
	UART5SMENn  = 20
	I2C1SMENn   = 21
	I2C2SMENn   = 22
	I2C3SMENn   = 23
	CAN1SMENn   = 25
	CAN2SMENn   = 26
	PWRSMENn    = 28
	DAC1SMENn   = 29
	OPAMPSMENn  = 30
	LPTIM1SMENn = 31
)

const (
	LPUART1SMEN APB1SMENR2 = 0x01 << 0 //+ Low power UART 1 clocks enable during Sleep and Stop modes
	I2C4SMEN    APB1SMENR2 = 0x01 << 1 //+ I2C4 clocks enable during Sleep and Stop modes
	SWPMI1SMEN  APB1SMENR2 = 0x01 << 2 //+ Single wire protocol clocks enable during Sleep and Stop modes
	LPTIM2SMEN  APB1SMENR2 = 0x01 << 5 //+ LPTIM2SMEN
)

const (
	LPUART1SMENn = 0
	I2C4SMENn    = 1
	SWPMI1SMENn  = 2
	LPTIM2SMENn  = 5
)

const (
	SYSCFGSMEN APB2SMENR = 0x01 << 0  //+ SYSCFG clocks enable during Sleep and Stop modes
	SDMMCSMEN  APB2SMENR = 0x01 << 10 //+ SDMMC clocks enable during Sleep and Stop modes
	TIM1SMEN   APB2SMENR = 0x01 << 11 //+ TIM1 timer clocks enable during Sleep and Stop modes
	SPI1SMEN   APB2SMENR = 0x01 << 12 //+ SPI1 clocks enable during Sleep and Stop modes
	TIM8SMEN   APB2SMENR = 0x01 << 13 //+ TIM8 timer clocks enable during Sleep and Stop modes
	USART1SMEN APB2SMENR = 0x01 << 14 //+ USART1clocks enable during Sleep and Stop modes
	TIM15SMEN  APB2SMENR = 0x01 << 16 //+ TIM15 timer clocks enable during Sleep and Stop modes
	TIM16SMEN  APB2SMENR = 0x01 << 17 //+ TIM16 timer clocks enable during Sleep and Stop modes
	TIM17SMEN  APB2SMENR = 0x01 << 18 //+ TIM17 timer clocks enable during Sleep and Stop modes
	SAI1SMEN   APB2SMENR = 0x01 << 21 //+ SAI1 clocks enable during Sleep and Stop modes
	SAI2SMEN   APB2SMENR = 0x01 << 22 //+ SAI2 clocks enable during Sleep and Stop modes
	DFSDMSMEN  APB2SMENR = 0x01 << 24 //+ DFSDM timer clocks enable during Sleep and Stop modes
)

const (
	SYSCFGSMENn = 0
	SDMMCSMENn  = 10
	TIM1SMENn   = 11
	SPI1SMENn   = 12
	TIM8SMENn   = 13
	USART1SMENn = 14
	TIM15SMENn  = 16
	TIM16SMENn  = 17
	TIM17SMENn  = 18
	SAI1SMENn   = 21
	SAI2SMENn   = 22
	DFSDMSMENn  = 24
)

const (
	USART1SEL  CCIPR = 0x03 << 0  //+ USART1 clock source selection
	USART2SEL  CCIPR = 0x03 << 2  //+ USART2 clock source selection
	USART3SEL  CCIPR = 0x03 << 4  //+ USART3 clock source selection
	UART4SEL   CCIPR = 0x03 << 6  //+ UART4 clock source selection
	UART5SEL   CCIPR = 0x03 << 8  //+ UART5 clock source selection
	LPUART1SEL CCIPR = 0x03 << 10 //+ LPUART1 clock source selection
	I2C1SEL    CCIPR = 0x03 << 12 //+ I2C1 clock source selection
	I2C2SEL    CCIPR = 0x03 << 14 //+ I2C2 clock source selection
	I2C3SEL    CCIPR = 0x03 << 16 //+ I2C3 clock source selection
	LPTIM1SEL  CCIPR = 0x03 << 18 //+ Low power timer 1 clock source selection
	LPTIM2SEL  CCIPR = 0x03 << 20 //+ Low power timer 2 clock source selection
	SAI1SEL    CCIPR = 0x03 << 22 //+ SAI1 clock source selection
	SAI2SEL    CCIPR = 0x03 << 24 //+ SAI2 clock source selection
	CLK48SEL   CCIPR = 0x03 << 26 //+ 48 MHz clock source selection
	ADCSEL     CCIPR = 0x03 << 28 //+ ADCs clock source selection
	SWPMI1SEL  CCIPR = 0x01 << 30 //+ SWPMI1 clock source selection
	DFSDMSEL   CCIPR = 0x01 << 31 //+ DFSDM clock source selection
)

const (
	USART1SELn  = 0
	USART2SELn  = 2
	USART3SELn  = 4
	UART4SELn   = 6
	UART5SELn   = 8
	LPUART1SELn = 10
	I2C1SELn    = 12
	I2C2SELn    = 14
	I2C3SELn    = 16
	LPTIM1SELn  = 18
	LPTIM2SELn  = 20
	SAI1SELn    = 22
	SAI2SELn    = 24
	CLK48SELn   = 26
	ADCSELn     = 28
	SWPMI1SELn  = 30
	DFSDMSELn   = 31
)

const (
	LSEON    BDCR = 0x01 << 0  //+ LSE oscillator enable
	LSERDY   BDCR = 0x01 << 1  //+ LSE oscillator ready
	LSEBYP   BDCR = 0x01 << 2  //+ LSE oscillator bypass
	LSEDRV   BDCR = 0x03 << 3  //+ SE oscillator drive capability
	LSECSSON BDCR = 0x01 << 5  //+ LSECSSON
	LSECSSD  BDCR = 0x01 << 6  //+ LSECSSD
	RTCSEL   BDCR = 0x03 << 8  //+ RTC clock source selection
	RTCEN    BDCR = 0x01 << 15 //+ RTC clock enable
	BDRST    BDCR = 0x01 << 16 //+ Backup domain software reset
	LSCOEN   BDCR = 0x01 << 24 //+ Low speed clock output enable
	LSCOSEL  BDCR = 0x01 << 25 //+ Low speed clock output selection
)

const (
	LSEONn    = 0
	LSERDYn   = 1
	LSEBYPn   = 2
	LSEDRVn   = 3
	LSECSSONn = 5
	LSECSSDn  = 6
	RTCSELn   = 8
	RTCENn    = 15
	BDRSTn    = 16
	LSCOENn   = 24
	LSCOSELn  = 25
)

const (
	LSION        CSR = 0x01 << 0  //+ LSI oscillator enable
	LSIRDY       CSR = 0x01 << 1  //+ LSI oscillator ready
	MSISRANGE    CSR = 0x0F << 8  //+ SI range after Standby mode
	RMVF         CSR = 0x01 << 23 //+ Remove reset flag
	FIREWALLRSTF CSR = 0x01 << 24 //+ Firewall reset flag
	OBLRSTF      CSR = 0x01 << 25 //+ Option byte loader reset flag
	PINRSTF      CSR = 0x01 << 26 //+ Pin reset flag
	BORRSTF      CSR = 0x01 << 27 //+ BOR flag
	SFTRSTF      CSR = 0x01 << 28 //+ Software reset flag
	IWDGRSTF     CSR = 0x01 << 29 //+ Independent window watchdog reset flag
	WWDGRSTF     CSR = 0x01 << 30 //+ Window watchdog reset flag
	LPWRSTF      CSR = 0x01 << 31 //+ Low-power reset flag
)

const (
	LSIONn        = 0
	LSIRDYn       = 1
	MSISRANGEn    = 8
	RMVFn         = 23
	FIREWALLRSTFn = 24
	OBLRSTFn      = 25
	PINRSTFn      = 26
	BORRSTFn      = 27
	SFTRSTFn      = 28
	IWDGRSTFn     = 29
	WWDGRSTFn     = 30
	LPWRSTFn      = 31
)
