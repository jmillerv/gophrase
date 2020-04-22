package generate

import "testing"

func TestKeySize (t *testing.T){
	type tables struct {
		first string
		second string
	}

	var testValues = tables{"a", "d"}

	if keySize(testValues.first) != 4 {
		t.Errorf("keySize() failed expected 4, got %v", keySize(testValues.first))
	}

	if keySize(testValues.second) != 5 {
		t.Errorf("keySize() failed expected 5, got %v", keySize(testValues.second))
	}
}
