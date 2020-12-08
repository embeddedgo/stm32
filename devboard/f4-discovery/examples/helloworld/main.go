// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import _ "github.com/embeddedgo/stm32/devboard/f4-discovery/board/init"

func main() {
	for {
		println("Hello, World!")
	}
}
