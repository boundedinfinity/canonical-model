package usa

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type VoterRegistration struct {
	Id                id.Id                   `json:"id,omitempty"`
	VoterId           string                  `json:"voter-id,omitempty"`
	Precinct          string                  `json:"precinct,omitempty"`
	IssuedBy          contact.Contact         `json:"issued-by,omitempty"`
	IssuedOn          rfc3339date.Rfc3339Date `json:"issued-on,omitempty"`
	IssuedTo          contact.Contact         `json:"issued-to,omitempty"`
	PollingPlace      contact.Contact         `json:"polling-place,omitempty"`
	ElectionDistricts []string                `json:"election-districts,omitempty"`
}
