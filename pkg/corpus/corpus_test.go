package corpus

import (
	"github.com/gophrase/internal"
	"testing"
)

// TODO write TestGetWordList and test the bytes
func TestGetWord(t *testing.T) {
	key := 2215
	wordlist := "a"
	word := GetWord(key, wordlist)
	if word != "dandelion" {
		t.Errorf("GetWord() failed, expected key 2215 to get 'dandelion', got %s", word)
	}

}

func TestSetWordList(t *testing.T) {
	type tables struct {
		first  string
		second string
		third  string
		fourth string
		fifth  string
	}

	var testValues = tables{"a", "b", "c", "d", ""}

	if SetWordList(testValues.first) != internal.EFF_SHORT_2 {
		t.Errorf("SetWordList() failed expected 'eff_short_wordlist_2_0.json', but got %s", SetWordList(testValues.first))
	}
	if SetWordList(testValues.second) != internal.EFF_SHORT_1 {
		t.Errorf("SetWordList() failed expected 'eff_short_wordlist_1.json', but got %s", SetWordList(testValues.second))
	}
	if SetWordList(testValues.third) != internal.EFF_LARGE {
		t.Errorf("SetWordList() failed expected 'eff_large_wordlist.json', but got %s", SetWordList(testValues.third))
	}
	if SetWordList(testValues.fourth) != internal.REINHOLD {
		t.Errorf("SetWordList() failed expected 'reinhold_wordlist.json', but got %s", SetWordList(testValues.fourth))
	}
	if SetWordList(testValues.fifth) != internal.EFF_SHORT_2 {
		t.Errorf("SetWordList() failed expected 'eff_short_wordlist_2_0.json', but got %s", SetWordList(testValues.fifth))
	}
}
