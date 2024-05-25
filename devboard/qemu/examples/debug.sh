#!/bin/sh

name=$(basename $(pwd))

qemu-system-arm -machine netduinoplus2 -cpu cortex-m4 -gdb tcp::3333 -S -nographic -monitor none --semihosting-config enable=on,target=native,userspace=on -kernel $name.elf