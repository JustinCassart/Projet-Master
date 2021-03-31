package main

import (
	"fmt"
	"math"
	"stringmatching/algo"
)

func main() {
	var text string = "la belle vache"
	var pattern string = "vache"
	fmt.Println("Recherche de VACHE dans LA BELLE VACHE")
	algo.Shiftand(text, pattern, math.MaxInt64)
	algo.Shiftor(text, pattern)
	text = "baabac"
	pattern = "aaba"
	fmt.Println("Recherche de AABA dans BAABAC")
	algo.Shiftand(text, pattern, math.MaxInt64)
	algo.Shiftor(text, pattern)
	text = "La grenouille mange le grillon. Le grillon mange la fleur"
	pattern = "grillon"
	fmt.Println("Recherhe de GRILLON dans LA GRENOUILLE MANGE LE GRILLON. LE GRILLON MANGE LA FLEUR")
	algo.Shiftand(text, pattern, math.MaxInt64)
	algo.Shiftor(text, pattern)
	text = "ababba"
	pattern = "abba"
	fmt.Println("Recherche pour abba dans ababba")
	fmt.Println("maxsize 2", algo.Shiftand(text, pattern, 2))
	fmt.Println("maxsize 3", algo.Shiftand(text, pattern, 3))
	fmt.Println("maxsize 1", algo.Shiftand(text, pattern, 1))
	fmt.Println("maxsize 10", algo.Shiftand(text, pattern, 10))
}
