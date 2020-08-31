#!/bin/sh

stm32flash -b 115200 -R -i rts,-dtr,dtr:-dtr,-rts /dev/ttyUSB0
