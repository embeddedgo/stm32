#!/bin/sh

INTERFACE=stlink
TARGET=stm32l4x
TRACECLKIN=80000000
RESET='srst_only srst_nogate'

. ../../../../../scripts/load-oocd.sh
