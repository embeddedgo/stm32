#!/bin/sh

set -e

if [ -z "$OBJCOPY" ]; then
	OBJCOPY=objcopy
fi

name=$(basename $(pwd))
$OBJCOPY -O binary $name.elf $name.bin
stm32flash -R -w $name.bin -b 460800 /dev/ttyACM0

# Example flashing speeds (binary size: 557912 bytes)
#
# 115200 baud:  72s
# 230400 baud:  45s
# 460800 baud:  32s
