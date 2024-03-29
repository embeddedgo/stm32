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

	_ "github.com/embeddedgo/stm32/devboard/minipro-f405/board/system"
)

var prompt = "date!> "

func main() {
	t0, _ := rtcst.LoadTime(0)
	u0 := time.Unix(0, 0)
	if !t0.Equal(u0) {
		time.Set(u0, t0)
		timeSet()
	}
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Buffer(nil, 256)
		fmt.Print("Type help for help.\n\n", prompt)
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
		{"date", date, "print or set the system date and time"},
		{"help", help, "print command description"},
		{"relay", relay, "control the relays"},
		{"sun", date, "print sunrise and sunset time"},
	}
}

func runCmd(args []string) {
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
