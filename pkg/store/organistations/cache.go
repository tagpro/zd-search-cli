package organistations

//go:generate mockgen -source=cache.go -destination=testdata/mocks/cache.go -package=orgcache . Cache

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/tagpro/zd-search-cli/pkg/zerror"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"
)

const (
	ID            = "_id"
	URL           = "url"
	ExternalID    = "external_id"
	Name          = "name"
	DomainNames   = "domain_names"
	CreatedAt     = "created_at"
	Details       = "details"
	SharedTickets = "shared_tickets"
	Tags          = "tags"
)

var keys = []string{
	ID,
	URL,
	ExternalID,
	Name,
	DomainNames,
	CreatedAt,
	Details,
	SharedTickets,
	Tags,
}

// GetKeys returns the list of all the keys on which organisation is cached on
func GetKeys() []string {
	return keys
}

// Cache is used to optimise all the organisations by creating an index and fetch a list of Organisations using
// keys from the list of available keys from JSON and the value for the key
type Cache interface {
	GetOrganisations(key, input string) (Organisations, error)
	Optimise() error
}

type cache struct {
	organisations Organisations
	data          map[string]map[string]Organisations
}

// GetOrganisations takes in key, the search term in the question, and search value to return the list of organisations
func (c *cache) GetOrganisations(key, input string) (Organisations, error) {
	if _, ok := c.data[key]; !ok {
		return nil, fmt.Errorf("invalid field name %s", key)
	}
	if orgs, ok := c.data[key][input]; !ok {
		return nil, fmt.Errorf("%w", zerror.ErrNotFound)
	} else {
		return orgs, nil
	}
}

// Optimise creates an index from a (key, value) to list of organisations
func (c *cache) Optimise() error {
	if c.organisations == nil || c.data == nil {
		return errors.New("cache not initialised")
	}
	for _, o := range c.organisations {
		if err := c.addOrganisation(o); err != nil {
			return fmt.Errorf("couldn't load all organisations: %w", err)
		}
	}
	return nil
}

// addOrganisation adds an organisation into the cache (the key, value map)
func (c *cache) addOrganisation(org *Organisation) error {
	if c.data == nil {
		return fmt.Errorf("cache data not initialised")
	}
	// Insert _id
	if _, ok := c.data[ID]; !ok {
		c.data[ID] = map[string]Organisations{}
	}
	c.data[ID][strconv.Itoa(org.ID)] = append(c.data[ID][strconv.Itoa(org.ID)], org)

	// Insert url
	if _, ok := c.data[URL]; !ok {
		c.data[URL] = map[string]Organisations{}
	}
	c.data[URL][org.URL] = append(c.data[URL][org.URL], org)

	// Insert external ID
	if _, ok := c.data[ExternalID]; !ok {
		c.data[ExternalID] = map[string]Organisations{}
	}
	c.data[ExternalID][org.ExternalID] = append(c.data[ExternalID][org.ExternalID], org)

	// Insert name
	if _, ok := c.data[Name]; !ok {
		c.data[Name] = map[string]Organisations{}
	}
	c.data[Name][org.Name] = append(c.data[Name][org.Name], org)

	// Insert domain names
	if _, ok := c.data[DomainNames]; !ok {
		c.data[DomainNames] = map[string]Organisations{}
	}
	for _, domain := range org.DomainNames {
		c.data[DomainNames][domain] = append(c.data[DomainNames][domain], org)
	}

	// Insert created_at
	if _, ok := c.data[CreatedAt]; !ok {
		c.data[CreatedAt] = map[string]Organisations{}
	}
	c.data[CreatedAt][org.CreatedAt.Format(jsontime.ZDTimeFormat)] = append(c.data[CreatedAt][org.CreatedAt.Format(jsontime.ZDTimeFormat)], org)

	// Insert details
	if _, ok := c.data[Details]; !ok {
		c.data[Details] = map[string]Organisations{}
	}
	c.data[Details][org.Details] = append(c.data[Details][org.Details], org)

	// Insert shared tickets
	if _, ok := c.data[SharedTickets]; !ok {
		c.data[SharedTickets] = map[string]Organisations{}
	}
	c.data[SharedTickets][strconv.FormatBool(org.SharedTickets)] = append(c.data[SharedTickets][strconv.FormatBool(org.SharedTickets)], org)

	// Insert tags
	if _, ok := c.data[Tags]; !ok {
		c.data[Tags] = map[string]Organisations{}
	}
	for _, tag := range org.Tags {
		c.data[Tags][tag] = append(c.data[Tags][tag], org)
	}
	return nil
}

func newCache(organisations Organisations) *cache {
	return &cache{organisations: organisations, data: map[string]map[string]Organisations{}}
}
