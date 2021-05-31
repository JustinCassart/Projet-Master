package algo

import (
	"stringmatching/utils"
)

// preshiftOr computes the mask for a string
// such that its length is equal or lower than
// the size of a computer word
func preShiftOr(pattern *string) *utils.Mask {
	mask := utils.CreateMask(true, len(*pattern))
	for i := 0; i < len(*pattern); i++ {
		var value uint = 1 << i
		value ^= mask.Get((*pattern)[i])
		mask.Set((*pattern)[i], value)
	}
	return mask
}

// PreShiftOr computes the masks for the pattern
// It returns a slice of masks
// If the length of the pattern is less or equal than a computer word one
// then the slice contains the single mask representing the pattern
// otherwise, the slice contains the mask for each subpattern s
// such tath the length of s is at most 64
func PreShiftOr(pattern *string) []*utils.Mask {
	return utils.CreateMaskBy(preShiftOr, pattern)
}

// ShiftOr performs the shiftor string matching algorithm
// It searches for all occurrences of the pattern in the text
// It returns a slice of integer such that each integer represents
// the position of the first symbol of an occurrence
func ShiftOr(text, pattern string, masks []*utils.Mask) []int {
	occ := []int{}
	var d uint = masks[0].Default()
	var match uint = 1 << (masks[0].Size() - 1)
	for i := 0; i < len(text); i++ {
		nextStateOr(masks[0], text[i], &d)
		if d < match {
			if len(masks) == 1 {
				occ = append(occ, i-masks[0].Size()+1)
			} else {
				if checkocc(masks, &text, 1, i+1, func(match, d uint) bool { return d < match }) {
					occ = append(occ, i-masks[0].Size()+1)
				}
			}
		}
	}
	return occ
}

// nextOrState calculates the next state of the search
// accordint to the current state, the symbol reading in the text
// and the mask of the current subpattern
func nextStateOr(mask *utils.Mask, char byte, state *uint) {
	var reset uint = 1 << mask.Size()
	*state <<= 1 // propagation of the previous value of the state
	if *state >= reset {
		// We want that 1111 << 1 = 1110
		// and not 1111 << 1 = 11110
		// so the most significant bit must be reset to 0
		*state ^= reset
	}
	// The new state is compare with the value in the mask
	*state |= mask.Get(char)
}
