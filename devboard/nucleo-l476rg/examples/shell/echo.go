// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

const echoUsage = `
echo [TEXT ...]
`

func echo(args []string) {
	if len(args) == 1 {
		fmt.Print(echoUsage)
		return
	}
	for _, arg := range args[1:] {
		fmt.Print(arg, " ")
	}
	fmt.Print("\n")
}
