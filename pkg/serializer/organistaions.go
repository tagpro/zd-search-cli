package serializer

import (
	"errors"
	"strconv"
	"strings"

	"github.com/tagpro/zd-search-cli/pkg/zerror"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"
	orgstore "github.com/tagpro/zd-search-cli/pkg/store/organistations"
	ticketstore "github.com/tagpro/zd-search-cli/pkg/store/tickets"
	userstore "github.com/tagpro/zd-search-cli/pkg/store/users"
)

func (s *serializer) printOrganisations(organisations orgstore.Organisations) error {

	for _, org := range organisations {
		// Print Org info
		var printData []kv
		printData = append(printData,
			kv{Id, strconv.Itoa(org.Id)},
			kv{Url, org.Url},
			kv{ExternalId, org.ExternalId},
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
		users, err := s.store.GetUsers(userstore.OrganizationId, strconv.Itoa(org.Id))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for i, user := range users {
			printData = append(printData, kv{strconv.Itoa(i), user.Name})
		}

		pprint("Users for organisation: "+org.Name, printData...)

		// Print Tickets
		printData = []kv{}
		tickets, err := s.store.GetTickets(ticketstore.OrganizationId, strconv.Itoa(org.Id))
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
