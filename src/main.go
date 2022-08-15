package main

import (
	"embed"
	_ "embed"
	"github.com/jmillerv/gophrase/command"
	"github.com/jmillerv/gophrase/config"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// App information constants
const (
	AppName      = "Go Phrase"
	AppUsage     = "CLI for generating secure, memorable passwords"
	AppVersion   = "v2.0.0 | release: 8.14.22"
	AppUsageText = "gophrase gen [word count] [word list] [--flags]"
)

//go:embed assets/*
var staticFS embed.FS

func main() {
	config.Assets = staticFS
	app := *cli.NewApp()
	app.Name = AppName
	app.Usage = AppUsage
	app.UsageText = AppUsageText
	app.Version = AppVersion
	app.HideHelp = false
	app.HideVersion = false
	app.Commands = command.Commands
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
