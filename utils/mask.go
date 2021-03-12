package utils

import (
	"fmt"
)

// Mask is a structure countening
// a defaultmask (ex : 1111111)
// a mask
type Mask struct {
	defaultMask uint
	masks       map[byte]uint
}

// CreateMask is used to create a new mask
func CreateMask(length int, value bool) Mask {
	var mask Mask
	mask.masks = make(map[byte]uint)
	var n uint
	if value {
		n = ^n
	}
	n ^= n << uint(length)
	mask.defaultMask = n
	return mask
}

// AddMask adds an entry
func AddMask(mask Mask, key byte) {
	mask.masks[key] = mask.defaultMask
}

// Set changes value of an entry
func Set(mask Mask, key byte, value uint) {
	mask.masks[key] = value
}

// Get returns the value of the entry
// or the default mask if the entry doesn't exist
func Get(mask Mask, key byte) uint {
	if value, ok := mask.masks[key]; ok {
		return value
	}
	return mask.defaultMask
}

// Default returns the default mask
func Default(mask Mask) uint {
	return mask.defaultMask
}

// Display is TODO
func Display(mask Mask) {
	for key, value := range mask.masks {
		fmt.Print(string(key), " ")
		fmt.Printf("%05b\n", value)
	}
}
