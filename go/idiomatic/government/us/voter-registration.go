package us

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/contact"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/rfc3339date"
)

type VoterRegistration struct {
	Id                ider.Id                 `json:"id,omitempty"`
	VoterId           string                  `json:"voter-id,omitempty"`
	Precinct          string                  `json:"precinct,omitempty"`
	IssuedBy          contact.Contact         `json:"issued-by,omitempty"`
	IssuedOn          rfc3339date.Rfc3339Date `json:"issued-on,omitempty"`
	IssuedTo          contact.Contact         `json:"issued-to,omitempty"`
	PollingPlace      contact.Contact         `json:"polling-place,omitempty"`
	ElectionDistricts []string                `json:"election-districts,omitempty"`
}
