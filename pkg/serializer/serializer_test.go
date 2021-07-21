package serializer

import (
	"errors"
	"testing"

	ticketstore "github.com/tagpro/zd-search-cli/pkg/store/tickets"

	userstore "github.com/tagpro/zd-search-cli/pkg/store/users"

	orgstore "github.com/tagpro/zd-search-cli/pkg/store/organistations"

	"github.com/golang/mock/gomock"
	tassert "github.com/stretchr/testify/assert"
	store "github.com/tagpro/zd-search-cli/pkg/store/testdata/mocks"
)

func TestSearchEntity_failNoOrganisation(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)

	gomock.InOrder(
		mockStore.EXPECT().
			GetOrganisations("key", "value").
			Return(nil, errors.New("error")),
	)
	s := &serializer{store: mockStore}
	assert.EqualError(s.SearchEntity(SearchCriteria{Organisations, "key", "value"}), "error")
}

func TestSearchEntity_failPrintOrganisation(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)

	gomock.InOrder(
		mockStore.EXPECT().
			GetOrganisations("key", "value").
			Return(orgstore.Organisations{{}}, nil),
		mockStore.EXPECT().
			GetUsers("organization_id", "0").
			Return(nil, errors.New("printError")),
	)
	s := &serializer{store: mockStore}
	assert.EqualError(s.SearchEntity(SearchCriteria{Organisations, "key", "value"}), "printError")
}

func TestSearchEntity_successfulPrintOrganisation(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)

	gomock.InOrder(
		mockStore.EXPECT().
			GetOrganisations("key", "value").
			Return(orgstore.Organisations{}, nil),
	)
	s := &serializer{store: mockStore}
	assert.Nil(s.SearchEntity(SearchCriteria{Organisations, "key", "value"}))
}

func TestSearchEntity_failNoUser(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)

	gomock.InOrder(
		mockStore.EXPECT().
			GetUsers("key", "value").
			Return(nil, errors.New("error")),
	)
	s := &serializer{store: mockStore}
	assert.EqualError(s.SearchEntity(SearchCriteria{Users, "key", "value"}), "error")
}

func TestSearchEntity_failPrintUser(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)

	gomock.InOrder(
		mockStore.EXPECT().
			GetUsers("key", "value").
			Return(userstore.Users{{}}, nil),
		mockStore.EXPECT().
			GetOrganisations("_id", "0").
			Return(nil, errors.New("printError")),
	)
	s := &serializer{store: mockStore}
	assert.EqualError(s.SearchEntity(SearchCriteria{Users, "key", "value"}), "printError")
}

func TestSearchEntity_successfulPrintUser(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)

	gomock.InOrder(
		mockStore.EXPECT().
			GetUsers("key", "value").
			Return(userstore.Users{}, nil),
	)
	s := &serializer{store: mockStore}
	assert.Nil(s.SearchEntity(SearchCriteria{Users, "key", "value"}))
}

func TestSearchEntity_failNoTicket(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)

	gomock.InOrder(
		mockStore.EXPECT().
			GetTickets("key", "value").
			Return(nil, errors.New("error")),
	)
	s := &serializer{store: mockStore}
	assert.EqualError(s.SearchEntity(SearchCriteria{Tickets, "key", "value"}), "error")
}

func TestSearchEntity_failPrintTicket(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)

	gomock.InOrder(
		mockStore.EXPECT().
			GetTickets("key", "value").
			Return(ticketstore.Tickets{{}}, nil),
		mockStore.EXPECT().
			GetUsers("_id", "0").
			Return(nil, errors.New("printError")),
	)
	s := &serializer{store: mockStore}
	assert.EqualError(s.SearchEntity(SearchCriteria{Tickets, "key", "value"}), "printError")
}

func TestSearchEntity_successfulPrintTicket(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)

	gomock.InOrder(
		mockStore.EXPECT().
			GetTickets("key", "value").
			Return(ticketstore.Tickets{}, nil),
	)
	s := &serializer{store: mockStore}
	assert.Nil(s.SearchEntity(SearchCriteria{Tickets, "key", "value"}))
}

func TestNewSerializer(t *testing.T) {
	assert := tassert.New(t)
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)

	s := NewSerializer(mockStore)
	assert.Equal(s, &serializer{mockStore})
}

func TestGetEntities(t *testing.T) {
	tassert.Equal(t, []string{"Users", "Tickets", "Organisations"}, GetEntities())
}

func TestToEntity(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  Entity
		err   string
	}{
		{
			name:  "failed to return for wrong value",
			input: "foo",
			want:  "",
			err:   "invalid entity",
		},
		{
			name:  "failed to return for empty value",
			input: "",
			want:  "",
			err:   "invalid entity",
		},
		{
			name:  "successfully converts Users",
			input: "Users",
			want:  "Users",
			err:   "",
		},
		{
			name:  "successfully converts Tickets",
			input: "Tickets",
			want:  "Tickets",
			err:   "",
		},
		{
			name:  "successfully converts Organisations",
			input: "Organisations",
			want:  "Organisations",
			err:   "",
		},
	}
	for _, tcase := range cases {
		t.Run(tcase.name, func(t *testing.T) {
			assert := tassert.New(t)
			got, gotErr := ToEntity(tcase.input)
			if tcase.err != "" {
				assert.EqualError(gotErr, tcase.err)
			}
			assert.Equal(tcase.want, got)
		})
	}
}
