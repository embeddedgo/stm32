package gpio

import (
	"unsafe"

	"github.com/embeddedgo/stm32/p/bus"
	"github.com/embeddedgo/stm32/p/gpio"
	"github.com/embeddedgo/stm32/p/mmap"
)

// Pnum is the number of available GPIO ports.
const Pnum = pnum

// Port represents a GPIO port.
type Port struct {
	registers
}

// P returns n-th GPIO port (0 means port A, 1 means port B, etc.).
func P(n int) *Port {
	if uint(n) > pnum {
		return nil
	}
	return (*Port)(unsafe.Pointer(mmap.GPIOA_BASE + uintptr(n)*pstep))
}

// Num returns the port number: 0 for A, 1 for B, etc.
func (p *Port) Num() int {
	return int(portnum(p))
}

// Bus returns a bus to which p is connected.
func (p *Port) Bus() bus.Bus {
	return (*gpio.Periph)(unsafe.Pointer(p)).Bus()
}

// EnableClock enables clock for port p.
// lp determines whether the clock remains on in low power (sleep) mode.
func (p *Port) EnableClock(lp bool) {
	enableClock(p, lp)
}

// DisableClock disables clock for port p.
func (p *Port) DisableClock() {
	disableClock(p)
}

// Reset resets port p.
func (p *Port) Reset() {
	reset(p)
}

// Pins is a bitmask which represents the pins of a GPIO port.
type Pins uint16

const (
	Pin0 Pins = 1 << iota
	Pin1
	Pin2
	Pin3
	Pin4
	Pin5
	Pin6
	Pin7
	Pin8
	Pin9
	Pin10
	Pin11
	Pin12
	Pin13
	Pin14
	Pin15
)

func (p *Port) Pin(id int) Pin {
	ptr := uintptr(unsafe.Pointer(p))
	return Pin{ptr | uintptr(id&0xf)}
}
