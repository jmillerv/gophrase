package command

import (
	"fmt"
	"github.com/gophrase/pkg/generate"
	"github.com/urfave/cli/v2"
	"strconv"
)

var Commands = []*cli.Command{
	{
		Name:    "Password",
		Aliases: []string{"gen"},
		Usage:   "gen <int>",
		Action: func(c *cli.Context) error {
			var i int
			var wl string
			i, _ = strconv.Atoi(c.Args().Get(0))
			if i == 0 {
				i = 3
			}
			wl = c.Args().Get(1)
			if wl == "" {
				wl = "a"
			}
			var capital bool
			if !c.Bool("capital") {
				capital = false
			} else {
				capital = true
			}
			password := generate.Password(i, wl, capital)
			fmt.Println(password)
			return nil
		},
		Flags: PasswordFlags,
	},
}
