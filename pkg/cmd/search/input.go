package search

import (
	"errors"
	"fmt"

	"github.com/tagpro/zd-search-cli/pkg/zerror"

	"github.com/tagpro/zd-search-cli/pkg/serializer"

	"github.com/manifoldco/promptui"
)

const (
	searchZendesk = "Search Zendesk"
	showFields    = "Show list of searchable fields"
	quit          = "Quit"
)

var primaryOptions = []string{
	searchZendesk,
	showFields,
	quit,
}

// serve defines and runs a single flow of prompts to serve the users request
func (a *app) serve() error {
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
		//TODO: Get all the fields in serializer and print it
		fmt.Println("Showing all the fields")
	case quit:
		return fmt.Errorf("%w", zerror.ErrQuit)
	default:
		fmt.Println("Invalid")
	}
	return nil
}

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
