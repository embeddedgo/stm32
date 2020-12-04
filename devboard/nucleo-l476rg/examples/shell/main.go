// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	_ "github.com/embeddedgo/stm32/devboard/nucleo-l476rg/board/init"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Buffer(nil, 256)
		fmt.Print("\n\nSimple shell. Type help for more information.\n\n> ")
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
		{"write", write, "write file"},
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

func help(args []string) {
	for i := 0; i < len(commands); i++ {
		cmd := &commands[i]
		if len(args) == 2 {
			if cmd.name == args[1] {
				fmt.Printf("Usage: ")
				cmd.f(nil)
				return
			}
			continue
		}
		fmt.Printf("%-8s %s\n", cmd.name, cmd.brief)
	}
}

func isErr(err error) bool {
	if err == nil {
		return false
	}
	fmt.Fprintf(os.Stderr, "\n%v\n", err)
	return true
}
