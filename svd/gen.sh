#!/bin/sh

set -e

cd ../../../embeddedgo/stm32/p
rm -rf *

svdxgen github.com/embeddedgo/stm32/p ../svd/*.svd

for p in dma exti flash gpio pwr rcc syscfg; do
	cd $p
	xgen *.go
	cd ..
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


