package generate

import (
	"github.com/gophrase/internal/corpus"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var onlyOnce sync.Once

func key(wordList string) int {
	var key int
	keySize := keySize(wordList)
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})
	die := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < keySize; i++ {
		key = (key * 10) + die[rand.Intn(len(die))]
	}
	return key
}

func keySize(wordList string) int {
	if wordList == "a" || wordList == "b" {
		return 4
	} else {
		return 5
	}
}

func Password(wordCount int, wordList string) string {
	var password []string
	for i := 1; i <= wordCount; i++ {
		key := key(wordList)
		word := corpus.GetWord(key, wordList)
		password = append(password, word)
	}
	return strings.Join(password[:], "")
}
