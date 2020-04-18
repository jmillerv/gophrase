package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func getWord(key int) string {
	words := make(map[int]string)
	corpus, err := ioutil.ReadFile("./assets/eff_short_wordlist_2_0.json")
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
	// min & max are determined by the upper and lower bound of a six sided die.
	min := 1
	max := 6
	var keys []string
	var numbers []int
	for i := 1; len(numbers) <= 3; i++ {
		number := rand.Intn((max - min + 1) + min)
		if number > 0 {
			numbers = append(numbers, number)
		}
	}
	for _, n := range numbers{
		keys = append(keys, strconv.Itoa(n))
	}
	key, _ := strconv.Atoi(strings.Join(keys[:], ""))
	return key
}

func GeneratePassword(wordCount int) string {
	var password []string
	var key int
	for i := 1; i <= wordCount; i++ {
		key = makeKey()
		word := getWord(key)
		password = append(password, word)
	}
	return strings.Join(password[:], "")
}
