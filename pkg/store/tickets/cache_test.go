package tickets

import (
	"errors"
	"testing"

	tassert "github.com/stretchr/testify/assert"
	"github.com/tagpro/zd-search-cli/pkg/jsontime"
	"github.com/tagpro/zd-search-cli/pkg/zerror"
)

func TestGetTickets(t *testing.T) {
	t.Run("Fails to return with invalid key name", TestGetTickets_invalidKey)
	t.Run("Fails to return with invalid value", TestGetTickets_notFound)
	t.Run("Successfully returns ticketanistaions list", TestGetTickets_Found)
}

func TestGetTickets_invalidKey(t *testing.T) {
	assert := tassert.New(t)
	c := &cache{}
	tickets, err := c.GetTickets("foo", "bar")
	assert.Nil(tickets)
	assert.EqualError(err, "invalid field name foo")
}

func TestGetTickets_notFound(t *testing.T) {
	assert := tassert.New(t)
	c := &cache{data: map[string]map[string]Tickets{"foo": {}}}
	tickets, err := c.GetTickets("foo", "invalid")
	assert.Nil(tickets)
	assert.EqualError(err, "no results found")
	assert.True(errors.Is(err, zerror.ErrNotFound))
}

func TestGetTickets_Found(t *testing.T) {
	assert := tassert.New(t)
	var want Tickets
	c := &cache{data: map[string]map[string]Tickets{
		"foo": {
			"bar": want,
		},
	}}
	got, err := c.GetTickets("foo", "bar")
	assert.Equal(want, got)
	assert.NoError(err)
}

func TestAddTicket(t *testing.T) {
	t.Run("Fails to load with invalid cache", TestAddTicket_invalidCache)
	t.Run("Successfully loads the ticket", TestAddTicket_successful)
}

func TestAddTicket_invalidCache(t *testing.T) {
	assert := tassert.New(t)
	c := &cache{}
	err := c.addTicket(&Ticket{})
	assert.EqualError(err, "cache data not initialised")
}

func TestAddTicket_successful(t *testing.T) {
	assert := tassert.New(t)
	ticket := &Ticket{
		Id:             "id",
		Url:            "url",
		ExternalId:     "external id",
		CreatedAt:      jsontime.Time{},
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
		DueAt:          jsontime.Time{},
		Via:            "web",
	}
	c := &cache{
		tickets: Tickets{ticket},
		data:    map[string]map[string]Tickets{},
	}
	want := map[string]map[string]Tickets{
		Id:             {"id": Tickets{ticket}},
		Url:            {"url": Tickets{ticket}},
		ExternalId:     {"external id": Tickets{ticket}},
		CreatedAt:      {"0001-01-01T00:00:00 +00:00": Tickets{ticket}},
		Type:           {"incident": Tickets{ticket}},
		Subject:        {"subject": Tickets{ticket}},
		Description:    {"just a description?": Tickets{ticket}},
		Priority:       {"medium-high": Tickets{ticket}},
		Status:         {"pending": Tickets{ticket}},
		SubmitterId:    {"1": Tickets{ticket}},
		AssigneeId:     {"2": Tickets{ticket}},
		OrganizationId: {"100": Tickets{ticket}},
		Tags:           {"kokomo": Tickets{ticket}, "bahamas": Tickets{ticket}},
		HasIncidents:   {"true": Tickets{ticket}},
		DueAt:          {"0001-01-01T00:00:00 +00:00": Tickets{ticket}},
		Via:            {"web": Tickets{ticket}},
	}
	err := c.addTicket(ticket)
	assert.NoError(err)
	assert.EqualValues(want, c.data)
}

func TestOptimise(t *testing.T) {
	cases := []struct {
		name string
		c    cache
		err  string
	}{
		{
			name: "fails with no tickets in cache",
			c:    cache{data: map[string]map[string]Tickets{}},
			err:  "cache not initialised",
		},
		{
			name: "fails with no data in cache",
			c:    cache{tickets: Tickets{}},
			err:  "cache not initialised",
		},
		{
			name: "successfully loads data in the cache",
			c:    cache{tickets: Tickets{{}}, data: map[string]map[string]Tickets{}},
			err:  "",
		},
	}
	for _, tcase := range cases {
		t.Run(tcase.name, func(t *testing.T) {
			assert := tassert.New(t)
			err := tcase.c.Optimise()
			if tcase.err != "" {
				assert.EqualError(err, tcase.err)
			} else {
				assert.NoError(err)
			}
		})
	}
}

func TestGetKeys(t *testing.T) {
	want := []string{
		"_id",
		"url",
		"external_id",
		"created_at",
		"type",
		"subject",
		"description",
		"priority",
		"status",
		"submitter_id",
		"assignee_id",
		"organization_id",
		"tags",
		"has_incidents",
		"due_at",
		"via",
	}
	tassert.Equal(t, want, GetKeys())
}
