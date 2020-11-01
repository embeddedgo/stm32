#!/bin/sh

INTERFACE=stlink
TARGET=stm32f4x
RESET='srst_only srst_nogate'

. ../../../../../scripts/debug-oocd.sh
