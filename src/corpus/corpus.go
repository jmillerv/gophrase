package corpus

import (
	"encoding/json"
	"github.com/jmillerv/gophrase/config"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func getSpecialCharList() []byte {
	fileLocation, err := config.Assets.ReadFile("../" + config.Characters)
	if err != nil {
		log.WithField("function", "getSpecialCharList").WithError(err).Fatal("unable to load special characters")
	}
	return fileLocation
}

// GetSpecialChar retrieves a special character from the special_characters.json
func GetSpecialChar() string {
	rand.Seed(time.Now().UnixNano())
	chars := make(map[int]string)
	corpus := getSpecialCharList()
	err := json.Unmarshal(corpus, &chars)
	if err != nil {
		log.WithError(err).Fatal("json.Unmarshal")
	}
	key := rand.Intn(len(chars)) + 1
	return chars[key]
}

func getWordList(wordlist string) []byte {
	fileLocation, err := config.Assets.ReadFile(wordlist)
	if err != nil {
		log.WithError(err).Fatal("unable to read file")
	}
	return fileLocation
}

// GetWord retrieves a word from the passed in wordlist
func GetWord(key int, wordlist string) string {
	words := make(map[int]string)
	corpus := getWordList(wordlist)
	err := json.Unmarshal(corpus, &words)
	if err != nil {
		log.WithField("function", "GetWord").WithError(err).Fatal("json.Unmarshal")
	}
	word := words[key]
	return word
}

// SetWordList updates the LoadedConfig.Wordlist property
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

// PrintWordListOptions prints wordlist documentation to the console.
func PrintWordListOptions() []byte {
	fileLocation, err := config.Assets.ReadFile(config.ListOptions)
	if err != nil {
		log.WithField("function", "PrintWordListOptions").WithError(err).Error("unable to read file")
	}
	return fileLocation
}
