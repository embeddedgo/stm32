#!/bin/sh

GOTARGET=stm32f407
GOTEXT=0x8000000
GOMEM=0x20000000:128K,0x10000000:64K
GOSTRIPFN=1

. ../../../../../scripts/build.sh $@
