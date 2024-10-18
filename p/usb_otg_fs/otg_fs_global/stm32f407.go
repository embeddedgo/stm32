// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f407

// Package otg_fs_global provides access to the registers of the OTG_FS_GLOBAL peripheral.
//
// Instances:
//
//	OTG_FS_GLOBAL  OTG_FS_GLOBAL_BASE  -  OTG_FS_WKUP,OTG_FS  USB on the go full speed
//
// Registers:
//
//	0x000 32  FS_GOTGCTL           OTG_FS control and status register (OTG_FS_GOTGCTL)
//	0x004 32  FS_GOTGINT           OTG_FS interrupt register (OTG_FS_GOTGINT)
//	0x008 32  FS_GAHBCFG           OTG_FS AHB configuration register (OTG_FS_GAHBCFG)
//	0x00C 32  FS_GUSBCFG           OTG_FS USB configuration register (OTG_FS_GUSBCFG)
//	0x010 32  FS_GRSTCTL           OTG_FS reset register (OTG_FS_GRSTCTL)
//	0x014 32  FS_GINTSTS           OTG_FS core interrupt register (OTG_FS_GINTSTS)
//	0x018 32  FS_GINTMSK           OTG_FS interrupt mask register (OTG_FS_GINTMSK)
//	0x01C 32  FS_GRXSTSR_Device    OTG_FS Receive status debug read(Device mode)
//	0x01C 32  FS_GRXSTSR_Host      OTG_FS Receive status debug read(Host mode)
//	0x024 32  FS_GRXFSIZ           OTG_FS Receive FIFO size register (OTG_FS_GRXFSIZ)
//	0x028 32  FS_GNPTXFSIZ_Device  OTG_FS non-periodic transmit FIFO size register (Device mode)
//	0x028 32  FS_GNPTXFSIZ_Host    OTG_FS non-periodic transmit FIFO size register (Host mode)
//	0x02C 32  FS_GNPTXSTS          OTG_FS non-periodic transmit FIFO/queue status register (OTG_FS_GNPTXSTS)
//	0x038 32  FS_GCCFG             OTG_FS general core configuration register (OTG_FS_GCCFG)
//	0x03C 32  FS_CID               core ID register
//	0x100 32  FS_HPTXFSIZ          OTG_FS Host periodic transmit FIFO size register (OTG_FS_HPTXFSIZ)
//	0x104 32  FS_DIEPTXF1          OTG_FS device IN endpoint transmit FIFO size register (OTG_FS_DIEPTXF2)
//	0x108 32  FS_DIEPTXF2          OTG_FS device IN endpoint transmit FIFO size register (OTG_FS_DIEPTXF3)
//	0x10C 32  FS_DIEPTXF3          OTG_FS device IN endpoint transmit FIFO size register (OTG_FS_DIEPTXF4)
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package otg_fs_global

const (
	SRQSCS  FS_GOTGCTL = 0x01 << 0  //+ Session request success
	SRQ     FS_GOTGCTL = 0x01 << 1  //+ Session request
	HNGSCS  FS_GOTGCTL = 0x01 << 8  //+ Host negotiation success
	HNPRQ   FS_GOTGCTL = 0x01 << 9  //+ HNP request
	HSHNPEN FS_GOTGCTL = 0x01 << 10 //+ Host set HNP enable
	DHNPEN  FS_GOTGCTL = 0x01 << 11 //+ Device HNP enabled
	CIDSTS  FS_GOTGCTL = 0x01 << 16 //+ Connector ID status
	DBCT    FS_GOTGCTL = 0x01 << 17 //+ Long/short debounce time
	ASVLD   FS_GOTGCTL = 0x01 << 18 //+ A-session valid
	BSVLD   FS_GOTGCTL = 0x01 << 19 //+ B-session valid
)

const (
	SRQSCSn  = 0
	SRQn     = 1
	HNGSCSn  = 8
	HNPRQn   = 9
	HSHNPENn = 10
	DHNPENn  = 11
	CIDSTSn  = 16
	DBCTn    = 17
	ASVLDn   = 18
	BSVLDn   = 19
)

const (
	SEDET   FS_GOTGINT = 0x01 << 2  //+ Session end detected
	SRSSCHG FS_GOTGINT = 0x01 << 8  //+ Session request success status change
	HNSSCHG FS_GOTGINT = 0x01 << 9  //+ Host negotiation success status change
	HNGDET  FS_GOTGINT = 0x01 << 17 //+ Host negotiation detected
	ADTOCHG FS_GOTGINT = 0x01 << 18 //+ A-device timeout change
	DBCDNE  FS_GOTGINT = 0x01 << 19 //+ Debounce done
)

const (
	SEDETn   = 2
	SRSSCHGn = 8
	HNSSCHGn = 9
	HNGDETn  = 17
	ADTOCHGn = 18
	DBCDNEn  = 19
)

const (
	GINT     FS_GAHBCFG = 0x01 << 0 //+ Global interrupt mask
	TXFELVL  FS_GAHBCFG = 0x01 << 7 //+ TxFIFO empty level
	PTXFELVL FS_GAHBCFG = 0x01 << 8 //+ Periodic TxFIFO empty level
)

const (
	GINTn     = 0
	TXFELVLn  = 7
	PTXFELVLn = 8
)

const (
	TOCAL  FS_GUSBCFG = 0x07 << 0  //+ FS timeout calibration
	PHYSEL FS_GUSBCFG = 0x01 << 6  //+ Full Speed serial transceiver select
	SRPCAP FS_GUSBCFG = 0x01 << 8  //+ SRP-capable
	HNPCAP FS_GUSBCFG = 0x01 << 9  //+ HNP-capable
	TRDT   FS_GUSBCFG = 0x0F << 10 //+ USB turnaround time
	FHMOD  FS_GUSBCFG = 0x01 << 29 //+ Force host mode
	FDMOD  FS_GUSBCFG = 0x01 << 30 //+ Force device mode
	CTXPKT FS_GUSBCFG = 0x01 << 31 //+ Corrupt Tx packet
)

const (
	TOCALn  = 0
	PHYSELn = 6
	SRPCAPn = 8
	HNPCAPn = 9
	TRDTn   = 10
	FHMODn  = 29
	FDMODn  = 30
	CTXPKTn = 31
)

const (
	CSRST   FS_GRSTCTL = 0x01 << 0  //+ Core soft reset
	HSRST   FS_GRSTCTL = 0x01 << 1  //+ HCLK soft reset
	FCRST   FS_GRSTCTL = 0x01 << 2  //+ Host frame counter reset
	RXFFLSH FS_GRSTCTL = 0x01 << 4  //+ RxFIFO flush
	TXFFLSH FS_GRSTCTL = 0x01 << 5  //+ TxFIFO flush
	TXFNUM  FS_GRSTCTL = 0x1F << 6  //+ TxFIFO number
	AHBIDL  FS_GRSTCTL = 0x01 << 31 //+ AHB master idle
)

const (
	CSRSTn   = 0
	HSRSTn   = 1
	FCRSTn   = 2
	RXFFLSHn = 4
	TXFFLSHn = 5
	TXFNUMn  = 6
	AHBIDLn  = 31
)

const (
	CMOD               FS_GINTSTS = 0x01 << 0  //+ Current mode of operation
	MMIS               FS_GINTSTS = 0x01 << 1  //+ Mode mismatch interrupt
	OTGINT             FS_GINTSTS = 0x01 << 2  //+ OTG interrupt
	SOF                FS_GINTSTS = 0x01 << 3  //+ Start of frame
	RXFLVL             FS_GINTSTS = 0x01 << 4  //+ RxFIFO non-empty
	NPTXFE             FS_GINTSTS = 0x01 << 5  //+ Non-periodic TxFIFO empty
	GINAKEFF           FS_GINTSTS = 0x01 << 6  //+ Global IN non-periodic NAK effective
	GOUTNAKEFF         FS_GINTSTS = 0x01 << 7  //+ Global OUT NAK effective
	ESUSP              FS_GINTSTS = 0x01 << 10 //+ Early suspend
	USBSUSP            FS_GINTSTS = 0x01 << 11 //+ USB suspend
	USBRST             FS_GINTSTS = 0x01 << 12 //+ USB reset
	ENUMDNE            FS_GINTSTS = 0x01 << 13 //+ Enumeration done
	ISOODRP            FS_GINTSTS = 0x01 << 14 //+ Isochronous OUT packet dropped interrupt
	EOPF               FS_GINTSTS = 0x01 << 15 //+ End of periodic frame interrupt
	IEPINT             FS_GINTSTS = 0x01 << 18 //+ IN endpoint interrupt
	OEPINT             FS_GINTSTS = 0x01 << 19 //+ OUT endpoint interrupt
	IISOIXFR           FS_GINTSTS = 0x01 << 20 //+ Incomplete isochronous IN transfer
	IPXFR_INCOMPISOOUT FS_GINTSTS = 0x01 << 21 //+ Incomplete periodic transfer(Host mode)/Incomplete isochronous OUT transfer(Device mode)
	HPRTINT            FS_GINTSTS = 0x01 << 24 //+ Host port interrupt
	HCINT              FS_GINTSTS = 0x01 << 25 //+ Host channels interrupt
	PTXFE              FS_GINTSTS = 0x01 << 26 //+ Periodic TxFIFO empty
	CIDSCHG            FS_GINTSTS = 0x01 << 28 //+ Connector ID status change
	DISCINT            FS_GINTSTS = 0x01 << 29 //+ Disconnect detected interrupt
	SRQINT             FS_GINTSTS = 0x01 << 30 //+ Session request/new session detected interrupt
	WKUPINT            FS_GINTSTS = 0x01 << 31 //+ Resume/remote wakeup detected interrupt
)

const (
	CMODn               = 0
	MMISn               = 1
	OTGINTn             = 2
	SOFn                = 3
	RXFLVLn             = 4
	NPTXFEn             = 5
	GINAKEFFn           = 6
	GOUTNAKEFFn         = 7
	ESUSPn              = 10
	USBSUSPn            = 11
	USBRSTn             = 12
	ENUMDNEn            = 13
	ISOODRPn            = 14
	EOPFn               = 15
	IEPINTn             = 18
	OEPINTn             = 19
	IISOIXFRn           = 20
	IPXFR_INCOMPISOOUTn = 21
	HPRTINTn            = 24
	HCINTn              = 25
	PTXFEn              = 26
	CIDSCHGn            = 28
	DISCINTn            = 29
	SRQINTn             = 30
	WKUPINTn            = 31
)

const (
	MMISM            FS_GINTMSK = 0x01 << 1  //+ Mode mismatch interrupt mask
	OTGINT           FS_GINTMSK = 0x01 << 2  //+ OTG interrupt mask
	SOFM             FS_GINTMSK = 0x01 << 3  //+ Start of frame mask
	RXFLVLM          FS_GINTMSK = 0x01 << 4  //+ Receive FIFO non-empty mask
	NPTXFEM          FS_GINTMSK = 0x01 << 5  //+ Non-periodic TxFIFO empty mask
	GINAKEFFM        FS_GINTMSK = 0x01 << 6  //+ Global non-periodic IN NAK effective mask
	GONAKEFFM        FS_GINTMSK = 0x01 << 7  //+ Global OUT NAK effective mask
	ESUSPM           FS_GINTMSK = 0x01 << 10 //+ Early suspend mask
	USBSUSPM         FS_GINTMSK = 0x01 << 11 //+ USB suspend mask
	USBRST           FS_GINTMSK = 0x01 << 12 //+ USB reset mask
	ENUMDNEM         FS_GINTMSK = 0x01 << 13 //+ Enumeration done mask
	ISOODRPM         FS_GINTMSK = 0x01 << 14 //+ Isochronous OUT packet dropped interrupt mask
	EOPFM            FS_GINTMSK = 0x01 << 15 //+ End of periodic frame interrupt mask
	EPMISM           FS_GINTMSK = 0x01 << 17 //+ Endpoint mismatch interrupt mask
	IEPINT           FS_GINTMSK = 0x01 << 18 //+ IN endpoints interrupt mask
	OEPINT           FS_GINTMSK = 0x01 << 19 //+ OUT endpoints interrupt mask
	IISOIXFRM        FS_GINTMSK = 0x01 << 20 //+ Incomplete isochronous IN transfer mask
	IPXFRM_IISOOXFRM FS_GINTMSK = 0x01 << 21 //+ Incomplete periodic transfer mask(Host mode)/Incomplete isochronous OUT transfer mask(Device mode)
	PRTIM            FS_GINTMSK = 0x01 << 24 //+ Host port interrupt mask
	HCIM             FS_GINTMSK = 0x01 << 25 //+ Host channels interrupt mask
	PTXFEM           FS_GINTMSK = 0x01 << 26 //+ Periodic TxFIFO empty mask
	CIDSCHGM         FS_GINTMSK = 0x01 << 28 //+ Connector ID status change mask
	DISCINT          FS_GINTMSK = 0x01 << 29 //+ Disconnect detected interrupt mask
	SRQIM            FS_GINTMSK = 0x01 << 30 //+ Session request/new session detected interrupt mask
	WUIM             FS_GINTMSK = 0x01 << 31 //+ Resume/remote wakeup detected interrupt mask
)

const (
	MMISMn            = 1
	OTGINTn           = 2
	SOFMn             = 3
	RXFLVLMn          = 4
	NPTXFEMn          = 5
	GINAKEFFMn        = 6
	GONAKEFFMn        = 7
	ESUSPMn           = 10
	USBSUSPMn         = 11
	USBRSTn           = 12
	ENUMDNEMn         = 13
	ISOODRPMn         = 14
	EOPFMn            = 15
	EPMISMn           = 17
	IEPINTn           = 18
	OEPINTn           = 19
	IISOIXFRMn        = 20
	IPXFRM_IISOOXFRMn = 21
	PRTIMn            = 24
	HCIMn             = 25
	PTXFEMn           = 26
	CIDSCHGMn         = 28
	DISCINTn          = 29
	SRQIMn            = 30
	WUIMn             = 31
)

const (
	EPNUM  FS_GRXSTSR_Device = 0x0F << 0  //+ Endpoint number
	BCNT   FS_GRXSTSR_Device = 0x7FF << 4 //+ Byte count
	DPID   FS_GRXSTSR_Device = 0x03 << 15 //+ Data PID
	PKTSTS FS_GRXSTSR_Device = 0x0F << 17 //+ Packet status
	FRMNUM FS_GRXSTSR_Device = 0x0F << 21 //+ Frame number
)

const (
	EPNUMn  = 0
	BCNTn   = 4
	DPIDn   = 15
	PKTSTSn = 17
	FRMNUMn = 21
)

const (
	EPNUM  FS_GRXSTSR_Host = 0x0F << 0  //+ Endpoint number
	BCNT   FS_GRXSTSR_Host = 0x7FF << 4 //+ Byte count
	DPID   FS_GRXSTSR_Host = 0x03 << 15 //+ Data PID
	PKTSTS FS_GRXSTSR_Host = 0x0F << 17 //+ Packet status
	FRMNUM FS_GRXSTSR_Host = 0x0F << 21 //+ Frame number
)

const (
	EPNUMn  = 0
	BCNTn   = 4
	DPIDn   = 15
	PKTSTSn = 17
	FRMNUMn = 21
)

const (
	RXFD FS_GRXFSIZ = 0xFFFF << 0 //+ RxFIFO depth
)

const (
	RXFDn = 0
)

const (
	TX0FSA FS_GNPTXFSIZ_Device = 0xFFFF << 0  //+ Endpoint 0 transmit RAM start address
	TX0FD  FS_GNPTXFSIZ_Device = 0xFFFF << 16 //+ Endpoint 0 TxFIFO depth
)

const (
	TX0FSAn = 0
	TX0FDn  = 16
)

const (
	NPTXFSA FS_GNPTXFSIZ_Host = 0xFFFF << 0  //+ Non-periodic transmit RAM start address
	NPTXFD  FS_GNPTXFSIZ_Host = 0xFFFF << 16 //+ Non-periodic TxFIFO depth
)

const (
	NPTXFSAn = 0
	NPTXFDn  = 16
)

const (
	NPTXFSAV FS_GNPTXSTS = 0xFFFF << 0 //+ Non-periodic TxFIFO space available
	NPTQXSAV FS_GNPTXSTS = 0xFF << 16  //+ Non-periodic transmit request queue space available
	NPTXQTOP FS_GNPTXSTS = 0x7F << 24  //+ Top of the non-periodic transmit request queue
)

const (
	NPTXFSAVn = 0
	NPTQXSAVn = 16
	NPTXQTOPn = 24
)

const (
	PWRDWN   FS_GCCFG = 0x01 << 16 //+ Power down
	VBUSASEN FS_GCCFG = 0x01 << 18 //+ Enable the VBUS sensing device
	VBUSBSEN FS_GCCFG = 0x01 << 19 //+ Enable the VBUS sensing device
	SOFOUTEN FS_GCCFG = 0x01 << 20 //+ SOF output enable
)

const (
	PWRDWNn   = 16
	VBUSASENn = 18
	VBUSBSENn = 19
	SOFOUTENn = 20
)

const (
	PRODUCT_ID FS_CID = 0xFFFFFFFF << 0 //+ Product ID field
)

const (
	PRODUCT_IDn = 0
)

const (
	PTXSA   FS_HPTXFSIZ = 0xFFFF << 0  //+ Host periodic TxFIFO start address
	PTXFSIZ FS_HPTXFSIZ = 0xFFFF << 16 //+ Host periodic TxFIFO depth
)

const (
	PTXSAn   = 0
	PTXFSIZn = 16
)

const (
	INEPTXSA FS_DIEPTXF1 = 0xFFFF << 0  //+ IN endpoint FIFO2 transmit RAM start address
	INEPTXFD FS_DIEPTXF1 = 0xFFFF << 16 //+ IN endpoint TxFIFO depth
)

const (
	INEPTXSAn = 0
	INEPTXFDn = 16
)

const (
	INEPTXSA FS_DIEPTXF2 = 0xFFFF << 0  //+ IN endpoint FIFO3 transmit RAM start address
	INEPTXFD FS_DIEPTXF2 = 0xFFFF << 16 //+ IN endpoint TxFIFO depth
)

const (
	INEPTXSAn = 0
	INEPTXFDn = 16
)

const (
	INEPTXSA FS_DIEPTXF3 = 0xFFFF << 0  //+ IN endpoint FIFO4 transmit RAM start address
	INEPTXFD FS_DIEPTXF3 = 0xFFFF << 16 //+ IN endpoint TxFIFO depth
)

const (
	INEPTXSAn = 0
	INEPTXFDn = 16
)
