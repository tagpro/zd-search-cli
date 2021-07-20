package organistations

import (
	"testing"
	"time"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"

	tassert "github.com/stretchr/testify/assert"
)

var valid = &cache{
	organisations: Organisations{{
		Id:            101,
		Url:           "url",
		ExternalId:    "external id",
		Name:          "Mega Corp",
		DomainNames:   []string{"corp.com", "bar.com"},
		CreatedAt:     jsontime.Time{},
		Details:       "MegaCorp",
		SharedTickets: false,
		Tags:          []string{"Foo", "Bar"},
	}},
	data: map[string]map[string]Organisations{},
}

func TestLoadOrganisations(t *testing.T) {
	t.Run("Fails for dir path", TestLoadOrganisation_emptyPath)
	t.Run("Fails for bad path", TestLoadOrganisation_badPath)
	t.Run("Fails for invalid JSON file", TestLoadOrganisation_invalidJSONFile)
	t.Run("Fails to unmarshal JSON data to struct", TestLoadOrganisation_badJSONValues)
	t.Run("Successfully loads JSON file", TestLoadOrganisation_validFile)
}

func TestLoadOrganisation_emptyPath(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	cache, err := LoadOrganisations("./testdata")
	assert.Nil(cache)
	assert.EqualError(err, "error reading JSON file: read ./testdata: is a directory")
}

func TestLoadOrganisation_badPath(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	cache, err := LoadOrganisations("testdata/nofile.json")
	assert.Nil(cache)
	assert.EqualError(err, "error reading JSON file: open testdata/nofile.json: no such file or directory")
}

func TestLoadOrganisation_invalidJSONFile(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	cache, err := LoadOrganisations("testdata/invalid.json")
	assert.Nil(cache)
	assert.EqualError(err, "error parsing JSON file: unexpected end of JSON input")
}

func TestLoadOrganisation_badJSONValues(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	cache, err := LoadOrganisations("testdata/bad.json")
	assert.Nil(cache)
	assert.EqualError(err, "error parsing JSON file: json: cannot unmarshal string into Go struct field Organisation._id of type int")
}

func TestLoadOrganisation_validFile(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	expectedTime, err := time.Parse(jsontime.ZDTimeFormat, "2016-05-21T11:10:28 -10:00")
	assert.NoError(err)
	valid.organisations[0].CreatedAt.Time = expectedTime
	cache, err := LoadOrganisations("testdata/valid.json")
	assert.Equal(valid, cache)
	assert.NoError(err)
}
