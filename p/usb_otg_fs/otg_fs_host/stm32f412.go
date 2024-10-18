// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32f412

// Package otg_fs_host provides access to the registers of the OTG_FS_HOST peripheral.
//
// Instances:
//
//	OTG_FS_HOST  OTG_FS_HOST_BASE  -  -  USB on the go full speed
//
// Registers:
//
//	0x000 32  FS_HCFG       OTG_FS host configuration register (OTG_FS_HCFG)
//	0x004 32  HFIR          OTG_FS Host frame interval register
//	0x008 32  FS_HFNUM      OTG_FS host frame number/frame time remaining register (OTG_FS_HFNUM)
//	0x010 32  FS_HPTXSTS    OTG_FS_Host periodic transmit FIFO/queue status register (OTG_FS_HPTXSTS)
//	0x014 32  HAINT         OTG_FS Host all channels interrupt register
//	0x018 32  HAINTMSK      OTG_FS host all channels interrupt mask register
//	0x040 32  FS_HPRT       OTG_FS host port control and status register (OTG_FS_HPRT)
//	0x100 32  FS_HCCHAR0    OTG_FS host channel-0 characteristics register (OTG_FS_HCCHAR0)
//	0x108 32  FS_HCINT0     OTG_FS host channel-0 interrupt register (OTG_FS_HCINT0)
//	0x10C 32  FS_HCINTMSK0  OTG_FS host channel-0 mask register (OTG_FS_HCINTMSK0)
//	0x110 32  FS_HCTSIZ0    OTG_FS host channel-0 transfer size register
//	0x120 32  FS_HCCHAR1    OTG_FS host channel-1 characteristics register (OTG_FS_HCCHAR1)
//	0x128 32  FS_HCINT1     OTG_FS host channel-1 interrupt register (OTG_FS_HCINT1)
//	0x12C 32  FS_HCINTMSK1  OTG_FS host channel-1 mask register (OTG_FS_HCINTMSK1)
//	0x130 32  FS_HCTSIZ1    OTG_FS host channel-1 transfer size register
//	0x140 32  FS_HCCHAR2    OTG_FS host channel-2 characteristics register (OTG_FS_HCCHAR2)
//	0x148 32  FS_HCINT2     OTG_FS host channel-2 interrupt register (OTG_FS_HCINT2)
//	0x14C 32  FS_HCINTMSK2  OTG_FS host channel-2 mask register (OTG_FS_HCINTMSK2)
//	0x150 32  FS_HCTSIZ2    OTG_FS host channel-2 transfer size register
//	0x160 32  FS_HCCHAR3    OTG_FS host channel-3 characteristics register (OTG_FS_HCCHAR3)
//	0x168 32  FS_HCINT3     OTG_FS host channel-3 interrupt register (OTG_FS_HCINT3)
//	0x16C 32  FS_HCINTMSK3  OTG_FS host channel-3 mask register (OTG_FS_HCINTMSK3)
//	0x170 32  FS_HCTSIZ3    OTG_FS host channel-3 transfer size register
//	0x180 32  FS_HCCHAR4    OTG_FS host channel-4 characteristics register (OTG_FS_HCCHAR4)
//	0x188 32  FS_HCINT4     OTG_FS host channel-4 interrupt register (OTG_FS_HCINT4)
//	0x18C 32  FS_HCINTMSK4  OTG_FS host channel-4 mask register (OTG_FS_HCINTMSK4)
//	0x190 32  FS_HCTSIZ4    OTG_FS host channel-x transfer size register
//	0x1A0 32  FS_HCCHAR5    OTG_FS host channel-5 characteristics register (OTG_FS_HCCHAR5)
//	0x1A8 32  FS_HCINT5     OTG_FS host channel-5 interrupt register (OTG_FS_HCINT5)
//	0x1AC 32  FS_HCINTMSK5  OTG_FS host channel-5 mask register (OTG_FS_HCINTMSK5)
//	0x1B0 32  FS_HCTSIZ5    OTG_FS host channel-5 transfer size register
//	0x1C0 32  FS_HCCHAR6    OTG_FS host channel-6 characteristics register (OTG_FS_HCCHAR6)
//	0x1C8 32  FS_HCINT6     OTG_FS host channel-6 interrupt register (OTG_FS_HCINT6)
//	0x1CC 32  FS_HCINTMSK6  OTG_FS host channel-6 mask register (OTG_FS_HCINTMSK6)
//	0x1D0 32  FS_HCTSIZ6    OTG_FS host channel-6 transfer size register
//	0x1E0 32  FS_HCCHAR7    OTG_FS host channel-7 characteristics register (OTG_FS_HCCHAR7)
//	0x1E8 32  FS_HCINT7     OTG_FS host channel-7 interrupt register (OTG_FS_HCINT7)
//	0x1EC 32  FS_HCINTMSK7  OTG_FS host channel-7 mask register (OTG_FS_HCINTMSK7)
//	0x1F0 32  FS_HCTSIZ7    OTG_FS host channel-7 transfer size register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package otg_fs_host

const (
	FSLSPCS FS_HCFG = 0x03 << 0 //+ FS/LS PHY clock select
	FSLSS   FS_HCFG = 0x01 << 2 //+ FS- and LS-only support
)

const (
	FSLSPCSn = 0
	FSLSSn   = 2
)

const (
	FRIVL HFIR = 0xFFFF << 0 //+ Frame interval
)

const (
	FRIVLn = 0
)

const (
	FRNUM FS_HFNUM = 0xFFFF << 0  //+ Frame number
	FTREM FS_HFNUM = 0xFFFF << 16 //+ Frame time remaining
)

const (
	FRNUMn = 0
	FTREMn = 16
)

const (
	PTXFSAVL FS_HPTXSTS = 0xFFFF << 0 //+ Periodic transmit data FIFO space available
	PTXQSAV  FS_HPTXSTS = 0xFF << 16  //+ Periodic transmit request queue space available
	PTXQTOP  FS_HPTXSTS = 0xFF << 24  //+ Top of the periodic transmit request queue
)

const (
	PTXFSAVLn = 0
	PTXQSAVn  = 16
	PTXQTOPn  = 24
)

const (
	HAINT HAINT = 0xFFFF << 0 //+ Channel interrupts
)

const (
	HAINTn = 0
)

const (
	HAINTM HAINTMSK = 0xFFFF << 0 //+ Channel interrupt mask
)

const (
	HAINTMn = 0
)

const (
	PCSTS   FS_HPRT = 0x01 << 0  //+ Port connect status
	PCDET   FS_HPRT = 0x01 << 1  //+ Port connect detected
	PENA    FS_HPRT = 0x01 << 2  //+ Port enable
	PENCHNG FS_HPRT = 0x01 << 3  //+ Port enable/disable change
	POCA    FS_HPRT = 0x01 << 4  //+ Port overcurrent active
	POCCHNG FS_HPRT = 0x01 << 5  //+ Port overcurrent change
	PRES    FS_HPRT = 0x01 << 6  //+ Port resume
	PSUSP   FS_HPRT = 0x01 << 7  //+ Port suspend
	PRST    FS_HPRT = 0x01 << 8  //+ Port reset
	PLSTS   FS_HPRT = 0x03 << 10 //+ Port line status
	PPWR    FS_HPRT = 0x01 << 12 //+ Port power
	PTCTL   FS_HPRT = 0x0F << 13 //+ Port test control
	PSPD    FS_HPRT = 0x03 << 17 //+ Port speed
)

const (
	PCSTSn   = 0
	PCDETn   = 1
	PENAn    = 2
	PENCHNGn = 3
	POCAn    = 4
	POCCHNGn = 5
	PRESn    = 6
	PSUSPn   = 7
	PRSTn    = 8
	PLSTSn   = 10
	PPWRn    = 12
	PTCTLn   = 13
	PSPDn    = 17
)

const (
	MPSIZ  FS_HCCHAR0 = 0x7FF << 0 //+ Maximum packet size
	EPNUM  FS_HCCHAR0 = 0x0F << 11 //+ Endpoint number
	EPDIR  FS_HCCHAR0 = 0x01 << 15 //+ Endpoint direction
	LSDEV  FS_HCCHAR0 = 0x01 << 17 //+ Low-speed device
	EPTYP  FS_HCCHAR0 = 0x03 << 18 //+ Endpoint type
	MCNT   FS_HCCHAR0 = 0x03 << 20 //+ Multicount
	DAD    FS_HCCHAR0 = 0x7F << 22 //+ Device address
	ODDFRM FS_HCCHAR0 = 0x01 << 29 //+ Odd frame
	CHDIS  FS_HCCHAR0 = 0x01 << 30 //+ Channel disable
	CHENA  FS_HCCHAR0 = 0x01 << 31 //+ Channel enable
)

const (
	MPSIZn  = 0
	EPNUMn  = 11
	EPDIRn  = 15
	LSDEVn  = 17
	EPTYPn  = 18
	MCNTn   = 20
	DADn    = 22
	ODDFRMn = 29
	CHDISn  = 30
	CHENAn  = 31
)

const (
	XFRC  FS_HCINT0 = 0x01 << 0  //+ Transfer completed
	CHH   FS_HCINT0 = 0x01 << 1  //+ Channel halted
	STALL FS_HCINT0 = 0x01 << 3  //+ STALL response received interrupt
	NAK   FS_HCINT0 = 0x01 << 4  //+ NAK response received interrupt
	ACK   FS_HCINT0 = 0x01 << 5  //+ ACK response received/transmitted interrupt
	TXERR FS_HCINT0 = 0x01 << 7  //+ Transaction error
	BBERR FS_HCINT0 = 0x01 << 8  //+ Babble error
	FRMOR FS_HCINT0 = 0x01 << 9  //+ Frame overrun
	DTERR FS_HCINT0 = 0x01 << 10 //+ Data toggle error
)

const (
	XFRCn  = 0
	CHHn   = 1
	STALLn = 3
	NAKn   = 4
	ACKn   = 5
	TXERRn = 7
	BBERRn = 8
	FRMORn = 9
	DTERRn = 10
)

const (
	XFRCM  FS_HCINTMSK0 = 0x01 << 0  //+ Transfer completed mask
	CHHM   FS_HCINTMSK0 = 0x01 << 1  //+ Channel halted mask
	STALLM FS_HCINTMSK0 = 0x01 << 3  //+ STALL response received interrupt mask
	NAKM   FS_HCINTMSK0 = 0x01 << 4  //+ NAK response received interrupt mask
	ACKM   FS_HCINTMSK0 = 0x01 << 5  //+ ACK response received/transmitted interrupt mask
	NYET   FS_HCINTMSK0 = 0x01 << 6  //+ response received interrupt mask
	TXERRM FS_HCINTMSK0 = 0x01 << 7  //+ Transaction error mask
	BBERRM FS_HCINTMSK0 = 0x01 << 8  //+ Babble error mask
	FRMORM FS_HCINTMSK0 = 0x01 << 9  //+ Frame overrun mask
	DTERRM FS_HCINTMSK0 = 0x01 << 10 //+ Data toggle error mask
)

const (
	XFRCMn  = 0
	CHHMn   = 1
	STALLMn = 3
	NAKMn   = 4
	ACKMn   = 5
	NYETn   = 6
	TXERRMn = 7
	BBERRMn = 8
	FRMORMn = 9
	DTERRMn = 10
)

const (
	XFRSIZ FS_HCTSIZ0 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT FS_HCTSIZ0 = 0x3FF << 19  //+ Packet count
	DPID   FS_HCTSIZ0 = 0x03 << 29   //+ Data PID
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	DPIDn   = 29
)

const (
	MPSIZ  FS_HCCHAR1 = 0x7FF << 0 //+ Maximum packet size
	EPNUM  FS_HCCHAR1 = 0x0F << 11 //+ Endpoint number
	EPDIR  FS_HCCHAR1 = 0x01 << 15 //+ Endpoint direction
	LSDEV  FS_HCCHAR1 = 0x01 << 17 //+ Low-speed device
	EPTYP  FS_HCCHAR1 = 0x03 << 18 //+ Endpoint type
	MCNT   FS_HCCHAR1 = 0x03 << 20 //+ Multicount
	DAD    FS_HCCHAR1 = 0x7F << 22 //+ Device address
	ODDFRM FS_HCCHAR1 = 0x01 << 29 //+ Odd frame
	CHDIS  FS_HCCHAR1 = 0x01 << 30 //+ Channel disable
	CHENA  FS_HCCHAR1 = 0x01 << 31 //+ Channel enable
)

const (
	MPSIZn  = 0
	EPNUMn  = 11
	EPDIRn  = 15
	LSDEVn  = 17
	EPTYPn  = 18
	MCNTn   = 20
	DADn    = 22
	ODDFRMn = 29
	CHDISn  = 30
	CHENAn  = 31
)

const (
	XFRC  FS_HCINT1 = 0x01 << 0  //+ Transfer completed
	CHH   FS_HCINT1 = 0x01 << 1  //+ Channel halted
	STALL FS_HCINT1 = 0x01 << 3  //+ STALL response received interrupt
	NAK   FS_HCINT1 = 0x01 << 4  //+ NAK response received interrupt
	ACK   FS_HCINT1 = 0x01 << 5  //+ ACK response received/transmitted interrupt
	TXERR FS_HCINT1 = 0x01 << 7  //+ Transaction error
	BBERR FS_HCINT1 = 0x01 << 8  //+ Babble error
	FRMOR FS_HCINT1 = 0x01 << 9  //+ Frame overrun
	DTERR FS_HCINT1 = 0x01 << 10 //+ Data toggle error
)

const (
	XFRCn  = 0
	CHHn   = 1
	STALLn = 3
	NAKn   = 4
	ACKn   = 5
	TXERRn = 7
	BBERRn = 8
	FRMORn = 9
	DTERRn = 10
)

const (
	XFRCM  FS_HCINTMSK1 = 0x01 << 0  //+ Transfer completed mask
	CHHM   FS_HCINTMSK1 = 0x01 << 1  //+ Channel halted mask
	STALLM FS_HCINTMSK1 = 0x01 << 3  //+ STALL response received interrupt mask
	NAKM   FS_HCINTMSK1 = 0x01 << 4  //+ NAK response received interrupt mask
	ACKM   FS_HCINTMSK1 = 0x01 << 5  //+ ACK response received/transmitted interrupt mask
	NYET   FS_HCINTMSK1 = 0x01 << 6  //+ response received interrupt mask
	TXERRM FS_HCINTMSK1 = 0x01 << 7  //+ Transaction error mask
	BBERRM FS_HCINTMSK1 = 0x01 << 8  //+ Babble error mask
	FRMORM FS_HCINTMSK1 = 0x01 << 9  //+ Frame overrun mask
	DTERRM FS_HCINTMSK1 = 0x01 << 10 //+ Data toggle error mask
)

const (
	XFRCMn  = 0
	CHHMn   = 1
	STALLMn = 3
	NAKMn   = 4
	ACKMn   = 5
	NYETn   = 6
	TXERRMn = 7
	BBERRMn = 8
	FRMORMn = 9
	DTERRMn = 10
)

const (
	XFRSIZ FS_HCTSIZ1 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT FS_HCTSIZ1 = 0x3FF << 19  //+ Packet count
	DPID   FS_HCTSIZ1 = 0x03 << 29   //+ Data PID
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	DPIDn   = 29
)

const (
	MPSIZ  FS_HCCHAR2 = 0x7FF << 0 //+ Maximum packet size
	EPNUM  FS_HCCHAR2 = 0x0F << 11 //+ Endpoint number
	EPDIR  FS_HCCHAR2 = 0x01 << 15 //+ Endpoint direction
	LSDEV  FS_HCCHAR2 = 0x01 << 17 //+ Low-speed device
	EPTYP  FS_HCCHAR2 = 0x03 << 18 //+ Endpoint type
	MCNT   FS_HCCHAR2 = 0x03 << 20 //+ Multicount
	DAD    FS_HCCHAR2 = 0x7F << 22 //+ Device address
	ODDFRM FS_HCCHAR2 = 0x01 << 29 //+ Odd frame
	CHDIS  FS_HCCHAR2 = 0x01 << 30 //+ Channel disable
	CHENA  FS_HCCHAR2 = 0x01 << 31 //+ Channel enable
)

const (
	MPSIZn  = 0
	EPNUMn  = 11
	EPDIRn  = 15
	LSDEVn  = 17
	EPTYPn  = 18
	MCNTn   = 20
	DADn    = 22
	ODDFRMn = 29
	CHDISn  = 30
	CHENAn  = 31
)

const (
	XFRC  FS_HCINT2 = 0x01 << 0  //+ Transfer completed
	CHH   FS_HCINT2 = 0x01 << 1  //+ Channel halted
	STALL FS_HCINT2 = 0x01 << 3  //+ STALL response received interrupt
	NAK   FS_HCINT2 = 0x01 << 4  //+ NAK response received interrupt
	ACK   FS_HCINT2 = 0x01 << 5  //+ ACK response received/transmitted interrupt
	TXERR FS_HCINT2 = 0x01 << 7  //+ Transaction error
	BBERR FS_HCINT2 = 0x01 << 8  //+ Babble error
	FRMOR FS_HCINT2 = 0x01 << 9  //+ Frame overrun
	DTERR FS_HCINT2 = 0x01 << 10 //+ Data toggle error
)

const (
	XFRCn  = 0
	CHHn   = 1
	STALLn = 3
	NAKn   = 4
	ACKn   = 5
	TXERRn = 7
	BBERRn = 8
	FRMORn = 9
	DTERRn = 10
)

const (
	XFRCM  FS_HCINTMSK2 = 0x01 << 0  //+ Transfer completed mask
	CHHM   FS_HCINTMSK2 = 0x01 << 1  //+ Channel halted mask
	STALLM FS_HCINTMSK2 = 0x01 << 3  //+ STALL response received interrupt mask
	NAKM   FS_HCINTMSK2 = 0x01 << 4  //+ NAK response received interrupt mask
	ACKM   FS_HCINTMSK2 = 0x01 << 5  //+ ACK response received/transmitted interrupt mask
	NYET   FS_HCINTMSK2 = 0x01 << 6  //+ response received interrupt mask
	TXERRM FS_HCINTMSK2 = 0x01 << 7  //+ Transaction error mask
	BBERRM FS_HCINTMSK2 = 0x01 << 8  //+ Babble error mask
	FRMORM FS_HCINTMSK2 = 0x01 << 9  //+ Frame overrun mask
	DTERRM FS_HCINTMSK2 = 0x01 << 10 //+ Data toggle error mask
)

const (
	XFRCMn  = 0
	CHHMn   = 1
	STALLMn = 3
	NAKMn   = 4
	ACKMn   = 5
	NYETn   = 6
	TXERRMn = 7
	BBERRMn = 8
	FRMORMn = 9
	DTERRMn = 10
)

const (
	XFRSIZ FS_HCTSIZ2 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT FS_HCTSIZ2 = 0x3FF << 19  //+ Packet count
	DPID   FS_HCTSIZ2 = 0x03 << 29   //+ Data PID
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	DPIDn   = 29
)

const (
	MPSIZ  FS_HCCHAR3 = 0x7FF << 0 //+ Maximum packet size
	EPNUM  FS_HCCHAR3 = 0x0F << 11 //+ Endpoint number
	EPDIR  FS_HCCHAR3 = 0x01 << 15 //+ Endpoint direction
	LSDEV  FS_HCCHAR3 = 0x01 << 17 //+ Low-speed device
	EPTYP  FS_HCCHAR3 = 0x03 << 18 //+ Endpoint type
	MCNT   FS_HCCHAR3 = 0x03 << 20 //+ Multicount
	DAD    FS_HCCHAR3 = 0x7F << 22 //+ Device address
	ODDFRM FS_HCCHAR3 = 0x01 << 29 //+ Odd frame
	CHDIS  FS_HCCHAR3 = 0x01 << 30 //+ Channel disable
	CHENA  FS_HCCHAR3 = 0x01 << 31 //+ Channel enable
)

const (
	MPSIZn  = 0
	EPNUMn  = 11
	EPDIRn  = 15
	LSDEVn  = 17
	EPTYPn  = 18
	MCNTn   = 20
	DADn    = 22
	ODDFRMn = 29
	CHDISn  = 30
	CHENAn  = 31
)

const (
	XFRC  FS_HCINT3 = 0x01 << 0  //+ Transfer completed
	CHH   FS_HCINT3 = 0x01 << 1  //+ Channel halted
	STALL FS_HCINT3 = 0x01 << 3  //+ STALL response received interrupt
	NAK   FS_HCINT3 = 0x01 << 4  //+ NAK response received interrupt
	ACK   FS_HCINT3 = 0x01 << 5  //+ ACK response received/transmitted interrupt
	TXERR FS_HCINT3 = 0x01 << 7  //+ Transaction error
	BBERR FS_HCINT3 = 0x01 << 8  //+ Babble error
	FRMOR FS_HCINT3 = 0x01 << 9  //+ Frame overrun
	DTERR FS_HCINT3 = 0x01 << 10 //+ Data toggle error
)

const (
	XFRCn  = 0
	CHHn   = 1
	STALLn = 3
	NAKn   = 4
	ACKn   = 5
	TXERRn = 7
	BBERRn = 8
	FRMORn = 9
	DTERRn = 10
)

const (
	XFRCM  FS_HCINTMSK3 = 0x01 << 0  //+ Transfer completed mask
	CHHM   FS_HCINTMSK3 = 0x01 << 1  //+ Channel halted mask
	STALLM FS_HCINTMSK3 = 0x01 << 3  //+ STALL response received interrupt mask
	NAKM   FS_HCINTMSK3 = 0x01 << 4  //+ NAK response received interrupt mask
	ACKM   FS_HCINTMSK3 = 0x01 << 5  //+ ACK response received/transmitted interrupt mask
	NYET   FS_HCINTMSK3 = 0x01 << 6  //+ response received interrupt mask
	TXERRM FS_HCINTMSK3 = 0x01 << 7  //+ Transaction error mask
	BBERRM FS_HCINTMSK3 = 0x01 << 8  //+ Babble error mask
	FRMORM FS_HCINTMSK3 = 0x01 << 9  //+ Frame overrun mask
	DTERRM FS_HCINTMSK3 = 0x01 << 10 //+ Data toggle error mask
)

const (
	XFRCMn  = 0
	CHHMn   = 1
	STALLMn = 3
	NAKMn   = 4
	ACKMn   = 5
	NYETn   = 6
	TXERRMn = 7
	BBERRMn = 8
	FRMORMn = 9
	DTERRMn = 10
)

const (
	XFRSIZ FS_HCTSIZ3 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT FS_HCTSIZ3 = 0x3FF << 19  //+ Packet count
	DPID   FS_HCTSIZ3 = 0x03 << 29   //+ Data PID
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	DPIDn   = 29
)

const (
	MPSIZ  FS_HCCHAR4 = 0x7FF << 0 //+ Maximum packet size
	EPNUM  FS_HCCHAR4 = 0x0F << 11 //+ Endpoint number
	EPDIR  FS_HCCHAR4 = 0x01 << 15 //+ Endpoint direction
	LSDEV  FS_HCCHAR4 = 0x01 << 17 //+ Low-speed device
	EPTYP  FS_HCCHAR4 = 0x03 << 18 //+ Endpoint type
	MCNT   FS_HCCHAR4 = 0x03 << 20 //+ Multicount
	DAD    FS_HCCHAR4 = 0x7F << 22 //+ Device address
	ODDFRM FS_HCCHAR4 = 0x01 << 29 //+ Odd frame
	CHDIS  FS_HCCHAR4 = 0x01 << 30 //+ Channel disable
	CHENA  FS_HCCHAR4 = 0x01 << 31 //+ Channel enable
)

const (
	MPSIZn  = 0
	EPNUMn  = 11
	EPDIRn  = 15
	LSDEVn  = 17
	EPTYPn  = 18
	MCNTn   = 20
	DADn    = 22
	ODDFRMn = 29
	CHDISn  = 30
	CHENAn  = 31
)

const (
	XFRC  FS_HCINT4 = 0x01 << 0  //+ Transfer completed
	CHH   FS_HCINT4 = 0x01 << 1  //+ Channel halted
	STALL FS_HCINT4 = 0x01 << 3  //+ STALL response received interrupt
	NAK   FS_HCINT4 = 0x01 << 4  //+ NAK response received interrupt
	ACK   FS_HCINT4 = 0x01 << 5  //+ ACK response received/transmitted interrupt
	TXERR FS_HCINT4 = 0x01 << 7  //+ Transaction error
	BBERR FS_HCINT4 = 0x01 << 8  //+ Babble error
	FRMOR FS_HCINT4 = 0x01 << 9  //+ Frame overrun
	DTERR FS_HCINT4 = 0x01 << 10 //+ Data toggle error
)

const (
	XFRCn  = 0
	CHHn   = 1
	STALLn = 3
	NAKn   = 4
	ACKn   = 5
	TXERRn = 7
	BBERRn = 8
	FRMORn = 9
	DTERRn = 10
)

const (
	XFRCM  FS_HCINTMSK4 = 0x01 << 0  //+ Transfer completed mask
	CHHM   FS_HCINTMSK4 = 0x01 << 1  //+ Channel halted mask
	STALLM FS_HCINTMSK4 = 0x01 << 3  //+ STALL response received interrupt mask
	NAKM   FS_HCINTMSK4 = 0x01 << 4  //+ NAK response received interrupt mask
	ACKM   FS_HCINTMSK4 = 0x01 << 5  //+ ACK response received/transmitted interrupt mask
	NYET   FS_HCINTMSK4 = 0x01 << 6  //+ response received interrupt mask
	TXERRM FS_HCINTMSK4 = 0x01 << 7  //+ Transaction error mask
	BBERRM FS_HCINTMSK4 = 0x01 << 8  //+ Babble error mask
	FRMORM FS_HCINTMSK4 = 0x01 << 9  //+ Frame overrun mask
	DTERRM FS_HCINTMSK4 = 0x01 << 10 //+ Data toggle error mask
)

const (
	XFRCMn  = 0
	CHHMn   = 1
	STALLMn = 3
	NAKMn   = 4
	ACKMn   = 5
	NYETn   = 6
	TXERRMn = 7
	BBERRMn = 8
	FRMORMn = 9
	DTERRMn = 10
)

const (
	XFRSIZ FS_HCTSIZ4 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT FS_HCTSIZ4 = 0x3FF << 19  //+ Packet count
	DPID   FS_HCTSIZ4 = 0x03 << 29   //+ Data PID
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	DPIDn   = 29
)

const (
	MPSIZ  FS_HCCHAR5 = 0x7FF << 0 //+ Maximum packet size
	EPNUM  FS_HCCHAR5 = 0x0F << 11 //+ Endpoint number
	EPDIR  FS_HCCHAR5 = 0x01 << 15 //+ Endpoint direction
	LSDEV  FS_HCCHAR5 = 0x01 << 17 //+ Low-speed device
	EPTYP  FS_HCCHAR5 = 0x03 << 18 //+ Endpoint type
	MCNT   FS_HCCHAR5 = 0x03 << 20 //+ Multicount
	DAD    FS_HCCHAR5 = 0x7F << 22 //+ Device address
	ODDFRM FS_HCCHAR5 = 0x01 << 29 //+ Odd frame
	CHDIS  FS_HCCHAR5 = 0x01 << 30 //+ Channel disable
	CHENA  FS_HCCHAR5 = 0x01 << 31 //+ Channel enable
)

const (
	MPSIZn  = 0
	EPNUMn  = 11
	EPDIRn  = 15
	LSDEVn  = 17
	EPTYPn  = 18
	MCNTn   = 20
	DADn    = 22
	ODDFRMn = 29
	CHDISn  = 30
	CHENAn  = 31
)

const (
	XFRC  FS_HCINT5 = 0x01 << 0  //+ Transfer completed
	CHH   FS_HCINT5 = 0x01 << 1  //+ Channel halted
	STALL FS_HCINT5 = 0x01 << 3  //+ STALL response received interrupt
	NAK   FS_HCINT5 = 0x01 << 4  //+ NAK response received interrupt
	ACK   FS_HCINT5 = 0x01 << 5  //+ ACK response received/transmitted interrupt
	TXERR FS_HCINT5 = 0x01 << 7  //+ Transaction error
	BBERR FS_HCINT5 = 0x01 << 8  //+ Babble error
	FRMOR FS_HCINT5 = 0x01 << 9  //+ Frame overrun
	DTERR FS_HCINT5 = 0x01 << 10 //+ Data toggle error
)

const (
	XFRCn  = 0
	CHHn   = 1
	STALLn = 3
	NAKn   = 4
	ACKn   = 5
	TXERRn = 7
	BBERRn = 8
	FRMORn = 9
	DTERRn = 10
)

const (
	XFRCM  FS_HCINTMSK5 = 0x01 << 0  //+ Transfer completed mask
	CHHM   FS_HCINTMSK5 = 0x01 << 1  //+ Channel halted mask
	STALLM FS_HCINTMSK5 = 0x01 << 3  //+ STALL response received interrupt mask
	NAKM   FS_HCINTMSK5 = 0x01 << 4  //+ NAK response received interrupt mask
	ACKM   FS_HCINTMSK5 = 0x01 << 5  //+ ACK response received/transmitted interrupt mask
	NYET   FS_HCINTMSK5 = 0x01 << 6  //+ response received interrupt mask
	TXERRM FS_HCINTMSK5 = 0x01 << 7  //+ Transaction error mask
	BBERRM FS_HCINTMSK5 = 0x01 << 8  //+ Babble error mask
	FRMORM FS_HCINTMSK5 = 0x01 << 9  //+ Frame overrun mask
	DTERRM FS_HCINTMSK5 = 0x01 << 10 //+ Data toggle error mask
)

const (
	XFRCMn  = 0
	CHHMn   = 1
	STALLMn = 3
	NAKMn   = 4
	ACKMn   = 5
	NYETn   = 6
	TXERRMn = 7
	BBERRMn = 8
	FRMORMn = 9
	DTERRMn = 10
)

const (
	XFRSIZ FS_HCTSIZ5 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT FS_HCTSIZ5 = 0x3FF << 19  //+ Packet count
	DPID   FS_HCTSIZ5 = 0x03 << 29   //+ Data PID
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	DPIDn   = 29
)

const (
	MPSIZ  FS_HCCHAR6 = 0x7FF << 0 //+ Maximum packet size
	EPNUM  FS_HCCHAR6 = 0x0F << 11 //+ Endpoint number
	EPDIR  FS_HCCHAR6 = 0x01 << 15 //+ Endpoint direction
	LSDEV  FS_HCCHAR6 = 0x01 << 17 //+ Low-speed device
	EPTYP  FS_HCCHAR6 = 0x03 << 18 //+ Endpoint type
	MCNT   FS_HCCHAR6 = 0x03 << 20 //+ Multicount
	DAD    FS_HCCHAR6 = 0x7F << 22 //+ Device address
	ODDFRM FS_HCCHAR6 = 0x01 << 29 //+ Odd frame
	CHDIS  FS_HCCHAR6 = 0x01 << 30 //+ Channel disable
	CHENA  FS_HCCHAR6 = 0x01 << 31 //+ Channel enable
)

const (
	MPSIZn  = 0
	EPNUMn  = 11
	EPDIRn  = 15
	LSDEVn  = 17
	EPTYPn  = 18
	MCNTn   = 20
	DADn    = 22
	ODDFRMn = 29
	CHDISn  = 30
	CHENAn  = 31
)

const (
	XFRC  FS_HCINT6 = 0x01 << 0  //+ Transfer completed
	CHH   FS_HCINT6 = 0x01 << 1  //+ Channel halted
	STALL FS_HCINT6 = 0x01 << 3  //+ STALL response received interrupt
	NAK   FS_HCINT6 = 0x01 << 4  //+ NAK response received interrupt
	ACK   FS_HCINT6 = 0x01 << 5  //+ ACK response received/transmitted interrupt
	TXERR FS_HCINT6 = 0x01 << 7  //+ Transaction error
	BBERR FS_HCINT6 = 0x01 << 8  //+ Babble error
	FRMOR FS_HCINT6 = 0x01 << 9  //+ Frame overrun
	DTERR FS_HCINT6 = 0x01 << 10 //+ Data toggle error
)

const (
	XFRCn  = 0
	CHHn   = 1
	STALLn = 3
	NAKn   = 4
	ACKn   = 5
	TXERRn = 7
	BBERRn = 8
	FRMORn = 9
	DTERRn = 10
)

const (
	XFRCM  FS_HCINTMSK6 = 0x01 << 0  //+ Transfer completed mask
	CHHM   FS_HCINTMSK6 = 0x01 << 1  //+ Channel halted mask
	STALLM FS_HCINTMSK6 = 0x01 << 3  //+ STALL response received interrupt mask
	NAKM   FS_HCINTMSK6 = 0x01 << 4  //+ NAK response received interrupt mask
	ACKM   FS_HCINTMSK6 = 0x01 << 5  //+ ACK response received/transmitted interrupt mask
	NYET   FS_HCINTMSK6 = 0x01 << 6  //+ response received interrupt mask
	TXERRM FS_HCINTMSK6 = 0x01 << 7  //+ Transaction error mask
	BBERRM FS_HCINTMSK6 = 0x01 << 8  //+ Babble error mask
	FRMORM FS_HCINTMSK6 = 0x01 << 9  //+ Frame overrun mask
	DTERRM FS_HCINTMSK6 = 0x01 << 10 //+ Data toggle error mask
)

const (
	XFRCMn  = 0
	CHHMn   = 1
	STALLMn = 3
	NAKMn   = 4
	ACKMn   = 5
	NYETn   = 6
	TXERRMn = 7
	BBERRMn = 8
	FRMORMn = 9
	DTERRMn = 10
)

const (
	XFRSIZ FS_HCTSIZ6 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT FS_HCTSIZ6 = 0x3FF << 19  //+ Packet count
	DPID   FS_HCTSIZ6 = 0x03 << 29   //+ Data PID
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	DPIDn   = 29
)

const (
	MPSIZ  FS_HCCHAR7 = 0x7FF << 0 //+ Maximum packet size
	EPNUM  FS_HCCHAR7 = 0x0F << 11 //+ Endpoint number
	EPDIR  FS_HCCHAR7 = 0x01 << 15 //+ Endpoint direction
	LSDEV  FS_HCCHAR7 = 0x01 << 17 //+ Low-speed device
	EPTYP  FS_HCCHAR7 = 0x03 << 18 //+ Endpoint type
	MCNT   FS_HCCHAR7 = 0x03 << 20 //+ Multicount
	DAD    FS_HCCHAR7 = 0x7F << 22 //+ Device address
	ODDFRM FS_HCCHAR7 = 0x01 << 29 //+ Odd frame
	CHDIS  FS_HCCHAR7 = 0x01 << 30 //+ Channel disable
	CHENA  FS_HCCHAR7 = 0x01 << 31 //+ Channel enable
)

const (
	MPSIZn  = 0
	EPNUMn  = 11
	EPDIRn  = 15
	LSDEVn  = 17
	EPTYPn  = 18
	MCNTn   = 20
	DADn    = 22
	ODDFRMn = 29
	CHDISn  = 30
	CHENAn  = 31
)

const (
	XFRC  FS_HCINT7 = 0x01 << 0  //+ Transfer completed
	CHH   FS_HCINT7 = 0x01 << 1  //+ Channel halted
	STALL FS_HCINT7 = 0x01 << 3  //+ STALL response received interrupt
	NAK   FS_HCINT7 = 0x01 << 4  //+ NAK response received interrupt
	ACK   FS_HCINT7 = 0x01 << 5  //+ ACK response received/transmitted interrupt
	TXERR FS_HCINT7 = 0x01 << 7  //+ Transaction error
	BBERR FS_HCINT7 = 0x01 << 8  //+ Babble error
	FRMOR FS_HCINT7 = 0x01 << 9  //+ Frame overrun
	DTERR FS_HCINT7 = 0x01 << 10 //+ Data toggle error
)

const (
	XFRCn  = 0
	CHHn   = 1
	STALLn = 3
	NAKn   = 4
	ACKn   = 5
	TXERRn = 7
	BBERRn = 8
	FRMORn = 9
	DTERRn = 10
)

const (
	XFRCM  FS_HCINTMSK7 = 0x01 << 0  //+ Transfer completed mask
	CHHM   FS_HCINTMSK7 = 0x01 << 1  //+ Channel halted mask
	STALLM FS_HCINTMSK7 = 0x01 << 3  //+ STALL response received interrupt mask
	NAKM   FS_HCINTMSK7 = 0x01 << 4  //+ NAK response received interrupt mask
	ACKM   FS_HCINTMSK7 = 0x01 << 5  //+ ACK response received/transmitted interrupt mask
	NYET   FS_HCINTMSK7 = 0x01 << 6  //+ response received interrupt mask
	TXERRM FS_HCINTMSK7 = 0x01 << 7  //+ Transaction error mask
	BBERRM FS_HCINTMSK7 = 0x01 << 8  //+ Babble error mask
	FRMORM FS_HCINTMSK7 = 0x01 << 9  //+ Frame overrun mask
	DTERRM FS_HCINTMSK7 = 0x01 << 10 //+ Data toggle error mask
)

const (
	XFRCMn  = 0
	CHHMn   = 1
	STALLMn = 3
	NAKMn   = 4
	ACKMn   = 5
	NYETn   = 6
	TXERRMn = 7
	BBERRMn = 8
	FRMORMn = 9
	DTERRMn = 10
)

const (
	XFRSIZ FS_HCTSIZ7 = 0x7FFFF << 0 //+ Transfer size
	PKTCNT FS_HCTSIZ7 = 0x3FF << 19  //+ Packet count
	DPID   FS_HCTSIZ7 = 0x03 << 29   //+ Data PID
)

const (
	XFRSIZn = 0
	PKTCNTn = 19
	DPIDn   = 29
)
