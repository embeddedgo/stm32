#!/bin/sh

# Connect BT0 and 3V3 pins to enable DFU bootloader

set -e

if [ -z "$OBJCOPY" ]; then
	OBJCOPY=objcopy
fi

name=$(basename $(pwd))
$OBJCOPY -O binary $name.elf $name.bin
dfu-util -a 0 -s 0x8000000 -D $name.bin
rm $name.bin
