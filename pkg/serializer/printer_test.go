package serializer

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"

	orgstore "github.com/tagpro/zd-search-cli/pkg/store/organistations"
	ticketstore "github.com/tagpro/zd-search-cli/pkg/store/tickets"
	userstore "github.com/tagpro/zd-search-cli/pkg/store/users"

	"github.com/golang/mock/gomock"
	tassert "github.com/stretchr/testify/assert"
	store "github.com/tagpro/zd-search-cli/pkg/store/testdata/mocks"

	"github.com/stretchr/testify/require"
)

func TestPrintOrganisations(t *testing.T) {

	want := `Organisation
_id                  | 1
url                  | url
external_id          | external id
name                 | Mega Corp
domain_names         | corp.com, bar.com
created_at           | 2016-05-21T11:10:28 -10:00
details              | MegaCorp
shared_tickets       | false
tags                 | Foo, Bar
Users for organisation: Mega Corp
0                    | user name
Tickets for organisation: Mega Corp
0                    | subject
`
	assert := tassert.New(t)
	realStdout := os.Stdout
	defer func() { os.Stdout = realStdout }()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout

	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)
	mockStore.EXPECT().GetUsers("organization_id", "1").Return(userstore.Users{&userstore.User{Name: "user name"}}, nil)
	mockStore.EXPECT().GetTickets("organization_id", "1").Return(ticketstore.Tickets{{Subject: "subject"}}, nil)

	createdAt, err := time.Parse(jsontime.ZDTimeFormat, "2016-05-21T11:10:28 -10:00")
	assert.NoError(err)
	orgs := orgstore.Organisations{
		{
			ID:            1,
			URL:           "url",
			ExternalID:    "external id",
			Name:          "Mega Corp",
			DomainNames:   []string{"corp.com", "bar.com"},
			CreatedAt:     jsontime.Time{Time: createdAt},
			Details:       "MegaCorp",
			SharedTickets: false,
			Tags:          []string{"Foo", "Bar"},
		},
	}
	assert.NoError(printOrganisations(mockStore, orgs))
	require.NoError(t, fakeStdout.Close())
	printed, err := ioutil.ReadAll(r)
	assert.NoError(err)
	assert.Equal(want, string(printed))
}

func TestPrintUser(t *testing.T) {

	want := `User
_id                  | 1
url                  | url
external_id          | external id
name                 | John Doe
alias                | Foo bar
created_at           | 2016-04-15T05:19:46 -10:00
active               | true
verified             | false
shared               | true
locale               | en-AU
timezone             | Sri Lanka
last_login_at        | 2013-08-04T01:03:27 -10:00
email                | foo@bar.com
phone                | 1234-567-789
signature            | Don't Worry Be Happy!
organization_id      | 101
tags                 | Sutton, Forrest
suspended            | true
role                 | admin
Organisation for user: John Doe
0                    | Org 101
Submitted tickets from User: John Doe
0                    | subject 1
Assigned tickets to User: John Doe
0                    | subject 2
`
	assert := tassert.New(t)
	realStdout := os.Stdout
	defer func() { os.Stdout = realStdout }()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout

	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)
	mockStore.EXPECT().GetOrganisations("_id", "101").Return(orgstore.Organisations{&orgstore.Organisation{Name: "Org 101"}}, nil)
	mockStore.EXPECT().GetTickets("submitter_id", "1").Return(ticketstore.Tickets{{Subject: "subject 1"}}, nil)
	mockStore.EXPECT().GetTickets("assignee_id", "1").Return(ticketstore.Tickets{{Subject: "subject 2"}}, nil)

	createdAt, err := time.Parse(jsontime.ZDTimeFormat, "2016-04-15T05:19:46 -10:00")
	assert.NoError(err)
	lastLoginAt, err := time.Parse(jsontime.ZDTimeFormat, "2013-08-04T01:03:27 -10:00")
	assert.NoError(err)

	users := userstore.Users{
		{
			ID:             1,
			URL:            "url",
			ExternalID:     "external id",
			Name:           "John Doe",
			Alias:          "Foo bar",
			CreatedAt:      jsontime.Time{Time: createdAt},
			Active:         true,
			Verified:       false,
			Shared:         true,
			Locale:         "en-AU",
			Timezone:       "Sri Lanka",
			LastLoginAt:    jsontime.Time{Time: lastLoginAt},
			Email:          "foo@bar.com",
			Phone:          "1234-567-789",
			Signature:      "Don't Worry Be Happy!",
			OrganizationID: 101,
			Tags:           []string{"Sutton", "Forrest"},
			Suspended:      true,
			Role:           "admin",
		},
	}
	assert.NoError(printUsers(mockStore, users))
	require.NoError(t, fakeStdout.Close())
	printed, err := ioutil.ReadAll(r)
	assert.NoError(err)
	assert.Equal(want, string(printed))
}

func TestPrintTickets(t *testing.T) {

	want := `Ticket
_id                  | 1
url                  | url
external_id          | external id
created_at           | 2016-04-28T11:19:34 -10:00
type                 | incident
subject              | subject
description          | just a description?
priority             | medium-high
status               | pending
submitter_id         | 2
assignee_id          | 3
organization_id      | 4
tags                 | kokomo, bahamas
has_incidents        | true
due_at               | 2016-07-31T02:37:50 -10:00
via                  | web
Users for ticket: subject
Submitter            | user name 1
Assignee             | user name 2
Organisation for ticket: subject
0                    | Org 101
`
	assert := tassert.New(t)
	realStdout := os.Stdout
	defer func() { os.Stdout = realStdout }()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout

	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)
	mockStore.EXPECT().GetOrganisations("_id", "4").Return(orgstore.Organisations{&orgstore.Organisation{Name: "Org 101"}}, nil)
	mockStore.EXPECT().GetUsers("_id", "2").Return(userstore.Users{&userstore.User{Name: "user name 1"}}, nil)
	mockStore.EXPECT().GetUsers("_id", "3").Return(userstore.Users{&userstore.User{Name: "user name 2"}}, nil)

	createdAt, err := time.Parse(jsontime.ZDTimeFormat, "2016-04-28T11:19:34 -10:00")
	assert.NoError(err)
	dueAt, err := time.Parse(jsontime.ZDTimeFormat, "2016-07-31T02:37:50 -10:00")
	assert.NoError(err)

	tickets := ticketstore.Tickets{
		{
			ID:             "1",
			URL:            "url",
			ExternalID:     "external id",
			CreatedAt:      jsontime.Time{Time: createdAt},
			Type:           "incident",
			Subject:        "subject",
			Description:    "just a description?",
			Priority:       "medium-high",
			Status:         "pending",
			SubmitterID:    2,
			AssigneeID:     3,
			OrganizationID: 4,
			Tags:           []string{"kokomo", "bahamas"},
			HasIncidents:   true,
			DueAt:          jsontime.Time{Time: dueAt},
			Via:            "web",
		},
	}
	assert.NoError(printTickets(mockStore, tickets))
	require.NoError(t, fakeStdout.Close())
	printed, err := ioutil.ReadAll(r)
	assert.NoError(err)
	assert.Equal(want, string(printed))
}
