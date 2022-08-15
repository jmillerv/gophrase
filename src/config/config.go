package config

import (
	"embed"
	"github.com/jmillerv/go-utilities/format"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// Static asset constants
const (
	EffLarge      = "assets/wordlists/eff_large_wordlist.json"
	EffShort1     = "assets/wordlists/eff_short_wordlist_1.json"
	EffShort2     = "assets/wordlists/eff_short_wordlist_2.json"
	Reinhold      = "assets/wordlists/reinhold_wordlist.json"
	Characters    = "assets/wordlists/special_characters.json"
	ListOptions   = "assets/list_options.txt"
	Configuration = "assets/config/config.yaml"
)

var Assets embed.FS
var LoadedConfig = new(Config)

type Config struct {
	WordCount int    `yaml:"WordCount"`
	WordList  string `yaml:"WordList"`
	Capital   bool   `yaml:"Capital"`
	Special   bool   `yaml:"Special"`
	Number    bool   `yaml:"Number"`
}

func (c *Config) PrintConfig() {
	log.Print(format.StructToIndentedString(c))
}

func SetConfigDefaults(config *Config) {
	fileLocation := Configuration // TODO create better implementation
	file, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(fileLocation, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadConfig() {
	fileLocation, err := Assets.ReadFile(Configuration)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(fileLocation, &LoadedConfig)
	if err != nil {
		log.Fatal(err)
	}
}
