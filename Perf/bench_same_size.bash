#!/bin/bash
# $1 -> sous-dossier dans Result
# $2 -> taille de l'alphabet
# $3 -> taille min du motif
# $4 -> taille max du motif
# $5 -> pas
if [ ! -d $1 ]
then
	mkdir -v $1
fi
for ((j=$3; j<=$4; j+=$5))
do
    echo $2 $j 
    go test $PWD/tests -run=NONE -bench=PatternEqualText1 -benchmem -args $2 $j $j >> $1/same1\_$2\_$j\_$j.text
    go test $PWD/tests -run=NONE -bench=PatternEqualText2 -benchmem -args $2 $j $j >> $1/same2\_$2\_$j\_$j.text
done
