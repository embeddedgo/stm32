// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build stm32h7x3

package dma

import "sync"

var chanAlloc = struct {
	mask [2]uint8
	mx   sync.Mutex
}{mask: [2]uint8{0xff, 0xff}}

// AllocChannel allocates a free channel (stream) in the controller. It also
// enables clock to the controller if it's the first channel allocated in the
// controller. AllocChannel returns an invalid channel if there is no free
// channel to be allocated. Use Channel.Free to free
// an unused channel.
func (d *Controller) AllocChannel() (ch Channel) {
	sn := 0
	cam := &chanAlloc.mask[dnum(d)]
	chanAlloc.mx.Lock()
	if *cam != 0 {
		mask := uint8(1)
		if *cam == 0xff {
			d.EnableClock(true) // first channel
			*cam = 0xfe
		} else {
			// Find a free channel.
			for *cam&mask == 0 {
				mask <<= 1
				sn++
			}
			*cam &^= mask
		}
		ch = d.Channel(sn, 0)
	}
	chanAlloc.mx.Unlock()
	return
}

// Free
func (c Channel) Free() {
	mask := uint8(1 << csnum(c))
	d := cctrl(c)
	cam := &chanAlloc.mask[dnum(d)]
	chanAlloc.mx.Lock()
	*cam |= mask
	if *cam == 0xff {
		d.DisableClock() // disable unused controller
	}
	chanAlloc.mx.Unlock()
}
