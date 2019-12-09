package main

import (
	"embedded/rtos"

	"github.com/embeddedgo/stm32/hal/exti"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi1"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/init"
)

const (
	CTRL_REG4 = 0x20
	CTRL_REG3 = 0x23
	OUT_X_L   = 0x28
)

var accel *LIS3DSH

type LIS3DSH struct {
	d   *spi.Driver
	cs  gpio.Pin
	dri exti.Lines
	drn rtos.Note
}

func NewLIS3DSH(d *spi.Driver, cs gpio.Pin, dri exti.Lines) *LIS3DSH {
	return &LIS3DSH{d: d, cs: cs, dri: dri}
}

func (a *LIS3DSH) Init() {
	rtos.Nanosleep(10e6 - rtos.Nanotime())
	a.cs.Clear()
	a.d.WriteRead([]byte{CTRL_REG4, 0x67}, nil)
	a.cs.Set()
	a.cs.Clear()
	a.d.WriteRead([]byte{CTRL_REG3, 0xC8}, nil)
	a.cs.Set()
}

func (a *LIS3DSH) ReadXYZ() (x, y, z int) {
	a.drn.Clear()
	a.dri.EnableIRQ()
	a.drn.Sleep(-1)
	buf := [1 + 6]byte{1<<7 | OUT_X_L}
	a.cs.Clear()
	a.d.WriteRead(buf[:1], buf[:])
	a.cs.Set()
	x = int(int16(buf[1])|int16(buf[2])<<8) * 2000 / 32768
	y = int(int16(buf[3])|int16(buf[4])<<8) * 2000 / 32768
	z = int(int16(buf[5])|int16(buf[6])<<8) * 2000 / 32768
	return
}

func (a *LIS3DSH) ISR() {
	a.dri.DisableIRQ()
	a.dri.ClearPending()
	a.drn.Wakeup()
}

func main() {
	// Allocate GPIO pins

	pa := gpio.A()
	pa.EnableClock(true)
	sck := pa.Pin(5)
	miso := pa.Pin(6)
	mosi := pa.Pin(7)

	pe := gpio.E()
	pe.EnableClock(true)
	dr := pe.Pin(0)
	cs := pe.Pin(3)

	// Configure SPI pins

	spi1.UsePinsMaster(sck, mosi, miso)
	cs.Set() // CS active state is low
	cs.Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.High})

	// Configure EXTI

	dr.Setup(&gpio.Config{Mode: gpio.In})
	dri := exti.Lines(1 << dr.Index())
	dri.Connect(dr.Port())
	dri.EnableRiseTrig()
	irq.EXTI0.Enable(rtos.IntPrioLow)

	// Configure and enable SPI

	d := spi1.Driver()
	d.Setup(spi.Master|spi.CPOL1|spi.CPHA1|spi.SoftSS|spi.ISSHigh, 10e6)
	d.SetWordSize(8)
	d.Enable()

	// Reading acceleration data using DRDY interrupt

	accel = NewLIS3DSH(d, cs, dri)
	accel.Init()
	for {
		show(accel.ReadXYZ())
	}
}

func show(x, y, z int) {
	gauge("x", x)
	gauge("y", y)
	gauge("z", z)
	print("+-----------------------------------------+\n")
	rtos.Nanosleep(100e6)
}

func gauge(name string, v int) {
	const bar = "--------------------"
	const spc = "                    "
	v /= 100
	print("|")
	switch {
	case v > 0:
		if v > 20 {
			v = 20
		}
		print(spc, name, bar[:v], spc[v:])
	case v < 0:
		if v < -20 {
			v = -20
		}
		print(spc[:20+v], bar[20+v:], name, spc)
	default: // v == 0:
		print(spc, name, spc)
	}
	print("|\n")
}

//go:interrupthandler
func EXTI0_Handler() { accel.ISR() }
