package main

import (
	"log"

	_ "github.com/golang/mock/gomock"
	flag "github.com/spf13/pflag"
	"github.com/tagpro/zd-search-cli/pkg/cmd/search"
)

var (
	userFile   string
	ticketFile string
	orgFile    string
	help       bool
)

func init() {
	flag.StringVarP(&userFile, "users", "u", "./data/users.json", "Path of the users file to import")
	flag.StringVarP(&ticketFile, "tickets", "t", "./data/tickets.json", "Path of the tickets file to import")
	flag.StringVarP(&orgFile, "organisations", "o", "./data/organizations.json", "Path of the organizations file to import")
	flag.BoolVarP(&help, "help", "h", false, "Print helpful information")
}

func main() {
	flag.Parse()
	if help {
		search.Help()
		return
	}
	app, err := search.NewSearchApp(userFile, ticketFile, orgFile)
	if err != nil {
		log.Fatalf("Couldn't start Application: %v", err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf("Got an unexpected error: %v", err)
	}
}
