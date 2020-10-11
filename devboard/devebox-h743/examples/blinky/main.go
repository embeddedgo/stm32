// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "github.com/embeddedgo/stm32/hal/system"

func main() {
	system.SetupPLL(25, 5, 80, 1)
	for i := 0; ;i++ {
		println(i)
	}
}
