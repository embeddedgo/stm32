// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

const rmUsage = `
rm NAME
`

func rm(args []string) {
	if len(args) != 2 {
		fmt.Print(rmUsage)
		return
	}
	isErr(os.Remove(args[1]))
}
