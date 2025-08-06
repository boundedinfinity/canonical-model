package usa

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type GovernmentRepresentation struct {
	Id      id.Id                      `json:"id,omitempty"`
	Federal []GovernmentRepresentative `json:"federal,omitempty"`
	State   []GovernmentRepresentative `json:"state,omitempty"`
}

type GovernmentRepresentative struct {
	Id        id.Id                   `json:"id,omitempty"`
	Office    string                  `json:"office,omitempty"`
	District  string                  `json:"district,omitempty"`
	Contact   contact.Contact         `json:"contact,omitempty"`
	TermStart rfc3339date.Rfc3339Date `json:"term-start,omitempty"`
	TermEnd   contact.Contact         `json:"term-end,omitempty"`
}
