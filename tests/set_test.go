package tests

import (
	"stringmatching/utils"
	"testing"
)

func TestValues(t *testing.T) {
	set := utils.CreateSet()
	set.Add('b')
	set.Add('c')
	current := set.Values()
	expected := []byte{'b', 'c'}
	CheckSlice(t, expected, current)
}
