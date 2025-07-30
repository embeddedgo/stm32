#!/bin/sh

set -e

if [ -z "$OBJCOPY" ]; then
	OBJCOPY=objcopy
fi

name=$(basename $(pwd))

egtool bin
dfu-util -a 0 -s 0x08000000 -D $name.bin
rm $name.bin
