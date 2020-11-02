#!/bin/sh

INTERFACE=stlink
TARGET=stm32l4x
TRACECLKIN=80000000
RESET=srst_only

. ../../../../../scripts/swo-oocd.sh
