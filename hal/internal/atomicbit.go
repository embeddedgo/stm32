// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"embedded/mmio"
	"sync/atomic"
	"unsafe"
)

//go:nosplit
func AtomicStoreBits[T mmio.T32](r *mmio.R32[T], mask, bits T) {
	bits &= mask
	for {
		old := r.Load()
		if atomic.CompareAndSwapUint32(
			(*uint32)(unsafe.Pointer(r)),
			uint32(old),
			uint32(old&^mask|bits),
		) {
			return
		}
	}
}

/*
//go:nosplit
func AtomicCompareAndSwapBits[T mmio.T32](r *mmio.R32[T], mask, old, new T) bool {
	old &= mask
	new &= mask
	cur := r.Load()
	if cur&mask != old {
		return false
	}
	return atomic.CompareAndSwapUint32(
		(*uint32)(unsafe.Pointer(r)),
		uint32(cur),
		uint32(cur&^mask|new),
	)
}
*/
