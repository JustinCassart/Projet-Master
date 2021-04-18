package tests

import (
	"stringmatching/algo"
	"testing"
)

func TestPreprocessingKMP(t *testing.T) {
	pattern := "aaba"
	current := algo.PreKMP(pattern)
	expected := []int{-1, -1, 1, -1, 1}
	CheckSlice(t, current, expected)
}

func TestKMP(t *testing.T) {
	pattern := "aaba"
	text := "ababaaba"
	current := algo.KMP(text, pattern)
	expected := []int{4}
	CheckSlice(t, current, expected)
}
