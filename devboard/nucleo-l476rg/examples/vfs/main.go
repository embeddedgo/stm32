// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"fmt"
	"os"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/init"
)

func lsmount() {
	fmt.Printf("prefix                       fstype     fsname       opencnt\n")
	fmt.Printf("------------------------------------------------------------\n")
	for _, mp := range rtos.Mounts() {
		fmt.Printf(
			"%-28s %-10s %-12s %7d\n",
			mp.Prefix, mp.FS.Type(), mp.FS.Name(), mp.OpenCount,
		)
	}
	fmt.Printf("------------------------------------------------------------\n")
}

func main() {
	lsmount()
	os.Open("/dev/console")
	fmt.Println()
	lsmount()
	println("END")
}
