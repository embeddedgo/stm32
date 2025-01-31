// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package sdmmc provides access to the registers of the SDMMC peripheral.
//
// Instances:
//
//	SDMMC1  SDMMC1_BASE  AHB3  SDMMC1+,SDMMC+  SDMMC1
//	SDMMC2  SDMMC2_BASE  AHB2  SDMMC1+,SDMMC+  SDMMC1
//
// Registers:
//
//	0x000 32  POWER       SDMMC power control register
//	0x004 32  CLKCR       The SDMMC_CLKCR register controls the SDMMC_CK output clock, the SDMMC_RX_CLK receive clock, and the bus width.
//	0x008 32  ARGR        The SDMMC_ARGR register contains a 32-bit command argument, which is sent to a card as part of a command message.
//	0x00C 32  CMDR        The SDMMC_CMDR register contains the command index and command type bits. The command index is sent to a card as part of a command message. The command type bits control the command path state machine (CPSM).
//	0x010 32  RESPCMDR    SDMMC command response register
//	0x014 32  RESP1R      The SDMMC_RESP1/2/3/4R registers contain the status of a card, which is part of the received response.
//	0x018 32  RESP2R      The SDMMC_RESP1/2/3/4R registers contain the status of a card, which is part of the received response.
//	0x01C 32  RESP3R      The SDMMC_RESP1/2/3/4R registers contain the status of a card, which is part of the received response.
//	0x020 32  RESP4R      The SDMMC_RESP1/2/3/4R registers contain the status of a card, which is part of the received response.
//	0x024 32  DTIMER      The SDMMC_DTIMER register contains the data timeout period, in card bus clock periods. A counter loads the value from the SDMMC_DTIMER register, and starts decrementing when the data path state machine (DPSM) enters the Wait_R or Busy state. If the timer reaches 0 while the DPSM is in either of these states, the timeout status flag is set.
//	0x028 32  DLENR       The SDMMC_DLENR register contains the number of data bytes to be transferred. The value is loaded into the data counter when data transfer starts.
//	0x02C 32  DCTRL       The SDMMC_DCTRL register control the data path state machine (DPSM).
//	0x030 32  DCNTR       The SDMMC_DCNTR register loads the value from the data length register (see SDMMC_DLENR) when the DPSM moves from the Idle state to the Wait_R or Wait_S state. As data is transferred, the counter decrements the value until it reaches 0. The DPSM then moves to the Idle state and when there has been no error, the data status end flag (DATAEND) is set.
//	0x034 32  STAR        The SDMMC_STAR register is a read-only register. It contains two types of flag:Static flags (bits [29,21,11:0]): these bits remain asserted until they are cleared by writing to the SDMMC interrupt Clear register (see SDMMC_ICR)Dynamic flags (bits [20:12]): these bits change state depending on the state of the underlying logic (for example, FIFO full and empty flags are asserted and de-asserted as data while written to the FIFO)
//	0x038 32  ICR         The SDMMC_ICR register is a write-only register. Writing a bit with 1 clears the corresponding bit in the SDMMC_STAR status register.
//	0x03C 32  MASKR       The interrupt mask register determines which status flags generate an interrupt request by setting the corresponding bit to 1.
//	0x040 32  ACKTIMER    The SDMMC_ACKTIMER register contains the acknowledgment timeout period, in SDMMC_CK bus clock periods. A counter loads the value from the SDMMC_ACKTIMER register, and starts decrementing when the data path state machine (DPSM) enters the Wait_Ack state. If the timer reaches 0 while the DPSM is in this states, the acknowledgment timeout status flag is set.
//	0x050 32  IDMACTRLR   The receive and transmit FIFOs can be read or written as 32-bit wide registers. The FIFOs contain 32 entries on 32 sequential addresses. This allows the CPU to use its load and store multiple operands to read from/write to the FIFO.
//	0x054 32  IDMABSIZER  The SDMMC_IDMABSIZER register contains the buffers size when in double buffer configuration.
//	0x058 32  IDMABASE0R  The SDMMC_IDMABASE0R register contains the memory buffer base address in single buffer configuration and the buffer 0 base address in double buffer configuration.
//	0x05C 32  IDMABASE1R  The SDMMC_IDMABASE1R register contains the double buffer configuration second buffer memory base address.
//	0x080 32  FIFOR       The receive and transmit FIFOs can be only read or written as word (32-bit) wide registers. The FIFOs contain 16 entries on sequential addresses. This allows the CPU to use its load and store multiple operands to read from/write to the FIFO.When accessing SDMMC_FIFOR with half word or byte access an AHB bus fault is generated.
//	0x3F4 32  VER         SDMMC IP version register
//	0x3F8 32  ID          SDMMC IP identification register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/bus
//	github.com/embeddedgo/stm32/p/mmap
package sdmmc

const (
	PWRCTRL   POWER = 0x03 << 0 //+ SDMMC state control bits. These bits can only be written when the SDMMC is not in the power-on state (PWRCTRL?11). These bits are used to define the functional state of the SDMMC signals: Any further write will be ignored, PWRCTRL value will keep 11.
	VSWITCH   POWER = 0x01 << 2 //+ Voltage switch sequence start. This bit is used to start the timing critical section of the voltage switch sequence:
	VSWITCHEN POWER = 0x01 << 3 //+ Voltage switch procedure enable. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). This bit is used to stop the SDMMC_CK after the voltage switch command response:
	DIRPOL    POWER = 0x01 << 4 //+ Data and command direction signals polarity selection. This bit can only be written when the SDMMC is in the power-off state (PWRCTRL = 00).
)

const (
	PWRCTRLn   = 0
	VSWITCHn   = 2
	VSWITCHENn = 3
	DIRPOLn    = 4
)

const (
	CLKDIV   CLKCR = 0x3FF << 0 //+ Clock divide factor This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0). This field defines the divide factor between the input clock (SDMMCCLK) and the output clock (SDMMC_CK): SDMMC_CK frequency = SDMMCCLK / [2 * CLKDIV]. 0xx: etc.. xxx: etc..
	PWRSAV   CLKCR = 0x01 << 12 //+ Power saving configuration bit This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0) For power saving, the SDMMC_CK clock output can be disabled when the bus is idle by setting PWRSAV:
	WIDBUS   CLKCR = 0x03 << 14 //+ Wide bus mode enable bit This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0)
	NEGEDGE  CLKCR = 0x01 << 16 //+ SDMMC_CK dephasing selection bit for data and Command. This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0). When clock division = 1 (CLKDIV = 0), this bit has no effect. Data and Command change on SDMMC_CK falling edge. When clock division &gt;1 (CLKDIV &gt; 0) &amp; DDR = 0: - SDMMC_CK edge occurs on SDMMCCLK rising edge. When clock division >1 (CLKDIV > 0) & DDR = 1: - Data changed on the SDMMCCLK falling edge succeeding a SDMMC_CK edge. - SDMMC_CK edge occurs on SDMMCCLK rising edge. - Data changed on the SDMMC_CK falling edge succeeding a SDMMC_CK edge. - SDMMC_CK edge occurs on SDMMCCLK rising edge.
	HWFC_EN  CLKCR = 0x01 << 17 //+ Hardware flow control enable This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0) When Hardware flow control is enabled, the meaning of the TXFIFOE and RXFIFOF flags change, please see SDMMC status register definition in Section56.8.11.
	DDR      CLKCR = 0x01 << 18 //+ Data rate signaling selection This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0) DDR rate shall only be selected with 4-bit or 8-bit wide bus mode. (WIDBUS &gt; 00). DDR = 1 has no effect when WIDBUS = 00 (1-bit wide bus). DDR rate shall only be selected with clock division &gt;1. (CLKDIV &gt; 0)
	BUSSPEED CLKCR = 0x01 << 19 //+ Bus speed mode selection between DS, HS, SDR12, SDR25 and SDR50, DDR50, SDR104. This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0)
	SELCLKRX CLKCR = 0x03 << 20 //+ Receive clock selection. These bits can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0)
)

const (
	CLKDIVn   = 0
	PWRSAVn   = 12
	WIDBUSn   = 14
	NEGEDGEn  = 16
	HWFC_ENn  = 17
	DDRn      = 18
	BUSSPEEDn = 19
	SELCLKRXn = 20
)

const (
	CMDARG ARGR = 0xFFFFFFFF << 0 //+ Command argument. These bits can only be written by firmware when CPSM is disabled (CPSMEN = 0). Command argument sent to a card as part of a command message. If a command contains an argument, it must be loaded into this register before writing a command to the command register.
)

const (
	CMDARGn = 0
)

const (
	CMDINDEX   CMDR = 0x3F << 0  //+ Command index. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). The command index is sent to the card as part of a command message.
	CMDTRANS   CMDR = 0x01 << 6  //+ The CPSM treats the command as a data transfer command, stops the interrupt period, and signals DataEnable to the DPSM This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). If this bit is set, the CPSM issues an end of interrupt period and issues DataEnable signal to the DPSM when the command is sent.
	CMDSTOP    CMDR = 0x01 << 7  //+ The CPSM treats the command as a Stop Transmission command and signals Abort to the DPSM. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). If this bit is set, the CPSM issues the Abort signal to the DPSM when the command is sent.
	WAITRESP   CMDR = 0x03 << 8  //+ Wait for response bits. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). They are used to configure whether the CPSM is to wait for a response, and if yes, which kind of response.
	WAITINT    CMDR = 0x01 << 10 //+ CPSM waits for interrupt request. If this bit is set, the CPSM disables command timeout and waits for an card interrupt request (Response). If this bit is cleared in the CPSM Wait state, will cause the abort of the interrupt mode.
	WAITPEND   CMDR = 0x01 << 11 //+ CPSM Waits for end of data transfer (CmdPend internal signal) from DPSM. This bit when set, the CPSM waits for the end of data transfer trigger before it starts sending a command. WAITPEND is only taken into account when DTMODE = MMC stream data transfer, WIDBUS = 1-bit wide bus mode, DPSMACT = 1 and DTDIR = from host to card.
	CPSMEN     CMDR = 0x01 << 12 //+ Command path state machine (CPSM) Enable bit This bit is written 1 by firmware, and cleared by hardware when the CPSM enters the Idle state. If this bit is set, the CPSM is enabled. When DTEN = 1, no command will be transfered nor boot procedure will be started. CPSMEN is cleared to 0.
	DTHOLD     CMDR = 0x01 << 13 //+ Hold new data block transmission and reception in the DPSM. If this bit is set, the DPSM will not move from the Wait_S state to the Send state or from the Wait_R state to the Receive state.
	BOOTMODE   CMDR = 0x01 << 14 //+ Select the boot mode procedure to be used. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0)
	BOOTEN     CMDR = 0x01 << 15 //+ Enable boot mode procedure.
	CMDSUSPEND CMDR = 0x01 << 16 //+ The CPSM treats the command as a Suspend or Resume command and signals interrupt period start/end. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). CMDSUSPEND = 1 and CMDTRANS = 0 Suspend command, start interrupt period when response bit BS=0. CMDSUSPEND = 1 and CMDTRANS = 1 Resume command with data, end interrupt period when response bit DF=1.
)

const (
	CMDINDEXn   = 0
	CMDTRANSn   = 6
	CMDSTOPn    = 7
	WAITRESPn   = 8
	WAITINTn    = 10
	WAITPENDn   = 11
	CPSMENn     = 12
	DTHOLDn     = 13
	BOOTMODEn   = 14
	BOOTENn     = 15
	CMDSUSPENDn = 16
)

const (
	RESPCMD RESPCMDR = 0x3F << 0 //+ Response command index
)

const (
	RESPCMDn = 0
)

const (
	CARDSTATUS1 RESP1R = 0xFFFFFFFF << 0 //+ see Table 432
)

const (
	CARDSTATUS1n = 0
)

const (
	CARDSTATUS2 RESP2R = 0xFFFFFFFF << 0 //+ see Table404.
)

const (
	CARDSTATUS2n = 0
)

const (
	CARDSTATUS3 RESP3R = 0xFFFFFFFF << 0 //+ see Table404.
)

const (
	CARDSTATUS3n = 0
)

const (
	CARDSTATUS4 RESP4R = 0xFFFFFFFF << 0 //+ see Table404.
)

const (
	CARDSTATUS4n = 0
)

const (
	DATATIME DTIMER = 0xFFFFFFFF << 0 //+ Data and R1b busy timeout period This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0). Data and R1b busy timeout period expressed in card bus clock periods.
)

const (
	DATATIMEn = 0
)

const (
	DATALENGTH DLENR = 0x1FFFFFF << 0 //+ Data length value This register can only be written by firmware when DPSM is inactive (DPSMACT = 0). Number of data bytes to be transferred. When DDR = 1 DATALENGTH is truncated to a multiple of 2. (The last odd byte is not transfered) When DATALENGTH = 0 no data will be transfered, when requested by a CPSMEN and CMDTRANS = 1 also no command will be transfered. DTEN and CPSMEN are cleared to 0.
)

const (
	DATALENGTHn = 0
)

const (
	DTEN       DCTRL = 0x01 << 0  //+ Data transfer enable bit This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). This bit is cleared by Hardware when data transfer completes. This bit shall only be used to transfer data when no associated data transfer command is used, i.e. shall not be used with SD or eMMC cards.
	DTDIR      DCTRL = 0x01 << 1  //+ Data transfer direction selection This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
	DTMODE     DCTRL = 0x03 << 2  //+ Data transfer mode selection. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
	DBLOCKSIZE DCTRL = 0x0F << 4  //+ Data block size This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). Define the data block length when the block data transfer mode is selected: When DATALENGTH is not a multiple of DBLOCKSIZE, the transfered data is truncated at a multiple of DBLOCKSIZE. (Any remain data will not be transfered.) When DDR = 1, DBLOCKSIZE = 0000 shall not be used. (No data will be transfered)
	RWSTART    DCTRL = 0x01 << 8  //+ Read wait start. If this bit is set, read wait operation starts.
	RWSTOP     DCTRL = 0x01 << 9  //+ Read wait stop This bit is written by firmware and auto cleared by hardware when the DPSM moves from the READ_WAIT state to the WAIT_R or IDLE state.
	RWMOD      DCTRL = 0x01 << 10 //+ Read wait mode. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
	SDIOEN     DCTRL = 0x01 << 11 //+ SD I/O interrupt enable functions This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). If this bit is set, the DPSM enables the SD I/O card specific interrupt operation.
	BOOTACKEN  DCTRL = 0x01 << 12 //+ Enable the reception of the boot acknowledgment. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
	FIFORST    DCTRL = 0x01 << 13 //+ FIFO reset, will flush any remaining data. This bit can only be written by firmware when IDMAEN= 0 and DPSM is active (DPSMACT = 1). This bit will only take effect when a transfer error or transfer hold occurs.
)

const (
	DTENn       = 0
	DTDIRn      = 1
	DTMODEn     = 2
	DBLOCKSIZEn = 4
	RWSTARTn    = 8
	RWSTOPn     = 9
	RWMODn      = 10
	SDIOENn     = 11
	BOOTACKENn  = 12
	FIFORSTn    = 13
)

const (
	DATACOUNT DCNTR = 0x1FFFFFF << 0 //+ Data count value When read, the number of remaining data bytes to be transferred is returned. Write has no effect.
)

const (
	DATACOUNTn = 0
)

const (
	CCRCFAIL   STAR = 0x01 << 0  //+ Command response received (CRC check failed). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	DCRCFAIL   STAR = 0x01 << 1  //+ Data block sent/received (CRC check failed). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	CTIMEOUT   STAR = 0x01 << 2  //+ Command response timeout. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR. The Command Timeout period has a fixed value of 64 SDMMC_CK clock periods.
	DTIMEOUT   STAR = 0x01 << 3  //+ Data timeout. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	TXUNDERR   STAR = 0x01 << 4  //+ Transmit FIFO underrun error or IDMA read transfer error. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	RXOVERR    STAR = 0x01 << 5  //+ Received FIFO overrun error or IDMA write transfer error. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	CMDREND    STAR = 0x01 << 6  //+ Command response received (CRC check passed, or no CRC). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	CMDSENT    STAR = 0x01 << 7  //+ Command sent (no response required). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	DATAEND    STAR = 0x01 << 8  //+ Data transfer ended correctly. (data counter, DATACOUNT is zero and no errors occur). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	DHOLD      STAR = 0x01 << 9  //+ Data transfer Hold. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	DBCKEND    STAR = 0x01 << 10 //+ Data block sent/received. (CRC check passed) and DPSM moves to the READWAIT state. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	DABORT     STAR = 0x01 << 11 //+ Data transfer aborted by CMD12. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	DPSMACT    STAR = 0x01 << 12 //+ Data path state machine active, i.e. not in Idle state. This is a hardware status flag only, does not generate an interrupt.
	CPSMACT    STAR = 0x01 << 13 //+ Command path state machine active, i.e. not in Idle state. This is a hardware status flag only, does not generate an interrupt.
	TXFIFOHE   STAR = 0x01 << 14 //+ Transmit FIFO half empty At least half the number of words can be written into the FIFO. This bit is cleared when the FIFO becomes half+1 full.
	RXFIFOHF   STAR = 0x01 << 15 //+ Receive FIFO half full There are at least half the number of words in the FIFO. This bit is cleared when the FIFO becomes half+1 empty.
	TXFIFOF    STAR = 0x01 << 16 //+ Transmit FIFO full This is a hardware status flag only, does not generate an interrupt. This bit is cleared when one FIFO location becomes empty.
	RXFIFOF    STAR = 0x01 << 17 //+ Receive FIFO full This bit is cleared when one FIFO location becomes empty.
	TXFIFOE    STAR = 0x01 << 18 //+ Transmit FIFO empty This bit is cleared when one FIFO location becomes full.
	RXFIFOE    STAR = 0x01 << 19 //+ Receive FIFO empty This is a hardware status flag only, does not generate an interrupt. This bit is cleared when one FIFO location becomes full.
	BUSYD0     STAR = 0x01 << 20 //+ Inverted value of SDMMC_D0 line (Busy), sampled at the end of a CMD response and a second time 2 SDMMC_CK cycles after the CMD response. This bit is reset to not busy when the SDMMCD0 line changes from busy to not busy. This bit does not signal busy due to data transfer. This is a hardware status flag only, it does not generate an interrupt.
	BUSYD0END  STAR = 0x01 << 21 //+ end of SDMMC_D0 Busy following a CMD response detected. This indicates only end of busy following a CMD response. This bit does not signal busy due to data transfer. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	SDIOIT     STAR = 0x01 << 22 //+ SDIO interrupt received. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	ACKFAIL    STAR = 0x01 << 23 //+ Boot acknowledgment received (boot acknowledgment check fail). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	ACKTIMEOUT STAR = 0x01 << 24 //+ Boot acknowledgment timeout. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	VSWEND     STAR = 0x01 << 25 //+ Voltage switch critical timing section completion. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	CKSTOP     STAR = 0x01 << 26 //+ SDMMC_CK stopped in Voltage switch procedure. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	IDMATE     STAR = 0x01 << 27 //+ IDMA transfer error. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
	IDMABTC    STAR = 0x01 << 28 //+ IDMA buffer transfer complete. interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
)

const (
	CCRCFAILn   = 0
	DCRCFAILn   = 1
	CTIMEOUTn   = 2
	DTIMEOUTn   = 3
	TXUNDERRn   = 4
	RXOVERRn    = 5
	CMDRENDn    = 6
	CMDSENTn    = 7
	DATAENDn    = 8
	DHOLDn      = 9
	DBCKENDn    = 10
	DABORTn     = 11
	DPSMACTn    = 12
	CPSMACTn    = 13
	TXFIFOHEn   = 14
	RXFIFOHFn   = 15
	TXFIFOFn    = 16
	RXFIFOFn    = 17
	TXFIFOEn    = 18
	RXFIFOEn    = 19
	BUSYD0n     = 20
	BUSYD0ENDn  = 21
	SDIOITn     = 22
	ACKFAILn    = 23
	ACKTIMEOUTn = 24
	VSWENDn     = 25
	CKSTOPn     = 26
	IDMATEn     = 27
	IDMABTCn    = 28
)

const (
	CCRCFAILC   ICR = 0x01 << 0  //+ CCRCFAIL flag clear bit Set by software to clear the CCRCFAIL flag.
	DCRCFAILC   ICR = 0x01 << 1  //+ DCRCFAIL flag clear bit Set by software to clear the DCRCFAIL flag.
	CTIMEOUTC   ICR = 0x01 << 2  //+ CTIMEOUT flag clear bit Set by software to clear the CTIMEOUT flag.
	DTIMEOUTC   ICR = 0x01 << 3  //+ DTIMEOUT flag clear bit Set by software to clear the DTIMEOUT flag.
	TXUNDERRC   ICR = 0x01 << 4  //+ TXUNDERR flag clear bit Set by software to clear TXUNDERR flag.
	RXOVERRC    ICR = 0x01 << 5  //+ RXOVERR flag clear bit Set by software to clear the RXOVERR flag.
	CMDRENDC    ICR = 0x01 << 6  //+ CMDREND flag clear bit Set by software to clear the CMDREND flag.
	CMDSENTC    ICR = 0x01 << 7  //+ CMDSENT flag clear bit Set by software to clear the CMDSENT flag.
	DATAENDC    ICR = 0x01 << 8  //+ DATAEND flag clear bit Set by software to clear the DATAEND flag.
	DHOLDC      ICR = 0x01 << 9  //+ DHOLD flag clear bit Set by software to clear the DHOLD flag.
	DBCKENDC    ICR = 0x01 << 10 //+ DBCKEND flag clear bit Set by software to clear the DBCKEND flag.
	DABORTC     ICR = 0x01 << 11 //+ DABORT flag clear bit Set by software to clear the DABORT flag.
	BUSYD0ENDC  ICR = 0x01 << 21 //+ BUSYD0END flag clear bit Set by software to clear the BUSYD0END flag.
	SDIOITC     ICR = 0x01 << 22 //+ SDIOIT flag clear bit Set by software to clear the SDIOIT flag.
	ACKFAILC    ICR = 0x01 << 23 //+ ACKFAIL flag clear bit Set by software to clear the ACKFAIL flag.
	ACKTIMEOUTC ICR = 0x01 << 24 //+ ACKTIMEOUT flag clear bit Set by software to clear the ACKTIMEOUT flag.
	VSWENDC     ICR = 0x01 << 25 //+ VSWEND flag clear bit Set by software to clear the VSWEND flag.
	CKSTOPC     ICR = 0x01 << 26 //+ CKSTOP flag clear bit Set by software to clear the CKSTOP flag.
	IDMATEC     ICR = 0x01 << 27 //+ IDMA transfer error clear bit Set by software to clear the IDMATE flag.
	IDMABTCC    ICR = 0x01 << 28 //+ IDMA buffer transfer complete clear bit Set by software to clear the IDMABTC flag.
)

const (
	CCRCFAILCn   = 0
	DCRCFAILCn   = 1
	CTIMEOUTCn   = 2
	DTIMEOUTCn   = 3
	TXUNDERRCn   = 4
	RXOVERRCn    = 5
	CMDRENDCn    = 6
	CMDSENTCn    = 7
	DATAENDCn    = 8
	DHOLDCn      = 9
	DBCKENDCn    = 10
	DABORTCn     = 11
	BUSYD0ENDCn  = 21
	SDIOITCn     = 22
	ACKFAILCn    = 23
	ACKTIMEOUTCn = 24
	VSWENDCn     = 25
	CKSTOPCn     = 26
	IDMATECn     = 27
	IDMABTCCn    = 28
)

const (
	CCRCFAILIE   MASKR = 0x01 << 0  //+ Command CRC fail interrupt enable Set and cleared by software to enable/disable interrupt caused by command CRC failure.
	DCRCFAILIE   MASKR = 0x01 << 1  //+ Data CRC fail interrupt enable Set and cleared by software to enable/disable interrupt caused by data CRC failure.
	CTIMEOUTIE   MASKR = 0x01 << 2  //+ Command timeout interrupt enable Set and cleared by software to enable/disable interrupt caused by command timeout.
	DTIMEOUTIE   MASKR = 0x01 << 3  //+ Data timeout interrupt enable Set and cleared by software to enable/disable interrupt caused by data timeout.
	TXUNDERRIE   MASKR = 0x01 << 4  //+ Tx FIFO underrun error interrupt enable Set and cleared by software to enable/disable interrupt caused by Tx FIFO underrun error.
	RXOVERRIE    MASKR = 0x01 << 5  //+ Rx FIFO overrun error interrupt enable Set and cleared by software to enable/disable interrupt caused by Rx FIFO overrun error.
	CMDRENDIE    MASKR = 0x01 << 6  //+ Command response received interrupt enable Set and cleared by software to enable/disable interrupt caused by receiving command response.
	CMDSENTIE    MASKR = 0x01 << 7  //+ Command sent interrupt enable Set and cleared by software to enable/disable interrupt caused by sending command.
	DATAENDIE    MASKR = 0x01 << 8  //+ Data end interrupt enable Set and cleared by software to enable/disable interrupt caused by data end.
	DHOLDIE      MASKR = 0x01 << 9  //+ Data hold interrupt enable Set and cleared by software to enable/disable the interrupt generated when sending new data is hold in the DPSM Wait_S state.
	DBCKENDIE    MASKR = 0x01 << 10 //+ Data block end interrupt enable Set and cleared by software to enable/disable interrupt caused by data block end.
	DABORTIE     MASKR = 0x01 << 11 //+ Data transfer aborted interrupt enable Set and cleared by software to enable/disable interrupt caused by a data transfer being aborted.
	TXFIFOHEIE   MASKR = 0x01 << 14 //+ Tx FIFO half empty interrupt enable Set and cleared by software to enable/disable interrupt caused by Tx FIFO half empty.
	RXFIFOHFIE   MASKR = 0x01 << 15 //+ Rx FIFO half full interrupt enable Set and cleared by software to enable/disable interrupt caused by Rx FIFO half full.
	RXFIFOFIE    MASKR = 0x01 << 17 //+ Rx FIFO full interrupt enable Set and cleared by software to enable/disable interrupt caused by Rx FIFO full.
	TXFIFOEIE    MASKR = 0x01 << 18 //+ Tx FIFO empty interrupt enable Set and cleared by software to enable/disable interrupt caused by Tx FIFO empty.
	BUSYD0ENDIE  MASKR = 0x01 << 21 //+ BUSYD0END interrupt enable Set and cleared by software to enable/disable the interrupt generated when SDMMC_D0 signal changes from busy to NOT busy following a CMD response.
	SDIOITIE     MASKR = 0x01 << 22 //+ SDIO mode interrupt received interrupt enable Set and cleared by software to enable/disable the interrupt generated when receiving the SDIO mode interrupt.
	ACKFAILIE    MASKR = 0x01 << 23 //+ Acknowledgment Fail interrupt enable Set and cleared by software to enable/disable interrupt caused by acknowledgment Fail.
	ACKTIMEOUTIE MASKR = 0x01 << 24 //+ Acknowledgment timeout interrupt enable Set and cleared by software to enable/disable interrupt caused by acknowledgment timeout.
	VSWENDIE     MASKR = 0x01 << 25 //+ Voltage switch critical timing section completion interrupt enable Set and cleared by software to enable/disable the interrupt generated when voltage switch critical timing section completion.
	CKSTOPIE     MASKR = 0x01 << 26 //+ Voltage Switch clock stopped interrupt enable Set and cleared by software to enable/disable interrupt caused by Voltage Switch clock stopped.
	IDMABTCIE    MASKR = 0x01 << 28 //+ IDMA buffer transfer complete interrupt enable Set and cleared by software to enable/disable the interrupt generated when the IDMA has transferred all data belonging to a memory buffer.
)

const (
	CCRCFAILIEn   = 0
	DCRCFAILIEn   = 1
	CTIMEOUTIEn   = 2
	DTIMEOUTIEn   = 3
	TXUNDERRIEn   = 4
	RXOVERRIEn    = 5
	CMDRENDIEn    = 6
	CMDSENTIEn    = 7
	DATAENDIEn    = 8
	DHOLDIEn      = 9
	DBCKENDIEn    = 10
	DABORTIEn     = 11
	TXFIFOHEIEn   = 14
	RXFIFOHFIEn   = 15
	RXFIFOFIEn    = 17
	TXFIFOEIEn    = 18
	BUSYD0ENDIEn  = 21
	SDIOITIEn     = 22
	ACKFAILIEn    = 23
	ACKTIMEOUTIEn = 24
	VSWENDIEn     = 25
	CKSTOPIEn     = 26
	IDMABTCIEn    = 28
)

const (
	ACKTIME ACKTIMER = 0x1FFFFFF << 0 //+ Boot acknowledgment timeout period This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). Boot acknowledgment timeout period expressed in card bus clock periods.
)

const (
	ACKTIMEn = 0
)

const (
	IDMAEN    IDMACTRLR = 0x01 << 0 //+ IDMA enable This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
	IDMABMODE IDMACTRLR = 0x01 << 1 //+ Buffer mode selection. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
	IDMABACT  IDMACTRLR = 0x01 << 2 //+ Double buffer mode active buffer indication This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). When IDMA is enabled this bit is toggled by hardware.
)

const (
	IDMAENn    = 0
	IDMABMODEn = 1
	IDMABACTn  = 2
)

const (
	IDMABNDT IDMABSIZER = 0xFF << 5 //+ Number of transfers per buffer. This 8-bit value shall be multiplied by 8 to get the size of the buffer in 32-bit words and by 32 to get the size of the buffer in bytes. Example: IDMABNDT = 0x01: buffer size = 8 words = 32 bytes. These bits can only be written by firmware when DPSM is inactive (DPSMACT = 0).
)

const (
	IDMABNDTn = 5
)

const (
	IDMABASE0 IDMABASE0R = 0xFFFFFFFF << 0 //+ Buffer 0 memory base address bits [31:2], shall be word aligned (bit [1:0] are always 0 and read only). This register can be written by firmware when DPSM is inactive (DPSMACT = 0), and can dynamically be written by firmware when DPSM active (DPSMACT = 1) and memory buffer 0 is inactive (IDMABACT = 1).
)

const (
	IDMABASE0n = 0
)

const (
	IDMABASE1 IDMABASE1R = 0xFFFFFFFF << 0 //+ Buffer 1 memory base address, shall be word aligned (bit [1:0] are always 0 and read only). This register can be written by firmware when DPSM is inactive (DPSMACT = 0), and can dynamically be written by firmware when DPSM active (DPSMACT = 1) and memory buffer 1 is inactive (IDMABACT = 0).
)

const (
	IDMABASE1n = 0
)

const (
	FIFODATA FIFOR = 0xFFFFFFFF << 0 //+ Receive and transmit FIFO data This register can only be read or written by firmware when the DPSM is active (DPSMACT=1). The FIFO data occupies 16 entries of 32-bit words.
)

const (
	FIFODATAn = 0
)

const (
	MINREV VER = 0x0F << 0 //+ IP minor revision number.
	MAJREV VER = 0x0F << 4 //+ IP major revision number.
)

const (
	MINREVn = 0
	MAJREVn = 4
)

const (
	IP_ID ID = 0xFFFFFFFF << 0 //+ SDMMC IP identification.
)

const (
	IP_IDn = 0
)
