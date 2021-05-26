#!/bin/bash
for ((i=2; i<=26; i*=3))
do
    for ((j=5; j<=5000; j*=10))
    do
        echo $i $j
        go test -run=NONE -bench=Pre2 -benchmem -args $i $j >> ../Results/PreProcessing/pre2\_$i\_$j.text
    done
done