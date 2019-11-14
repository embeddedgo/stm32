#!/bin/sh

GOTARGET=stm32f215
GOMEM=0x20000000:128K
GOTEXT=0x8000000
IRQNAMES=../../../../hal/irq

. ../../../../../scripts/build.sh

