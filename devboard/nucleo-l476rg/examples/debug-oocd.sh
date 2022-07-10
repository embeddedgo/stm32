#!/bin/sh

TARGET=stm32l4x
RESET=srst_only

. $(emgo env GOROOT)/../scripts/debug-oocd.sh
