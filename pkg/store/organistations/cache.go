package organistations

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/tagpro/zd-search-cli/pkg/zerror"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"
)

const (
	Id            = "Id"
	Url           = "Url"
	ExternalId    = "ExternalId"
	Name          = "Name"
	DomainNames   = "DomainNames"
	CreatedAt     = "CreatedAt"
	Details       = "Details"
	SharedTickets = "SharedTickets"
	Tags          = "Tags"
)

var keys = []string{
	Id,
	Url,
	ExternalId,
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

type Cache interface {
	GetOrganisations(key, input string) (Organisations, error)
	Optimise() error
}

type cache struct {
	organisations Organisations
	data          map[string]map[string]Organisations
}

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

func (c *cache) addOrganisation(org *Organisation) error {
	// Insert _id
	if _, ok := c.data[Id]; !ok {
		c.data[Id] = map[string]Organisations{}
	}
	c.data[Id][strconv.Itoa(org.Id)] = append(c.data[Id][strconv.Itoa(org.Id)], org)

	// Insert url
	if _, ok := c.data[Url]; !ok {
		c.data[Url] = map[string]Organisations{}
	}
	c.data[Url][org.Url] = append(c.data[Url][org.Url], org)

	// Insert external Id
	if _, ok := c.data[ExternalId]; !ok {
		c.data[ExternalId] = map[string]Organisations{}
	}
	c.data[ExternalId][org.ExternalId] = append(c.data[ExternalId][org.ExternalId], org)

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

func (c *cache) Optimise() error {
	if c.organisations == nil {
		return errors.New("organisations not loaded")
	}
	for _, o := range c.organisations {
		if err := c.addOrganisation(o); err != nil {
			return fmt.Errorf("couldn't load all organisations: %w", err)
		}
	}
	return nil
}

func newCache(organisations Organisations) *cache {
	return &cache{organisations: organisations, data: map[string]map[string]Organisations{}}
}
