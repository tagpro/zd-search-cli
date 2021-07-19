package store

import (
	orgstore "github.com/tagpro/zd-search-cli/pkg/store/organistations"
	ticketstore "github.com/tagpro/zd-search-cli/pkg/store/tickets"
	userstore "github.com/tagpro/zd-search-cli/pkg/store/users"
)

type Files struct {
	UsersFile         string
	TicketsFile       string
	OrganisationsFile string
}

type Keys struct {
	Organisation []string
	User         []string
	Ticket       []string
}

type Store interface {
	GetOrganisations(key, input string) (orgstore.Organisations, error)
	GetUsers(key, input string) (userstore.Users, error)
	GetTickets(key, input string) (ticketstore.Tickets, error)
	GetKeys() Keys
}

type store struct {
	organisations orgstore.Cache
	users         userstore.Cache
	tickets       ticketstore.Cache
}

func (s *store) GetOrganisations(key, input string) (orgstore.Organisations, error) {
	return s.organisations.GetOrganisations(key, input)
}

func (s *store) GetUsers(key, input string) (userstore.Users, error) {
	return s.users.GetUsers(key, input)
}

func (s *store) GetTickets(key, input string) (ticketstore.Tickets, error) {
	return s.tickets.GetTickets(key, input)
}

func (s *store) GetKeys() Keys {
	return Keys{
		Organisation: orgstore.GetKeys(),
		User:         userstore.GetKeys(),
		Ticket:       ticketstore.GetKeys(),
	}
}

// init optimises the data by create a cache for all the data.
// It will optimise organisations, then link users to organisations and then link tickets to users and organisations
func (s *store) init() error {
	if err := s.organisations.Optimise(); err != nil {
		return err
	}
	if err := s.users.Optimise(); err != nil {
		return err
	}
	if err := s.tickets.Optimise(); err != nil {
		return err
	}
	return nil
}

func NewStore(f Files) (Store, error) {
	// TODO: Load files in go routines (errgroup)
	orgs, err := orgstore.LoadOrganisations(f.OrganisationsFile)
	if err != nil {
		return nil, err
	}
	users, err := userstore.LoadUsers(f.UsersFile)
	if err != nil {
		return nil, err
	}
	tickets, err := ticketstore.LoadTickets(f.TicketsFile)
	if err != nil {
		return nil, err
	}
	s := &store{organisations: orgs, users: users, tickets: tickets}
	if err := s.init(); err != nil {
		return nil, err
	}
	return s, nil
}
