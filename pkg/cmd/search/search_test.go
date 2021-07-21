package search

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"
)

func TestNewSearchApp_failStore(t *testing.T) {
	assert := tassert.New(t)
	app, err := NewSearchApp("", "", "")
	assert.EqualError(err, "failed to create store: error reading JSON file: open : no such file or directory")
	assert.Nil(app)
}

func TestNewSearchApp(t *testing.T) {
	assert := tassert.New(t)
	app, err := NewSearchApp("./testdata/valid.json", "./testdata/valid.json", "./testdata/valid.json")
	assert.NoError(err)
	assert.NotNil(app)
}
