#!/bin/sh

set -e

cd ../../../embeddedgo/stm32/p
rm -rf *

svdxgen github.com/embeddedgo/stm32/p ../svd/*.svd

for p in dma exti flash gpio pwr rcc rtc spi syscfg tim/tim1 tim/tim2 tim/tim3 tim/tim5 tim/tim6; do
	cd $p
	xgen *.go
	for f in *.go; do
		GOTARGET=$(basename $f .go)
		case $GOTARGET in
		stm32f*|stm32l*)
			GOOS=noos GOARCH=thumb $(emgo env GOROOT)/bin/go build -tags $GOTARGET
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


