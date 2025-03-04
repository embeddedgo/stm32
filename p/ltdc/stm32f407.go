// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f407

// Package ltdc provides access to the registers of the LTDC peripheral.
//
// Instances:
//
//	LTDC  LTDC_BASE  -  LCD_TFT,LCD_TFT_1  LCD-TFT Controller
//
// Registers:
//
//	0x008 32  SSCR      Synchronization Size Configuration Register
//	0x00C 32  BPCR      Back Porch Configuration Register
//	0x010 32  AWCR      Active Width Configuration Register
//	0x014 32  TWCR      Total Width Configuration Register
//	0x018 32  GCR       Global Control Register
//	0x024 32  SRCR      Shadow Reload Configuration Register
//	0x02C 32  BCCR      Background Color Configuration Register
//	0x034 32  IER       Interrupt Enable Register
//	0x038 32  ISR       Interrupt Status Register
//	0x03C 32  ICR       Interrupt Clear Register
//	0x040 32  LIPCR     Line Interrupt Position Configuration Register
//	0x044 32  CPSR      Current Position Status Register
//	0x048 32  CDSR      Current Display Status Register
//	0x084 32  L1CR      Layerx Control Register
//	0x088 32  L1WHPCR   Layerx Window Horizontal Position Configuration Register
//	0x08C 32  L1WVPCR   Layerx Window Vertical Position Configuration Register
//	0x090 32  L1CKCR    Layerx Color Keying Configuration Register
//	0x094 32  L1PFCR    Layerx Pixel Format Configuration Register
//	0x098 32  L1CACR    Layerx Constant Alpha Configuration Register
//	0x09C 32  L1DCCR    Layerx Default Color Configuration Register
//	0x0A0 32  L1BFCR    Layerx Blending Factors Configuration Register
//	0x0AC 32  L1CFBAR   Layerx Color Frame Buffer Address Register
//	0x0B0 32  L1CFBLR   Layerx Color Frame Buffer Length Register
//	0x0B4 32  L1CFBLNR  Layerx ColorFrame Buffer Line Number Register
//	0x0C4 32  L1CLUTWR  Layerx CLUT Write Register
//	0x104 32  L2CR      Layerx Control Register
//	0x108 32  L2WHPCR   Layerx Window Horizontal Position Configuration Register
//	0x10C 32  L2WVPCR   Layerx Window Vertical Position Configuration Register
//	0x110 32  L2CKCR    Layerx Color Keying Configuration Register
//	0x114 32  L2PFCR    Layerx Pixel Format Configuration Register
//	0x118 32  L2CACR    Layerx Constant Alpha Configuration Register
//	0x11C 32  L2DCCR    Layerx Default Color Configuration Register
//	0x120 32  L2BFCR    Layerx Blending Factors Configuration Register
//	0x12C 32  L2CFBAR   Layerx Color Frame Buffer Address Register
//	0x130 32  L2CFBLR   Layerx Color Frame Buffer Length Register
//	0x134 32  L2CFBLNR  Layerx ColorFrame Buffer Line Number Register
//	0x144 32  L2CLUTWR  Layerx CLUT Write Register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package ltdc

const (
	VSH SSCR = 0x7FF << 0  //+ Vertical Synchronization Height (in units of horizontal scan line)
	HSW SSCR = 0x3FF << 16 //+ Horizontal Synchronization Width (in units of pixel clock period)
)

const (
	VSHn = 0
	HSWn = 16
)

const (
	AVBP BPCR = 0x7FF << 0  //+ Accumulated Vertical back porch (in units of horizontal scan line)
	AHBP BPCR = 0x3FF << 16 //+ Accumulated Horizontal back porch (in units of pixel clock period)
)

const (
	AVBPn = 0
	AHBPn = 16
)

const (
	AAH AWCR = 0x7FF << 0  //+ Accumulated Active Height (in units of horizontal scan line)
	AAV AWCR = 0x3FF << 16 //+ AAV
)

const (
	AAHn = 0
	AAVn = 16
)

const (
	TOTALH TWCR = 0x7FF << 0  //+ Total Height (in units of horizontal scan line)
	TOTALW TWCR = 0x3FF << 16 //+ Total Width (in units of pixel clock period)
)

const (
	TOTALHn = 0
	TOTALWn = 16
)

const (
	LTDCEN GCR = 0x01 << 0  //+ LCD-TFT controller enable bit
	DBW    GCR = 0x07 << 4  //+ Dither Blue Width
	DGW    GCR = 0x07 << 8  //+ Dither Green Width
	DRW    GCR = 0x07 << 12 //+ Dither Red Width
	DEN    GCR = 0x01 << 16 //+ Dither Enable
	PCPOL  GCR = 0x01 << 28 //+ Pixel Clock Polarity
	DEPOL  GCR = 0x01 << 29 //+ Data Enable Polarity
	VSPOL  GCR = 0x01 << 30 //+ Vertical Synchronization Polarity
	HSPOL  GCR = 0x01 << 31 //+ Horizontal Synchronization Polarity
)

const (
	LTDCENn = 0
	DBWn    = 4
	DGWn    = 8
	DRWn    = 12
	DENn    = 16
	PCPOLn  = 28
	DEPOLn  = 29
	VSPOLn  = 30
	HSPOLn  = 31
)

const (
	IMR SRCR = 0x01 << 0 //+ Immediate Reload
	VBR SRCR = 0x01 << 1 //+ Vertical Blanking Reload
)

const (
	IMRn = 0
	VBRn = 1
)

const (
	BC BCCR = 0xFFFFFF << 0 //+ Background Color Red value
)

const (
	BCn = 0
)

const (
	LIE    IER = 0x01 << 0 //+ Line Interrupt Enable
	FUIE   IER = 0x01 << 1 //+ FIFO Underrun Interrupt Enable
	TERRIE IER = 0x01 << 2 //+ Transfer Error Interrupt Enable
	RRIE   IER = 0x01 << 3 //+ Register Reload interrupt enable
)

const (
	LIEn    = 0
	FUIEn   = 1
	TERRIEn = 2
	RRIEn   = 3
)

const (
	LIF    ISR = 0x01 << 0 //+ Line Interrupt flag
	FUIF   ISR = 0x01 << 1 //+ FIFO Underrun Interrupt flag
	TERRIF ISR = 0x01 << 2 //+ Transfer Error interrupt flag
	RRIF   ISR = 0x01 << 3 //+ Register Reload Interrupt Flag
)

const (
	LIFn    = 0
	FUIFn   = 1
	TERRIFn = 2
	RRIFn   = 3
)

const (
	CLIF    ICR = 0x01 << 0 //+ Clears the Line Interrupt Flag
	CFUIF   ICR = 0x01 << 1 //+ Clears the FIFO Underrun Interrupt flag
	CTERRIF ICR = 0x01 << 2 //+ Clears the Transfer Error Interrupt Flag
	CRRIF   ICR = 0x01 << 3 //+ Clears Register Reload Interrupt Flag
)

const (
	CLIFn    = 0
	CFUIFn   = 1
	CTERRIFn = 2
	CRRIFn   = 3
)

const (
	LIPOS LIPCR = 0x7FF << 0 //+ Line Interrupt Position
)

const (
	LIPOSn = 0
)

const (
	CYPOS CPSR = 0xFFFF << 0  //+ Current Y Position
	CXPOS CPSR = 0xFFFF << 16 //+ Current X Position
)

const (
	CYPOSn = 0
	CXPOSn = 16
)

const (
	VDES   CDSR = 0x01 << 0 //+ Vertical Data Enable display Status
	HDES   CDSR = 0x01 << 1 //+ Horizontal Data Enable display Status
	VSYNCS CDSR = 0x01 << 2 //+ Vertical Synchronization display Status
	HSYNCS CDSR = 0x01 << 3 //+ Horizontal Synchronization display Status
)

const (
	VDESn   = 0
	HDESn   = 1
	VSYNCSn = 2
	HSYNCSn = 3
)

const (
	LEN    L1CR = 0x01 << 0 //+ Layer Enable
	COLKEN L1CR = 0x01 << 1 //+ Color Keying Enable
	CLUTEN L1CR = 0x01 << 4 //+ Color Look-Up Table Enable
)

const (
	LENn    = 0
	COLKENn = 1
	CLUTENn = 4
)

const (
	WHSTPOS L1WHPCR = 0xFFF << 0  //+ Window Horizontal Start Position
	WHSPPOS L1WHPCR = 0xFFF << 16 //+ Window Horizontal Stop Position
)

const (
	WHSTPOSn = 0
	WHSPPOSn = 16
)

const (
	WVSTPOS L1WVPCR = 0x7FF << 0  //+ Window Vertical Start Position
	WVSPPOS L1WVPCR = 0x7FF << 16 //+ Window Vertical Stop Position
)

const (
	WVSTPOSn = 0
	WVSPPOSn = 16
)

const (
	CKBLUE  L1CKCR = 0xFF << 0  //+ Color Key Blue value
	CKGREEN L1CKCR = 0xFF << 8  //+ Color Key Green value
	CKRED   L1CKCR = 0xFF << 16 //+ Color Key Red value
)

const (
	CKBLUEn  = 0
	CKGREENn = 8
	CKREDn   = 16
)

const (
	PF L1PFCR = 0x07 << 0 //+ Pixel Format
)

const (
	PFn = 0
)

const (
	CONSTA L1CACR = 0xFF << 0 //+ Constant Alpha
)

const (
	CONSTAn = 0
)

const (
	DCBLUE  L1DCCR = 0xFF << 0  //+ Default Color Blue
	DCGREEN L1DCCR = 0xFF << 8  //+ Default Color Green
	DCRED   L1DCCR = 0xFF << 16 //+ Default Color Red
	DCALPHA L1DCCR = 0xFF << 24 //+ Default Color Alpha
)

const (
	DCBLUEn  = 0
	DCGREENn = 8
	DCREDn   = 16
	DCALPHAn = 24
)

const (
	BF2 L1BFCR = 0x07 << 0 //+ Blending Factor 2
	BF1 L1BFCR = 0x07 << 8 //+ Blending Factor 1
)

const (
	BF2n = 0
	BF1n = 8
)

const (
	CFBADD L1CFBAR = 0xFFFFFFFF << 0 //+ Color Frame Buffer Start Address
)

const (
	CFBADDn = 0
)

const (
	CFBLL L1CFBLR = 0x1FFF << 0  //+ Color Frame Buffer Line Length
	CFBP  L1CFBLR = 0x1FFF << 16 //+ Color Frame Buffer Pitch in bytes
)

const (
	CFBLLn = 0
	CFBPn  = 16
)

const (
	CFBLNBR L1CFBLNR = 0x7FF << 0 //+ Frame Buffer Line Number
)

const (
	CFBLNBRn = 0
)

const (
	BLUE    L1CLUTWR = 0xFF << 0  //+ Blue value
	GREEN   L1CLUTWR = 0xFF << 8  //+ Green value
	RED     L1CLUTWR = 0xFF << 16 //+ Red value
	CLUTADD L1CLUTWR = 0xFF << 24 //+ CLUT Address
)

const (
	BLUEn    = 0
	GREENn   = 8
	REDn     = 16
	CLUTADDn = 24
)

const (
	LEN    L2CR = 0x01 << 0 //+ Layer Enable
	COLKEN L2CR = 0x01 << 1 //+ Color Keying Enable
	CLUTEN L2CR = 0x01 << 4 //+ Color Look-Up Table Enable
)

const (
	LENn    = 0
	COLKENn = 1
	CLUTENn = 4
)

const (
	WHSTPOS L2WHPCR = 0xFFF << 0  //+ Window Horizontal Start Position
	WHSPPOS L2WHPCR = 0xFFF << 16 //+ Window Horizontal Stop Position
)

const (
	WHSTPOSn = 0
	WHSPPOSn = 16
)

const (
	WVSTPOS L2WVPCR = 0x7FF << 0  //+ Window Vertical Start Position
	WVSPPOS L2WVPCR = 0x7FF << 16 //+ Window Vertical Stop Position
)

const (
	WVSTPOSn = 0
	WVSPPOSn = 16
)

const (
	CKBLUE  L2CKCR = 0xFF << 0   //+ Color Key Blue value
	CKGREEN L2CKCR = 0x7F << 8   //+ Color Key Green value
	CKRED   L2CKCR = 0x1FF << 15 //+ Color Key Red value
)

const (
	CKBLUEn  = 0
	CKGREENn = 8
	CKREDn   = 15
)

const (
	PF L2PFCR = 0x07 << 0 //+ Pixel Format
)

const (
	PFn = 0
)

const (
	CONSTA L2CACR = 0xFF << 0 //+ Constant Alpha
)

const (
	CONSTAn = 0
)

const (
	DCBLUE  L2DCCR = 0xFF << 0  //+ Default Color Blue
	DCGREEN L2DCCR = 0xFF << 8  //+ Default Color Green
	DCRED   L2DCCR = 0xFF << 16 //+ Default Color Red
	DCALPHA L2DCCR = 0xFF << 24 //+ Default Color Alpha
)

const (
	DCBLUEn  = 0
	DCGREENn = 8
	DCREDn   = 16
	DCALPHAn = 24
)

const (
	BF2 L2BFCR = 0x07 << 0 //+ Blending Factor 2
	BF1 L2BFCR = 0x07 << 8 //+ Blending Factor 1
)

const (
	BF2n = 0
	BF1n = 8
)

const (
	CFBADD L2CFBAR = 0xFFFFFFFF << 0 //+ Color Frame Buffer Start Address
)

const (
	CFBADDn = 0
)

const (
	CFBLL L2CFBLR = 0x1FFF << 0  //+ Color Frame Buffer Line Length
	CFBP  L2CFBLR = 0x1FFF << 16 //+ Color Frame Buffer Pitch in bytes
)

const (
	CFBLLn = 0
	CFBPn  = 16
)

const (
	CFBLNBR L2CFBLNR = 0x7FF << 0 //+ Frame Buffer Line Number
)

const (
	CFBLNBRn = 0
)

const (
	BLUE    L2CLUTWR = 0xFF << 0  //+ Blue value
	GREEN   L2CLUTWR = 0xFF << 8  //+ Green value
	RED     L2CLUTWR = 0xFF << 16 //+ Red value
	CLUTADD L2CLUTWR = 0xFF << 24 //+ CLUT Address
)

const (
	BLUEn    = 0
	GREENn   = 8
	REDn     = 16
	CLUTADDn = 24
)
