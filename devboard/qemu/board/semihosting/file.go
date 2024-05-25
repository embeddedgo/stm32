// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package semihosting provieds access to files located on a debuging host.
// Debuger or emulator like QEMU must support it.
//
// Example QEMU semihosting options:
//
//	-semihosting --semihosting-config enable=on,target=native,userspace=on
//
// https://github.com/ARM-software/abi-aa/blob/main/semihosting/semihosting.rst
package semihosting

import (
	"fmt"
	"io"
	"unsafe"
)

type Error struct {
	no int
}

func (err *Error) Error() string {
	return fmt.Sprint("semihosting error: ", err.no)
}

func hostError() *Error {
	return &Error{hostCall(0x13, nil)}
}

// File represents an open file on debuging host.
type File struct {
	fd int
}

// Fd returns the integer Unix file descriptor referencing the open file.
func (f *File) Fd() int {
	return f.fd
}

// FopenMode describes file read/write mode. It mimics the meaningo of mode
// parameter of C standard library fopen function.
type FopenMode int

const (
	R   FopenMode = 0  // Open text file for reading from beggining of file.
	RB  FopenMode = 1  // Like R but for binary file.
	RW  FopenMode = 2  // Open text file for read/writing at beggining of file.
	RWB FopenMode = 3  // Like RW but for binary file.
	W   FopenMode = 4  // Truncate or create text file for writing.
	WB  FopenMode = 5  // Like W but for binary file.
	WR  FopenMode = 6  // Truncate or create text file for writing and reading.
	WRB FopenMode = 7  // Like WR but for binary file.
	A   FopenMode = 8  // Open or create text file for appending.
	AB  FopenMode = 9  // Like A but for binary file.
	AR  FopenMode = 10 // Open or create text file for appending and reading.
	ARB FopenMode = 11 // Like Ap but for binary file.
)

// OpenFile opens the named file for operations specified by mode. Use name
// ":tt" to read/write from/to standard input/output.
func OpenFile(name string, mode FopenMode) (f *File, err error) {
	type oargs struct {
		path    *byte
		mode    FopenMode
		pathLen int
	}
	args := &oargs{
		unsafe.StringData(name + "\x00"),
		mode,
		len(name),
	}
	fd := hostCall(0x01, unsafe.Pointer(args))
	if fd == -1 {
		err = hostError()
		return
	}
	f = &File{fd}
	return
}

// Close closes file.
func (f File) Close() error {
	if hostCall(0x02, unsafe.Pointer(&f.fd)) == -1 {
		return hostError()
	}
	return nil
}

type rwargs struct {
	fd int
	p  *byte
	n  int
}

// WriteString implements io.StringWriter interface.
func (f *File) WriteString(s string) (n int, err error) {
	if len(s) == 0 {
		return
	}
	args := &rwargs{
		f.fd,
		unsafe.StringData(s),
		len(s),
	}
	notWritten := hostCall(0x05, unsafe.Pointer(args))
	n = len(s) - notWritten
	if notWritten != 0 {
		err = hostError()
	}
	return
}

// Write implements io.Writer interface.
func (f *File) Write(p []byte) (int, error) {
	return f.WriteString(*(*string)(unsafe.Pointer(&p)))
}

// Read implements io.Reader interface.
func (f *File) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}
	args := &rwargs{
		f.fd,
		unsafe.SliceData(p),
		len(p),
	}
	notRead := hostCall(0x06, unsafe.Pointer(args))
	n = len(p) - notRead
	if n == 0 {
		err = io.EOF
	}
	return
}
