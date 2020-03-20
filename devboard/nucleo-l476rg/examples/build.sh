#!/bin/sh

GOTARGET=stm32l4x6
GOMEM=0x20000000:96K,0x10000000:32K
GOTEXT=0x8000000

. ../../../../../scripts/build.sh $@
