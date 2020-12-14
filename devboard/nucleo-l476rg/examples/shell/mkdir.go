// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

const mkdirUsage = `
mkdir NAME
`

func mkdir(args []string) {
	if len(args) != 2 {
		fmt.Print(mkdirUsage)
		return
	}
	isErr(os.Mkdir(args[1], 0755))
}
