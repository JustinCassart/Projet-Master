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
	size        int             // sizes of the last substring
	// value       bool            // boolean used to create default masks
}

// methods used for mask

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

func (mask *Mask) Keys() []byte {
	keys := make([]byte, len(mask.masks))
	i := 0
	for key := range mask.masks {
		keys[i] = key
		i += 1
	}
	return keys
}

func CreateMaskBy(prepro func(pattern string) *Mask, pattern string) []*Mask {
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
		end := int(math.Min(float64(begin+bits.UintSize), float64(len(pattern))))
		mask := prepro(pattern[begin:end])
		masks[i] = mask
	}
	return masks
}

func (mask *Mask) Resize(size int) {
	mask.size = size
}

// Methods used for multimask

func CreateMultiMask(n int) *MultiMask {
	var mask MultiMask
	mask.masks = make(map[byte][]uint)
	mask.size = 0
	mask.defaultMask = make([]uint, n)
	return &mask
}

func (mask *MultiMask) Default() []uint {
	return mask.defaultMask
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

func (mask *MultiMask) Size() int {
	return mask.size
}

func (mask *MultiMask) SetDefault(value []uint) {
	mask.defaultMask = value
}

func (mask *MultiMask) SetValue(value []uint, key byte) {
	mask.masks[key] = value
}

func (mask *MultiMask) SetSub(value uint, key byte, index int) {
	if _, ok := mask.masks[key]; !ok {
		mask.masks[key] = make([]uint, len(mask.defaultMask))
		copy(mask.masks[key], mask.defaultMask)
	}
	mask.masks[key][index] = value
}

func (mask *MultiMask) SetLastSize(size int) {
	mask.size = size
}

func (multimask *MultiMask) insertMasks(masks []*Mask, alphabet *Set) {
	for i, mask := range masks {
		// for each mask we save its default value
		// into the default masks slice of the multimask
		// because each subpattern may have different size
		multimask.defaultMask[i] = mask.Default()
		// similary we save each size
		// to be used into the search part
		// multimask.sizes[i] = mask.size
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

func (mask *MultiMask) Get(key byte) []uint {
	if value, ok := mask.masks[key]; ok {
		return value
	}
	return mask.defaultMask
}

func (mask *MultiMask) GetSub(key byte, index int) uint {
	if value, ok := mask.masks[key]; ok {
		return value[index]
	}
	return mask.defaultMask[index]
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
