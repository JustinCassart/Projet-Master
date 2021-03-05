package utils

//Mask is a structure
type Mask struct {
	defaultMask int
	masks       map[int]int
}

// CreateMaks is used to create a new mask
func CreateMask(length int, value bool) Mask {
	var mask Mask
	var n int
	if value {
		n = 1
	} else {
		n = 0
	}
	for i := 1; i < length; i++ {
		n |= n << 1
	}
	mask.defaultMask = n
	return mask
}

func AddMask(mask Mask, key int) {
	mask.masks[key] = mask.defaultMask
}

func Set(mask Mask, key int, value int) {
	mask.masks[key] = value
}

func Get(mask Mask, key int) int {
	if value, ok := mask.masks[key]; ok {
		return value
	}
	return mask.defaultMask
}
