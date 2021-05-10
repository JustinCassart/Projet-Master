package main

import (
	"fmt"
	"reflect"
)

// func compare(text string, pattern string, size int) {
// 	fmt.Printf("Recherche de %s dans %s (taille max = %d)\n", pattern, text, size)
// 	fmt.Printf("Occurrence trouvée par shiftand : %v\n", algo.ShiftAnd(text, pattern, size))
// 	fmt.Printf("Occurrence trouvée par shiftand : %v\n\n", algo.ShiftOr(text, pattern))
// }

type super struct {
	elements map[int]int
}

type imp struct {
	super
}

type imp2 struct {
	elements map[int][]int
	super
}

func (s super) keys() []int {
	keys := []int{}
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

func set(array []int, pos int, value int) {
	array[pos] = value
}

func app(array []int, value int) {
	array = append(array, value)
}

func showType(inter interface{}) {
	data := reflect.ValueOf(inter)
	fmt.Println(data.Kind())
}

func check(inter1, inter2 interface{}) bool {
	array1 := reflect.ValueOf(inter1)
	array2 := reflect.ValueOf(inter2)
	if array1.Kind() != reflect.Slice {
		return false
	}
	if array1.Kind() != array2.Kind() {
		return false
	}
	if array1.Len() != array2.Len() {
		return false
	}
	for i := 0; i < array1.Len(); i++ {
		fmt.Println("index : ", array1.Index(i).Interface(), " ", array2.Index(i).Interface())
		if array1.Index(i).Interface() != array2.Index(i).Interface() {
			return false
		}
	}
	return true
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {

	// text := "ababaa"
	// pattern := "abaa"
	// start := time.Now()
	// algo.KMP(text, pattern)
	// fmt.Printf("Execution time for KMP : %d\n", time.Since(start))
	// start = time.Now()
	// algo.KR(text, pattern)
	// fmt.Printf("Execution time for KR : %d\n", time.Since(start))
	// start = time.Now()
	// algo.ShiftOr(text, pattern)
	// fmt.Printf("Execution time for ShiftOR : %d\n", time.Since(start))
	// start = time.Now()
	// algo.Naive(text, pattern)
	// fmt.Printf("Execution time for Naive : %d\n", time.Since(start))
	// for c := '1'; c <= '9'; c++ {
	// 	fmt.Printf("%c\n", c)
	// }
	// s := super{elements: make(map[int]int)}
	// s.elements[1] = 2
	// s.elements[2] = 3
	// i1 := imp{super: super{elements: make(map[int]int)}}
	// i1.super.elements[3] = 4
	// i2 := imp2{elements: make(map[int][]int)}
	// i2.elements[1] = []int{1, 2, 3}
	// i2.elements[2] = []int{2, 3, 4}
	// fmt.Println(s.keys())
	// fmt.Println(i1.keys())
	// fmt.Println(i2.keys())
	// s1 := []int{1, 2}
	// s2 := []int{1, 2}
	// s3 := []int{2, 1}
	// fmt.Println(check(s1, s2))
	// fmt.Println(check(s1, s3))
	arr := make([]int, 5)
	set(arr, 3, 25)
	fmt.Println(arr)
	app(arr, 45)
	fmt.Println(arr)
	nums := []int{1, 2, 3}
	fmt.Println(sum(nums...))
	fmt.Println(sum(1, 2, 3))
	var x uint = 1<<64 - 1
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", (x<<1)|1)
	var y uint = 1 << 63
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", y<<1)

}
