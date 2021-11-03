// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/embeddedgo/display/pixd"
	"github.com/embeddedgo/display/pixd/driver/tftdrv/ili9486"
	"github.com/embeddedgo/stm32/dci/tftdci"
	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/spi"
	"github.com/embeddedgo/stm32/hal/spi/spi2"

	_ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/init"
)

func rndMax(rnd int64, max int) (int, int64) {
	r := rnd / int64(max)
	return int(rnd - int64(max)*r), r
}

// rndPoint takes 18 LS-bits from rnd to generate random point
func rndPoint(rnd int64, max image.Point) (image.Point, int64) {
	var p image.Point
	p.X, rnd = rndMax(rnd, max.X)
	p.Y, rnd = rndMax(rnd, max.Y)
	return p, rnd
}

// rndRGBA takes 18 LS-bits from rnd to generate r, g, b not greather than a
func rndRGBA(rnd int64, a uint8) (color.RGBA, int64) {
	r, rnd := rndMax(rnd, int(a))
	g, rnd := rndMax(rnd, int(a))
	b, rnd := rndMax(rnd, int(a))
	return color.RGBA{uint8(r), uint8(g), uint8(b), a}, rnd
}

func main() {
	pb := gpio.B()
	pb.EnableClock(true)
	csn := pb.Pin(12)
	sck := pb.Pin(13)
	miso := pb.Pin(14)
	mosi := pb.Pin(15)

	pd := gpio.D()
	pd.EnableClock(true)
	reset := pd.Pin(8)
	dc := pd.Pin(10)

	reset.Clear()
	reset.Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.Low}) // assert RESET
	time.Sleep(time.Millisecond)
	reset.Set() // deassert RESET
	time.Sleep(5 * time.Millisecond)

	spidrv := spi2.Driver()
	spidrv.UsePinMaster(sck, spi.SCK)
	spidrv.UsePinMaster(miso, spi.MISO)
	spidrv.UsePinMaster(mosi, spi.MOSI)
	spidrv.UsePinMaster(csn, spi.NSS) // use hardware slave-select

	dci := tftdci.NewSPI(spidrv, dc, spi.CPOL0|spi.CPHA0, 21e6, 21e6)
	//dci.UseCSN(csn, false) // use software slave-select
	//drv := ili9341.NewOver(dci)
	//drv.Init(ili9341.GFX)
	drv := ili9486.New(dci)
	drv.Init(ili9486.MSP4022)
	disp := pixd.NewDisplay(drv)

	a := disp.NewArea(disp.Bounds())
	a.SetColorRGBA(77, 78, 79, 255)
	a.Fill(a.Bounds())
	a.SetRect(disp.Bounds().Sub(image.Pt(0, 0)))
	r := a.Bounds()

	/*
		a.SetColorRGBA(0, 0, 0, 255)
		a.Fill(a.Bounds())

		a.SetColorRGBA(200, 0, 0, 255)
		a.FillEllipse(image.Pt(160, 110), 100, 100)
		a.SetColorRGBA(255, 0, 0, 255)
		a.DrawEllipse(image.Pt(160, 110), 100, 100)

		a.SetColorRGBA(200, 0, 0, 255)
		a.FillEllipse(image.Pt(160, 220), 9, 9)
		a.SetColorRGBA(255, 0, 0, 255)
		a.DrawEllipse(image.Pt(160, 220), 9, 9)

		a.SetColorRGBA(200, 0, 0, 255)
		a.FillEllipse(image.Pt(160, 240), 8, 8)
		a.SetColorRGBA(255, 0, 0, 255)
		a.DrawEllipse(image.Pt(160, 240), 8, 8)

		a.SetColorRGBA(200, 0, 0, 255)
		a.FillEllipse(image.Pt(160, 260), 7, 7)
		a.SetColorRGBA(255, 0, 0, 255)
		a.DrawEllipse(image.Pt(160, 260), 7, 7)

		a.SetColorRGBA(200, 0, 0, 255)
		a.FillEllipse(image.Pt(160, 280), 6, 6)
		a.SetColorRGBA(255, 0, 0, 255)
		a.DrawEllipse(image.Pt(160, 280), 6, 6)

		a.SetColorRGBA(200, 0, 0, 255)
		a.FillEllipse(image.Pt(160, 300), 5, 5)
		a.SetColorRGBA(255, 0, 0, 255)
		a.DrawEllipse(image.Pt(160, 300), 5, 5)

		a.SetColorRGBA(200, 0, 0, 255)
		a.FillEllipse(image.Pt(160, 320), 4, 4)
		a.SetColorRGBA(255, 0, 0, 255)
		a.DrawEllipse(image.Pt(160, 320), 4, 4)

		a.SetColorRGBA(200, 0, 0, 255)
		a.FillEllipse(image.Pt(160, 340), 3, 3)
		a.SetColorRGBA(255, 0, 0, 255)
		a.DrawEllipse(image.Pt(160, 340), 3, 3)

		a.SetColorRGBA(200, 0, 0, 255)
		a.FillEllipse(image.Pt(160, 360), 2, 2)
		a.SetColorRGBA(255, 0, 0, 255)
		a.DrawEllipse(image.Pt(160, 360), 2, 2)

		a.SetColorRGBA(200, 0, 0, 255)
		a.FillEllipse(image.Pt(160, 380), 1, 1)
		a.SetColorRGBA(255, 0, 0, 255)
		a.DrawEllipse(image.Pt(160, 380), 1, 1)

		time.Sleep(10 * time.Second)
	*/

	for {
		a.SetColorRGBA(0, 0, 0, 255)
		a.Fill(a.Bounds())
		a.SetColorRGBA(255, 0, 0, 255)
		a.Fill(image.Rect(10, 10, 50, 200))
		a.SetColorRGBA(0, 255, 0, 255)
		a.Fill(image.Rect(60, 10, 100, 200))
		a.SetColorRGBA(0, 0, 255, 255)
		a.Fill(image.Rect(110, 10, 150, 200))
		a.SetColorRGBA(128, 0, 128, 128)
		a.RoundRect(image.Pt(10, 50), image.Pt(230, 100), 0, 0, false)
		a.RoundRect(image.Pt(230, 150), image.Pt(10, 200), 0, 0, true)
		time.Sleep(time.Second)

		a.SetColorRGBA(255, 255, 255, 255)
		t1 := time.Now()
		for i := 0; i < 200; i++ {
			rnd := rand.Int63()
			p0, rnd := rndPoint(rnd, r.Max)
			p1, rnd := rndPoint(rnd, r.Max)
			p2, rnd := rndPoint(rnd, r.Max)

			a.SetColorRGBA(uint8(rnd), uint8(rnd), uint8(rnd), 255)
			a.Triangle(p0, p1, p2, true)

			p0, rnd = rndPoint(rnd, r.Max)
			p1, rnd = rndPoint(rnd, r.Max)
			p2, _ = rndPoint(rnd, r.Max)

			rnd >>= 8
			a.SetColorRGBA(uint8(rnd), uint8(rnd), uint8(rnd), 255)
			a.Triangle(p0, p1, p2, true)

			//rnd = rand.Int63()
			//c, rnd := rndRGBA(rnd, 192)

			//a.SetColorRGBA(0, 0, 0, 255)
			//a.Fill(a.Bounds())
			//a.SetColorRGBA(0, 255, 0, 255)
			//a.Triangle(p0, p1, p2, false)
			//a.SetColorRGBA(255, 0, 0, 255)
			//a.Triangle(p0, p1, p2, true)

			//a.Fill(image.Rect(x-w, y-h, x+w, y+h))
			//a.DrawPoint(image.Pt(x, y+h), w)
			//a.Ellipse(image.Pt(x, y+h), w, h, true)
			//a.SetColorRGBA(0, 0, 0, 192)
			//a.Ellipse(image.Pt(x, y+h), w, h, false)

			//r = byte(rnd&31) << 2
			//rnd >>= 5
			//g = byte(rnd&31) << 2
			//rnd >>= 5
			//b = byte(rnd&31) << 2

			//a.SetColorRGBA(r, g, b, 192)
			//a.Fill(image.Rect(x-h, y-w, x+h, y+w))
			//a.DrawPoint(image.Pt(x+w, y), h)
			//a.Ellipse(image.Pt(x+w, y), h, w, true)
			//a.SetColorRGBA(0, 0, 0, 192)
			//a.Ellipse(image.Pt(x+w, y), h, w, false)
		}
		a.Flush()
		t2 := time.Now()
		println(t2.Sub(t1))
	}
	/*
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
			if w == bl {
				buf4[i*4+3] = 0
			} else {
				buf4[i*4+3] = 255
			}
		}

		img16 := &pixd.RGB16{
			Rect:   image.Rectangle{Max: image.Pt(18, 18)},
			Stride: 2 * 18,
			Pix:    buf2,
		}
		img24 := &pixd.RGB{
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
		mask := &image.Uniform{color.Alpha{150}}
		_ = mask

		for {
			a.Fill(a.Bounds())
			t1 := time.Now()
			for i := 0; i < 2000; i++ {
				rnd := rand.Int()
				x, rnd := rnd%r.Max.X, rnd/r.Max.X
				y, rnd := rnd%r.Max.Y, rnd/r.Max.Y
				a.Draw(
					image.Rectangle{image.Pt(x, y), image.Pt(x+18, y+18)},
					img32, image.Point{},
					nil, image.Point{},
					draw.Src,
				)
			}
			t2 := time.Now()
			println(t2.Sub(t1))
			time.Sleep(2 * time.Second)
		}
	*/
}
