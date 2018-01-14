package prose

import (
	"regexp"
)

type Entity struct {
	Name      string
	IsPresent *regexp.Regexp
}

func NewEntity(name string, regex *regexp.Regexp) *Entity {
	return &Entity{name, regex}
}
