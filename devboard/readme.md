## Support for STM32 development boards

### Directory structure

Every board directory contains a set of packages (in *board* subdirectory) that provides the interface to the peripherals available on the board (for now the support is modest: only LEDs and buttons).

The board/init package, when imported, configures the whole system for typical usage. If you use any other package from *board* directory the board/init package is imported implicitly to ensure the board is properly configured.

The *examples* subdirectory as the name suggests includes sample code, but also scripts and configuration that help to build, load and debug.

There is also *doc* subdirectory that contain useful information and other resources about this development board.

### Supported boards

[az3166](az3166): MXCHIP Azure IoT DevKit based on [EMW3166](https://en.mxchip.com/productinfo/244866.html) ([STM32F412RG](https://www.st.com/en/microcontrollers/stm32f412rg.html) + [BCM43362](http://www.cypress.com/products/wi-fi) + 2MB QSPI Flash), [website](https://www.mxchip.com/az3166)

![AZ3166](az3166/doc/board.jpg)

[devebox-h743](devebox-h743): DevEBox STM32H743 ([STM32H743VIT6](https://www.st.com/en/microcontrollers-microprocessors/stm32h743vi.html) + 8MB QSPI Flash [W25Q64JVSIQ](https://www.winbond.com/resource-files/w25q64jv%20spi%20%20%20revc%2006032016%20kms.pdf)), [website](https://mcudev.taobao.com/)

![DevEBox STM32H743](devebox-h743/doc/board.jpg)

[emw3162](emw3162): MXCHIP EMW3162 ([STM32F205RGT6](https://www.st.com/en/microcontrollers/stm32f205rg.html) + [BCM43362](http://www.cypress.com/products/wi-fi)), [website](https://en.mxchip.com/productinfo/244895.html)

![EMW3162](emw3162/doc/board.jpg)

[f4-discovery](f4-discovery): STM32F4-Discovery, [website](https://www.st.com/web/catalog/tools/FM116/SC959/SS1532/PF252419)

![STM32F4-Discovery](f4-discovery/doc/board.jpg)

[minipro-f405](minipro-f405): STM32_MiNi_Pro ([STM32F405RGT6](https://www.st.com/en/microcontrollers-microprocessors/stm32f405rg.html) + 2MB SPI Flash [W25Q16JVSIQ](https://www.winbond.com/resource-files/w25q16jv%20spi%20revg%2003222018%20plus.pdf) + USB to UART/RESET/BOOT0 [CH340C](http://www.wch-ic.com/products/CH340.html)), [website](https://mcudev.taobao.com/)

![minipro-f405](minipro-f405/doc/board.jpg)

[nucleo-l476rg](nucleo-l476rg): NUCLEO-L476RG, [website](https://www.st.com/en/evaluation-tools/nucleo-l476rg.html)

![NUCLEO-L476RG](nucleo-l476rg/doc/board.jpg)

[nucleo-l496zg](nucleo-l496zg): NUCLEO-L496ZG, [website](https://www.st.com/en/evaluation-tools/nucleo-l496zg.html)

![NUCLEO-L496ZG](nucleo-l496zg/doc/board.jpg)
