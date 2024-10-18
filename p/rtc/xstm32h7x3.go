// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build stm32h7x3

package rtc

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	TR       mmio.R32[TR]
	DR       mmio.R32[DR]
	CR       mmio.R32[CR]
	ISR      mmio.R32[ISR]
	PRER     mmio.R32[PRER]
	WUTR     mmio.R32[WUTR]
	_        uint32
	ALRMAR   mmio.R32[ALRMR]
	ALRMBR   mmio.R32[ALRMR]
	WPR      mmio.R32[WPR]
	SSR      mmio.R32[SSR]
	SHIFTR   mmio.R32[SHIFTR]
	TSTR     mmio.R32[TR]
	TSDR     mmio.R32[DR]
	TSSSR    mmio.R32[SSR]
	CALR     mmio.R32[CALR]
	TAMPCR   mmio.R32[TAMPCR]
	ALRMASSR mmio.R32[ALRMSSR]
	ALRMBSSR mmio.R32[ALRMSSR]
	OR       mmio.R32[OR]
	BKPR     [32]mmio.R32[uint32]
}

func RTC() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.RTC_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type TR uint32

func SU_(p *Periph) mmio.RM32[TR]  { return mmio.RM32[TR]{&p.TR, SU} }
func ST_(p *Periph) mmio.RM32[TR]  { return mmio.RM32[TR]{&p.TR, ST} }
func MNU_(p *Periph) mmio.RM32[TR] { return mmio.RM32[TR]{&p.TR, MNU} }
func MNT_(p *Periph) mmio.RM32[TR] { return mmio.RM32[TR]{&p.TR, MNT} }
func HU_(p *Periph) mmio.RM32[TR]  { return mmio.RM32[TR]{&p.TR, HU} }
func HT_(p *Periph) mmio.RM32[TR]  { return mmio.RM32[TR]{&p.TR, HT} }
func PM_(p *Periph) mmio.RM32[TR]  { return mmio.RM32[TR]{&p.TR, PM} }

type DR uint32

func DU_(p *Periph) mmio.RM32[DR]  { return mmio.RM32[DR]{&p.DR, DU} }
func DT_(p *Periph) mmio.RM32[DR]  { return mmio.RM32[DR]{&p.DR, DT} }
func MU_(p *Periph) mmio.RM32[DR]  { return mmio.RM32[DR]{&p.DR, MU} }
func MT_(p *Periph) mmio.RM32[DR]  { return mmio.RM32[DR]{&p.DR, MT} }
func WDU_(p *Periph) mmio.RM32[DR] { return mmio.RM32[DR]{&p.DR, WDU} }
func YU_(p *Periph) mmio.RM32[DR]  { return mmio.RM32[DR]{&p.DR, YU} }
func YT_(p *Periph) mmio.RM32[DR]  { return mmio.RM32[DR]{&p.DR, YT} }

type CR uint32

func WUCKSEL_(p *Periph) mmio.RM32[CR] { return mmio.RM32[CR]{&p.CR, WUCKSEL} }
func TSEDGE_(p *Periph) mmio.RM32[CR]  { return mmio.RM32[CR]{&p.CR, TSEDGE} }
func REFCKON_(p *Periph) mmio.RM32[CR] { return mmio.RM32[CR]{&p.CR, REFCKON} }
func BYPSHAD_(p *Periph) mmio.RM32[CR] { return mmio.RM32[CR]{&p.CR, BYPSHAD} }
func FMT_(p *Periph) mmio.RM32[CR]     { return mmio.RM32[CR]{&p.CR, FMT} }
func ALRAE_(p *Periph) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.CR, ALRAE} }
func ALRBE_(p *Periph) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.CR, ALRBE} }
func WUTE_(p *Periph) mmio.RM32[CR]    { return mmio.RM32[CR]{&p.CR, WUTE} }
func TSE_(p *Periph) mmio.RM32[CR]     { return mmio.RM32[CR]{&p.CR, TSE} }
func ALRAIE_(p *Periph) mmio.RM32[CR]  { return mmio.RM32[CR]{&p.CR, ALRAIE} }
func ALRBIE_(p *Periph) mmio.RM32[CR]  { return mmio.RM32[CR]{&p.CR, ALRBIE} }
func WUTIE_(p *Periph) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.CR, WUTIE} }
func TSIE_(p *Periph) mmio.RM32[CR]    { return mmio.RM32[CR]{&p.CR, TSIE} }
func ADD1H_(p *Periph) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.CR, ADD1H} }
func SUB1H_(p *Periph) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.CR, SUB1H} }
func BKP_(p *Periph) mmio.RM32[CR]     { return mmio.RM32[CR]{&p.CR, BKP} }
func COSEL_(p *Periph) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.CR, COSEL} }
func POL_(p *Periph) mmio.RM32[CR]     { return mmio.RM32[CR]{&p.CR, POL} }
func OSEL_(p *Periph) mmio.RM32[CR]    { return mmio.RM32[CR]{&p.CR, OSEL} }
func COE_(p *Periph) mmio.RM32[CR]     { return mmio.RM32[CR]{&p.CR, COE} }
func ITSE_(p *Periph) mmio.RM32[CR]    { return mmio.RM32[CR]{&p.CR, ITSE} }

type ISR uint32

func ALRAWF_(p *Periph) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR, ALRAWF} }
func ALRBWF_(p *Periph) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR, ALRBWF} }
func WUTWF_(p *Periph) mmio.RM32[ISR]   { return mmio.RM32[ISR]{&p.ISR, WUTWF} }
func SHPF_(p *Periph) mmio.RM32[ISR]    { return mmio.RM32[ISR]{&p.ISR, SHPF} }
func INITS_(p *Periph) mmio.RM32[ISR]   { return mmio.RM32[ISR]{&p.ISR, INITS} }
func RSF_(p *Periph) mmio.RM32[ISR]     { return mmio.RM32[ISR]{&p.ISR, RSF} }
func INITF_(p *Periph) mmio.RM32[ISR]   { return mmio.RM32[ISR]{&p.ISR, INITF} }
func INIT_(p *Periph) mmio.RM32[ISR]    { return mmio.RM32[ISR]{&p.ISR, INIT} }
func ALRAF_(p *Periph) mmio.RM32[ISR]   { return mmio.RM32[ISR]{&p.ISR, ALRAF} }
func ALRBF_(p *Periph) mmio.RM32[ISR]   { return mmio.RM32[ISR]{&p.ISR, ALRBF} }
func WUTF_(p *Periph) mmio.RM32[ISR]    { return mmio.RM32[ISR]{&p.ISR, WUTF} }
func TSF_(p *Periph) mmio.RM32[ISR]     { return mmio.RM32[ISR]{&p.ISR, TSF} }
func TSOVF_(p *Periph) mmio.RM32[ISR]   { return mmio.RM32[ISR]{&p.ISR, TSOVF} }
func TAMP1F_(p *Periph) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR, TAMP1F} }
func TAMP2F_(p *Periph) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR, TAMP2F} }
func TAMP3F_(p *Periph) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR, TAMP3F} }
func RECALPF_(p *Periph) mmio.RM32[ISR] { return mmio.RM32[ISR]{&p.ISR, RECALPF} }
func ITSF_(p *Periph) mmio.RM32[ISR]    { return mmio.RM32[ISR]{&p.ISR, ITSF} }

type PRER uint32

func PREDIV_S_(p *Periph) mmio.RM32[PRER] { return mmio.RM32[PRER]{&p.PRER, PREDIV_S} }
func PREDIV_A_(p *Periph) mmio.RM32[PRER] { return mmio.RM32[PRER]{&p.PRER, PREDIV_A} }

type WUTR uint32

func WUT_(p *Periph) mmio.RM32[WUTR] { return mmio.RM32[WUTR]{&p.WUTR, WUT} }

type ALRMR uint32

type WPR uint32

func KEY_(p *Periph) mmio.RM32[WPR] { return mmio.RM32[WPR]{&p.WPR, KEY} }

type SSR uint32

func SS_(p *Periph) mmio.RM32[SSR] { return mmio.RM32[SSR]{&p.SSR, SS} }

type SHIFTR uint32

func SUBFS_(p *Periph) mmio.RM32[SHIFTR] { return mmio.RM32[SHIFTR]{&p.SHIFTR, SUBFS} }
func ADD1S_(p *Periph) mmio.RM32[SHIFTR] { return mmio.RM32[SHIFTR]{&p.SHIFTR, ADD1S} }

type CALR uint32

func CALM_(p *Periph) mmio.RM32[CALR]   { return mmio.RM32[CALR]{&p.CALR, CALM} }
func CALW16_(p *Periph) mmio.RM32[CALR] { return mmio.RM32[CALR]{&p.CALR, CALW16} }
func CALW8_(p *Periph) mmio.RM32[CALR]  { return mmio.RM32[CALR]{&p.CALR, CALW8} }
func CALP_(p *Periph) mmio.RM32[CALR]   { return mmio.RM32[CALR]{&p.CALR, CALP} }

type TAMPCR uint32

func TAMP1E_(p *Periph) mmio.RM32[TAMPCR]       { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP1E} }
func TAMP1TRG_(p *Periph) mmio.RM32[TAMPCR]     { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP1TRG} }
func TAMPIE_(p *Periph) mmio.RM32[TAMPCR]       { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMPIE} }
func TAMP2E_(p *Periph) mmio.RM32[TAMPCR]       { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP2E} }
func TAMP2TRG_(p *Periph) mmio.RM32[TAMPCR]     { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP2TRG} }
func TAMP3E_(p *Periph) mmio.RM32[TAMPCR]       { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP3E} }
func TAMP3TRG_(p *Periph) mmio.RM32[TAMPCR]     { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP3TRG} }
func TAMPTS_(p *Periph) mmio.RM32[TAMPCR]       { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMPTS} }
func TAMPFREQ_(p *Periph) mmio.RM32[TAMPCR]     { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMPFREQ} }
func TAMPFLT_(p *Periph) mmio.RM32[TAMPCR]      { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMPFLT} }
func TAMPPRCH_(p *Periph) mmio.RM32[TAMPCR]     { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMPPRCH} }
func TAMPPUDIS_(p *Periph) mmio.RM32[TAMPCR]    { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMPPUDIS} }
func TAMP1IE_(p *Periph) mmio.RM32[TAMPCR]      { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP1IE} }
func TAMP1NOERASE_(p *Periph) mmio.RM32[TAMPCR] { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP1NOERASE} }
func TAMP1MF_(p *Periph) mmio.RM32[TAMPCR]      { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP1MF} }
func TAMP2IE_(p *Periph) mmio.RM32[TAMPCR]      { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP2IE} }
func TAMP2NOERASE_(p *Periph) mmio.RM32[TAMPCR] { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP2NOERASE} }
func TAMP2MF_(p *Periph) mmio.RM32[TAMPCR]      { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP2MF} }
func TAMP3IE_(p *Periph) mmio.RM32[TAMPCR]      { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP3IE} }
func TAMP3NOERASE_(p *Periph) mmio.RM32[TAMPCR] { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP3NOERASE} }
func TAMP3MF_(p *Periph) mmio.RM32[TAMPCR]      { return mmio.RM32[TAMPCR]{&p.TAMPCR, TAMP3MF} }

type ALRMSSR uint32

type OR uint32

func RTC_ALARM_TYPE_(p *Periph) mmio.RM32[OR] { return mmio.RM32[OR]{&p.OR, RTC_ALARM_TYPE} }
func RTC_OUT_RMP_(p *Periph) mmio.RM32[OR]    { return mmio.RM32[OR]{&p.OR, RTC_OUT_RMP} }
