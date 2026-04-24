package fiction

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/model/person/name"
)

type Charactor struct {
	Id      ider.Id     `json:"id"`
	Name    name.Name   `json:"legal-name"`
	Aliases []name.Name `json:"aliases"`
}
