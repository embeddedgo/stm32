// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package apb

import "unsafe"

const apbBase = 0x40000000

func rnbn(addr unsafe.Pointer) (rn, bn uintptr) {
	n := (uintptr(addr) - apbBase) / 0x400
	return n / 32, n % 32
}
