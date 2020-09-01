package logging

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/fatih/color"
)

type (
	// Padding for logs
	Padding int
	// Action is a function that returns error
	Action func() error
)

const (
	// DefaultInitialPadding is the default padding in the log library.
	DefaultInitialPadding = 3
	// ExtraPadding is the double of the DefaultInitialPadding.
	ExtraPadding = DefaultInitialPadding * 2
)

// Log pretty prints the given action and its title.
func Log(title string, next Action) error {
	defer func() {
		cli.Default.Padding = int(DefaultInitialPadding)
	}()

	cli.Default.Padding = int(ExtraPadding)

	log.Infof(color.New(color.Bold).Sprint(title))
	cli.Default.Padding = int(ExtraPadding + DefaultInitialPadding)

	return next()
}
