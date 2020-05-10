package argument

import (
	"fmt"
	"github.com/gophrase/internal/config"
	"github.com/gophrase/pkg/corpus"
	"github.com/gophrase/pkg/entropy"
	"github.com/gophrase/pkg/generate"
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
		Name:    "Wordlist Options",
		Aliases: []string{"opts"},
		Usage:   "View the wordlist options for passphrase generation",
		Action: func(c *cli.Context) error {
			fmt.Print(string(corpus.PrintWordListOptions()))
			return nil
		},
	},
	{
		Name:    "List Defaults",
		Aliases: []string{"ld"},
		Usage:   "Set default options for word count and word list",
		Action: func(c *cli.Context) error {
			config.PrintConfigDefaults(config.Defaults)
			return nil
		},
	},
}
