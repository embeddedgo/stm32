#!/bin/sh

INTERFACE=stlink
TARGET=stm32f4x
TRACECLKIN=168000000
RESET=srst_only

. ../../../../../scripts/load-oocd.sh
