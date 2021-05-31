package algo

import (
	"math"
	"math/bits"
	"stringmatching/utils"
)

// PreShiftOr creates the mask B for a pattern
// such that its length is less or equal to 64
// For example the mask for aaba is
// B[a] = 1011
// B[c] = 0010
func preShiftAnd(pattern *string) *utils.Mask {
	mask := utils.CreateMask(false, len(*pattern))
	for i := 0; i < len(*pattern); i++ {
		var value uint = 1 << i
		value |= mask.Get((*pattern)[i])
		mask.Set((*pattern)[i], value)
	}
	return mask
}

// PreShiftAnd computes the masks for the pattern
// It returns a slice of mask
// If the length of the pattern is less or equal than the size of a word
// then the slice contains the single mask representing the pattern
// otherwise, the slice contains the mask for each subpattern s
// such tath the length of s is at most the size of a word
func PreShiftAnd(pattern *string) []*utils.Mask {
	return utils.CreateMaskBy(preShiftAnd, pattern)
}

func ShiftAnd(text, pattern string, masks []*utils.Mask) []int {
	occ := []int{}
	var match uint = 1 << (masks[0].Size() - 1)
	var d uint = 0
	for i := 0; i < len(text); i++ {
		nextStateAnd(masks[0], text[i], &d)
		if d&match != 0 {
			if len(masks) == 1 {
				occ = append(occ, i-masks[0].Size()+1)
			} else {
				if checkocc(masks, &text, 1, i+1, func(match, d uint) bool { return d&match != 0 }) {
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

func checkocc(masks []*utils.Mask, text *string, id, begin int, check func(match, d uint) bool) bool {
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
		if check(match, d) {
			if id != len(masks)-1 {
				return checkocc(masks, text, id+1, begin+i+1, check)
			}
			return true
		}
	}
	return false
}

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

func ShiftAndMultiMask(text, pattern *string, mask *utils.MultiMask) []int {
	occ := []int{}
	// mask := PreShiftAndMultiMask(pattern)
	d := make([]uint, utils.NSubPattern(pattern))
	var match uint = 1 << (mask.Size() - 1)
	for i := 0; i < len(*text); i++ {
		utils.ArrayShift(&d, mask.Size())
		value := mask.Get((*text)[i])
		utils.ArrayOp(func(i int, arrays ...*[]uint) uint {
			return (*arrays[0])[i] & (*arrays[1])[i]
		}, &d, &value)
		if d[0]&match != 0 {
			occ = append(occ, i-len(*pattern)+1)
		}
	}
	return occ
}
