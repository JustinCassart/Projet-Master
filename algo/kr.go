package algo

func Hash(word string) int {
	id := 0
	for i := 0; i < len(word); i++ {
		id += int(word[i]) << (len(word) - 1 - i)
	}
	return id
}

func NextHash(text string, id, pos, patternLen int) int {
	return (id-int(text[pos-1])<<(patternLen-1))<<1 + int(text[patternLen+pos-1])

}

func checkOcc(id, patternID, pos int, occ *[]int, text, pattern string) {
	if id == patternID && Test(text, pattern, pos) {
		*occ = append(*occ, pos)
	}
}

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
