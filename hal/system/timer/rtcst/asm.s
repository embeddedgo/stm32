// +build noos

#include "textflag.h"

#define ICSR_ADDR 0xE000ED04
#define ICSR_PENDSVSET (1<<28)

TEXT ·schedule(SB),NOSPLIT|NOFRAME,$0-0
	MOVW  $ICSR_ADDR, R0
	MOVW  $ICSR_PENDSVSET, R1
	MOVW  R1, (R0)
	SEV   // see ARM Errata 563915
	RET
