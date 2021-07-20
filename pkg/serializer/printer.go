package serializer

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"
	"github.com/tagpro/zd-search-cli/pkg/store"
	orgstore "github.com/tagpro/zd-search-cli/pkg/store/organistations"
	ticketstore "github.com/tagpro/zd-search-cli/pkg/store/tickets"
	userstore "github.com/tagpro/zd-search-cli/pkg/store/users"
	"github.com/tagpro/zd-search-cli/pkg/zerror"
)

// Common printing keys are defined here
const (
	ID             = "_id"
	URL            = "url"
	ExternalID     = "external_id"
	Name           = "name"
	DomainNames    = "domain_names"
	CreatedAt      = "created_at"
	Details        = "details"
	SharedTickets  = "shared_tickets"
	Tags           = "tags"
	Type           = "type"
	Subject        = "subject"
	Description    = "description"
	Priority       = "priority"
	Status         = "status"
	SubmitterID    = "submitter_id"
	AssigneeID     = "assignee_id"
	OrganizationID = "organization_id"
	HasIncidents   = "has_incidents"
	DueAt          = "due_at"
	Via            = "via"
	Alias          = "alias"
	Active         = "active"
	Verified       = "verified"
	Shared         = "shared"
	Locale         = "locale"
	Timezone       = "timezone"
	LastLoginAt    = "last_login_at"
	Email          = "email"
	Phone          = "phone"
	Signature      = "signature"
	Suspended      = "suspended"
	Role           = "role"
)

type kv struct {
	key   string
	value string
}

// pprint takes in a title and a list of key value pairs and pretty prints it as a table with 2 columns.
func pprint(title string, kvs ...kv) {
	red := color.New(color.FgRed)
	cyan := color.New(color.FgCyan)
	red.Fprintln(os.Stdout, title)
	for _, data := range kvs {
		cyan.Fprintf(os.Stdout, "%-20s | %s\n", data.key, data.value)
	}
}

func printOrganisations(s store.Store, organisations orgstore.Organisations) error {

	for _, org := range organisations {
		// Print Org info
		var printData []kv
		printData = append(printData,
			kv{ID, strconv.Itoa(org.ID)},
			kv{URL, org.URL},
			kv{ExternalID, org.ExternalID},
			kv{Name, org.Name},
			kv{DomainNames, strings.Join(org.DomainNames, ", ")},
			kv{CreatedAt, org.CreatedAt.Format(jsontime.ZDTimeFormat)},
			kv{Details, org.Details},
			kv{SharedTickets, strconv.FormatBool(org.SharedTickets)},
			kv{Tags, strings.Join(org.Tags, ", ")},
		)
		pprint("Organisation", printData...)

		// Print Users
		printData = []kv{}
		users, err := s.GetUsers(userstore.OrganizationID, strconv.Itoa(org.ID))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for i, user := range users {
			printData = append(printData, kv{strconv.Itoa(i), user.Name})
		}

		pprint("Users for organisation: "+org.Name, printData...)

		// Print Tickets
		printData = []kv{}
		tickets, err := s.GetTickets(ticketstore.OrganizationID, strconv.Itoa(org.ID))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for i, ticket := range tickets {
			printData = append(printData, kv{strconv.Itoa(i), ticket.Subject})
		}

		pprint("Tickets for organisation: "+org.Name, printData...)
	}
	return nil
}

func printUsers(s store.Store, users userstore.Users) error {
	for _, user := range users {
		// Print User info
		var printData []kv
		printData = append(printData,
			kv{ID, strconv.Itoa(user.ID)},
			kv{URL, user.URL},
			kv{ExternalID, user.ExternalID},
			kv{Name, user.Name},
			kv{Alias, user.Alias},
			kv{CreatedAt, user.CreatedAt.Format(jsontime.ZDTimeFormat)},
			kv{Active, strconv.FormatBool(user.Active)},
			kv{Verified, strconv.FormatBool(user.Verified)},
			kv{Shared, strconv.FormatBool(user.Shared)},
			kv{Locale, user.Locale},
			kv{Timezone, user.Timezone},
			kv{LastLoginAt, user.LastLoginAt.Format(jsontime.ZDTimeFormat)},
			kv{Email, user.Email},
			kv{Phone, user.Phone},
			kv{Signature, user.Signature},
			kv{OrganizationID, strconv.Itoa(user.OrganizationID)},
			kv{Tags, strings.Join(user.Tags, ", ")},
			kv{Suspended, strconv.FormatBool(user.Suspended)},
			kv{Role, user.Role},
		)
		pprint("User", printData...)

		// Print Organisation
		printData = []kv{}
		orgs, err := s.GetOrganisations(orgstore.ID, strconv.Itoa(user.OrganizationID))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for _, org := range orgs {
			printData = append(printData, kv{strconv.Itoa(org.ID), org.Name})
		}

		pprint("Organisation for user: "+user.Name, printData...)

		// Print submitted tickets
		printData = []kv{}
		tickets, err := s.GetTickets(ticketstore.SubmitterID, strconv.Itoa(user.ID))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for i, ticket := range tickets {
			printData = append(printData, kv{strconv.Itoa(i), ticket.Subject})
		}
		pprint("Submitted tickets from User: "+user.Name, printData...)

		// Print assigned tickets
		printData = []kv{}
		tickets, err = s.GetTickets(ticketstore.AssigneeID, strconv.Itoa(user.ID))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for i, ticket := range tickets {
			printData = append(printData, kv{strconv.Itoa(i), ticket.Subject})
		}

		pprint("Assigned tickets to User: "+user.Name, printData...)
	}
	return nil
}

func printTickets(s store.Store, tickets ticketstore.Tickets) error {
	for _, ticket := range tickets {
		// Print Org info
		var printData []kv
		printData = append(printData,

			kv{ID, ticket.ID},
			kv{URL, ticket.URL},
			kv{ExternalID, ticket.ExternalID},
			kv{CreatedAt, ticket.CreatedAt.Format(jsontime.ZDTimeFormat)},
			kv{Type, ticket.Type},
			kv{Subject, ticket.Subject},
			kv{Description, ticket.Description},
			kv{Priority, ticket.Priority},
			kv{Status, ticket.Status},
			kv{SubmitterID, strconv.Itoa(ticket.SubmitterID)},
			kv{AssigneeID, strconv.Itoa(ticket.AssigneeID)},
			kv{OrganizationID, strconv.Itoa(ticket.OrganizationID)},
			kv{Tags, strings.Join(ticket.Tags, ", ")},
			kv{HasIncidents, strconv.FormatBool(ticket.HasIncidents)},
			kv{DueAt, ticket.DueAt.Format(jsontime.ZDTimeFormat)},
			kv{Via, ticket.Via},
		)
		pprint("Ticket", printData...)

		// Print Users
		printData = []kv{}
		users, err := s.GetUsers(userstore.ID, strconv.Itoa(ticket.SubmitterID))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for _, user := range users {
			printData = append(printData, kv{"Submitter", user.Name})
		}
		users, err = s.GetUsers(userstore.ID, strconv.Itoa(ticket.AssigneeID))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for _, user := range users {
			printData = append(printData, kv{"Assignee", user.Name})
		}

		pprint("Users for ticket: "+ticket.Subject, printData...)

		// Print Organisation
		printData = []kv{}
		orgs, err := s.GetOrganisations(orgstore.ID, strconv.Itoa(ticket.OrganizationID))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for _, org := range orgs {
			printData = append(printData, kv{strconv.Itoa(org.ID), org.Name})
		}

		pprint("Organisation for ticket: "+ticket.Subject, printData...)
	}
	return nil
}
