// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtcst

import (
	"testing"
	"time"
)

func TestMonoday(t *testing.T) {
	date := time.Date(2000, 3, 1, 0, 0, 0, 0, time.UTC)
	start := date.Unix()
	for i := 0; i <= 100*365; i++ {
		y, m, d := date.Date()
		d1 := (date.Unix() - start) / (24 * 60 * 60)
		d2 := int64(monoday(uint(y-2000), uint(m), uint(d)))
		if d1 != d2 {
			t.Errorf("%v: %d != %d", date, d1, d2)
		}
		date = date.Add(24 * time.Hour)
	}
}

func TestDayofmonth(t *testing.T) {
	date := time.Date(2000, 3, 1, 0, 0, 0, 0, time.UTC)
	start := date.Unix()
	for i := 0; i < 100*365+24; i++ {
		d1 := uint(date.Day())
		d2 := dayofmonth(uint(date.Unix()-start) / (24 * 60 * 60))
		if d1 != d2 {
			t.Errorf("%v: %d != %d", date, d1, d2)
		}
		date = date.Add(24 * time.Hour)
	}
}
