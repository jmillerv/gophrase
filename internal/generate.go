package internal

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

var onlyOnce sync.Once

func makeKey() int {
	var key int
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})
	die := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 4; i++ {
		key = (key * 10) + die[rand.Intn(len(die))]
	}
	return key
}

func GeneratePassword(wordCount int) string {
	var password []string
	for i := 1; i <= wordCount; i++ {
		key := makeKey()
		word := GetWord(key)
		password = append(password, word)
	}
	return strings.Join(password[:], "")
}
