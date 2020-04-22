package main

import (
	"github.com/gophrase/internal"
	"github.com/gophrase/internal/command"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := *cli.NewApp()
	app.Name = internal.APP_NAME
	app.Usage = internal.APP_USAGE
	app.UsageText = internal.APP_USAGETEXT
	app.Version = internal.APP_VERSION
	app.HideHelp = false
	app.HideVersion = false

	app.Commands = command.Commands // this is weird because go build tells me the opposite.
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
