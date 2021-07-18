package organistations

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"
	"github.com/tagpro/zd-search-cli/pkg/store/tickets"
	"github.com/tagpro/zd-search-cli/pkg/store/users"
)

type Organisation struct {
	Id            int             `json:"_id"`
	Url           string          `json:"url"`
	ExternalId    string          `json:"external_id"`
	Name          string          `json:"name"`
	DomainNames   []string        `json:"domain_names"`
	CreatedAt     jsontime.Time   `json:"created_at"`
	Details       string          `json:"details"`
	SharedTickets bool            `json:"shared_tickets"`
	Tags          []string        `json:"tags"`
	Users         users.Users     `json:"-"`
	Tickets       tickets.Tickets `json:"-"`
}

func (o *Organisation) AddUser(user *users.User) {
	o.Users = append(o.Users, user)
}

func (o *Organisation) AddTicket(ticket *tickets.Ticket) {
	o.Tickets = append(o.Tickets, ticket)
}

type Organisations []*Organisation

func LoadOrganisations(path string) (Cache, error) {
	var organisations Organisations
	if path == "" {
		return nil, errors.New("no file path available")
	}

	jsonFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON file: %w", err)
	}
	err = json.Unmarshal(jsonFile, &organisations)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON file: %w", err)
	}
	return newCache(organisations), nil
}