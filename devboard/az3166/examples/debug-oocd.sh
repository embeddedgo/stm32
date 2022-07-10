#!/bin/sh

TARGET=stm32f4x
RESET=srst_only

. $(emgo env GOROOT)/../scripts/debug-oocd.sh
