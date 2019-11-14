// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f215

// Package otg_fs_pwrclk provides access to the registers of the OTG_FS_PWRCLK peripheral.
//
// Instances:
//  OTG_FS_PWRCLK  OTG_FS_PWRCLK_BASE  -  -  USB on the go full speed
// Registers:
//  0x000 32  FS_PCGCCTL  OTG_FS power and clock gating control register (OTG_FS_PCGCCTL)
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package otg_fs_pwrclk

const (
	STPPCLK  FS_PCGCCTL = 0x01 << 0 //+ Stop PHY clock
	GATEHCLK FS_PCGCCTL = 0x01 << 1 //+ Gate HCLK
	PHYSUSP  FS_PCGCCTL = 0x01 << 4 //+ PHY Suspended
)

const (
	STPPCLKn  = 0
	GATEHCLKn = 1
	PHYSUSPn  = 4
)
