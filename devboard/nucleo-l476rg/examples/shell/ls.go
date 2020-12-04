// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

const lsUsage = `
ls DIR

List directory contents.
`

func ls(args []string) {
	if len(args) != 2 {
		fmt.Print(lsUsage)
	}
	f, err := os.Open(args[1])
	if isErr(err) {
		return
	}
	defer isErr(f.Close())
	list, err := f.Readdir(-1)
	if isErr(err) {
		return
	}
	for _, fi := range list {
		fmt.Printf("%v %s\n", fi.Mode(), fi.Name())
	}
}
