package gitignore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	assert := assert.New(t)
	tt := Templates{
		{Name: "Foo"},
		{Name: "Test"},
		{Name: "Sort"},
		{Name: "Bar"},
	}

	tt.Sort()

	assert.Equal("Bar", tt[0].Name)
	assert.Equal("Foo", tt[1].Name)
	assert.Equal("Sort", tt[2].Name)
	assert.Equal("Test", tt[3].Name)
}
