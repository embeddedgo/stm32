#!/bin/sh

TARGET=stm32h7x_dual_bank
TRACECLKIN=200000000
RESET=none

. $(emgo env GOROOT)/../scripts/load-oocd.sh
