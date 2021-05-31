package tests

import (
	"stringmatching/algo"
	"stringmatching/utils"
	"testing"
)

func TestPreExtendedShiftOr(t *testing.T) {
	pattern := "a[ab]*c?[a-c]+"
	currentMaskB, currentMaskR, currentI, currentF, currentO, currentSize := algo.PreExtendedShiftOr(&pattern)
	expectedMaskB := utils.CreateMask(true, 4)
	expectedMaskB.Set('a', 4)
	expectedMaskB.Set('b', 5)
	expectedMaskB.Set('c', 3)
	expectedMaskR := utils.CreateMask(true, 4)
	expectedMaskR.Set('a', 5)
	expectedMaskR.Set('b', 5)
	expectedMaskR.Set('c', 7)
	var expectedI, expectedF, expectedO uint = 14, 11, 9
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

func TestExtendedShiftOr(t *testing.T) {
	pattern := "a[ab]*c?a+"
	text := "acaabca"
	current := algo.ExtendedShiftOr(pattern, text)
	expected := []int{2, 3, 6}
	CheckSlice(t, expected, current)
}
