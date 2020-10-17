#!/bin/sh

INTERFACE=stlink-dap
TARGET=stm32h7x_dual_bank
TRACECLKIN=200000000
RESET=none

. ../../../../../scripts/load-oocd.sh
