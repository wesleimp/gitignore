package generate

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleimp/gitignore/internal/testlib"
)

func TestFileDoesNotExists(t *testing.T) {
	assert.NoError(t, os.Remove("testdata/.gitignore"))
	defer createFile(t)

	app := testlib.NewApp(Command)
	args := []string{"generate", "-l", "", "-p", "testdata"}
	assert.NoError(t, testlib.Run(app, args))
}

func TestFileExistsWithoutForce(t *testing.T) {
	app := testlib.NewApp(Command)
	args := []string{"generate", "-l", "", "-p", "testdata"}
	assert.EqualError(t, testlib.Run(app, args), ErrFileExists.Error())
}

func createFile(t *testing.T) {
	_, err := os.Create("testdata/.gitignore")
	assert.NoError(t, err)
}
