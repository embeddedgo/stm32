// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !stm32h7x3

package nocache

import "unsafe"

func alloc(align, size uintptr) unsafe.Pointer {
	mask := align - 1
	if mask < 3 {
		mask = 3
	}
	arr := make([]uint32, (size+mask)/4)
	p := unsafe.Pointer(&arr[0])
	o := -uintptr(p) & mask
	return unsafe.Add(p, o)
}
