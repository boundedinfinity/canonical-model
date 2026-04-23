package physical_card

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/contact"
	"github.com/boundedinfinity/canonical_model/idiomatic/id"
	"github.com/boundedinfinity/canonical_model/idiomatic/insurance"
	"github.com/boundedinfinity/canonical_model/idiomatic/location"
	"github.com/boundedinfinity/canonical_model/idiomatic/vehicle"
	"github.com/boundedinfinity/rfc3339date"
)

type AutoInsuranceCard struct {
	Id             id.Id                         `json:"id,omitempty"`
	State          location.State                `json:"state,omitempty"`
	Insured        []contact.Contact             `json:"insured,omitempty"`
	PolicyNumber   string                        `json:"policy-number,omitempty"`
	Issuer         insurance.AutoInsurancePolicy `json:"issuer,omitempty"`
	Vehicle        vehicle.Vehicle               `json:"vehicle,omitempty"`
	OriginDate     rfc3339date.Rfc3339Date       `json:"origin-date,omitempty"`
	IssueDate      rfc3339date.Rfc3339Date       `json:"issue-date,omitempty"`
	ExpirationDate rfc3339date.Rfc3339Date       `json:"expiration-date,omitempty"`
}
