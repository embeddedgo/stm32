#!/bin/sh

GOTARGET=stm32l4x6
GOMEM=0x20000000:320K

. $(emgo env GOROOT)/../scripts/build.sh $@
