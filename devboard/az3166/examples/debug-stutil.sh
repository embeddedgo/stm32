#!/bin/sh

set -e

gdb_cmd=gdb
if command -v arm-none-eabi-gdb; then
	gdb_cmd=arm-none-eabi-gdb
elif command -v gdb-multiarch; then
	gdb_cmd=gdb-multiarch
fi

brkpnt=6
wchpnt=4

setsid st-util >/dev/null 2>&1 </dev/null &

trap /bin/true INT

gdb_cmd --tui \
	-ex 'target extended-remote localhost:4242' \
	-ex 'set mem inaccessible-by-default off' \
	-ex "set remote hardware-breakpoint-limit $brkpnt" \
	-ex "set remote hardware-watchpoint-limit $wchpnt" \
	-ex 'set history save on' \
	-ex 'set history filename ~/.gdb-history-embeddedgo' \
	-ex 'set history size 1000' \
	"$(basename $(pwd)).elf"

killall st-util
