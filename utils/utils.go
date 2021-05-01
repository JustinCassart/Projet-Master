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
func ArrayShift(array *[]uint, sizes *[]int) {
	for i := 0; i < len(*array); i++ {
		var max uint = 1 << (*sizes)[i]
		v := (*array)[i]
		v <<= 1
		if i == len(*array)-1 {
			// We are in presence with the first subpattern
			// So we must or the shift value with 1
			v |= 1
		}
		if v >= max {
			if i != 0 {
				// we must report the overlap 1
				// into the previous state
				(*array)[i-1] |= 1
			}
			v ^= max
		}
		(*array)[i] = v
	}
}

func ArrayOp(array1, array2 *[]uint, operation func(array1, array2 *[]uint, i int) uint) {
	if len(*array1) != len(*array2) {
		panic("The size of the first array differ from the second one")
	}
	for i := 0; i < len(*array1); i++ {
		(*array1)[i] = operation(array1, array2, i)
	}
}

func NSubPattern(pattern *string) int {
	return int(math.Ceil(float64(len(*pattern)) / bits.UintSize))
}
