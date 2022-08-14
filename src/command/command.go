package command

import (
	"fmt"
	"github.com/jmillerv/gophrase/config"
	"github.com/jmillerv/gophrase/corpus"
	"github.com/jmillerv/gophrase/entropy"
	"github.com/jmillerv/gophrase/generate"
	"github.com/jmillerv/gophrase/handlers"
	"github.com/urfave/cli/v2"
	"strconv"
)

var Commands = []*cli.Command{
	{
		Name:    "generate",
		Aliases: []string{"gen"},
		Usage:   "gen [int]",
		Action: func(c *cli.Context) error {
			// TODO input validator to clean up section
			config.LoadConfigDefaults()
			p := generate.Params{}
			p.WordCount, _ = strconv.Atoi(c.Args().Get(0))
			if p.WordCount == 0 {
				p.WordCount = config.Defaults.WordCount
			}
			p.WordList = c.Args().Get(1)
			if p.WordList == "" {
				p.WordList = config.Defaults.WordList
			}
			if c.Bool("capital") {
				p.Capitals = true
			} else {
				p.Capitals = false
			}
			if c.Bool("special") {
				p.SpecialChars = true
			} else {
				p.SpecialChars = false
			}
			if c.Bool("number") {
				p.Numbers = true
			} else {
				p.Numbers = false
			}
			password := generate.Password(&p)
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
		Action: func(c *cli.Context) error {
			p := generate.Params{}
			p.WordCount, _ = strconv.Atoi(c.Args().Get(0))
			if p.WordCount == 0 {
				p.WordCount = config.Defaults.WordCount
			}
			p.WordList = corpus.SetWordList(c.Args().Get(1))
			if p.WordList == "" {
				p.WordList = config.Defaults.WordList
			}
			config.SetConfigDefaults(config.Defaults, p.WordCount, p.WordList)
			return nil
		},
	},
	{
		Name:    "list-defaults",
		Aliases: []string{"ld"},
		Usage:   "Print default options for word count and word list",
		Action: func(c *cli.Context) error {
			config.LoadConfigDefaults()
			config.PrintConfigDefaults(config.Defaults)
			return nil
		},
	},
	{
		Name:    "run-server",
		Aliases: []string{"serve"},
		Usage:   "Start a gophrase server",
		Action: func(c *cli.Context) error {
			handlers.RunServer()
			return nil
		},
	},
}
