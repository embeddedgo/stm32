// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build stm32g471xx

package rtc

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	TR       RTR
	DR       RDR
	SSR      RSSR
	ICSR     RICSR
	PRER     RPRER
	WUTR     RWUTR
	CR       RCR
	_        [2]uint32
	WPR      RWPR
	CALR     RCALR
	SHIFTR   RSHIFTR
	TSTR     RTR
	TSDR     RDR
	TSSSR    RSSR
	_        uint32
	ALRMAR   RALRMR
	ALRMASSR RALRMSSR
	ALRMBR   RALRMR
	ALRMBSSR RALRMSSR
	SR       RSR
	MISR     RMISR
	_        uint32
	SCR      RSCR
}

func RTC() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.RTC_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type TR uint32

type RTR struct{ mmio.U32 }

func (r *RTR) LoadBits(mask TR) TR  { return TR(r.U32.LoadBits(uint32(mask))) }
func (r *RTR) StoreBits(mask, b TR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTR) SetBits(mask TR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTR) ClearBits(mask TR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTR) Load() TR             { return TR(r.U32.Load()) }
func (r *RTR) Store(b TR)           { r.U32.Store(uint32(b)) }

type RMTR struct{ mmio.UM32 }

func (rm RMTR) Load() TR   { return TR(rm.UM32.Load()) }
func (rm RMTR) Store(b TR) { rm.UM32.Store(uint32(b)) }

func SU_(p *Periph) RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(SU)}}
}

func ST_(p *Periph) RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(ST)}}
}

func MNU_(p *Periph) RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(MNU)}}
}

func MNT_(p *Periph) RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(MNT)}}
}

func HU_(p *Periph) RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(HU)}}
}

func HT_(p *Periph) RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(HT)}}
}

func PM_(p *Periph) RMTR {
	return RMTR{mmio.UM32{&p.TR.U32, uint32(PM)}}
}

type DR uint32

type RDR struct{ mmio.U32 }

func (r *RDR) LoadBits(mask DR) DR  { return DR(r.U32.LoadBits(uint32(mask))) }
func (r *RDR) StoreBits(mask, b DR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDR) SetBits(mask DR)      { r.U32.SetBits(uint32(mask)) }
func (r *RDR) ClearBits(mask DR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDR) Load() DR             { return DR(r.U32.Load()) }
func (r *RDR) Store(b DR)           { r.U32.Store(uint32(b)) }

type RMDR struct{ mmio.UM32 }

func (rm RMDR) Load() DR   { return DR(rm.UM32.Load()) }
func (rm RMDR) Store(b DR) { rm.UM32.Store(uint32(b)) }

func DU_(p *Periph) RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(DU)}}
}

func DT_(p *Periph) RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(DT)}}
}

func MU_(p *Periph) RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(MU)}}
}

func MT_(p *Periph) RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(MT)}}
}

func WDU_(p *Periph) RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(WDU)}}
}

func YU_(p *Periph) RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(YU)}}
}

func YT_(p *Periph) RMDR {
	return RMDR{mmio.UM32{&p.DR.U32, uint32(YT)}}
}

type SSR uint32

type RSSR struct{ mmio.U32 }

func (r *RSSR) LoadBits(mask SSR) SSR { return SSR(r.U32.LoadBits(uint32(mask))) }
func (r *RSSR) StoreBits(mask, b SSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSSR) SetBits(mask SSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSSR) ClearBits(mask SSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSSR) Load() SSR             { return SSR(r.U32.Load()) }
func (r *RSSR) Store(b SSR)           { r.U32.Store(uint32(b)) }

type RMSSR struct{ mmio.UM32 }

func (rm RMSSR) Load() SSR   { return SSR(rm.UM32.Load()) }
func (rm RMSSR) Store(b SSR) { rm.UM32.Store(uint32(b)) }

func SS_(p *Periph) RMSSR {
	return RMSSR{mmio.UM32{&p.SSR.U32, uint32(SS)}}
}

type ICSR uint32

type RICSR struct{ mmio.U32 }

func (r *RICSR) LoadBits(mask ICSR) ICSR { return ICSR(r.U32.LoadBits(uint32(mask))) }
func (r *RICSR) StoreBits(mask, b ICSR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RICSR) SetBits(mask ICSR)       { r.U32.SetBits(uint32(mask)) }
func (r *RICSR) ClearBits(mask ICSR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RICSR) Load() ICSR              { return ICSR(r.U32.Load()) }
func (r *RICSR) Store(b ICSR)            { r.U32.Store(uint32(b)) }

type RMICSR struct{ mmio.UM32 }

func (rm RMICSR) Load() ICSR   { return ICSR(rm.UM32.Load()) }
func (rm RMICSR) Store(b ICSR) { rm.UM32.Store(uint32(b)) }

func ALRAWF_(p *Periph) RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(ALRAWF)}}
}

func ALRBWF_(p *Periph) RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(ALRBWF)}}
}

func WUTWF_(p *Periph) RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(WUTWF)}}
}

func SHPF_(p *Periph) RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(SHPF)}}
}

func INITS_(p *Periph) RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(INITS)}}
}

func RSF_(p *Periph) RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(RSF)}}
}

func INITF_(p *Periph) RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(INITF)}}
}

func INIT_(p *Periph) RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(INIT)}}
}

func RECALPF_(p *Periph) RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(RECALPF)}}
}

type PRER uint32

type RPRER struct{ mmio.U32 }

func (r *RPRER) LoadBits(mask PRER) PRER { return PRER(r.U32.LoadBits(uint32(mask))) }
func (r *RPRER) StoreBits(mask, b PRER)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPRER) SetBits(mask PRER)       { r.U32.SetBits(uint32(mask)) }
func (r *RPRER) ClearBits(mask PRER)     { r.U32.ClearBits(uint32(mask)) }
func (r *RPRER) Load() PRER              { return PRER(r.U32.Load()) }
func (r *RPRER) Store(b PRER)            { r.U32.Store(uint32(b)) }

type RMPRER struct{ mmio.UM32 }

func (rm RMPRER) Load() PRER   { return PRER(rm.UM32.Load()) }
func (rm RMPRER) Store(b PRER) { rm.UM32.Store(uint32(b)) }

func PREDIV_S_(p *Periph) RMPRER {
	return RMPRER{mmio.UM32{&p.PRER.U32, uint32(PREDIV_S)}}
}

func PREDIV_A_(p *Periph) RMPRER {
	return RMPRER{mmio.UM32{&p.PRER.U32, uint32(PREDIV_A)}}
}

type WUTR uint32

type RWUTR struct{ mmio.U32 }

func (r *RWUTR) LoadBits(mask WUTR) WUTR { return WUTR(r.U32.LoadBits(uint32(mask))) }
func (r *RWUTR) StoreBits(mask, b WUTR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWUTR) SetBits(mask WUTR)       { r.U32.SetBits(uint32(mask)) }
func (r *RWUTR) ClearBits(mask WUTR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RWUTR) Load() WUTR              { return WUTR(r.U32.Load()) }
func (r *RWUTR) Store(b WUTR)            { r.U32.Store(uint32(b)) }

type RMWUTR struct{ mmio.UM32 }

func (rm RMWUTR) Load() WUTR   { return WUTR(rm.UM32.Load()) }
func (rm RMWUTR) Store(b WUTR) { rm.UM32.Store(uint32(b)) }

func WUT_(p *Periph) RMWUTR {
	return RMWUTR{mmio.UM32{&p.WUTR.U32, uint32(WUT)}}
}

type CR uint32

type RCR struct{ mmio.U32 }

func (r *RCR) LoadBits(mask CR) CR  { return CR(r.U32.LoadBits(uint32(mask))) }
func (r *RCR) StoreBits(mask, b CR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR) SetBits(mask CR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR) ClearBits(mask CR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR) Load() CR             { return CR(r.U32.Load()) }
func (r *RCR) Store(b CR)           { r.U32.Store(uint32(b)) }

type RMCR struct{ mmio.UM32 }

func (rm RMCR) Load() CR   { return CR(rm.UM32.Load()) }
func (rm RMCR) Store(b CR) { rm.UM32.Store(uint32(b)) }

func WCKSEL_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(WCKSEL)}}
}

func TSEDGE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TSEDGE)}}
}

func REFCKON_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(REFCKON)}}
}

func BYPSHAD_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(BYPSHAD)}}
}

func FMT_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(FMT)}}
}

func ALRAE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALRAE)}}
}

func ALRBE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALRBE)}}
}

func WUTE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(WUTE)}}
}

func TSE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TSE)}}
}

func ALRAIE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALRAIE)}}
}

func ALRBIE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ALRBIE)}}
}

func WUTIE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(WUTIE)}}
}

func TSIE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TSIE)}}
}

func ADD1H_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ADD1H)}}
}

func SUB1H_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(SUB1H)}}
}

func BKP_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(BKP)}}
}

func COSEL_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(COSEL)}}
}

func POL_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(POL)}}
}

func OSEL_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(OSEL)}}
}

func COE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(COE)}}
}

func ITSE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ITSE)}}
}

func TAMPTS_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TAMPTS)}}
}

func TAMPOE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TAMPOE)}}
}

func TAMPALRM_PU_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TAMPALRM_PU)}}
}

func TAMPALRM_TYPE_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TAMPALRM_TYPE)}}
}

func OUT2EN_(p *Periph) RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(OUT2EN)}}
}

type WPR uint32

type RWPR struct{ mmio.U32 }

func (r *RWPR) LoadBits(mask WPR) WPR { return WPR(r.U32.LoadBits(uint32(mask))) }
func (r *RWPR) StoreBits(mask, b WPR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWPR) SetBits(mask WPR)      { r.U32.SetBits(uint32(mask)) }
func (r *RWPR) ClearBits(mask WPR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RWPR) Load() WPR             { return WPR(r.U32.Load()) }
func (r *RWPR) Store(b WPR)           { r.U32.Store(uint32(b)) }

type RMWPR struct{ mmio.UM32 }

func (rm RMWPR) Load() WPR   { return WPR(rm.UM32.Load()) }
func (rm RMWPR) Store(b WPR) { rm.UM32.Store(uint32(b)) }

func KEY_(p *Periph) RMWPR {
	return RMWPR{mmio.UM32{&p.WPR.U32, uint32(KEY)}}
}

type CALR uint32

type RCALR struct{ mmio.U32 }

func (r *RCALR) LoadBits(mask CALR) CALR { return CALR(r.U32.LoadBits(uint32(mask))) }
func (r *RCALR) StoreBits(mask, b CALR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCALR) SetBits(mask CALR)       { r.U32.SetBits(uint32(mask)) }
func (r *RCALR) ClearBits(mask CALR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RCALR) Load() CALR              { return CALR(r.U32.Load()) }
func (r *RCALR) Store(b CALR)            { r.U32.Store(uint32(b)) }

type RMCALR struct{ mmio.UM32 }

func (rm RMCALR) Load() CALR   { return CALR(rm.UM32.Load()) }
func (rm RMCALR) Store(b CALR) { rm.UM32.Store(uint32(b)) }

func CALM_(p *Periph) RMCALR {
	return RMCALR{mmio.UM32{&p.CALR.U32, uint32(CALM)}}
}

func CALW16_(p *Periph) RMCALR {
	return RMCALR{mmio.UM32{&p.CALR.U32, uint32(CALW16)}}
}

func CALW8_(p *Periph) RMCALR {
	return RMCALR{mmio.UM32{&p.CALR.U32, uint32(CALW8)}}
}

func CALP_(p *Periph) RMCALR {
	return RMCALR{mmio.UM32{&p.CALR.U32, uint32(CALP)}}
}

type SHIFTR uint32

type RSHIFTR struct{ mmio.U32 }

func (r *RSHIFTR) LoadBits(mask SHIFTR) SHIFTR { return SHIFTR(r.U32.LoadBits(uint32(mask))) }
func (r *RSHIFTR) StoreBits(mask, b SHIFTR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSHIFTR) SetBits(mask SHIFTR)         { r.U32.SetBits(uint32(mask)) }
func (r *RSHIFTR) ClearBits(mask SHIFTR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RSHIFTR) Load() SHIFTR                { return SHIFTR(r.U32.Load()) }
func (r *RSHIFTR) Store(b SHIFTR)              { r.U32.Store(uint32(b)) }

type RMSHIFTR struct{ mmio.UM32 }

func (rm RMSHIFTR) Load() SHIFTR   { return SHIFTR(rm.UM32.Load()) }
func (rm RMSHIFTR) Store(b SHIFTR) { rm.UM32.Store(uint32(b)) }

func SUBFS_(p *Periph) RMSHIFTR {
	return RMSHIFTR{mmio.UM32{&p.SHIFTR.U32, uint32(SUBFS)}}
}

func ADD1S_(p *Periph) RMSHIFTR {
	return RMSHIFTR{mmio.UM32{&p.SHIFTR.U32, uint32(ADD1S)}}
}

type ALRMR uint32

type RALRMR struct{ mmio.U32 }

func (r *RALRMR) LoadBits(mask ALRMR) ALRMR { return ALRMR(r.U32.LoadBits(uint32(mask))) }
func (r *RALRMR) StoreBits(mask, b ALRMR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RALRMR) SetBits(mask ALRMR)        { r.U32.SetBits(uint32(mask)) }
func (r *RALRMR) ClearBits(mask ALRMR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RALRMR) Load() ALRMR               { return ALRMR(r.U32.Load()) }
func (r *RALRMR) Store(b ALRMR)             { r.U32.Store(uint32(b)) }

type RMALRMR struct{ mmio.UM32 }

func (rm RMALRMR) Load() ALRMR   { return ALRMR(rm.UM32.Load()) }
func (rm RMALRMR) Store(b ALRMR) { rm.UM32.Store(uint32(b)) }

type ALRMSSR uint32

type RALRMSSR struct{ mmio.U32 }

func (r *RALRMSSR) LoadBits(mask ALRMSSR) ALRMSSR { return ALRMSSR(r.U32.LoadBits(uint32(mask))) }
func (r *RALRMSSR) StoreBits(mask, b ALRMSSR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RALRMSSR) SetBits(mask ALRMSSR)          { r.U32.SetBits(uint32(mask)) }
func (r *RALRMSSR) ClearBits(mask ALRMSSR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RALRMSSR) Load() ALRMSSR                 { return ALRMSSR(r.U32.Load()) }
func (r *RALRMSSR) Store(b ALRMSSR)               { r.U32.Store(uint32(b)) }

type RMALRMSSR struct{ mmio.UM32 }

func (rm RMALRMSSR) Load() ALRMSSR   { return ALRMSSR(rm.UM32.Load()) }
func (rm RMALRMSSR) Store(b ALRMSSR) { rm.UM32.Store(uint32(b)) }

type SR uint32

type RSR struct{ mmio.U32 }

func (r *RSR) LoadBits(mask SR) SR  { return SR(r.U32.LoadBits(uint32(mask))) }
func (r *RSR) StoreBits(mask, b SR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR) SetBits(mask SR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR) ClearBits(mask SR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR) Load() SR             { return SR(r.U32.Load()) }
func (r *RSR) Store(b SR)           { r.U32.Store(uint32(b)) }

type RMSR struct{ mmio.UM32 }

func (rm RMSR) Load() SR   { return SR(rm.UM32.Load()) }
func (rm RMSR) Store(b SR) { rm.UM32.Store(uint32(b)) }

func ALRAF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(ALRAF)}}
}

func ALRBF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(ALRBF)}}
}

func WUTF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(WUTF)}}
}

func TSF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TSF)}}
}

func TSOVF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TSOVF)}}
}

func ITSF_(p *Periph) RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(ITSF)}}
}

type MISR uint32

type RMISR struct{ mmio.U32 }

func (r *RMISR) LoadBits(mask MISR) MISR { return MISR(r.U32.LoadBits(uint32(mask))) }
func (r *RMISR) StoreBits(mask, b MISR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMISR) SetBits(mask MISR)       { r.U32.SetBits(uint32(mask)) }
func (r *RMISR) ClearBits(mask MISR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RMISR) Load() MISR              { return MISR(r.U32.Load()) }
func (r *RMISR) Store(b MISR)            { r.U32.Store(uint32(b)) }

type RMMISR struct{ mmio.UM32 }

func (rm RMMISR) Load() MISR   { return MISR(rm.UM32.Load()) }
func (rm RMMISR) Store(b MISR) { rm.UM32.Store(uint32(b)) }

func ALRAMF_(p *Periph) RMMISR {
	return RMMISR{mmio.UM32{&p.MISR.U32, uint32(ALRAMF)}}
}

func ALRBMF_(p *Periph) RMMISR {
	return RMMISR{mmio.UM32{&p.MISR.U32, uint32(ALRBMF)}}
}

func WUTMF_(p *Periph) RMMISR {
	return RMMISR{mmio.UM32{&p.MISR.U32, uint32(WUTMF)}}
}

func TSMF_(p *Periph) RMMISR {
	return RMMISR{mmio.UM32{&p.MISR.U32, uint32(TSMF)}}
}

func TSOVMF_(p *Periph) RMMISR {
	return RMMISR{mmio.UM32{&p.MISR.U32, uint32(TSOVMF)}}
}

func ITSMF_(p *Periph) RMMISR {
	return RMMISR{mmio.UM32{&p.MISR.U32, uint32(ITSMF)}}
}

type SCR uint32

type RSCR struct{ mmio.U32 }

func (r *RSCR) LoadBits(mask SCR) SCR { return SCR(r.U32.LoadBits(uint32(mask))) }
func (r *RSCR) StoreBits(mask, b SCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSCR) SetBits(mask SCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSCR) ClearBits(mask SCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSCR) Load() SCR             { return SCR(r.U32.Load()) }
func (r *RSCR) Store(b SCR)           { r.U32.Store(uint32(b)) }

type RMSCR struct{ mmio.UM32 }

func (rm RMSCR) Load() SCR   { return SCR(rm.UM32.Load()) }
func (rm RMSCR) Store(b SCR) { rm.UM32.Store(uint32(b)) }

func CALRAF_(p *Periph) RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(CALRAF)}}
}

func CALRBF_(p *Periph) RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(CALRBF)}}
}

func CWUTF_(p *Periph) RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(CWUTF)}}
}

func CTSF_(p *Periph) RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(CTSF)}}
}

func CTSOVF_(p *Periph) RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(CTSOVF)}}
}

func CITSF_(p *Periph) RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(CITSF)}}
}
