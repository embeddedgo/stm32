// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtcst

// Monoday returns the day number since 0000-03-01 where yy-mm-dd can be any
// valid date from 0000-03-01 to 0100-02-29.
//
// Monoday is simplified version of the algorithm published by Josh Haberman
// https://blog.reverberate.org/2020/05/12/optimizing-date-algorithms.html
func monoday(yy, mm, dd uint) uint {
	madj := mm - 3 // March 1 will be the day 0 of year
	if madj > mm {
		yy--
		madj += 12
	}
	monthdays := (madj*62719 + 769) >> 11
	leapdays := yy >> 2
	return yy*365 + leapdays + monthdays + dd - 1
}

// Dayofmonth returns the day of month for a given monoday.
//
// Dayofmonth is simplified version of the algorithm published by Howard Hinnant
// http://howardhinnant.github.io/date_algorithms.html#civil_from_days
func dayofmonth(monoday uint) uint {
	leapdays := monoday / 1460
	yy := (monoday - leapdays) / 365
	doy := monoday - (365*yy + yy>>2)
	madj := (535*doy + 332) >> 14
	monthdays := (madj*62719 + 769) >> 11
	return doy - monthdays + 1
}

/*
// Unixday returns the day number since 1970-01-01 UTC for a given UTC date.
// It returns 0 for 1970-01-01, positive numbers for later dates and
// negative numbers for earlier dates (the lower bound is -4799-01-01).
//
// Unixday is based on the algorithm published by Josh Haberman:
// https://blog.reverberate.org/2020/05/12/optimizing-date-algorithms.html
func unixday(year int, month, day uint) int {
	yadj := uint(year + 4800)
	madj := month - 3 // March 1 will be the day 0 of year
	if madj > month {
		yadj--
		madj += 12
	}
	monthdays := (madj*62719 + 769) >> 11
	leapdays := yadj/4 - yadj/100 + yadj/400
	return int(yadj*365+leapdays+monthdays+day) - (2472632 + 1)
}

// Dayofmonth returns the day of month for a given Unix day.
//
// Dayofmonth is based on the algorithm published by Howard Hinnant:
// http://howardhinnant.github.io/date_algorithms.html#civil_from_days
func dayofmonth(unixday int) uint {
	const eradays = 400*365 + 97
	doe := uint(unixday+2472632) % eradays
	leapdays := doe/(eradays/100) - doe/(eradays/4) + doe/(eradays-1)
	yoe := (doe - leapdays) / 365
	doy := doe - (365*yoe + yoe/4 - yoe/100)
	madj := (535*doy + 332) >> 14
	monthdays := (madj*62719 + 769) >> 11
	return doy - monthdays + 1
}
*/
