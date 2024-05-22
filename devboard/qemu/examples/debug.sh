#!/bin/sh

name=$(basename $(pwd))

# -nographic
# -monitor none

qemu-system-arm -machine netduinoplus2 -cpu cortex-m4 -gdb tcp::3333 -S -kernel $name.elf
