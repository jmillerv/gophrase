package corpus

import (
	"encoding/json"
	"github.com/jmillerv/gophrase/config"
	"log"
	"math/rand"
	"time"
)

func getSpecialCharList() []byte {
	fileLocation, err := config.Assets.ReadFile(config.Characters)
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
	fileLocation, err := config.Assets.ReadFile(SetWordList(wordlist))
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
	switch list := wordlist; list {
	case "a":
		return config.EffShort1
	case "b":
		return config.EffShort2
	case "c":
		return config.EffLarge
	case "d":
		return config.Reinhold
	default:
		return config.LoadedConfig.WordList
	}
}

func PrintWordListOptions() []byte {
	fileLocation, err := config.Assets.ReadFile(config.ListOptions)
	if err != nil {
		log.Fatal(err)
	}
	return fileLocation
}
