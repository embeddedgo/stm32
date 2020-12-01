// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"embedded/rtos"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/embeddedgo/fs/ramfs"
	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/init"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Buffer(nil, 256)
		fmt.Print("Simple shell. Type help for more information.\n\n> ")
		for scanner.Scan() {
			args := strings.Fields(scanner.Text())
			if len(args) >= 1 {
				runCmd(args)
			}
			fmt.Print("> ")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

type command struct {
	name  string
	f     func(args []string)
	brief string
}

var commands []command

func init() {
	// avoid initialization loop
	commands = []command{
		{"help", help, "print a list of available commands"},
		{"mount", mount, "mount a file system or print all mount points"},
	}
}

var ErrCmdNotFound = errors.New("command not found")

func runCmd(args []string) {
	for i := 0; i < len(commands); i++ {
		cmd := &commands[i]
		if cmd.name == args[0] {
			cmd.f(args)
			return
		}
	}
	fmt.Fprint(os.Stderr, args[0], ": unknown command\n")
}

func help(_ []string) {
	for i := 0; i < len(commands); i++ {
		cmd := &commands[i]
		fmt.Printf("%-8s %s\n", cmd.name, cmd.brief)
	}
}

const head = `
prefix                       fstype     fsname      opencnt      available used
-------------------------------------------------------------------------------
`

func mount(args []string) {
	if len(args) <= 1 {
		fmt.Printf(head)
		for _, mp := range rtos.Mounts() {
			_, _, used, max := mp.FS.Usage()
			fmt.Printf(
				"%-28s %-10s %-12s %6d %9d %9d\n",
				mp.Prefix, mp.FS.Type(), mp.FS.Name(), mp.OpenCount, max-used,
				used,
			)
		}
		return
	}
	if args[1] == "-h" || len(args) != 3 {
		fmt.Print(args[0], " FSTYPE(FSARGS) PREFIX\n")
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
