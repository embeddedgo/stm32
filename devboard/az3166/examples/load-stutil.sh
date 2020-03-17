#!/bin/sh

set -e

name=$(basename $(pwd))
arm-none-eabi-objcopy -O binary $name.elf $name.bin
st-flash --reset write $name.bin 0x8000000
rm $name.bin