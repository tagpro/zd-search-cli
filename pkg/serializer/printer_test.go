package serializer

import (
	"io/ioutil"
	"os"
	"testing"

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
url                  | 
external_id          | 
name                 | Denzesk
domain_names         | 
created_at           | 0001-01-01T00:00:00 +00:00
details              | 
shared_tickets       | false
tags                 | 
Users for organisation: Denzesk
0                    | user name
Tickets for organisation: Denzesk
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

	orgs := orgstore.Organisations{{ID: 1, Name: "Denzesk"}}
	assert.NoError(printOrganisations(mockStore, orgs))
	require.NoError(t, fakeStdout.Close())
	printed, err := ioutil.ReadAll(r)
	assert.NoError(err)
	assert.Equal(want, string(printed))
}

func TestPrintUser(t *testing.T) {

	want := `User
_id                  | 1
url                  | 
external_id          | 
name                 | user name
alias                | 
created_at           | 0001-01-01T00:00:00 +00:00
active               | false
verified             | false
shared               | false
locale               | 
timezone             | 
last_login_at        | 0001-01-01T00:00:00 +00:00
email                | 
phone                | 
signature            | 
organization_id      | 101
tags                 | 
suspended            | false
role                 | 
Organisation for user: user name
0                    | Org 101
Submitted tickets from User: user name
0                    | subject 1
Assigned tickets to User: user name
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

	users := userstore.Users{{ID: 1, OrganizationID: 101, Name: "user name"}}
	assert.NoError(printUsers(mockStore, users))
	require.NoError(t, fakeStdout.Close())
	printed, err := ioutil.ReadAll(r)
	assert.NoError(err)
	assert.Equal(want, string(printed))
}

func TestPrintTickets(t *testing.T) {

	want := `Ticket
_id                  | 1
url                  | 
external_id          | 
created_at           | 0001-01-01T00:00:00 +00:00
type                 | 
subject              | subject
description          | 
priority             | 
status               | 
submitter_id         | 2
assignee_id          | 3
organization_id      | 4
tags                 | 
has_incidents        | false
due_at               | 0001-01-01T00:00:00 +00:00
via                  | 
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

	tickets := ticketstore.Tickets{{ID: "1", Subject: "subject", SubmitterID: 2, AssigneeID: 3, OrganizationID: 4}}
	assert.NoError(printTickets(mockStore, tickets))
	require.NoError(t, fakeStdout.Close())
	printed, err := ioutil.ReadAll(r)
	assert.NoError(err)
	assert.Equal(want, string(printed))
}
