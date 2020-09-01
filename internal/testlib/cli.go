package testlib

import "github.com/urfave/cli/v2"

const name = "testapp"

// NewApp creates a new test app
func NewApp(command *cli.Command) *cli.App {
	app := &cli.App{
		Name:    name,
		Usage:   "Test App",
		Version: "1.0.0",
		OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
			return nil
		},
		Commands: []*cli.Command{
			command,
		},
	}

	return app
}

// Run app
func Run(app *cli.App, args []string) error {
	aa := append([]string{name}, args...)
	return app.Run(aa)
}
