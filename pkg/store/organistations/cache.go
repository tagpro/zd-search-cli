package organistations

import (
	"fmt"
	"strconv"

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
	GetOrganisations(key, input string) ([]Organisation, error)
	AddOrganisation(org Organisation) error
}

type cache struct {
	data map[string]map[string][]*Organisation
}

func (c *cache) GetOrganisations(key, input string) ([]*Organisation, error) {
	if _, ok := c.data[key]; !ok {
		return nil, fmt.Errorf("invalid key name %s", key)
	}
	if org, ok := c.data[key][input]; !ok {
		return nil, fmt.Errorf("No results found")
	} else {
		return org, nil
	}
}

func (c *cache) AddOrganisation(org Organisation) error {
	// Insert _id
	if _, ok := c.data[Id]; !ok {
		c.data[Id] = map[string][]*Organisation{}
	}
	c.data[Id][strconv.Itoa(org.Id)] = append(c.data[Id][strconv.Itoa(org.Id)], &org)

	// Insert url
	if _, ok := c.data[Url]; !ok {
		c.data[Url] = map[string][]*Organisation{}
	}
	c.data[Url][org.Url] = append(c.data[Url][org.Url], &org)

	// Insert external Id
	if _, ok := c.data[ExternalId]; !ok {
		c.data[ExternalId] = map[string][]*Organisation{}
	}
	c.data[ExternalId][org.ExternalId] = append(c.data[ExternalId][org.ExternalId], &org)

	// Insert name
	if _, ok := c.data[Name]; !ok {
		c.data[Name] = map[string][]*Organisation{}
	}
	c.data[Name][org.Name] = append(c.data[Name][org.Name], &org)

	// Insert domain names
	if _, ok := c.data[DomainNames]; !ok {
		c.data[DomainNames] = map[string][]*Organisation{}
	}
	for _, domain := range org.DomainNames {
		c.data[DomainNames][domain] = append(c.data[DomainNames][domain], &org)
	}

	// Insert created_at
	if _, ok := c.data[CreatedAt]; !ok {
		c.data[CreatedAt] = map[string][]*Organisation{}
	}
	c.data[CreatedAt][org.CreatedAt.Format(jsontime.ZDTimeFormat)] = append(c.data[CreatedAt][org.CreatedAt.Format(jsontime.ZDTimeFormat)], &org)

	// Insert details
	if _, ok := c.data[Details]; !ok {
		c.data[Details] = map[string][]*Organisation{}
	}
	c.data[Details][org.Details] = append(c.data[Details][org.Details], &org)

	// Insert shared tickets
	if _, ok := c.data[SharedTickets]; !ok {
		c.data[SharedTickets] = map[string][]*Organisation{}
	}
	c.data[SharedTickets][strconv.FormatBool(org.SharedTickets)] = append(c.data[SharedTickets][strconv.FormatBool(org.SharedTickets)], &org)

	// Insert tags
	if _, ok := c.data[Tags]; !ok {
		c.data[Tags] = map[string][]*Organisation{}
	}
	for _, tag := range org.Tags {
		c.data[Tags][tag] = append(c.data[Tags][tag], &org)
	}
	return nil
}
