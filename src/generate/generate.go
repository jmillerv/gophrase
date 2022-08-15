package generate

import (
	"github.com/jmillerv/gophrase/config"
	"github.com/jmillerv/gophrase/corpus"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Params struct {
	WordCount    int
	WordList     string
	Capitals     bool
	SpecialChars bool
	Numbers      bool
}

// TODO There is a chance this doesn't capitalize. Need to make a validator.
func Capitals(passphrase []string) string {
	rand.Seed(time.Now().UnixNano())
	var opts int
	capsPassphrase := strings.Join(passphrase[:], "")
	for i := 0; i < len(capsPassphrase); i++ {
		opts = choice()
		if opts == 1 {
			letter := strings.ToUpper(string(capsPassphrase[i]))
			capsPassphrase = strings.Replace(capsPassphrase, string(capsPassphrase[i]), letter, i)
		} else {
			letter := strings.ToLower(string(capsPassphrase[i]))
			capsPassphrase = strings.Replace(capsPassphrase, string(capsPassphrase[i]), letter, i)
		}
	}
	return capsPassphrase
}

func choice() int {
	rand.Seed(time.Now().UnixNano())
	opts := []int{1, 2, 3, 4}
	return opts[rand.Intn(len(opts))]
}

func HandleFlags(p *Params, passphrase []string) string {
	var returnValue string
	if p.SpecialChars == true {
		passphrase = SpecialCharacters(passphrase)

	}
	if p.Numbers == true {
		passphrase = Numbers(passphrase)
	}
	returnValue = strings.Join(passphrase[:], "")
	if p.Capitals == true {
		returnValue = Capitals(passphrase)
	}
	return returnValue
}

func key(wordList string) int {
	var key int
	keySize := keySize(wordList)
	rand.Seed(time.Now().UnixNano())
	die := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < keySize; i++ {
		key = (key * 10) + die[rand.Intn(len(die))]
	}
	return key
}

// keySize determines the size of the key for random generation based on the wordlist used.
func keySize(wordList string) int {
	if wordList == config.EffShort1 || wordList == config.EffShort2 {
		return 4
	} else {
		return 5
	}
}

func Password(p *Params) string {
	var passphrase []string
	for i := 1; i <= p.WordCount; i++ {
		key := key(p.WordList)
		word := corpus.GetWord(key, p.WordList)
		passphrase = append(passphrase, word)
	}
	returnValue := strings.Join(passphrase[:], "")
	if p.Capitals == true || p.SpecialChars == true || p.Numbers == true {
		returnValue = HandleFlags(p, passphrase)
	}
	return returnValue
}

func SpecialCharacters(passphrase []string) []string {
	rand.Seed(time.Now().UnixNano())
	passphrase = append(passphrase, "")
	charCount := rand.Intn(len(passphrase)) + 1 // add a random number of special characters
	for i := 0; i < charCount; i++ {
		char := corpus.GetSpecialChar()
		passphrase = append(passphrase, char)
	}
	// shuffle the special characters so they aren't just at the end of a string array
	rand.Shuffle(len(passphrase), func(i, j int) { passphrase[i], passphrase[j] = passphrase[j], passphrase[i] })
	return passphrase
}

func Numbers(passphrase []string) []string {
	rand.Seed(time.Now().UnixNano())
	charCount := rand.Intn(len(passphrase)) + 1
	for i := 0; i < charCount; i++ {
		num := strconv.Itoa(rand.Intn(100))
		passphrase = append(passphrase, num)
	}
	rand.Shuffle(len(passphrase), func(i, j int) { passphrase[i], passphrase[j] = passphrase[j], passphrase[i] })
	return passphrase
}
