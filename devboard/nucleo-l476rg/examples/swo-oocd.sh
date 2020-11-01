#!/bin/sh

INTERFACE=stlink
TARGET=stm32l4x
RESET='srst_only srst_nogate'
TRACECLKIN=80000000

. ../../../../../scripts/swo-oocd.sh
