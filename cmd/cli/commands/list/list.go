package list

import (
	"strings"

	"github.com/apex/log"
	"github.com/urfave/cli/v2"
	"github.com/wesleimp/gitignore/internal/gitignore"
	"github.com/wesleimp/gitignore/internal/logging"
)

// Command list al lavailable templates
var Command = &cli.Command{
	Name:    "list",
	Usage:   "List all available templates",
	Aliases: []string{"l"},
	Action: func(c *cli.Context) error {
		err := logging.Log("listing templates...", listTemplates)
		return err
	},
}

func listTemplates() error {
	tt, err := gitignore.GetTemplates()
	if err != nil {
		return err
	}

	tt.Sort()
	for _, t := range tt {
		var name = strings.ReplaceAll(t.Name, ".gitignore", "")
		log.Info(name)
	}
	return nil
}
