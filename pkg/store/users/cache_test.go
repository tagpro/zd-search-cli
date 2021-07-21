package users

import (
	"errors"
	"testing"

	tassert "github.com/stretchr/testify/assert"
	"github.com/tagpro/zd-search-cli/pkg/jsontime"
	"github.com/tagpro/zd-search-cli/pkg/zerror"
)

func TestGetUsers(t *testing.T) {
	t.Run("Fails to return with invalid key name", TestGetUsers_invalidKey)
	t.Run("Fails to return with invalid value", TestGetUsers_notFound)
	t.Run("Successfully returns users list", TestGetUsers_Found)
}

func TestGetUsers_invalidKey(t *testing.T) {
	assert := tassert.New(t)
	c := &cache{}
	users, err := c.GetUsers("foo", "bar")
	assert.Nil(users)
	assert.EqualError(err, "invalid field name foo")
}

func TestGetUsers_notFound(t *testing.T) {
	assert := tassert.New(t)
	c := &cache{data: map[string]map[string]Users{"foo": {}}}
	users, err := c.GetUsers("foo", "invalid")
	assert.Nil(users)
	assert.EqualError(err, "no results found")
	assert.True(errors.Is(err, zerror.ErrNotFound))
}

func TestGetUsers_Found(t *testing.T) {
	assert := tassert.New(t)
	var want Users
	c := &cache{data: map[string]map[string]Users{
		"foo": {
			"bar": want,
		},
	}}
	got, err := c.GetUsers("foo", "bar")
	assert.Equal(want, got)
	assert.NoError(err)
}

func TestAddUser(t *testing.T) {
	t.Run("Fails to load with invalid cache", TestAddUser_invalidCache)
	t.Run("Successfully loads the user", TestAddUser_successful)
}

func TestAddUser_invalidCache(t *testing.T) {
	assert := tassert.New(t)
	c := &cache{}
	err := c.addUser(&User{})
	assert.EqualError(err, "cache data not initialised")
}

func TestAddUser_successful(t *testing.T) {
	assert := tassert.New(t)
	user := &User{
		ID:             1,
		URL:            "url",
		ExternalID:     "external id",
		Name:           "John Doe",
		Alias:          "Foo bar",
		CreatedAt:      jsontime.Time{},
		Active:         true,
		Verified:       false,
		Shared:         true,
		Locale:         "en-AU",
		Timezone:       "Sri Lanka",
		LastLoginAt:    jsontime.Time{},
		Email:          "foo@bar.com",
		Phone:          "1234-567-789",
		Signature:      "Don't Worry Be Happy!",
		OrganizationID: 100,
		Tags:           []string{"Sutton", "Forrest"},
		Suspended:      true,
		Role:           "admin",
	}
	c := &cache{
		users: Users{user},
		data:  map[string]map[string]Users{},
	}
	want := map[string]map[string]Users{
		ID:             {"1": Users{user}},
		URL:            {"url": Users{user}},
		ExternalID:     {"external id": Users{user}},
		Name:           {"John Doe": Users{user}},
		Alias:          {"Foo bar": Users{user}},
		CreatedAt:      {"0001-01-01T00:00:00 +00:00": Users{user}},
		Active:         {"true": Users{user}},
		Verified:       {"false": Users{user}},
		Shared:         {"true": Users{user}},
		Locale:         {"en-AU": Users{user}},
		Timezone:       {"Sri Lanka": Users{user}},
		LastLoginAt:    {"0001-01-01T00:00:00 +00:00": Users{user}},
		Email:          {"foo@bar.com": Users{user}},
		Phone:          {"1234-567-789": Users{user}},
		Signature:      {"Don't Worry Be Happy!": Users{user}},
		OrganizationID: {"100": Users{user}},
		Tags:           {"Sutton": Users{user}, "Forrest": Users{user}},
		Suspended:      {"true": Users{user}},
		Role:           {"admin": Users{user}},
	}
	err := c.addUser(user)
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
			name: "fails with no users in cache",
			c:    cache{data: map[string]map[string]Users{}},
			err:  "cache not initialised",
		},
		{
			name: "fails with no data in cache",
			c:    cache{users: Users{}},
			err:  "cache not initialised",
		},
		{
			name: "successfully loads data in the cache",
			c:    cache{users: Users{{}}, data: map[string]map[string]Users{}},
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
		"alias",
		"created_at",
		"active",
		"verified",
		"shared",
		"locale",
		"timezone",
		"last_login_at",
		"email",
		"phone",
		"signature",
		"organization_id",
		"tags",
		"suspended",
		"role",
	}
	tassert.Equal(t, want, GetKeys())
}
