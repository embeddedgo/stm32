// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"fmt"
	"io"
	"os"

	"github.com/embeddedgo/fs/semihostfs"
)

func panicErr(what string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%s: %v\n", what, err))
	}
}

func fatalErr(what string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", what, err)
		os.Stderr.Close()
	}
}

func debug(fd int, p []byte) int {
	n, _ := os.Stderr.Write(p)
	return n
}

func main() {
	fsys := semihostfs.New("SH1", "shdir")
	rtos.Mount(fsys, "/host")

	var err error
	os.Stderr, err = os.Open("/host/:stderr")
	panicErr(":stderr", err)
	rtos.SetSystemWriter(debug)
	os.Stdout, err = os.Open("/host/:stdout")
	fatalErr(":stdout", err)
	os.Stdin, err = os.Open("/host/:stdin")
	fatalErr(":stdin", err)

	fmt.Fprintln(os.Stderr, "Hello over stderr!")

	var x float64
	fmt.Print("Enter a number: ")
	fmt.Scanf("%g", &x)

	fname := "/host/x.txt"

	fmt.Println("Writing", x, "to the", fname, "file.")
	f, err := os.Create(fname)
	fatalErr("create", err)
	_, err = fmt.Fprintln(f, x)
	fatalErr("write", err)
	err = f.Close()
	fatalErr("close", err)

	fmt.Println("Reading from the", fname, "file:")
	f, err = os.Open(fname)
	fatalErr("open", err)
	_, err = io.Copy(os.Stdout, f)
	fatalErr("copy", err)
	err = f.Close()
	fatalErr("close", err)

	fmt.Println("Delete", fname)
	err = os.Remove(fname)
	fatalErr("remove", err)

	fmt.Println("Exit")
	os.Stderr.Close() // graceful exit
}
