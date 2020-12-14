// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
)

const catUsage = `
cat FILENAME1 [FILENAME2 ...]
`

func cat(args []string) {
	if len(args) < 2 {
		fmt.Print(catUsage)
		return
	}
	for _, name := range args[1:] {
		f, err := os.Open(name)
		if isErr(err) {
			continue
		}
		buf := make([]byte, 64)
		for {
			n, err1 := f.Read(buf)
			_, err2 := os.Stdout.Write(buf[:n])
			if isErr(err2) || err1 == io.EOF || isErr(err1) {
				break
			}
		}
		isErr(f.Close())
	}
}
