// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build stm32f215

package dma

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/mmap"
)

type Periph struct {
	ISR  [2]mmio.R32[ISR]
	IFCR [2]mmio.R32[IFCR]
	S    [8]SS
}

func DMA1() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.DMA1_BASE))) }
func DMA2() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.DMA2_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	return bus.AHB1
}

type ISR uint32

func FEIF0_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], FEIF0} }
func DMEIF0_(p *Periph, i int) mmio.RM32[ISR] { return mmio.RM32[ISR]{&p.ISR[i], DMEIF0} }
func TEIF0_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], TEIF0} }
func HTIF0_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], HTIF0} }
func TCIF0_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], TCIF0} }
func FEIF1_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], FEIF1} }
func DMEIF1_(p *Periph, i int) mmio.RM32[ISR] { return mmio.RM32[ISR]{&p.ISR[i], DMEIF1} }
func TEIF1_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], TEIF1} }
func HTIF1_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], HTIF1} }
func TCIF1_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], TCIF1} }
func FEIF2_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], FEIF2} }
func DMEIF2_(p *Periph, i int) mmio.RM32[ISR] { return mmio.RM32[ISR]{&p.ISR[i], DMEIF2} }
func TEIF2_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], TEIF2} }
func HTIF2_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], HTIF2} }
func TCIF2_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], TCIF2} }
func FEIF3_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], FEIF3} }
func DMEIF3_(p *Periph, i int) mmio.RM32[ISR] { return mmio.RM32[ISR]{&p.ISR[i], DMEIF3} }
func TEIF3_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], TEIF3} }
func HTIF3_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], HTIF3} }
func TCIF3_(p *Periph, i int) mmio.RM32[ISR]  { return mmio.RM32[ISR]{&p.ISR[i], TCIF3} }

type IFCR uint32

func CFEIF0_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CFEIF0} }
func CDMEIF0_(p *Periph, i int) mmio.RM32[IFCR] { return mmio.RM32[IFCR]{&p.IFCR[i], CDMEIF0} }
func CTEIF0_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CTEIF0} }
func CHTIF0_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CHTIF0} }
func CTCIF0_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CTCIF0} }
func CFEIF1_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CFEIF1} }
func CDMEIF1_(p *Periph, i int) mmio.RM32[IFCR] { return mmio.RM32[IFCR]{&p.IFCR[i], CDMEIF1} }
func CTEIF1_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CTEIF1} }
func CHTIF1_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CHTIF1} }
func CTCIF1_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CTCIF1} }
func CFEIF2_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CFEIF2} }
func CDMEIF2_(p *Periph, i int) mmio.RM32[IFCR] { return mmio.RM32[IFCR]{&p.IFCR[i], CDMEIF2} }
func CTEIF2_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CTEIF2} }
func CHTIF2_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CHTIF2} }
func CTCIF2_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CTCIF2} }
func CFEIF3_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CFEIF3} }
func CDMEIF3_(p *Periph, i int) mmio.RM32[IFCR] { return mmio.RM32[IFCR]{&p.IFCR[i], CDMEIF3} }
func CTEIF3_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CTEIF3} }
func CHTIF3_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CHTIF3} }
func CTCIF3_(p *Periph, i int) mmio.RM32[IFCR]  { return mmio.RM32[IFCR]{&p.IFCR[i], CTCIF3} }

type SS struct {
	CR   mmio.R32[CR]
	NDTR mmio.R32[NDTR]
	PAR  mmio.R32[PAR]
	M0AR mmio.R32[M0AR]
	M1AR mmio.R32[M1AR]
	FCR  mmio.R32[FCR]
}

type CR uint32

func EN_(p *Periph, i int) mmio.RM32[CR]     { return mmio.RM32[CR]{&p.S[i].CR, EN} }
func DMEIE_(p *Periph, i int) mmio.RM32[CR]  { return mmio.RM32[CR]{&p.S[i].CR, DMEIE} }
func TEIE_(p *Periph, i int) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.S[i].CR, TEIE} }
func HTIE_(p *Periph, i int) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.S[i].CR, HTIE} }
func TCIE_(p *Periph, i int) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.S[i].CR, TCIE} }
func PFCTRL_(p *Periph, i int) mmio.RM32[CR] { return mmio.RM32[CR]{&p.S[i].CR, PFCTRL} }
func DIR_(p *Periph, i int) mmio.RM32[CR]    { return mmio.RM32[CR]{&p.S[i].CR, DIR} }
func CIRC_(p *Periph, i int) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.S[i].CR, CIRC} }
func PINC_(p *Periph, i int) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.S[i].CR, PINC} }
func MINC_(p *Periph, i int) mmio.RM32[CR]   { return mmio.RM32[CR]{&p.S[i].CR, MINC} }
func PSIZE_(p *Periph, i int) mmio.RM32[CR]  { return mmio.RM32[CR]{&p.S[i].CR, PSIZE} }
func MSIZE_(p *Periph, i int) mmio.RM32[CR]  { return mmio.RM32[CR]{&p.S[i].CR, MSIZE} }
func PINCOS_(p *Periph, i int) mmio.RM32[CR] { return mmio.RM32[CR]{&p.S[i].CR, PINCOS} }
func PL_(p *Periph, i int) mmio.RM32[CR]     { return mmio.RM32[CR]{&p.S[i].CR, PL} }
func DBM_(p *Periph, i int) mmio.RM32[CR]    { return mmio.RM32[CR]{&p.S[i].CR, DBM} }
func CT_(p *Periph, i int) mmio.RM32[CR]     { return mmio.RM32[CR]{&p.S[i].CR, CT} }
func PBURST_(p *Periph, i int) mmio.RM32[CR] { return mmio.RM32[CR]{&p.S[i].CR, PBURST} }
func MBURST_(p *Periph, i int) mmio.RM32[CR] { return mmio.RM32[CR]{&p.S[i].CR, MBURST} }
func CHSEL_(p *Periph, i int) mmio.RM32[CR]  { return mmio.RM32[CR]{&p.S[i].CR, CHSEL} }

type NDTR uint32

func NDT_(p *Periph, i int) mmio.RM32[NDTR] { return mmio.RM32[NDTR]{&p.S[i].NDTR, NDT} }

type PAR uint32

func PA_(p *Periph, i int) mmio.RM32[PAR] { return mmio.RM32[PAR]{&p.S[i].PAR, PA} }

type M0AR uint32

func M0A_(p *Periph, i int) mmio.RM32[M0AR] { return mmio.RM32[M0AR]{&p.S[i].M0AR, M0A} }

type M1AR uint32

func M1A_(p *Periph, i int) mmio.RM32[M1AR] { return mmio.RM32[M1AR]{&p.S[i].M1AR, M1A} }

type FCR uint32

func FTH_(p *Periph, i int) mmio.RM32[FCR]   { return mmio.RM32[FCR]{&p.S[i].FCR, FTH} }
func DMDIS_(p *Periph, i int) mmio.RM32[FCR] { return mmio.RM32[FCR]{&p.S[i].FCR, DMDIS} }
func FS_(p *Periph, i int) mmio.RM32[FCR]    { return mmio.RM32[FCR]{&p.S[i].FCR, FS} }
func FEIE_(p *Periph, i int) mmio.RM32[FCR]  { return mmio.RM32[FCR]{&p.S[i].FCR, FEIE} }
