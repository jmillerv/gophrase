package config

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"io/ioutil"
	"log"
)

// Static asset constants
const EFF_LARGE = "eff_large_wordlist.json"
const EFF_SHORT_1 = "eff_short_wordlist_1.json"
const EFF_SHORT_2 = "eff_short_wordlist_2_0.json"
const REINHOLD = "reinhold_wordlist.json"
const CHARACTERS = "special_characters.json"
const LIST_OPTIONS = "list_options.txt"
const DEFAULTS = "defaults.json"

type Default struct {
	WordCount int    `json:"wordCount"`
	WordList  string `json:"wordList"`
}

var Assets = packr.New("assets", "../../assets")
var Defaults = new(Default)

func SetConfigDefaults(config *Default, count int, list string) {
	fileLocation := DEFAULTS // TODO create better implementation
	config.WordCount = count
	config.WordList = list
	file, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("../../assets/"+fileLocation, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadConfigDefaults() {
	fileLocation, err := Assets.Find(DEFAULTS)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(fileLocation, &Defaults)
	if err != nil {
		log.Fatal(err)
	}
}

func PrintConfigDefaults(config *Default) {
	fmt.Printf("Word Count: %d \nWord List: %s \n", config.WordCount, config.WordList)
}
