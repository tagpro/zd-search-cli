package search

import (
	"fmt"

	"github.com/tagpro/zd-search-cli/pkg/store"
)

// Cli is definition of a CLI application for search
type Cli interface {
	Run() error
}

// app is a basic implementation of the search app which fulfils Cli interface
type app struct {
	store store.Store
}

func (a *app) Run() error {
	fmt.Println("Starting the app...")

	// TODO: Load and parse the data

	err := a.store.LoadData()
	if err != nil {
		return err
	}

	// TODO: Create prompts to read user input
	// TODO: Create logic to handle different inputs
	return nil
}

func NewSearchApp(usersFile, ticketsFile, organisationFile string) Cli {
	return &app{
		store: store.NewStore(store.Files{
			UsersFile:         usersFile,
			TicketsFile:       ticketsFile,
			OrganisationsFile: organisationFile,
		}),
	}
}

func Help() {
	fmt.Println(`TODO: Help
Is
Coming`)
}
