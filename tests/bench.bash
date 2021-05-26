#!/bin/bash
# $1 -> taille de l'alphabet
# $2 -> taille min du motif
# $3 -> taille max du motif
# $4 -> step
# $5 -> sous-dossier dans Result

for ((j=$2; j<=$3; j+=$4))
do
    echo $1 $j 
    go test -run=NONE -bench=PatternEqualText1 -benchmem -args $1 $j $j >> ../Results/$5/same1\_$1\_$j\_$j.text
    go test -run=NONE -bench=PatternEqualText2 -benchmem -args $1 $j $j >> ../Results/$5/same2\_$1\_$j\_$j.text
done