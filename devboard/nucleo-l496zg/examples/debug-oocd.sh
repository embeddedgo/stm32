#!/bin/sh

INTERFACE=stlink
TARGET=stm32l4x
RESET='srst_only srst_nogate'

. ../../../../../scripts/debug-oocd.sh
