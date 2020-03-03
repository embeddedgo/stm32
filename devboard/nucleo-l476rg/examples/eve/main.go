// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"time"

	"github.com/embeddedgo/display/eve"
	"github.com/embeddedgo/stm32/dci/evedci"
	"github.com/embeddedgo/stm32/hal/exti"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/irq"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi2"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/init"
)

var eveInt rtos.Note

func main() {
	// Allocate GPIO pins

	pb := gpio.B()
	pb.EnableClock(true)
	intn := pb.Pin(1)
	pdn := pb.Pin(11)
	csn := pb.Pin(12)
	sck := pb.Pin(13)
	miso := pb.Pin(14)
	mosi := pb.Pin(15)

	// EVE driver

	evedci.UseIntPin(intn)
	irq.EXTI1.Enable(rtos.IntPrioLowest)

	spi2.UsePinMaster(spi.NSS, csn)
	spi2.UsePinMaster(spi.SCK, sck)
	spi2.UsePinMaster(spi.MOSI, mosi)
	spi2.UsePinMaster(spi.MISO, miso)

	lcd := eve.NewDriver(evedci.NewSPI(spi2.Driver(), pdn), 256)
	lcd.SetNote(&eveInt)

	////

	checkErr(lcd.Init(&eve.Default480x272, nil))
	lcd.SetBacklight(64)

	width, height := lcd.Width(), lcd.Height()

	dl := lcd.DL(-1)
	dl.Clear(eve.CST)
	dl.Begin(eve.POINTS)
	dl.Vertex2f(200<<4, 100<<4)
	dl.Display()
	dl.SwapDL()

	waitTouch(lcd)

	dl = lcd.DL(-1)
	dl.Clear(eve.CST)
	dl.Begin(eve.POINTS)
	dl.PointSize(70 << 4)
	dl.Vertex2f(200<<4, 100<<4)
	dl.Display()
	dl.SwapDL()

	waitTouch(lcd)

	dl = lcd.DL(-1)
	dl.Clear(eve.CST)
	dl.Begin(eve.POINTS)
	dl.PointSize(70 << 4)
	dl.Vertex2f(200<<4, 100<<4)
	dl.ColorRGB(0x0000FF)
	dl.PointSize(50 << 4)
	dl.Vertex2f(240<<4, 150<<4)
	dl.Display()
	dl.SwapDL()

	waitTouch(lcd)

	dl = lcd.DL(-1)
	dl.Clear(eve.CST)
	dl.Begin(eve.POINTS)
	dl.PointSize(70 << 4)
	dl.Vertex2f(200<<4, 100<<4)
	dl.ColorRGB(0x0000FF)
	dl.ColorA(128)
	dl.PointSize(50 << 4)
	dl.Vertex2f(240<<4, 150<<4)
	dl.Display()
	dl.SwapDL()

	waitTouch(lcd)

	dl = lcd.DL(-1)
	dl.Clear(eve.CST)
	dl.Begin(eve.BITMAPS)
	dl.BitmapHandle(31)
	dl.Cell('E')
	dl.Vertex2f(200<<4, 100<<4)
	dl.Display()
	dl.SwapDL()

	waitTouch(lcd)

	dl = lcd.DL(-1)
	dl.Clear(eve.CST)
	dl.Begin(eve.BITMAPS)
	dl.BitmapHandle(31)
	dl.Cell('E')
	dl.Vertex2f(200<<4, 100<<4)
	dl.Cell('V')
	dl.Vertex2f(224<<4, 100<<4)
	dl.Cell('E')
	dl.Vertex2f(250<<4, 100<<4)
	dl.Display()
	dl.SwapDL()

	waitTouch(lcd)

	ce := lcd.CE()
	ce.DLStart()
	ce.Clear(eve.CST)
	ce.TextString(width/2, height/2, 31, eve.OPT_CENTER, "Hello world!")
	ce.Display()
	ce.Swap()
	ce.Close()

	waitTouch(lcd)
	time.Sleep(200 * time.Millisecond)

calibration:
	ce = lcd.CE()
	ce.DLStart()
	ce.Clear(eve.CST)
	ce.TextString(
		width/2, height/2, 30, eve.OPT_CENTER,
		"Touch panel calibration",
	)
	addr := ce.Calibrate()
	ce.Close()

	if lcd.ReadUint32(addr) == 0 {
		ce = lcd.CE()
		ce.DLStart()
		ce.Clear(eve.CST)
		ce.TextString(
			width/2, height/2, 30, eve.OPT_CENTER,
			"Calibration failed!",
		)
		ce.Close()
		time.Sleep(2 * time.Second)
		goto calibration
	}

	w := lcd.W(0)
	w.WriteString(gopherMask)
	addr = w.Close()
	addr1 := (len(gopherMask) + 3) &^ 3
	println(addr, addr1)

	ce = lcd.CE()
	ce.DLStart()
	ce.LoadImageString(addr, eve.OPT_RGB565, gopher)
	ce.Close() // A lot of data sent. Ensure free space.

	ce = lcd.CE()
	ce.BitmapHandle(1)
	ce.BitmapLayout(eve.L1, 216/8, 251)
	ce.BitmapSize(eve.DEFAULT, 211, 251)
	ce.Clear(eve.CST)
	ce.Gradient(0, 0, 0x001155, 0, height, 0x772200)
	ce.Begin(eve.BITMAPS)
	ce.ColorMask(eve.A)
	ce.Clear(eve.C)
	ce.BitmapHandle(1)
	ce.Vertex2f(31*16, 21*16)
	ce.ColorMask(eve.RGBA)
	ce.BlendFunc(eve.DST_ALPHA, eve.ONE_MINUS_DST_ALPHA)
	ce.BitmapHandle(0)
	ce.Vertex2f(0, 0)
	ce.Display()
	ce.Swap()
	ce.Close()
}

func checkErr(err error) {
	if err == nil {
		return
	}
	println(err.Error())
	for {
	}
}

func waitTouch(lcd *eve.Driver) {
	time.Sleep(100 * time.Millisecond)
	lcd.ClearInt(eve.IntTouch)
	lcd.WaitInt(eve.IntTouch)
}

//go:interrupthandler
func EXTI1_Handler() {
	exti.L1.ClearPending()
	eveInt.Wakeup()
}
