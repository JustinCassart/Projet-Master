#!/bin/bash
for ((i=2; i<=26; i*=3))
do
    for ((j=5; j<=5000; j*=10))
    do
        for ((k=$j; k<=50000; k*=10))
        do
            echo $i $j $k $step
            go test -run=NONE -bench=PatternEqualText1 -benchmem -args $i $j $k >> ../Results/Search/patternequaltext1\_$i\_$j\_$k.text
        done
    done
done