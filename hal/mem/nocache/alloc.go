// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package nocache allows to allocate variables and slices of arbitrary type in
// the DMA-capable non-cached memory. The allocation is permanent (there is no
// way to free allocated memory).
//
// The size of the non-cached memory is MCU specific (256 KB in STM32H743).
//
// If the MCU doesn't use data cache and there is no other memory available
// in the system then normal heap allocation is used.
package nocache

import (
	"math/bits"
	"unsafe"
)

// New works like new(T) but guarantees that the allocated variable has the
// given alignmet. Align must be a power of two.
func New[T any](align uintptr) (ptr *T) {
	if bits.OnesCount32(uint32(align)) != 1 {
		panic("bad align")
	}
	if a := unsafe.Alignof(*ptr); a > align {
		align = a
	}
	return (*T)(alloc(align, unsafe.Sizeof(*ptr)))
}

// MakeSlice works like make([]T, len, cap) but guarantees that the returned
// slice has the given alignmet. Align must be a power of two.
func MakeSlice[T any](align uintptr, len, cap int) (slice []T) {
	if bits.OnesCount32(uint32(align)) != 1 {
		panic("bad align")
	}
	if a := unsafe.Alignof(slice[0]); a > align {
		align = a
	}
	ptr := alloc(align, unsafe.Sizeof(slice[0])*uintptr(cap))
	return unsafe.Slice((*T)(ptr), cap)[:len]
}
