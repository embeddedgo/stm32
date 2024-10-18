// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32l4x6

// Package otg_fs_device provides access to the registers of the OTG_FS_DEVICE peripheral.
//
// Instances:
//
//	OTG_FS_DEVICE  OTG_FS_DEVICE_BASE  -  -  USB on the go full speed
//
// Registers:
//
//	0x000 32  FS_DCFG      OTG_FS device configuration register (OTG_FS_DCFG)
//	0x004 32  FS_DCTL      OTG_FS device control register (OTG_FS_DCTL)
//	0x008 32  FS_DSTS      OTG_FS device status register (OTG_FS_DSTS)
//	0x010 32  FS_DIEPMSK   OTG_FS device IN endpoint common interrupt mask register (OTG_FS_DIEPMSK)
//	0x014 32  FS_DOEPMSK   OTG_FS device OUT endpoint common interrupt mask register (OTG_FS_DOEPMSK)
//	0x018 32  FS_DAINT     OTG_FS device all endpoints interrupt register (OTG_FS_DAINT)
//	0x01C 32  FS_DAINTMSK  OTG_FS all endpoints interrupt mask register (OTG_FS_DAINTMSK)
//	0x028 32  DVBUSDIS     OTG_FS device VBUS discharge time register
//	0x02C 32  DVBUSPULSE   OTG_FS device VBUS pulsing time register
//	0x034 32  DIEPEMPMSK   OTG_FS device IN endpoint FIFO empty interrupt mask register
//	0x100 32  FS_DIEPCTL0  OTG_FS device control IN endpoint 0 control register (OTG_FS_DIEPCTL0)
//	0x108 32  DIEPINT0     device endpoint-x interrupt register
//	0x110 32  DIEPTSIZ0    device endpoint-0 transfer size register
//	0x118 32  DTXFSTS0     OTG_FS device IN endpoint transmit FIFO status register
//	0x120 32  DIEPCTL1     OTG device endpoint-1 control register
//	0x128 32  DIEPINT1     device endpoint-1 interrupt register
//	0x130 32  DIEPTSIZ1    device endpoint-1 transfer size register
//	0x138 32  DTXFSTS1     OTG_FS device IN endpoint transmit FIFO status register
//	0x140 32  DIEPCTL2     OTG device endpoint-2 control register
//	0x148 32  DIEPINT2     device endpoint-2 interrupt register
//	0x150 32  DIEPTSIZ2    device endpoint-2 transfer size register
//	0x158 32  DTXFSTS2     OTG_FS device IN endpoint transmit FIFO status register
//	0x160 32  DIEPCTL3     OTG device endpoint-3 control register
//	0x168 32  DIEPINT3     device endpoint-3 interrupt register
//	0x170 32  DIEPTSIZ3    device endpoint-3 transfer size register
//	0x178 32  DTXFSTS3     OTG_FS device IN endpoint transmit FIFO status register
//	0x300 32  DOEPCTL0     device endpoint-0 control register
//	0x308 32  DOEPINT0     device endpoint-0 interrupt register
//	0x310 32  DOEPTSIZ0    device OUT endpoint-0 transfer size register
//	0x320 32  DOEPCTL1     device endpoint-1 control register
//	0x328 32  DOEPINT1     device endpoint-1 interrupt register
//	0x330 32  DOEPTSIZ1    device OUT endpoint-1 transfer size register
//	0x340 32  DOEPCTL2     device endpoint-2 control register
//	0x348 32  DOEPINT2     device endpoint-2 interrupt register
//	0x350 32  DOEPTSIZ2    device OUT endpoint-2 transfer size register
//	0x360 32  DOEPCTL3     device endpoint-3 control register
//	0x368 32  DOEPINT3     device endpoint-3 interrupt register
//	0x370 32  DOEPTSIZ3    device OUT endpoint-3 transfer size register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package otg_fs_device

const (
	DSPD     FS_DCFG = 0x03 << 0  //+ Device speed
	NZLSOHSK FS_DCFG = 0x01 << 2  //+ Non-zero-length status OUT handshake
	DAD      FS_DCFG = 0x7F << 4  //+ Device address
	PFIVL    FS_DCFG = 0x03 << 11 //+ Periodic frame interval
)

const (
	DSPDn     = 0
	NZLSOHSKn = 2
	DADn      = 4
	PFIVLn    = 11
)

const (
	RWUSIG   FS_DCTL = 0x01 << 0  //+ Remote wakeup signaling
	SDIS     FS_DCTL = 0x01 << 1  //+ Soft disconnect
	GINSTS   FS_DCTL = 0x01 << 2  //+ Global IN NAK status
	GONSTS   FS_DCTL = 0x01 << 3  //+ Global OUT NAK status
	TCTL     FS_DCTL = 0x07 << 4  //+ Test control
	SGINAK   FS_DCTL = 0x01 << 7  //+ Set global IN NAK
	CGINAK   FS_DCTL = 0x01 << 8  //+ Clear global IN NAK
	SGONAK   FS_DCTL = 0x01 << 9  //+ Set global OUT NAK
	CGONAK   FS_DCTL = 0x01 << 10 //+ Clear global OUT NAK
	POPRGDNE FS_DCTL = 0x01 << 11 //+ Power-on programming done
)

const (
	RWUSIGn   = 0
	SDISn     = 1
	GINSTSn   = 2
	GONSTSn   = 3
	TCTLn     = 4
	SGINAKn   = 7
	CGINAKn   = 8
	SGONAKn   = 9
	CGONAKn   = 10
	POPRGDNEn = 11
)

const (
	SUSPSTS FS_DSTS = 0x01 << 0   //+ Suspend status
	ENUMSPD FS_DSTS = 0x03 << 1   //+ Enumerated speed
	EERR    FS_DSTS = 0x01 << 3   //+ Erratic error
	FNSOF   FS_DSTS = 0x3FFF << 8 //+ Frame number of the received SOF
)

const (
	SUSPSTSn = 0
	ENUMSPDn = 1
	EERRn    = 3
	FNSOFn   = 8
)

const (
	XFRCM     FS_DIEPMSK = 0x01 << 0 //+ Transfer completed interrupt mask
	EPDM      FS_DIEPMSK = 0x01 << 1 //+ Endpoint disabled interrupt mask
	TOM       FS_DIEPMSK = 0x01 << 3 //+ Timeout condition mask (Non-isochronous endpoints)
	ITTXFEMSK FS_DIEPMSK = 0x01 << 4 //+ IN token received when TxFIFO empty mask
	INEPNMM   FS_DIEPMSK = 0x01 << 5 //+ IN token received with EP mismatch mask
	INEPNEM   FS_DIEPMSK = 0x01 << 6 //+ IN endpoint NAK effective mask
)

const (
	XFRCMn     = 0
	EPDMn      = 1
	TOMn       = 3
	ITTXFEMSKn = 4
	INEPNMMn   = 5
	INEPNEMn   = 6
)

const (
	XFRCM  FS_DOEPMSK = 0x01 << 0 //+ Transfer completed interrupt mask
	EPDM   FS_DOEPMSK = 0x01 << 1 //+ Endpoint disabled interrupt mask
	STUPM  FS_DOEPMSK = 0x01 << 3 //+ SETUP phase done mask
	OTEPDM FS_DOEPMSK = 0x01 << 4 //+ OUT token received when endpoint disabled mask
)

const (
	XFRCMn  = 0
	EPDMn   = 1
	STUPMn  = 3
	OTEPDMn = 4
)

const (
	IEPINT FS_DAINT = 0xFFFF << 0  //+ IN endpoint interrupt bits
	OEPINT FS_DAINT = 0xFFFF << 16 //+ OUT endpoint interrupt bits
)

const (
	IEPINTn = 0
	OEPINTn = 16
)

const (
	IEPM   FS_DAINTMSK = 0xFFFF << 0  //+ IN EP interrupt mask bits
	OEPINT FS_DAINTMSK = 0xFFFF << 16 //+ OUT endpoint interrupt bits
)

const (
	IEPMn   = 0
	OEPINTn = 16
)

const (
	VBUSDT DVBUSDIS = 0xFFFF << 0 //+ Device VBUS discharge time
)

const (
	VBUSDTn = 0
)

const (
	DVBUSP DVBUSPULSE = 0xFFF << 0 //+ Device VBUS pulsing time
)

const (
	DVBUSPn = 0
)

const (
	INEPTXFEM DIEPEMPMSK = 0xFFFF << 0 //+ IN EP Tx FIFO empty interrupt mask bits
)

const (
	INEPTXFEMn = 0
)

const (
	MPSIZ  FS_DIEPCTL0 = 0x03 << 0  //+ Maximum packet size
	USBAEP FS_DIEPCTL0 = 0x01 << 15 //+ USB active endpoint
	NAKSTS FS_DIEPCTL0 = 0x01 << 17 //+ NAK status
	EPTYP  FS_DIEPCTL0 = 0x03 << 18 //+ Endpoint type
	STALL  FS_DIEPCTL0 = 0x01 << 21 //+ STALL handshake
	TXFNUM FS_DIEPCTL0 = 0x0F << 22 //+ TxFIFO number
	CNAK   FS_DIEPCTL0 = 0x01 << 26 //+ Clear NAK
	SNAK   FS_DIEPCTL0 = 0x01 << 27 //+ Set NAK
	EPDIS  FS_DIEPCTL0 = 0x01 << 30 //+ Endpoint disable
	EPENA  FS_DIEPCTL0 = 0x01 << 31 //+ Endpoint enable
)

const (
	MPSIZn  = 0
	USBAEPn = 15
	NAKSTSn = 17
	EPTYPn  = 18
	STALLn  = 21
	TXFNUMn = 22
	CNAKn   = 26
	SNAKn   = 27
	EPDISn  = 30
	EPENAn  = 31
)

const (
	XFRC   DIEPINT0 = 0x01 << 0 //+ XFRC
	EPDISD DIEPINT0 = 0x01 << 1 //+ EPDISD
	TOC    DIEPINT0 = 0x01 << 3 //+ TOC
	ITTXFE DIEPINT0 = 0x01 << 4 //+ ITTXFE
	INEPNE DIEPINT0 = 0x01 << 6 //+ INEPNE
	TXFE   DIEPINT0 = 0x01 << 7 //+ TXFE
)

const (
	XFRCn   = 0
	EPDISDn = 1
	TOCn    = 3
	ITTXFEn = 4
	INEPNEn = 6
	TXFEn   = 7
)

const (
	XFRSIZ DIEPTSIZ0 = 0x7F << 0  //+ Transfer size
	PKTCNT DIEPTSIZ0 = 0x03 << 19 //+ Packet count
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
)

const (
	INEPTFSAV DTXFSTS0 = 0xFFFF << 0 //+ IN endpoint TxFIFO space available
)

const (
	INEPTFSAVn = 0
)

const (
	MPSIZ          DIEPCTL1 = 0x7FF << 0 //+ MPSIZ
	USBAEP         DIEPCTL1 = 0x01 << 15 //+ USBAEP
	EONUM_DPID     DIEPCTL1 = 0x01 << 16 //+ EONUM/DPID
	NAKSTS         DIEPCTL1 = 0x01 << 17 //+ NAKSTS
	EPTYP          DIEPCTL1 = 0x03 << 18 //+ EPTYP
	Stall          DIEPCTL1 = 0x01 << 21 //+ Stall
	TXFNUM         DIEPCTL1 = 0x0F << 22 //+ TXFNUM
	CNAK           DIEPCTL1 = 0x01 << 26 //+ CNAK
	SNAK           DIEPCTL1 = 0x01 << 27 //+ SNAK
	SD0PID_SEVNFRM DIEPCTL1 = 0x01 << 28 //+ SD0PID/SEVNFRM
	SODDFRM_SD1PID DIEPCTL1 = 0x01 << 29 //+ SODDFRM/SD1PID
	EPDIS          DIEPCTL1 = 0x01 << 30 //+ EPDIS
	EPENA          DIEPCTL1 = 0x01 << 31 //+ EPENA
)

const (
	MPSIZn          = 0
	USBAEPn         = 15
	EONUM_DPIDn     = 16
	NAKSTSn         = 17
	EPTYPn          = 18
	Stalln          = 21
	TXFNUMn         = 22
	CNAKn           = 26
	SNAKn           = 27
	SD0PID_SEVNFRMn = 28
	SODDFRM_SD1PIDn = 29
	EPDISn          = 30
	EPENAn          = 31
)

const (
	XFRC   DIEPINT1 = 0x01 << 0 //+ XFRC
	EPDISD DIEPINT1 = 0x01 << 1 //+ EPDISD
	TOC    DIEPINT1 = 0x01 << 3 //+ TOC
	ITTXFE DIEPINT1 = 0x01 << 4 //+ ITTXFE
	INEPNE DIEPINT1 = 0x01 << 6 //+ INEPNE
	TXFE   DIEPINT1 = 0x01 << 7 //+ TXFE
)

const (
	XFRCn   = 0
	EPDISDn = 1
	TOCn    = 3
	ITTXFEn = 4
	INEPNEn = 6
	TXFEn   = 7
)

const (
	XFRSIZ DIEPTSIZ1 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT DIEPTSIZ1 = 0x3FF << 19  //+ Packet count
	MCNT   DIEPTSIZ1 = 0x03 << 29   //+ Multi count
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	MCNTn   = 29
)

const (
	INEPTFSAV DTXFSTS1 = 0xFFFF << 0 //+ IN endpoint TxFIFO space available
)

const (
	INEPTFSAVn = 0
)

const (
	MPSIZ          DIEPCTL2 = 0x7FF << 0 //+ MPSIZ
	USBAEP         DIEPCTL2 = 0x01 << 15 //+ USBAEP
	EONUM_DPID     DIEPCTL2 = 0x01 << 16 //+ EONUM/DPID
	NAKSTS         DIEPCTL2 = 0x01 << 17 //+ NAKSTS
	EPTYP          DIEPCTL2 = 0x03 << 18 //+ EPTYP
	Stall          DIEPCTL2 = 0x01 << 21 //+ Stall
	TXFNUM         DIEPCTL2 = 0x0F << 22 //+ TXFNUM
	CNAK           DIEPCTL2 = 0x01 << 26 //+ CNAK
	SNAK           DIEPCTL2 = 0x01 << 27 //+ SNAK
	SD0PID_SEVNFRM DIEPCTL2 = 0x01 << 28 //+ SD0PID/SEVNFRM
	SODDFRM        DIEPCTL2 = 0x01 << 29 //+ SODDFRM
	EPDIS          DIEPCTL2 = 0x01 << 30 //+ EPDIS
	EPENA          DIEPCTL2 = 0x01 << 31 //+ EPENA
)

const (
	MPSIZn          = 0
	USBAEPn         = 15
	EONUM_DPIDn     = 16
	NAKSTSn         = 17
	EPTYPn          = 18
	Stalln          = 21
	TXFNUMn         = 22
	CNAKn           = 26
	SNAKn           = 27
	SD0PID_SEVNFRMn = 28
	SODDFRMn        = 29
	EPDISn          = 30
	EPENAn          = 31
)

const (
	XFRC   DIEPINT2 = 0x01 << 0 //+ XFRC
	EPDISD DIEPINT2 = 0x01 << 1 //+ EPDISD
	TOC    DIEPINT2 = 0x01 << 3 //+ TOC
	ITTXFE DIEPINT2 = 0x01 << 4 //+ ITTXFE
	INEPNE DIEPINT2 = 0x01 << 6 //+ INEPNE
	TXFE   DIEPINT2 = 0x01 << 7 //+ TXFE
)

const (
	XFRCn   = 0
	EPDISDn = 1
	TOCn    = 3
	ITTXFEn = 4
	INEPNEn = 6
	TXFEn   = 7
)

const (
	XFRSIZ DIEPTSIZ2 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT DIEPTSIZ2 = 0x3FF << 19  //+ Packet count
	MCNT   DIEPTSIZ2 = 0x03 << 29   //+ Multi count
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	MCNTn   = 29
)

const (
	INEPTFSAV DTXFSTS2 = 0xFFFF << 0 //+ IN endpoint TxFIFO space available
)

const (
	INEPTFSAVn = 0
)

const (
	MPSIZ          DIEPCTL3 = 0x7FF << 0 //+ MPSIZ
	USBAEP         DIEPCTL3 = 0x01 << 15 //+ USBAEP
	EONUM_DPID     DIEPCTL3 = 0x01 << 16 //+ EONUM/DPID
	NAKSTS         DIEPCTL3 = 0x01 << 17 //+ NAKSTS
	EPTYP          DIEPCTL3 = 0x03 << 18 //+ EPTYP
	Stall          DIEPCTL3 = 0x01 << 21 //+ Stall
	TXFNUM         DIEPCTL3 = 0x0F << 22 //+ TXFNUM
	CNAK           DIEPCTL3 = 0x01 << 26 //+ CNAK
	SNAK           DIEPCTL3 = 0x01 << 27 //+ SNAK
	SD0PID_SEVNFRM DIEPCTL3 = 0x01 << 28 //+ SD0PID/SEVNFRM
	SODDFRM        DIEPCTL3 = 0x01 << 29 //+ SODDFRM
	EPDIS          DIEPCTL3 = 0x01 << 30 //+ EPDIS
	EPENA          DIEPCTL3 = 0x01 << 31 //+ EPENA
)

const (
	MPSIZn          = 0
	USBAEPn         = 15
	EONUM_DPIDn     = 16
	NAKSTSn         = 17
	EPTYPn          = 18
	Stalln          = 21
	TXFNUMn         = 22
	CNAKn           = 26
	SNAKn           = 27
	SD0PID_SEVNFRMn = 28
	SODDFRMn        = 29
	EPDISn          = 30
	EPENAn          = 31
)

const (
	XFRC   DIEPINT3 = 0x01 << 0 //+ XFRC
	EPDISD DIEPINT3 = 0x01 << 1 //+ EPDISD
	TOC    DIEPINT3 = 0x01 << 3 //+ TOC
	ITTXFE DIEPINT3 = 0x01 << 4 //+ ITTXFE
	INEPNE DIEPINT3 = 0x01 << 6 //+ INEPNE
	TXFE   DIEPINT3 = 0x01 << 7 //+ TXFE
)

const (
	XFRCn   = 0
	EPDISDn = 1
	TOCn    = 3
	ITTXFEn = 4
	INEPNEn = 6
	TXFEn   = 7
)

const (
	XFRSIZ DIEPTSIZ3 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT DIEPTSIZ3 = 0x3FF << 19  //+ Packet count
	MCNT   DIEPTSIZ3 = 0x03 << 29   //+ Multi count
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	MCNTn   = 29
)

const (
	INEPTFSAV DTXFSTS3 = 0xFFFF << 0 //+ IN endpoint TxFIFO space available
)

const (
	INEPTFSAVn = 0
)

const (
	MPSIZ  DOEPCTL0 = 0x03 << 0  //+ MPSIZ
	USBAEP DOEPCTL0 = 0x01 << 15 //+ USBAEP
	NAKSTS DOEPCTL0 = 0x01 << 17 //+ NAKSTS
	EPTYP  DOEPCTL0 = 0x03 << 18 //+ EPTYP
	SNPM   DOEPCTL0 = 0x01 << 20 //+ SNPM
	Stall  DOEPCTL0 = 0x01 << 21 //+ Stall
	CNAK   DOEPCTL0 = 0x01 << 26 //+ CNAK
	SNAK   DOEPCTL0 = 0x01 << 27 //+ SNAK
	EPDIS  DOEPCTL0 = 0x01 << 30 //+ EPDIS
	EPENA  DOEPCTL0 = 0x01 << 31 //+ EPENA
)

const (
	MPSIZn  = 0
	USBAEPn = 15
	NAKSTSn = 17
	EPTYPn  = 18
	SNPMn   = 20
	Stalln  = 21
	CNAKn   = 26
	SNAKn   = 27
	EPDISn  = 30
	EPENAn  = 31
)

const (
	XFRC    DOEPINT0 = 0x01 << 0 //+ XFRC
	EPDISD  DOEPINT0 = 0x01 << 1 //+ EPDISD
	STUP    DOEPINT0 = 0x01 << 3 //+ STUP
	OTEPDIS DOEPINT0 = 0x01 << 4 //+ OTEPDIS
	B2BSTUP DOEPINT0 = 0x01 << 6 //+ B2BSTUP
)

const (
	XFRCn    = 0
	EPDISDn  = 1
	STUPn    = 3
	OTEPDISn = 4
	B2BSTUPn = 6
)

const (
	XFRSIZ  DOEPTSIZ0 = 0x7F << 0  //+ Transfer size
	PKTCNT  DOEPTSIZ0 = 0x01 << 19 //+ Packet count
	STUPCNT DOEPTSIZ0 = 0x03 << 29 //+ SETUP packet count
)

const (
	XFRSIZn  = 0
	PKTCNTn  = 19
	STUPCNTn = 29
)

const (
	MPSIZ          DOEPCTL1 = 0x7FF << 0 //+ MPSIZ
	USBAEP         DOEPCTL1 = 0x01 << 15 //+ USBAEP
	EONUM_DPID     DOEPCTL1 = 0x01 << 16 //+ EONUM/DPID
	NAKSTS         DOEPCTL1 = 0x01 << 17 //+ NAKSTS
	EPTYP          DOEPCTL1 = 0x03 << 18 //+ EPTYP
	SNPM           DOEPCTL1 = 0x01 << 20 //+ SNPM
	Stall          DOEPCTL1 = 0x01 << 21 //+ Stall
	CNAK           DOEPCTL1 = 0x01 << 26 //+ CNAK
	SNAK           DOEPCTL1 = 0x01 << 27 //+ SNAK
	SD0PID_SEVNFRM DOEPCTL1 = 0x01 << 28 //+ SD0PID/SEVNFRM
	SODDFRM        DOEPCTL1 = 0x01 << 29 //+ SODDFRM
	EPDIS          DOEPCTL1 = 0x01 << 30 //+ EPDIS
	EPENA          DOEPCTL1 = 0x01 << 31 //+ EPENA
)

const (
	MPSIZn          = 0
	USBAEPn         = 15
	EONUM_DPIDn     = 16
	NAKSTSn         = 17
	EPTYPn          = 18
	SNPMn           = 20
	Stalln          = 21
	CNAKn           = 26
	SNAKn           = 27
	SD0PID_SEVNFRMn = 28
	SODDFRMn        = 29
	EPDISn          = 30
	EPENAn          = 31
)

const (
	XFRC    DOEPINT1 = 0x01 << 0 //+ XFRC
	EPDISD  DOEPINT1 = 0x01 << 1 //+ EPDISD
	STUP    DOEPINT1 = 0x01 << 3 //+ STUP
	OTEPDIS DOEPINT1 = 0x01 << 4 //+ OTEPDIS
	B2BSTUP DOEPINT1 = 0x01 << 6 //+ B2BSTUP
)

const (
	XFRCn    = 0
	EPDISDn  = 1
	STUPn    = 3
	OTEPDISn = 4
	B2BSTUPn = 6
)

const (
	XFRSIZ         DOEPTSIZ1 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT         DOEPTSIZ1 = 0x3FF << 19  //+ Packet count
	RXDPID_STUPCNT DOEPTSIZ1 = 0x03 << 29   //+ Received data PID/SETUP packet count
)

const (
	XFRSIZn         = 0
	PKTCNTn         = 19
	RXDPID_STUPCNTn = 29
)

const (
	MPSIZ          DOEPCTL2 = 0x7FF << 0 //+ MPSIZ
	USBAEP         DOEPCTL2 = 0x01 << 15 //+ USBAEP
	EONUM_DPID     DOEPCTL2 = 0x01 << 16 //+ EONUM/DPID
	NAKSTS         DOEPCTL2 = 0x01 << 17 //+ NAKSTS
	EPTYP          DOEPCTL2 = 0x03 << 18 //+ EPTYP
	SNPM           DOEPCTL2 = 0x01 << 20 //+ SNPM
	Stall          DOEPCTL2 = 0x01 << 21 //+ Stall
	CNAK           DOEPCTL2 = 0x01 << 26 //+ CNAK
	SNAK           DOEPCTL2 = 0x01 << 27 //+ SNAK
	SD0PID_SEVNFRM DOEPCTL2 = 0x01 << 28 //+ SD0PID/SEVNFRM
	SODDFRM        DOEPCTL2 = 0x01 << 29 //+ SODDFRM
	EPDIS          DOEPCTL2 = 0x01 << 30 //+ EPDIS
	EPENA          DOEPCTL2 = 0x01 << 31 //+ EPENA
)

const (
	MPSIZn          = 0
	USBAEPn         = 15
	EONUM_DPIDn     = 16
	NAKSTSn         = 17
	EPTYPn          = 18
	SNPMn           = 20
	Stalln          = 21
	CNAKn           = 26
	SNAKn           = 27
	SD0PID_SEVNFRMn = 28
	SODDFRMn        = 29
	EPDISn          = 30
	EPENAn          = 31
)

const (
	XFRC    DOEPINT2 = 0x01 << 0 //+ XFRC
	EPDISD  DOEPINT2 = 0x01 << 1 //+ EPDISD
	STUP    DOEPINT2 = 0x01 << 3 //+ STUP
	OTEPDIS DOEPINT2 = 0x01 << 4 //+ OTEPDIS
	B2BSTUP DOEPINT2 = 0x01 << 6 //+ B2BSTUP
)

const (
	XFRCn    = 0
	EPDISDn  = 1
	STUPn    = 3
	OTEPDISn = 4
	B2BSTUPn = 6
)

const (
	XFRSIZ         DOEPTSIZ2 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT         DOEPTSIZ2 = 0x3FF << 19  //+ Packet count
	RXDPID_STUPCNT DOEPTSIZ2 = 0x03 << 29   //+ Received data PID/SETUP packet count
)

const (
	XFRSIZn         = 0
	PKTCNTn         = 19
	RXDPID_STUPCNTn = 29
)

const (
	MPSIZ          DOEPCTL3 = 0x7FF << 0 //+ MPSIZ
	USBAEP         DOEPCTL3 = 0x01 << 15 //+ USBAEP
	EONUM_DPID     DOEPCTL3 = 0x01 << 16 //+ EONUM/DPID
	NAKSTS         DOEPCTL3 = 0x01 << 17 //+ NAKSTS
	EPTYP          DOEPCTL3 = 0x03 << 18 //+ EPTYP
	SNPM           DOEPCTL3 = 0x01 << 20 //+ SNPM
	Stall          DOEPCTL3 = 0x01 << 21 //+ Stall
	CNAK           DOEPCTL3 = 0x01 << 26 //+ CNAK
	SNAK           DOEPCTL3 = 0x01 << 27 //+ SNAK
	SD0PID_SEVNFRM DOEPCTL3 = 0x01 << 28 //+ SD0PID/SEVNFRM
	SODDFRM        DOEPCTL3 = 0x01 << 29 //+ SODDFRM
	EPDIS          DOEPCTL3 = 0x01 << 30 //+ EPDIS
	EPENA          DOEPCTL3 = 0x01 << 31 //+ EPENA
)

const (
	MPSIZn          = 0
	USBAEPn         = 15
	EONUM_DPIDn     = 16
	NAKSTSn         = 17
	EPTYPn          = 18
	SNPMn           = 20
	Stalln          = 21
	CNAKn           = 26
	SNAKn           = 27
	SD0PID_SEVNFRMn = 28
	SODDFRMn        = 29
	EPDISn          = 30
	EPENAn          = 31
)

const (
	XFRC    DOEPINT3 = 0x01 << 0 //+ XFRC
	EPDISD  DOEPINT3 = 0x01 << 1 //+ EPDISD
	STUP    DOEPINT3 = 0x01 << 3 //+ STUP
	OTEPDIS DOEPINT3 = 0x01 << 4 //+ OTEPDIS
	B2BSTUP DOEPINT3 = 0x01 << 6 //+ B2BSTUP
)

const (
	XFRCn    = 0
	EPDISDn  = 1
	STUPn    = 3
	OTEPDISn = 4
	B2BSTUPn = 6
)

const (
	XFRSIZ         DOEPTSIZ3 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT         DOEPTSIZ3 = 0x3FF << 19  //+ Packet count
	RXDPID_STUPCNT DOEPTSIZ3 = 0x03 << 29   //+ Received data PID/SETUP packet count
)

const (
	XFRSIZn         = 0
	PKTCNTn         = 19
	RXDPID_STUPCNTn = 29
)
