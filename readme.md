## Support for the STM32 microcontrollers

Embedded Go supports the STM32 microcontrollers based on the ARM Cortex-M3 or higher CPUs (Cortex-M0 is not supported).

As the generated binaries contain the full Go runtime (support for threads, goroutines, garbage collection) and the information about types and functions (required by the reflection, garbage collector and tracebacks) the minimal requirements are quiet high: 1 MiB flash and 128 KiB RAM. This is enough for simple programs that don't import many external packages. But to fully appreciate the Embedded Go programming you need an MCU with at least 2 MiB of flash and 256 KiB of RAM.

For now the preferred MCUs are STM32H743/753 with 2 MiB flash. They are very capable with the fast CPU and 64-bit FPU fully supported by Embedded Go. There are easily available and quiet cheap development boards based on the STM32H743VI MCU. We will use [such board](devboard/devebox-h743/doc) in the *Getting started* section below.

### Prerequisites

1. Go complier.

   You can download it from [go.dev/dl](https://go.dev/dl/).

2. Git command.

   To instll git on Linux use the package manager provided by your Linux distribution (apt, pacman, rpm, ...).

   Windows users may check the [git for Windows](https://gitforwindows.org/) website.

   The Mac users may use the git command provided by the [Xcode](https://developer.apple.com/xcode/) commandline tools. Another way is to use the [Homebrew](https://brew.sh/) package manager.

### Getting started

1. Install the Embedded Go toolchain.

   Make sure the `$GOPATH/bin` directory is in your `PATH`, as tools installed with the `go install` command will be placed here. If you didn't set the `GOPATH` environment variable manually you can find its default value using the `go env GOPATH` command.

   Then install the Embedded Go toolchain using the following two commands:

   ```sh
   go install github.com/embeddedgo/dl/go1.24.5-embedded@latest
   go1.24.5-embedded download
   ```

2. Install egtool.

   ```sh
   go install github.com/embeddedgo/tools/egtool@latest
   ```

3. Create a project directory containing the `main.go` file with your first Go program for the DevEBox STM32H743 board .

   ```go
   package main

   import (
   	"time"

   	"github.com/embeddedgo/stm32/devboard/devebox-h743/board/leds"
   )

   func main() {
   	for {
   		leds.User.Toggle()
   		time.Sleep(time.Second/2)
   	}
   }
   ```

4. Initialize your project.

   ```sh
   go mod init firstprog
   go mod tidy
   ```

5. Copy the [go.env](devboard/devebox-h743/examples/go.env) file suitable for our board.

6. Compile your first program.

   ```sh
   export GOENV=go.env
   go build
   ```

   or

   ```sh
   GOENV=go.env go build
   ```

   or

   ```sh
   egtool build
   ```

   The last one is like `GOENV=go.env go build` but looks for the `go.env` file up the current module directory tree.

7. Connect your board to the computer in the BOOT mode. In case of the DevEBox board use a jumper or any other tool to shorting BTO and 3V3 pins while connecting it to the USB (you can leave these pins shorted while developing and use the reset button to enter the BOOT mode).

8. Load and run.

   ```sh
   egtool load
   ```

9. See the [Embedded Go](https://embeddedgo.github.io/) website for more information.

### Examples

See more example code for [supported develompent boards](devboard).
