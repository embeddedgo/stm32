#!/bin/sh

GOTARGET=stm32f412
GOMEM=0x20000000:256K
GOTEXT=0x8000000

. ../../../../../scripts/build.sh $@
