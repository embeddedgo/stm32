#!/bin/sh

GOTARGET=stm32f215
GOMEM=0x20000000:128K

. $(emgo env GOROOT)/../scripts/build.sh $@
