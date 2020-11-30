// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"embedded/rtos"
	"fmt"
	"os"

	"github.com/embeddedgo/fs/ramfs"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/init"
)

const head = `
prefix                       fstype     fsname      opencnt      used available
-------------------------------------------------------------------------------
`

func lsmount() {
	fmt.Printf(head)
	for _, mp := range rtos.Mounts() {
		_, _, used, avail := mp.FS.Usage()
		fmt.Printf(
			"%-28s %-10s %-12s %6d %9d %9d\n",
			mp.Prefix, mp.FS.Type(), mp.FS.Name(), mp.OpenCount, used, avail,
		)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	checkErr(rtos.Mount(ramfs.New(1024), "/tmp"))
	lsmount()
	buf := make([]byte, 64)
	for {
		n, err := os.Stdin.Read(buf)
		checkErr(err)
		fmt.Printf("read: \"%s\"\n", buf[:n])
		for i, b := range buf[:n] {
			fmt.Printf("%d: '%c' %#x\n", i, b, b)
		}
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	checkErr(scanner.Err())
	lsmount()
	os.Open("/dev/console")
	fmt.Println()
	lsmount()
}
