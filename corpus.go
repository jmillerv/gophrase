package main

import (
	"encoding/json"
	"github.com/gobuffalo/packr/v2"
	"log"
)

func getCorpus() []byte {
	box := packr.New("assets", "./assets")
	fileLocation, err := box.Find("eff_short_wordlist_2_0.json")
	if err != nil {
		log.Fatal(err)
	}
	return fileLocation
}

func GetWord(key int) string {
	words := make(map[int]string)
	corpus := getCorpus()
	err := json.Unmarshal(corpus, &words)
	if err != nil {
		log.Fatal(err)
	}
	word := words[key]
	return word
}
