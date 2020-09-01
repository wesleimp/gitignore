package append

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/wesleimp/gitignore/internal/gitignore"
	"github.com/wesleimp/gitignore/internal/logging"
)

// Command add
var Command = &cli.Command{
	Name:    "append",
	Usage:   "Append to the existing .gitignore file",
	Aliases: []string{"a"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "path",
			Aliases: []string{"p"},
			Usage:   "Path to .gitignore file",
			Value:   ".",
		},
		&cli.StringSliceFlag{
			Name:    "lang",
			Aliases: []string{"l"},
			Usage:   "Language template",
		},
		&cli.StringSliceFlag{
			Name:    "text",
			Aliases: []string{"t"},
			Usage:   "Custom text",
		},
	},
	Action: func(c *cli.Context) error {
		return logging.Log("appending...", func() error {
			file := filepath.Join(c.String("path"), ".gitignore")
			if _, err := os.Stat(file); os.IsNotExist(err) {
				return errors.New(`the .gitignore file does not exists, use the "generate" command to create it`)
			}

			f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return errors.Wrap(err, "couldn't open the .gitignore file")
			}
			defer f.Close()

			if c.IsSet("lang") {
				log.Info("appending langs")
				content, err := gitignore.DownloadTemplates(c.StringSlice("lang"))
				if err != nil {
					return errors.Wrap(err, "error appending languages")
				}

				f.WriteString(content)
			}

			if c.IsSet("text") {
				log.Info("appending text")
				content := appendText(c.StringSlice("text"))
				f.WriteString(content)
			}

			return nil
		})
	},
}

func appendText(texts []string) string {
	var buff bytes.Buffer

	for _, t := range texts {
		buff.WriteString(t)
		buff.WriteString("\n")
	}
	buff.WriteString("\n")

	return buff.String()
}
