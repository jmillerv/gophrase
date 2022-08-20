package corpus

import (
	"github.com/jmillerv/gophrase/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetWordList(t *testing.T) {
	type tables struct {
		first  string
		second string
		third  string
		fourth string
	}

	var testValues = tables{"a", "b", "c", "d"}
	assert.ObjectsAreEqual(testValues.first, config.EffShort1)
	assert.ObjectsAreEqual(testValues.second, config.EffShort2)
	assert.ObjectsAreEqual(testValues.third, config.EffLarge)
	assert.ObjectsAreEqual(testValues.fourth, config.Reinhold)
}

func TestGetWord(t *testing.T) {
	type args struct {
		key      int
		wordlist string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success: Retrieves Word From Wordlist",
			args: args{
				key:      2251,
				wordlist: "a",
			},
			want: "dandy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			word := GetWord(tt.args.key, tt.args.wordlist)
			assert.ObjectsAreEqual("dandy", word)
		})
	}
}

// TODO write TestGetWordList and test the bytes
