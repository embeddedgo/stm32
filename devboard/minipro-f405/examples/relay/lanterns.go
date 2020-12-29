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

	w := math.Acos(((-0.014485726138606466 - math.Sin(l.latitude)*sind) / (math.Cos(l.latitude) * cosd)))

	// Calculate transit and daytime

	tranUnix := (Jtran - (2440587.5 - 2451545.0)) * (24 * 3600)
	tranSec, tranNsec := math.Modf(tranUnix)

	transit = time.Unix(int64(tranSec), int64(tranNsec)).In(t.Location())
	if w > 0 {
		daytime = time.Duration(w * (float64(24*time.Hour) / math.Pi))
	}
	return
}

type Lanterns struct {
	relays    []gpio.Pin
	latitude  float64
	longitude float64
	reset     chan struct{}
}

func (l *Lanterns) Run() {
	timer := time.NewTimer(0)
	for {
		// turn off the lanterns
		for _, r := range l.relays {
			r.Set()
		}
		now := time.Now()
		transit, daytime := l.TransitDaytime(now)
		sunset := transit.Add(daytime/2 + 15*time.Minute)
		duration := (24*time.Hour - daytime) / 3
		delay := sunset.Sub(now)
		if !timer.Stop() {
			<-timer.C
		}
	again:
		timer.Reset(delay)
		select {
		case <-timer.C:
			if duration == 0 {
				break
			}
			// turn on the lanterns
			for _, r := range l.relays {
				r.Clear()
			}
			delay = duration
			duration = 0
			goto again
		case <-l.reset:
		}
	}
}

func (l *Lanterns) Reset() {
	l.reset <- struct{}{}
}

var lanterns = Lanterns{
	relays:    relays[0:1],
	latitude:  (51 + 43/60.0) * deg,
	longitude: (19 + 38/60.0) * deg,
	reset:     make(chan struct{}, 1),
}
