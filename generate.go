package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

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
	for _, n := range numbers {
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
		word := GetWord(key)
		password = append(password, word)
	}
	return strings.Join(password[:], "")
}
