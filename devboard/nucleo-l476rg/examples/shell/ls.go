// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/fs"
	"os"
)

const lsUsage = `
ls DIR
`

func ls(args []string) {
	if len(args) != 2 {
		fmt.Print(lsUsage)
		return
	}
	f, err := os.Open(args[1])
	if isErr(err) {
		return
	}
	{
		fi, err := f.Stat()
		if isErr(err) {
			goto close
		}
		var list []fs.FileInfo
		if fi.IsDir() {
			list, err = f.Readdir(-1)
			if isErr(err) {
				goto close
			}
		} else {
			list = []fs.FileInfo{fi}
		}
		for _, fi := range list {
			fmt.Printf("%v %7d %s %s\n",
				fi.Mode(), fi.Size(),
				fi.ModTime().Format(timeLayout), fi.Name(),
			)
		}
	}
close:
	isErr(f.Close())
}
