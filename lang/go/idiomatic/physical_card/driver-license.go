package physical_card

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/location"
)

type DriverLicense struct {
	Id             id.Id                   `json:"id,omitempty"`
	State          location.State          `json:"state,omitempty"`
	Person         contact.Contact         `json:"person,omitempty"`
	Number         string                  `json:"number,omitempty"`
	Class          string                  `json:"class,omitempty"`
	Donor          DonorStatus             `json:"donor,omitempty"`
	IssueDate      rfc3339date.Rfc3339Date `json:"issue-date,omitempty"`
	ExpirationDate rfc3339date.Rfc3339Date `json:"expiration-date,omitempty"`
}
