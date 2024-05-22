// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	_ "github.com/embeddedgo/stm32/devboard/qemu/board/system"
)

func main() {
	for {
		fmt.Println("Hello World!")
	}
}
