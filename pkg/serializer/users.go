package serializer

import (
	"errors"
	"strconv"
	"strings"

	"github.com/tagpro/zd-search-cli/pkg/zerror"

	ticketstore "github.com/tagpro/zd-search-cli/pkg/store/tickets"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"
	orgstore "github.com/tagpro/zd-search-cli/pkg/store/organistations"
	userstore "github.com/tagpro/zd-search-cli/pkg/store/users"
)

func (s *serializer) printUsers(users userstore.Users) error {
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
		orgs, err := s.store.GetOrganisations(orgstore.ID, strconv.Itoa(user.OrganizationID))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for _, org := range orgs {
			printData = append(printData, kv{strconv.Itoa(org.ID), org.Name})
		}

		pprint("Organisation for user: "+user.Name, printData...)

		// Print submitted tickets
		printData = []kv{}
		tickets, err := s.store.GetTickets(ticketstore.SubmitterID, strconv.Itoa(user.ID))
		if err != nil && !errors.Is(err, zerror.ErrNotFound) {
			return err
		}
		for i, ticket := range tickets {
			printData = append(printData, kv{strconv.Itoa(i), ticket.Subject})
		}
		pprint("Submitted tickets from User: "+user.Name, printData...)

		// Print assigned tickets
		printData = []kv{}
		tickets, err = s.store.GetTickets(ticketstore.AssigneeID, strconv.Itoa(user.ID))
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
