package algo

func PreMultiShiftOr(sizes *[]int, init, final *uint) {
	t := 0
	for _, size := range *sizes {
		*init ^= 1 << t
		t += size
		*final ^= 1 << (t - 1)
	}
}

func MultiShiftOr(text string, patterns []string) [][]int {
	occ := make([][]int, len(patterns))
	sizes, bigPattern := bigPattern(&patterns)
	masks := PreShiftOr(&bigPattern)
	var init uint = (1 << len(bigPattern)) - 1
	var final uint = (1 << len(bigPattern)) - 1
	var match uint = (1 << len(bigPattern)) - 1
	PreMultiShiftOr(&sizes, &init, &final)
	var d uint = (1 << len(bigPattern)) - 1
	for i := 0; i < len(text); i++ {
		d = ((d << 1) & init) | masks[0].Get(text[i])
		if (d | final) != match {
			s := 0
			for pos, size := range sizes {
				s += size
				if d|(1<<(s-1)) > d {
					occ[len(occ)-pos-1] = append(occ[len(occ)-pos-1], i-size+1)
				}
			}
		}
	}
	return occ
}
