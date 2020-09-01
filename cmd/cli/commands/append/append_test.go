package append

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleimp/gitignore/internal/testlib"
)

func prepareTest(t *testing.T) {
	err := os.Remove("./testdata/.gitignore")
	assert.NoError(t, err)

	f, err := os.Create("./testdata/.gitignore")
	assert.NoError(t, err)
	defer f.Close()

	_, err = f.WriteString("Foo\nBar")
	assert.NoError(t, err)
}

func TestAppendText(t *testing.T) {
	assert := assert.New(t)
	prepareTest(t)

	app := testlib.NewApp(Command)
	assert.NotNil(app)
	assert.NoError(testlib.Run(app, []string{"append", "-p", "./testdata", "-t", "Test.exe"}))

	bts, err := ioutil.ReadFile("./testdata/.gitignore")
	assert.NoError(err)
	assert.Contains(string(bts), "Test.exe")
}
