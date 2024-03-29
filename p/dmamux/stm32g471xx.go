// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build stm32g471xx

// Package dmamux provides access to the registers of the DMAMUX peripheral.
//
// Instances:
//  DMAMUX  DMAMUX_BASE  AHB1  DMAMUX_OVR
// Registers:
//  0x000 32  C0CR   DMAMux - DMA request line multiplexer channel x control register
//  0x004 32  C1CR   DMAMux - DMA request line multiplexer channel x control register
//  0x008 32  C2CR   DMAMux - DMA request line multiplexer channel x control register
//  0x00C 32  C3CR   DMAMux - DMA request line multiplexer channel x control register
//  0x010 32  C4CR   DMAMux - DMA request line multiplexer channel x control register
//  0x014 32  C5CR   DMAMux - DMA request line multiplexer channel x control register
//  0x018 32  C6CR   DMAMux - DMA request line multiplexer channel x control register
//  0x01C 32  C7CR   DMAMux - DMA request line multiplexer channel x control register
//  0x020 32  C8CR   DMAMux - DMA request line multiplexer channel x control register
//  0x024 32  C9CR   DMAMux - DMA request line multiplexer channel x control register
//  0x028 32  C10CR  DMAMux - DMA request line multiplexer channel x control register
//  0x02C 32  C11CR  DMAMux - DMA request line multiplexer channel x control register
//  0x030 32  C12CR  DMAMux - DMA request line multiplexer channel x control register
//  0x034 32  C13CR  DMAMux - DMA request line multiplexer channel x control register
//  0x038 32  C14CR  DMAMux - DMA request line multiplexer channel x control register
//  0x03C 32  C15CR  DMAMux - DMA request line multiplexer channel x control register
//  0x080 32  CSR    DMAMUX request line multiplexer interrupt channel status register
//  0x084 32  CFR    DMAMUX request line multiplexer interrupt clear flag register
//  0x100 32  RG0CR  DMAMux - DMA request generator channel x control register
//  0x104 32  RG1CR  DMAMux - DMA request generator channel x control register
//  0x108 32  RG2CR  DMAMux - DMA request generator channel x control register
//  0x10C 32  RG3CR  DMAMux - DMA request generator channel x control register
//  0x140 32  RGSR   DMAMux - DMA request generator status register
//  0x144 32  RGCFR  DMAMux - DMA request generator clear flag register
// Import:
//  github.com/embeddedgo/stm32/p/bus
//  github.com/embeddedgo/stm32/p/mmap
package dmamux

const (
	DMAREQ_ID C0CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C0CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C0CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C0CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C0CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C0CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C0CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C1CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C1CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C1CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C1CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C1CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C1CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C1CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C2CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C2CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C2CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C2CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C2CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C2CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C2CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C3CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C3CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C3CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C3CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C3CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C3CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C3CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C4CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C4CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C4CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C4CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C4CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C4CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C4CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C5CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C5CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C5CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C5CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C5CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C5CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C5CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C6CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C6CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C6CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C6CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C6CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C6CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C6CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C7CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C7CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C7CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C7CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C7CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C7CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C7CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C8CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C8CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C8CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C8CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C8CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C8CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C8CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C9CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C9CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C9CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C9CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C9CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C9CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C9CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C10CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C10CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C10CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C10CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C10CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C10CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C10CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C11CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C11CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C11CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C11CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C11CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C11CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C11CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C12CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C12CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C12CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C12CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C12CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C12CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C12CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C13CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C13CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C13CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C13CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C13CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C13CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C13CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C14CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C14CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C14CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C14CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C14CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C14CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C14CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	DMAREQ_ID C15CR = 0x7F << 0  //+ Input DMA request line selected
	SOIE      C15CR = 0x01 << 8  //+ Interrupt enable at synchronization event overrun
	EGE       C15CR = 0x01 << 9  //+ Event generation enable/disable
	SE        C15CR = 0x01 << 16 //+ Synchronous operating mode enable/disable
	SPOL      C15CR = 0x03 << 17 //+ Synchronization event type selector Defines the synchronization event on the selected synchronization input:
	NBREQ     C15CR = 0x1F << 19 //+ Number of DMA requests to forward Defines the number of DMA requests forwarded before output event is generated. In synchronous mode, it also defines the number of DMA requests to forward after a synchronization event, then stop forwarding. The actual number of DMA requests forwarded is NBREQ+1. Note: This field can only be written when both SE and EGE bits are reset.
	SYNC_ID   C15CR = 0x1F << 24 //+ Synchronization input selected
)

const (
	DMAREQ_IDn = 0
	SOIEn      = 8
	EGEn       = 9
	SEn        = 16
	SPOLn      = 17
	NBREQn     = 19
	SYNC_IDn   = 24
)

const (
	SOF CSR = 0xFFFF << 0 //+ Synchronization overrun event flag
)

const (
	SOFn = 0
)

const (
	CSOF CFR = 0xFFFF << 0 //+ Clear synchronization overrun event flag
)

const (
	CSOFn = 0
)

const (
	SIG_ID RG0CR = 0x1F << 0  //+ DMA request trigger input selected
	OIE    RG0CR = 0x01 << 8  //+ Interrupt enable at trigger event overrun
	GE     RG0CR = 0x01 << 16 //+ DMA request generator channel enable/disable
	GPOL   RG0CR = 0x03 << 17 //+ DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
	GNBREQ RG0CR = 0x1F << 19 //+ Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
)

const (
	SIG_IDn = 0
	OIEn    = 8
	GEn     = 16
	GPOLn   = 17
	GNBREQn = 19
)

const (
	SIG_ID RG1CR = 0x1F << 0  //+ DMA request trigger input selected
	OIE    RG1CR = 0x01 << 8  //+ Interrupt enable at trigger event overrun
	GE     RG1CR = 0x01 << 16 //+ DMA request generator channel enable/disable
	GPOL   RG1CR = 0x03 << 17 //+ DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
	GNBREQ RG1CR = 0x1F << 19 //+ Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
)

const (
	SIG_IDn = 0
	OIEn    = 8
	GEn     = 16
	GPOLn   = 17
	GNBREQn = 19
)

const (
	SIG_ID RG2CR = 0x1F << 0  //+ DMA request trigger input selected
	OIE    RG2CR = 0x01 << 8  //+ Interrupt enable at trigger event overrun
	GE     RG2CR = 0x01 << 16 //+ DMA request generator channel enable/disable
	GPOL   RG2CR = 0x03 << 17 //+ DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
	GNBREQ RG2CR = 0x1F << 19 //+ Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
)

const (
	SIG_IDn = 0
	OIEn    = 8
	GEn     = 16
	GPOLn   = 17
	GNBREQn = 19
)

const (
	SIG_ID RG3CR = 0x1F << 0  //+ DMA request trigger input selected
	OIE    RG3CR = 0x01 << 8  //+ Interrupt enable at trigger event overrun
	GE     RG3CR = 0x01 << 16 //+ DMA request generator channel enable/disable
	GPOL   RG3CR = 0x03 << 17 //+ DMA request generator trigger event type selection Defines the trigger event on the selected DMA request trigger input
	GNBREQ RG3CR = 0x1F << 19 //+ Number of DMA requests to generate Defines the number of DMA requests generated after a trigger event, then stop generating. The actual number of generated DMA requests is GNBREQ+1. Note: This field can only be written when GE bit is reset.
)

const (
	SIG_IDn = 0
	OIEn    = 8
	GEn     = 16
	GPOLn   = 17
	GNBREQn = 19
)

const (
	OF RGSR = 0x0F << 0 //+ Trigger event overrun flag The flag is set when a trigger event occurs on DMA request generator channel x, while the DMA request generator counter value is lower than GNBREQ. The flag is cleared by writing 1 to the corresponding COFx bit in DMAMUX_RGCFR register.
)

const (
	OFn = 0
)

const (
	COF RGCFR = 0x0F << 0 //+ Clear trigger event overrun flag Upon setting, this bit clears the corresponding overrun flag OFx in the DMAMUX_RGCSR register.
)

const (
	COFn = 0
)
