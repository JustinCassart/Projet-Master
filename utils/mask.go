package utils

import (
	"fmt"
	"math"
	"math/bits"
)

// Mask is a structure countening
// a defaultmask (ex : 1111111)
// a mask
type Mask struct {
	defaultMask uint          // default mask for non given symboles
	masks       map[byte]uint // mask for given symboles
	size        int           // size of the string
}

type MultiMask struct {
	defaultMask []uint          // default mask for non given symboles
	masks       map[byte][]uint // mask for given symboles for each pattern
	sizes       []int           // sizes of the differents masks
	// value       bool            // boolean used to create default masks
}

func CreateMultiMask(n int) *MultiMask {
	var mask MultiMask
	mask.masks = make(map[byte][]uint)
	mask.sizes = make([]int, n)
	mask.defaultMask = make([]uint, n)
	return &mask
}

func (mask *MultiMask) Default() *[]uint {
	return &mask.defaultMask
}

func CreateMultiMaskBy(prepro func(pattern string) *Mask, pattern string) *MultiMask {
	if len(pattern) > 64 {
		// computes the number of sub-patterns
		n := int(math.Ceil(float64(len(pattern)) / bits.UintSize))
		multimask := CreateMultiMask(n)
		alphabet := CreateSet()
		masks := make([]*Mask, n)
		for i := 0; i < n; i++ {
			// begin corresponds to the first index
			// of the factor in the pattern
			begin := i * bits.UintSize
			// end corresponds to the last index + 1
			// of the factor in the pattern
			// it is at most equal to the lengteh of the pattern
			end := int(math.Min(float64((i+1)*bits.UintSize),
				float64(len(pattern))))
			mask := prepro(pattern[begin:end])
			masks[n-i-1] = mask
			for _, key := range mask.Keys() {
				alphabet.Add(key)
			}
		}
		multimask.insertMasks(masks, alphabet)
		return multimask
	}
	multimask := CreateMultiMask(1)
	masks := []*Mask{prepro(pattern)}
	fmt.Printf("masks : %v\n", masks[0])
	fmt.Printf("keys : %v\n", masks[0].Keys())
	alphabet := CreateSet()
	for _, key := range masks[0].Keys() {
		fmt.Printf("OK : %c\n", key)
		alphabet.Add(key)
	}
	multimask.insertMasks(masks, alphabet)
	return multimask
}

func (mask *MultiMask) Size() []int {
	return mask.sizes
}

func (multimask *MultiMask) insertMasks(masks []*Mask, alphabet *Set) {
	for i, mask := range masks {
		// for each mask we save its default value
		// into the default masks slice of the multimask
		// because each subpattern may have different size
		multimask.defaultMask[i] = mask.Default()
		// similary we save each size
		// to be used into the search part
		multimask.sizes[i] = mask.size
		// we must initialise the masks
		// for each symbol of the alphabet
		// used for the whole pattern
		// with the defaults values
		for _, alpha := range alphabet.Values() {
			if i == 0 {
				// the map is not yet initialised
				multimask.masks[alpha] = make([]uint, len(masks))
			}
			multimask.masks[alpha][i] = mask.Default()
		}
	}
	// We can know set the real values
	// for each masks
	for i, mask := range masks {
		for key, value := range mask.masks {
			multimask.masks[key][i] = value
		}
	}
}

func (mask *MultiMask) Get(key byte) *[]uint {
	arr := mask.masks[key]
	return &arr
}

// CreateMask is used to create a new mask
func CreateMask(value bool, size int) *Mask {
	var mask Mask
	mask.masks = make(map[byte]uint)
	mask.size = size
	mask.defaultMask = getDefault(value, size)
	return &mask
}

func getDefault(value bool, size int) uint {
	var n uint
	if value {
		n = uint(1<<size - 1)
	}
	return n
}

// Size returns the number of bits
// used for each entry
func (mask *Mask) Size() int {
	return mask.size
}

// Set replaces the previous value
// for the entry key
// with the new value
func (mask *Mask) Set(key byte, value uint) {
	mask.masks[key] = value
}

// Get returns the value of the entry
// or the default mask if the entry doesn't exist
func (mask *Mask) Get(key byte) uint {
	if value, ok := mask.masks[key]; ok {
		return value
	}
	return mask.defaultMask
}

// Default returns the default mask
func (mask *Mask) Default() uint {
	return mask.defaultMask
}

// Display is TODO
func (mask *Mask) Display() {
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

func (mask *Mask) Keys() []byte {
	keys := make([]byte, len(mask.masks))
	i := 0
	for key := range mask.masks {
		keys[i] = key
		i += 1
	}
	return keys
}

func (mask *MultiMask) Keys() []byte {
	keys := make([]byte, len(mask.masks))
	i := 0
	for key := range mask.masks {
		keys[i] = key
		i += 1
	}
	return keys
}

func CreateMaskBy(prepro func(pattern string) *Mask, pattern string) []*Mask {
	if len(pattern) > 64 {
		// computes the number of sub-patterns
		n := int(math.Ceil(float64(len(pattern)) / bits.UintSize))
		// creates a slice with n masks
		masks := make([]*Mask, n)
		for i := 0; i < n; i++ {
			// begin corresponds to the first index
			// of the factor in the pattern
			begin := i * bits.UintSize
			// end corresponds to the last index + 1
			// of the factor in the pattern
			// it is at most equal to the lengteh of the pattern
			end := int(math.Min(float64((i+1)*bits.UintSize),
				float64(len(pattern))))
			mask := prepro(pattern[begin:end])
			masks[i] = mask
		}
		return masks
	}
	return []*Mask{prepro(pattern)}
}

func (mask *Mask) Resize(size int) {
	mask.size = size
}
