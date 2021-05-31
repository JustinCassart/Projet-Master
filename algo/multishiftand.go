package algo

// PreMultiShiftAnd compute the init set
// and the match set such that if we have a set
// P = {P1, P2, ..., Pr} of patterns
// we have init = 0^{mr-1}1 ... 0^{m2-1}1 0^{m1-1}1
// where mr is the size of the rth pattern
// reset = 10^{mr-1} ... 10^{m2-1} 10^{m1-1}
func PreMultiShiftAnd(sizes *[]int, init, final *uint) {
	t := 0
	for _, size := range *sizes {
		*init |= 1 << t
		t += size
		*final |= 1 << (t - 1)
	}
}

func bigPattern(patterns *[]string) ([]int, string) {
	bigPattern := ""
	sizes := make([]int, len(*patterns))
	for i := len(*patterns) - 1; i >= 0; i-- {
		pattern := (*patterns)[i]
		bigPattern += pattern
		sizes[len(*patterns)-i-1] = len(pattern)
	}
	return sizes, bigPattern
}

// MultiShiftAnd performs the ShiftAnd string matching algorithm
// for a set of patterns
func MultiShiftAnd(text string, patterns []string) [][]int {
	sizes, bigPattern := bigPattern(&patterns)
	masks := PreShiftAnd(&bigPattern)
	var init uint  // give where begin a pattern in the concatenation
	var final uint // says where a pattern finish
	PreMultiShiftAnd(&sizes, &init, &final)
	var d uint
	occ := make([][]int, len(patterns)) // occurrences found for each pattern
	for i := 0; i < len(text); i++ {
		d = ((d << 1) | init) & masks[0].Get(text[i])
		if d&final != 0 {
			s := 0
			for pos, size := range sizes {
				s += size
				if d&(1<<(s-1)) != 0 {
					occ[len(occ)-pos-1] = append(occ[len(occ)-pos-1], i-size+1)
				}
			}
		}
	}
	return occ
}
