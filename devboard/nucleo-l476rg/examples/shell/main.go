// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Shell shows how to use the Embedded Go virtual filesystem.
//
// Two different filesystems are used here. First one is the console filesystem
// you can find in fs/termfs package. It implemnts os.Stdin and os.Stdout files
// as the Go runtime expects. The second one is the in RAM filesystem you can
// find in the fs/ramfs package. This program allows to mount a mulitple such
// filesystems and work with files through them (little RAM doesn't allow too
// much).
//
// Since the STM32L476 has little RAM, the board/init package uses a simplified
// console filesystem (termfs.LightFS) by default. The simple console provided
// doesn't support neither CR/LF conversion nor remote echo. Use options of
// your terminal emulator to provide local echo and convert CR and LF chars
// properly. The working example for Linux is:
//
//	picocom -b 115200 --echo --imap lfcrlf --omap crlf /dev/ttyACM0
//
// Example:
//
//	date!> date
//	1970-01-01 00:33:23 UTC
//	date!> help date
//
//	date
//	date YYYY-MM-DD hh:mm:ss
//	date!> date 2023-04-06 13:15:11
//	> date
//	2023-04-06 13:15:19 UTC
//	> help mount
//	mount
//	mount FSTYPE(FSARGS) PREFIX
//	Supported filesystems:
//
//	ramfs(SIZE[,NAME])  in RAM file system
//
//	> mount ramfs(512,FS1) /myfs
//	> mount ramfs(100,FS2) /mnt/fs2
//	> mount
//
//	prefix                    fstype     fsname      opencnt available      used
//	----------------------------------------------------------------------------
//	/dev/console              lterm      USART2            2         0        -1
//	/myfs                     ram        FS1               0       512         0
//	/mnt/fs2                  ram        FS2               0       100         0
//	> echo Hello World >/myfs/file1.txt
//	> mount
//
//	prefix                   fstype     fsname      opencnt available      	used
//	----------------------------------------------------------------------------
//	/dev/console              lterm      USART2            2         0        -1
//	/myfs                     ram        FS1               0       428        84
//	/mnt/fs2                  ram        FS2               0       100         0
//	> ls /myfs
//	-rw-rw-rw-      13 2023-04-06 13:16:25 UTC file1.txt
//	> cat /myfs/file1.txt
//	Hello World
//	>
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
			"\n\nSimple shell. Type help for help.\n\n",
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
		{"help", help, "print command description"},
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
	if os.Stdout != stdout {
		os.Stdout.Close()
		os.Stdout = stdout
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
