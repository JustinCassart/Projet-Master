package main

import (
	"fmt"
	"math"
	"stringmatching/algo"
)

func compare(text string, pattern string, size int) {
	fmt.Printf("Recherche de %s dans %s (taille max = %d)\n", pattern, text, size)
	fmt.Printf("Occurrence trouvée par shiftand : %v\n", algo.ShiftAnd(text, pattern, size))
	fmt.Printf("Occurrence trouvée par shiftand : %v\n\n", algo.ShiftOr(text, pattern, size))
}

func main() {
	// var text string = "la belle vache"
	// var pattern string = "vache"
	// fmt.Println("Recherche de VACHE dans LA BELLE VACHE")
	// fmt.Println(algo.Shiftand(text, pattern, math.MaxInt64))
	// fmt.Println(algo.Shiftor(text, pattern, math.MaxInt64))
	// text = "baabac"
	// pattern = "aaba"
	// fmt.Println("Recherche de aaba dans baabac")
	// fmt.Println(algo.Shiftand(text, pattern, math.MaxInt64))
	// fmt.Println(algo.Shiftor(text, pattern, math.MaxInt64))
	// text = "baabac"
	// pattern = "aaba"
	// fmt.Println("Recherche de AABA dans BAABAC")
	// fmt.Println(algo.Shiftand(text, pattern, math.MaxInt64))
	// fmt.Println(algo.Shiftor(text, pattern, math.MaxInt64))
	// text = "La grenouille mange le grillon. Le grillon mange la fleur"
	// pattern = "grillon"
	// fmt.Println("Recherhe de GRILLON dans LA GRENOUILLE MANGE LE GRILLON. LE GRILLON MANGE LA FLEUR")
	// fmt.Println(algo.Shiftand(text, pattern, math.MaxInt64))
	// fmt.Println(algo.Shiftor(text, pattern, math.MaxInt64))
	// text = "ababba"
	// pattern = "abba"
	// fmt.Println("Recherche pour abba dans ababba")
	// fmt.Println("maxsize 2", algo.Shiftand(text, pattern, 2))
	// fmt.Println("maxsize 2", algo.Shiftor(text, pattern, 2))
	// fmt.Println("maxsize 3", algo.Shiftand(text, pattern, 3))
	// fmt.Println("maxsize 3", algo.Shiftor(text, pattern, 3))
	// fmt.Println("maxsize 1", algo.Shiftand(text, pattern, 1))
	// fmt.Println("maxsize 1", algo.Shiftor(text, pattern, 1))
	// fmt.Println("maxsize 10", algo.Shiftand(text, pattern, 10))
	// fmt.Println("maxsize 10", algo.Shiftor(text, pattern, 10))
	compare("baabac", "aaba", math.MaxInt64)
	compare("baabac", "aaba", 1)
	compare("baabac", "aaba", 2)
	compare("baabac", "aaba", 3)
	compare("baabac", "aaba", 4)
	compare("baabac", "a", 1)
	compare("baabac", "ba", 1)
}
