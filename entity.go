package prose

import (
	"regexp"
)

// Entity represents an entity
type Entity struct {
	Name      string
	IsPresent *regexp.Regexp
}

// NewEntity returns a new entity with name and regexp in params
// The regexp will be used to know if the entity is present or not
// in a string
func NewEntity(name string, regex *regexp.Regexp) *Entity {
	return &Entity{name, regex}
}
