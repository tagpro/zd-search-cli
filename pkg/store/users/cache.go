package users

//go:generate mockgen -source=cache.go -destination=testdata/mocks/cache.go -package=mocks . Cache

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"

	"github.com/tagpro/zd-search-cli/pkg/zerror"
)

const (
	ID             = "_id"
	URL            = "url"
	ExternalID     = "external_id"
	Name           = "name"
	Alias          = "alias"
	CreatedAt      = "created_at"
	Active         = "active"
	Verified       = "verified"
	Shared         = "shared"
	Locale         = "locale"
	Timezone       = "timezone"
	LastLoginAt    = "last_login_at"
	Email          = "email"
	Phone          = "phone"
	Signature      = "signature"
	OrganizationID = "organization_id"
	Tags           = "tags"
	Suspended      = "suspended"
	Role           = "role"
)

func GetKeys() []string {
	return []string{
		ID,
		URL,
		ExternalID,
		Name,
		Alias,
		CreatedAt,
		Active,
		Verified,
		Shared,
		Locale,
		Timezone,
		LastLoginAt,
		Email,
		Phone,
		Signature,
		OrganizationID,
		Tags,
		Suspended,
		Role,
	}
}

type Cache interface {
	GetUsers(key, input string) (Users, error)
	Optimise() error
}

type cache struct {
	users Users
	data  map[string]map[string]Users
}

func (c *cache) GetUsers(key, input string) (Users, error) {
	if _, ok := c.data[key]; !ok {
		return nil, fmt.Errorf("invalid field name %s", key)
	}
	if users, ok := c.data[key][input]; !ok {
		return nil, fmt.Errorf("%w", zerror.ErrNotFound)
	} else {
		return users, nil
	}
}

func (c *cache) Optimise() error {
	if c.users == nil || c.data == nil {
		return errors.New("cache not initialised")
	}
	for _, user := range c.users {
		if err := c.addUser(user); err != nil {
			return fmt.Errorf("failed to load all users: %w", err)
		}
	}
	return nil
}

func (c *cache) addUser(user *User) error {
	if c.data == nil {
		return fmt.Errorf("cache data not initialised")
	}
	// Insert _id
	if _, ok := c.data[ID]; !ok {
		c.data[ID] = map[string]Users{}
	}
	c.data[ID][strconv.Itoa(user.ID)] = append(c.data[ID][strconv.Itoa(user.ID)], user)

	// Insert url
	if _, ok := c.data[URL]; !ok {
		c.data[URL] = map[string]Users{}
	}
	c.data[URL][user.URL] = append(c.data[URL][user.URL], user)

	// Insert external_id
	if _, ok := c.data[ExternalID]; !ok {
		c.data[ExternalID] = map[string]Users{}
	}
	c.data[ExternalID][user.ExternalID] = append(c.data[ExternalID][user.ExternalID], user)

	// Insert name
	if _, ok := c.data[Name]; !ok {
		c.data[Name] = map[string]Users{}
	}
	c.data[Name][user.Name] = append(c.data[Name][user.Name], user)

	// Insert alias

	if _, ok := c.data[Alias]; !ok {
		c.data[Alias] = map[string]Users{}
	}
	c.data[Alias][user.Alias] = append(c.data[Alias][user.Alias], user)

	// Insert created_at
	if _, ok := c.data[CreatedAt]; !ok {
		c.data[CreatedAt] = map[string]Users{}
	}
	c.data[CreatedAt][user.CreatedAt.Format(jsontime.ZDTimeFormat)] = append(c.data[CreatedAt][user.CreatedAt.Format(jsontime.ZDTimeFormat)], user)

	// Insert active

	if _, ok := c.data[Active]; !ok {
		c.data[Active] = map[string]Users{}
	}
	c.data[Active][strconv.FormatBool(user.Active)] = append(c.data[Active][strconv.FormatBool(user.Active)], user)

	// Insert verified
	if _, ok := c.data[Verified]; !ok {
		c.data[Verified] = map[string]Users{}
	}
	c.data[Verified][strconv.FormatBool(user.Verified)] = append(c.data[Verified][strconv.FormatBool(user.Verified)], user)

	// Insert shared
	if _, ok := c.data[Shared]; !ok {
		c.data[Shared] = map[string]Users{}
	}
	c.data[Shared][strconv.FormatBool(user.Shared)] = append(c.data[Shared][strconv.FormatBool(user.Shared)], user)

	// Insert locale
	if _, ok := c.data[Locale]; !ok {
		c.data[Locale] = map[string]Users{}
	}
	c.data[Locale][user.Locale] = append(c.data[Locale][user.Locale], user)

	// Insert timezone
	if _, ok := c.data[Timezone]; !ok {
		c.data[Timezone] = map[string]Users{}
	}
	c.data[Timezone][user.Timezone] = append(c.data[Timezone][user.Timezone], user)

	// Insert last_login_at
	if _, ok := c.data[LastLoginAt]; !ok {
		c.data[LastLoginAt] = map[string]Users{}
	}
	c.data[LastLoginAt][user.LastLoginAt.Format(jsontime.ZDTimeFormat)] = append(c.data[LastLoginAt][user.LastLoginAt.Format(jsontime.ZDTimeFormat)], user)

	// Insert email
	if _, ok := c.data[Email]; !ok {
		c.data[Email] = map[string]Users{}
	}
	c.data[Email][user.Email] = append(c.data[Email][user.Email], user)

	// Insert phone
	if _, ok := c.data[Phone]; !ok {
		c.data[Phone] = map[string]Users{}
	}
	c.data[Phone][user.Phone] = append(c.data[Phone][user.Phone], user)
	// Insert signature
	if _, ok := c.data[Signature]; !ok {
		c.data[Signature] = map[string]Users{}
	}
	c.data[Signature][user.Signature] = append(c.data[Signature][user.Signature], user)
	// Insert organization_id
	if _, ok := c.data[OrganizationID]; !ok {
		c.data[OrganizationID] = map[string]Users{}
	}
	c.data[OrganizationID][strconv.Itoa(user.OrganizationID)] = append(c.data[OrganizationID][strconv.Itoa(user.OrganizationID)], user)
	// Insert tags
	if _, ok := c.data[Tags]; !ok {
		c.data[Tags] = map[string]Users{}
	}
	for _, tag := range user.Tags {
		c.data[Tags][tag] = append(c.data[Tags][tag], user)
	}
	// Insert suspended
	if _, ok := c.data[Suspended]; !ok {
		c.data[Suspended] = map[string]Users{}
	}
	c.data[Suspended][strconv.FormatBool(user.Suspended)] = append(c.data[Suspended][strconv.FormatBool(user.Suspended)], user)

	// Insert role
	if _, ok := c.data[Role]; !ok {
		c.data[Role] = map[string]Users{}
	}
	c.data[Role][user.Role] = append(c.data[Role][user.Role], user)
	return nil
}

func newCache(users Users) *cache {
	return &cache{users: users, data: map[string]map[string]Users{}}
}
