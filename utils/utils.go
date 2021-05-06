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

// ArrayShift shifts the given array
// such that the
// func ArrayShift(array []uint, sizes []int) {
// 	for i := 0; i < len(array); i++ {
// 		// var max uint = uint(float(math.Max(1 << sizes[i]),
// 		var max uint = 1 << sizes[i]
// 		v := array[i]
// 		v <<= 1
// 		if i == len(array)-1 {
// 			// We are in presence with the first subpattern
// 			// So we must or the shift value with 1
// 			v |= 1
// 		}
// 		if (max == 0 && array[i] >= 1<<(bits.UintSize-1)) ||
// 			(max > 0 && v >= max) {
// 			// The size of the subpattern has the max size
// 			// So we check if the substate is greather or equal
// 			// to 1 << (w-1) where w is the size of a computer
// 			// word. In this case we must report the overflow
// 			// to the previous substate.
// 			if i != 0 {
// 				// we must report the overflow
// 				// into the previous state
// 				array[i-1] |= 1
// 			}
// 		}
// 		array[i] = v
// 	}
// }

func ArrayShift(array []uint, sizes []int) {
	for i := 0; i < len(array); i++ {
		state := array[i]
		var max uint = 1 << sizes[i]
		state <<= 1
		if sizes[i] < bits.UintSize {
			if state >= max {
				state = state ^ max
				if i != 0 {
					array[i-1] |= 1
				}
			}
		} else {
			if array[i] >= 1<<(bits.UintSize-1) {
				if i != 0 {
					array[i-1] |= 1
				}
			}
		}
		if i == len(array)-1 {
			state |= 1
		}
		array[i] = state
	}
}

func ArrayOp(operation func(i int, arrays ...[]uint) uint, arrays ...[]uint) {
	for i := 0; i < len(arrays[0]); i++ {
		arrays[0][i] = operation(i, arrays...)
	}
}

// func ArrayOp(array1, array2 []uint, operation func(array1, array2 []uint, i int) uint) {
// 	if len(array1) != len(array2) {
// 		panic("The size of the first array differ from the second one")
// 	}
// 	for i := 0; i < len(array1); i++ {
// 		array1[i] = operation(array1, array2, i)
// 	}
// }

// func ArrayOp2(array1, array2 *[]uint, operation func(array1, array2 *[]uint, i int) uint) {
// 	if len(*array1) != len(*array2) {
// 		panic("The size of the first array differ from the second one")
// 	}
// 	for i := 0; i < len(*array1); i++ {
// 		(*array1)[i] = operation(array1, array2, i)
// 	}
// }

func NSubPattern(pattern *string) int {
	return int(math.Ceil(float64(len(*pattern)) / bits.UintSize))
}
