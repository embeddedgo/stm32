// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/embeddedgo/stm32/hal/gpio"
	"github.com/embeddedgo/stm32/hal/usart"
	"github.com/embeddedgo/stm32/hal/usart/usart1"

	_ "github.com/embeddedgo/stm32/devboard/devebox-h743/board/system"
)

func main() {
	// GPIO pins assignment
	pa := gpio.PA()
	pa.EnableClock(true)
	tx := pa.Pin(9)
	rx := pa.Pin(10)

	// Configure and enable USART
	u := usart1.Driver()
	u.UsePin(tx, usart.TXD)
	u.UsePin(rx, usart.RXD)
	u.Setup(usart.Word8b, 115200)
	u.EnableTx()
	u.EnableRx(64)

	// Speed test

	u.WriteString("\r\nWrite speed test...\r\n\n")
	time.Sleep(time.Second)

	n := 40
	s := "00000000001111111111222222222233333333334444444444555555555566666666667777777777\r\n"
	br := u.Periph().Baudrate()
	for k := 0; k < 2; k++ {
		t := time.Now()
		for i := 0; i < n; i++ {
			u.WriteString(s)
		}
		dt := int64(time.Now().Sub(t))
		bps := (int64(n*len(s))*1e9 + dt/2) / dt
		fmt.Fprintf(u, "baud rate: %d b/s (%d B/s),  speed: %d b/s (%d B/s)\r\n\n",
			br, br/8, bps*8, bps)
		time.Sleep(2 * time.Second)
	}

	// Read test

	s = "<=[+](*)->0123456789abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ"
	u.WriteString(s)
	u.WriteString(s)
	u.WriteString("\r\n\nRead test. ")
	u.WriteString("Use keyboard or paste some text (eg. the above line).\r\n\n")

	buf := make([]byte, 128)
	for {
		n, err := u.Read(buf)
		if n != 0 {
			fmt.Fprintf(u, "read%3d: %s\r\n", n, buf[:n])
		}
		if err != nil {
			u.WriteString("err: " + err.Error() + "\r\n")
		}
	}
}
