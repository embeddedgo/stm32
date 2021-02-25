// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This program tests the thread scheduler according to saving/restoring
// a floating-point context.
package main

import (
	"runtime"
	"sync/atomic"
	"time"

	"github.com/embeddedgo/stm32/devboard/devebox-h743/board/leds"
)

var (
	i32 uint32
	f32 = [...]struct{ a, b, c float32 }{
		{2, 1, 0}, {4, 2, 0}, {6, 3, 0}, {8, 4, 0}, {10, 5, 0}, {12, 6, 0},
	}
	i64 uint32
	f64 = [...]struct{ a, b, c float64 }{
		{3, 1, 0}, {6, 2, 0}, {9, 3, 0}, {12, 4, 0}, {15, 5, 0}, {18, 6, 0},
	}
)

func div32() {
	runtime.LockOSThread()
	time.Sleep(100 * time.Millisecond) // avoid StopTheWorld in busy loop
	for {
		for i := range f32 {
			e := &f32[i]
			e.c = e.a / e.b
		}
		for i := range f32 {
			e := &f32[i]
			if e.c != 2 {
				panic("fp32 error")
			}
			e.c = 0
		}
		atomic.StoreUint32(&i32, i32+1)
	}
}

func div64() {
	runtime.LockOSThread()
	time.Sleep(100 * time.Millisecond) // avoid StopTheWorld in busy loop
	for {
		for i := range f64 {
			e := &f64[i]
			e.c = e.a / e.b
		}
		for i := range f64 {
			e := &f64[i]
			if e.c != 3 {
				panic("fp64 error")
			}
			e.c = 0
		}
		atomic.StoreUint32(&i64, i64+1)
	}
}

func blink(n int) {
	for i := 0; i < n; i++ {
		leds.User.SetOn()
		time.Sleep(50 * time.Millisecond)
		leds.User.SetOff()
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	runtime.GOMAXPROCS(3)
	go div32()
	go div64()
	var p32, p64 uint32
	for {
		i := atomic.LoadUint32(&i32)
		if i != p32 {
			p32 = i
			blink(1) // blink once if div32 runs
		}
		time.Sleep(time.Second)
		i = atomic.LoadUint32(&i64)
		if i != p64 {
			p64 = i
			blink(2) // blink twice if div64 runs
		}
		time.Sleep(time.Second)
	}
}
