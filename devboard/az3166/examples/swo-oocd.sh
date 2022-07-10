#!/bin/sh

TARGET=stm32f4x
RESET=srst_only
TRACECLKIN=120000000

. $(emgo env GOROOT)/../scripts/swo-oocd.sh
