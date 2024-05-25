// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package semihosting provieds access to files located on a debuging host.
// Debuger or emulator like QEMU must support it.
//
// Example QEMU semihosting options:
//
// 	 -semihosting --semihosting-config enable=on,target=native,userspace=on
//
// https://github.com/ARM-software/abi-aa/blob/main/semihosting/semihosting.rst
package semihosting
