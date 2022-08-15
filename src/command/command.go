package command

import (
	"fmt"
	"github.com/jmillerv/gophrase/config"
	"github.com/jmillerv/gophrase/corpus"
	"github.com/jmillerv/gophrase/entropy"
	"github.com/jmillerv/gophrase/generate"
	"github.com/urfave/cli/v2"
	"strconv"
)

const (
	defaultWordCount = 5
)

// Commands hold an array of cli.Commands which are used to run the CLI application.
var Commands = []*cli.Command{
	{
		Name:    "generate",
		Aliases: []string{"gen"},
		Usage:   "gen [int]",
		Action: func(c *cli.Context) error {
			// TODO input validator to clean up section
			config.LoadConfig()
			p := &generate.Params{}
			p.WordCount, _ = strconv.Atoi(c.Args().Get(0))
			if p.WordCount == 0 {
				p.WordCount = config.LoadedConfig.WordCount
			}
			p.WordList = corpus.SetWordList(c.Args().Get(1))
			if p.WordList == "" {
				p.WordList = config.LoadedConfig.WordList
			}
			if c.Bool("capital") || config.LoadedConfig.Capital {
				p.Capitals = true
			} else {
				p.Capitals = false
			}
			if c.Bool("special") || config.LoadedConfig.Special {
				p.SpecialChars = true
			} else {
				p.SpecialChars = false
			}
			if c.Bool("number") || config.LoadedConfig.Number {
				p.Numbers = true
			} else {
				p.Numbers = false
			}
			password := generate.Password(p)
			entropy.PrintEntropy(password)
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "capital",
				Aliases: []string{"c"},
				Usage:   "Add random capitalization to your passwords",
			},
			&cli.BoolFlag{
				Name:    "special",
				Aliases: []string{"s"},
				Usage:   "Add random special characters to your passwords",
			},
			&cli.BoolFlag{
				Name:    "number",
				Aliases: []string{"n"},
				Usage:   "Add numbers to your passwords",
			},
		},
	},
	{
		Name:    "wordlist-options",
		Aliases: []string{"opts"},
		Usage:   "View the wordlist options for passphrase generation",
		Action: func(c *cli.Context) error {
			fmt.Print(string(corpus.PrintWordListOptions()))
			return nil
		},
	},
	{
		Name:    "set-defaults",
		Aliases: []string{"sd"},
		Usage:   "Set default options for word count and word list",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "capital",
				Aliases: []string{"c"},
				Usage:   "Add random capitalization to your passwords",
			},
			&cli.BoolFlag{
				Name:    "special",
				Aliases: []string{"s"},
				Usage:   "Add random special characters to your passwords",
			},
			&cli.BoolFlag{
				Name:    "number",
				Aliases: []string{"n"},
				Usage:   "Add numbers to your passwords",
			},
		},
		Action: func(c *cli.Context) error {
			conf := &config.Config{}
			conf.WordCount, _ = strconv.Atoi(c.Args().Get(0))
			if conf.WordCount == 0 {
				conf.WordCount = defaultWordCount
			}
			conf.WordList = corpus.SetWordList(c.Args().Get(1))
			if conf.WordList == "" {
				conf.WordList = config.EffShort1
			}
			if c.Bool("capital") {
				conf.Capital = true
			}
			if c.Bool("special") {
				conf.Special = true
			}
			if c.Bool("number") {
				conf.Number = true
			}
			config.SetConfig(conf)
			return nil
		},
	},
	{
		Name:    "list-defaults",
		Aliases: []string{"ld"},
		Usage:   "Print default options for word count and word list",
		Action: func(c *cli.Context) error {
			config.LoadConfig()
			config.LoadedConfig.PrintConfig()
			return nil
		},
	},
}
