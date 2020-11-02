#!/bin/sh

INTERFACE=stlink
TARGET=stm32f4x
RESET=srst_only
TRACECLKIN=120000000

. ../../../../../scripts/swo-oocd.sh
