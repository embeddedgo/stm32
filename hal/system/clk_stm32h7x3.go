// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build stm32h7x3

package system

const (
	HSIClk = 64e6 // Hz
	CSIClk = 4e6  // Hz
)

const (
	// VOS0
	maxSysClk0 = 480e6 // Hz
	maxAHBClk0 = 240e6 // Hz
	maxAPBClk0 = 120e6 // Hz

	// VOS1
	maxSysClk1 = 400e6 // Hz
	maxAHBClk1 = 200e6 // Hz
	maxAPBClk1 = 100e6 // Hz

	// VOS2
	maxSysClk2 = 300e6 // Hz
	maxAHBClk2 = 150e6 // Hz
	maxAPBClk2 = 75e6  // Hz

	// VOS3
	maxSysClk3 = 200e6 // Hz
	maxAHBClk3 = 100e6 // Hz
	maxAPBClk3 = 50e6  // Hz
)
