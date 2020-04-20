package main

import (
	"github.com/gophrase/internal"
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

	cli.AppHelpTemplate = `
		{{.Name}} - {{.Usage}}
	
		{{.UsageText}}{{if len .Authors}}
	
		{{range .Authors}}{{ . }}{{end}} {{end}}{{if .Commands}} {{.Version}}{{end}}
	`
	app.Commands = internal.Commands // this is weird because go build tells me the opposite.
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
