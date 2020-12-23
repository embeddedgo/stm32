// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

const timeLayout = "2006-01-02 15:04:05"

func isErr(err error) bool {
	if err == nil {
		return false
	}
	fmt.Fprintf(os.Stderr, "\n%v\n", err)
	return true
}
