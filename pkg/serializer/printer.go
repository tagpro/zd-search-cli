package serializer

import (
	"github.com/fatih/color"
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
	color.Red(title)
	for _, data := range kvs {
		color.Cyan("%-20s | %s", data.key, data.value)
	}
}
