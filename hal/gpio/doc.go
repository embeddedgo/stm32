// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gpio provides interface to configure and control GPIO ports, pins and
// alternate functions.
//
// This package handles the whole STM32 family in uniform way, even though the
// GPIO design used in STM32F1xx series is significantly different than the one
// used in newer series.
package gpio
