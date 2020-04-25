package command

import 	"github.com/urfave/cli/v2"

var PasswordFlags = []cli.Flag {
	&cli.BoolFlag{
	Name: "capital",
	Aliases: []string{"c"},
	Usage: "Add random capitalization to your passwords",
	},
}
