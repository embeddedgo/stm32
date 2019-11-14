// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build stm32f215

// Package otg_hs_pwrclk provides access to the registers of the OTG_HS_PWRCLK peripheral.
//
// Instances:
//  OTG_HS_PWRCLK  OTG_HS_PWRCLK_BASE  -  -  USB on the go high speed
// Registers:
//  0x000 32  OTG_HS_PCGCR  Power and clock gating control register
// Import:
//  github.com/embeddedgo/stm32/p/mmap
package otg_hs_pwrclk

const (
	STPPCLK  OTG_HS_PCGCR = 0x01 << 0 //+ Stop PHY clock
	GATEHCLK OTG_HS_PCGCR = 0x01 << 1 //+ Gate HCLK
	PHYSUSP  OTG_HS_PCGCR = 0x01 << 4 //+ PHY suspended
)

const (
	STPPCLKn  = 0
	GATEHCLKn = 1
	PHYSUSPn  = 4
)
