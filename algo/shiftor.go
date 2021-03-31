package algo

import (
	"fmt"
	"stringmatching/utils"
)

func preproc(pattern string) utils.Mask {
	mask := utils.CreateMask(len(pattern), true, len(pattern))
	for i := 0; i < len(pattern); i++ {
		var value uint = 1 << i
		value ^= utils.Get(mask, pattern[i])
		utils.Set(mask, pattern[i], value)
	}
	return mask
}

// Shiftor is TODO
func Shiftor(text string, pattern string) {
	mask := preproc(pattern)
	var d uint = utils.Default(mask)
	var match uint = 1 << (len(pattern) - 1)
	var reset uint = 1 << len(pattern)
	for i := 0; i < len(text)-len(pattern)+1; i++ {
		d <<= 1 // Propagation de la valeur précédente
		if d >= reset {
			d ^= reset // Suppression du 1 propagé plus loin que la longueur du pattern (1111 << 1 = 1110)
		}
		d |= utils.Get(mask, text[i])
		if d < match {
			fmt.Println("occurence en position", i-len(pattern)+1)
		}
	}
}
