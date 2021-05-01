package tests

import (
	"stringmatching/algo"
	"testing"
)

func TestHash(t *testing.T) {
	word := "abaa"
	currentValue := algo.Hash(word)
	expectedValue := 1459
	if currentValue != expectedValue {
		t.Errorf("Expected id %d but found %d\n", expectedValue, currentValue)
	}
}

func TestNextHash(t *testing.T) {
	text := "abaab"
	patternLen := 4
	id := algo.Hash(text[:4])
	currentNextID := algo.NextHash(text, id, 1, patternLen)
	expectedNextID := algo.Hash(text[1:])
	if currentNextID != expectedNextID {
		t.Errorf("Expected next id %d but found %d\n", expectedNextID, currentNextID)
	}
}

func TestKR(t *testing.T) {
	pattern := "aaba"
	text := "ababaaba"
	current := algo.KR(text, pattern)
	expected := []int{4}
	CheckSlice(t, current, expected)
}
