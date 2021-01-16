// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/embeddedgo/stm32/hal/system/timer/rtcst"
)

const dateUsage = `
date
date YYYY-MM-DD hh:mm:ss
sun
`

// time.Parse is too heavy
func parse(s, sep string) (v [3]int, err error) {
	fields := strings.Split(s, sep)
	if len(fields) != 3 {
		err = errors.New("bad layout")
		return
	}
	for i, f := range fields {
		var u uint64
		u, err = strconv.ParseUint(f, 10, 16)
		if err != nil {
			return
		}
		v[i] = int(u)
	}
	return
}

func date(args []string) {
	now := time.Now()
	switch len(args) {
	case 3:
		ymd, err := parse(args[1], "-")
		if isErr(err) {
			break
		}
		hms, err := parse(args[2], ":")
		if isErr(err) {
			break
		}
		//offset, err := strconv.ParseInt(args[4], 10, 8)
		//if isErr(err) {
		//	break
		//}
		//time.Local = time.FixedZone(args[3], int(offset))
		t := time.Date(
			ymd[0], time.Month(ymd[1]), ymd[2],
			hms[0], hms[1], hms[2], 0, time.Local,
		)
		t0 := time.Set(now, t)
		rtcst.StoreTime(t0, 0)
		if prompt != "> " {
			timeSet()
		} else {
			lanterns.Reset()
		}
	case 1:
		if args[0] == "date" {
			fmt.Println(now.Format(timeLayout))
		} else {
			for i := 0; i < 2; i++  {
				if i != 0 {
					fmt.Println()
				}
				transit, daytime := lanterns.TransitDaytime(now)
				half := daytime / 2
				fmt.Println("sunrise:", transit.Add(-half).Format(timeLayout))
				fmt.Println("transit:", transit.Format(timeLayout))
				fmt.Println("sunset: ", transit.Add(half).Format(timeLayout))
				now = now.Add(24 * time.Hour)
			}
		}
	default:
		fmt.Print(dateUsage)
	}
}

func timeSet() {
	go lanterns.Run()
	prompt = "> "
}
