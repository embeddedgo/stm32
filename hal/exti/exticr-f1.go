// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package exti

import (
	"stm32/hal/raw/afio"
	"stm32/hal/raw/rcc"
)

func exticr(n int) *mmio.R32[uint32] {
	return &afio.AFIO.EXTICR[n]
}

func exticrEna() {
	rcc.RCC.AFIOEN().AtomicSet()
	rcc.RCC.APB2ENR.Load()
}

func exticrDis() {
	rcc.RCC.AFIOEN().AtomicClear()
}
