package utils

import (
	"fmt"
	"math"
	"math/bits"
	"reflect"
	"sort"
)

func SortSlice(array interface{}) {
	arr := reflect.ValueOf(array)
	if arr.Kind() != reflect.Slice {
		panic("array is not a slice")
	}
	sort.SliceStable(array, func(i, j int) bool {
		return fmt.Sprintf("%v", arr.Index(i).Interface()) < fmt.Sprintf("%v", arr.Index(j).Interface())
	})
}

// ArrayShift performs the shift between the elements
// of a slice
func ArrayShift(array *[]uint, size int) {
	for i := 0; i < len(*array); i++ {
		var max uint = 1 << size
		state := (*array)[i]
		state <<= 1
		if i == 0 {
			// We must check if the size of the state
			// overflows the number of bits (size)
			// No report is needed
			if state >= max {
				// It's work when the size is equal to
				// the one of a computer word too.
				// Because in this cas max is 0.
				// So we xor with 0
				state ^= max
			}
		} else {
			// By definition the size of internal word are
			// equal to the one of a computer word
			if (*array)[i] >= 1<<(bits.UintSize-1) {
				(*array)[i-1] |= 1
			}
		}
		if i == len(*array)-1 {
			// We add one in the last position
			// to do the suppposition of a new occurrence.
			state |= 1
		}
		(*array)[i] = state
	}
}

// ArrayOp performs a given operation over each element of a slice
func ArrayOp(operation func(i int, arrays ...*[]uint) uint, arrays ...*[]uint) {
	for i := 0; i < len(*arrays[0]); i++ {
		(*arrays[0])[i] = operation(i, arrays...)
	}
}

// NSubPattern computes the number of computer words needed
// to represent the pattern
func NSubPattern(pattern *string) int {
	return int(math.Ceil(float64(len(*pattern)) / bits.UintSize))
}
