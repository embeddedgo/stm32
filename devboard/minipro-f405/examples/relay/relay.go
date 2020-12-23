// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strings"

	"github.com/embeddedgo/stm32/hal/gpio"
)

var relays = [...]gpio.Pin{
	gpio.C().Pin(0),
	gpio.C().Pin(1),
	gpio.C().Pin(4),
	gpio.C().Pin(5),
}

func init() {
	cfg := &gpio.Config{Mode: gpio.Out, Driver: gpio.OpenDrain, Speed: gpio.Low}
	for _, pin := range relays {
		pin.Port().EnableClock(true)
		pin.Set()
		pin.Setup(cfg)
	}
}

const relayUsage = `
relay [N:{on|off} ...]
`

func relay(args []string) {
	var badArg string
	switch len(args) {
	case 0:
		fmt.Print(relayUsage)
	case 1:
		for i, pin := range relays {
			state := "off"
			if pin.LoadOut() == 0 {
				state = "on"
			}
			fmt.Printf(" %d:%s", i+1, state)
		}
		fmt.Println()
	default:
		for _, a := range args[1:] {
			relay := strings.Split(a, ":")
			if len(relay) != 2 || len(relay[0]) != 1 || len(relay[1]) > 3 {
				badArg = a
				break
			}
			i := int(relay[0][0]) - '1'
			if uint(i) >= uint(len(relays)) {
				badArg = a
				break
			}
			pin := relays[i]
			switch relay[1] {
			case "on":
				pin.Clear()
			case "off":
				pin.Set()
			default:
				badArg = a
				break
			}
		}
	}
	if badArg != "" {
		fmt.Println("bad argument:", badArg)
		fmt.Println()
		fmt.Print(relayUsage)
	}
}
