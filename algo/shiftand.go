package algo

import (
	"fmt"
	"math"
	"math/bits"
	"stringmatching/utils"
)

// PreShiftOr creates the mask B for a pattern
// such that its length is less or equal to 64
// For example the mask for aaba is
// B[a] = 1011
// B[c] = 0010
func preShiftAnd(pattern string) *utils.Mask {
	mask := utils.CreateMask(false, len(pattern))
	for i := 0; i < len(pattern); i++ {
		var value uint = 1 << i
		value |= mask.Get(pattern[i])
		mask.Set(pattern[i], value)
	}
	return mask
}

// PreShiftAnd computes the masks for the pattern
// It returns a slice of mask
// If the length of the pattern is less or equal than the size of a word
// then the slice contains the single mask representing the pattern
// otherwise, the slice contains the mask for each subpattern s
// such tath the length of s is at most the size of a word
func PreShiftAnd(pattern string) []*utils.Mask {
	return utils.CreateMaskBy(preShiftAnd, pattern)
}

func ShiftAnd(text, pattern string) []int {
	masks := PreShiftAnd(pattern)
	occ := []int{}
	var match uint = 1 << (masks[0].Size() - 1)
	var d uint = 0
	for i := 0; i < len(text); i++ {
		nextStateAnd(masks[0], text[i], &d)
		if d&match != 0 {
			if len(masks) == 1 {
				occ = append(occ, i-masks[0].Size()+1)
			} else {
				if occand(masks, &text, 1, i+1) {
					occ = append(occ, i-masks[0].Size()+1)
				}
			}
		}
	}
	return occ
}

func nextStateAnd(mask *utils.Mask, char byte, state *uint) {
	*state = ((*state << 1) | 1) & mask.Get(char)
}

func occand(masks []*utils.Mask, text *string, id, begin int) bool {
	if begin >= len(*text) {
		return false
	}
	var d uint = 0
	var match uint = 1 << (masks[id].Size() - 1)
	for i := 0; i < masks[id].Size(); i++ {
		if begin+i >= len(*text) {
			return false
		}
		if masks[id].Get((*text)[begin+i]) == 0 {
			return false
		}
		nextStateAnd(masks[id], (*text)[begin+i], &d)
		if d&match != 0 {
			if id != len(masks)-1 {
				return occand(masks, text, id+1, begin+i+1)
			}
			return true
		}
	}
	return false
}

// PreMultiShiftAnd compute the init set
// and the match set such that if we have a set
// P = {P1, P2, ..., Pr} of patterns
// we have init = 0^{mr-1}1 ... 0^{m2-1}1 0^{m1-1}1
// where mr is the size of the rth pattern
// reset = 10^{mr-1} ... 10^{m2-1} 10^{m1-1}
func PreMutliShiftAnd(sizes *[]int, init, match, d *uint) {
	for _, size := range *sizes {
		*init += 1 << *d
		*d += uint(size)
		*match += 1 << (*d - 1)
	}
}

func nextStateMultiAnd(mask *utils.Mask, state, init *uint, char byte) {
	var max uint = 1 << mask.Size()
	*state <<= 1
	if *state >= max {
		*state ^= max
	}
	*state |= *init
	*state &= mask.Get(char)
}

// MultiShiftAnd performs the ShiftAnd string matching algorithm
// for a set of patterns
func MultiShiftAnd(text string, patterns []string) [][]int {
	bigPattern := ""
	sizes := make([]int, len(patterns))
	for i, pattern := range patterns {
		bigPattern += pattern
		sizes[i] = len(pattern)
	}
	masks := PreShiftAnd(bigPattern)
	var init uint  // give where begin a pattern in the concatenation
	var match uint // tel if a pattern match
	var d uint     // shift used to construct the init and match sets
	PreMutliShiftAnd(&sizes, &init, &match, &d)
	var state uint                      // state
	occ := make([][]int, len(patterns)) // occurrences found for each pattern
	for i := 0; i < len(text); i++ {
		nextStateMultiAnd(masks[0], &state, &init, text[i])
		if state&match != 0 {
			s := 0
			for pos, size := range sizes {
				s += size
				if state&(1<<(s-1)) != 0 {
					occ[pos] = append(occ[pos], i-size+1)
				}
			}
		}
	}
	return occ
}

func extendedShiftAnd(pattern string, size int, i *int, mask *utils.Mask) {
	for pattern[*i] != ']' {
		if pattern[*i] == '-' {
			fmt.Printf("begin : %c, end : %c\n", pattern[*i-1], pattern[*i+1])
			begin := pattern[*i-1] + 1
			end := pattern[*i+1]
			for c := begin; c <= end; c++ {
				fmt.Printf("c : %c\n", c)
				var value uint = 1 << size
				value |= mask.Get(c)
				mask.Set(c, value)
			}
			*i += 2
			break
		}
		var value uint = 1 << size
		value |= mask.Get(pattern[*i])
		mask.Set(pattern[*i], value)
		*i += 1
	}
}

// func gapsShiftAnd(pattern *string, i, initGap, endGap *int, mask *utils.Mask) {

// }

func PreExtendedShiftAnd(pattern *string, initGap, endGap *int) *utils.Mask {
	mask := utils.CreateMask(false, 0)
	size := 0
	for i := 0; i < len(*pattern); i++ {
		if (*pattern)[i] == '[' {
			// class of symboles
			i += 1
			extendedShiftAnd(*pattern, size, &i, mask)
		} else if (*pattern)[i] == 'x' && (*pattern)[i+1] == '(' {
			// gaps

		} else {
			// simple symbole
			var value uint = 1 << size
			value |= mask.Get((*pattern)[i])
			mask.Set((*pattern)[i], value)
		}
		size += 1
	}
	mask.Resize(size)
	return mask
}

// func ExtendedShiftAnd(text, pattern string) {
// 	var initGap, endGap int
// 	mask := PreExtendedShiftAnd(&pattern, &initGap, &endGap)
// }

func PreShiftAndMultiMask(pattern *string) *utils.MultiMask {
	n := utils.NSubPattern(pattern)
	mask := utils.CreateMultiMask(n)
	mask.SetDefault(make([]uint, n))
	for i := 0; i < n; i++ {
		begin := i * bits.UintSize
		end := int(math.Min(float64(begin+bits.UintSize), float64(len(*pattern))))
		sub := (*pattern)[begin:end]
		for j := 0; j < len(sub); j++ {
			key := sub[j]
			index := n - i - 1
			value := mask.GetSub(key, index)
			value |= 1 << j
			mask.SetSub(value, key, index)
		}
		if i == n-1 {
			// We must save the size of the last subpattern
			mask.SetLastSize(end - begin)
		}
	}
	return mask
}

func ShiftAndMultiMask(text, pattern string) []int {
	occ := []int{}
	mask := PreShiftAndMultiMask(&pattern)
	d := make([]uint, utils.NSubPattern(&pattern))
	var match uint = 1 << (mask.Size() - 1)
	for i := 0; i < len(text); i++ {
		utils.ArrayShift(d, mask.Size())
		utils.ArrayOp(func(i int, arrays ...[]uint) uint {
			return arrays[0][i] & arrays[1][i]
		}, d, mask.Get(text[i]))
		if d[0]&match != 0 {
			occ = append(occ, i-len(pattern)+1)
		}
	}
	return occ
}
