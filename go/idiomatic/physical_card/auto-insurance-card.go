package physical_card

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/contact"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/insurance"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/location/state"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/vehicle"
	"github.com/boundedinfinity/rfc3339date"
)

type AutoInsuranceCard struct {
	Id             ider.Id                       `json:"id,omitempty"`
	State          state.State                   `json:"state,omitempty"`
	Insured        []contact.Contact             `json:"insured,omitempty"`
	PolicyNumber   string                        `json:"policy-number,omitempty"`
	Issuer         insurance.AutoInsurancePolicy `json:"issuer,omitempty"`
	Vehicle        vehicle.Vehicle               `json:"vehicle,omitempty"`
	OriginDate     rfc3339date.Rfc3339Date       `json:"origin-date,omitempty"`
	IssueDate      rfc3339date.Rfc3339Date       `json:"issue-date,omitempty"`
	ExpirationDate rfc3339date.Rfc3339Date       `json:"expiration-date,omitempty"`
}
