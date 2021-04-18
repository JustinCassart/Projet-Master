package tests

import (
	"stringmatching/algo"
	"stringmatching/utils"
	"testing"
)

func TestPreOr(t *testing.T) {
	pattern := "aaba"
	currentMask := algo.PreShiftOr(pattern)
	expectedMask := utils.CreateMask(true, 4)
	expectedMask.Set('a', 4)
	expectedMask.Set('b', 11)
	for _, v := range []byte{'a', 'b', 'c'} {
		currentV := currentMask.Get(v)
		expectedV := expectedMask.Get(v)
		if currentV != expectedV {
			t.Errorf("Expected %b but found %b for key %c", expectedV, currentV, v)
		}
	}
}

func TestShiftOr(t *testing.T) {
	pattern := "aaba"
	text := "ababaaba"
	current := algo.ShiftOr(text, pattern, 5)
	expected := []int{4}
	CheckSlice(t, current, expected)
}

func TestMultiShiftOr(t *testing.T) {
	pattern := "aaba"
	text := "ababaaba"
	current := algo.ShiftOr(text, pattern, 2)
	expected := []int{4}
	CheckSlice(t, current, expected)
}
