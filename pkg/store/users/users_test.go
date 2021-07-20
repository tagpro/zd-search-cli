package users

import (
	"testing"
	"time"

	tassert "github.com/stretchr/testify/assert"
	"github.com/tagpro/zd-search-cli/pkg/jsontime"
)

type testcase struct {
	name string
	test func(t *testing.T)
}

func TestLoadUsers(t *testing.T) {
	t.Run("Fails for dir path", TestLoadUser_emptyPath)
	t.Run("Fails for bad path", TestLoadUser_badPath)
	t.Run("Fails for invalid JSON file", TestLoadUser_invalidJSONFile)
	t.Run("Fails to unmarshal JSON data to struct", TestLoadUser_badJSONValues)
	t.Run("Successfully loads JSON file", TestLoadUser_validFile)
}

func TestLoadUser_emptyPath(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	cache, err := LoadUsers("./testdata")
	assert.Nil(cache)
	assert.EqualError(err, "error reading JSON file: read ./testdata: is a directory")
}

func TestLoadUser_badPath(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	cache, err := LoadUsers("testdata/nofile.json")
	assert.Nil(cache)
	assert.EqualError(err, "error reading JSON file: open testdata/nofile.json: no such file or directory")
}

func TestLoadUser_invalidJSONFile(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	cache, err := LoadUsers("testdata/invalid.json")
	assert.Nil(cache)
	assert.EqualError(err, "error parsing JSON file: unexpected end of JSON input")
}

func TestLoadUser_badJSONValues(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	cache, err := LoadUsers("testdata/bad.json")
	assert.Nil(cache)
	assert.EqualError(err, "error parsing JSON file: json: cannot unmarshal string into Go struct field User._id of type int")
}

func TestLoadUser_validFile(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)

	createdAt, err := time.Parse(jsontime.ZDTimeFormat, "2016-04-15T05:19:46 -10:00")
	assert.NoError(err)
	lastLoginAt, err := time.Parse(jsontime.ZDTimeFormat, "2013-08-04T01:03:27 -10:00")
	assert.NoError(err)

	var want = &cache{
		users: Users{{
			Id:             1,
			Url:            "url",
			ExternalId:     "external id",
			Name:           "John Doe",
			Alias:          "Foo bar",
			CreatedAt:      jsontime.Time{Time: createdAt},
			Active:         true,
			Verified:       false,
			Shared:         true,
			Locale:         "en-AU",
			Timezone:       "Sri Lanka",
			LastLoginAt:    jsontime.Time{Time: lastLoginAt},
			Email:          "foo@bar.com",
			Phone:          "1234-567-789",
			Signature:      "Don't Worry Be Happy!",
			OrganizationId: 100,
			Tags:           []string{"Sutton", "Forrest"},
			Suspended:      true,
			Role:           "admin",
		}},
		data: map[string]map[string]Users{},
	}

	got, err := LoadUsers("testdata/valid.json")
	assert.Equal(want, got)
	assert.NoError(err)
}
