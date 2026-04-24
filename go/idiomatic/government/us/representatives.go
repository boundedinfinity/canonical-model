package us

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/contact"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/rfc3339date"
)

type GovernmentRepresentation struct {
	Id      ider.Id                    `json:"id,omitempty"`
	Federal []GovernmentRepresentative `json:"federal,omitempty"`
	State   []GovernmentRepresentative `json:"state,omitempty"`
}

type GovernmentRepresentative struct {
	Id        ider.Id                 `json:"id,omitempty"`
	Office    string                  `json:"office,omitempty"`
	District  string                  `json:"district,omitempty"`
	Contact   contact.Contact         `json:"contact,omitempty"`
	TermStart rfc3339date.Rfc3339Date `json:"term-start,omitempty"`
	TermEnd   rfc3339date.Rfc3339Date `json:"term-end,omitempty"`
}
