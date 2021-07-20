package tickets

import (
	"testing"
	"time"

	tassert "github.com/stretchr/testify/assert"
	"github.com/tagpro/zd-search-cli/pkg/jsontime"
)

func TestLoadTickets(t *testing.T) {
	t.Run("Fails for dir path", TestLoadTicket_emptyPath)
	t.Run("Fails for bad path", TestLoadTicket_badPath)
	t.Run("Fails for invalid JSON file", TestLoadTicket_invalidJSONFile)
	t.Run("Fails to unmarshal JSON data to struct", TestLoadTicket_badJSONValues)
	t.Run("Successfully loads JSON file", TestLoadTicket_validFile)
}

func TestLoadTicket_emptyPath(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	cache, err := LoadTickets("./testdata")
	assert.Nil(cache)
	assert.EqualError(err, "error reading JSON file: read ./testdata: is a directory")
}

func TestLoadTicket_badPath(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	cache, err := LoadTickets("testdata/nofile.json")
	assert.Nil(cache)
	assert.EqualError(err, "error reading JSON file: open testdata/nofile.json: no such file or directory")
}

func TestLoadTicket_invalidJSONFile(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	cache, err := LoadTickets("testdata/invalid.json")
	assert.Nil(cache)
	assert.EqualError(err, "error parsing JSON file: unexpected end of JSON input")
}

func TestLoadTicket_badJSONValues(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)
	cache, err := LoadTickets("testdata/bad.json")
	assert.Nil(cache)
	assert.EqualError(err, "error parsing JSON file: json: cannot unmarshal number into Go struct field Ticket._id of type string")
}

func TestLoadTicket_validFile(t *testing.T) {
	t.Parallel()
	assert := tassert.New(t)

	createdAt, err := time.Parse(jsontime.ZDTimeFormat, "2016-04-28T11:19:34 -10:00")
	assert.NoError(err)
	dueAt, err := time.Parse(jsontime.ZDTimeFormat, "2016-07-31T02:37:50 -10:00")
	assert.NoError(err)

	var valid = &cache{
		tickets: Tickets{{
			Id:             "id",
			Url:            "url",
			ExternalId:     "external id",
			CreatedAt:      jsontime.Time{Time: createdAt},
			Type:           "incident",
			Subject:        "subject",
			Description:    "just a description?",
			Priority:       "medium-high",
			Status:         "pending",
			SubmitterId:    1,
			AssigneeId:     2,
			OrganizationId: 100,
			Tags:           []string{"kokomo", "bahamas"},
			HasIncidents:   true,
			DueAt:          jsontime.Time{Time: dueAt},
			Via:            "web",
		}},
		data: map[string]map[string]Tickets{},
	}
	cache, err := LoadTickets("testdata/valid.json")
	assert.Equal(valid, cache)
	assert.NoError(err)
}
