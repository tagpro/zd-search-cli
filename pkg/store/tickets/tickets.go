package tickets

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"
)

type Ticket struct {
	ID             string        `json:"_id"`
	URL            string        `json:"url"`
	ExternalID     string        `json:"external_id"`
	CreatedAt      jsontime.Time `json:"created_at"`
	Type           string        `json:"type"`
	Subject        string        `json:"subject"`
	Description    string        `json:"description"`
	Priority       string        `json:"priority"`
	Status         string        `json:"status"`
	SubmitterID    int           `json:"submitter_id"`
	AssigneeID     int           `json:"assignee_id"`
	OrganizationID int           `json:"organization_id"`
	Tags           []string      `json:"tags"`
	HasIncidents   bool          `json:"has_incidents"`
	DueAt          jsontime.Time `json:"due_at"`
	Via            string        `json:"via"`
}

type Tickets []*Ticket

func LoadTickets(path string) (Cache, error) {
	var tickets Tickets

	jsonFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON file: %w", err)
	}
	err = json.Unmarshal(jsonFile, &tickets)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON file: %w", err)
	}
	return newCache(tickets), nil
}
