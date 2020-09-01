package testlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestCreateNewApp(t *testing.T) {
	app := NewApp(&cli.Command{})
	assert.NotNil(t, app)
	assert.Equal(t, app.Name, name)
}

func TestRunAppVersion(t *testing.T) {
	app := NewApp(&cli.Command{})
	assert.NoError(t, app.Run([]string{"-v"}))
}
