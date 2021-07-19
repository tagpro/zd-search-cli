package search

import (
	"errors"
	"fmt"

	"github.com/tagpro/zd-search-cli/pkg/zerror"

	"github.com/tagpro/zd-search-cli/pkg/serializer"

	"github.com/tagpro/zd-search-cli/pkg/store"
)

// Cli is definition of a CLI application for search
type Cli interface {
	Run() error
}

// app is a basic implementation of the search app which fulfils Cli interface
type app struct {
	serializer serializer.Serializer
}

func (a *app) Run() error {
	fmt.Println("Starting the app . . .")
	for {
		err := a.serve()
		if err != nil && errors.Is(err, zerror.ErrQuit) {
			fmt.Println("Quitting . . .")
			return nil
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}

func NewSearchApp(usersFile, ticketsFile, organisationFile string) (Cli, error) {
	s, err := store.NewStore(store.Files{
		UsersFile:         usersFile,
		TicketsFile:       ticketsFile,
		OrganisationsFile: organisationFile,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create store: %w", err)
	}
	return &app{serializer.NewSerializer(s)}, nil
}

func Help() {
	fmt.Println(`TODO: Help
Is
Coming`)
}
