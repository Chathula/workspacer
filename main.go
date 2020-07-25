package main

import (
	"fmt"
	"os"

	"github.com/chathula/workspacer/app"
	"github.com/chathula/workspacer/command"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {
	cliApp := cli.NewApp()

	cliApp.Name = "workspacer"
	cliApp.Description = "Setup DEV workspace without hassle"
	cliApp.Usage = "Automated DEV Workspace setup"
	cliApp.Version = app.GetVersion()
	cliApp.Authors = []*cli.Author{
		{
			Name:  "Chathula Sampath",
			Email: "schathula@gmail.com",
		},
	}

	cliApp.Commands = []*cli.Command{
		{
			Name:      "install",
			Usage:     "Install/Set/Clone all services, tools, environment variables and git repositories",
			ArgsUsage: "<file>|<gist-url>",
			Action: func(c *cli.Context) error {
				if c.Args().Len() == 0 {
					return fmt.Errorf("require argument <%s>", color.RedString("file|gist-url"))
				}
				return command.Install(c.Args().First())
			},
		},
	}

	cli.AppHelpTemplate = `NAME:
    {{.Name}} - {{.Usage}}
    {{if len .Authors}}
AUTHOR:
    {{range .Authors}}{{ . }}{{end}}
    {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}} {{ .ArgsUsage }}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
    {{range .VisibleFlags}}{{.}}
    {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
    {{.Copyright}}
    {{end}}{{if .Version}}
VERSION:
    {{.Version}}
    {{end}}
EXAMPLES:
    {{.Name}} install <file.yml>
    {{.Name}} install <gist-url>
SOURCE CODE:
    https://github.com/chathula/workspacer
`

	if err := cliApp.Run(os.Args); err != nil {
		if os.Getenv("DEBUG") != "" {
			fmt.Printf("%+v\n", err)
		} else {
			fmt.Println(err.Error())
			fmt.Printf("run with environment variables %s to print more information\n", color.GreenString("DEBUG=1"))
		}
		os.Exit(1)
	}
}
