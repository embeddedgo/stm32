// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/embeddedgo/stm32/hal/system/timer/rtcst"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/init"
)

var prompt = "date!> "

func main() {
	t0, _ := rtcst.LoadTime(0)
	u0 := time.Unix(0, 0)
	if !t0.Equal(u0) {
		time.Set(u0, t0)
		prompt = "> "
	}
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Buffer(nil, 256)
		fmt.Print(
			"\n\nSimple shell. Type help for more information.\n\n",
			prompt,
		)
		for scanner.Scan() {
			args := strings.Fields(scanner.Text())
			if len(args) >= 1 {
				runCmd(args)
			}
			fmt.Print(prompt)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprint(os.Stderr, err)
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
		{"cat", cat, "print files on the standard output"},
		{"date", date, "print or set the system date and time"},
		{"echo", echo, "display a line of tex"},
		{"help", help, "print information about command"},
		{"ls", ls, "list directory content"},
		{"mkdir", mkdir, "make directory"},
		{"mount", mount, "mount a file system"},
		{"rename", rename, "rename file"},
		{"rm", rm, "remove file"},
	}
}

func runCmd(args []string) {
	stdout := os.Stdout
	if m := len(args) - 1; m > 0 && len(args[m]) > 1 && args[m][0] == '>' {
		name := args[m][1:]
		args = args[:m]
		flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
		if len(name) > 1 && name[0] == '>' {
			name = name[1:]
			flags = os.O_WRONLY | os.O_CREATE | os.O_APPEND
		}
		f, err := os.OpenFile(name, flags, 0644)
		if isErr(err) {
			return
		}
		os.Stdout = f
	}
	var cmd *command
	for i := 0; i < len(commands); i++ {
		if commands[i].name == args[0] {
			cmd = &commands[i]
			break
		}
	}
	if cmd == nil {
		fmt.Fprint(os.Stderr, args[0], ": unknown command\n")
	} else {
		cmd.f(args)
	}
	os.Stdout = stdout
}

func help(args []string) {
	for i := 0; i < len(commands); i++ {
		cmd := &commands[i]
		if len(args) == 2 {
			if cmd.name == args[1] {
				cmd.f(nil)
				return
			}
			continue
		}
		fmt.Printf("%-8s %s\n", cmd.name, cmd.brief)
	}
}
