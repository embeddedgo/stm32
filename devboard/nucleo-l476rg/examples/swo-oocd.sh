#!/bin/sh

TARGET=stm32l4x
TRACECLKIN=80000000
RESET=srst_only

. $(emgo env GOROOT)/../scripts/swo-oocd.sh
