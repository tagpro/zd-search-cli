package store

import (
	"errors"
	"testing"

	ticketcache "github.com/tagpro/zd-search-cli/pkg/store/tickets/testdata/mocks"
	usercache "github.com/tagpro/zd-search-cli/pkg/store/users/testdata/mocks"

	orgcache "github.com/tagpro/zd-search-cli/pkg/store/organistations/testdata/mocks"

	"github.com/golang/mock/gomock"

	tassert "github.com/stretchr/testify/assert"

	"github.com/tagpro/zd-search-cli/pkg/store/organistations"
	"github.com/tagpro/zd-search-cli/pkg/store/tickets"
	"github.com/tagpro/zd-search-cli/pkg/store/users"
)

func TestNewStore(t *testing.T) {
	cases := []struct {
		name  string
		f     Files
		store Store
		err   string
	}{
		{
			name: "Fails to load with invalid org file",
			f: Files{
				UsersFile:         "./testdata/valid.json",
				TicketsFile:       "./testdata/valid.json",
				OrganisationsFile: "./testdata/invalid.json",
			},
			store: nil,
			err:   "error parsing JSON file: unexpected end of JSON input",
		},
		{
			name: "Fails to load with invalid org file",
			f: Files{
				UsersFile:         "./testdata/valid.json",
				TicketsFile:       "./testdata/valid.json",
				OrganisationsFile: "./testdata/invalid.json",
			},
			store: nil,
			err:   "error parsing JSON file: unexpected end of JSON input",
		},
		{
			name: "Fails to load with invalid user file",
			f: Files{
				UsersFile:         "./testdata/invalid.json",
				TicketsFile:       "./testdata/valid.json",
				OrganisationsFile: "./testdata/valid.json",
			},
			store: nil,
			err:   "error parsing JSON file: unexpected end of JSON input",
		},
		{
			name: "Fails to load with invalid tickets file",
			f: Files{
				UsersFile:         "./testdata/valid.json",
				TicketsFile:       "./testdata/invalid.json",
				OrganisationsFile: "./testdata/valid.json",
			},
			store: nil,
			err:   "error parsing JSON file: unexpected end of JSON input",
		},
		{
			name: "Successfully creates the store",
			f: Files{
				UsersFile:         "./testdata/valid.json",
				TicketsFile:       "./testdata/valid.json",
				OrganisationsFile: "./testdata/valid.json",
			},
			store: &store{},
			err:   "",
		},
	}
	for _, tcase := range cases {
		t.Run(tcase.name, func(t *testing.T) {
			assert := tassert.New(t)
			s, err := NewStore(tcase.f)
			if tcase.err != "" {
				assert.EqualError(err, tcase.err)
			}
			if tcase.store != nil {
				want, ok := s.(*store)
				assert.True(ok)
				if assert.NotNil(want) {
					assert.NotNil(want.tickets)
					assert.NotNil(want.users)
					assert.NotNil(want.organisations)
				}
			}
		})
	}
}

func TestStore_GetKeys(t *testing.T) {
	assert := tassert.New(t)
	store := &store{}
	keys := Keys{
		Organisation: []string{"_id", "url", "external_id", "name", "domain_names", "created_at", "details", "shared_tickets", "tags"},
		User:         []string{"_id", "url", "external_id", "name", "alias", "created_at", "active", "verified", "shared", "locale", "timezone", "last_login_at", "email", "phone", "signature", "organization_id", "tags", "suspended", "role"},
		Ticket:       []string{"_id", "url", "external_id", "created_at", "type", "subject", "description", "priority", "status", "submitter_id", "assignee_id", "organization_id", "tags", "has_incidents", "due_at", "via"},
	}
	assert.Equal(keys, store.GetKeys())
}

func TestStore_GetOrganisations(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	want := organistations.Organisations{}
	wantErr := errors.New("want error")
	mockOrg := orgcache.NewMockCache(ctrl)
	mockOrg.EXPECT().
		GetOrganisations("key", "value").
		Return(want, wantErr).
		Times(1)

	s := &store{
		organisations: mockOrg,
		users:         nil,
		tickets:       nil,
	}
	got, gotErr := s.GetOrganisations("key", "value")
	assert.Equal(want, got)
	assert.Equal(wantErr, gotErr)
}

func TestStore_GetTickets(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	want := tickets.Tickets{}
	wantErr := errors.New("want error")
	mockTicket := ticketcache.NewMockCache(ctrl)
	mockTicket.EXPECT().
		GetTickets("key", "value").
		Return(want, wantErr).
		Times(1)

	s := &store{
		organisations: nil,
		users:         nil,
		tickets:       mockTicket,
	}
	got, gotErr := s.GetTickets("key", "value")
	assert.Equal(want, got)
	assert.Equal(wantErr, gotErr)
}

func TestStore_GetUsers(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	want := users.Users{}
	wantErr := errors.New("want error")
	mockUser := usercache.NewMockCache(ctrl)
	mockUser.EXPECT().
		GetUsers("key", "value").
		Return(want, wantErr).
		Times(1)

	s := &store{
		organisations: nil,
		users:         mockUser,
		tickets:       nil,
	}
	got, gotErr := s.GetUsers("key", "value")
	assert.Equal(want, got)
	assert.Equal(wantErr, gotErr)
}

func TestStore_init(t *testing.T) {
	t.Run("Fails to optimise organisation data", TestStore_organisationError)
	t.Run("Fails to optimise user data", TestStore_userError)
	t.Run("Fails to optimise ticket data", TestStore_ticketError)
	t.Run("Successfully optimises all data", TestStore_initSuccessful)
}

func TestStore_organisationError(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockOrg := orgcache.NewMockCache(ctrl)
	mockTicket := ticketcache.NewMockCache(ctrl)
	mockUser := usercache.NewMockCache(ctrl)
	s := &store{mockOrg, mockUser, mockTicket}
	gomock.InOrder(
		mockOrg.EXPECT().Optimise().Return(errors.New("error")).Times(1),
		mockUser.EXPECT().Optimise().Times(0),
		mockTicket.EXPECT().Optimise().Times(0),
	)
	assert.EqualError(s.init(), "error")
}

func TestStore_userError(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockOrg := orgcache.NewMockCache(ctrl)
	mockTicket := ticketcache.NewMockCache(ctrl)
	mockUser := usercache.NewMockCache(ctrl)
	s := &store{mockOrg, mockUser, mockTicket}
	gomock.InOrder(
		mockOrg.EXPECT().Optimise().Times(1),
		mockUser.EXPECT().Optimise().Return(errors.New("error")).Times(1),
		mockTicket.EXPECT().Optimise().Times(0),
	)
	assert.EqualError(s.init(), "error")
}

func TestStore_ticketError(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockOrg := orgcache.NewMockCache(ctrl)
	mockTicket := ticketcache.NewMockCache(ctrl)
	mockUser := usercache.NewMockCache(ctrl)
	s := &store{mockOrg, mockUser, mockTicket}
	gomock.InOrder(
		mockOrg.EXPECT().Optimise().Times(1),
		mockUser.EXPECT().Optimise().Times(1),
		mockTicket.EXPECT().Optimise().Return(errors.New("error")).Times(1),
	)
	assert.EqualError(s.init(), "error")
}

func TestStore_initSuccessful(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockOrg := orgcache.NewMockCache(ctrl)
	mockTicket := ticketcache.NewMockCache(ctrl)
	mockUser := usercache.NewMockCache(ctrl)
	s := &store{mockOrg, mockUser, mockTicket}
	gomock.InOrder(
		mockOrg.EXPECT().Optimise().Times(1),
		mockUser.EXPECT().Optimise().Times(1),
		mockTicket.EXPECT().Optimise().Times(1),
	)
	assert.Nil(s.init())
}
