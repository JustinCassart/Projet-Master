package algo

import (
	"math"
	"stringmatching/utils"
)

// preproc permet de créer le mask T
// pour un motif donné
func prepoc(pattern string) utils.Mask {
	mask := utils.CreateMask(len(pattern), false, len(pattern))
	for i := 0; i < len(pattern); i++ {
		var value uint = 1 << (len(pattern) - 1 - i)
		value |= utils.Get(mask, pattern[i])
		utils.Set(mask, pattern[i], value)
	}
	return mask
}

// Shiftand permet de trouver toutes les occurrences
// d'un motif pattern dans un texte text
// La taille maximale du motif est donnée par maxsize
func Shiftand(text string, pattern string, maxsize int) []int {
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
		mask := prepoc(pattern[i:int(math.Min(float64(i+maxsize), float64(len(pattern))))])
		masks[index] = mask
		index++
	}
	return multiShiftand(masks, text)
}

// checkocc permet de vérifier si
// lorsque le premier motif donne une occurrence
// les autres motifs induisent aussi l'occurrence
func checkocc(masks []utils.Mask, text string, id int, begin int) int {
	if begin >= len(text) {
		return 0
	}
	// On a besoin d'un nouveau point de départ
	// Suppons que le masque soit formé de deux bits
	// Par exemple T[a] = 01 et T[b] = 10
	// Alors le masque D de départ est 100
	// Ainsi lors du décalage d >>= 1
	// On aura bien 10
	var d uint = 1 << masks[id].Size
	for i := 0; i < masks[id].Size; i++ {
		begin += i
		d = nextState(masks[id], text[begin], d)
		if (d & 1) != 0 {
			if id != len(masks)-1 {
				// Il reste encore des sous-motifs
				// Il faut donc encore vérifier
				// Qu'ils induisent bien une occurrence
				sub := checkocc(masks, text, id+1, begin+1)
				if sub == 0 {
					// Aucune occurrence n'est trouvée
					return 0
				}
				return sub + masks[id].Size
			}
			return masks[id].Size
		}
	}
	return 0
}

// nextState permet de calculer l'état D suivant
// selon le masque T utilisé (mask)
// le caractère lu dans le texte (char)
// et l'état courant (state)
func nextState(mask utils.Mask, char byte, state uint) uint {
	var reset uint = 1 << (mask.Size - 1)
	state >>= 1                    // Propagation valeur précédente
	state |= reset                 // Ajout du 1 au début (0000 >> 1 = 1000)
	state &= utils.Get(mask, char) // Comparaison avec la valeur courante du texte
	return state
}

// multiShiftand est l'algorithme permettant
// de vraiment trouver les occurrences
func multiShiftand(masks []utils.Mask, text string) []int {
	occ := []int{} // liste contenant les positions des occurrences trouvées
	var d uint     // le masque D initialisé à 0
	for i := 0; i < len(text); i++ {
		d = nextState(masks[0], text[i], d)
		if (d & 1) != 0 {
			if len(masks) == 1 {
				// Le motif tient en un mot machine
				occ = append(occ, i-masks[0].Size+1)
			} else {
				// Dans le cas d'une occurrence
				// Mais que plusieurs mots machines sont utilisés
				// Il faut vérifier que les autres mots induisent
				// Aussi une occurrence
				check := checkocc(masks, text, 1, i+1)
				if check > 0 {
					// Une occurrence a bien été trouée
					occ = append(occ, i-masks[0].Size+1)
				}
			}
		}
	}
	return occ
}
