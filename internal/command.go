package internal

import (
	"fmt"
	"github.com/urfave/cli"
	"strconv"
)

var Commands = []cli.Command{
	{
		Name:    "GeneratePassword",
		Aliases: []string{"gen"},
		Usage:   "gen <int>",
		Action: func(c *cli.Context) error {
			i, _ := strconv.Atoi(c.Args().Get(0))
			wl := c.Args().Get(1)
			if wl == "" {
				wl = "a"
			}
			password := GeneratePassword(i, wl)
			fmt.Println(password)
			return nil
		},
	},
}
