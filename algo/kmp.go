package algo

// PreKMP computes the next table
// for the pattern
func PreKMP(pattern string) []int {
	m := len(pattern)
	next := make([]int, m+1)
	d := -1
	j := 0
	next[0] = -1
	for j < m {
		for d > -1 && pattern[j] != pattern[d] {
			d = next[d]
		}
		d += 1
		j += 1
		if j >= m || pattern[j] != pattern[d] {
			next[j] = d
		} else {
			next[j] = next[d]
		}
	}
	return next
}

// KMP finds all occurrence
// of a pattern in a text
// using the knuth, Morris and Pratt's algorithm
func KMP(text, pattern string) []int {
	occ := []int{}
	next := PreKMP(pattern)
	n := len(text)
	m := len(pattern)
	i := 0
	j := 0
	for i+m-j <= n {
		for j != -1 && pattern[j] != text[i] {
			j = next[j]
		}
		i += 1
		j += 1
		if j == m {
			occ = append(occ, i-j)
			j = next[m]
		}
	}
	return occ
}
