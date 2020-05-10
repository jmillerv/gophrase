package config

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/gophrase/internal"
	"log"
	"strconv"
)

type Default struct {
	WordCount int
	WordList  string
}

var Assets = packr.New("assets", "../../assets")

var Defaults = Default{
	WordCount: 3,
	WordList:  internal.EFF_SHORT_2,
}

func SetConfigDefaults(Defaults *Default, count int, list string) {

}

func LoadConfigDefaults() {
	fileLocation, err := Assets.Find(internal.DEFAULTS)
	if err != nil {
		log.Fatal(err)
	}
	settings := make(map[int]string)
	err = json.Unmarshal(fileLocation, &settings)
	if err != nil {
		log.Fatal(err)
	}
	Defaults.WordCount, _ = strconv.Atoi(settings[0])
	Defaults.WordList = settings[1]
}

func PrintConfigDefaults(config Default) {
	fmt.Printf("Word Count: %d \nWord List: %s \n", config.WordCount, config.WordList)
}
