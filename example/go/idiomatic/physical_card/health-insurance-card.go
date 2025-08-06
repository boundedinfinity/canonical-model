package physical_card

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/insurance"
	"github.com/boundedinfinity/schema/idiomatic/location"
)

type HealthInsuranceCard struct {
	Id             id.Id                         `json:"id,omitempty"`
	State          location.State                `json:"state,omitempty"`
	Insured        []contact.Contact             `json:"insured,omitempty"`
	InsuranceId    string                        `json:"insurance-id,omitempty"`
	GroupNumber    string                        `json:"group-number,omitempty"`
	Issuer         insurance.AutoInsurancePolicy `json:"issuer,omitempty"`
	IssueDate      rfc3339date.Rfc3339Date       `json:"issue-date,omitempty"`
	ExpirationDate rfc3339date.Rfc3339Date       `json:"expiration-date,omitempty"`
}

var _ id.TypeNamer = &HealthInsuranceCard{}

func (t HealthInsuranceCard) TypeName() string {
	return id.TypeNamers.Dotted(HealthInsuranceCard{})
}
