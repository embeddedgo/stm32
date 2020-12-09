// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
)

const writeUsage = `
create FILENAME
write  FILENAME

Create/write reads text from the standard input and writes it to a file. Use
CTRL+D to close the file and exit.
`

func write(args []string) {
	if len(args) != 2 {
		fmt.Fprint(os.Stdout, writeUsage)
		return
	}
	flags := os.O_WRONLY
	if args[0] == "create" {
		flags |= os.O_CREATE | os.O_EXCL
	}
	f, err := os.OpenFile(args[1], flags, 0666)
	if isErr(err) {
		return
	}
	buf := make([]byte, 64)
	for {
		n, err1 := os.Stdin.Read(buf)
		_, err2 := f.Write(buf[:n])
		if isErr(err2) || err1 == io.EOF || isErr(err1) {
			break
		}
	}
	isErr(f.Close())
}
