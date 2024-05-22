#!/bin/sh

name=$(basename $(pwd))

qemu-system-arm -machine netduinoplus2 -cpu cortex-m4 -nographic -monitor none -serial stdio -kernel $name.elf
