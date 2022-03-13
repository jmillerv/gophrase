package corpus

import (
	"encoding/json"
	"github.com/jmillerv/gophrase/config"
	"log"
	"math/rand"
	"time"
)

func getSpecialCharList() []byte {
	fileLocation, err := config.Assets.Find(config.CHARACTERS)
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
	fileLocation, err := config.Assets.Find(SetWordList(wordlist))
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

func SetWordList(wordlist string) string {
	// This is organized by my personal preference.
	switch list := wordlist; list {
	case "a":
		return config.EFF_SHORT_2
	case "b":
		return config.EFF_SHORT_1
	case "c":
		return config.EFF_LARGE
	case "d":
		return config.REINHOLD
	default:
		return config.Defaults.WordList
	}
}

func PrintWordListOptions() []byte {
	fileLocation, err := config.Assets.Find(config.LIST_OPTIONS)
	if err != nil {
		log.Fatal(err)
	}
	return fileLocation
}
