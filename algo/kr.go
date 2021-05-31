package algo

// Hash computes the id of a word
func Hash(word string) int {
	id := 0
	for i := 0; i < len(word); i++ {
		id += int(word[i]) << (len(word) - 1 - i)
	}
	return id
}

// NextHash computes the id of window from the prevuous one
func NextHash(text string, id, pos, patternLen int) int {
	return (id-int(text[pos-1])<<(patternLen-1))<<1 + int(text[patternLen+pos-1])

}

// checkOcc first checks if the id of the window is identical
// to that of the pattern. If it is, it checks if the word
// contained in the window is the pattern.
func checkOcc(id, patternID, pos int, occ *[]int, text, pattern string) {
	if id == patternID && Test(text, pattern, pos) {
		*occ = append(*occ, pos)
	}
}

// KR finds all instances of a simple pattern
// using the Karp, Rabin's algorithm
func KR(text, pattern string) []int {
	occ := []int{}
	patternID := Hash(pattern)
	id := Hash(text[:len(pattern)])
	checkOcc(id, patternID, 0, &occ, text, pattern)
	for i := 1; i < len(text)-len(pattern)+1; i++ {
		id = NextHash(text, id, i, len(pattern))
		checkOcc(id, patternID, i, &occ, text, pattern)
	}
	return occ
}
