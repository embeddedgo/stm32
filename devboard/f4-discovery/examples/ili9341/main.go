// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image"
	"image/draw"
	"math/rand"
	"time"

	"github.com/embeddedgo/display/pixd"
	"github.com/embeddedgo/display/pixd/driver/tftdrv/ili9341"
	"github.com/embeddedgo/stm32/dci/tftdci"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi2"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/init"
)

func main() {
	pb := gpio.B()
	pb.EnableClock(true)
	sck := pb.Pin(13)
	miso := pb.Pin(14)
	mosi := pb.Pin(15)

	pd := gpio.D()
	pd.EnableClock(true)
	reset := pd.Pin(8)
	cs := pd.Pin(9)
	dc := pd.Pin(10)

	cfg := &gpio.Config{Mode: gpio.Out, Speed: gpio.Low}
	dc.Clear()
	dc.Setup(cfg)
	reset.Clear() // assert RESET
	reset.Setup(cfg)
	cs.Clear() // assert CS
	cs.Setup(cfg)

	spidrv := spi2.Driver()
	spidrv.UsePinMaster(sck, spi.SCK)
	spidrv.UsePinMaster(miso, spi.MISO)
	spidrv.UsePinMaster(mosi, spi.MOSI)
	spidrv.Setup(spi.Master|spi.MSBF|spi.CPOL0|spi.CPHA0|spi.SoftSS|spi.ISSHigh, 21e6)
	spidrv.Enable()

	// reset the display controller
	reset.Clear()
	time.Sleep(time.Millisecond)
	reset.Set()

	wh := uint16(0xffff)
	bl := uint16(0)
	gr := uint16(31 << 11)
	re := uint16(31)

	buf := []uint16{
		wh, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, wh,
		bl, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, bl,
		bl, wh, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, wh, bl,
		bl, wh, bl, re, re, re, re, re, re, re, re, re, re, re, re, bl, wh, bl,
		bl, wh, bl, re, re, re, re, re, re, re, re, re, re, re, re, bl, wh, bl,
		bl, wh, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, wh, bl,
		bl, wh, bl, gr, gr, gr, gr, gr, gr, gr, gr, gr, gr, gr, gr, bl, wh, bl,
		bl, wh, bl, gr, gr, gr, gr, gr, gr, gr, gr, gr, gr, gr, gr, bl, wh, bl,
		bl, wh, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, wh, bl,
		bl, wh, bl, bl, re, re, re, re, re, re, re, re, re, bl, bl, bl, wh, bl,
		bl, wh, bl, bl, re, re, re, re, re, re, re, re, re, bl, bl, bl, wh, bl,
		bl, wh, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, wh, bl,
		bl, wh, bl, bl, bl, gr, gr, gr, gr, gr, gr, gr, gr, gr, bl, bl, wh, bl,
		bl, wh, bl, bl, bl, gr, gr, gr, gr, gr, gr, gr, gr, gr, bl, bl, wh, bl,
		bl, wh, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, wh, bl,
		bl, wh, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, wh, bl,
		bl, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, wh, bl,
		wh, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, bl, wh,
	}
	buf2 := make([]byte, 2*len(buf))
	buf3 := make([]byte, 3*len(buf))
	buf4 := make([]byte, 4*len(buf))

	for i, w := range buf {
		buf2[i*2+0] = byte(w >> 8)
		buf2[i*2+1] = byte(w)

		buf3[i*3+0] = byte(w>>11) << 3
		buf3[i*3+1] = byte(w>>5&0x3f) << 2
		buf3[i*3+2] = byte(w&0x1f) << 3

		buf4[i*4+0] = buf3[i*3+0]
		buf4[i*4+1] = buf3[i*3+1]
		buf4[i*4+2] = buf3[i*3+2]
		buf4[i*4+3] = 255
	}

	img16 := &pixd.RGB16{
		Rect:   image.Rectangle{Max: image.Point{18, 18}},
		Stride: 2 * 18,
		Pix:    buf2,
	}
	img24 := &pixd.RGB16{
		Rect:   img16.Bounds(),
		Stride: 3 * 18,
		Pix:    buf3,
	}
	img32 := &image.RGBA{
		Rect:   img24.Bounds(),
		Stride: 4 * 18,
		Pix:    buf4,
	}
	_ = img32

	drv := ili9341.New(tftdci.NewSPI(spidrv, dc))
	drv.Init(false)
	drv.SetMADCTL(ili9341.BGR | ili9341.MX)
	disp := pixd.NewDisplay(drv)

	a := disp.NewArea(disp.Bounds())
	a.SetColorRGB(77, 78, 79)
	a.Fill(a.Bounds())
	a.SetRect(disp.Bounds().Sub(image.Point{20, 20}))
	r := a.Bounds()
	t1 := time.Now()
	for i := 0; i < 50; i++ {
		rnd := rand.Int()
		x, rnd := rnd%r.Max.X, rnd/r.Max.X
		y, rnd := rnd%r.Max.Y, rnd/r.Max.Y
		ra, rnd := rnd&127, rnd>>7
		rb := rnd & 127
		rnd = rand.Int()
		a.SetColorRGB(uint8(rnd&^7), uint8(rnd>>8&^3), uint8(rnd>>16&^7))
		a.FillEllipse(image.Pt(x, y), ra, rb)
	}
	t2 := time.Now()
	println(t2.Sub(t1))

	a.SetColorRGB(24, 128, 40)
	for {
		a.Fill(a.Bounds())
		t1 = time.Now()
		for i := 0; i < 1000; i++ {
			rnd := rand.Int()
			x, rnd := rnd%r.Max.X, rnd/r.Max.X
			y, rnd := rnd%r.Max.Y, rnd/r.Max.Y
			a.Draw(
				image.Rectangle{image.Pt(x, y), image.Pt(x+18, y+18)},
				img16, image.Point{},
				nil, image.Point{},
				draw.Src,
			)
		}
		t2 = time.Now()
		println(t2.Sub(t1))
		time.Sleep(2 * time.Second)
	}
}
