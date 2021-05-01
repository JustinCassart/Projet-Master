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

func TestMaxSize1(t *testing.T) {
	mask := utils.CreateMask(true, 64)
	const expect uint = 1<<64 - 1
	if mask.Default() != expect {
		t.Errorf("\nExpected  %b as default mask\nbut found %b", expect, mask.Default())
	}
}

func TestMaxSize0(t *testing.T) {
	mask := utils.CreateMask(false, 64)
	const expect uint = 0
	if mask.Default() != expect {
		t.Errorf("\nExpected  %b as default mask\nbut found %b", expect, mask.Default())
	}
}
