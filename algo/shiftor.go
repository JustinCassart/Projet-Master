package algo

import (
	"fmt"
	"math"
	"math/bits"
	"stringmatching/utils"
)

// compute the masks for a string
// such that its length is equal or lower than 64
func preShiftOr(pattern string) *utils.Mask {
	fmt.Println("len", len(pattern), pattern)
	mask := utils.CreateMask(true, len(pattern))
	for i := 0; i < len(pattern); i++ {
		var value uint = 1 << i
		value ^= mask.Get(pattern[i])
		mask.Set(pattern[i], value)
	}
	return mask
}

// PreShiftOr computes the masks for the pattern
// It returns a slice of mask
// If the length of the pattern is less or equal than 64
// then the slice contains the single mask representing the pattern
// otherwise, the slice contains the mask for each subpattern s
// such tath the length of s is at most 64
func PreShiftOr(pattern string) []*utils.Mask {
	if len(pattern) > 64 {
		// computes the number of sub-patterns
		n := int(math.Ceil(float64(len(pattern)) / bits.UintSize))
		// creates a slice with n masks
		masks := make([]*utils.Mask, n)
		for i := 0; i < n; i++ {
			// begin corresponds to the first index
			// of the factor in the pattern
			begin := i * bits.UintSize
			// end corresponds to the last index + 1
			// of the factor in the pattern
			// it is at most equal to the lengteh of the pattern
			end := int(math.Min(float64((i+1)*bits.UintSize),
				float64(len(pattern))))
			mask := preShiftOr(pattern[begin:end])
			masks[i] = mask
		}
		return masks
	}
	return []*utils.Mask{preShiftOr(pattern)}
}

// ShiftOr performs the shiftor string matching algorithm
// It searches for all occurrences of the pattern in the text
// It returns a slice of integer such that each integer represents
// the position of the first symbol of an occurrence
func ShiftOr(text string, pattern string) []int {
	masks := PreShiftOr(pattern)
	return shiftor(masks, text)
}

// nextOrState calculates the next state of the search
// accordint to the current state, the symbol reading in the text
// and the mask of the current subpattern
func nextState1(mask *utils.Mask, char byte, state uint) uint {
	var reset uint = 1 << mask.Size()
	state <<= 1 // propagation of the previous value of the state
	if state >= reset {
		// We want that 1111 << 1 = 1110
		// and not 1111 << 1 = 11110
		// so the most significant bit must be reset to 0
		state ^= reset
	}
	// The new state is compare with the value in the mask
	state |= mask.Get(char)
	return state
}

// shiftor is the main algorithm performing the search
// It is able to perform the search for one or more subpatterns
func shiftor(masks []*utils.Mask, text string) []int {
	occ := []int{}
	var d uint = masks[0].Default()
	var match uint = 1 << (masks[0].Size() - 1)
	for i := 0; i < len(text); i++ {
		d = nextState1(masks[0], text[i], d)
		if d < match {
			if len(masks) == 1 {
				occ = append(occ, i-masks[0].Size()+1)
			} else {
				// If there are more than one subpattern
				// we must check if the others subpatterns
				// confirm the potential occurrence
				check := occor(masks, text, 1, i+1)
				if check > 0 {
					// The potential occurrence
					// is a real occurrence
					occ = append(occ, i-masks[0].Size()+1)
				}
			}
		}
	}
	return occ
}

// occor checks if a potential occurrence is a real one
func occor(masks []*utils.Mask, text string, id int, begin int) int {
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
		if begin >= len(text) {
			return 0
		}
		d = nextState1(masks[id], text[begin], d)
		if d < match {
			if id != len(masks)-1 {
				// Il reste encore des sous-motifs
				// Il faut donc encore vérifier
				// Qu'ils induisent bien une occurrence
				sub := occor(masks, text, id+1, begin+1)
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
