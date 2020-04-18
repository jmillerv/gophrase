package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func getWord(key int) string {
	words := make(map[int]string)
	corpus, err := ioutil.ReadFile("./assets/shortwords.json")
	if err != nil {
		log.Println("Error reading json file")
	}
	err = json.Unmarshal(corpus, &words)
	if err != nil {
		log.Println("Error unmarshalling json")
	}
	word := words[key]
	return word
}

func makeKey() int {
	rand.Seed(time.Now().UnixNano())
	min := 1111 // lowest key number in assets/shortwords.json
	max := 6666 // highest key number in assets/shortwords.json
	key := rand.Intn((max - min) + min)
	return key
}

func GeneratePassword(wordCount int) string {
	var password []string
	var key int
	for i := 0; i < wordCount; i++ {
		key = makeKey()
		word := getWord(key)
		password = append(password, word)
		fmt.Println(password)
	}
	return strings.Join(password[:], "")
}
