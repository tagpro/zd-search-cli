package serializer

//go:generate mockgen -source=serializer.go -destination=./testdata/mocks/serializer.go -package=serializer . Serializer

import (
	"errors"

	"github.com/tagpro/zd-search-cli/pkg/store"
)

type Entity string

const (
	Users         Entity = "Users"
	Tickets       Entity = "Tickets"
	Organisations Entity = "Organisations"
)

func GetEntities() []string {
	return []string{
		string(Users),
		string(Tickets),
		string(Organisations),
	}
}

func ToEntity(e string) (Entity, error) {
	if e != string(Users) && e != string(Tickets) && e != string(Organisations) {
		return "", errors.New("invalid entity")
	}
	return Entity(e), nil
}

type SearchCriteria struct {
	Entity Entity
	Field  string
	Value  string
}

type Serializer interface {
	SearchEntity(criteria SearchCriteria) error
	PrintKeys()
}

type serializer struct {
	store store.Store
}

func (s *serializer) SearchEntity(criteria SearchCriteria) error {
	switch criteria.Entity {
	case Organisations:
		orgs, err := s.store.GetOrganisations(criteria.Field, criteria.Value)
		if err != nil {
			return err
		}
		err = printOrganisations(s.store, orgs)
		if err != nil {
			return err
		}
	case Users:
		users, err := s.store.GetUsers(criteria.Field, criteria.Value)
		if err != nil {
			return err
		}
		err = printUsers(s.store, users)
		if err != nil {
			return err
		}
	case Tickets:
		tickets, err := s.store.GetTickets(criteria.Field, criteria.Value)
		if err != nil {
			return err
		}
		err = printTickets(s.store, tickets)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *serializer) PrintKeys() {
	keys := s.store.GetKeys()
	printList("List of acceptable keys to search on:")
	printList(string(Organisations), keys.Organisation...)
	printList(string(Tickets), keys.Ticket...)
	printList(string(Users), keys.User...)
}

func NewSerializer(s store.Store) Serializer {
	return &serializer{store: s}
}
