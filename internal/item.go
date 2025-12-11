// Package entities holds all the entities that are shared across subdomains.
package tavern

import "github.com/google/uuid"

// Item is an entity that represents an item in all domains.
type Item struct {
	// ID an identifier of the entity
	ID          uuid.UUID
	Name        string
	Description string
}
