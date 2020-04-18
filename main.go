package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := *cli.NewApp()
	app.Name = APP_NAME
	app.Usage = APP_USAGE
	app.UsageText = APP_USAGETEXT
	app.Version = APP_VERSION
	app.HideHelp = false
	app.HideVersion = false

	cli.AppHelpTemplate = `
		{{.Name}} - {{.Usage}}
	
		{{.UsageText}}{{if len .Authors}}
	
		{{range .Authors}}{{ . }}{{end}} {{end}}{{if .Commands}} {{.Version}}{{end}}
	`
	app.Commands = Commands
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
