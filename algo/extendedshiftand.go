package algo

import (
	"stringmatching/utils"
)

func preExtendedShiftAndClasses(pattern *string, i, pos int, mask *utils.Mask) int {
	if (*pattern)[i] == '[' {
		// it is a class of symboles
		if (*pattern)[i+2] == '-' {
			// it is a class like [a-z]
			first := (*pattern)[i+1]
			last := (*pattern)[i+3]
			for c := first; c <= last; c++ {
				v := mask.Get(c)
				v |= 1 << pos
				mask.Set(c, v)
			}
			return i + 4
		} else {
			// it is a class like [abc]
			i += 1 // we start with the first elment of the class
			for (*pattern)[i] != ']' {
				c := (*pattern)[i]
				v := mask.Get(c)
				v |= 1 << pos
				mask.Set(c, v)
				i += 1
			}
			return i
		}
	} else {
		// it is a single symbol
		c := (*pattern)[i]
		v := mask.Get(c)
		v |= 1 << pos
		mask.Set(c, v)
		return i
	}
}

func nextIsOptionnal(pattern *string, i int) bool {
	if (*pattern)[i+2] == '?' || (*pattern)[i+2] == '*' {
		return true
	} else if (*pattern)[i+1] == '[' {
		i += 1
		for (*pattern)[i] != ']' {
			i += 1
		}
		if (*pattern)[i+1] == '?' || (*pattern)[i+1] == '*' {
			return true
		}
	}
	return false
}

// PreExtendedShiftAnd computes the preprocessing
// for the ExtendedShiftAnd algorithm
// It returns the masks maskB, maskR, I, F and O
func PreExtendedShiftAnd(pattern *string) (*utils.Mask, *utils.Mask, uint, uint, uint, int) {
	maskB := utils.CreateMask(false, 0)
	maskR := utils.CreateMask(false, 0)
	var I, F, O uint = 0, 0, 0
	pos := 0
	lastClass := 0 // position of the last class
	for i := 0; i < len(*pattern); i++ {
		switch (*pattern)[i] {
		case '?':
			O |= 1 << (pos - 1)
			if i+2 < len(*pattern) && !nextIsOptionnal(pattern, i) {
				F |= 1 << uint(pos-1)
			}
		case '+':
			preExtendedShiftAndClasses(pattern, lastClass, pos-1, maskR)
		case '*':
			O |= 1 << (pos - 1)
			preExtendedShiftAndClasses(pattern, lastClass, pos-1, maskR)
			if i+2 < len(*pattern) && !nextIsOptionnal(pattern, i) {
				F |= 1 << uint(pos-1)
			}
		default:
			lastClass = i
			i = preExtendedShiftAndClasses(pattern, i, pos, maskB)
			if lastClass+2 < len(*pattern) && nextIsOptionnal(pattern, lastClass) {
				I |= 1 << pos
			}
			pos += 1
		}
	}
	return maskB, maskR, I, F, O, pos
}

func ExtendedShiftAnd(pattern, text string) []int {
	maskB, maskR, I, F, O, size := PreExtendedShiftAnd(&pattern)
	occ := []int{}
	var d uint = 0
	var match uint = 1 << (size - 1)
	for i := 0; i < len(text); i++ {
		d1 := ((d << 1) | 1) & maskB.Get(text[i])
		d2 := d & maskR.Get(text[i])
		d = d1 | d2
		df := d | F
		d |= O & ((^(df - I)) ^ df)
		if (d & match) != 0 {
			occ = append(occ, i)
		}
	}
	return occ
}
