#!/bin/sh

GOTARGET=stm32l4x6
GOMEM=0x20000000:320K
GOTEXT=0x8000000

. ../../../../../scripts/build.sh $@
