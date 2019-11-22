// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import "unsafe"

func BoolUint32(x bool) uint32 {
	return uint32(uint8(*(*uint8)(unsafe.Pointer(&x))))
}
