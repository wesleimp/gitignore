package generate

import (
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/wesleimp/gitignore/internal/gitignore"
	"github.com/wesleimp/gitignore/internal/logging"
)

// ErrFileExists return if file already exists
var ErrFileExists = errors.New(".gitignore file already exists. Use --force flag to override it")

// Command for generate
var Command = &cli.Command{
	Name:    "generate",
	Aliases: []string{"g"},
	Usage:   "Generate .gitignore file",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "lang",
			Aliases:  []string{"l"},
			Required: true,
			Usage:    "Templates to generate file",
		},
		&cli.StringFlag{
			Name:    "path",
			Aliases: []string{"p"},
			Usage:   "Path to .gitignore file",
			Value:   ".",
		},
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "Force create a new .gitignore file. This flag will override the existent file",
			Value:   false,
		},
	},
	Action: func(c *cli.Context) error {
		return logging.Log("generating file...", func() error {
			if err := checkFile(c); err != nil {
				return err
			}

			return generate(c)
		})
	},
}

func generate(c *cli.Context) error {
	dest := c.String("path")
	langs := c.StringSlice("lang")

	log.WithField("path", dest).Info("creating file")
	f, err := os.OpenFile(filepath.Join(dest, ".gitignore"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_EXCL, 0644)
	if err != nil {
		return errors.Wrap(err, "error openning .gitignore file")
	}
	defer f.Close()

	content, err := gitignore.DownloadTemplates(langs)
	if err != nil {
		return err
	}

	_, err = f.WriteString(content)
	return err
}

func checkFile(c *cli.Context) error {
	dest := c.String("path")
	force := c.Bool("force")

	log.Info("checking .gitignore file")
	_, err := os.Stat(filepath.Join(dest, ".gitignore"))
	if os.IsNotExist(err) {
		log.WithField("path", dest+".gitignore").Debug(".gitignore does not exists")
		return nil
	}

	if force {
		log.Warn("--force is set, removing current .gitignore")
		return os.Remove(filepath.Join(dest, ".gitignore"))
	}

	return ErrFileExists
}
