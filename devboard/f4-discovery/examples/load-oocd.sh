#!/bin/sh

TARGET=stm32f4x
TRACECLKIN=168000000
RESET='srst_only srst_nogate connect_assert_srst'

. ../../../../scripts/load-oocd.sh
