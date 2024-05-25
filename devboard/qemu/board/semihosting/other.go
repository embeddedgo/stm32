// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semihosting

import "unsafe"

//go:noescape
func hostCall(cmd int, arg unsafe.Pointer) int

// SysExit software reasons
const (
	SysExitBreakPoint          = 0x20020
	SysExitWatchPoint          = 0x20021
	SysExitStepComplete        = 0x20022
	SysExitRunTimeErrorUnknown = 0x20023
	SysExitInternalError       = 0x20024
	SysExitUserInterruption    = 0x20025
	SysExitApplicationExit     = 0x20026
	SysExitStackOverflow       = 0x20027
	SysExitDivisionByZero      = 0x20028
	SysExitOSSpecific          = 0x20029
)

// SysExit hardware reasons
const (
	SysExitBranchThroughZero = 0x20000
	SysExitUndefinedInstr    = 0x20001
	SysExitSoftwareInterrupt = 0x20002
	SysExitPrefetchAbort     = 0x20003
	SysExitDataAbort         = 0x20004
	SysExitAddressException  = 0x20005
	SysExitIRQ               = 0x20006
	SysExitFIQ               = 0x20007
)

func SysExit(reason int) {
	hostCall(0x18, unsafe.Pointer(uintptr(reason)))
}

func Exit() {
	hostCall(0x18, unsafe.Pointer(uintptr(SysExitApplicationExit)))
}
