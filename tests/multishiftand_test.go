package tests

import (
	"stringmatching/algo"
	"testing"
)

func TestPreProMultiShiftAnd(t *testing.T) {
	var currentInit, currentFinal uint
	algo.PreMultiShiftAnd(&[]int{5, 4, 4}, &currentInit, &currentFinal)
	var expectedInit, expectedFinal uint = 545, 4368
	if currentFinal != expectedFinal {
		t.Errorf("Final mask error : \nexpected  %b\nbut found %b", expectedFinal, currentFinal)
	}
	if currentInit != expectedInit {
		t.Errorf("Init mask error : expected %b but found %b", expectedInit, currentInit)
	}
}

func TestMultiShiftAnd(t *testing.T) {
	patterns := []string{"abaa", "aaba", "acbba"}
	text := "abaabacbba"
	current := algo.MultiShiftAnd(text, patterns)
	expected := [][]int{{0}, {2}, {5}}
	CheckTable(t, expected, current)
}
