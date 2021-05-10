package tests

import (
	"reflect"
	"stringmatching/utils"
	"testing"
)

func CheckSlice(t *testing.T, expected, current interface{}) {
	expectedSlice := reflect.ValueOf(expected)
	currentSlice := reflect.ValueOf(current)
	if expectedSlice.Kind() != reflect.Slice {
		t.Error("The expected value is not a slice")
	}
	if currentSlice.Kind() != reflect.Slice {
		t.Error("The current value is not a slice")
	}
	if expectedSlice.Len() != currentSlice.Len() {
		t.Errorf("Length of slices differ : expected %d but found %d", expectedSlice.Len(), currentSlice.Len())
	}
	for i := 0; i < expectedSlice.Len(); i++ {
		expectedValue := expectedSlice.Index(i).Interface()
		currentValue := currentSlice.Index(i).Interface()
		if currentValue != expectedValue {
			t.Errorf("Arguments %d differ : expected %b but found %b", i, expectedValue, currentValue)
		}
	}
}

func CheckTable(t *testing.T, expectedTable, currentTable [][]int) {
	if len(expectedTable) != len(currentTable) {
		t.Errorf("Table size error : expected %d but found %d", len(expectedTable), len(currentTable))
	}
	for i := 0; i < len(expectedTable); i++ {
		t.Logf("expectedSlice %v currentSlice %v", expectedTable[i], currentTable[i])
		CheckSlice(t, expectedTable[i], currentTable[i])
	}
}

func CheckMask(t *testing.T, expectedMask, currentMask *utils.Mask) {
	if expectedMask.Size() != currentMask.Size() {
		t.Errorf("Expected size %d but found %d", expectedMask.Size(), currentMask.Size())
	}
	if len(expectedMask.Keys()) != len(currentMask.Keys()) {
		t.Errorf("Number of keys error : Expected %d by found %d", len(expectedMask.Keys()), len(currentMask.Keys()))
	}
	if expectedMask.Default() != currentMask.Default() {
		t.Errorf("Default mask error : Expected %b but found %b", expectedMask.Default(), currentMask.Default())
	}
	for _, key := range currentMask.Keys() {
		expectedValue := expectedMask.Get(key)
		currentValue := currentMask.Get(key)
		if expectedValue != currentValue {
			t.Errorf("Key error %c : Expected %b but found %b", key, expectedValue, currentValue)
		}
	}
}

func CheckMasks(t *testing.T, expectedMasks, currentMasks []*utils.Mask) {
	if len(expectedMasks) != len(currentMasks) {
		t.Errorf("Number of masks error : Expected %d but found %d", len(expectedMasks), len(currentMasks))
	}
	for i := 0; i < len(expectedMasks); i++ {
		CheckMask(t, expectedMasks[i], currentMasks[i])
	}
}

// func CheckMultiMask(t *testing.T, expectedMask, currentMask *utils.MultiMask) {
// 	if len(*currentMask.Size()) != len(*expectedMask.Size()) {
// 		t.Errorf("Number of sizes error : Expected %d but found %d", len(*expectedMask.Size()), len(*currentMask.Size()))
// 	}
// 	length := len(*currentMask.Size())
// 	for i := 0; i < length; i++ {
// 		currentSize := (*currentMask.Size())[i]
// 		expectedSize := (*expectedMask.Size())[i]
// 		if currentSize != expectedSize {
// 			t.Errorf("Expected size %d but found %d", expectedSize, currentSize)
// 		}
// 	}
// 	for _, key := range expectedMask.Keys() {
// 		currentValueL := currentMask.Get(key)
// 		expectedValueL := expectedMask.Get(key)
// 		for i := 0; i < len(expectedValueL); i++ {
// 			currentValue := currentValueL[i]
// 			expectedValue := expectedValueL[i]
// 			if currentValue != expectedValue {
// 				t.Errorf("Key error %c : Expected %v but found %v", key, expectedValueL, currentValueL)
// 				break
// 			}
// 		}
// 	}
// }
