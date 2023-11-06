package physical_card

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/audit"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/insurance"
	"github.com/boundedinfinity/schema/idiomatic/location"
	"github.com/boundedinfinity/schema/idiomatic/vehicle"
)

type AutoInsuranceCard struct {
	Id             id.Id                      `json:"id,omitempty"`
	State          location.State             `json:"state,omitempty"`
	Insured        []contact.Contact          `json:"insured,omitempty"`
	PolicyNumber   string                     `json:"policy-number,omitempty"`
	Issuer         insurance.InsuranceCompany `json:"issuer,omitempty"`
	Vehicle        vehicle.Vehicle            `json:"vehicle,omitempty"`
	OriginDate     rfc3339date.Rfc3339Date    `json:"origin-date,omitempty"`
	IssueDate      rfc3339date.Rfc3339Date    `json:"issue-date,omitempty"`
	ExpirationDate rfc3339date.Rfc3339Date    `json:"expiration-date,omitempty"`
	Audit          audit.Audit                `json:"audit,omitempty"`
}
