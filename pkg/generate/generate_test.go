package generate

import (
	"fmt"
	"math"
	"testing"
	"unicode"
)

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

func TestKey (t *testing.T){
	type tables struct {
		first string
		second string
	}

	var testValues = tables{"a", "d"}
	key1 := key(testValues.first)
	// check if key is 4 digit number
	if getDigit(key1, 4) == 0 && getDigit(key1, 5) == 0 {
		t.Errorf("key() failed expected key with 4 digits")
	}
	// check if key is 5 digit number
	key2 := key(testValues.second)
	if getDigit(key2, 5) == 0 {
		t.Errorf("key() failed expected key with 5 digits")
	}
}

func TestBasicPassword(t *testing.T) {
	passphrase := Password(3, "d", false)
	if len(passphrase) <= 0 {
		t.Errorf("Password() failed expected string, got empty return")
	}
}

func TestCapitalPassword(t *testing.T){
	passphrase := Password(5, "a", true)
	runes := []rune(passphrase)
	var results []string
	fmt.Println(passphrase)
	for i := 0; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) == true {
			results = append(results,"true")
		} else {
			results = append(results,"false")
		}
	}
	if !contains(results, "true"){
		t.Errorf("Capital Password failed expected uppercase letter got %s", passphrase)
	}
}

func getDigit(num, length int) int {
	r := num % int(math.Pow(10, float64(length)))
	return r / int(math.Pow(10, float64(length-1)))
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}