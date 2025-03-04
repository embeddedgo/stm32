// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32h7x3

// Package ethernet_mac provides access to the registers of the Ethernet_MAC peripheral.
//
// Instances:
//
//	Ethernet_MAC  Ethernet_MAC_BASE  -  ETH  Ethernet: media access control (MAC)
//
// Registers:
//
//	0x000 32  MACCR                               Operating mode configuration register
//	0x004 32  MACECR                              Extended operating mode configuration register
//	0x008 32  MACPFR                              Packet filtering control register
//	0x00C 32  MACWTR                              Watchdog timeout register
//	0x010 32  MACHT0R                             Hash Table 0 register
//	0x014 32  MACHT1R                             Hash Table 1 register
//	0x050 32  MACVTR                              VLAN tag register
//	0x058 32  MACVHTR                             VLAN Hash table register
//	0x060 32  MACVIR                              VLAN inclusion register
//	0x064 32  MACIVIR                             Inner VLAN inclusion register
//	0x070 32  MACQTxFCR                           Tx Queue flow control register
//	0x090 32  MACRxFCR                            Rx flow control register
//	0x0B0 32  MACISR                              Interrupt status register
//	0x0B4 32  MACIER                              Interrupt enable register
//	0x0B8 32  MACRxTxSR                           Rx Tx status register
//	0x0C0 32  MACPCSR                             PMT control status register
//	0x0C4 32  MACRWKPFR                           Remove wakeup packet filter register
//	0x0D0 32  MACLCSR                             LPI control status register
//	0x0D4 32  MACLTCR                             LPI timers control register
//	0x0D8 32  MACLETR                             LPI entry timer register
//	0x0DC 32  MAC1USTCR                           1-microsecond-tick counter register
//	0x110 32  MACVR                               Version register
//	0x114 32  MACDR                               Debug register
//	0x120 32  MACHWF1R                            HW feature 1 register
//	0x124 32  MACHWF2R                            HW feature 2 register
//	0x200 32  MACMDIOAR                           MDIO address register
//	0x204 32  MACMDIODR                           MDIO data register
//	0x300 32  MACA0HR                             Address 0 high register
//	0x304 32  MACA0LR                             Address 0 low register
//	0x308 32  MACA1HR                             Address 1 high register
//	0x30C 32  MACA1LR                             Address 1 low register
//	0x310 32  MACA2HR                             Address 2 high register
//	0x314 32  MACA2LR                             Address 2 low register
//	0x318 32  MACA3HR                             Address 3 high register
//	0x31C 32  MACA3LR                             Address 3 low register
//	0x700 32  MMC_CONTROL                         MMC control register
//	0x704 32  MMC_RX_INTERRUPT                    MMC Rx interrupt register
//	0x708 32  MMC_TX_INTERRUPT                    MMC Tx interrupt register
//	0x70C 32  MMC_RX_INTERRUPT_MASK               MMC Rx interrupt mask register
//	0x710 32  MMC_TX_INTERRUPT_MASK               MMC Tx interrupt mask register
//	0x74C 32  TX_SINGLE_COLLISION_GOOD_PACKETS    Tx single collision good packets register
//	0x750 32  TX_MULTIPLE_COLLISION_GOOD_PACKETS  Tx multiple collision good packets register
//	0x768 32  TX_PACKET_COUNT_GOOD                Tx packet count good register
//	0x794 32  RX_CRC_ERROR_PACKETS                Rx CRC error packets register
//	0x798 32  RX_ALIGNMENT_ERROR_PACKETS          Rx alignment error packets register
//	0x7C4 32  RX_UNICAST_PACKETS_GOOD             Rx unicast packets good register
//	0x7EC 32  TX_LPI_USEC_CNTR                    Tx LPI microsecond timer register
//	0x7F0 32  TX_LPI_TRAN_CNTR                    Tx LPI transition counter register
//	0x7F4 32  RX_LPI_USEC_CNTR                    Rx LPI microsecond counter register
//	0x7F8 32  RX_LPI_TRAN_CNTR                    Rx LPI transition counter register
//	0x900 32  MACL3L4C0R                          L3 and L4 control 0 register
//	0x904 32  MACL4A0R                            Layer4 address filter 0 register
//	0x910 32  MACL3A00R                           MACL3A00R
//	0x914 32  MACL3A10R                           Layer3 address 1 filter 0 register
//	0x918 32  MACL3A20                            Layer3 Address 2 filter 0 register
//	0x91C 32  MACL3A30                            Layer3 Address 3 filter 0 register
//	0x930 32  MACL3L4C1R                          L3 and L4 control 1 register
//	0x934 32  MACL4A1R                            Layer 4 address filter 1 register
//	0x940 32  MACL3A01R                           Layer3 address 0 filter 1 Register
//	0x944 32  MACL3A11R                           Layer3 address 1 filter 1 register
//	0x948 32  MACL3A21R                           Layer3 address 2 filter 1 Register
//	0x94C 32  MACL3A31R                           Layer3 address 3 filter 1 register
//	0xAE0 32  MACARPAR                            ARP address register
//	0xB00 32  MACTSCR                             Timestamp control Register
//	0xB04 32  MACSSIR                             Sub-second increment register
//	0xB08 32  MACSTSR                             System time seconds register
//	0xB0C 32  MACSTNR                             System time nanoseconds register
//	0xB10 32  MACSTSUR                            System time seconds update register
//	0xB14 32  MACSTNUR                            System time nanoseconds update register
//	0xB18 32  MACTSAR                             Timestamp addend register
//	0xB20 32  MACTSSR                             Timestamp status register
//	0xB30 32  MACTxTSSNR                          Tx timestamp status nanoseconds register
//	0xB34 32  MACTxTSSSR                          Tx timestamp status seconds register
//	0xB40 32  MACACR                              Auxiliary control register
//	0xB48 32  MACATSNR                            Auxiliary timestamp nanoseconds register
//	0xB4C 32  MACATSSR                            Auxiliary timestamp seconds register
//	0xB50 32  MACTSIACR                           Timestamp Ingress asymmetric correction register
//	0xB54 32  MACTSEACR                           Timestamp Egress asymmetric correction register
//	0xB58 32  MACTSICNR                           Timestamp Ingress correction nanosecond register
//	0xB5C 32  MACTSECNR                           Timestamp Egress correction nanosecond register
//	0xB70 32  MACPPSCR                            PPS control register
//	0xB80 32  MACPPSTTSR                          PPS target time seconds register
//	0xB84 32  MACPPSTTNR                          PPS target time nanoseconds register
//	0xB88 32  MACPPSIR                            PPS interval register
//	0xB8C 32  MACPPSWR                            PPS width register
//	0xBC0 32  MACPOCR                             PTP Offload control register
//	0xBC4 32  MACSPI0R                            PTP Source Port Identity 0 Register
//	0xBC8 32  MACSPI1R                            PTP Source port identity 1 register
//	0xBCC 32  MACSPI2R                            PTP Source port identity 2 register
//	0xBD0 32  MACLMIR                             Log message interval register
//
// Import:
//
//	github.com/embeddedgo/stm32/p/mmap
package ethernet_mac

const (
	RE     MACCR = 0x01 << 0  //+ Receiver Enable
	TE     MACCR = 0x01 << 1  //+ TE
	PRELEN MACCR = 0x03 << 2  //+ PRELEN
	DC     MACCR = 0x01 << 4  //+ DC
	BL     MACCR = 0x03 << 5  //+ BL
	DR     MACCR = 0x01 << 8  //+ DR
	DCRS   MACCR = 0x01 << 9  //+ DCRS
	DO     MACCR = 0x01 << 10 //+ DO
	ECRSFD MACCR = 0x01 << 11 //+ ECRSFD
	LM     MACCR = 0x01 << 12 //+ LM
	DM     MACCR = 0x01 << 13 //+ DM
	FES    MACCR = 0x01 << 14 //+ FES
	JE     MACCR = 0x01 << 16 //+ JE
	JD     MACCR = 0x01 << 17 //+ JD
	WD     MACCR = 0x01 << 19 //+ WD
	ACS    MACCR = 0x01 << 20 //+ ACS
	CST    MACCR = 0x01 << 21 //+ CST
	S2KP   MACCR = 0x01 << 22 //+ S2KP
	GPSLCE MACCR = 0x01 << 23 //+ GPSLCE
	IPG    MACCR = 0x07 << 24 //+ IPG
	IPC    MACCR = 0x01 << 27 //+ IPC
	SARC   MACCR = 0x07 << 28 //+ SARC
	ARPEN  MACCR = 0x01 << 31 //+ ARPEN
)

const (
	REn     = 0
	TEn     = 1
	PRELENn = 2
	DCn     = 4
	BLn     = 5
	DRn     = 8
	DCRSn   = 9
	DOn     = 10
	ECRSFDn = 11
	LMn     = 12
	DMn     = 13
	FESn    = 14
	JEn     = 16
	JDn     = 17
	WDn     = 19
	ACSn    = 20
	CSTn    = 21
	S2KPn   = 22
	GPSLCEn = 23
	IPGn    = 24
	IPCn    = 27
	SARCn   = 28
	ARPENn  = 31
)

const (
	GPSL   MACECR = 0x3FFF << 0 //+ GPSL
	DCRCC  MACECR = 0x01 << 16  //+ DCRCC
	SPEN   MACECR = 0x01 << 17  //+ SPEN
	USP    MACECR = 0x01 << 18  //+ USP
	EIPGEN MACECR = 0x01 << 24  //+ EIPGEN
	EIPG   MACECR = 0x1F << 25  //+ EIPG
)

const (
	GPSLn   = 0
	DCRCCn  = 16
	SPENn   = 17
	USPn    = 18
	EIPGENn = 24
	EIPGn   = 25
)

const (
	PR   MACPFR = 0x01 << 0  //+ PR
	HUC  MACPFR = 0x01 << 1  //+ HUC
	HMC  MACPFR = 0x01 << 2  //+ HMC
	DAIF MACPFR = 0x01 << 3  //+ DAIF
	PM   MACPFR = 0x01 << 4  //+ PM
	DBF  MACPFR = 0x01 << 5  //+ DBF
	PCF  MACPFR = 0x03 << 6  //+ PCF
	SAIF MACPFR = 0x01 << 8  //+ SAIF
	SAF  MACPFR = 0x01 << 9  //+ SAF
	HPF  MACPFR = 0x01 << 10 //+ HPF
	VTFE MACPFR = 0x01 << 16 //+ VTFE
	IPFE MACPFR = 0x01 << 20 //+ IPFE
	DNTU MACPFR = 0x01 << 21 //+ DNTU
	RA   MACPFR = 0x01 << 31 //+ RA
)

const (
	PRn   = 0
	HUCn  = 1
	HMCn  = 2
	DAIFn = 3
	PMn   = 4
	DBFn  = 5
	PCFn  = 6
	SAIFn = 8
	SAFn  = 9
	HPFn  = 10
	VTFEn = 16
	IPFEn = 20
	DNTUn = 21
	RAn   = 31
)

const (
	WTO MACWTR = 0x0F << 0 //+ WTO
	PWE MACWTR = 0x01 << 8 //+ PWE
)

const (
	WTOn = 0
	PWEn = 8
)

const (
	HT31T0 MACHT0R = 0xFFFFFFFF << 0 //+ HT31T0
)

const (
	HT31T0n = 0
)

const (
	HT63T32 MACHT1R = 0xFFFFFFFF << 0 //+ HT63T32
)

const (
	HT63T32n = 0
)

const (
	VL      MACVTR = 0xFFFF << 0 //+ VL
	ETV     MACVTR = 0x01 << 16  //+ ETV
	VTIM    MACVTR = 0x01 << 17  //+ VTIM
	ESVL    MACVTR = 0x01 << 18  //+ ESVL
	ERSVLM  MACVTR = 0x01 << 19  //+ ERSVLM
	DOVLTC  MACVTR = 0x01 << 20  //+ DOVLTC
	EVLS    MACVTR = 0x03 << 21  //+ EVLS
	EVLRXS  MACVTR = 0x01 << 24  //+ EVLRXS
	VTHM    MACVTR = 0x01 << 25  //+ VTHM
	EDVLP   MACVTR = 0x01 << 26  //+ EDVLP
	ERIVLT  MACVTR = 0x01 << 27  //+ ERIVLT
	EIVLS   MACVTR = 0x03 << 28  //+ EIVLS
	EIVLRXS MACVTR = 0x01 << 31  //+ EIVLRXS
)

const (
	VLn      = 0
	ETVn     = 16
	VTIMn    = 17
	ESVLn    = 18
	ERSVLMn  = 19
	DOVLTCn  = 20
	EVLSn    = 21
	EVLRXSn  = 24
	VTHMn    = 25
	EDVLPn   = 26
	ERIVLTn  = 27
	EIVLSn   = 28
	EIVLRXSn = 31
)

const (
	VLHT MACVHTR = 0xFFFF << 0 //+ VLHT
)

const (
	VLHTn = 0
)

const (
	VLT  MACVIR = 0xFFFF << 0 //+ VLT
	VLC  MACVIR = 0x03 << 16  //+ VLC
	VLP  MACVIR = 0x01 << 18  //+ VLP
	CSVL MACVIR = 0x01 << 19  //+ CSVL
	VLTI MACVIR = 0x01 << 20  //+ VLTI
)

const (
	VLTn  = 0
	VLCn  = 16
	VLPn  = 18
	CSVLn = 19
	VLTIn = 20
)

const (
	VLT  MACIVIR = 0xFFFF << 0 //+ VLT
	VLC  MACIVIR = 0x03 << 16  //+ VLC
	VLP  MACIVIR = 0x01 << 18  //+ VLP
	CSVL MACIVIR = 0x01 << 19  //+ CSVL
	VLTI MACIVIR = 0x01 << 20  //+ VLTI
)

const (
	VLTn  = 0
	VLCn  = 16
	VLPn  = 18
	CSVLn = 19
	VLTIn = 20
)

const (
	FCB_BPA MACQTxFCR = 0x01 << 0    //+ FCB_BPA
	TFE     MACQTxFCR = 0x01 << 1    //+ TFE
	PLT     MACQTxFCR = 0x07 << 4    //+ PLT
	DZPQ    MACQTxFCR = 0x01 << 7    //+ DZPQ
	PT      MACQTxFCR = 0xFFFF << 16 //+ PT
)

const (
	FCB_BPAn = 0
	TFEn     = 1
	PLTn     = 4
	DZPQn    = 7
	PTn      = 16
)

const (
	RFE MACRxFCR = 0x01 << 0 //+ RFE
	UP  MACRxFCR = 0x01 << 1 //+ UP
)

const (
	RFEn = 0
	UPn  = 1
)

const (
	PHYIS   MACISR = 0x01 << 3  //+ PHYIS
	PMTIS   MACISR = 0x01 << 4  //+ PMTIS
	LPIIS   MACISR = 0x01 << 5  //+ LPIIS
	MMCIS   MACISR = 0x01 << 8  //+ MMCIS
	MMCRXIS MACISR = 0x01 << 9  //+ MMCRXIS
	MMCTXIS MACISR = 0x01 << 10 //+ MMCTXIS
	TSIS    MACISR = 0x01 << 12 //+ TSIS
	TXSTSIS MACISR = 0x01 << 13 //+ TXSTSIS
	RXSTSIS MACISR = 0x01 << 14 //+ RXSTSIS
)

const (
	PHYISn   = 3
	PMTISn   = 4
	LPIISn   = 5
	MMCISn   = 8
	MMCRXISn = 9
	MMCTXISn = 10
	TSISn    = 12
	TXSTSISn = 13
	RXSTSISn = 14
)

const (
	PHYIE   MACIER = 0x01 << 3  //+ PHYIE
	PMTIE   MACIER = 0x01 << 4  //+ PMTIE
	LPIIE   MACIER = 0x01 << 5  //+ LPIIE
	TSIE    MACIER = 0x01 << 12 //+ TSIE
	TXSTSIE MACIER = 0x01 << 13 //+ TXSTSIE
	RXSTSIE MACIER = 0x01 << 14 //+ RXSTSIE
)

const (
	PHYIEn   = 3
	PMTIEn   = 4
	LPIIEn   = 5
	TSIEn    = 12
	TXSTSIEn = 13
	RXSTSIEn = 14
)

const (
	TJT   MACRxTxSR = 0x01 << 0 //+ TJT
	NCARR MACRxTxSR = 0x01 << 1 //+ NCARR
	LCARR MACRxTxSR = 0x01 << 2 //+ LCARR
	EXDEF MACRxTxSR = 0x01 << 3 //+ EXDEF
	LCOL  MACRxTxSR = 0x01 << 4 //+ LCOL
	EXCOL MACRxTxSR = 0x01 << 5 //+ LCOL
	RWT   MACRxTxSR = 0x01 << 8 //+ RWT
)

const (
	TJTn   = 0
	NCARRn = 1
	LCARRn = 2
	EXDEFn = 3
	LCOLn  = 4
	EXCOLn = 5
	RWTn   = 8
)

const (
	PWRDWN     MACPCSR = 0x01 << 0  //+ PWRDWN
	MGKPKTEN   MACPCSR = 0x01 << 1  //+ MGKPKTEN
	RWKPKTEN   MACPCSR = 0x01 << 2  //+ RWKPKTEN
	MGKPRCVD   MACPCSR = 0x01 << 5  //+ MGKPRCVD
	RWKPRCVD   MACPCSR = 0x01 << 6  //+ RWKPRCVD
	GLBLUCAST  MACPCSR = 0x01 << 9  //+ GLBLUCAST
	RWKPFE     MACPCSR = 0x01 << 10 //+ RWKPFE
	RWKPTR     MACPCSR = 0x1F << 24 //+ RWKPTR
	RWKFILTRST MACPCSR = 0x01 << 31 //+ RWKFILTRST
)

const (
	PWRDWNn     = 0
	MGKPKTENn   = 1
	RWKPKTENn   = 2
	MGKPRCVDn   = 5
	RWKPRCVDn   = 6
	GLBLUCASTn  = 9
	RWKPFEn     = 10
	RWKPTRn     = 24
	RWKFILTRSTn = 31
)

const (
	MACRWKPFR MACRWKPFR = 0xFFFFFFFF << 0 //+ MACRWKPFR
)

const (
	MACRWKPFRn = 0
)

const (
	TLPIEN MACLCSR = 0x01 << 0  //+ TLPIEN
	TLPIEX MACLCSR = 0x01 << 1  //+ TLPIEX
	RLPIEN MACLCSR = 0x01 << 2  //+ RLPIEN
	RLPIEX MACLCSR = 0x01 << 3  //+ RLPIEX
	TLPIST MACLCSR = 0x01 << 8  //+ TLPIST
	RLPIST MACLCSR = 0x01 << 9  //+ RLPIST
	LPIEN  MACLCSR = 0x01 << 16 //+ LPIEN
	PLS    MACLCSR = 0x01 << 17 //+ PLS
	PLSEN  MACLCSR = 0x01 << 18 //+ PLSEN
	LPITXA MACLCSR = 0x01 << 19 //+ LPITXA
	LPITE  MACLCSR = 0x01 << 20 //+ LPITE
)

const (
	TLPIENn = 0
	TLPIEXn = 1
	RLPIENn = 2
	RLPIEXn = 3
	TLPISTn = 8
	RLPISTn = 9
	LPIENn  = 16
	PLSn    = 17
	PLSENn  = 18
	LPITXAn = 19
	LPITEn  = 20
)

const (
	TWT MACLTCR = 0xFFFF << 0 //+ TWT
	LST MACLTCR = 0x3FF << 16 //+ LST
)

const (
	TWTn = 0
	LSTn = 16
)

const (
	LPIET MACLETR = 0x1FFFF << 0 //+ LPIET
)

const (
	LPIETn = 0
)

const (
	TIC_1US_CNTR MAC1USTCR = 0xFFF << 0 //+ TIC_1US_CNTR
)

const (
	TIC_1US_CNTRn = 0
)

const (
	SNPSVER MACVR = 0xFF << 0 //+ SNPSVER
	USERVER MACVR = 0xFF << 8 //+ USERVER
)

const (
	SNPSVERn = 0
	USERVERn = 8
)

const (
	RPESTS   MACDR = 0x01 << 0  //+ RPESTS
	RFCFCSTS MACDR = 0x03 << 1  //+ RFCFCSTS
	TPESTS   MACDR = 0x01 << 16 //+ TPESTS
	TFCSTS   MACDR = 0x03 << 17 //+ TFCSTS
)

const (
	RPESTSn   = 0
	RFCFCSTSn = 1
	TPESTSn   = 16
	TFCSTSn   = 17
)

const (
	RXFIFOSIZE MACHWF1R = 0x1F << 0  //+ RXFIFOSIZE
	TXFIFOSIZE MACHWF1R = 0x1F << 6  //+ TXFIFOSIZE
	OSTEN      MACHWF1R = 0x01 << 11 //+ OSTEN
	PTOEN      MACHWF1R = 0x01 << 12 //+ PTOEN
	ADVTHWORD  MACHWF1R = 0x01 << 13 //+ ADVTHWORD
	ADDR64     MACHWF1R = 0x03 << 14 //+ ADDR64
	DCBEN      MACHWF1R = 0x01 << 16 //+ DCBEN
	SPHEN      MACHWF1R = 0x01 << 17 //+ SPHEN
	TSOEN      MACHWF1R = 0x01 << 18 //+ TSOEN
	DBGMEMA    MACHWF1R = 0x01 << 19 //+ DBGMEMA
	AVSEL      MACHWF1R = 0x01 << 20 //+ AVSEL
	HASHTBLSZ  MACHWF1R = 0x03 << 24 //+ HASHTBLSZ
	L3L4FNUM   MACHWF1R = 0x0F << 27 //+ L3L4FNUM
)

const (
	RXFIFOSIZEn = 0
	TXFIFOSIZEn = 6
	OSTENn      = 11
	PTOENn      = 12
	ADVTHWORDn  = 13
	ADDR64n     = 14
	DCBENn      = 16
	SPHENn      = 17
	TSOENn      = 18
	DBGMEMAn    = 19
	AVSELn      = 20
	HASHTBLSZn  = 24
	L3L4FNUMn   = 27
)

const (
	RXQCNT     MACHWF2R = 0x0F << 0  //+ RXQCNT
	TXQCNT     MACHWF2R = 0x0F << 6  //+ TXQCNT
	RXCHCNT    MACHWF2R = 0x0F << 12 //+ RXCHCNT
	TXCHCNT    MACHWF2R = 0x0F << 18 //+ TXCHCNT
	PPSOUTNUM  MACHWF2R = 0x07 << 24 //+ PPSOUTNUM
	AUXSNAPNUM MACHWF2R = 0x07 << 28 //+ AUXSNAPNUM
)

const (
	RXQCNTn     = 0
	TXQCNTn     = 6
	RXCHCNTn    = 12
	TXCHCNTn    = 18
	PPSOUTNUMn  = 24
	AUXSNAPNUMn = 28
)

const (
	MB   MACMDIOAR = 0x01 << 0  //+ MB
	C45E MACMDIOAR = 0x01 << 1  //+ C45E
	GOC  MACMDIOAR = 0x03 << 2  //+ GOC
	SKAP MACMDIOAR = 0x01 << 4  //+ SKAP
	CR   MACMDIOAR = 0x0F << 8  //+ CR
	NTC  MACMDIOAR = 0x07 << 12 //+ NTC
	RDA  MACMDIOAR = 0x1F << 16 //+ RDA
	PA   MACMDIOAR = 0x1F << 21 //+ PA
	BTB  MACMDIOAR = 0x01 << 26 //+ BTB
	PSE  MACMDIOAR = 0x01 << 27 //+ PSE
)

const (
	MBn   = 0
	C45En = 1
	GOCn  = 2
	SKAPn = 4
	CRn   = 8
	NTCn  = 12
	RDAn  = 16
	PAn   = 21
	BTBn  = 26
	PSEn  = 27
)

const (
	MD MACMDIODR = 0xFFFF << 0  //+ MD
	RA MACMDIODR = 0xFFFF << 16 //+ RA
)

const (
	MDn = 0
	RAn = 16
)

const (
	ADDRHI MACA0HR = 0xFFFF << 0 //+ ADDRHI
	AE     MACA0HR = 0x01 << 31  //+ AE
)

const (
	ADDRHIn = 0
	AEn     = 31
)

const (
	ADDRLO MACA0LR = 0xFFFFFFFF << 0 //+ ADDRLO
)

const (
	ADDRLOn = 0
)

const (
	ADDRHI MACA1HR = 0xFFFF << 0 //+ ADDRHI
	MBC    MACA1HR = 0x3F << 24  //+ MBC
	SA     MACA1HR = 0x01 << 30  //+ SA
	AE     MACA1HR = 0x01 << 31  //+ AE
)

const (
	ADDRHIn = 0
	MBCn    = 24
	SAn     = 30
	AEn     = 31
)

const (
	ADDRLO MACA1LR = 0xFFFFFFFF << 0 //+ ADDRLO
)

const (
	ADDRLOn = 0
)

const (
	ADDRHI MACA2HR = 0xFFFF << 0 //+ ADDRHI
	MBC    MACA2HR = 0x3F << 24  //+ MBC
	SA     MACA2HR = 0x01 << 30  //+ SA
	AE     MACA2HR = 0x01 << 31  //+ AE
)

const (
	ADDRHIn = 0
	MBCn    = 24
	SAn     = 30
	AEn     = 31
)

const (
	ADDRLO MACA2LR = 0xFFFFFFFF << 0 //+ ADDRLO
)

const (
	ADDRLOn = 0
)

const (
	ADDRHI MACA3HR = 0xFFFF << 0 //+ ADDRHI
	MBC    MACA3HR = 0x3F << 24  //+ MBC
	SA     MACA3HR = 0x01 << 30  //+ SA
	AE     MACA3HR = 0x01 << 31  //+ AE
)

const (
	ADDRHIn = 0
	MBCn    = 24
	SAn     = 30
	AEn     = 31
)

const (
	ADDRLO MACA3LR = 0xFFFFFFFF << 0 //+ ADDRLO
)

const (
	ADDRLOn = 0
)

const (
	CNTRST     MMC_CONTROL = 0x01 << 0 //+ CNTRST
	CNTSTOPRO  MMC_CONTROL = 0x01 << 1 //+ CNTSTOPRO
	RSTONRD    MMC_CONTROL = 0x01 << 2 //+ RSTONRD
	CNTFREEZ   MMC_CONTROL = 0x01 << 3 //+ CNTFREEZ
	CNTPRST    MMC_CONTROL = 0x01 << 4 //+ CNTPRST
	CNTPRSTLVL MMC_CONTROL = 0x01 << 5 //+ CNTPRSTLVL
	UCDBC      MMC_CONTROL = 0x01 << 8 //+ UCDBC
)

const (
	CNTRSTn     = 0
	CNTSTOPROn  = 1
	RSTONRDn    = 2
	CNTFREEZn   = 3
	CNTPRSTn    = 4
	CNTPRSTLVLn = 5
	UCDBCn      = 8
)

const (
	RXCRCERPIS  MMC_RX_INTERRUPT = 0x01 << 5  //+ RXCRCERPIS
	RXALGNERPIS MMC_RX_INTERRUPT = 0x01 << 6  //+ RXALGNERPIS
	RXUCGPIS    MMC_RX_INTERRUPT = 0x01 << 17 //+ RXUCGPIS
	RXLPIUSCIS  MMC_RX_INTERRUPT = 0x01 << 26 //+ RXLPIUSCIS
	RXLPITRCIS  MMC_RX_INTERRUPT = 0x01 << 27 //+ RXLPITRCIS
)

const (
	RXCRCERPISn  = 5
	RXALGNERPISn = 6
	RXUCGPISn    = 17
	RXLPIUSCISn  = 26
	RXLPITRCISn  = 27
)

const (
	TXSCOLGPIS MMC_TX_INTERRUPT = 0x01 << 14 //+ TXSCOLGPIS
	TXMCOLGPIS MMC_TX_INTERRUPT = 0x01 << 15 //+ TXMCOLGPIS
	TXGPKTIS   MMC_TX_INTERRUPT = 0x01 << 21 //+ TXGPKTIS
	TXLPIUSCIS MMC_TX_INTERRUPT = 0x01 << 26 //+ TXLPIUSCIS
	TXLPITRCIS MMC_TX_INTERRUPT = 0x01 << 27 //+ TXLPITRCIS
)

const (
	TXSCOLGPISn = 14
	TXMCOLGPISn = 15
	TXGPKTISn   = 21
	TXLPIUSCISn = 26
	TXLPITRCISn = 27
)

const (
	RXCRCERPIM  MMC_RX_INTERRUPT_MASK = 0x01 << 5  //+ RXCRCERPIM
	RXALGNERPIM MMC_RX_INTERRUPT_MASK = 0x01 << 6  //+ RXALGNERPIM
	RXUCGPIM    MMC_RX_INTERRUPT_MASK = 0x01 << 17 //+ RXUCGPIM
	RXLPIUSCIM  MMC_RX_INTERRUPT_MASK = 0x01 << 26 //+ RXLPIUSCIM
	RXLPITRCIM  MMC_RX_INTERRUPT_MASK = 0x01 << 27 //+ RXLPITRCIM
)

const (
	RXCRCERPIMn  = 5
	RXALGNERPIMn = 6
	RXUCGPIMn    = 17
	RXLPIUSCIMn  = 26
	RXLPITRCIMn  = 27
)

const (
	TXSCOLGPIM MMC_TX_INTERRUPT_MASK = 0x01 << 14 //+ TXSCOLGPIM
	TXMCOLGPIM MMC_TX_INTERRUPT_MASK = 0x01 << 15 //+ TXMCOLGPIM
	TXGPKTIM   MMC_TX_INTERRUPT_MASK = 0x01 << 21 //+ TXGPKTIM
	TXLPIUSCIM MMC_TX_INTERRUPT_MASK = 0x01 << 26 //+ TXLPIUSCIM
	TXLPITRCIM MMC_TX_INTERRUPT_MASK = 0x01 << 27 //+ TXLPITRCIM
)

const (
	TXSCOLGPIMn = 14
	TXMCOLGPIMn = 15
	TXGPKTIMn   = 21
	TXLPIUSCIMn = 26
	TXLPITRCIMn = 27
)

const (
	TXSNGLCOLG TX_SINGLE_COLLISION_GOOD_PACKETS = 0xFFFFFFFF << 0 //+ TXSNGLCOLG
)

const (
	TXSNGLCOLGn = 0
)

const (
	TXMULTCOLG TX_MULTIPLE_COLLISION_GOOD_PACKETS = 0xFFFFFFFF << 0 //+ TXMULTCOLG
)

const (
	TXMULTCOLGn = 0
)

const (
	TXPKTG TX_PACKET_COUNT_GOOD = 0xFFFFFFFF << 0 //+ TXPKTG
)

const (
	TXPKTGn = 0
)

const (
	RXCRCERR RX_CRC_ERROR_PACKETS = 0xFFFFFFFF << 0 //+ RXCRCERR
)

const (
	RXCRCERRn = 0
)

const (
	RXALGNERR RX_ALIGNMENT_ERROR_PACKETS = 0xFFFFFFFF << 0 //+ RXALGNERR
)

const (
	RXALGNERRn = 0
)

const (
	RXUCASTG RX_UNICAST_PACKETS_GOOD = 0xFFFFFFFF << 0 //+ RXUCASTG
)

const (
	RXUCASTGn = 0
)

const (
	TXLPIUSC TX_LPI_USEC_CNTR = 0xFFFFFFFF << 0 //+ TXLPIUSC
)

const (
	TXLPIUSCn = 0
)

const (
	TXLPITRC TX_LPI_TRAN_CNTR = 0xFFFFFFFF << 0 //+ TXLPITRC
)

const (
	TXLPITRCn = 0
)

const (
	RXLPIUSC RX_LPI_USEC_CNTR = 0xFFFFFFFF << 0 //+ RXLPIUSC
)

const (
	RXLPIUSCn = 0
)

const (
	RXLPITRC RX_LPI_TRAN_CNTR = 0xFFFFFFFF << 0 //+ RXLPITRC
)

const (
	RXLPITRCn = 0
)

const (
	L3PEN0  MACL3L4C0R = 0x01 << 0  //+ L3PEN0
	L3SAM0  MACL3L4C0R = 0x01 << 2  //+ L3SAM0
	L3SAIM0 MACL3L4C0R = 0x01 << 3  //+ L3SAIM0
	L3DAM0  MACL3L4C0R = 0x01 << 4  //+ L3DAM0
	L3DAIM0 MACL3L4C0R = 0x01 << 5  //+ L3DAIM0
	L3HSBM0 MACL3L4C0R = 0x1F << 6  //+ L3HSBM0
	L3HDBM0 MACL3L4C0R = 0x1F << 11 //+ L3HDBM0
	L4PEN0  MACL3L4C0R = 0x01 << 16 //+ L4PEN0
	L4SPM0  MACL3L4C0R = 0x01 << 18 //+ L4SPM0
	L4SPIM0 MACL3L4C0R = 0x01 << 19 //+ L4SPIM0
	L4DPM0  MACL3L4C0R = 0x01 << 20 //+ L4DPM0
	L4DPIM0 MACL3L4C0R = 0x01 << 21 //+ L4DPIM0
)

const (
	L3PEN0n  = 0
	L3SAM0n  = 2
	L3SAIM0n = 3
	L3DAM0n  = 4
	L3DAIM0n = 5
	L3HSBM0n = 6
	L3HDBM0n = 11
	L4PEN0n  = 16
	L4SPM0n  = 18
	L4SPIM0n = 19
	L4DPM0n  = 20
	L4DPIM0n = 21
)

const (
	L4SP0 MACL4A0R = 0xFFFF << 0  //+ L4SP0
	L4DP0 MACL4A0R = 0xFFFF << 16 //+ L4DP0
)

const (
	L4SP0n = 0
	L4DP0n = 16
)

const (
	L3A00 MACL3A00R = 0xFFFFFFFF << 0 //+ L3A00
)

const (
	L3A00n = 0
)

const (
	L3A10 MACL3A10R = 0xFFFFFFFF << 0 //+ L3A10
)

const (
	L3A10n = 0
)

const (
	L3A20 MACL3A20 = 0xFFFFFFFF << 0 //+ L3A20
)

const (
	L3A20n = 0
)

const (
	L3A30 MACL3A30 = 0xFFFFFFFF << 0 //+ L3A30
)

const (
	L3A30n = 0
)

const (
	L3PEN1  MACL3L4C1R = 0x01 << 0  //+ L3PEN1
	L3SAM1  MACL3L4C1R = 0x01 << 2  //+ L3SAM1
	L3SAIM1 MACL3L4C1R = 0x01 << 3  //+ L3SAIM1
	L3DAM1  MACL3L4C1R = 0x01 << 4  //+ L3DAM1
	L3DAIM1 MACL3L4C1R = 0x01 << 5  //+ L3DAIM1
	L3HSBM1 MACL3L4C1R = 0x1F << 6  //+ L3HSBM1
	L3HDBM1 MACL3L4C1R = 0x1F << 11 //+ L3HDBM1
	L4PEN1  MACL3L4C1R = 0x01 << 16 //+ L4PEN1
	L4SPM1  MACL3L4C1R = 0x01 << 18 //+ L4SPM1
	L4SPIM1 MACL3L4C1R = 0x01 << 19 //+ L4SPIM1
	L4DPM1  MACL3L4C1R = 0x01 << 20 //+ L4DPM1
	L4DPIM1 MACL3L4C1R = 0x01 << 21 //+ L4DPIM1
)

const (
	L3PEN1n  = 0
	L3SAM1n  = 2
	L3SAIM1n = 3
	L3DAM1n  = 4
	L3DAIM1n = 5
	L3HSBM1n = 6
	L3HDBM1n = 11
	L4PEN1n  = 16
	L4SPM1n  = 18
	L4SPIM1n = 19
	L4DPM1n  = 20
	L4DPIM1n = 21
)

const (
	L4SP1 MACL4A1R = 0xFFFF << 0  //+ L4SP1
	L4DP1 MACL4A1R = 0xFFFF << 16 //+ L4DP1
)

const (
	L4SP1n = 0
	L4DP1n = 16
)

const (
	L3A01 MACL3A01R = 0xFFFFFFFF << 0 //+ L3A01
)

const (
	L3A01n = 0
)

const (
	L3A11 MACL3A11R = 0xFFFFFFFF << 0 //+ L3A11
)

const (
	L3A11n = 0
)

const (
	L3A21 MACL3A21R = 0xFFFFFFFF << 0 //+ L3A21
)

const (
	L3A21n = 0
)

const (
	L3A31 MACL3A31R = 0xFFFFFFFF << 0 //+ L3A31
)

const (
	L3A31n = 0
)

const (
	ARPPA MACARPAR = 0xFFFFFFFF << 0 //+ ARPPA
)

const (
	ARPPAn = 0
)

const (
	TSENA       MACTSCR = 0x01 << 0  //+ TSENA
	TSCFUPDT    MACTSCR = 0x01 << 1  //+ TSCFUPDT
	TSINIT      MACTSCR = 0x01 << 2  //+ TSINIT
	TSUPDT      MACTSCR = 0x01 << 3  //+ TSUPDT
	TSADDREG    MACTSCR = 0x01 << 5  //+ TSADDREG
	TSENALL     MACTSCR = 0x01 << 8  //+ TSENALL
	TSCTRLSSR   MACTSCR = 0x01 << 9  //+ TSCTRLSSR
	TSVER2ENA   MACTSCR = 0x01 << 10 //+ TSVER2ENA
	TSIPENA     MACTSCR = 0x01 << 11 //+ TSIPENA
	TSIPV6ENA   MACTSCR = 0x01 << 12 //+ TSIPV6ENA
	TSIPV4ENA   MACTSCR = 0x01 << 13 //+ TSIPV4ENA
	TSEVNTENA   MACTSCR = 0x01 << 14 //+ TSEVNTENA
	TSMSTRENA   MACTSCR = 0x01 << 15 //+ TSMSTRENA
	SNAPTYPSEL  MACTSCR = 0x03 << 16 //+ SNAPTYPSEL
	TSENMACADDR MACTSCR = 0x01 << 18 //+ TSENMACADDR
	CSC         MACTSCR = 0x01 << 19 //+ CSC
	TXTSSTSM    MACTSCR = 0x01 << 24 //+ TXTSSTSM
)

const (
	TSENAn       = 0
	TSCFUPDTn    = 1
	TSINITn      = 2
	TSUPDTn      = 3
	TSADDREGn    = 5
	TSENALLn     = 8
	TSCTRLSSRn   = 9
	TSVER2ENAn   = 10
	TSIPENAn     = 11
	TSIPV6ENAn   = 12
	TSIPV4ENAn   = 13
	TSEVNTENAn   = 14
	TSMSTRENAn   = 15
	SNAPTYPSELn  = 16
	TSENMACADDRn = 18
	CSCn         = 19
	TXTSSTSMn    = 24
)

const (
	SNSINC MACSSIR = 0xFF << 8  //+ SNSINC
	SSINC  MACSSIR = 0xFF << 16 //+ SSINC
)

const (
	SNSINCn = 8
	SSINCn  = 16
)

const (
	TSS MACSTSR = 0xFFFFFFFF << 0 //+ TSS
)

const (
	TSSn = 0
)

const (
	TSSS MACSTNR = 0x7FFFFFFF << 0 //+ TSSS
)

const (
	TSSSn = 0
)

const (
	TSS MACSTSUR = 0xFFFFFFFF << 0 //+ TSS
)

const (
	TSSn = 0
)

const (
	TSSS   MACSTNUR = 0x7FFFFFFF << 0 //+ TSSS
	ADDSUB MACSTNUR = 0x01 << 31      //+ ADDSUB
)

const (
	TSSSn   = 0
	ADDSUBn = 31
)

const (
	TSAR MACTSAR = 0xFFFFFFFF << 0 //+ TSAR
)

const (
	TSARn = 0
)

const (
	TSSOVF     MACTSSR = 0x01 << 0  //+ TSSOVF
	TSTARGT0   MACTSSR = 0x01 << 1  //+ TSTARGT0
	AUXTSTRIG  MACTSSR = 0x01 << 2  //+ AUXTSTRIG
	TSTRGTERR0 MACTSSR = 0x01 << 3  //+ TSTRGTERR0
	TXTSSIS    MACTSSR = 0x01 << 15 //+ TXTSSIS
	ATSSTN     MACTSSR = 0x0F << 16 //+ ATSSTN
	ATSSTM     MACTSSR = 0x01 << 24 //+ ATSSTM
	ATSNS      MACTSSR = 0x1F << 25 //+ ATSNS
)

const (
	TSSOVFn     = 0
	TSTARGT0n   = 1
	AUXTSTRIGn  = 2
	TSTRGTERR0n = 3
	TXTSSISn    = 15
	ATSSTNn     = 16
	ATSSTMn     = 24
	ATSNSn      = 25
)

const (
	TXTSSLO  MACTxTSSNR = 0x7FFFFFFF << 0 //+ TXTSSLO
	TXTSSMIS MACTxTSSNR = 0x01 << 31      //+ TXTSSMIS
)

const (
	TXTSSLOn  = 0
	TXTSSMISn = 31
)

const (
	TXTSSHI MACTxTSSSR = 0xFFFFFFFF << 0 //+ TXTSSHI
)

const (
	TXTSSHIn = 0
)

const (
	ATSFC  MACACR = 0x01 << 0 //+ ATSFC
	ATSEN0 MACACR = 0x01 << 4 //+ ATSEN0
	ATSEN1 MACACR = 0x01 << 5 //+ ATSEN1
	ATSEN2 MACACR = 0x01 << 6 //+ ATSEN2
	ATSEN3 MACACR = 0x01 << 7 //+ ATSEN3
)

const (
	ATSFCn  = 0
	ATSEN0n = 4
	ATSEN1n = 5
	ATSEN2n = 6
	ATSEN3n = 7
)

const (
	AUXTSLO MACATSNR = 0x7FFFFFFF << 0 //+ AUXTSLO
)

const (
	AUXTSLOn = 0
)

const (
	AUXTSHI MACATSSR = 0xFFFFFFFF << 0 //+ AUXTSHI
)

const (
	AUXTSHIn = 0
)

const (
	OSTIAC MACTSIACR = 0xFFFFFFFF << 0 //+ OSTIAC
)

const (
	OSTIACn = 0
)

const (
	OSTEAC MACTSEACR = 0xFFFFFFFF << 0 //+ OSTEAC
)

const (
	OSTEACn = 0
)

const (
	TSIC MACTSICNR = 0xFFFFFFFF << 0 //+ TSIC
)

const (
	TSICn = 0
)

const (
	TSEC MACTSECNR = 0xFFFFFFFF << 0 //+ TSEC
)

const (
	TSECn = 0
)

const (
	PPSCTRL     MACPPSCR = 0x0F << 0 //+ PPSCTRL
	PPSEN0      MACPPSCR = 0x01 << 4 //+ PPSEN0
	TRGTMODSEL0 MACPPSCR = 0x03 << 5 //+ TRGTMODSEL0
)

const (
	PPSCTRLn     = 0
	PPSEN0n      = 4
	TRGTMODSEL0n = 5
)

const (
	TSTRH0 MACPPSTTSR = 0x7FFFFFFF << 0 //+ TSTRH0
)

const (
	TSTRH0n = 0
)

const (
	TTSL0     MACPPSTTNR = 0x7FFFFFFF << 0 //+ TTSL0
	TRGTBUSY0 MACPPSTTNR = 0x01 << 31      //+ TRGTBUSY0
)

const (
	TTSL0n     = 0
	TRGTBUSY0n = 31
)

const (
	PPSINT0 MACPPSIR = 0xFFFFFFFF << 0 //+ PPSINT0
)

const (
	PPSINT0n = 0
)

const (
	PPSWIDTH0 MACPPSWR = 0xFFFFFFFF << 0 //+ PPSWIDTH0
)

const (
	PPSWIDTH0n = 0
)

const (
	PTOEN      MACPOCR = 0x01 << 0 //+ PTOEN
	ASYNCEN    MACPOCR = 0x01 << 1 //+ ASYNCEN
	APDREQEN   MACPOCR = 0x01 << 2 //+ APDREQEN
	ASYNCTRIG  MACPOCR = 0x01 << 4 //+ ASYNCTRIG
	APDREQTRIG MACPOCR = 0x01 << 5 //+ APDREQTRIG
	DRRDIS     MACPOCR = 0x01 << 6 //+ DRRDIS
	DN         MACPOCR = 0xFF << 8 //+ DN
)

const (
	PTOENn      = 0
	ASYNCENn    = 1
	APDREQENn   = 2
	ASYNCTRIGn  = 4
	APDREQTRIGn = 5
	DRRDISn     = 6
	DNn         = 8
)

const (
	SPI0 MACSPI0R = 0xFFFFFFFF << 0 //+ SPI0
)

const (
	SPI0n = 0
)

const (
	SPI1 MACSPI1R = 0xFFFFFFFF << 0 //+ SPI1
)

const (
	SPI1n = 0
)

const (
	SPI2 MACSPI2R = 0xFFFF << 0 //+ SPI2
)

const (
	SPI2n = 0
)

const (
	LSI     MACLMIR = 0xFF << 0  //+ LSI
	DRSYNCR MACLMIR = 0x07 << 8  //+ DRSYNCR
	LMPDRI  MACLMIR = 0xFF << 24 //+ LMPDRI
)

const (
	LSIn     = 0
	DRSYNCRn = 8
	LMPDRIn  = 24
)
