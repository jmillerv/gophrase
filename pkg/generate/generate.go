package generate

import (
	"github.com/gophrase/pkg/corpus"
	"math/rand"
	"strings"
	"time"
)


func key(wordList string) int {
	var key int
	keySize := keySize(wordList)
	rand.Seed(time.Now().UnixNano())
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
// There is a chance this doesn't capitalize. Need to make a validator.
func capitals(password []string) string {
	rand.Seed(time.Now().UnixNano())
	var opts int
	capsPassword := strings.Join(password[:], "")
	for i := 0; i < len(capsPassword); i++ {
		opts = choice()
		if opts == 1 {
			letter := strings.ToUpper(string(capsPassword[i]))
			capsPassword = strings.Replace(capsPassword, string(capsPassword[i]), letter, i)
		} else {
			letter := strings.ToLower(string(capsPassword[i]))
			capsPassword = strings.Replace(capsPassword, string(capsPassword[i]), letter, i)
		}
	}
	return capsPassword
}

func choice() int {
	rand.Seed(time.Now().UnixNano())
	opts := []int{1,2,3,4}
	return opts[rand.Intn(len(opts))]
}

func Password(wordCount int, wordList string, caps bool) string {
	var password []string
	for i := 1; i <= wordCount; i++ {
		key := key(wordList)
		word := corpus.GetWord(key, wordList)
		password = append(password, word)
	}
	returnValue := strings.Join(password[:], "")
	if caps == true {
		returnValue = capitals(password)
	}
	return returnValue
}

