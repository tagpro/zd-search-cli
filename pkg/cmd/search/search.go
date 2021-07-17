package search

import "fmt"

// Cli is definition of a CLI application for search
type Cli interface {
	Run() error
}

// app is a basic implementation of the search app which fulfils Cli interface
type app struct {
	usersFile         string
	ticketsFile       string
	organisationsFile string
}

func (a *app) Run() error {
	fmt.Println("Starting the app...")
	// TODO: Load and parse the data
	// TODO: Create prompts to read user input
	// TODO: Create logic to handle different inputs
	return nil
}

func NewSearchApp(usersFile, ticketsFile, organisationFile string) Cli {
	return &app{
		usersFile:         usersFile,
		ticketsFile:       ticketsFile,
		organisationsFile: organisationFile,
	}
}

func Help() {
	fmt.Println(`TODO: Help
Is
Coming`)
}
