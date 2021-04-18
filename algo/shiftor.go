package algo

import (
	"math"
	"stringmatching/utils"
)

func PreShiftOr(pattern string) utils.Mask {
	mask := utils.CreateMask(true, len(pattern))
	for i := 0; i < len(pattern); i++ {
		var value uint = 1 << i
		value ^= mask.Get(pattern[i])
		mask.Set(pattern[i], value)
	}
	return mask
}

func ShiftOr(text string, pattern string, maxsize int) []int {
	// D'abord on calcule le nombre de sous-motifs nécessaire
	// En fonction de la taille maximale
	l := float64(len(pattern))
	m := float64(maxsize)
	n := l / m
	// On crée une liste pour les mask de chaque sous-motifs
	masks := make([]utils.Mask, int(math.Ceil(n)))
	index := 0
	for i := 0; i < len(pattern); i += maxsize {
		// On calcule chaque masque pour chaque sous-motif
		mask := PreShiftOr(pattern[i:int(math.Min(float64(i+maxsize), float64(len(pattern))))])
		masks[index] = mask
		index++
	}
	return shiftor(masks, text)
}

// nextState permet de calculer l'état D suivant
// selon le masque T utilisé (mask)
// le caractère lu dans le texte (char)
// et l'état courant (state)
func nextState1(mask utils.Mask, char byte, state uint) uint {
	var reset uint = 1 << mask.Size()
	state <<= 1 // Propagation valeur précédente
	if state >= reset {
		state ^= reset // Suppression du 1 propagé plus loin que la longueur du pattern (1111 << 1 = 1110)

	}
	state |= mask.Get(char) // Comparaison avec la valeur courante du texte
	return state
}

// Shiftor is TODO
func shiftor(masks []utils.Mask, text string) []int {
	occ := []int{}
	var d uint = masks[0].Default()
	var match uint = 1 << (masks[0].Size() - 1)
	for i := 0; i < len(text); i++ {
		d = nextState1(masks[0], text[i], d)
		if d < match {
			if len(masks) == 1 {
				occ = append(occ, i-masks[0].Size()+1)
			} else {
				check := occand(masks, text, 1, i+1)
				if check > 0 {
					occ = append(occ, i-masks[0].Size()+1)
				}
			}
		}
	}
	return occ
}

// checkocc permet de vérifier si
// lorsque le premier motif donne une occurrence
// les autres motifs induisent aussi l'occurrence
func occand(masks []utils.Mask, text string, id int, begin int) int {
	if begin >= len(text) {
		return 0
	}
	// On a besoin d'un nouveau point de départ
	// Suppons que le masque soit formé de deux bits
	// Par exemple T[a] = 01 et T[b] = 10
	// Alors le masque D de départ est 100
	// Ainsi lors du décalage d >>= 1
	// On aura bien 10
	var d uint = masks[id].Default()
	var match uint = 1 << (masks[id].Size() - 1)
	for i := 0; i < masks[id].Size(); i++ {
		begin += i
		d = nextState1(masks[id], text[begin], d)
		if d < match {
			if id != len(masks)-1 {
				// Il reste encore des sous-motifs
				// Il faut donc encore vérifier
				// Qu'ils induisent bien une occurrence
				sub := occand(masks, text, id+1, begin+1)
				if sub == 0 {
					// Aucune occurrence n'est trouvée
					return 0
				}
				return sub + masks[id].Size()
			}
			return masks[id].Size()
		}
	}
	return 0
}
