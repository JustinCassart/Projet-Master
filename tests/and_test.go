package tests

import (
	"stringmatching/algo"
	"stringmatching/utils"
	"testing"
)

func TestPreAnd(t *testing.T) {
	pattern := "aaba"
	currentMask := algo.PreShiftOr(pattern)
	expectedMask := utils.CreateMask(false, 4)
	CheckMask(t, expectedMask, currentMask, []byte{'a', 'b', 'c'})
}
