package main

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

var onlyOnce sync.Once
func makeKey() int {
	var key int
	onlyOnce.Do(func(){
		rand.Seed(time.Now().UnixNano())
	})
	die := []int{1,2,3,4,5,6}
	a := die[rand.Intn(len(die))] * 1000
	b := die[rand.Intn(len(die))] * 100
	c := die[rand.Intn(len(die))] * 10
	d := die[rand.Intn(len(die))]
	key += a+b+c+d
	return key
}

func GeneratePassword(wordCount int) string {
	var password []string
	for i := 1; i <= wordCount; i++ {
		key := makeKey()
		word := GetWord(key)
		password = append(password, word)
	}
	return strings.Join(password[:], "")
}
