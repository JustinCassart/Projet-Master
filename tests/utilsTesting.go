package tests

import (
	"stringmatching/utils"
	"testing"
)

func CheckSlice(t *testing.T, slice1, slice2 []int) {
	if len(slice1) != len(slice2) {
		t.Errorf("Len of slice 1 = %d, len of slice 2 = %d", len(slice1), len(slice2))
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			t.Errorf("Arguments %d differ : %d != %d", i, slice1[i], slice2[i])
		}
	}
}

func CheckMask(t *testing.T, mask1, mask2 utils.Mask, keys []byte) {
	if mask1.Default() != mask2.Default() {
		t.Errorf("Expected %b as default mask but found %b", mask1.Default(), mask2.Default())
	}
	for _, key := range keys {
		v1 := mask1.Get(key)
		v2 := mask2.Get(key)
		if v1 != v2 {
			t.Errorf("Key %c : expected %b but found %b", key, v1, v2)
		}
	}
}
