#!/bin/sh

TARGET=stm32l4x
TRACECLKIN=80000000
RESET='srst_only srst_nogate connect_assert_srst'

. $(emgo env GOROOT)/../scripts/load-oocd.sh
