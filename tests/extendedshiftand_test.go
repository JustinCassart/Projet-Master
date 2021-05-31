package tests

import (
	"stringmatching/algo"
	"stringmatching/utils"
	"testing"
)

func TestPreExtendedShiftAnd(t *testing.T) {
	pattern := "a[ab]*c?[a-c]+"
	currentMaskB, currentMaskR, currentI, currentF, currentO, currentSize := algo.PreExtendedShiftAnd(&pattern)
	expectedMaskB := utils.CreateMask(false, 0)
	expectedMaskB.Set('a', 11)
	expectedMaskB.Set('b', 10)
	expectedMaskB.Set('c', 12)
	expectedMaskR := utils.CreateMask(false, 0)
	expectedMaskR.Set('a', 10)
	expectedMaskR.Set('b', 10)
	expectedMaskR.Set('c', 8)
	var expectedI, expectedF, expectedO uint = 1, 4, 6
	var expectedSize = 4
	if currentI != expectedI {
		t.Errorf("Mask I error : expected %b but found %b", expectedI, currentI)
	}
	if currentF != expectedF {
		t.Errorf("Mask F error : expected %b but found %b", expectedF, currentF)
	}
	if currentO != expectedO {
		t.Errorf("Mask O error : expected %b but found %b", currentO, expectedO)
	}
	if currentSize != expectedSize {
		t.Errorf("Size error : expected %d but found %d", expectedSize, currentSize)
	}
	CheckMask(t, expectedMaskB, currentMaskB)
	CheckMask(t, expectedMaskR, currentMaskR)
}

func TestExtendedShiftAnd(t *testing.T) {
	pattern := "a[ab]*c?a+"
	text := "acaabca"
	current := algo.ExtendedShiftAnd(pattern, text)
	expected := []int{2, 3, 6}
	CheckSlice(t, expected, current)
}
