// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
)

const readUsage = `
read FILENAME

Read reads the content of a file and writes it to the standard output.
`

func read(args []string) {
	if len(args) != 2 {
		fmt.Fprint(os.Stdout, readUsage)
		return
	}
	f, err := os.Open(args[1])
	if isErr(err) {
		return
	}
	defer isErr(f.Close())
	buf := make([]byte, 64)
	for {
		n, err1 := f.Read(buf)
		_, err2 := os.Stdout.Write(buf[:n])
		if isErr(err2) || err1 == io.EOF || isErr(err1) {
			return
		}
	}
}
