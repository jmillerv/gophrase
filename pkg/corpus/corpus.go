package corpus

import (
	"encoding/json"
	"github.com/gophrase/internal"
	"github.com/gophrase/internal/config"
	"log"
	"math/rand"
	"time"
)

func getSpecialCharList() []byte {
	fileLocation, err := config.Assets.Find(internal.CHARACTERS)
	if err != nil {
		log.Fatal(err)
	}
	return fileLocation
}

func GetSpecialChar() string {
	rand.Seed(time.Now().UnixNano())
	chars := make(map[int]string)
	corpus := getSpecialCharList()
	err := json.Unmarshal(corpus, &chars)
	if err != nil {
		log.Fatal(err)
	}
	key := rand.Intn(len(chars)) + 1
	return chars[key]
}

func getWordList(wordlist string) []byte {
	fileLocation, err := config.Assets.Find(setWordList(wordlist))
	if err != nil {
		log.Fatal(err)
	}
	return fileLocation
}

func GetWord(key int, wordlist string) string {
	words := make(map[int]string)
	corpus := getWordList(wordlist)
	err := json.Unmarshal(corpus, &words)
	if err != nil {
		log.Fatal(err)
	}
	word := words[key]
	return word
}

func setWordList(wordlist string) string {
	// This is organized by my personal preference.
	switch list := wordlist; list {
	case "a":
		return internal.EFF_SHORT_2
	case "b":
		return internal.EFF_SHORT_1
	case "c":
		return internal.EFF_LARGE
	case "d":
		return internal.REINHOLD
	default:
		return config.Defaults.WordList
	}
}

func PrintWordListOptions() []byte {
	fileLocation, err := config.Assets.Find(internal.LIST_OPTIONS)
	if err != nil {
		log.Fatal(err)
	}
	return fileLocation
}
