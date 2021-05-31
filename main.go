package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"stringmatching/algo"
)

func main() {
	pattern := "Le professeur considéra pendant quelques instants cette série de caractères"
	b, err := ioutil.ReadFile("textes/jules-verne-voyage-au-centre-de-la-terre.txt")
	if err != nil {
		log.Print(err)
	}
	text := string(b)
	fmt.Printf("Recherche de la phrase '%s' dans le livre 'voyage au centre de la terre' de Jules Verne en utilisant différents algorithmes\n", pattern)
	fmt.Println("Le résultat donné est la position du début de l'occurrence")
	fmt.Printf("naive : %v\n", algo.Naive(text, pattern))
	fmt.Printf("kmp   : %v\n", algo.KMP(text, pattern))
	fmt.Printf("kr    : %v\n", algo.KR(text, pattern))
	fmt.Printf("shift-and : %v\n", algo.ShiftAnd(text, pattern, algo.PreShiftAnd(&pattern)))
	fmt.Printf("shift-or  : %v\n", algo.ShiftOr(text, pattern, algo.PreShiftOr(&pattern)))
}
