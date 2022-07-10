#!/bin/sh

TARGET=stm32f2x
TRACECLKIN=120000000
RESET=none

. $(emgo env GOROOT)/../scripts/swo-oocd.sh
