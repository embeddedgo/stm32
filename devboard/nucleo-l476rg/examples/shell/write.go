// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
)

const writeUsage = `write FILENAME

Write writes text from standard input to a file. Use CTRL+D to exit.
`

func write(args []string) {
	if len(args) != 2 {
		fmt.Print(writeUsage)
		return
	}
	f, err := os.OpenFile(args[1], os.O_CREATE, 0666)
	if isErr(err) {
		return
	}
	defer f.Close()
	buf := make([]byte, 64)
	for {
		n, err1 := os.Stdin.Read(buf)
		_, err2 := f.Write(buf[:n])
		if isErr(err2) || err1 == io.EOF || isErr(err1) {
			return
		}
	}
}
