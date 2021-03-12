package main

import (
	"fmt"
	"stringmatching/algo"
)

func main() {
	var text string = "la belle vache"
	var pattern string = "vache"
	fmt.Println("Recherche de VACHE dans LA BELLE VACHE")
	algo.Shiftand(text, pattern)
	algo.Shiftor(text, pattern)
	text = "baabac"
	pattern = "aaba"
	fmt.Println("Recherche de AABA dans BAABAC")
	algo.Shiftand(text, pattern)
	algo.Shiftor(text, pattern)
	text = "La grenouille mange le grillon. Le grillon mange la fleur"
	pattern = "grillon"
	fmt.Println("Recherhe de GRILLON dans LA GRENOUILLE MANGE LE GRILLON. LE GRILLON MANGE LA FLEUR")
	algo.Shiftand(text, pattern)
	algo.Shiftor(text, pattern)
}
