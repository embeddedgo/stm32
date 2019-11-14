package gpio

import (
	"unsafe"

	"github.com/embeddedgo/stm32/p/mmap"
)

const pstep = mmap.GPIOB_BASE - mmap.GPIOA_BASE

func portnum(p *Port) uint {
	return uint((uintptr(unsafe.Pointer(p)) - mmap.GPIOA_BASE) / pstep)
}
