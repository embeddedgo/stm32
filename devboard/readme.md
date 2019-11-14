## Support for STM32 development boards

### The idea

The idea is to simplyfy the use of popular STM32 development boards.

Every board directory contains the `board` package that provides the interface to the peripherals available on the board (for now the support is modest: only LEDs and buttons). The provided interface tries to be simple and uniform accros all development boards with the same or similar peripherals. For example, the `Setup` function allows to configure the whole system for typical usage without going into details about available clock sources and their parameters.

The _examples_ subdirectory as the name suggests includes sample code, but also scripts and configuration that help to build, load and debug.

There is also _doc_ subdirectory that contain useful information and other resources about this development board.

### Supported boards

[emw3162](emw3162): MXCHIP EMW3162 ([STM32F205RGT6](https://www.st.com/en/microcontrollers/stm32f205rg.html) + [BCM43362](http://www.cypress.com/products/wi-fi)), [website](http://en.mxchip.com/product/wifi_product/39)

![EMW3162](emw3162/doc/board.jpg)

[f4-discovery](f4-discovery): STM32F4-Discovery, [website](http://www.st.com/web/catalog/tools/FM116/SC959/SS1532/PF252419)

![STM32F4-Discovery](f4-discovery/doc/board.jpg)

[nucleo-l476rg](nucleo-l476rg]): NUCLEO-L476RG, [website](http://www.st.com/en/evaluation-tools/nucleo-l476rg.html)

![NUCLEO-L476RG](nucleo-l476rg/doc/board.jpg)
