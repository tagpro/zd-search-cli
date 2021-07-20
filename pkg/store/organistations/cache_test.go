package organistations

import (
	"errors"
	"testing"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"

	"github.com/tagpro/zd-search-cli/pkg/zerror"

	tassert "github.com/stretchr/testify/assert"
)

func TestGetOrganisations(t *testing.T) {
	t.Run("Fails to return with invalid key name", TestGetOrganisations_invalidKey)
	t.Run("Fails to return with invalid value", TestGetOrganisations_notFound)
	t.Run("Successfully returns organisations list", TestGetOrganisations_Found)
}

func TestGetOrganisations_invalidKey(t *testing.T) {
	assert := tassert.New(t)
	c := &cache{}
	orgs, err := c.GetOrganisations("foo", "bar")
	assert.Nil(orgs)
	assert.EqualError(err, "invalid field name foo")
}

func TestGetOrganisations_notFound(t *testing.T) {
	assert := tassert.New(t)
	c := &cache{data: map[string]map[string]Organisations{"foo": {}}}
	orgs, err := c.GetOrganisations("foo", "invalid")
	assert.Nil(orgs)
	assert.EqualError(err, "no results found")
	assert.True(errors.Is(err, zerror.ErrNotFound))
}

func TestGetOrganisations_Found(t *testing.T) {
	assert := tassert.New(t)
	var want Organisations
	c := &cache{data: map[string]map[string]Organisations{
		"foo": {
			"bar": want,
		},
	}}
	got, err := c.GetOrganisations("foo", "bar")
	assert.Equal(want, got)
	assert.NoError(err)
}

func TestAddOrganisation(t *testing.T) {
	t.Run("Fails to load with invalid cache", TestAddOrganisation_invalidCache)
	t.Run("Successfully loads the org", TestAddOrganisation_successful)
}

func TestAddOrganisation_invalidCache(t *testing.T) {
	assert := tassert.New(t)
	c := &cache{}
	err := c.addOrganisation(&Organisation{})
	assert.EqualError(err, "cache data not initialised")
}

func TestAddOrganisation_successful(t *testing.T) {
	assert := tassert.New(t)
	org := &Organisation{
		ID:            101,
		URL:           "url",
		ExternalID:    "external id",
		Name:          "Mega Corp",
		DomainNames:   []string{"corp.com", "bar.com"},
		CreatedAt:     jsontime.Time{},
		Details:       "MegaCorp",
		SharedTickets: false,
		Tags:          []string{"Foo", "Bar"},
	}
	c := &cache{
		organisations: Organisations{org},
		data:          map[string]map[string]Organisations{},
	}
	want := map[string]map[string]Organisations{
		ID:            {"101": Organisations{org}},
		URL:           {"url": Organisations{org}},
		ExternalID:    {"external id": Organisations{org}},
		Name:          {"Mega Corp": Organisations{org}},
		DomainNames:   {"corp.com": Organisations{org}, "bar.com": Organisations{org}},
		CreatedAt:     {"0001-01-01T00:00:00 +00:00": Organisations{org}},
		Details:       {"MegaCorp": Organisations{org}},
		SharedTickets: {"false": Organisations{org}},
		Tags:          {"Foo": Organisations{org}, "Bar": Organisations{org}},
	}
	err := c.addOrganisation(org)
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
			name: "fails with no organisations in cache",
			c:    cache{data: map[string]map[string]Organisations{}},
			err:  "cache not initialised",
		},
		{
			name: "fails with no data in cache",
			c:    cache{organisations: Organisations{}},
			err:  "cache not initialised",
		},
		{
			name: "successfully loads data in the cache",
			c:    cache{organisations: Organisations{{}}, data: map[string]map[string]Organisations{}},
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
		"name",
		"domain_names",
		"created_at",
		"details",
		"shared_tickets",
		"tags",
	}
	tassert.Equal(t, want, GetKeys())
}
