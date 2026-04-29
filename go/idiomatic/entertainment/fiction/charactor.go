package fiction

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person/name"
)

type Character struct {
	Id      ider.Id     `json:"id"`
	Name    name.Name   `json:"legal-name"`
	Aliases []name.Name `json:"aliases"`
}
