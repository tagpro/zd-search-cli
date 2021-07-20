package serializer

import (
	"errors"
	"strconv"
	"strings"

	"github.com/tagpro/zd-search-cli/pkg/zerror"

	orgstore "github.com/tagpro/zd-search-cli/pkg/store/organistations"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"
	ticketstore "github.com/tagpro/zd-search-cli/pkg/store/tickets"
	userstore "github.com/tagpro/zd-search-cli/pkg/store/users"
)

func (s *serializer) printTickets(tickets ticketstore.Tickets) error {
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
		users, err := s.store.GetUsers(userstore.ID, strconv.Itoa(ticket.SubmitterID))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for _, user := range users {
			printData = append(printData, kv{"Submitter", user.Name})
		}
		users, err = s.store.GetUsers(userstore.ID, strconv.Itoa(ticket.AssigneeID))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for _, user := range users {
			printData = append(printData, kv{"Assignee", user.Name})
		}

		pprint("Users for ticket: "+ticket.Subject, printData...)

		// Print Organisation
		printData = []kv{}
		orgs, err := s.store.GetOrganisations(orgstore.ID, strconv.Itoa(ticket.OrganizationID))
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
