#!/bin/sh

name=$(basename $(pwd))

# -serial stdio

qemu-system-arm -machine netduinoplus2 -cpu cortex-m4 -nographic -monitor none -serial none --semihosting-config enable=on,target=native,userspace=on -kernel $name.elf
