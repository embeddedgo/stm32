// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/embeddedgo/fs/ramfs"
)

const mountUsage = `
mount
mount FSTYPE(FSARGS) PREFIX

Supported filesystems:

ramfs(SIZE)  filesystem in RAM
`

const head = `
prefix                        fstype     fsname      opencnt available      used
--------------------------------------------------------------------------------
`

func mount(args []string) {
	if len(args) == 1 {
		fmt.Printf(head)
		for _, mp := range rtos.Mounts() {
			_, _, used, max := mp.FS.Usage()
			fmt.Printf(
				"%-29s %-10s %-12s %6d %9d %9d\n",
				mp.Prefix, mp.FS.Type(), mp.FS.Name(), mp.OpenCount, max-used,
				used,
			)
		}
		return
	}
	if len(args) != 3 {
		fmt.Print(mountUsage)
		return
	}
	var fsargs []string
	fstype := args[1]
	prefix := args[2]
	if i := strings.IndexByte(fstype, '('); i > 0 {
		if fstype[len(fstype)-1] != ')' {
			fmt.Fprint(os.Stderr, "argument list must end with ')'")
			return
		}
		fsargs = strings.Split(fstype[i+1:len(fstype)-1], ",")
		fstype = fstype[:i]
	}
	var fsys rtos.FS
	switch fstype {
	case "ramfs":
		size := 1024
		if len(fsargs) > 0 {
			var err error
			size, err = strconv.Atoi(fsargs[0])
			if err != nil {
				fmt.Println("bad FS size:", err)
				return
			}
		}
		fsys = ramfs.New(int64(size))
	default:
		fmt.Fprint(os.Stderr, "unknown file system type: ", fstype,
			".\nUse one of:\n",
			"ramfs(SIZE)\n")
		return
	}
	rtos.Mount(fsys, prefix)
}
