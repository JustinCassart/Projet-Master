package utils

import (
	"fmt"
	"math"
)

// Mask is a structure countening
// a defaultmask (ex : 1111111)
// a mask
type Mask struct {
	defaultMask uint
	masks       map[byte]uint
	size        int
}

// CreateMask is used to create a new mask
func CreateMask(value bool, size int) Mask {
	var mask Mask
	mask.masks = make(map[byte]uint)
	mask.size = size
	var n uint
	if value {
		n = uint(math.Pow(2, float64(size))) - 1
	}
	mask.defaultMask = n
	return mask
}

// Size returns the number of bits
// used for each entry
func (mask Mask) Size() int {
	return mask.size
}

// SetNew adds a new entry for the given key
// with the default mask as value
func (mask Mask) SetNew(key byte) {
	mask.masks[key] = mask.defaultMask
}

// Set replaces the previous value
// for the entry key
// with the new value
func (mask Mask) Set(key byte, value uint) {
	mask.masks[key] = value
}

// Get returns the value of the entry
// or the default mask if the entry doesn't exist
func (mask Mask) Get(key byte) uint {
	if value, ok := mask.masks[key]; ok {
		return value
	}
	return mask.defaultMask
}

// Default returns the default mask
func (mask Mask) Default() uint {
	return mask.defaultMask
}

// Display is TODO
func (mask Mask) Display() {
	for key, value := range mask.masks {
		fmt.Println(string(key), " -> ", getSTR(value, mask.size))
	}
}

func getSTR(n uint, size int) string {
	bin := fmt.Sprintf("%b", n)
	for len(bin) != size {
		bin = "0" + bin
	}
	return bin
}
