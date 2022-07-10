#!/bin/sh

TARGET=stm32f4x
TRACECLKIN=168000000
RESET=srst_only

. $(emgo env GOROOT)/../scripts/load-oocd.sh
