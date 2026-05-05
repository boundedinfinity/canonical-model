package physical_card

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/contact"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/identifier"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/location/state"
	"github.com/boundedinfinity/rfc3339date"
)

type DriverLicense struct {
	Id     ider.Id         `json:"id,omitempty"`
	State  state.State     `json:"state,omitempty"`
	Person contact.Contact `json:"person,omitempty"`
	Number string          `json:"number,omitempty"`
	Class  string          `json:"class,omitempty"`
	// Donor          DonorStatus             `json:"donor,omitempty"`
	IssueDate      rfc3339date.Rfc3339Date `json:"issue-date,omitempty"`
	ExpirationDate rfc3339date.Rfc3339Date `json:"expiration-date,omitempty"`
}

var _ identifier.TypeNamer = &DriverLicense{}

func (t DriverLicense) TypeName() string {
	return identifier.TypeNamers.Dotted(DriverLicense{})
}
