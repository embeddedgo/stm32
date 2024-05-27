// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"fmt"
	"path/filepath"

	"github.com/embeddedgo/stm32/devboard/qemu/board/semihosting"
)

var stdin, stdout, stderr *semihosting.File

func fatalErr(what string, err error) {
	if err != nil {
		fmt.Fprintf(stderr, "%s: %v\n", what, err)
		semihosting.Exit()
	}
}

func init() {
	var err error

	stderr, err = semihosting.OpenFile(":tt", semihosting.A)
	if err != nil {
		panic("open stderr: " + err.Error())
	}

	stdout, err = semihosting.OpenFile(":tt", semihosting.W)
	fatalErr("open stdout", err)

	stdin, err = semihosting.OpenFile(":tt", semihosting.R)
	fatalErr("open stdin", err)
}

func main() {
	var buf [80]byte

	stdout.WriteString("enter: ")
	n, err := stdin.Read(buf[:])
	fatalErr("read", err)
	stdout.WriteString("read: ")
	stdout.Write(buf[:n])

	stdout.WriteString("write file\n")
	f, err := semihosting.OpenFile(filepath.Join("shdir", "test.txt"), semihosting.W)
	fatalErr("open file for writting", err)
	_, err = f.WriteString("12345678\n")
	fatalErr("write to file", err)
	f.Close()

	stdout.WriteString("read file: ")
	f, err = semihosting.OpenFile("test.txt", semihosting.R)
	fatalErr("open file for reading", err)
	n, err = f.Read(buf[:])
	fatalErr("read from file", err)
	f.Close()
	stdout.Write(buf[:n])

	semihosting.Exit()
}
