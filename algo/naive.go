package algo

// Test checks if the word contained in the
// search window
func Test(text string, pattern string, pos int) bool {
	for i := 0; i < len(pattern); i++ {
		if pattern[i] != text[pos+i] {
			return false
		}
	}
	return true
}

// Naive finds all instances of a pattern
// in a text using the brute force algorithm
func Naive(text string, pattern string) []int {
	occurrences := []int{}
	for i := 0; i < len(text)-len(pattern)+1; i++ {
		if Test(text, pattern, i) {
			occurrences = append(occurrences, i)
		}
	}
	return occurrences
}
