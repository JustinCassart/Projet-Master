package utils

type Set struct {
	values map[byte]struct{}
}

func CreateSet() *Set {
	set := Set{}
	set.values = make(map[byte]struct{})
	return &set
}

func (set *Set) Add(key byte) {
	set.values[key] = struct{}{}
}

func (set *Set) Values() []byte {
	values := make([]byte, len(set.values))
	i := 0
	for key := range set.values {
		values[i] = key
		i += 1
	}
	return values
}
