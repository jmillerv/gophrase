package corpus

import (
	"encoding/json"
	"github.com/gobuffalo/packr/v2"
	"github.com/gophrase/internal"
	"log"
)

func getCorpus(wordlist string) []byte {
	box := packr.New("assets", "../assets")
	fileLocation, err := box.Find(setWordList(wordlist))
	if err != nil {
		log.Fatal(err)
	}
	return fileLocation
}

func GetWord(key int, wordlist string) string {
	words := make(map[int]string)
	corpus := getCorpus(wordlist)
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
		return internal.EFF_SHORT_2
	}
}
