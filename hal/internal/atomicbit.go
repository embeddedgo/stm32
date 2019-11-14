package internal

import (
	"embedded/mmio"
	"sync/atomic"
	"unsafe"
)

//go:nosplit
func AtomicStoreBits(r *mmio.U32, mask, bits uint32) {
	for {
		old := r.Load()
		if atomic.CompareAndSwapUint32(
			(*uint32)(unsafe.Pointer(r)), old, old&^mask|bits&mask,
		) {
			return
		}
	}
}

//go:nosplit
func AtomicSetBits(r *mmio.U32, mask uint32) {
	for {
		old := r.Load()
		if atomic.CompareAndSwapUint32(
			(*uint32)(unsafe.Pointer(r)), old, old|mask,
		) {
			return
		}
	}
}

//go:nosplit
func AtomicClearBits(r *mmio.U32, mask uint32) {
	for {
		old := r.Load()
		if atomic.CompareAndSwapUint32(
			(*uint32)(unsafe.Pointer(r)), old, old&^mask,
		) {
			return
		}
	}
}
