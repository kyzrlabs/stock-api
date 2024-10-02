package datasource

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Datasource interface {
	// Create adds a new JSON entry.
	Create(key string, value json.RawMessage) error

	// Read retrieves a JSON entry by its key.
	Read(id uuid.UUID) (json.RawMessage, error)

	// Update modifies an existing JSON entry.
	Update(key string, value json.RawMessage) error

	// Delete removes a JSON entry by its key.
	Delete(key string) error

	// List retrieves all JSON entries.
	List() (json.RawMessage, error)
}
