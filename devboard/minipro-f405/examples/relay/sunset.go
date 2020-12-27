// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math"
	"time"
)

const (
	deg = math.Pi / 180
	lat = (51 + 43/60.0) * deg
	lng = (19 + 38/60.0) * deg
)

// TransitDaytime returns the Sun transit time and the daytime for a given
// latiude, longitude (in radians) and t.UTC().Date() day. Based on
// https://en.wikipedia.org/wiki/Sunrise_equation.
func transitDaytime(latiude, laongitude float64, t time.Time) (transit time.Time, daytime time.Duration) {
	// Calculate current Julian day

	n := t.Unix()/(24*3600) + (2440588 - 2451545)

	// Mean solar noon

	J := float64(n) - laongitude/(2*math.Pi)

	// Solar mean anomaly

	M := 357.5291*deg + (0.98560028*deg)*J

	// Equation of the center

	C := (1.9148*deg)*math.Sin(M) +
		(0.0200*deg)*math.Sin(2*M) +
		(0.0003*deg)*math.Sin(3*M)

	// Ecliptic longitude

	L := M + C + (180+102.9372)*deg

	// Solar transit

	Jtran := J + 0.0054*math.Sin(M) - 0.0069*math.Sin(2*L)

	// Declination of the Sun

	sind := math.Sin(L) * 0.39778850739794974
	cosd := math.Sqrt(1 - sind*sind)

	// Hour angle

	w := math.Acos(((-0.014485726138606466 - math.Sin(latiude)*sind) / (math.Cos(latiude) * cosd)))

	// Calculate transit and daytime

	tranUnix := (Jtran - (2440587.5 - 2451545.0)) * (24 * 3600)
	tranSec, tranNsec := math.Modf(tranUnix)
	transit = time.Unix(int64(tranSec), int64(tranNsec)).In(t.Location())
	if w > 0 {
		daytime = time.Duration(w * (float64(24*time.Hour) / math.Pi))
	}

	return
}

func sunsetDelay() time.Duration {
	now := time.Now()
	transit, daytime := transitDaytime(lat, lng, now)
	sunset := transit.Add(daytime/2 + 15*time.Minute)
	return sunset.Sub(now)
}

/*
func sunset() {
	var on, off *time.Timer
	for {
		select {
		}
	}
}
*/
