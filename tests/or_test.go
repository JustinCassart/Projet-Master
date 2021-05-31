package tests

import (
	"stringmatching/algo"
	"stringmatching/utils"
	"testing"
)

func TestPreOr(t *testing.T) {
	pattern := "aaba"
	currentMask := algo.PreShiftOr(&pattern)
	expectedMask := utils.CreateMask(true, 4)
	expectedMask.Set('a', 4)
	expectedMask.Set('b', 11)
	if len(currentMask) != 1 {
		t.Errorf("Size of masks error : expected 1 but found %d", len(currentMask))
	}
	CheckMask(t, expectedMask, currentMask[0])
}

func TestPreOr64(t *testing.T) {
	pattern := "aaaaaaaa"
	for len(pattern) < 64 {
		pattern += "aaaaaaaa"
	}
	pattern += "bbb"
	currentMasks := algo.PreShiftOr(&pattern)
	expectedMasks := make([]*utils.Mask, 2)
	expectedMasks[0] = utils.CreateMask(true, 64)
	expectedMasks[0].Set('a', 0)
	expectedMasks[1] = utils.CreateMask(true, 3)
	expectedMasks[1].Set('b', 0)
	CheckMasks(t, expectedMasks, currentMasks)
}

func TestShiftOr(t *testing.T) {
	pattern := "aaba"
	texts := "ababaaba"
	current := algo.ShiftOr(texts, pattern, algo.PreShiftOr(&pattern))
	expected := []int{4}
	CheckSlice(t, current, expected)
}
