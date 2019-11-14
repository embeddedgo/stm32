// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package exti

import (
	"mmio"

	"stm32/hal/raw/afio"
	"stm32/hal/raw/rcc"
)

func exticr(n int) *mmio.U32 {
	return (*mmio.U32)(&afio.AFIO.EXTICR[n].U32)
}

func exticrEna() {
	rcc.RCC.AFIOEN().AtomicSet()
	rcc.RCC.APB2ENR.Load()
}

func exticrDis() {
	rcc.RCC.AFIOEN().AtomicClear()
}
