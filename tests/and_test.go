package tests

import (
	"math/bits"
	"stringmatching/algo"
	"stringmatching/utils"
	"testing"
)

func TestPreAnd(t *testing.T) {
	pattern := "aaba"
	currentMask := algo.PreShiftAnd(pattern)
	expectedMask := make([]*utils.Mask, 1)
	expectedMask[0] = utils.CreateMask(false, 4)
	expectedMask[0].Set('a', 11)
	expectedMask[0].Set('b', 4)
	CheckMasks(t, expectedMask, currentMask)
}

func TestPreAnd64(t *testing.T) {
	pattern := "aaaaaaaa"
	for len(pattern) < bits.UintSize {
		pattern += "aaaaaaaa"
	}
	pattern += "aaba"
	currentMasks := algo.PreShiftAnd(pattern)
	expectedMasks := make([]*utils.Mask, 2)
	expectedMasks[0] = utils.CreateMask(false, bits.UintSize)
	expectedMasks[0].Set('a', 1<<bits.UintSize-1)
	expectedMasks[1] = utils.CreateMask(false, 4)
	expectedMasks[1].Set('a', 11)
	expectedMasks[1].Set('b', 4)
	CheckMasks(t, expectedMasks, currentMasks)
}

func TestShiftAnd(t *testing.T) {
	pattern := "aaba"
	text := "ababaaba"
	current := algo.ShiftAnd(text, pattern)
	expected := []int{4}
	CheckSlice(t, expected, current)
}

func TestBigShiftAnd(t *testing.T) {
	pattern := "abbaabba"
	for len(pattern) < 128 {
		pattern += "abbaabba"
	}
	text := pattern
	current := algo.ShiftAnd(text, pattern)
	expected := []int{0}
	CheckSlice(t, expected, current)
}

func TestPreMultiShiftAnd(t *testing.T) {
	sizes := []int{4, 3, 2, 5}
	var init, match, d uint
	algo.PreMutliShiftAnd(&sizes, &init, &match, &d)
	if init != 657 {
		t.Errorf("\nExpected  %b as init set\nbut found %b\n", 1169, init)
	}
	if match != 8520 {
		t.Errorf("\nExpected  %b as match set\nbut found %b\n", 8362, match)
	}
}

func TestMultiShiftAnd(t *testing.T) {
	patterns := []string{"ab", "ba", "abba"}
	text := "abcbacabba"
	current := algo.MultiShiftAnd(text, patterns)
	expected := [][]int{{0, 6}, {3, 8}, {6}}
	CheckTable(t, expected, current)
}

func TestPreMultiMask(t *testing.T) {
	pattern := "abba"
	current := algo.PreShiftAndMultiMask(pattern)
	expectedSize := []int{4}
	CheckSlice(t, expectedSize, *current.Size())
	expectedKeys := []byte{'a', 'b'}
	currentKeys := current.Keys()
	utils.SortSlice(currentKeys)
	CheckSlice(t, expectedKeys, currentKeys)
	if len(currentKeys) != len(expectedKeys) {
		t.Errorf("Len error : expected %d but found %d", len(expectedKeys), len(currentKeys))
	}
	expectedValues := make(map[byte][]uint)
	expectedValues['a'] = []uint{9}
	expectedValues['b'] = []uint{6}
	for _, key := range expectedKeys {
		CheckSlice(t, expectedValues[key], current.Get(key))
	}
}

func TestPreMultiMaskShiftAnd64(t *testing.T) {
	pattern := "aaba"
	for len(pattern) < 68 {
		pattern += "aaba"
	}
	current := algo.PreShiftAndMultiMask(pattern)
	expectedSizes := []int{4, 64}
	CheckSlice(t, expectedSizes, *current.Size())
	expectedKeys := []byte{'a', 'b'}
	currentKeys := current.Keys()
	utils.SortSlice(currentKeys)
	if len(expectedKeys) != len(currentKeys) {
		t.Errorf("Size error : Expected %d keys but found %d", len(expectedKeys), len(currentKeys))
	}
	expectedValues := make(map[byte][]uint)
	expectedValues['a'] = []uint{13527612320720337851, 11}
	expectedValues['b'] = []uint{4919131752989213764, 4}
	for key, value := range expectedValues {
		CheckSlice(t, value, current.Get(key))
	}
}

func TestMultiMaskShiftAnd(t *testing.T) {
	pattern := "aaba"
	text := "ababaaba"
	current := algo.ShiftAndMultiMask(text, pattern)
	expected := []int{4}
	CheckSlice(t, expected, current)
}

func TestBigShiftAndMultiMask(t *testing.T) {
	pattern := "abbaabba"
	for len(pattern) < 128 {
		pattern += "abbaabba"
	}
	text := pattern
	current := algo.ShiftAndMultiMask(text, pattern)
	expected := []int{0}
	CheckSlice(t, expected, current)
}

// func TestPreClassesShiftAnd(t *testing.T) {
// 	pattern := "[123]/[1-3]/199[89]"
// 	current := algo.PreClassesShiftAnd(pattern)
// 	expected := utils.CreateMask(false, 8)
// 	expected.Set('1', 21)
// 	expected.Set('2', 5)
// 	expected.Set('3', 5)
// 	expected.Set('8', 128)
// 	expected.Set('9', 224)
// 	expected.Set('/', 10)
// 	CheckMask(t, expected, current)
// }
