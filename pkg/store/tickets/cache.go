package tickets

//go:generate mockgen -source=cache.go -destination=testdata/mocks/cache.go -package=mocks . Cache

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"

	"github.com/tagpro/zd-search-cli/pkg/zerror"
)

const (
	Id             = "_id"
	Url            = "url"
	ExternalId     = "external_id"
	CreatedAt      = "created_at"
	Type           = "type"
	Subject        = "subject"
	Description    = "description"
	Priority       = "priority"
	Status         = "status"
	SubmitterId    = "submitter_id"
	AssigneeId     = "assignee_id"
	OrganizationId = "organization_id"
	Tags           = "tags"
	HasIncidents   = "has_incidents"
	DueAt          = "due_at"
	Via            = "via"
)

func GetKeys() []string {
	return []string{
		Id,
		Url,
		ExternalId,
		CreatedAt,
		Type,
		Subject,
		Description,
		Priority,
		Status,
		SubmitterId,
		AssigneeId,
		OrganizationId,
		Tags,
		HasIncidents,
		DueAt,
		Via,
	}
}

type Cache interface {
	GetTickets(key, input string) (Tickets, error)
	Optimise() error
}

type cache struct {
	tickets Tickets
	data    map[string]map[string]Tickets
}

func (c *cache) GetTickets(key, input string) (Tickets, error) {
	if _, ok := c.data[key]; !ok {
		return nil, fmt.Errorf("invalid field name %s", key)
	}
	if tickets, ok := c.data[key][input]; !ok {
		return nil, fmt.Errorf("%w", zerror.ErrNotFound)
	} else {
		return tickets, nil
	}
}

func (c *cache) Optimise() error {
	if c.tickets == nil || c.data == nil {
		return errors.New("cache not initialised")
	}
	for _, ticket := range c.tickets {
		if err := c.addTicket(ticket); err != nil {
			return fmt.Errorf("failed to load all tickets: %w", err)
		}
	}
	return nil
}
func (c *cache) addTicket(ticket *Ticket) error {
	if c.data == nil {
		return fmt.Errorf("cache data not initialised")
	}
	//insert _id
	if _, ok := c.data[Id]; !ok {
		c.data[Id] = map[string]Tickets{}
	}
	c.data[Id][ticket.Id] = append(c.data[Id][ticket.Id], ticket)

	//insert url
	if _, ok := c.data[Url]; !ok {
		c.data[Url] = map[string]Tickets{}
	}
	c.data[Url][ticket.Url] = append(c.data[Url][ticket.Url], ticket)
	//insert external_id
	if _, ok := c.data[ExternalId]; !ok {
		c.data[ExternalId] = map[string]Tickets{}
	}
	c.data[ExternalId][ticket.ExternalId] = append(c.data[ExternalId][ticket.ExternalId], ticket)
	//insert created_at
	if _, ok := c.data[CreatedAt]; !ok {
		c.data[CreatedAt] = map[string]Tickets{}
	}
	c.data[CreatedAt][ticket.CreatedAt.Format(jsontime.ZDTimeFormat)] = append(c.data[CreatedAt][ticket.CreatedAt.Format(jsontime.ZDTimeFormat)], ticket)
	//insert type
	if _, ok := c.data[Type]; !ok {
		c.data[Type] = map[string]Tickets{}
	}
	c.data[Type][ticket.Type] = append(c.data[Type][ticket.Type], ticket)
	//insert subject
	if _, ok := c.data[Subject]; !ok {
		c.data[Subject] = map[string]Tickets{}
	}
	c.data[Subject][ticket.Subject] = append(c.data[Subject][ticket.Subject], ticket)
	//insert description
	if _, ok := c.data[Description]; !ok {
		c.data[Description] = map[string]Tickets{}
	}
	c.data[Description][ticket.Description] = append(c.data[Description][ticket.Description], ticket)
	//insert priority
	if _, ok := c.data[Priority]; !ok {
		c.data[Priority] = map[string]Tickets{}
	}
	c.data[Priority][ticket.Priority] = append(c.data[Priority][ticket.Priority], ticket)
	//insert status
	if _, ok := c.data[Status]; !ok {
		c.data[Status] = map[string]Tickets{}
	}
	c.data[Status][ticket.Status] = append(c.data[Status][ticket.Status], ticket)
	//insert submitter_id
	if _, ok := c.data[SubmitterId]; !ok {
		c.data[SubmitterId] = map[string]Tickets{}
	}
	c.data[SubmitterId][strconv.Itoa(ticket.SubmitterId)] = append(c.data[SubmitterId][strconv.Itoa(ticket.SubmitterId)], ticket)
	//insert assignee_id
	if _, ok := c.data[AssigneeId]; !ok {
		c.data[AssigneeId] = map[string]Tickets{}
	}
	c.data[AssigneeId][strconv.Itoa(ticket.AssigneeId)] = append(c.data[AssigneeId][strconv.Itoa(ticket.AssigneeId)], ticket)
	//insert organization_id
	if _, ok := c.data[OrganizationId]; !ok {
		c.data[OrganizationId] = map[string]Tickets{}
	}
	c.data[OrganizationId][strconv.Itoa(ticket.OrganizationId)] = append(c.data[OrganizationId][strconv.Itoa(ticket.OrganizationId)], ticket)
	//insert tags
	if _, ok := c.data[Tags]; !ok {
		c.data[Tags] = map[string]Tickets{}
	}
	for _, tag := range ticket.Tags {
		c.data[Tags][tag] = append(c.data[Tags][tag], ticket)
	}
	//insert has_incidents
	if _, ok := c.data[HasIncidents]; !ok {
		c.data[HasIncidents] = map[string]Tickets{}
	}
	c.data[HasIncidents][strconv.FormatBool(ticket.HasIncidents)] = append(c.data[HasIncidents][strconv.FormatBool(ticket.HasIncidents)], ticket)
	//insert due_at
	if _, ok := c.data[DueAt]; !ok {
		c.data[DueAt] = map[string]Tickets{}
	}
	c.data[DueAt][ticket.DueAt.Format(jsontime.ZDTimeFormat)] = append(c.data[DueAt][ticket.DueAt.Format(jsontime.ZDTimeFormat)], ticket)
	//insert via
	if _, ok := c.data[Via]; !ok {
		c.data[Via] = map[string]Tickets{}
	}
	c.data[Via][ticket.Via] = append(c.data[Via][ticket.Via], ticket)
	return nil
}

func newCache(tickets Tickets) *cache {
	return &cache{tickets: tickets, data: map[string]map[string]Tickets{}}
}
