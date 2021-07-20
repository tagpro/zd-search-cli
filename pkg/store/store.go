package store

//go:generate mockgen -source=store.go -destination=./testdata/mocks/store.go -package=store . Store

import (
	"context"

	orgstore "github.com/tagpro/zd-search-cli/pkg/store/organistations"
	ticketstore "github.com/tagpro/zd-search-cli/pkg/store/tickets"
	userstore "github.com/tagpro/zd-search-cli/pkg/store/users"
	"golang.org/x/sync/errgroup"
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

// init optimises the data by create a cache(a reverse index of the values) for all the data.
func (s *store) init() error {
	if err := s.organisations.Optimise(); err != nil {
		return err
	}
	if err := s.users.Optimise(); err != nil {
		return err
	}
	return s.tickets.Optimise()
}

func NewStore(f Files) (Store, error) {
	g, _ := errgroup.WithContext(context.TODO())
	var orgs orgstore.Cache
	var users userstore.Cache
	var tickets ticketstore.Cache
	g.Go(func() error {
		var err error
		orgs, err = orgstore.LoadOrganisations(f.OrganisationsFile)
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		users, err = userstore.LoadUsers(f.UsersFile)
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		tickets, err = ticketstore.LoadTickets(f.TicketsFile)
		if err != nil {
			return err
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}
	s := &store{organisations: orgs, users: users, tickets: tickets}
	if err := s.init(); err != nil {
		return nil, err
	}
	return s, nil
}
