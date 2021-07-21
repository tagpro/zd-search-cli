package search

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/tagpro/zd-search-cli/pkg/serializer"
	"github.com/tagpro/zd-search-cli/pkg/store"
	"github.com/tagpro/zd-search-cli/pkg/zerror"
)

const (
	searchZendesk = "Search Zendesk"
	showFields    = "Show list of searchable fields"
	quit          = "Quit"
)

// Cli is definition of a CLI application for search
type Cli interface {
	Run() error
}

// app is a basic implementation of the search app which fulfils Cli interface
type app struct {
	serializer serializer.Serializer
}

// Run creates the interface to take user input and process it. It wraps an internal method that handles the input logic
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

// serve defines and runs a single flow of prompts to serve the users request. It is responsible to create a
// user flow for a single operation
func (a *app) serve() error {
	primaryOptions := []string{
		searchZendesk,
		showFields,
		quit,
	}

	prompt := promptui.Select{
		Label: "Please choose an option",
		Items: primaryOptions,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}
	return a.handlePrimaryResponse(result)
}

// handlePrimaryResponse uses the output from users selection prompt to read the result and perform logic based on the
// chosen input
func (a *app) handlePrimaryResponse(result string) error {
	switch result {
	case searchZendesk:
		criteria, err := getSearchCriteria()
		if err != nil {
			return err
		}
		err = a.serializer.SearchEntity(*criteria)
		if errors.Is(err, zerror.ErrNotFound) {
			fmt.Println("No results found.")
			return nil
		}
		return err
	case showFields:
		a.serializer.PrintKeys()
	case quit:
		return fmt.Errorf("%w", zerror.ErrQuit)
	default:
		fmt.Println("Invalid input")
	}
	return nil
}

// getSearchCriteria will as user to choose the values to search for and create the search criteria required by the
// serializer as an input.
func getSearchCriteria() (criteria *serializer.SearchCriteria, err error) {
	entityPrompt := promptui.Select{
		Label: "Please select a category",
		Items: serializer.GetEntities(),
	}
	_, entityResult, err := entityPrompt.Run()
	if err != nil {
		return nil, err
	}
	entity, err := serializer.ToEntity(entityResult)
	if err != nil {
		return nil, err
	}
	fieldPrompt := promptui.Prompt{
		Label: "Enter the field name",
	}
	field, err := fieldPrompt.Run()
	if err != nil {
		return nil, err
	}

	valuePrompt := promptui.Prompt{
		Label: "Please enter the value to search on",
	}
	value, err := valuePrompt.Run()
	if err != nil {
		return nil, err
	}
	return &serializer.SearchCriteria{Entity: entity, Field: field, Value: value}, nil
}

// NewSearchApp initialises the store and returns a Cli application
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
