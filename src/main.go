package main

import (
	"github.com/jmillerv/gophrase/src/command"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// App information constants
const APP_NAME = "Go Phrase"
const APP_USAGE = "CLI for generating secure, memorable passwords"
const APP_VERSION = "v2.0.0 | release: 3.13.22"
const APP_USAGETEXT = "gophrase gen [word count] [word list] [--flags]"

// go:embed

func main() {
	app := *cli.NewApp()
	app.Name = APP_NAME
	app.Usage = APP_USAGE
	app.UsageText = APP_USAGETEXT
	app.Version = APP_VERSION
	app.HideHelp = false
	app.HideVersion = false
	app.Commands = command.Commands
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
