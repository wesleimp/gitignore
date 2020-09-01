package cli

import (
	"github.com/urfave/cli/v2"

	"github.com/wesleimp/gitignore/cmd/cli/commands/append"
	"github.com/wesleimp/gitignore/cmd/cli/commands/generate"
	"github.com/wesleimp/gitignore/cmd/cli/commands/list"
)

// Execute cli
func Execute(version string, args []string) error {
	app := &cli.App{
		Name:    "gitignore",
		Usage:   "Gitignore generator",
		Version: version,
		Authors: []*cli.Author{{
			Name:  "Weslei Juan Moser Pereira",
			Email: "wesleimsr@gmail.com",
		}},
		Commands: []*cli.Command{
			list.Command,
			generate.Command,
			append.Command,
		},
	}

	return app.Run(args)
}
