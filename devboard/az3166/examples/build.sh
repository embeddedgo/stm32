#!/bin/sh

GOTARGET=stm32f412
GOMEM=0x20000000:256K

. $(emgo env GOROOT)/../scripts/build.sh $@
