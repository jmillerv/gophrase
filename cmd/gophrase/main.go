package main

import (
	"github.com/gophrase/internal"
	"github.com/gophrase/internal/argument"
	"github.com/urfave/cli/v2"
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
	app.Commands = argument.Commands
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
