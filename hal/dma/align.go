// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dma

import "unsafe"

// An MemAlign is the prefered memory alignment for DMA operations. It's always
// power of 2 and defined in such a way that a fully MemAlign aligned (top and
// bottom) memory region corresponds to the complete data cache lines.
const MemAlign = cacheLineSize

// A CacheMaint indicates whether DMA requires cache maintenance.
const CacheMaint = cacheMaint

// AlignOffsets calculatest the start and end offsets to the MemAlign aligned
// portion of the memory described by ptr and size.
func AlignOffsets(ptr unsafe.Pointer, size uintptr) (start, end uintptr) {
	const alignMask = MemAlign - 1
	p := uintptr(ptr)
	start = -p & alignMask
	end = size - (p+size)&alignMask
	return
}

func alloc(size uintptr) unsafe.Pointer {
	size = (size + (MemAlign - 1)) &^ (MemAlign - 1)
	size += MemAlign // extra space for address alignment
	buf := make([]byte, size)
	addr := uintptr(unsafe.Pointer(&buf[0]))
	addr = (addr + (MemAlign - 1)) &^ (MemAlign - 1)
	return unsafe.Pointer(addr)
}

// New works like new(T) but guarantees that the allocated variable has the
// prefered DMA alignment (see MemAlign).
func New[T any]() (ptr *T) {
	if MemAlign == 1 {
		return new(T)
	}
	return (*T)(alloc(unsafe.Sizeof(*ptr)))
}

// MakeSlice works like make([]T, len, cap) but guarantees that the returned
// slice has the prefered DMA alignment (see MemAlign).
func MakeSlice[T any](len, cap int) (slice []T) {
	if MemAlign == 1 {
		return make([]T, len, cap)
	}
	ptr := alloc(unsafe.Sizeof(slice[0]) * uintptr(cap))
	return unsafe.Slice((*T)(ptr), cap)[:len]
}
