package tests

import (
	"stringmatching/utils"
	"testing"
)

func TestDefaultMask1(t *testing.T) {
	mask := utils.CreateMask(true, 5)
	if mask.Default() != 31 {
		t.Errorf("expeted 31 and found %d", mask.Default())
	}
}

func TestDefaultMask0(t *testing.T) {
	mask := utils.CreateMask(false, 5)
	if mask.Default() != 0 {
		t.Errorf("expected 0 and found %d", mask.Default())
	}
}
