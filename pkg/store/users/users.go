package users

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tagpro/zd-search-cli/pkg/jsontime"
)

type User struct {
	ID             int           `json:"_id"`
	URL            string        `json:"url"`
	ExternalID     string        `json:"external_id"`
	Name           string        `json:"name"`
	Alias          string        `json:"alias"`
	CreatedAt      jsontime.Time `json:"created_at"`
	Active         bool          `json:"active"`
	Verified       bool          `json:"verified"`
	Shared         bool          `json:"shared"`
	Locale         string        `json:"locale"`
	Timezone       string        `json:"timezone"`
	LastLoginAt    jsontime.Time `json:"last_login_at"`
	Email          string        `json:"email"`
	Phone          string        `json:"phone"`
	Signature      string        `json:"signature"`
	OrganizationID int           `json:"organization_id"`
	Tags           []string      `json:"tags"`
	Suspended      bool          `json:"suspended"`
	Role           string        `json:"role"`
}

type Users []*User

// LoadUsers loads a JSON file into an unoptimised cache and returns it.
// It is advised to call `.Optimise()` on the returned cache to use the cache
func LoadUsers(path string) (Cache, error) {
	var users Users

	jsonFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON file: %w", err)
	}
	err = json.Unmarshal(jsonFile, &users)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON file: %w", err)
	}
	return newCache(users), nil
}
