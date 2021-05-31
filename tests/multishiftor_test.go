package tests

import (
	"stringmatching/algo"
	"testing"
)

func TestPreProMultiShiftOr(t *testing.T) {
	var currentInit, currentFinal uint = (1 << 13) - 1, (1 << 13) - 1
	algo.PreMultiShiftOr(&[]int{5, 4, 4}, &currentInit, &currentFinal)
	var expectedInit, expectedFinal uint = 7646, 3823
	if currentFinal != expectedFinal {
		t.Errorf("Final mask error : \nexpected  %b\nbut found %b", expectedFinal, currentFinal)
	}
	if currentInit != expectedInit {
		t.Errorf("Init mask error : \nexpected  %b\nbut found %b", expectedInit, currentInit)
	}
}

func TestMultiShiftOr(t *testing.T) {
	patterns := []string{"abaa", "aaba", "acbba"}
	text := "abaabacbba"
	current := algo.MultiShiftOr(text, patterns)
	expected := [][]int{{0}, {2}, {5}}
	CheckTable(t, expected, current)
}
