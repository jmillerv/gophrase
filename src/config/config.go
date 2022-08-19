package config

import (
	"embed"
	"github.com/jmillerv/go-utilities/format"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

// Static asset constants
const (
	EffLarge             = "assets/wordlists/eff_large_wordlist.json"
	EffShort1            = "assets/wordlists/eff_short_wordlist_1.json"
	EffShort2            = "assets/wordlists/eff_short_wordlist_2.json"
	Reinhold             = "assets/wordlists/reinhold_wordlist.json"
	Characters           = "assets/wordlists/special_characters.json"
	ListOptions          = "assets/list_options.txt"
	DefaultConfiguration = "assets/config/config.yaml"
	CustomConfiguration  = "config.yaml"
)

// Assets is the embed.FS file system to access embedded files throughout the application.
var Assets embed.FS

// LoadedConfig is the configuration for gophrase throuhgout the application.
var LoadedConfig = new(Config)

// Config holds the values from a configuration file for default generation settings.
type Config struct {
	WordCount int    `yaml:"WordCount"`
	WordList  string `yaml:"WordList"`
	Capital   bool   `yaml:"Capital"`
	Special   bool   `yaml:"Special"`
	Number    bool   `yaml:"Number"`
}

// PrintConfig outputs the current config to the console
func (c *Config) PrintConfig() {
	log.Print(format.StructToIndentedString(c))
}

// SetConfig takes in a config and writes it to the configuration file.
func SetConfig(config *Config) {
	file, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(CustomConfiguration, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// LoadConfig reads a configuration file and loads it into the LoadedConfig variable.
func LoadConfig() {
	customFileLocation, err := os.ReadFile(CustomConfiguration)
	if customFileLocation != nil {
		err = yaml.Unmarshal(customFileLocation, &LoadedConfig)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	defaultFileLocation, err := Assets.ReadFile(DefaultConfiguration)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(defaultFileLocation, &LoadedConfig)
	if err != nil {
		log.Fatal(err)
	}
}
