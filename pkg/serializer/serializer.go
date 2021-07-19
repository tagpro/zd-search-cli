package serializer

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
		err = s.printOrganisations(orgs)
		if err != nil {
			return err
		}
	case Users:
		users, err := s.store.GetUsers(criteria.Field, criteria.Value)
		if err != nil {
			return err
		}
		err = s.printUsers(users)
		if err != nil {
			return err
		}
	case Tickets:
		tickets, err := s.store.GetTickets(criteria.Field, criteria.Value)
		if err != nil {
			return err
		}
		err = s.printTickets(tickets)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewSerializer(s store.Store) Serializer {
	return &serializer{store: s}
}
