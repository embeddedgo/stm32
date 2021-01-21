// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math"
	"time"

	"github.com/embeddedgo/stm32/hal/gpio"
)

const deg = math.Pi / 180

// TransitDaytime returns the Sun transit time and the daytime for a given
// latiude, longitude (in radians) and t.UTC().Date() day. Based on
// https://en.wikipedia.org/wiki/Sunrise_equation.
func (l *Lanterns) TransitDaytime(t time.Time) (transit time.Time, daytime time.Duration) {
	// Calculate current Julian day

	n := t.Unix()/(24*3600) + (2440588 - 2451545)

	// Mean solar noon

	J := float64(n) - l.longitude/(2*math.Pi)

	// Solar mean anomaly

	M := 357.5291*deg + (0.98560028*deg)*J

	// Equation of the center

	sinm := math.Sin(M)
	sin2m := math.Sin(2 * M)
	sin3m := math.Sin(3 * M)

	C := (1.9148*deg)*sinm + (0.0200*deg)*sin2m + (0.0003*deg)*sin3m

	// Ecliptic longitude

	L := M + C + (180+102.9372)*deg

	// Solar transit

	Jtran := J + 0.0054*sinm - 0.0069*math.Sin(2*L)

	// Declination of the Sun

	sind := math.Sin(L) * 0.39778850739794974
	cosd := math.Sqrt(1 - sind*sind)

	// Hour angle

	cosw := ((l.sinw - math.Sin(l.latitude)*sind) / (math.Cos(l.latitude) * cosd))

	var w float64
	if cosw <= -1 {
		w = math.Pi
	} else if cosw < 1 {
		w = math.Acos(cosw)
	}

	// Calculate transit and daytime

	tranUnix := (Jtran - (2440587.5 - 2451545.0)) * (24 * 3600)
	tranSec, tranNsec := math.Modf(tranUnix)

	transit = time.Unix(int64(tranSec), int64(tranNsec)).In(t.Location())
	daytime = time.Duration(w * (float64(24*time.Hour) / math.Pi))
	return
}

type Lanterns struct {
	relays    []gpio.Pin
	latitude  float64
	longitude float64
	sinw      float64
	reset     chan struct{}
}

func (l *Lanterns) Run() {
	for {
		// turn off the lanterns
		for _, r := range l.relays {
			r.Set()
		}
		var daytime, delay time.Duration
		now := time.Now()
		day := now
		for {
			var transit time.Time
			transit, daytime = l.TransitDaytime(day)
			sunset := transit.Add(daytime / 2)
			delay = sunset.Sub(now)
			if delay >= 0 {
				break
			}
			day = day.Add(24 * time.Hour)
		}
		timer := time.NewTimer(delay)
		duration := (24*time.Hour - daytime) / 3
	again:
		select {
		case <-timer.C:
			if duration == 0 {
				break // turn off the lanterns
			}
			timer.Reset(duration)
			duration = 0
			// turn on the lanterns
			for _, r := range l.relays {
				r.Clear()
			}
			goto again
		case <-l.reset:
			if !timer.Stop() {
				<-timer.C
			}
		}
	}
}

func (l *Lanterns) Reset() {
	l.reset <- struct{}{}
}

const (
	sinGeom  = -0.01453808050249695 // sin(-0.833 * deg)
	sinCivil = -0.10452846326765347 // sin(-6 * deg)
	sinNauti = -0.20791169081775934 // sin(-12 * deg)
	sinAstro = -0.3090169943749474  // sin(-18 * deg)
)

var lanterns = Lanterns{
	relays:    relays[3 : 3+1],
	latitude:  (51 + 43/60.0) * deg,
	longitude: (19 + 38/60.0) * deg,
	sinw:      sinCivil,
	reset:     make(chan struct{}, 1),
}
