#!/bin/sh

set -e

if [ -z "$STM32TTY" ]; then
	STM32TTY=/dev/ttyUSB0
fi

if [ -z "$STM32BAUD" ]; then
	# Maximum officially supported speed is 115200 baud.
	# Flashing takes:
	# - 2m38s at  57600 (always works)
	# - 1m41s at 115200 (almost always works)
	# - 1m13s at 230400 (sometimes doesn't work)
	# - 0m58s at 460800 (works rarely)
	STM32BAUD=115200
fi

if [ -z "$OBJCOPY" ]; then
	OBJCOPY=objcopy
fi

name=$(basename $(pwd))
$OBJCOPY -O binary $name.elf $name.bin
stm32flash -b $STM32BAUD -R -i rts,-dtr,dtr:-dtr,-rts -w $name.bin $STM32TTY
rm $name.bin
