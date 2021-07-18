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
	Id            int              `json:"_id"`
	Url           string           `json:"url"`
	ExternalId    string           `json:"external_id"`
	Name          string           `json:"name"`
	DomainNames   []string         `json:"domain_names"`
	CreatedAt     jsontime.Time    `json:"created_at"`
	Details       string           `json:"details"`
	SharedTickets bool             `json:"shared_tickets"`
	Tags          []string         `json:"tags"`
	users         []users.User     `json:"-"`
	tickets       []tickets.Ticket `json:"-"`
}

func (o *Organisation) AddUser(user users.User) {
	o.users = append(o.users, user)
}

func (o *Organisation) AddTicket(ticket tickets.Ticket) {
	o.tickets = append(o.tickets, ticket)
}

type Organisations struct {
	orgs []*Organisation
}

var organisations Organisations

func GetOrganisations() Organisations {
	return organisations
}

func LoadOrganisations(path string) error {
	if path == "" {
		return errors.New("no file path available")
	}

	jsonFile, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %s", err)
	}
	err = json.Unmarshal(jsonFile, &organisations)
	if err != nil {
		return fmt.Errorf("error parsing JSON file: %s", err)
	}
	return nil
}
