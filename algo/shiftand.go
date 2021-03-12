package algo

import (
	"fmt"
	"stringmatching/utils"
)

func prepoc(pattern string) utils.Mask {
	mask := utils.CreateMask(len(pattern), false)
	for i := 0; i < len(pattern); i++ {
		var value uint = 1 << (len(pattern) - 1 - i)
		value |= utils.Get(mask, pattern[i])
		utils.Set(mask, pattern[i], value)
	}
	return mask
}

// Shiftand is TODO
func Shiftand(text string, pattern string) {
	mask := prepoc(pattern)
	var reset uint = 1 << (len(pattern) - 1)
	var d uint
	for i := 0; i < len(text); i++ {
		d >>= 1                       // Propagation valeur précédente
		d |= reset                    // Ajout du 1 au début (0000 >> 1 = 1000)
		d &= utils.Get(mask, text[i]) // Comparaison avec la valeur courante du texte
		if (d & 1 << len(pattern)) != 0 {
			fmt.Println("Occurence en position", i-len(pattern)+1)
		}
	}
	fmt.Println("Recherche terminée")
}
