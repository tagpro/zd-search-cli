package organistations

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"
)

type Organisation struct {
	ID            int           `json:"_id"`
	URL           string        `json:"url"`
	ExternalID    string        `json:"external_id"`
	Name          string        `json:"name"`
	DomainNames   []string      `json:"domain_names"`
	CreatedAt     jsontime.Time `json:"created_at"`
	Details       string        `json:"details"`
	SharedTickets bool          `json:"shared_tickets"`
	Tags          []string      `json:"tags"`
}

type Organisations []*Organisation

// LoadOrganisations loads a JSON file into an unoptimised cache and returns it.
// It is advised to call `.Optimise()` on the returned cache to use the cache
func LoadOrganisations(path string) (Cache, error) {
	var organisations Organisations

	jsonFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON file: %w", err)
	}
	err = json.Unmarshal(jsonFile, &organisations)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON file: %w", err)
	}
	return newCache(organisations), nil
}
