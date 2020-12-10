// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

const renameUsage = `
rename OLDNAME NEWNAME

Rename renames the file.
`

func rename(args []string) {
	if len(args) != 3 {
		fmt.Fprint(os.Stdout, renameUsage)
		return
	}
	isErr(os.Rename(args[1], args[2]))
}
