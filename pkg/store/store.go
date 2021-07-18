package store

import "github.com/tagpro/zd-search-cli/pkg/store/organistations"

type Files struct {
	UsersFile         string
	TicketsFile       string
	OrganisationsFile string
}

type Store struct {
	files Files
}

func (s *Store) LoadData() error {
	if err := organistations.LoadOrganisations(s.files.UsersFile); err != nil {
		return err
	}
	return nil
}

func NewStore(f Files) Store {
	return Store{f}
}
