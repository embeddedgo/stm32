// Copyright 2024 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"embedded/mmio"
	"sync"
)

var locks [4]sync.Mutex

func ExclusiveStoreBits[T mmio.T32](r *mmio.R32[T], mask, bits T) {
	mt := &locks[int(r.Addr()>>2)&(len(locks)-1)]
	mt.Lock()
	r.StoreBits(mask, bits)
	mt.Unlock()
}
