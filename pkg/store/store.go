package store

import (
	"github.com/tagpro/zd-search-cli/pkg/store/organistations"
)

type Files struct {
	UsersFile         string
	TicketsFile       string
	OrganisationsFile string
}

type Store interface {
	GetOrganisations(key, input string) ([]*organistations.Organisation, error)
}

type store struct {
	organisations organistations.Cache
}

func (s *store) GetOrganisations(key, input string) ([]*organistations.Organisation, error) {
	return s.organisations.GetOrganisations(key, input)
}

func (s *store) init() error {
	if err := s.organisations.Optimise(); err != nil {
		return err
	}
	return nil
}

func NewStore(f Files) (Store, error) {
	orgs, err := organistations.LoadOrganisations(f.OrganisationsFile)
	if err != nil {
		return nil, err
	}
	s := &store{organisations: orgs}
	if err := s.init(); err != nil {
		return nil, err
	}
	return s, nil
}
