package physical_card

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/contact"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/id"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/insurance"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/location"
	"github.com/boundedinfinity/rfc3339date"
)

type HealthInsuranceCard struct {
	Id             ider.Id                       `json:"id,omitempty"`
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
