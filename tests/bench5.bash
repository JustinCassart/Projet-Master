#!/bin/bash
for ((i=2; i<=26; i*=20))
do
    for ((j=1; j<=1000000; j+=5))
    do
        echo $i $j 
        go test -run=NONE -bench=PatternEqualText1 -benchmem -args $i $j $j >> ../Results/Same/same1\_$i\_$j\_$j.text
        go test -run=NONE -bench=PatternEqualText2 -benchmem -args $i $j $j >> ../Results/Same/same2\_$i\_$j\_$j.text
    done
done