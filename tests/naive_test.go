package tests

import (
	"stringmatching/algo"
	"testing"
)

func TestNaive(t *testing.T) {
	pattern := "aaba"
	text := "ababaaba"
	current := algo.Naive(text, pattern)
	expected := []int{4}
	CheckSlice(t, current, expected)
}
