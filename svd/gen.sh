#!/bin/sh

set -e

cd ../../../embeddedgo/stm32/p
rm -rf *

svdxgen github.com/embeddedgo/stm32/p ../svd/*.svd

for p in dma dmamux dmamux/dmamux1 dmamux/dmamux2 exti flash gpio pwr rcc rtc spi syscfg tim; do
	cd $p
	xgen -g *.go
	for f in *.go; do
		GOTARGET=$(basename $f .go)
		case $GOTARGET in
		stm32f*|stm32l*|stm32h*)
			GOTOOLCHAIN=go1.24.5-embedded GOOS=noos GOARCH=thumb go build -tags $GOTARGET
			;;
		esac
	done
	cd $OLDPWD
done

rm -f ../hal/irq/*

awkscript='{
	gsub("package irq", "package irq\n\nimport \"embedded/rtos\"", $0)
	gsub(" = ", " rtos.IRQ = ", $0)
	print
}'
cd irq
for f in *; do
	awk "$awkscript" $f >../../hal/irq/$f
done


