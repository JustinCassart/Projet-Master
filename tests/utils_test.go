package tests

import (
	"stringmatching/utils"
	"testing"
)

func TestSortInt(t *testing.T) {
	arrInt := []int{5, 2, 3}
	utils.SortSlice(arrInt)
	for i, v := range []int{2, 3, 5} {
		if arrInt[i] != v {
			t.Errorf("Value error in index %d : expected %v but found %v", i, v, arrInt[i])
		}
	}
}

func TestArrayShift(t *testing.T) {
	states := []uint{3, 2, 3}
	expected := []uint{3, 1, 3}
	utils.ArrayShift(&states, &[]int{2, 2, 2})
	CheckSlice(t, expected, states)
}

func TestArrayOp(t *testing.T) {
	array1 := []uint{1, 2, 3}
	array2 := []uint{3, 2, 1}
	utils.ArrayOp(&array1, &array2, func(array1, array2 *[]uint, i int) uint {
		return (*array1)[i] & (*array2)[i]
	})
	expected := []uint{1, 2, 1}
	CheckSlice(t, expected, array1)
}
