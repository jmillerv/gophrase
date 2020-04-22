package corpus

import (
	"github.com/gophrase/internal"
	"testing"
)

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
		fifth string
	}

	var testValues = tables{"a", "b", "c", "d", ""}

	if  setWordList(testValues.first) != internal.EFF_SHORT_2 {
		t.Errorf("setWordList() failed expected 'eff_short_wordlist_2_0.json', but got %s", setWordList(testValues.first))
	}
	if  setWordList(testValues.second) != internal.EFF_SHORT_1 {
		t.Errorf("setWordList() failed expected 'eff_short_wordlist_1.json', but got %s", setWordList(testValues.second))
	}
	if  setWordList(testValues.third) != internal.EFF_LARGE {
		t.Errorf("setWordList() failed expected 'eff_large_wordlist.json', but got %s", setWordList(testValues.third))
	}
	if  setWordList(testValues.fourth) != internal.REINHOLD {
		t.Errorf("setWordList() failed expected 'reinhold_wordlist.json', but got %s", setWordList(testValues.fourth))
	}
	if  setWordList(testValues.fifth) != internal.EFF_SHORT_2 {
		t.Errorf("setWordList() failed expected 'eff_short_wordlist_2_0.json', but got %s", setWordList(testValues.fifth))
	}
}