package config

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/gophrase/internal"
	"io/ioutil"
	"log"
)

type Default struct {
	WordCount int    `json:"wordCount"`
	WordList  string `json:"wordList"`
}

var Assets = packr.New("assets", "../../assets")
var Defaults = new(Default)

func SetConfigDefaults(config *Default, count int, list string) {
	fileLocation := internal.DEFAULTS //TODO create better implementation
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
	fileLocation, err := Assets.Find(internal.DEFAULTS)
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
