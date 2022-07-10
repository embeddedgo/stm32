#!/bin/sh

TARGET=stm32f4x
TRACECLKIN=168000000
RESET=none

. $(emgo env GOROOT)/../scripts/swo-oocd.sh
